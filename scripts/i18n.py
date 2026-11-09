# /// script
# requires-python = ">=3.11"
# dependencies = ["beautifulsoup4"]
# ///
"""Build the translated copies of every page under /<lang>/.

The English pages are the source. For each language in LANGS this script:

1. pulls the text out of every English page (each paragraph, heading, list item and so on
   as one piece of HTML, plus titles, descriptions, alt text and the JSON-LD names),
2. translates the pieces it hasn't seen before with the `claude` CLI, and keeps them in
   i18n/<lang>.json, so after editing a page only the changed text is translated again,
3. writes <lang>/<path>/index.html with links pointing at the same language,
4. adds hreflang links and the language switcher to every page, English included,
   and lists the translated pages in sitemap.xml.

    uv run scripts/i18n.py                 # every language
    uv run scripts/i18n.py zh-Hans         # one language
    uv run scripts/i18n.py --no-translate  # rebuild from i18n/*.json only

Run it after changing any English page (and after related_projects.py), then run
markdown_pages.py. To fix a translation, edit i18n/<lang>.json and rerun with
--no-translate. Running it twice changes nothing.
"""

import concurrent.futures
import html
import json
import pathlib
import re
import subprocess
import sys
import threading
from collections import Counter

from bs4 import BeautifulSoup, NavigableString, Tag

ROOT = pathlib.Path(__file__).resolve().parent.parent
SITE = "https://sawirstudio.com"
CACHE = ROOT / "i18n"
MODEL = "claude-sonnet-5-5"
WORKERS = 6

# code: (URL prefix, name in its own language, og:locale, how to write it for the model)
LANGS = {
    "zh-Hans": ("zh-hans", "简体中文", "zh_CN", "Simplified Chinese, as written in mainland China and Singapore. Address the reader as 您"),
    "zh-Hant": ("zh-hant", "繁體中文", "zh_TW", "Traditional Chinese, as written in Taiwan (use Taiwan terminology). Address the reader as 您"),
    "ja": ("ja", "日本語", "ja_JP", "Japanese for business readers in Japan, in polite です/ます style. Address the reader politely (お客様, or leave the subject out as natural Japanese does); use common katakana for tech terms (アプリ, ウェブサイト)"),
    "de": ("de", "Deutsch", "de_DE", "German for business readers in Germany, Austria and Switzerland. Address the reader formally as Sie. Keep common English tech terms Germans use (App, Website, Shop, Plugin, Backend); use German quotation marks „…“"),
}
RTL = {"ar", "he"}

INLINE = {"a", "abbr", "b", "bdi", "bdo", "br", "cite", "code", "data", "dfn", "em", "i", "img", "kbd",
          "mark", "q", "s", "samp", "small", "span", "strong", "sub", "sup", "time", "u", "var", "wbr"}
NO_TEXT = {"script", "style", "pre", "svg", "noscript", "template", "textarea", "code"}
ATTRS = ["title", "alt", "aria-label", "placeholder", "data-sending"]
META = [("name", "description"), ("property", "og:title"), ("property", "og:description"),
        ("name", "twitter:title"), ("name", "twitter:description")]
LD_KEYS = {"name", "headline", "alternativeHeadline", "description", "serviceType", "jobTitle", "text"}

ALT_START, ALT_END = "<!-- i18n:alternates -->", "<!-- /i18n:alternates -->"
SW_START, SW_END = "<!-- i18n:languages -->", "<!-- /i18n:languages -->"
MAP_START, MAP_END = "<!-- i18n -->", "<!-- /i18n -->"


def norm(s):
    return " ".join(s.split())


# ---------------------------------------------------------------- source pages

def source_pages():
    """(path, file) for every English page: path is the URL path, like /shopify/."""
    prefixes = {p for p, *_ in LANGS.values()}
    out = []
    for f in sorted(ROOT.glob("**/index.html")):
        parts = f.relative_to(ROOT).parts
        if any(p.startswith(".") for p in parts) or parts[0] in prefixes or parts[0] == "scripts":
            continue
        out.append(("/" + "/".join(parts[:-1]) + ("/" if len(parts) > 1 else ""), f))
    out.append(("/404.html", ROOT / "404.html"))
    return out


