---
title: Shopify Script Tags End March 2027 | SawirStudio
url: https://sawirstudio.com/posts/shopify-script-tags-stop-march-2027/
description: Shopify stops running online store script tags on March 1, 2027. Which apps it affects, how to check your store, and what replaces them.
---

# Shopify Script Tags Stop Working on March 1, 2027. What Store Owners Should Do

On March 1, 2027, Shopify stops injecting script tags into online storefronts. Anything on your store that still loads that way, a review widget, a chat bubble, a tracking pixel, will simply stop appearing. Nothing breaks loudly. Things just go missing.

Shopify announced this in [August 2026](https://shopify.dev/changelog/posts/online-store-script-tags-deprecation). Here's what it means if you run a store.

## What a script tag is

A script tag is a way for an app to add JavaScript to your storefront through the API, without touching your theme. That's why apps liked it, and also the problem: script tags don't show up in the theme editor, so you can't see or switch them off there. Uninstalled apps sometimes left theirs behind.

## The timeline

- **October 1, 2026:** apps can no longer create or update script tags, on any API version. Existing ones keep running.
- **March 1, 2027:** Shopify stops loading online store script tags.

Script tags on the order status page were retired earlier and on a separate schedule, so if something stopped showing there already, that's why.

## How to check your store

Shopify doesn't give merchants a list of script tags, so you have to look:

1. **List the apps that put something on your storefront.** Reviews, chat, pop-ups, upsells, loyalty, analytics, ads pixels. Apps that only work in the admin are not affected.
2. **Look at your theme editor's App embeds.** If an app appears there and is switched on, it is using the newer method and is fine.
3. **Look at the page source.** Open your store, view the source and search for `asyncLoad`. On most stores the script tag URLs are listed in that block. Each domain there usually points to an app.
4. **Ask the app developer.** “Does your app use script tags, and will it work after March 1, 2027?” is a question a maintained app can answer in one line.

Shopify has pushed apps toward theme app extensions for years, so most well-maintained apps are already fine. The risk is in old apps, custom apps someone built for you years ago, and leftovers from apps you no longer use.

## What replaces script tags

- **App embed blocks**, part of a theme app extension, for anything that shows on the page. You switch them on in the theme editor, which means you can finally see them.
- **Web pixels** for scripts that only collect analytics or conversion data. They run sandboxed and need nothing from you once the app supports them.

## If an app won't be updated

Replace it before March, or have its feature rebuilt. For a custom app, the work is usually moving the script into a theme app extension, or into a web pixel if it only tracks. While you're there, delete script tags from apps you uninstalled; the `scriptTagDelete` mutation still works.

If you have a custom Shopify app that needs moving, or you're not sure what's loading on your store, [send us your store address](https://sawirstudio.com/contact/) and we'll tell you what we find. More on our Shopify work is on [this page](https://sawirstudio.com/shopify/).

**Keep reading:** [Shopify's 2026-10 API: What Breaks in Your Integrations](https://sawirstudio.com/posts/shopify-2026-10-api-what-breaks/), or [Connecting Shopify to Odoo 19 Community, Step by Step](https://sawirstudio.com/posts/connect-shopify-to-odoo-19-community/).
