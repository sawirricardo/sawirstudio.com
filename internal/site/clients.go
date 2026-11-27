package site

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/a-h/templ"
	"gopkg.in/yaml.v3"
)

// Client is one client in content/clients.yaml. Its projects are the ones in projects.yaml
// whose client: is its Name.
type Client struct {
	Slug  string `yaml:"slug"`
	Name  string `yaml:"name"`
	Intro string `yaml:"intro"` // HTML, one or two sentences
	Quote string `yaml:"quote"` // HTML, an approved quote from the client, shown on its page
	// Logo is the client's logo for the homepage strip (/assets/images/clients/...). LogoDark
	// draws it in ink, for logos made for dark backgrounds; LogoLabel puts the name beside a
	// logo that is only a mark.
	Logo      string `yaml:"logo"`
	LogoDark  bool   `yaml:"logo_dark"`
	LogoLabel bool   `yaml:"logo_label"`
	// Brands are the client's other brands with their own logos, shown in the strip next to it.
	Brands []Brand `yaml:"brands"`
}

// Brand is one of a client's brands in the logo strip, linking to the client.
type Brand struct {
	Name string `yaml:"name"`
	Logo string `yaml:"logo"`
}

// LoadClients reads the client list and checks it against the projects: every project's
// client must be listed, and every listed client must have a project.
func LoadClients(file string, ps []Project) ([]Client, error) {
	src, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var cs []Client
	dec := yaml.NewDecoder(bytes.NewReader(src))
	dec.KnownFields(true)
	if err := dec.Decode(&cs); err != nil {
		return nil, fmt.Errorf("%s: %w", file, err)
	}
	names := map[string]bool{}
	for _, c := range cs {
		if c.Slug == "" || c.Name == "" || c.Intro == "" {
			return nil, fmt.Errorf("%s: every client needs a slug, name and intro", file)
		}
		if len(clientProjects(ps, c.Name)) == 0 {
			return nil, fmt.Errorf("%s: %s has no projects in projects.yaml", file, c.Name)
		}
		names[c.Name] = true
	}
	for _, p := range ps {
		if p.Client != "" && !names[p.Client] {
			return nil, fmt.Errorf("projects.yaml: %s's client %q isn't in %s", p.Slug, p.Client, file)
		}
	}
	return cs, nil
}

// clientProjects are the client's projects, in list order.
func clientProjects(ps []Project, name string) []Project {
	var out []Project
	for _, p := range ps {
		if p.Client == name {
			out = append(out, p)
		}
	}
	return out
}

// hasPage reports whether the client gets its own page: only with more than one project.
func (c Client) hasPage(ps []Project) bool { return len(clientProjects(ps, c.Name)) > 1 }

// clientURL is where a client's name links to: its own page, or its place on the clients page.
func clientURL(cs []Client, ps []Project, name string) string {
	for _, c := range cs {
		if c.Name == name {
			if c.hasPage(ps) {
				return "/clients/" + c.Slug + "/"
			}
			return "/clients/#" + c.Slug
		}
	}
	return "/clients/"
}

// clientPages are the pages the builder makes from the client list: /clients/, and a page for
// each client with more than one project.
func clientPages(cs []Client, ps []Project) []*Page {
	pages := []*Page{{
		Path:        "/clients/",
		Title:       "Clients | Sawir Studio",
		Description: "The companies and institutions Sawir Studio has built for, and what we built for each of them.",
		OGType:      "website",
		Main:        "<!-- @client-list -->",
	}}
	for _, c := range cs {
		if !c.hasPage(ps) {
			continue
		}
		var titles []string
		for _, p := range clientProjects(ps, c.Name) {
			titles = append(titles, p.Title)
		}
		pages = append(pages, &Page{
			Path:        "/clients/" + c.Slug + "/",
			Client:      c.Slug,
			Title:       "Our Work for " + c.Name + " | Sawir Studio",
			Description: "What Sawir Studio has built for " + c.Name + ": " + andList(titles) + ".",
			OGType:      "website",
			Main:        "<!-- @client-page -->",
		})
	}
	return pages
}

func andList(s []string) string {
	if len(s) < 2 {
		return strings.Join(s, "")
	}
	return strings.Join(s[:len(s)-1], ", ") + " and " + s[len(s)-1]
}

func clientBySlug(cs []Client, slug string) Client {
	for _, c := range cs {
		if c.Slug == slug {
			return c
		}
	}
	return Client{}
}

// tabindex takes the copy of the logo strip out of the tab order.
func tabindex(copy bool) templ.Attributes {
	if copy {
		return templ.Attributes{"tabindex": "-1"}
	}
	return templ.Attributes{}
}

// alt is a logo's alt text: the client's name, or empty in the hidden copy.
func alt(name string, copy bool) string {
	if copy {
		return ""
	}
	return name
}
