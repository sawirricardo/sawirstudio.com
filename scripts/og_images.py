# /// script
# requires-python = ">=3.11"
# dependencies = ["fonttools", "uharfbuzz", "resvg-py", "pillow"]
# ///
"""Build the share cards (og:image) and the image tags on every English page.

- Draws a 1200x630 card in assets/og/cards/ for each English page that has no picture of its own
  (home, blog posts, service pages, about, contact and so on): the Sawir Studio wordmark, the page's
  title in Bricolage Grotesque, and a label (Blog, Services). Project pages keep their own picture.
- Points each of those pages' og:image (and the JSON-LD "image") at its card.
- Adds og:image:width, og:image:height (read from the file), og:image:alt and twitter:image:alt to
  every English page. The alt is the page's og:title, so i18n.py already has it translated.

    uv run scripts/og_images.py
    uv run scripts/i18n.py --no-translate
    uv run scripts/markdown_pages.py

Rerun after adding a page or changing a page's title. Running it twice changes nothing.
"""

import html, io, math, pathlib, re, tempfile, urllib.request
import resvg_py
import uharfbuzz as hb
from fontTools.pens.svgPathPen import SVGPathPen
from fontTools.pens.transformPen import TransformPen
from fontTools.ttLib import TTFont
from fontTools.varLib.instancer import instantiateVariableFont
from PIL import Image

ROOT = pathlib.Path(__file__).resolve().parent.parent
SITE = "https://sawirstudio.com"
CARDS = ROOT / "assets/og/cards"
FONT_URL = "https://github.com/google/fonts/raw/main/ofl/bricolagegrotesque/BricolageGrotesque%5Bopsz,wdth,wght%5D.ttf"
LANG_DIRS = {"zh-hans", "zh-hant", "ja", "de", "fr", "ru", "hi", "nl", "it", "he", "ar"}
SERVICES = {"shopify", "odoo", "custom", "launch-and-growth", "forward-deployed-engineering"}
W, H = 1200, 630
BG, INK, MUTED = "#fffcf5", "#1f1d1a", "#6b665e"
MARKS = {"yellow": "#ffe580", "coral": "#ffc2b0", "blue": "#c4d8ff", "green": "#bdf0d3"}
# The home card is the home page's hero line; [color]text[/] is a highlighted run, as on the page.
HOME = "We build software [yellow]people use[/], for [coral]clients[/] and for [blue]ourselves[/]."


# ---------------------------------------------------------------- type

font_file = pathlib.Path(tempfile.mkdtemp()) / "bricolage.ttf"
urllib.request.urlretrieve(FONT_URL, font_file)
blob = hb.Blob.from_file_path(str(font_file))


class Face:
    def __init__(self, axes):
        tt = instantiateVariableFont(TTFont(font_file), axes)
        self.upm = tt["head"].unitsPerEm
        self.gs = tt.getGlyphSet()
        self.asc, self.desc = tt["hhea"].ascent, tt["hhea"].descent
        self.hb = hb.Font(hb.Face(blob))
        self.hb.set_variations(axes)

    def shape(self, text):
        buf = hb.Buffer(); buf.add_str(text); buf.guess_segment_properties()
        hb.shape(self.hb, buf)
        return [(self.hb.glyph_to_string(i.codepoint), p) for i, p in zip(buf.glyph_infos, buf.glyph_positions)]

    def width(self, text, size, track=0.0):
        items = self.shape(text)
        return (sum(p.x_advance for _, p in items) + track * self.upm * max(len(items) - 1, 0)) * size / self.upm

    def path(self, text, x, y, size, track=0.0):
        """SVG path data for text with its baseline at (x, y), in px."""
        pen = SVGPathPen(self.gs)
        s = size / self.upm
        cx = 0
        for name, p in self.shape(text):
            t = TransformPen(pen, (s, 0, 0, -s, x + (cx + p.x_offset) * s, y - p.y_offset * s))
            self.gs[name].draw(t)
            cx += p.x_advance + track * self.upm
        return pen.getCommands()


BOLD = Face({"wght": 800, "opsz": 96, "wdth": 100})
SEMI = Face({"wght": 600, "opsz": 24, "wdth": 100})
TRACK = -0.03  # letter-spacing of the site's headings, in em


