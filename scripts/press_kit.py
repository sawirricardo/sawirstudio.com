# /// script
# requires-python = ">=3.11"
# dependencies = ["fonttools", "uharfbuzz", "resvg-py", "beautifulsoup4"]
# ///
"""Build the press kit files in assets/press/.

- The wordmark ("Sawir" + "Studio" in the tilted yellow box, as in the site header), drawn from
  Bricolage Grotesque ExtraBold as outlines, so it doesn't need the font. Ink and on-dark versions.
- The icon from favicon.svg, with fixed colors (favicon.svg switches with the viewer's theme).
- A PNG of each, and sawirstudio-press-kit.zip with the logos, the founder photo, the app icons
  and about.txt (facts and boilerplate, taken from press/index.html).

    uv run scripts/press_kit.py

Rerun after changing the logo, the facts or the boilerplate on press/index.html.
"""

import math, pathlib, re, tempfile, urllib.request, zipfile
import uharfbuzz as hb
from bs4 import BeautifulSoup
from fontTools.ttLib import TTFont
from fontTools.varLib.instancer import instantiateVariableFont
from fontTools.pens.svgPathPen import SVGPathPen
from fontTools.pens.transformPen import TransformPen
import resvg_py

ROOT = pathlib.Path(__file__).resolve().parent.parent
FONT_URL = "https://github.com/google/fonts/raw/main/ofl/bricolagegrotesque/BricolageGrotesque%5Bopsz,wdth,wght%5D.ttf"
favicon = ROOT / "favicon.svg"
out = ROOT / "assets/press"
font_path = pathlib.Path(tempfile.mkdtemp()) / "bricolage.ttf"
urllib.request.urlretrieve(FONT_URL, font_path)
out.mkdir(parents=True, exist_ok=True)
AXES = {"wght": 800, "opsz": 30, "wdth": 100}

tt = instantiateVariableFont(TTFont(font_path), AXES)
upm = tt["head"].unitsPerEm
gs = tt.getGlyphSet()
asc, desc = tt["hhea"].ascent, tt["hhea"].descent  # desc negative

blob = hb.Blob.from_file_path(str(font_path))
hbfont = hb.Font(hb.Face(blob))
hbfont.set_variations(AXES)

def shape(text):
    buf = hb.Buffer(); buf.add_str(text); buf.guess_segment_properties()
    hb.shape(hbfont, buf)
    names = [hbfont.glyph_to_string(i.codepoint) for i in buf.glyph_infos]
    return list(zip(names, buf.glyph_positions))

def run_path(text, x0, y0, tracking):
    """Glyph outlines for text with baseline at (x0, y0) in y-down units. Returns (d, width)."""
    pen = SVGPathPen(gs)
    x = x0
    items = shape(text)
    for i, (name, pos) in enumerate(items):
        tp = TransformPen(pen, (1, 0, 0, -1, x + pos.x_offset, y0 - pos.y_offset))
        gs[name].draw(tp)
        x += pos.x_advance + (tracking if i < len(items) - 1 else 0)
    return pen.getCommands(), x - x0

em = upm
track = -0.03 * em                    # letter-spacing: -0.03em (applied between letters)
line = 1.1 * em                       # line-height: 1.1
base = (line - (asc - desc)) / 2 + asc  # baseline within a line box
border = 2 / 30 * em                  # 2px at 30px
pad = 0.18 * em
gap = 0.08 * em
radius = 0.3 * em

sawir_d, sawir_w = run_path("Sawir", 0, base, track)
sawir_w += track  # CSS adds spacing after the last letter of the run too
box_x = sawir_w + gap
studio_d, studio_w = run_path("Studio", box_x + border + pad, base, track)
box_w = border * 2 + pad * 2 + studio_w + track
box_h = line + border * 2
box_y = -border  # the span's border sits outside the line box
cx, cy = box_x + box_w / 2, box_y + box_h / 2

# bounding box incl. rotation of the box
ang = math.radians(-3)
corners = [(box_x, box_y), (box_x + box_w, box_y), (box_x, box_y + box_h), (box_x + box_w, box_y + box_h)]
rot = [(cx + (x - cx) * math.cos(ang) - (y - cy) * math.sin(ang), cy + (x - cx) * math.sin(ang) + (y - cy) * math.cos(ang)) for x, y in corners]
m = 0.06 * em
minx = min(0, *[p[0] for p in rot]) - m
maxx = max(sawir_w, *[p[0] for p in rot]) + m
miny = min(0, *[p[1] for p in rot]) - m
maxy = max(line, *[p[1] for p in rot]) + m
vb = f"{minx:.1f} {miny:.1f} {maxx - minx:.1f} {maxy - miny:.1f}"
f = lambda v: f"{v:.1f}"

