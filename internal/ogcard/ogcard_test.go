package ogcard

import (
	"bytes"
	"image/jpeg"
	"os"
	"strings"
	"testing"
)

func drawer(t *testing.T) *Drawer {
	wm, err := os.ReadFile("../../assets/press/sawirstudio-wordmark.svg")
	if err != nil {
		t.Fatal(err)
	}
	d, err := New(wm)
	if err != nil {
		t.Fatal(err)
	}
	return d
}

func TestCard(t *testing.T) {
	d := drawer(t)
	jpg, err := d.Card("Build Your Own CMS Like WordPress", "Blog")
	if err != nil {
		t.Fatal(err)
	}
	img, err := jpeg.Decode(bytes.NewReader(jpg))
	if err != nil {
		t.Fatal(err)
	}
	if b := img.Bounds(); b.Dx() != W || b.Dy() != H {
		t.Errorf("card is %v, want %dx%d", b, W, H)
	}
	again, _ := d.Card("Build Your Own CMS Like WordPress", "Blog")
	if !bytes.Equal(jpg, again) {
		t.Error("the same card should come out byte for byte the same")
	}
}

func TestSVG(t *testing.T) {
	d := drawer(t)
	svg := d.SVG(Home, "")
	// Three highlighted runs, rounded (oksvg needs ry), and the wordmark scaled on both axes.
	if n := strings.Count(svg, `ry="`); n < 4 {
		t.Errorf("%d rounded rects, want the wordmark's and three highlights", n)
	}
	if !strings.Contains(svg, "scale(0.04302 0.04302)") {
		t.Error("wordmark should be scaled on both axes")
	}
	if strings.Contains(d.SVG("Hi", ""), `rx="22"`) {
		t.Error("no label pill without a label")
	}
}

func TestLayoutWraps(t *testing.T) {
	d := drawer(t)
	lines := d.layout(runs("Why Shopify Payouts Don't Match Your Books, and How to Fix It"), 100, 900)
	if len(lines) < 2 {
		t.Errorf("a long title should wrap, got %d line(s)", len(lines))
	}
	for _, l := range lines {
		if w := d.lineWidth(l, 100); w > 900 && len(l) > 1 {
			t.Errorf("line %v is %.0fpx, wider than 900", l, w)
		}
	}
}
