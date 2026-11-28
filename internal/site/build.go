package site

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/a-h/templ"
	"github.com/sawirricardo/sawirstudio.com/internal/ogcard"
	"gopkg.in/yaml.v3"
)

// Build renders every page in root/content into root, next to the assets, draws their share
// cards (assets/og/cards/), and returns the files it wrote or removed (relative to root).
func Build(root string) ([]string, error) { return build(root, true) }

// Check returns the pages Build would change, without writing them.
func Check(root string) ([]string, error) { return build(root, false) }

func build(root string, write bool) ([]string, error) {
	opts, err := loadOptions(filepath.Join(root, "content", "site.yaml"))
	if err != nil {
		return nil, err
	}
	pages, err := LoadPages(filepath.Join(root, "content"))
	if err != nil {
		return nil, err
	}
	projects, err := LoadProjects(filepath.Join(root, "content", "projects.yaml"))
	if err != nil {
		return nil, err
	}
	clients, err := LoadClients(filepath.Join(root, "content", "clients.yaml"), projects)
	if err != nil {
		return nil, err
	}
	pages = append(pages, clientPages(clients, projects)...)
	versions := map[string]string{}
	for key, file := range map[string]string{"css": "assets/css/style.css", "js": "assets/js/main.js"} {
		if versions[key], err = fileVersion(filepath.Join(root, file)); err != nil {
			return nil, err
		}
	}
	written, err := drawCards(root, pages, write)
	if err != nil {
		return nil, fmt.Errorf("share cards: %w", err)
	}
	for _, p := range pages {
		w, h, err := imageSize(filepath.Join(root, filepath.FromSlash(p.ogImagePath())))
		if errors.Is(err, fs.ErrNotExist) && p.OGImage == "" && !write {
			w, h, err = ogcard.W, ogcard.H, nil // the card a real build would draw
		}
		if err != nil {
			return nil, fmt.Errorf("%s: share picture: %w", p.Path, err)
		}
		// Link previews cache the picture by URL, so its URL carries a hash of the picture.
		// (In a check before the card is drawn there's nothing to hash yet; Check compares
		// pages, which a real build then writes with the hash.)
		ogV, _ := fileVersion(filepath.Join(root, filepath.FromSlash(p.ogImagePath())))
		if p.Related && related(projects, slugOf(p.Path)) == nil {
			return nil, fmt.Errorf("%s: related: true, but it isn't in content/projects.yaml", p.Path)
		}
		var share []shareFormat
		if p.isPost() {
			share = shareFormats(root, p, w, h, ogV)
		}
		var buf bytes.Buffer
		d := pageData{Page: p, OGImageW: w, OGImageH: h, OGImageV: ogV, Share: share, Projects: projects, Clients: clients, Opts: opts, Versions: versions}
		if err := Layout(d).Render(context.Background(), &buf); err != nil {
			return nil, fmt.Errorf("%s: %w", p.Path, err)
		}
		buf.WriteByte('\n')
		out := filepath.Join(root, p.OutFile())
		if old, err := os.ReadFile(out); err == nil && bytes.Equal(old, buf.Bytes()) {
			continue
		}
		written = append(written, p.OutFile())
		if !write {
			continue
		}
		if err := os.MkdirAll(filepath.Dir(out), 0o755); err != nil {
			return nil, err
		}
		if err := os.WriteFile(out, buf.Bytes(), 0o644); err != nil {
			return nil, err
		}
	}
	return written, nil
}

// Placeholders a page's HTML can use for parts the builder renders.
var placeholders = map[string]func(d pageData) templ.Component{
	"<!-- @project-cards -->":   func(d pageData) templ.Component { return projectCards(d.Projects) },
	"<!-- @inquiry-form -->":    func(d pageData) templ.Component { return inquiryForm(d.Opts.InquiryAPI, d.Opts.TurnstileSitekey) },
	"<!-- @newsletter-form -->": func(d pageData) templ.Component { return newsletterForm() },
	"<!-- @client-list -->":     func(d pageData) templ.Component { return clientList(d.Clients, d.Projects) },
	"<!-- @client-logos -->":    func(d pageData) templ.Component { return clientLogos(d.Clients, d.Projects) },
	"<!-- @client-page -->": func(d pageData) templ.Component {
		c := clientBySlug(d.Clients, d.Page.Client)
		return clientPage(c, clientProjects(d.Projects, c.Name))
	},
}

// expand is the page's <main> HTML with its placeholders rendered, and on a blog post the
// share button at the end of its <article>.
func expand(d pageData) string {
	s := d.Page.Main
	for marker, c := range placeholders {
		if !strings.Contains(s, marker) {
			continue
		}
		s = strings.ReplaceAll(s, marker, render(c(d)))
	}
	if len(d.Share) > 0 {
		share := render(shareCard(d.Share)) + "\n"
		if i := strings.Index(s, "</article>"); i >= 0 {
			s = s[:i] + share + "        " + s[i:]
		} else {
			s += share
		}
	}
	return s
}

func render(c templ.Component) string {
	var buf bytes.Buffer
	if err := c.Render(context.Background(), &buf); err != nil {
		panic(err)
	}
	return buf.String()
}

// shareFormats are the sizes in a post's share dialog: its pictures in ogcard.Formats, then
// its link card (w×h, content hash ogV). In a check before the pictures are drawn their
// hashes are empty, like the link card's.
func shareFormats(root string, p *Page, w, h int, ogV string) []shareFormat {
	slug := path.Base(strings.Trim(p.Path, "/"))
	labels := map[string][2]string{"story": {"Story", "9:16"}, "feed": {"Feed", "4:5"}, "square": {"Square", "1:1"}}
	var out []shareFormat
	for _, f := range ogcard.Formats {
		src := "/assets/og/cards/" + p.shareCardName(f.Name) + ".jpg"
		v, _ := fileVersion(filepath.Join(root, filepath.FromSlash(src)))
		l := labels[f.Name]
		out = append(out, shareFormat{Key: f.Name, Label: l[0], Ratio: l[1], Src: src + "?v=" + v,
			File: "sawirstudio-" + slug + "-" + f.Name + ".jpg", W: f.W, H: f.H})
	}
	ext := path.Ext(p.ogImagePath())
	return append(out, shareFormat{Key: "link", Label: "Link", Ratio: "1.91:1", Src: p.ogImagePath() + "?v=" + ogV,
		File: "sawirstudio-" + slug + "-link" + ext, W: w, H: h})
}

func slugOf(path string) string {
	return strings.TrimPrefix(strings.Trim(path, "/"), "projects/")
}

func imageSize(file string) (int, int, error) {
	f, err := os.Open(file)
	if err != nil {
		return 0, 0, err
	}
	defer f.Close()
	c, _, err := image.DecodeConfig(f)
	return c.Width, c.Height, err
}

func loadOptions(file string) (Options, error) {
	var o Options
	src, err := os.ReadFile(file)
	if errors.Is(err, fs.ErrNotExist) {
		return o, nil
	}
	if err != nil {
		return o, err
	}
	dec := yaml.NewDecoder(bytes.NewReader(src))
	dec.KnownFields(true)
	if err := dec.Decode(&o); err != nil && !errors.Is(err, io.EOF) {
		return o, fmt.Errorf("%s: %w", file, err)
	}
	return o, nil
}

// fileVersion is the first 10 hex digits of the file's SHA-256.
func fileVersion(file string) (string, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	sum := sha256.Sum256(b)
	return hex.EncodeToString(sum[:])[:10], nil
}