def wordmark(word_color, border_color):
    return (f'<svg xmlns="http://www.w3.org/2000/svg" viewBox="{vb}">'
            f'<title>Sawir Studio</title>'
            f'<path fill="{word_color}" d="{sawir_d}"/>'
            f'<g transform="rotate(-3 {f(cx)} {f(cy)})">'
            f'<rect x="{f(box_x + border / 2)}" y="{f(box_y + border / 2)}" width="{f(box_w - border)}" height="{f(box_h - border)}" rx="{f(radius - border / 2)}" fill="#ffd84d" stroke="{border_color}" stroke-width="{f(border)}"/>'
            f'<path fill="#1f1d1a" d="{studio_d}"/></g></svg>\n')

fav = favicon.read_text()
fav = re.sub(r"<style>.*?</style>", "", fav)
def mark(shadow):
    return fav.replace('class="shadow" ', "").replace('fill="#1f1d1a"/><rect', f'fill="{shadow}"/><rect', 1).replace("<svg ", "<svg ", 1).replace('viewBox="0 0 64 64">', 'viewBox="0 0 64 64"><title>Sawir Studio</title>', 1)

files = {
    "sawirstudio-wordmark.svg": wordmark("#1f1d1a", "#1f1d1a"),
    "sawirstudio-wordmark-on-dark.svg": wordmark("#f5efe3", "#f5efe3"),
    "sawirstudio-mark.svg": mark("#1f1d1a"),
    "sawirstudio-mark-on-dark.svg": mark("#f5efe3"),
}
for name, svg in files.items():
    (out / name).write_text(svg)
    w = 2400 if "wordmark" in name else 1024
    png = resvg_py.svg_to_bytes(svg_string=svg, width=w)
    (out / name.replace(".svg", ".png")).write_bytes(bytes(png))


# The zip: logos, founder photo, app icons and about.txt.
soup = BeautifulSoup((ROOT / "press/index.html").read_text(), "html.parser")
clean = lambda el: re.sub(r" ([,.;:])", r"\1", re.sub(r"\s+", " ", el.get_text(" "))).strip()
facts = [f"{clean(li.strong)}: {clean(li)[len(clean(li.strong)):].strip()}" for li in soup.select(".press-facts li")]
paras = soup.select(".press-intro .prose p")
about = ("Sawir Studio press kit\nhttps://sawirstudio.com/press/\n\n" + "\n".join(facts)
         + "\n\nIn one line\n" + clean(paras[0]) + "\n\nIn a paragraph\n" + clean(paras[1])
         + "\n\nFounder\n" + clean(soup.select_one(".press-bio p"))
         + "\n\nColors\n" + "\n".join(f"{clean(li.strong)}: {clean(li.code)}" for li in soup.select(".swatches li"))
         + "\n\nType\n" + clean(soup.select_one(".swatches + .prose p")) + "\n")
zpath = out / "sawirstudio-press-kit.zip"
with zipfile.ZipFile(zpath, "w", zipfile.ZIP_DEFLATED) as z:
    def add(name, data):
        # A fixed date keeps the zip byte-identical when nothing changed.
        info = zipfile.ZipInfo(f"sawirstudio-press-kit/{name}", (2025, 1, 1, 0, 0, 0))
        info.compress_type = zipfile.ZIP_DEFLATED
        z.writestr(info, data)
    add("about.txt", about)
    for f in sorted(out.glob("sawirstudio-*.*")):
        if f.suffix in (".svg", ".png"):
            add(f"logo/{f.name}", f.read_bytes())
    add("founder/ricardo-sawir.webp", (ROOT / "assets/images/ricardo-sawir.webp").read_bytes())
    for a in soup.select(".app-row a"):
        img = ROOT / a.img["src"].lstrip("/")
        add(f"app-icons/{img.name}", img.read_bytes())
print("wrote", ", ".join(sorted(p.name for p in out.iterdir())))
