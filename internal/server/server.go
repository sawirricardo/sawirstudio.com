// Package server serves sawirstudio.com: the built site (embedded in the binary, or a folder
// in development), the same way the Cloudflare worker (worker.js) did, plus the contact
// form's API. Cloudflare sits in front and caches; the Cache-Control headers here say what
// it may keep.
package server

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"math"
	"mime"
	"net/http"
	"path"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Config is how the server runs.
type Config struct {
	Site fs.FS // the built site
	// TrustProxy: take the client's IP and scheme from CF-Connecting-IP, X-Forwarded-For and
	// X-Forwarded-Proto. Set it when the server is only reachable through Cloudflare.
	TrustProxy bool
	Inquiries  *Inquiries // nil: the contact API answers 503
	Log        *slog.Logger
}

// Cache lifetimes. Nothing is fingerprinted (style.css keeps its name), so browsers and
// Cloudflare revalidate every file with its ETag, as with the old Worker: a deploy shows
// up at once, and an unchanged file costs a 304.
const (
	pageCache  = "public, max-age=0, must-revalidate"
	assetCache = "public, max-age=0, must-revalidate"
)

// Content types for the site's files, the same on every host (Go's own table lacks some,
// and would fall back to the machine's /etc/mime.types).
func init() {
	for ext, typ := range map[string]string{
		".html": "text/html; charset=utf-8", ".css": "text/css; charset=utf-8",
		".js": "text/javascript; charset=utf-8", ".md": "text/markdown; charset=utf-8",
		".txt": "text/plain; charset=utf-8", ".xml": "application/xml", ".json": "application/json",
		".svg": "image/svg+xml", ".png": "image/png", ".jpg": "image/jpeg", ".webp": "image/webp",
		".ico": "image/vnd.microsoft.icon", ".zip": "application/zip",
	} {
		mime.AddExtensionType(ext, typ)
	}
}

// New returns the site's handler.
func New(cfg Config) (http.Handler, error) {
	if cfg.Log == nil {
		cfg.Log = slog.New(slog.DiscardHandler)
	}
	if _, err := fs.Stat(cfg.Site, "index.html"); err != nil {
		return nil, errors.New("the site has no index.html (build the binary with make dist, or pass -root)")
	}
	ignore, err := readIgnore(cfg.Site, ".assetsignore")
	if err != nil {
		return nil, err
	}
	redirects, err := readRedirects(cfg.Site, "_redirects")
	if err != nil {
		return nil, err
	}
	s := &static{site: cfg.Site, ignore: ignore, redirects: redirects, etags: map[string]string{}}
	mux := http.NewServeMux()
	mux.Handle("POST /api/inquiries", noStore(inquiryHandler(cfg)))
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		if cfg.Inquiries != nil {
			if err := cfg.Inquiries.Ping(r.Context()); err != nil {
				http.Error(w, "database: "+err.Error(), http.StatusServiceUnavailable)
				return
			}
		}
		fmt.Fprintln(w, "ok")
	})
	mux.Handle("/", s)
	return logRequests(cfg, secureHeaders(httpsOnly(cfg, mux))), nil
}

// static serves the site's files. Like worker.js, it answers Accept: text/markdown with a
// page's index.md, and points to the Markdown copy, llms.txt and the sitemap in Link headers.
type static struct {
	site      fs.FS
	ignore    map[string]bool
	redirects map[string]redirect

	mu    sync.Mutex
	etags map[string]string
}

type redirect struct {
	to   string
	code int
}