def strip_blocks(text):
    for a, b in ((ALT_START, ALT_END), (SW_START, SW_END)):
        text = re.sub(r"\n?[ \t]*" + re.escape(a) + r".*?" + re.escape(b), "", text, flags=re.S)
    return text


# ---------------------------------------------------------------- extraction

def has_block(el):
    return any(isinstance(d, Tag) and d.name not in INLINE for d in el.descendants)


def skipped(node):
    for p in node.parents:
        if p.name in NO_TEXT or p.get("translate") == "no" or "logo" in (p.get("class") or []):
            return True
    return False


def segments(soup):
    """The HTML pieces to translate, as (node, key). A node is a leaf block (its inner HTML
    is translated as one piece, so links and emphasis stay inside the sentence) or a bare
    text node that sits next to block elements."""
    seen, out = set(), []
    for text in soup.find_all(string=True):
        if type(text) is not NavigableString or not text.strip() or skipped(text):
            continue
        node = text
        while True:
            p = node.parent
            if p is None or p.name in ("body", "html", "head", "[document]") or has_block(p):
                break
            node = p
            if p.name not in INLINE:
                break
        if id(node) in seen:
            continue
        seen.add(id(node))
        if isinstance(node, NavigableString):
            key = norm(html.escape(str(node), quote=False))
        else:
            key = norm(node.decode_contents())
        if re.search(r"[A-Za-z]", re.sub(r"<[^>]+>", "", key)):
            out.append((node, key))
    return out


def ld_strings(obj, found):
    if isinstance(obj, dict):
        for k, v in obj.items():
            if k in LD_KEYS and isinstance(v, str):
                found.append(v)
            else:
                ld_strings(v, found)
    elif isinstance(obj, list):
        for v in obj:
            ld_strings(v, found)


def texts(soup):
    """Plain-text pieces: attributes, meta descriptions, data-search and JSON-LD strings."""
    out = []
    for attr in ATTRS:
        for el in soup.find_all(attrs={attr: True}):
            if not skipped(el):
                out.append(el[attr])
    for k, v in META:
        el = soup.find("meta", attrs={k: v})
        if el and el.get("content"):
            out.append(el["content"])
    for el in soup.find_all(attrs={"data-search": True}):
        out.append(el["data-search"])
    for s in soup.find_all("script", type="application/ld+json"):
        ld_strings(json.loads(s.string), out)
    return [norm(t) for t in out if re.search(r"[A-Za-z]", t)]


def pieces(path, f):
    soup = BeautifulSoup(strip_blocks(f.read_text()), "html.parser")
    return [k for _, k in segments(soup)], texts(soup)


# ---------------------------------------------------------------- translation

PROMPT = """You translate the website of SawirStudio, a small software studio (apps, web platforms, Shopify and Odoo integrations, WordPress plugins), into {lang}.

You get a JSON object: "page" is the page's URL, "html" maps ids to HTML fragments, "text" maps ids to plain text. Reply with one JSON object with the same two keys and the same ids, each value translated. Nothing else, no code fences.

Rules:
- Write natural, plain business copy a native reader would expect from a good local studio. Translate the meaning, not word for word. Keep the tone: direct, friendly, no hype.
- HTML fragments: keep every tag and attribute exactly as given (same tags, same href, class and other attributes, nothing added or dropped). You may move tags to fit the grammar. Translate only the human-readable text between them. Leave text inside <code> as is.
- Plain text: reply with plain text, no HTML and no entities.
- Use the target language's own punctuation. Chinese and Japanese use full-width marks: ，。：；！？（）「」 (and Chinese “” for quotes), never ASCII , : ; ! ? ( ) between CJK characters. Put a space between CJK characters and Latin words or numbers (Shopify 集成, 2019 年).
- Keep names as they are: SawirStudio, Sawir Studio, people's names, company and client names, product and app names (SawirPost, Sawirly, Quill, Par, CleanStorage, Sawir Renew and other Sawir plugins), platforms and tools (Shopify, Odoo, WooCommerce, WordPress, Stripe, Midtrans, iPaymu, App Store, iPhone, GitHub, LinkedIn, Instagram, ChatGPT, Claude, MCP), and code, file names, URLs, versions and numbers.
- Text that is only a name, a URL or code: return it unchanged.
- "text" ids starting with "s" are lists of search keywords: translate them into keywords in the target language, lowercase, space separated.
"""

