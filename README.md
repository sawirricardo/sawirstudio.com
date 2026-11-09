# sawirstudio.com

Plain static site: HTML, CSS and a little JavaScript. No build step, no dependencies.

```
index.html                 Home
about-us/index.html        About
contact/index.html         Contact (inquiry form)
projects/index.html        Project list (search + category filter)
projects/<slug>/index.html Project pages
posts/index.html           Blog list
posts/<slug>/index.html    Blog posts
404.html
assets/css/style.css       Site styles
assets/css/convertkit.css  Newsletter form styles (blog posts only)
assets/js/main.js          Google Analytics + project filter
assets/js/vendor/          Matter.js, vendored (physics for the app icons in the home hero; loaded only there)
assets/images/             Images
sitemap.xml, robots.txt
**/index.md, llms.txt    Markdown copies of the pages for AI agents (generated)
scripts/related_projects.py Rebuilds the "More projects" rows (not deployed)
scripts/markdown_pages.py   Rebuilds index.md for every page, and llms.txt (not deployed)
```

## Run locally

```bash
python3 -m http.server 8000
```

Open http://localhost:8000. Pages use root-relative paths (`/assets/...`), so open them through a server, not via `file://`.

## Adding content

Copy an existing page in `projects/` or `posts/` into a new folder, edit it, then link it from the
list page (and the home page for featured projects) and add its URL to `sitemap.xml`.

For a project, set `data-categories` on its card in `projects/index.html` (1 = Design, 2 = Website, 3 = Application)
and put lowercase searchable text in `data-search`.

The picture on a project's card grows into the picture at the top of its page (a view transition). Give the card's
`<img>` and the page's top picture (`.project-hero`, or `.app-icon` for our own apps) the same
`style="view-transition-name: media-<slug>"`, and the home page's picture too if the project is featured there.

Then rebuild the "More projects" row at the bottom of every project page, which copies cards from the list:

```bash
python3 scripts/related_projects.py
```

Run it again after renaming or removing a project. `scripts/` isn't deployed (see `.assetsignore`).

After changing any page's text, or adding a page to `sitemap.xml`, rebuild the Markdown copies and `llms.txt`:

```bash
uv run scripts/markdown_pages.py
```

`worker.js` serves a page's `index.md` to requests with `Accept: text/markdown` (agents), and adds `Link`
headers pointing to it, `llms.txt` and the sitemap. Check with https://isitagentready.com/sawirstudio.com.

## Translations

Every page also exists in Simplified Chinese (`/zh-hans/`), Traditional Chinese (`/zh-hant/`), Japanese (`/ja/`), German (`/de/`), French (`/fr/`), Russian (`/ru/`) and Hebrew (`/he/`, right-to-left), built from
the English pages by `scripts/i18n.py`. Don't edit the files under those folders; edit the English page and rebuild:

```bash
uv run scripts/i18n.py           # translates new or changed text with the claude CLI, then rebuilds
uv run scripts/markdown_pages.py
```

Translations are kept in `i18n/<lang>.json` (English HTML → translation), so only changed text is sent for
translation. To fix a translation, edit it there and run `uv run scripts/i18n.py --no-translate`. The script also
adds the `hreflang` links and the footer language list to every page, and the translated URLs to `sitemap.xml`.
To add a language, add it to `LANGS` in the script.

## Contact form

`contact/index.html` posts to the Google Form "SawirStudio inquiries"
(edit: https://docs.google.com/forms/d/1J5pXtNHc4mB3SXJNIkbOJRf6OHx-ToGfELj3XZoaHl8/edit).
Each field's `name` is that question's `entry.<id>`. If you add or change a question, get the new IDs with:

```bash
gog forms raw 1J5pXtNHc4mB3SXJNIkbOJRf6OHx-ToGfELj3XZoaHl8 \
  | python3 -c 'import json,sys; [print(int(i["questionItem"]["question"]["questionId"],16), i["title"]) for i in json.load(sys.stdin)["items"]]'
```

Read responses with `gog forms responses list 1J5pXtNHc4mB3SXJNIkbOJRf6OHx-ToGfELj3XZoaHl8`.

## Deploy

Hosted on Cloudflare (Pages, now served as a Worker with static assets), configured in `wrangler.jsonc`:

```bash
npx wrangler@4.152.0 deploy
```

This uploads the working tree; anything listed in `.assetsignore` (README, config, `.git`) is not served.
Live at https://sawirstudio.com via the custom domain in `wrangler.jsonc` (Wrangler manages its DNS record).