func (s *static) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		w.Header().Set("Allow", "GET, HEAD")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	p := path.Clean("/" + r.URL.Path)
	if strings.HasSuffix(r.URL.Path, "/") && p != "/" {
		p += "/"
	}
	if rd, ok := s.redirects[p]; ok {
		w.Header().Set("Cache-Control", assetCache)
		http.Redirect(w, r, rd.to, rd.code)
		return
	}
	if s.hidden(p) {
		s.notFound(w, r)
		return
	}
	name := strings.Trim(p, "/")
	if name == "" {
		name = "."
	}
	info, err := fs.Stat(s.site, name)
	if err != nil {
		s.notFound(w, r)
		return
	}
	if info.IsDir() {
		if !strings.HasSuffix(p, "/") {
			// /about-us -> /about-us/, keeping the query.
			to := p + "/"
			if r.URL.RawQuery != "" {
				to += "?" + r.URL.RawQuery
			}
			w.Header().Set("Cache-Control", assetCache)
			http.Redirect(w, r, to, http.StatusMovedPermanently)
			return
		}
		s.page(w, r, p, name)
		return
	}
	if strings.HasSuffix(name, ".html") {
		w.Header().Set("Cache-Control", pageCache)
	} else {
		w.Header().Set("Cache-Control", assetCache)
	}
	s.serveFile(w, r, name)
}

// page serves a directory page (a URL ending in /).
func (s *static) page(w http.ResponseWriter, r *http.Request, p, dir string) {
	h := w.Header()
	h.Set("Vary", "Accept")
	if acceptsMarkdown(r.Header.Get("Accept")) {
		if md, err := fs.ReadFile(s.site, path.Join(dir, "index.md")); err == nil {
			// Cloudflare doesn't vary its cache on Accept, so the Markdown answer must not be
			// cached there, or browsers would get it.
			h.Set("Cache-Control", "private, no-store")
			h.Set("Content-Type", "text/markdown; charset=utf-8")
			// A rough count (about four characters a token), as Markdown for Agents sends.
			h.Set("x-markdown-tokens", strconv.Itoa(int(math.Ceil(float64(len(md))/4))))
			h.Set("Link", fmt.Sprintf("<%s>; rel=\"canonical\"", p))
			h.Set("Content-Length", strconv.Itoa(len(md)))
			if r.Method != http.MethodHead {
				w.Write(md)
			}
			return
		}
	}
	index := path.Join(dir, "index.html")
	if _, err := fs.Stat(s.site, index); err != nil {
		s.notFound(w, r)
		return
	}
	h.Set("Cache-Control", pageCache)
	h.Add("Link", fmt.Sprintf("<%sindex.md>; rel=\"alternate\"; type=\"text/markdown\"", p))
	if p == "/" {
		h.Add("Link", `</llms.txt>; rel="describedby"; type="text/plain"`)
		h.Add("Link", `</sitemap.xml>; rel="sitemap"; type="application/xml"`)
	}
	h.Set("Content-Type", "text/html; charset=utf-8")
	s.serveFile(w, r, index)
}

// serveFile sends a file with an ETag (embedded files have no modification time), so
// browsers and Cloudflare can revalidate it cheaply.
func (s *static) serveFile(w http.ResponseWriter, r *http.Request, name string) {
	body, err := fs.ReadFile(s.site, name)
	if err != nil {
		s.notFound(w, r)
		return
	}
	w.Header().Set("ETag", s.etag(name, body))
	http.ServeContent(w, r, path.Base(name), time.Time{}, bytes.NewReader(body))
}

func (s *static) etag(name string, body []byte) string {
	s.mu.Lock()
	defer s.mu.Unlock()
	if e, ok := s.etags[name]; ok {
		return e
	}
	sum := sha256.Sum256(body)
	e := `"` + hex.EncodeToString(sum[:8]) + `"`
	s.etags[name] = e
	return e
}

func acceptsMarkdown(accept string) bool {
	for _, part := range strings.Split(accept, ",") {
		mt, _, _ := strings.Cut(strings.TrimSpace(part), ";")
		if strings.EqualFold(strings.TrimSpace(mt), "text/markdown") {
			return true
		}
	}
	return false
}