lock = threading.Lock()


def call_claude(lang_name, page, items_html, items_text):
    payload = json.dumps({"page": page, "html": items_html, "text": items_text}, ensure_ascii=False)
    r = subprocess.run(
        ["claude", "-p", "--model", MODEL, "--tools", "", "--strict-mcp-config", "--setting-sources", "",
         "--no-session-persistence", "--output-format", "json", "--system-prompt", PROMPT.format(lang=lang_name)],
        input=payload, capture_output=True, text=True, timeout=1200, cwd="/tmp",
    )
    out = json.loads(r.stdout)
    if out.get("is_error"):
        raise RuntimeError(out.get("result"))
    body = out["result"].strip()
    body = re.sub(r"^```(?:json)?\s*|\s*```$", "", body)
    return json.loads(body[body.index("{"): body.rindex("}") + 1])


def markup(s):
    tags = sorted(re.findall(r"<\s*/?\s*([a-zA-Z0-9]+)", s))
    attrs = sorted(re.findall(r'\s(href|class|id|src|for|lang|data-[\w-]+)="([^"]*)"', s))
    return tags, attrs


def translate(lang, jobs):
    """jobs: list of (page, html keys, text keys). Fills the cache for lang."""
    lang_name = LANGS[lang][3]
    path = CACHE / f"{lang}.json"
    cache = json.loads(path.read_text()) if path.exists() else {}

    def run(job):
        page, hkeys, tkeys = job
        hkeys = [k for k in dict.fromkeys(hkeys) if k not in cache]
        tkeys = [k for k in dict.fromkeys(tkeys) if k not in cache]
        for attempt in range(3):
            if not hkeys and not tkeys:
                return
            ih = {f"h{i}": k for i, k in enumerate(hkeys)}
            it = {(f"s{i}" if k in search_keys else f"t{i}"): k for i, k in enumerate(tkeys)}
            try:
                res = call_claude(lang_name, page, ih, it)
            except Exception as e:  # noqa: BLE001
                print(f"  {lang} {page}: attempt {attempt + 1} failed: {str(e)[:200]}", flush=True)
                continue
            got = {}
            for i, k in ih.items():
                v = (res.get("html") or {}).get(i)
                if isinstance(v, str) and v.strip() and markup(v) == markup(k):
                    got[k] = norm(v)
            for i, k in it.items():
                v = (res.get("text") or {}).get(i)
                if isinstance(v, str) and v.strip():
                    got[k] = norm(html.unescape(v))
            with lock:
                cache.update(got)
                path.write_text(json.dumps(cache, ensure_ascii=False, indent=1, sort_keys=True) + "\n")
            hkeys = [k for k in hkeys if k not in got]
            tkeys = [k for k in tkeys if k not in got]
        if hkeys or tkeys:
            print(f"  {lang} {page}: {len(hkeys) + len(tkeys)} pieces left in English", flush=True)

    # Split big pages so each call stays small enough to come back whole.
    batches = []
    for page, hkeys, tkeys in jobs:
        hkeys = [k for k in dict.fromkeys(hkeys) if k not in cache]
        tkeys = [k for k in dict.fromkeys(tkeys) if k not in cache]
        for i in range(0, max(len(hkeys), 1), 60):
            batches.append((page, hkeys[i:i + 60], tkeys if i == 0 else []))
    batches = [b for b in batches if b[1] or b[2]]
    if not batches:
        return cache
    print(f"{lang}: translating {sum(len(b[1]) + len(b[2]) for b in batches)} pieces in {len(batches)} calls", flush=True)
    with concurrent.futures.ThreadPoolExecutor(WORKERS) as pool:
        for n, _ in enumerate(pool.map(run, batches), 1):
            if n % 10 == 0:
                print(f"  {lang}: {n}/{len(batches)}", flush=True)
    return cache


