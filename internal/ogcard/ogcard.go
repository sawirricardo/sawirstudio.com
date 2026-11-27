// Package ogcard draws the share cards (og:image) for pages without a picture of their own:
// a 1200x630 card with the Sawir Studio wordmark, the page's title in Bricolage Grotesque,
// a label (Blog, Services), and the site's little doodles.
//
// The card is put together as SVG and rasterized with oksvg; the title's letters are the
// font's outlines (go-text/typesetting shapes the text like HarfBuzz and applies the
// variable font's weight and optical size). The font (SIL Open Font License, OFL.txt) is
// built in, so cards don't depend on what's installed. The JPEG bytes are stable on one
// machine; another CPU architecture may differ slightly and redraw them.
package ogcard

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-text/typesetting/di"
	"github.com/go-text/typesetting/font"
	ot "github.com/go-text/typesetting/font/opentype"
	"github.com/go-text/typesetting/language"
	"github.com/go-text/typesetting/shaping"
	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
	"golang.org/x/image/math/fixed"
)

//go:embed BricolageGrotesque.ttf
var fontData []byte

const (
	W, H  = 1200, 630
	bg    = "#fffcf5"
	ink   = "#1f1d1a"
	muted = "#6b665e"
	track = -0.03 // letter-spacing of the site's headings, in em
)

var marks = map[string]string{"yellow": "#ffe580", "coral": "#ffc2b0", "blue": "#c4d8ff", "green": "#bdf0d3"}

// Home is the home page's card title: its hero line. [color]text[/] is a highlighted run.
const Home = "We find what's [coral]slowing your business down[/], and [yellow]fix it[/]."

// Drawer draws cards. It isn't safe for concurrent use.
type Drawer struct {
	bold, semi *face
	wordmark   string // <g> with the wordmark, placed at the card's top left
}

// New prepares a Drawer; wordmarkSVG is assets/press/sawirstudio-wordmark.svg.
func New(wordmarkSVG []byte) (*Drawer, error) {
	bold, err := newFace(map[string]float32{"wght": 800, "opsz": 96, "wdth": 100})
	if err != nil {
		return nil, err
	}
	semi, err := newFace(map[string]float32{"wght": 600, "opsz": 24, "wdth": 100})
	if err != nil {
		return nil, err
	}
	wm, err := wordmark(string(wordmarkSVG), 72, 58, 66)
	if err != nil {
		return nil, err
	}
	return &Drawer{bold: bold, semi: semi, wordmark: wm}, nil
}

// Card returns the card for a page as a JPEG.
func (d *Drawer) Card(title, label string) ([]byte, error) {
	svg := d.SVG(title, label)
	icon, err := oksvg.ReadIconStream(strings.NewReader(svg), oksvg.StrictErrorMode)
	if err != nil {
		return nil, fmt.Errorf("card %q: %w", title, err)
	}
	img := image.NewRGBA(image.Rect(0, 0, W, H))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)
	icon.SetTarget(0, 0, W, H)
	icon.Draw(rasterx.NewDasher(W, H, rasterx.NewScannerGV(W, H, img, img.Bounds())), 1)
	var buf bytes.Buffer
	if err := jpeg.Encode(&buf, img, &jpeg.Options{Quality: 88}); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// SVG is the card as SVG.
func (d *Drawer) SVG(title, label string) string {
	const x, maxW = 72.0, 900.0
	var b strings.Builder
	fmt.Fprintf(&b, `<svg xmlns="http://www.w3.org/2000/svg" width="%d" height="%d" viewBox="0 0 %d %d">`, W, H, W, H)
	fmt.Fprintf(&b, `<rect width="%d" height="%d" fill="%s"/>`, W, H, bg)
	b.WriteString(decor())
	b.WriteString(d.wordmark)
	top := 196.0
	if label != "" {
		const size = 26.0
		lw := d.semi.width(label, size, 0)
		fmt.Fprintf(&b, `<rect x="%g" y="%g" width="%.1f" height="44" rx="22" ry="22" fill="%s" stroke="%s" stroke-width="3"/>`, x, top-52, lw+36, marks["yellow"], ink)
		fmt.Fprintf(&b, `<path fill="%s" d="%s"/>`, ink, d.semi.path(label, x+18, top-21, size, 0))
		top += 20
	}
	b.WriteString(d.title(title, x, top, maxW, 520-top))
	fmt.Fprintf(&b, `<path fill="%s" d="%s"/>`, muted, d.semi.path("sawirstudio.com", x, 568, 30, 0))
	b.WriteString("</svg>")
	return b.String()
}

// ---------------------------------------------------------------- type

type face struct {
	f         *font.Face
	upem      float64
	asc, desc float64
}