# ---------------------------------------------------------------- pieces

def wordmark(x, y, height):
    svg = (ROOT / "assets/press/sawirstudio-wordmark.svg").read_text()
    vx, vy, vw, vh = map(float, re.search(r'viewBox="([^"]+)"', svg)[1].split())
    inner = re.sub(r"^<svg[^>]*>|</svg>\s*$|<title>.*?</title>", "", svg.strip())
    s = height / vh
    return f'<g transform="translate({x} {y}) scale({s:.5f}) translate({-vx} {-vy})">{inner}</g>'


def star(cx, cy, r, rot=-12):
    pts = []
    for i in range(10):
        a = math.radians(rot - 90 + i * 36)
        rr = r if i % 2 == 0 else r * 0.45
        pts.append(f"{cx + rr * math.cos(a):.1f},{cy + rr * math.sin(a):.1f}")
    return f'<polygon points="{" ".join(pts)}" fill="#ffd84d" stroke="{INK}" stroke-width="6" stroke-linejoin="round"/>'


DECOR = (
    star(1100, 112, 46)
    + f'<path d="M1052 210v34M1035 227h34" stroke="#ff7a59" stroke-width="6" stroke-linecap="round"/>'
    + f'<circle cx="1112" cy="452" r="16" fill="#c4d8ff" stroke="{INK}" stroke-width="4"/>'
    + '<path d="M1010 540q20-22 40 0t40 0t40 0" fill="none" stroke="#3fc283" stroke-width="9" stroke-linecap="round"/>'
)


def runs(text):
    """[(word, mark color or None)] from text with [color]...[/] runs."""
    out = []
    for m in re.finditer(r"\[(\w+)\](.*?)\[/\]|([^\[]+)", text):
        color, words = (MARKS[m[1]], m[2]) if m[1] else (None, m[3])
        for w in words.split():
            if out and not color and re.fullmatch(r"[,.;:!?]+", w):
                out.append((w, "glue"))  # drawn right after the previous word, no space
            else:
                out.append((w, color))
    return out


def line_width(line, size):
    space = BOLD.width(" ", size)
    return sum(BOLD.width(t, size, TRACK) + (space if i and c != "glue" else 0) for i, (t, c) in enumerate(line))


def layout(words, size, max_w):
    lines, line = [], []
    for w in words:
        trial = line + [w]
        if line and w[1] != "glue" and line_width(trial, size) > max_w:
            lines.append(line); line = [w]
        else:
            line = trial
    return lines + [line] if line else lines


def title_svg(text, x, top, max_w, max_h):
    words = runs(text)
    for size in range(100, 47, -2):
        lh = size * 1.12
        lines = layout(words, size, max_w)
        if len(lines) * lh <= max_h and all(line_width(l, size) <= max_w for l in lines):
            break
    top += (max_h - len(lines) * lh) / 2  # centre the block in its area
    space = BOLD.width(" ", size)
    out, y = [], top + (size * 1.12 - (BOLD.asc - BOLD.desc) * size / BOLD.upm) / 2 + BOLD.asc * size / BOLD.upm
    pad, rad = 0.15 * size, 0.25 * size
    for line in lines:
        cx = x
        spans = []  # [x0, x1, color] of consecutive marked words
        for i, (w, color) in enumerate(line):
            ww = BOLD.width(w, size, TRACK)
            if color == "glue":
                cx -= space
                if line[i - 1][1] not in (None, "glue"):
                    cx += pad  # past the highlight's padding, as on the page
                color = None
            if color and spans and spans[-1][2] == color and i and line[i - 1][1] == color:
                spans[-1][1] = cx + ww
            elif color:
                spans.append([cx, cx + ww, color])
            out.append(f'<path fill="{INK}" d="{BOLD.path(w, cx, y, size, TRACK)}"/>')
            cx += ww + space
        boxes = "".join(
            f'<rect x="{x0 - pad:.1f}" y="{y - 0.86 * size:.1f}" width="{x1 - x0 + 2 * pad:.1f}" height="{1.1 * size:.1f}" rx="{rad:.1f}" fill="{c}"/>'
            for x0, x1, c in spans)
        out.insert(len(out) - len(line), boxes)
        y += size * 1.12
    return "".join(out)