search_keys = set()


# ---------------------------------------------------------------- rendering

def lang_url(prefix, href, pages):
    """The same link in the given language, if it's one of our pages."""
    m = re.match(r"^(https://sawirstudio\.com)?(/[^?#]*)(.*)$", href)
    if not m or m[2] not in pages:
        return href
    return (m[1] or "") + f"/{prefix}" + m[2] + m[3]


def ld_translate(obj, cache, prefix, pages):
    if isinstance(obj, dict):
        return {k: (cache.get(norm(v), v) if k in LD_KEYS and isinstance(v, str)
                    else lang_url(prefix, v, pages) if k in ("url", "item") and isinstance(v, str)
                    else ld_translate(v, cache, prefix, pages)) for k, v in obj.items()}
    if isinstance(obj, list):
        return [ld_translate(v, cache, prefix, pages) for v in obj]
    return obj


def fragment(s):
    return list(BeautifulSoup(s, "html.parser").contents)


def render(lang, path, src, cache, pages):
    prefix, _, locale, _ = LANGS[lang]
    soup = BeautifulSoup(src, "html.parser")

    for node, key in segments(soup):
        t = cache.get(key)
        if t is None:
            continue
        if isinstance(node, NavigableString):
            s = str(node)
            lead, trail = s[: len(s) - len(s.lstrip())], s[len(s.rstrip()):]
            new = fragment(t)
            node.replace_with(*([NavigableString(lead)] if lead else []), *new, *([NavigableString(trail)] if trail else []))
        else:
            node.clear()
            for n in fragment(t):
                node.append(n)

    for attr in ATTRS:
        for el in soup.find_all(attrs={attr: True}):
            el[attr] = cache.get(norm(el[attr]), el[attr])
    for k, v in META:
        el = soup.find("meta", attrs={k: v})
        if el and el.get("content"):
            el["content"] = cache.get(norm(el["content"]), el["content"])
    for el in soup.find_all(attrs={"data-search": True}):
        t = cache.get(norm(el["data-search"]))
        if t:
            el["data-search"] = t.lower() + " " + el["data-search"]
    for s in soup.find_all("script", type="application/ld+json"):
        data = ld_translate(json.loads(s.string), cache, prefix, pages)
        s.string = "\n    " + json.dumps(data, ensure_ascii=False, indent=2).replace("\n", "\n    ") + "\n    "

    soup.html["lang"] = lang
    if lang.split("-")[0] in RTL:
        soup.html["dir"] = "rtl"
    for el in soup.find_all(href=True):
        if el.name == "a":
            el["href"] = lang_url(prefix, el["href"], pages)
    for el in soup.find_all("link", rel="canonical"):
        el["href"] = lang_url(prefix, el["href"], pages)
    og = soup.find("meta", property="og:url")
    if og:
        og["content"] = lang_url(prefix, og["content"], pages)
    if not soup.find("meta", property="og:locale"):
        site_name = soup.find("meta", property="og:site_name")
        if site_name:
            tag = soup.new_tag("meta", property="og:locale", content=locale)
            site_name.insert_after(tag)
            site_name.insert_after("\n    ")
    return "<!DOCTYPE html>\n" + str(soup).removeprefix("<!DOCTYPE html>").lstrip("\n")


