"""Rebuild the "More projects" row at the bottom of every project page.

Each row shows the next three projects in projects/index.html that share a
category with the page, as copies of their cards in that list. Run it after
adding, renaming or removing a project:

    python3 scripts/related_projects.py

It rewrites the rows in place, so running it twice changes nothing.
"""

import pathlib
import re

ROOT = pathlib.Path(__file__).resolve().parent.parent
ANCHOR = '      <section class="section" id="cta">'
ROW = re.compile(r'      <section class="section related">\n.*?      </section>\n', re.S)


def cards():
    """The project list's cards, in order, as (slug, categories, card HTML)."""
    html = (ROOT / "projects/index.html").read_text()
    found = []
    for m in re.finditer(r'<a class="card" href="/projects/([^/]+)/" data-categories="([^"]*)".*?</a>', html, re.S):
        card = m[0]
        # The filter's data and the list's page-transition names belong to the list only.
        card = re.sub(r' data-categories="[^"]*"| data-search="[^"]*"', "", card)
        card = re.sub(r' style="view-transition-name: [^"]*"', "", card)
        # Under the row's own "More projects" heading the card titles are h3s.
        card = card.replace("<h2>", "<h3>").replace("</h2>", "</h3>")
        found.append((m[1], set(m[2].split(",")), card))
    return found


def row(picks):
    body = "\n".join("          " + card for _, _, card in picks)
    return (
        '      <section class="section related">\n'
        '        <h2 class="section-title">More projects</h2>\n'
        '        <div class="project-grid">\n'
        f"{body}\n"
        "        </div>\n"
        "      </section>\n"
    )


def main():
    listed = cards()
    for i, (slug, categories, _) in enumerate(listed):
        after = listed[i + 1:] + listed[:i]
        picks = [c for c in after if c[1] & categories][:3]
        picks += [c for c in after if c not in picks][: 3 - len(picks)]

        page = ROOT / f"projects/{slug}/index.html"
        html = ROW.sub("", page.read_text())
        if html.count(ANCHOR) != 1:
            raise SystemExit(f"{page}: expected one closing box ({ANCHOR.strip()})")
        page.write_text(html.replace(ANCHOR, row(picks) + ANCHOR))
        print(f"{slug}: {', '.join(c[0] for c in picks)}")


if __name__ == "__main__":
    main()