func newFace(axes map[string]float32) (*face, error) {
	ft, err := font.ParseTTF(bytes.NewReader(fontData))
	if err != nil {
		return nil, err
	}
	f := font.NewFace(ft.Font)
	var vs []font.Variation
	for tag, v := range axes {
		vs = append(vs, font.Variation{Tag: ot.MustNewTag(tag), Value: v})
	}
	f.SetVariations(vs)
	ext, _ := f.FontHExtents()
	return &face{f: f, upem: float64(f.Upem()), asc: float64(ext.Ascender), desc: float64(ext.Descender)}, nil
}

type glyph struct {
	id                  font.GID
	advance, xOff, yOff float64 // font units
}

func (fc *face) shape(text string) []glyph {
	rs := []rune(text)
	out := (&shaping.HarfbuzzShaper{}).Shape(shaping.Input{
		Text: rs, RunEnd: len(rs), Direction: di.DirectionLTR, Face: fc.f,
		Size: fixed.I(int(fc.upem)), Script: language.Latin, Language: language.NewLanguage("en"),
	})
	gs := make([]glyph, len(out.Glyphs))
	for i, g := range out.Glyphs {
		gs[i] = glyph{g.GlyphID, float64(g.XAdvance) / 64, float64(g.XOffset) / 64, float64(g.YOffset) / 64}
	}
	return gs
}

// width of text at size px, with letter-spacing tr (em) between glyphs.
func (fc *face) width(text string, size, tr float64) float64 {
	gs := fc.shape(text)
	sum := 0.0
	for _, g := range gs {
		sum += g.advance
	}
	return (sum + tr*fc.upem*math.Max(float64(len(gs)-1), 0)) * size / fc.upem
}

// path is SVG path data for text with its baseline at (x, y), in px.
func (fc *face) path(text string, x, y, size, tr float64) string {
	s := size / fc.upem
	var b strings.Builder
	cx := 0.0
	for _, g := range fc.shape(text) {
		ox, oy := x+(cx+g.xOff)*s, y-g.yOff*s
		pt := func(p ot.SegmentPoint) string {
			return num(ox+float64(p.X)*s) + " " + num(oy-float64(p.Y)*s)
		}
		if o, ok := fc.f.GlyphData(g.id).(font.GlyphOutline); ok {
			started := false
			for _, seg := range o.Segments {
				switch seg.Op {
				case ot.SegmentOpMoveTo:
					if started {
						b.WriteString("Z")
					}
					b.WriteString("M" + pt(seg.Args[0]))
					started = true
				case ot.SegmentOpLineTo:
					b.WriteString("L" + pt(seg.Args[0]))
				case ot.SegmentOpQuadTo:
					b.WriteString("Q" + pt(seg.Args[0]) + " " + pt(seg.Args[1]))
				case ot.SegmentOpCubeTo:
					b.WriteString("C" + pt(seg.Args[0]) + " " + pt(seg.Args[1]) + " " + pt(seg.Args[2]))
				}
			}
			if started {
				b.WriteString("Z")
			}
		}
		cx += g.advance + tr*fc.upem
	}
	return b.String()
}

func num(f float64) string { return strconv.FormatFloat(math.Round(f*100)/100, 'f', -1, 64) }

// ---------------------------------------------------------------- title

type word struct {
	text  string
	color string // a mark color, "" for none, or "glue": punctuation drawn right after the previous word
}

var runRE = regexp.MustCompile(`\[(\w+)\](.*?)\[/\]|([^\[]+)`)
var punctRE = regexp.MustCompile(`^[,.;:!?]+$`)

func runs(text string) []word {
	var out []word
	for _, m := range runRE.FindAllStringSubmatch(text, -1) {
		col, words := "", m[3]
		if m[1] != "" {
			col, words = marks[m[1]], m[2]
		}
		for _, w := range strings.Fields(words) {
			if len(out) > 0 && col == "" && punctRE.MatchString(w) {
				out = append(out, word{w, "glue"})
			} else {
				out = append(out, word{w, col})
			}
		}
	}
	return out
}

func (d *Drawer) lineWidth(line []word, size float64) float64 {
	space := d.bold.width(" ", size, 0)
	sum := 0.0
	for i, w := range line {
		sum += d.bold.width(w.text, size, track)
		if i > 0 && w.color != "glue" {
			sum += space
		}
	}
	return sum
}

func (d *Drawer) layout(words []word, size, maxW float64) [][]word {
	var lines [][]word
	var line []word
	for _, w := range words {
		trial := append(append([]word{}, line...), w)
		if len(line) > 0 && w.color != "glue" && d.lineWidth(trial, size) > maxW {
			lines = append(lines, line)
			line = []word{w}
		} else {
			line = trial
		}
	}
	if len(line) > 0 {
		lines = append(lines, line)
	}
	return lines
}