def add_blocks(text, path, pages, current):
    """hreflang links in <head> and the language list in the footer, for one page."""
    text = strip_blocks(text)
    if path in pages:
        def url(prefix):
            return SITE + (f"/{prefix}" if prefix else "") + path
        links = [f'    <link rel="alternate" hreflang="en" href="{url("")}" />']
        links += [f'    <link rel="alternate" hreflang="{code}" href="{url(p)}" />' for code, (p, *_) in LANGS.items()]
        links += [f'    <link rel="alternate" hreflang="x-default" href="{url("")}" />']
        block = f"    {ALT_START}\n" + "\n".join(links) + f"\n    {ALT_END}\n"
        text = re.sub(r"[ \t]*</head>", lambda m: block + m[0], text, count=1)

    href_path = path if path in pages else "/"
    items = [("en", "", "English")] + [(code, p, name) for code, (p, name, *_) in LANGS.items()]
    lis = []
    for code, prefix, name in items:
        href = (f"/{prefix}" if prefix else "") + href_path
        cur = ' aria-current="true"' if code == current else ""
        lis.append(f'        <li><a href="{href}" hreflang="{code}" lang="{code}"{cur}>{name}</a></li>')
    sw = (f"      {SW_START}\n      <ul class=\"lang-links\" aria-label=\"Language\">\n" + "\n".join(lis)
          + f"\n      </ul>\n      {SW_END}\n")
    return re.sub(r"[ \t]*</footer>", lambda m: sw + m[0], text, count=1)


def sitemap(pages):
    p = ROOT / "sitemap.xml"
    text = re.sub(r"\n?  " + re.escape(MAP_START) + r".*?" + re.escape(MAP_END), "", p.read_text(), flags=re.S)
    english = re.findall(r"<loc>https://sawirstudio\.com([^<]*)</loc>", text)
    urls = [f"  <url><loc>{SITE}/{prefix}{path}</loc></url>" for prefix, *_ in LANGS.values() for path in english if path in pages]
    block = f"  {MAP_START}\n" + "\n".join(urls) + f"\n  {MAP_END}\n"
    p.write_text(text.replace("</urlset>", block + "</urlset>", 1))


# ---------------------------------------------------------------- main

def main():
    args = [a for a in sys.argv[1:] if not a.startswith("--")]
    do_translate = "--no-translate" not in sys.argv
    langs = args or list(LANGS)
    CACHE.mkdir(exist_ok=True)

    src = source_pages()
    pages = {path for path, _ in src if path.endswith("/")}
    extracted = {}
    for path, f in src:
        h, t = pieces(path, f)
        extracted[path] = (h, t)
        soup = BeautifulSoup(strip_blocks(f.read_text()), "html.parser")
        for el in soup.find_all(attrs={"data-search": True}):
            search_keys.add(norm(el["data-search"]))

    for lang in langs:
        if do_translate:
            # Pieces on many pages (header, footer, shared cards) first, in one go, so they
            # read the same everywhere; then each page with what's left.
            counts = Counter(k for h, t in extracted.values() for k in set(h) | set(t))
            common_h = [k for k in counts if counts[k] >= 4 and any(k in h for h, _ in extracted.values())]
            common_t = [k for k in counts if counts[k] >= 4 and k not in common_h]
            translate(lang, [("(shared across pages)", common_h, common_t)])
            translate(lang, [(SITE + path, h, t) for path, (h, t) in extracted.items()])
        path = CACHE / f"{lang}.json"
        cache = json.loads(path.read_text()) if path.exists() else {}

        prefix = LANGS[lang][0]
        missing = 0
        for page, f in src:
            h, t = extracted[page]
            missing += sum(1 for k in h + t if k not in cache)
            out = render(lang, page, strip_blocks(f.read_text()), cache, pages)
            out = add_blocks(out, page, pages, lang)
            dest = ROOT / prefix / (page.lstrip("/") + ("index.html" if page.endswith("/") else ""))
            dest.parent.mkdir(parents=True, exist_ok=True)
            if not dest.exists() or dest.read_text() != out:
                dest.write_text(out)
        print(f"{lang}: {len(src)} pages written to /{prefix}/, {missing} pieces still in English")

    for page, f in src:
        text = f.read_text()
        new = add_blocks(text, page, pages, "en")
        if new != text:
            f.write_text(new)
    sitemap(pages)


if __name__ == "__main__":
    main()
