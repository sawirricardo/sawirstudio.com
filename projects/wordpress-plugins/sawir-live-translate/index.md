---
title: Sawir Live Translate | SawirStudio
url: https://sawirstudio.com/projects/wordpress-plugins/sawir-live-translate/
description: Sawir Live Translate translates your pages where they stand: pick a language, auto-translate what is visible, save it, and give visitors a language switcher.
---

# Sawir Live Translate

![Sawir Live Translate banner](https://sawirstudio.com/assets/images/plugins/sawir-live-translate-banner.webp)

Sawir Live Translate translates your pages where they stand: pick a language, auto-translate what is visible, save it, and give visitors a language switcher.

[Get it free on WordPress.org](https://wordpress.org/plugins/sawir-live-translate/)

## What it does

Open any page on your site, click "Translate this page" in the admin bar, choose a target language, and either auto-translate every visible string (via a free machine-translation service) or edit translations by hand while seeing the page update live.

Saved translations are rendered server-side into the HTML response (SEO-friendly: no flash of untranslated text, and hreflang alternates in the head), and visitors switch languages via a link-based switcher (bottom-right).

URLs use a language path prefix, e.g. /es/about for Spanish, /about for the source language (which stays canonical, un-prefixed). Language is resolved in this order: /es/ path → ?lang=es → saved cookie → browser Accept-Language → source. Root and prefix-less URLs auto-detect the visitor's browser language without redirecting (a Vary header is sent so shared caches behave).

1. Settings → Live Translate: set your source language and tick the languages to offer. 2. Visit a front-end page and click "🌐 Translate this page" in the admin bar. 3. Pick a language, hit "Auto-translate empty" (or edit strings), then "Save". 4. Visitors use the switcher (bottom-right); each language has its own /xx/ URL.

Note: language path prefixes assume a root-domain install and no top-level page whose slug is a two-letter language code.

## Questions

### How does it compare to alternatives?

Multilingual frameworks like WPML, Polylang, and TranslatePress treat translated content as first-class site structure, with per-language editing workflows, translation services, and deep page-builder and WooCommerce integration. Sawir Live Translate is deliberately lighter: open a page, translate what's visible (automatically or by hand), and serve it from a language-prefixed URL with hreflang tags. If you are building a fully multilingual site with translated menus, products, and editorial workflows, use one of the frameworks. If you want existing pages readable in another language this afternoon, that is what this plugin does.

### I switch languages but nothing changes

Almost always this means no translations are saved for that language yet. Check Settings → Live Translate → Saved translations — if the language isn't listed (or shows 0 strings), open the page editor again, auto-translate, and click Save. The live preview in the editor is not saved until you do.

### Some strings stay untranslated

The in-page editor only sees strings on the page you opened. Text that lives on other pages needs those pages translated too — or run the site-wide scan. Strings injected by sliders/scripts after page load may need a manual entry via "Edit all site strings".

### I edited my content and its translation disappeared

Translations are matched to the exact original text. If you rewrite a sentence, translate the new version once and save.

## Goes well with

- [Sawir Insights](https://sawirstudio.com/projects/wordpress-plugins/sawir-insights/): shows how visitors use your site, with heatmaps, session recordings, surveys, funnels and form analytics, and keeps all of it on your own server..
- [Sawir Media Folders](https://sawirstudio.com/projects/wordpress-plugins/sawir-media-folders/): brings folders to the WordPress media library, as many as you need, with drag and drop to file things away..

See all of our [free WordPress plugins](https://sawirstudio.com/projects/wordpress-plugins/).

- **Price** Free
- **Version** 1.1.0
- **Requires** WordPress 5.8, PHP 7.4
- **License** GPL-2.0-or-later
- **WordPress.org** <https://wordpress.org/plugins/sawir-live-translate/>

Enough about us.

## How can we help you?

[Start a project](https://sawirstudio.com/contact/)
