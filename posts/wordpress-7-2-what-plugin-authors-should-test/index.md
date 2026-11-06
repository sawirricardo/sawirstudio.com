---
title: WordPress 7.2: What Plugin Authors Should Test | SawirStudio
url: https://sawirstudio.com/posts/wordpress-7-2-what-plugin-authors-should-test/
description: WordPress 7.2 is due December 8, 2026. What changes for plugin authors, from Application Passwords and the proposed Secrets API to the Interactivity API router, and how to test now.
---

# WordPress 7.2 Lands December 8: What We're Testing in Our Plugins

WordPress 7.2 is scheduled for December 8, 2026. The first release candidate is due November 17, and that's when the feature list freezes and the Field Guide, the list of every developer note, comes out. Until then, betas are the cheapest place to find out your plugin has a problem.

We maintain [12 free plugins on WordPress.org](https://sawirstudio.com/projects/wordpress-plugins/), so this is our own checklist. Some items below were still proposals when we wrote this; the Field Guide is the final word.

## How to test today

- **WP-CLI** on a test site: `wp core update --version=7.2-beta4`, or whichever beta or RC is newest.
- **WordPress Beta Tester plugin**, set to “Bleeding edge” and “Beta/RC Only”.
- **WordPress Playground**, to click through in a browser with nothing installed.

Never on a live site. Turn on `WP_DEBUG` and `WP_DEBUG_LOG` and watch for deprecation notices; that's where most of the early warnings show up.

## What we're checking

### Application Passwords

7.2 tightens Application Passwords: better detection of local and HTTPS setups, an email when one is added, and safer handling of the default role. If your plugin or your customers' tools authenticate to the REST API this way, run that flow end to end.

### The Secrets API

A core way to store credentials, encrypted, so plugins stop keeping API keys as plain options. It was proposed for 7.2, with the admin screen pushed to 7.3. If your plugin stores API keys, watch for it in the Field Guide: once it ships, it's the place those keys should live. None of our free plugins store third-party keys today, so for us this is about client work.

### Sudo mode

Planned: asking a logged-in user to confirm their password again before highly privileged actions. If your plugin performs actions like that from its own screens or over AJAX, test what happens when core asks for a re-check in the middle.

### Interactivity API router

A new router for observing client-side navigation replaces the deprecated `state.navigation.hasStarted` and `hasFinished` pair. There's also a `data-wp-html` directive for rendering trusted markup.

### Admin icons and notices

Dashicons are being replaced with SVG icons in the admin bar and menu, and admin notices get better screen reader announcements. If your plugin styles its menu icon or targets Dashicon classes in admin CSS, look at it in the beta.

### DataViews and the Site Editor

Plugins can register their own screens in the Site Editor and control DataViews fields and layouts from the server. Nothing breaks here, but if your plugin has a list screen, it's worth knowing where things are going.

## After testing

Bump “Tested up to” in your readme once you've actually tested, not before. WordPress.org shows a warning to users when a plugin hasn't been tested with the last major releases, and that warning costs installs. Our plugins say “Tested up to: 7.1” right now; they'll move to 7.2 once they pass on the RC.

The new default theme, Ipsum, ships too. If your plugin outputs anything on the front end, check it there; it's what new sites will look like.

If you have a plugin that needs updating for 7.2, or a custom one nobody maintains anymore, [we can take a look](https://sawirstudio.com/contact/).

**Keep reading:** [Make Your WordPress Plugin Usable by AI Agents with the Abilities API](https://sawirstudio.com/posts/wordpress-abilities-api-for-plugins/), or [WooCommerce or Shopify for an Indonesian Business?](https://sawirstudio.com/posts/woocommerce-or-shopify-indonesia/).
