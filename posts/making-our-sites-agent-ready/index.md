---
title: Making Our Sites Agent-Ready | SawirStudio
url: https://sawirstudio.com/posts/making-our-sites-agent-ready/
description: How we made our product sites readable by AI agents: Markdown copies of every page, Link headers, llms.txt, Content Signals in robots.txt, an API catalog and Agent Skills.
---

# What It Took to Make Our Sites Agent-Ready

More and more of the visits to a website aren't people. They're AI agents reading pages to answer someone's question, or trying to use a product on someone's behalf. This week we went through this site and our product sites, [SawirPost](https://sawirstudio.com/projects/sawirpost/), [Sawirly](https://sawirstudio.com/projects/sawirly/), [Rindraw](https://sawirstudio.com/projects/rindraw/), [Math Daily](https://sawirstudio.com/projects/math-daily/) and others, to make them easier for agents to read and use.

We used [isitagentready.com](https://isitagentready.com) as the checklist. It grades a site in levels, and here's what each step took.

## Markdown copies of every page

An agent reading HTML wades through navigation, scripts and footers to find a few paragraphs. So every page on this site now has a Markdown copy next to it, at `index.md`, with only the page's main content. A small script builds them from the HTML, so they never drift.

Agents get them two ways: by adding `index.md` to a URL, or by asking for `Accept: text/markdown`. A tiny Cloudflare Worker in front of the static files checks that header and serves the Markdown copy, with a rough token count in an `x-markdown-tokens` header. Browsers still get the normal page.

```
curl -H "Accept: text/markdown" https://sawirstudio.com/shopify/
```

## Link headers

Every HTML page now says where its Markdown copy is, in an HTTP `Link` header with `rel="alternate"`. The home page also points to `llms.txt` and the sitemap. An agent can find everything from one request without guessing URLs.

## llms.txt

`llms.txt` is a plain-text index of the site for language models: a one-line summary and a list of every page with its description, linking to the Markdown copies. Ours is generated from the sitemap by the same script.

## Content Signals

`robots.txt` says who may crawl, not what they may do with what they read. [Content Signals](https://contentsignals.org/) adds that. Ours is one line, and allows all three uses, because we want to be found and quoted:

```
Content-Signal: ai-train=yes, search=yes, ai-input=yes
```

Those three steps took every site to Level 3.

## An API catalog and Agent Skills

SawirPost and Sawirly have APIs, so they went further. Each publishes an API catalog at `/.well-known/api-catalog` (RFC 9727): a small JSON document pointing to the OpenAPI description and the docs. An agent that lands on the site can find the API without reading a single page.

They also publish Agent Skills, short instructions for agents on how to do common tasks with the product. That took both to Level 4.

## Where we stopped

Level 5 wants agents to be able to sign up and authenticate on their own, through an OAuth agent registration flow. A page describing how to sign up with an email and password doesn't count. That's a real feature to build, with real security questions, so we're thinking it through instead of rushing it.

## Was it worth it?

For a static site, most of this was an afternoon: one script, one small Worker and a line in `robots.txt`. Nothing changed for people visiting the site. We think more questions will be answered by agents reading sites like ours, and a clean Markdown page is the easiest way to be quoted correctly.

We set up the same for clients as part of [Launch & growth](https://sawirstudio.com/launch-and-growth/).

**Keep reading:** [Make Your WordPress Plugin Usable by AI Agents with the Abilities API](https://sawirstudio.com/posts/wordpress-abilities-api-for-plugins/), or [The Throwable App Icons on Our Home Page](https://sawirstudio.com/posts/throwable-app-icons-matter-js/).
