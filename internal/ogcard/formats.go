package ogcard

import (
	"fmt"
	"strings"
)

// Format is a size of share picture made for posting rather than for link previews:
// an Instagram story, a feed post, a square. (The link format is the 1200x630 Card.)
type Format struct {
	Name string // file suffix and the share dialog's key: "story", "feed", "square"
	W, H int
	// Top and Bottom keep the content clear of what the app draws over the picture
	// (a story's progress bar and profile at the top, the reply bar at the bottom).
	Top, Bottom float64
	// Sticker is empty space kept under the text, where a story's link sticker goes.
	Sticker   float64
	TitleMax  int // biggest title size, px
	DescLines int // the description is cut to this many lines (0 leaves it out)
}

// Formats are drawn for every blog post, next to its link card.
var Formats = []Format{
	{Name: "story", W: 1080, H: 1920, Top: 250, Bottom: 250, Sticker: 300, TitleMax: 128, DescLines: 5},
	{Name: "feed", W: 1080, H: 1350, Top: 80, Bottom: 80, TitleMax: 112, DescLines: 4},
	{Name: "square", W: 1080, H: 1080, Top: 72, Bottom: 72, TitleMax: 100, DescLines: 2},
}

// Post returns the post's picture in format f as a JPEG.
func (d *Drawer) Post(f Format, title, desc, label string) ([]byte, error) {
	svg, err := d.PostSVG(f, title, desc, label)
	if err != nil {
		return nil, err
	}
	jpg, err := rasterize(svg, f.W, f.H)
	if err != nil {
		return nil, fmt.Errorf("%s card %q: %w", f.Name, title, err)
	}
	return jpg, nil
}

// PostSVG is the post's picture in format f as SVG: the wordmark at the top, then the label,
// the title and the description, centred in the space left, and the site's address at the
// bottom, with the doodles down the right edge.
func (d *Drawer) PostSVG(f Format, title, desc, label string) (string, error) {
	const x = 72.0
	w, h := float64(f.W), float64(f.H)
	maxW := w - 2*x
	var b strings.Builder
	fmt.Fprintf(&b, `<svg xmlns="http://www.w3.org/2000/svg" width="%d" height="%d" viewBox="0 0 %d %d">`, f.W, f.H, f.W, f.H)
	fmt.Fprintf(&b, `<rect width="%d" height="%d" fill="%s"/>`, f.W, f.H, bg)

	const wmH = 76.0
	wm, err := wordmark(d.wordmarkSVG, x, f.Top, wmH)
	if err != nil {
		return "", err
	}
	b.WriteString(wm)
	b.WriteString(star(w-120, f.Top+wmH/2, 50, -12))
	footY := h - f.Bottom
	fmt.Fprintf(&b, `<path fill="%s" d="%s"/>`, muted, d.semi.path("sawirstudio.com", x, footY, 36, 0))
	fmt.Fprintf(&b, `<path d="M%g %gq20-22 40 0t40 0t40 0" fill="none" stroke="#3fc283" stroke-width="9" stroke-linecap="round"/>`, w-212, footY-14)

	// The space for the text: under the wordmark, above the footer (and a story's sticker).
	top, bottom := f.Top+wmH+64, footY-48-f.Sticker
	const pillH, pillGap = 52.0, 36.0
	const descSize, descLH, descGap = 40.0, 40.0 * 1.4, 44.0
	var descLines []string
	if f.DescLines > 0 && strings.TrimSpace(desc) != "" {
		descLines = d.wrap(d.text, desc, descSize, maxW, f.DescLines)
	}
	room := bottom - top
	if label != "" {
		room -= pillH + pillGap
	}
	if len(descLines) > 0 {
		room -= descGap + float64(len(descLines))*descLH
	}
	lines, size, lh := d.fitTitle(title, maxW, room, f.TitleMax)
	block := float64(len(lines)) * lh
	if label != "" {
		block += pillH + pillGap
	}
	if len(descLines) > 0 {
		block += descGap + float64(len(descLines))*descLH
	}
	y := top + (bottom-top-block)/2

	if label != "" {
		const size = 30.0
		lw := d.semi.width(label, size, 0)
		fmt.Fprintf(&b, `<rect x="%g" y="%g" width="%.1f" height="%g" rx="%g" ry="%g" fill="%s" stroke="%s" stroke-width="3"/>`, x, y, lw+42, pillH, pillH/2, pillH/2, marks["yellow"], ink)
		fmt.Fprintf(&b, `<path fill="%s" d="%s"/>`, ink, d.semi.path(label, x+21, y+36, size, 0))
		y += pillH + pillGap
	}
	b.WriteString(d.drawTitle(lines, size, lh, x, y))
	y += float64(len(lines)) * lh
	if len(descLines) > 0 {
		// A short coral squiggle between the title and the description.
		fmt.Fprintf(&b, `<path d="M%g %gq15-14 30 0t30 0t30 0" fill="none" stroke="#ff7a59" stroke-width="7" stroke-linecap="round"/>`, x+4, y+descGap/2)
		y += descGap
		sc := descSize / d.text.upem
		base := y + (descLH-(d.text.asc-d.text.desc)*sc)/2 + d.text.asc*sc
		for _, l := range descLines {
			fmt.Fprintf(&b, `<path fill="%s" d="%s"/>`, muted, d.text.path(l, x, base, descSize, 0))
			base += descLH
		}
	}
	b.WriteString("</svg>")
	return b.String(), nil
}

// wrap breaks text into at most max lines of maxW px, ending a cut text with an ellipsis.
func (d *Drawer) wrap(fc *face, text string, size, maxW float64, max int) []string {
	var lines []string
	line := ""
	words := strings.Fields(text)
	for _, w := range words {
		trial := strings.TrimSpace(line + " " + w)
		if line != "" && fc.width(trial, size, 0) > maxW {
			lines = append(lines, line)
			if len(lines) == max {
				return ellipsize(fc, lines, size, maxW)
			}
			line = w
		} else {
			line = trial
		}
	}
	if line != "" {
		lines = append(lines, line)
	}
	return lines
}

// ellipsize ends the last line with "…", dropping words until it fits.
func ellipsize(fc *face, lines []string, size, maxW float64) []string {
	last := strings.Fields(lines[len(lines)-1])
	for len(last) > 1 && fc.width(strings.Join(last, " ")+"…", size, 0) > maxW {
		last = last[:len(last)-1]
	}
	s := strings.TrimRight(strings.Join(last, " "), ",.;:—-")
	lines[len(lines)-1] = s + "…"
	return lines
}
