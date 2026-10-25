# sawirstudio.com

Plain static site: HTML, CSS and a little JavaScript. No build step, no dependencies.

```
index.html                 Home
about-us/index.html        About
projects/index.html        Project list (search + category filter)
projects/<slug>/index.html Project pages
posts/index.html           Blog list
posts/<slug>/index.html    Blog posts
404.html
assets/css/style.css       Site styles
assets/css/convertkit.css  Newsletter form styles (blog posts only)
assets/js/main.js          Google Analytics + project filter
assets/images/             Images
sitemap.xml, robots.txt
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
