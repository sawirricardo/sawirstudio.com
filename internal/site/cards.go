package site

import (
	"bytes"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/sawirricardo/sawirstudio.com/internal/ogcard"
)

// Service pages get a "Services" label on their card.
var servicePages = map[string]bool{"shopify": true, "odoo": true, "custom": true, "launch-and-growth": true, "forward-deployed-engineering": true}

// cardText is the title and label drawn on a page's share card.
func (p *Page) cardText() (title, label string) {
	if p.cardName() == "home" {
		return ogcard.Home, ""
	}
	parts := strings.Split(strings.Trim(p.Path, "/"), "/")
	title = strings.TrimSuffix(strings.TrimSpace(p.ogTitle()), " | Sawir Studio")
	switch {
	case parts[0] == "posts" && len(parts) > 1:
		label = "Blog"
	case servicePages[parts[0]]:
		label = "Services"
	}
	return title, label
}

// drawCards draws assets/og/cards/<name>.jpg for every page without its own picture, and
// removes cards no page uses. It returns the files that changed (or would, without write).
func drawCards(root string, pages []*Page, write bool) ([]string, error) {
	dir := filepath.Join(root, "assets", "og", "cards")
	wm, err := os.ReadFile(filepath.Join(root, "assets", "press", "sawirstudio-wordmark.svg"))
	if err != nil {
		return nil, err
	}
	d, err := ogcard.New(wm)
	if err != nil {
		return nil, err
	}
	var changed []string
	used := map[string]bool{}
	for _, p := range pages {
		name := p.cardName()
		if p.OGImage != "" || used[name] {
			continue
		}
		used[name] = true
		jpg, err := d.Card(p.cardText())
		if err != nil {
			return nil, err
		}
		file := filepath.Join(dir, name+".jpg")
		if old, err := os.ReadFile(file); err == nil && bytes.Equal(old, jpg) {
			continue
		}
		rel, _ := filepath.Rel(root, file)
		changed = append(changed, rel)
		if write {
			if err := os.MkdirAll(dir, 0o755); err != nil {
				return nil, err
			}
			if err := os.WriteFile(file, jpg, 0o644); err != nil {
				return nil, err
			}
		}
	}
	entries, err := os.ReadDir(dir)
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		return nil, err
	}
	for _, e := range entries {
		name, ok := strings.CutSuffix(e.Name(), ".jpg")
		if !ok || used[name] {
			continue
		}
		rel, _ := filepath.Rel(root, filepath.Join(dir, e.Name()))
		changed = append(changed, rel)
		if write {
			if err := os.Remove(filepath.Join(dir, e.Name())); err != nil {
				return nil, err
			}
		}
	}
	return changed, nil
}
