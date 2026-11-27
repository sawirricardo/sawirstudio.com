package site

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"io/fs"
	"os"
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
		if p.Related && related(projects, slugOf(p.Path)) == nil {
			return nil, fmt.Errorf("%s: related: true, but it isn't in content/projects.yaml", p.Path)
		}
		var buf bytes.Buffer
		d := pageData{Page: p, OGImageW: w, OGImageH: h, Projects: projects, Clients: clients, Opts: opts}
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
	"<!-- @project-cards -->": func(d pageData) templ.Component { return projectCards(d.Projects) },
	"<!-- @inquiry-form -->":  func(d pageData) templ.Component { return inquiryForm(d.Opts.InquiryAPI) },
	"<!-- @client-list -->":   func(d pageData) templ.Component { return clientList(d.Clients, d.Projects) },
	"<!-- @client-logos -->":  func(d pageData) templ.Component { return clientLogos(d.Clients, d.Projects) },
	"<!-- @client-page -->": func(d pageData) templ.Component {
		c := clientBySlug(d.Clients, d.Page.Client)
		return clientPage(c, clientProjects(d.Projects, c.Name))
	},
}

// expand is the page's <main> HTML with its placeholders rendered.
func expand(d pageData) string {
	s := d.Page.Main
	for marker, c := range placeholders {
		if !strings.Contains(s, marker) {
			continue
		}
		var buf bytes.Buffer
		if err := c(d).Render(context.Background(), &buf); err != nil {
			panic(err)
		}
		s = strings.ReplaceAll(s, marker, buf.String())
	}
	return s
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
