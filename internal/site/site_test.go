package site

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRelated(t *testing.T) {
	ps := []Project{
		{Slug: "a", Categories: []int{3}},
		{Slug: "b", Categories: []int{2}},
		{Slug: "c", Categories: []int{3}},
		{Slug: "d", Categories: []int{1}},
		{Slug: "e", Categories: []int{3, 2}},
	}
	got := slugs(related(ps, "c"))
	// Same category first (e, then a, wrapping around), topped up with the next one (d).
	if got != "e a d" {
		t.Errorf("related(c) = %q, want %q", got, "e a d")
	}
	if related(ps, "missing") != nil {
		t.Error("related of an unknown slug should be nil")
	}
}

func TestSameClient(t *testing.T) {
	ps := []Project{
		{Slug: "a", Client: "X", Categories: []int{3}},
		{Slug: "b", Categories: []int{3}},
		{Slug: "c", Client: "X", Categories: []int{2}},
		{Slug: "d", Client: "Y", Categories: []int{3}},
	}
	if got := slugs(sameClient(ps, "a")); got != "c" {
		t.Errorf("sameClient(a) = %q, want %q", got, "c")
	}
	if got := slugs(sameClient(ps, "b")); got != "" {
		t.Errorf("sameClient(b) = %q, want none: our own work has no client", got)
	}
	// The same client's projects are in their own row, so "More projects" skips them.
	if got := slugs(related(ps, "a")); got != "b d" {
		t.Errorf("related(a) = %q, want %q", got, "b d")
	}
}

func slugs(ps []Project) string {
	s := make([]string, len(ps))
	for i, p := range ps {
		s[i] = p.Slug
	}
	return strings.Join(s, " ")
}

func TestNavCurrent(t *testing.T) {
	for _, c := range []struct{ path, href, want string }{
		{"/", "/", "page"},
		{"/projects/", "/projects/", "page"},
		{"/projects/rindraw/", "/projects/", "true"},
		{"/clients/peoplebond/", "/projects/", "true"},
		{"/posts/x/", "/posts/", "true"},
		{"/custom/", "/projects/", ""},
		{"/about-us/", "/about-us/", "page"},
	} {
		if got := navCurrent(c.path, c.href); got != c.want {
			t.Errorf("navCurrent(%q, %q) = %q, want %q", c.path, c.href, got, c.want)
		}
	}
}

// The repo's own content builds, and building twice changes nothing.
func TestBuildRepo(t *testing.T) {
	root := t.TempDir()
	repo, _ := filepath.Abs("../..")
	// content/ and assets/ from the repo, but the cards go to a fresh assets/og/cards/.
	link := func(from, to string) {
		t.Helper()
		if err := os.Symlink(filepath.Join(repo, from), filepath.Join(root, to)); err != nil {
			t.Fatal(err)
		}
	}
	link("content", "content")
	os.MkdirAll(filepath.Join(root, "assets", "og", "cards"), 0o755)
	entries, _ := os.ReadDir(filepath.Join(repo, "assets"))
	for _, e := range entries {
		if e.Name() != "og" {
			link(filepath.Join("assets", e.Name()), filepath.Join("assets", e.Name()))
		}
	}
	ogs, _ := os.ReadDir(filepath.Join(repo, "assets", "og"))
	for _, e := range ogs {
		if !e.IsDir() {
			link(filepath.Join("assets", "og", e.Name()), filepath.Join("assets", "og", e.Name()))
		}
	}
	os.WriteFile(filepath.Join(root, "assets", "og", "cards", "gone.jpg"), nil, 0o644)
	written, err := Build(root)
	if err != nil {
		t.Fatal(err)
	}
	if len(written) < 70 {
		t.Fatalf("built %d pages, want every page", len(written))
	}
	if _, err := os.Stat(filepath.Join(root, "assets/og/cards/posts.jpg")); err != nil {
		t.Error("the blog list's card wasn't drawn")
	}
	if _, err := os.Stat(filepath.Join(root, "assets/og/cards/gone.jpg")); err == nil {
		t.Error("a card no page uses should be removed")
	}
	if _, err := os.Stat(filepath.Join(root, "assets/og/cards/projects-rindraw.jpg")); err == nil {
		t.Error("a page with its own picture shouldn't get a card")
	}
	again, err := Build(root)
	if err != nil || len(again) != 0 {
		t.Fatalf("second build wrote %v (err %v), want nothing", again, err)
	}
	page, _ := os.ReadFile(filepath.Join(root, "projects/rindraw/index.html"))
	for _, want := range []string{
		`<link rel="canonical" href="https://sawirstudio.com/projects/rindraw/">`,
		`<meta property="og:image:width" content="512">`,
		`<section class="section related">`,
		`<a href="/projects/" aria-current="true">Projects</a>`,
		`<!-- i18n:alternates -->`,
		`id="cta"`,
	} {
		if !strings.Contains(string(page), want) {
			t.Errorf("rindraw page is missing %s", want)
		}
	}
	contact, _ := os.ReadFile(filepath.Join(root, "contact/index.html"))
	if !strings.Contains(string(contact), `id="inquiry-form"`) {
		t.Error("contact page has no form")
	}
}

func TestInquiryFormTarget(t *testing.T) {
	d := pageData{Page: &Page{Main: "<!-- @inquiry-form -->"}, Opts: Options{InquiryAPI: "/api/inquiries"}}
	got := expand(d)
	for _, want := range []string{`action="/api/inquiries"`, `name="email"`, `name="website"`, `data-endpoint="api"`} {
		if !strings.Contains(got, want) {
			t.Errorf("API form is missing %s", want)
		}
	}
	d.Opts.InquiryAPI = ""
	if got := expand(d); !strings.Contains(got, "docs.google.com/forms") || strings.Contains(got, `name="website"`) {
		t.Error("default form should post to the Google Form, without the honeypot")
	}
}

func TestCardText(t *testing.T) {
	for _, c := range []struct{ path, title, wantTitle, wantLabel string }{
		{"/posts/x/", "Hello | Sawir Studio", "Hello", "Blog"},
		{"/posts/", "Blog | Sawir Studio", "Blog", ""},
		{"/shopify/", "Shopify | Sawir Studio", "Shopify", "Services"},
		{"/about-us/", "About us | Sawir Studio", "About us", ""},
	} {
		p := &Page{Path: c.path, Title: c.title}
		if title, label := p.cardText(); title != c.wantTitle || label != c.wantLabel {
			t.Errorf("%s: %q %q, want %q %q", c.path, title, label, c.wantTitle, c.wantLabel)
		}
	}
	if title, _ := (&Page{Path: "/404.html", Title: "x"}).cardText(); title == "x" {
		t.Error("404 should use the home card")
	}
}

func TestClientURL(t *testing.T) {
	ps := []Project{{Slug: "a", Client: "X"}, {Slug: "b", Client: "X"}, {Slug: "c", Client: "Y"}}
	cs := []Client{{Slug: "x", Name: "X"}, {Slug: "y", Name: "Y"}}
	if got := clientURL(cs, ps, "X"); got != "/clients/x/" {
		t.Errorf("a client with two projects: %q, want its own page", got)
	}
	if got := clientURL(cs, ps, "Y"); got != "/clients/#y" {
		t.Errorf("a client with one project: %q, want its place on /clients/", got)
	}
	if pages := clientPages(cs, ps); len(pages) != 2 || pages[1].Path != "/clients/x/" {
		t.Errorf("clientPages: want /clients/ and /clients/x/ only")
	}
}