func (s *static) notFound(w http.ResponseWriter, r *http.Request) {
	// The nearest 404.html up the path, so /ja/nope/ gets the Japanese one.
	dir := path.Dir(strings.TrimSuffix(path.Clean("/"+r.URL.Path), "/") + "/x")
	var body []byte
	for {
		b, err := fs.ReadFile(s.site, strings.TrimPrefix(path.Join(dir, "404.html"), "/"))
		if err == nil {
			body = b
			break
		}
		if dir == "/" || dir == "." {
			break
		}
		dir = path.Dir(dir)
	}
	if body == nil {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Cache-Control", pageCache)
	w.WriteHeader(http.StatusNotFound)
	if r.Method != http.MethodHead {
		w.Write(body)
	}
}

// hidden reports whether a path must not be served: dotfiles and anything named in
// .assetsignore (the source, scripts and config next to the built pages).
func (s *static) hidden(p string) bool {
	for _, seg := range strings.Split(strings.Trim(p, "/"), "/") {
		if strings.HasPrefix(seg, ".") || s.ignore[seg] || seg == "_redirects" || seg == "_headers" {
			return true
		}
	}
	return false
}

// readIgnore reads .assetsignore: one name per line, matched against every path segment.
func readIgnore(site fs.FS, file string) (map[string]bool, error) {
	names := map[string]bool{}
	err := eachLine(site, file, func(line string) error {
		names[strings.Trim(line, "/")] = true
		return nil
	})
	return names, err
}

// readRedirects reads _redirects (Cloudflare's format): "from to [code]" per line.
func readRedirects(site fs.FS, file string) (map[string]redirect, error) {
	rs := map[string]redirect{}
	err := eachLine(site, file, func(line string) error {
		f := strings.Fields(line)
		if len(f) < 2 {
			return fmt.Errorf("%s: bad line %q", file, line)
		}
		code := http.StatusFound
		if len(f) > 2 {
			c, err := strconv.Atoi(f[2])
			if err != nil {
				return fmt.Errorf("%s: bad status in %q", file, line)
			}
			code = c
		}
		rs[f[0]] = redirect{f[1], code}
		return nil
	})
	return rs, err
}

func eachLine(site fs.FS, file string, fn func(string) error) error {
	f, err := site.Open(file)
	if errors.Is(err, fs.ErrNotExist) {
		return nil
	}
	if err != nil {
		return err
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		if err := fn(line); err != nil {
			return err
		}
	}
	return sc.Err()
}

// httpsOnly sends plain-HTTP requests to HTTPS, when a proxy in front says they came in
// over HTTP, and www.sawirstudio.com to sawirstudio.com.
func httpsOnly(cfg Config, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if host, ok := strings.CutPrefix(r.Host, "www."); ok {
			u := *r.URL
			u.Scheme, u.Host = "https", host
			w.Header().Set("Cache-Control", assetCache)
			http.Redirect(w, r, u.String(), http.StatusMovedPermanently)
			return
		}
		if cfg.TrustProxy && r.Header.Get("X-Forwarded-Proto") == "http" {
			u := *r.URL
			u.Scheme, u.Host = "https", r.Host
			http.Redirect(w, r, u.String(), http.StatusMovedPermanently)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func noStore(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		next.ServeHTTP(w, r)
	})
}

func secureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := w.Header()
		h.Set("X-Content-Type-Options", "nosniff")
		h.Set("Referrer-Policy", "strict-origin-when-cross-origin")
		h.Set("Strict-Transport-Security", "max-age=31536000")
		next.ServeHTTP(w, r)
	})
}

type statusWriter struct {
	http.ResponseWriter
	status int
}

func (s *statusWriter) WriteHeader(code int) {
	s.status = code
	s.ResponseWriter.WriteHeader(code)
}

func logRequests(cfg Config, next http.Handler) http.Handler {
	if cfg.Log == nil {
		return next
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		sw := &statusWriter{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(sw, r)
		cfg.Log.Info("request", "method", r.Method, "path", r.URL.Path, "status", sw.status,
			"ms", time.Since(start).Milliseconds(), "ip", clientIP(cfg, r))
	})
}