def card(title, label):
    x, max_w = 72, 900
    parts = [f'<rect width="{W}" height="{H}" fill="{BG}"/>', DECOR, wordmark(x, 58, 66)]
    top = 196
    if label:
        size = 26
        lw = SEMI.width(label, size)
        parts.append(f'<rect x="{x}" y="{top - 52}" width="{lw + 36:.1f}" height="44" rx="22" fill="{MARKS["yellow"]}" stroke="{INK}" stroke-width="3"/>')
        parts.append(f'<path fill="{INK}" d="{SEMI.path(label, x + 18, top - 21, size)}"/>')
        top += 20
    parts.append(title_svg(title, x, top, max_w, 520 - top))
    parts.append(f'<path fill="{MUTED}" d="{SEMI.path("sawirstudio.com", x, 568, 30)}"/>')
    svg = f'<svg xmlns="http://www.w3.org/2000/svg" width="{W}" height="{H}" viewBox="0 0 {W} {H}">{"".join(parts)}</svg>'
    png = bytes(resvg_py.svg_to_bytes(svg_string=svg, width=W))
    buf = io.BytesIO()
    Image.open(io.BytesIO(png)).convert("RGB").save(buf, "JPEG", quality=88, optimize=True)
    return buf.getvalue()


# ---------------------------------------------------------------- pages

def english_pages():
    for f in sorted(ROOT.glob("**/index.html")):
        parts = f.relative_to(ROOT).parts
        if any(p.startswith(".") for p in parts) or parts[0] in LANG_DIRS or parts[0] in ("scripts", "node_modules"):
            continue
        yield f
    yield ROOT / "404.html"


def meta(text, attr, key):
    m = re.search(rf'<meta {attr}="{re.escape(key)}" content="([^"]*)"', text)
    return m and m[1]


CARDS.mkdir(parents=True, exist_ok=True)
written, tagged = set(), 0
for f in english_pages():
    text = f.read_text()
    image = meta(text, "property", "og:image")
    og_title = meta(text, "property", "og:title")
    if not image or not og_title:
        print("skip (no og:image or og:title):", f.relative_to(ROOT)); continue
    rel = f.relative_to(ROOT)
    if image.endswith("/assets/og/default.jpg") or "/assets/og/cards/" in image:
        parts = rel.parts[:-1]
        if rel.name == "404.html" or not parts:
            name, title, label = "home", HOME, None
        else:
            name = "-".join(parts)
            title = re.sub(r"\s*\|\s*Sawir Studio$", "", html.unescape(og_title))
            label = "Blog" if parts[0] == "posts" and len(parts) > 1 else "Services" if parts[0] in SERVICES else None
        if name not in written:
            data = card(title, label)
            out = CARDS / f"{name}.jpg"
            if not out.exists() or out.read_bytes() != data:
                out.write_bytes(data)
            written.add(name)
        new = f"{SITE}/assets/og/cards/{name}.jpg"
        text = text.replace(image, new)
        image = new
    file = ROOT / image.removeprefix(SITE + "/")
    w, h = Image.open(file).size
    # Replace any earlier tags, then add them right after og:image.
    text = re.sub(r'\n[ \t]*<meta (?:property="og:image:(?:width|height|alt)"|name="twitter:image:alt")[^\n]*', "", text)
    m = re.search(r'\n([ \t]*)<meta property="og:image"[^\n]*', text)
    ind = m[1]
    tags = (f'\n{ind}<meta property="og:image:width" content="{w}" />'
            f'\n{ind}<meta property="og:image:height" content="{h}" />'
            f'\n{ind}<meta property="og:image:alt" content="{og_title}" />'
            f'\n{ind}<meta name="twitter:image:alt" content="{og_title}" />')
    text = text[:m.end()] + tags + text[m.end():]
    if text != f.read_text():
        f.write_text(text)
    tagged += 1

for old in CARDS.glob("*.jpg"):
    if old.stem not in written:
        old.unlink(); print("removed", old.relative_to(ROOT))
print(f"{len(written)} cards, {tagged} pages tagged")