// title draws the title as big as it fits (100px down to 48px), centred in its area.
func (d *Drawer) title(text string, x, top, maxW, maxH float64) string {
	words := runs(text)
	var size, lh float64
	var lines [][]word
	for s := 100; s > 47; s -= 2 {
		size, lh = float64(s), float64(s)*1.12
		lines = d.layout(words, size, maxW)
		fits := float64(len(lines))*lh <= maxH
		for _, l := range lines {
			fits = fits && d.lineWidth(l, size) <= maxW
		}
		if fits {
			break
		}
	}
	top += (maxH - float64(len(lines))*lh) / 2
	space := d.bold.width(" ", size, 0)
	sc := size / d.bold.upem
	y := top + (lh-(d.bold.asc-d.bold.desc)*sc)/2 + d.bold.asc*sc
	pad, rad := 0.15*size, 0.25*size
	var b strings.Builder
	for _, line := range lines {
		cx := x
		type span struct {
			x0, x1 float64
			color  string
		}
		var spans []span
		var glyphs strings.Builder
		for i, w := range line {
			ww := d.bold.width(w.text, size, track)
			col := w.color
			if col == "glue" {
				cx -= space
				if prev := line[i-1].color; prev != "" && prev != "glue" {
					cx += pad // past the highlight's padding, as on the page
				}
				col = ""
			}
			switch {
			case col != "" && len(spans) > 0 && spans[len(spans)-1].color == col && i > 0 && line[i-1].color == col:
				spans[len(spans)-1].x1 = cx + ww
			case col != "":
				spans = append(spans, span{cx, cx + ww, col})
			}
			fmt.Fprintf(&glyphs, `<path fill="%s" d="%s"/>`, ink, d.bold.path(w.text, cx, y, size, track))
			cx += ww + space
		}
		for _, s := range spans {
			fmt.Fprintf(&b, `<rect x="%.1f" y="%.1f" width="%.1f" height="%.1f" rx="%.1f" ry="%.1f" fill="%s"/>`,
				s.x0-pad, y-0.86*size, s.x1-s.x0+2*pad, 1.1*size, rad, rad, s.color)
		}
		b.WriteString(glyphs.String())
		y += lh
	}
	return b.String()
}

// ---------------------------------------------------------------- pieces

var (
	viewBoxRE = regexp.MustCompile(`viewBox="([^"]+)"`)
	svgTagRE  = regexp.MustCompile(`(?s)^<svg[^>]*>|</svg>\s*$|<title>.*?</title>`)
	rxOnlyRE  = regexp.MustCompile(`(\srx="([^"]*)")`)
	strokeRE  = regexp.MustCompile(`stroke-width="([^"]*)"`)
)

func wordmark(svg string, x, y, height float64) (string, error) {
	m := viewBoxRE.FindStringSubmatch(svg)
	if m == nil {
		return "", fmt.Errorf("wordmark: no viewBox")
	}
	var vx, vy, vw, vh float64
	if _, err := fmt.Sscan(m[1], &vx, &vy, &vw, &vh); err != nil {
		return "", fmt.Errorf("wordmark viewBox: %w", err)
	}
	inner := svgTagRE.ReplaceAllString(strings.TrimSpace(svg), "")
	// oksvg leaves ry at 0 when only rx is given (SVG says ry defaults to rx).
	inner = rxOnlyRE.ReplaceAllString(inner, `$1 ry="$2"`)
	s := height / vh
	// oksvg doesn't scale stroke widths by the transform, so scale them here.
	inner = strokeRE.ReplaceAllStringFunc(inner, func(m string) string {
		w, _ := strconv.ParseFloat(strokeRE.FindStringSubmatch(m)[1], 64)
		return fmt.Sprintf(`stroke-width="%.3f"`, w*s)
	})
	// Two scale arguments: oksvg reads scale(s) as scale(s, 0).
	return fmt.Sprintf(`<g transform="translate(%g %g) scale(%.5f %.5f) translate(%g %g)">%s</g>`, x, y, s, s, -vx, -vy, inner), nil
}

func star(cx, cy, r, rot float64) string {
	pts := make([]string, 10)
	for i := range pts {
		a := (rot - 90 + float64(i)*36) * math.Pi / 180
		rr := r
		if i%2 == 1 {
			rr = r * 0.45
		}
		pts[i] = fmt.Sprintf("%.1f,%.1f", cx+rr*math.Cos(a), cy+rr*math.Sin(a))
	}
	return fmt.Sprintf(`<polygon points="%s" fill="#ffd84d" stroke="%s" stroke-width="6" stroke-linejoin="round"/>`, strings.Join(pts, " "), ink)
}

func decor() string {
	return star(1100, 112, 46, -12) +
		`<path d="M1052 210v34M1035 227h34" stroke="#ff7a59" stroke-width="6" stroke-linecap="round"/>` +
		fmt.Sprintf(`<circle cx="1112" cy="452" r="16" fill="#c4d8ff" stroke="%s" stroke-width="4"/>`, ink) +
		`<path d="M1010 540q20-22 40 0t40 0t40 0" fill="none" stroke="#3fc283" stroke-width="9" stroke-linecap="round"/>`
}
