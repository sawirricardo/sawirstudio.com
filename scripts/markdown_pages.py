# /// script
# requires-python = ">=3.11"
# dependencies = ["beautifulsoup4", "markdownify"]
# ///
"""Write a Markdown copy of every page, and llms.txt, for AI agents.

Each <dir>/index.html gets a <dir>/index.md made from its <main>. worker.js serves
it when a request asks for Accept: text/markdown, so agents get the text without the
markup. llms.txt lists every page in sitemap.xml with its description. Run it after
changing any page's text:

    uv run scripts/markdown_pages.py

It rewrites the files in place, so running it twice changes nothing.
"""

import pathlib
import re

from bs4 import BeautifulSoup, NavigableString
from markdownify import MarkdownConverter

ROOT = pathlib.Path(__file__).resolve().parent.parent
SITE = "https://sawirstudio.com"


def space_out(main):
    """Cards often show a bold number or title and its text as separate elements with no space
    between them in the HTML; add one so the Markdown doesn't run them together ("**9**years")."""
    for el in main.find_all(["strong", "b", "span", "div", "p", "h2", "h3", "h4", "small"]):
        nxt = el.next_sibling
        if nxt is None:
            continue
        starts = str(nxt)[:1] if isinstance(nxt, NavigableString) else "x"
        if starts.isalnum():
            el.insert_after(" ")


class Converter(MarkdownConverter):
    def convert_a(self, el, text, parent_tags):
        href = el.get("href", "")
        if href.startswith("/"):
            el["href"] = SITE + href
        return super().convert_a(el, text, parent_tags)

    def convert_img(self, el, text, parent_tags):
        src = el.get("src", "")
        if src.startswith("/"):
            el["src"] = SITE + src
        # Decorative images (empty alt) say nothing to an agent.
        if not el.get("alt"):
            return ""
        return super().convert_img(el, text, parent_tags)


def page(html_path):
    """(url, title, description, markdown) for one page."""
    soup = BeautifulSoup(html_path.read_text(), "html.parser")
    title = soup.title.get_text(strip=True)
    desc = soup.find("meta", attrs={"name": "description"})
    desc = desc["content"].strip() if desc else ""
    url = soup.find("link", rel="canonical")["href"]

    main = soup.find("main")
    # Drawings, scripts, the breadcrumb trail and form controls aren't content.
    for el in main.select("svg, script, style, canvas, noscript, template, nav.breadcrumbs, form, [aria-hidden=true]"):
        el.decompose()
    space_out(main)
    body = Converter(heading_style="ATX", bullets="-", strip=["button"]).convert_soup(main)
    body = re.sub(r"[ \t]+\n", "\n", body)
    body = re.sub(r"\n{3,}", "\n\n", body).strip()

    front = f"---\ntitle: {title}\nurl: {url}\ndescription: {desc}\n---\n\n"
    return url, title, desc, front + body + "\n"


def llms_txt(pages):
    by_url = {url: (title, desc) for url, title, desc in pages}
    order = re.findall(r"<loc>([^<]+)</loc>", (ROOT / "sitemap.xml").read_text())
    groups = {"Studio": [], "Projects": [], "WordPress plugins": [], "Blog": []}
    for url in order:
        if url not in by_url:
            continue
        path = url.removeprefix(SITE)
        group = (
            "WordPress plugins" if path.startswith("/projects/wordpress-plugins/") and path != "/projects/wordpress-plugins/"
            else "Projects" if path.startswith("/projects/")
            else "Blog" if path.startswith("/posts/")
            else "Studio"
        )
        title, desc = by_url[url]
        name = title.split(" | ")[0]
        groups[group].append(f"- [{name}]({url}index.md): {desc}")

    home_title, home_desc = by_url[SITE + "/"]
    out = [f"# SawirStudio\n\n> {home_desc}\n"]
    out.append(
        "Every page below is also available as Markdown: request it with "
        "`Accept: text/markdown`, or add `index.md` to its URL as linked here.\n"
    )
    for name, lines in groups.items():
        if lines:
            out.append(f"## {name}\n\n" + "\n".join(lines) + "\n")
    return "\n".join(out)


def main():
    pages = []
    for html_path in sorted(ROOT.glob("**/index.html")):
        if any(part.startswith(".") for part in html_path.relative_to(ROOT).parts):
            continue
        url, title, desc, md = page(html_path)
        md_path = html_path.with_name("index.md")
        if not md_path.exists() or md_path.read_text() != md:
            md_path.write_text(md)
        pages.append((url, title, desc))

    txt = llms_txt(pages)
    llms = ROOT / "llms.txt"
    if not llms.exists() or llms.read_text() != txt:
        llms.write_text(txt)
    print(f"{len(pages)} pages, llms.txt")


if __name__ == "__main__":
    main()
