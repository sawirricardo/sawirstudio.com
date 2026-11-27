// Package site builds the English pages of sawirstudio.com from content/.
//
// Each page is one file under content/ (content/projects/rindraw.html is /projects/rindraw/,
// content/home.html is /, content/404.html is /404.html): YAML front matter with the page's
// title, description and share-card details, then the HTML that goes inside <main>.
// The layout, header, footer, project cards, "More projects" rows and the closing call to
// action come from the templ components in this package.
package site

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

const SiteURL = "https://sawirstudio.com"

// Page is one English page.
type Page struct {
	Path string `yaml:"-"` // URL path: "/", "/projects/rindraw/", or "/404.html"
	Main string `yaml:"-"` // the HTML inside <main>

	Title         string `yaml:"title"`
	Description   string `yaml:"description"`
	OGTitle       string `yaml:"og_title"`       // defaults to Title
	OGDescription string `yaml:"og_description"` // defaults to Description
	OGType        string `yaml:"og_type"`
	// OGImage is the share picture, like /assets/og/rindraw.jpg. Without it the page uses
	// its card drawn by the builder (internal/ogcard), /assets/og/cards/<name>.jpg.
	OGImage     string `yaml:"og_image"`
	TwitterCard string `yaml:"twitter_card"` // defaults to summary_large_image
	Robots      string `yaml:"robots"`
	ConvertKit  bool   `yaml:"convertkit"` // the newsletter form's styles and script (blog posts)
	LD          string `yaml:"ld"`         // JSON-LD, as written
	CTA         *bool  `yaml:"cta"`        // the "How can we help you?" box; on unless false
	Related     bool   `yaml:"related"`    // the "More projects" row (project pages)
}

func (p *Page) ogTitle() string {
	if p.OGTitle != "" {
		return p.OGTitle
	}
	return p.Title
}

func (p *Page) ogDescription() string {
	if p.OGDescription != "" {
		return p.OGDescription
	}
	return p.Description
}

func (p *Page) twitterCard() string {
	if p.TwitterCard != "" {
		return p.TwitterCard
	}
	return "summary_large_image"
}

func (p *Page) showCTA() bool { return p.CTA == nil || *p.CTA }

// IsPage reports whether the page is a directory page (everything but 404.html).
func (p *Page) IsPage() bool { return strings.HasSuffix(p.Path, "/") }

// cardName is the file name of the page's drawn share card.
func (p *Page) cardName() string {
	if !p.IsPage() || p.Path == "/" {
		return "home"
	}
	return strings.ReplaceAll(strings.Trim(p.Path, "/"), "/", "-")
}

// ogImagePath is the share picture's path from the site root.
func (p *Page) ogImagePath() string {
	if p.OGImage != "" {
		return p.OGImage
	}
	return "/assets/og/cards/" + p.cardName() + ".jpg"
}

// OutFile is where the built page goes, relative to the site root.
func (p *Page) OutFile() string {
	if !p.IsPage() {
		return strings.TrimPrefix(p.Path, "/")
	}
	return filepath.FromSlash(strings.TrimPrefix(p.Path, "/") + "index.html")
}

// LoadPages reads every page under dir.
func LoadPages(dir string) ([]*Page, error) {
	var pages []*Page
	dir, err := filepath.EvalSymlinks(dir)
	if err != nil {
		return nil, err
	}
	err = filepath.WalkDir(dir, func(file string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || filepath.Ext(file) != ".html" {
			return err
		}
		rel, _ := filepath.Rel(dir, file)
		p, err := loadPage(file, filepath.ToSlash(strings.TrimSuffix(rel, ".html")))
		if err != nil {
			return fmt.Errorf("%s: %w", file, err)
		}
		pages = append(pages, p)
		return nil
	})
	sort.Slice(pages, func(i, j int) bool { return pages[i].Path < pages[j].Path })
	return pages, err
}

func loadPage(file, rel string) (*Page, error) {
	src, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	front, body, ok := bytes.Cut(bytes.TrimPrefix(src, []byte("---\n")), []byte("\n---\n"))
	if !ok || !bytes.HasPrefix(src, []byte("---\n")) {
		return nil, fmt.Errorf("no front matter (the file must start with a --- block)")
	}
	p := &Page{Main: string(body)}
	dec := yaml.NewDecoder(bytes.NewReader(front))
	dec.KnownFields(true)
	if err := dec.Decode(p); err != nil {
		return nil, err
	}
	switch rel {
	case "home":
		p.Path = "/"
	case "404":
		p.Path = "/404.html"
	default:
		p.Path = "/" + rel + "/"
	}
	if p.Title == "" || p.Description == "" || p.OGType == "" {
		return nil, fmt.Errorf("title, description and og_type are required")
	}
	return p, nil
}

// Project is one card in the project list (content/projects.yaml).
type Project struct {
	Slug       string `yaml:"slug"`
	Title      string `yaml:"title"`
	Categories []int  `yaml:"categories"` // 1 Design, 2 Website, 3 Application
	Search     string `yaml:"search"`
	Image      string `yaml:"image"`
	Width      int    `yaml:"width"`
	Height     int    `yaml:"height"`
	Alt        string `yaml:"alt"` // defaults to Title
	App        bool   `yaml:"app"` // an app icon, shown in a rounded box
}

func (p Project) alt() string {
	if p.Alt != "" {
		return p.Alt
	}
	return p.Title
}

func (p Project) categories() string {
	s := make([]string, len(p.Categories))
	for i, c := range p.Categories {
		s[i] = fmt.Sprint(c)
	}
	return strings.Join(s, ",")
}

func (p Project) shares(q Project) bool {
	for _, a := range p.Categories {
		for _, b := range q.Categories {
			if a == b {
				return true
			}
		}
	}
	return false
}

// LoadProjects reads the project list.
func LoadProjects(file string) ([]Project, error) {
	src, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var ps []Project
	dec := yaml.NewDecoder(bytes.NewReader(src))
	dec.KnownFields(true)
	if err := dec.Decode(&ps); err != nil {
		return nil, fmt.Errorf("%s: %w", file, err)
	}
	return ps, nil
}

// related picks the "More projects" row for a project: the next three in the list that share
// a category with it, wrapping around, topped up with the next ones that don't.
func related(ps []Project, slug string) []Project {
	i := -1
	for j, p := range ps {
		if p.Slug == slug {
			i = j
		}
	}
	if i < 0 {
		return nil
	}
	after := append(append([]Project{}, ps[i+1:]...), ps[:i]...)
	var picks []Project
	picked := map[string]bool{}
	for _, p := range after {
		if len(picks) < 3 && p.shares(ps[i]) {
			picks = append(picks, p)
			picked[p.Slug] = true
		}
	}
	for _, p := range after {
		if len(picks) < 3 && !picked[p.Slug] {
			picks = append(picks, p)
		}
	}
	return picks
}
