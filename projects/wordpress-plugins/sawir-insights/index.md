---
title: Sawir Insights | SawirStudio
url: https://sawirstudio.com/projects/wordpress-plugins/sawir-insights/
description: Heatmaps, session recordings, surveys, funnels, and form analytics. Everything stays on your server.
---

# Sawir Insights

![Sawir Insights banner](https://sawirstudio.com/assets/images/plugins/sawir-insights-banner.webp)

Sawir Insights shows how visitors use your site, with heatmaps, session recordings, surveys, funnels and form analytics, and keeps all of it on your own server.

[Get it free on WordPress.org](https://wordpress.org/plugins/sawir-insights/)

## What it does

Sawir Insights is behavioral analytics for WordPress: heatmaps, session recordings, surveys, a feedback widget, form analytics, conversion funnels, and a dashboard that ties them together.

On multisite networks, each site tracks and configures its analytics independently.

In short: a self-hosted alternative to Hotjar or Microsoft Clarity, where visitor data stays in your own database.

The reason it exists is simple. The established tools in this category are good, but they work by shipping a record of everything your visitors do to someone else's servers. For a lot of sites — especially in the EU — that is somewhere between uncomfortable and unworkable. I wanted the same class of tooling where the data never leaves the site that collected it. So that is the constraint this plugin is built around: your database, your uploads directory, your retention rules. There is no SaaS behind it, no account to create, and no phone-home of any kind.

### What it does

- Click, movement, and scroll heatmaps, rendered over your live pages, split by desktop, tablet, and mobile.
- Session recordings with playback in wp-admin. Input values are always masked; recording is off until you turn it on.
- Popup surveys and a persistent feedback tab, with targeting by URL, device, traffic percentage, time on page, scroll depth, or exit intent.
- Form analytics: where people hesitate, what they leave empty, where they abandon. Selectors and timing only — never what anyone typed.
- Ordered conversion funnels with daily reports.
- A dashboard with traffic KPIs and period deltas, a daily trend chart, and device, browser, page, referrer, campaign, funnel, and recording breakdowns.

### What it deliberately does not do

I think you should know what a tool refuses to do before you install it.

- It does not store IP addresses. Not hashed, not truncated — at all.
- It does not record form values, ever. The capture layer sends metadata only, so there is nothing to leak.
- It does not load anything from a CDN or send anything to one. The rrweb replay library ships inside the plugin.
- It does not keep raw data forever. Events are aggregated and pruned on a retention schedule you control, and recordings have their own retention and a disk cap with oldest-first eviction, because a plugin that quietly fills your hosting disk is a bad plugin.

The frontend tracker is small vanilla JavaScript, batched over `sendBeacon`, with sampling controls. Visitors who are not sampled into recording never download the replay library at all.

### Privacy posture

Do Not Track is honored by default. A consent-mode API (`window.swinsConsent`) integrates with whatever consent manager you already run. Sessions, events, and recordings are wired into WordPress's own personal-data export and erasure tools. Whether your specific deployment is GDPR compliant depends on your notices and legal basis — that part no plugin can do for you — but the data-minimization work is done.

### Support

Support runs through the wordpress.org forum for this plugin, or email ricardo@sawirstudio.com.

## Questions

### How does it compare to alternatives?

Hosted analytics services like Hotjar and Microsoft Clarity are excellent tools with generous free tiers — and they work by sending a record of your visitors' behavior to their servers. Sawir Insights trades that cross-site infrastructure for one constraint: everything is collected, stored, and viewed on your own site. Choose a hosted service if you want its extras and don't mind third-party processing; choose Sawir Insights if visitor data leaving your server is a dealbreaker. Recordings and heatmaps live in your own database and uploads directory, so the retention settings matter more here than in a SaaS.

### Is Sawir Insights GDPR compliant?

The plugin is built for privacy-conscious deployments: data stays on your server, IP addresses are never stored, retention is configurable, Do Not Track is honored, and WordPress export/erasure tools are integrated. Compliance itself depends on your configuration, notices, legal basis, and consent implementation — consult qualified counsel for your site.

### How is this different from Hotjar or Microsoft Clarity?

Those are hosted services: capable tools, but your visitors' behavior is recorded onto their servers under their terms. Sawir Insights covers the same category — heatmaps, session recordings, surveys, funnels, form analytics — as a plugin that keeps everything on your own server, with no account, no external service, and no data sharing.

### Does visitor data leave my server?

No. The plugin sends no behavioral data to Sawir Studio or any analytics provider. Events live in your WordPress database; recording files live beneath your protected uploads directory.

## Goes well with

- [Sawir Pulse](https://sawirstudio.com/projects/wordpress-plugins/sawir-pulse/): shows shoppers what is really happening in your store, such as recent orders, people viewing a product and fresh reviews, as small notifications in the corner.
- [Sawir Live Translate](https://sawirstudio.com/projects/wordpress-plugins/sawir-live-translate/): translates your pages where they stand: pick a language, auto-translate what is visible, save it, and give visitors a language switcher..

See all of our [free WordPress plugins](https://sawirstudio.com/projects/wordpress-plugins/).

- **Price** Free
- **Version** 1.0.0
- **Requires** WordPress 6.2, PHP 7.4
- **License** GPLv2 or later
- **WordPress.org** <https://wordpress.org/plugins/sawir-insights/>

Enough about us.

## How can we help you?

[Start a project](https://sawirstudio.com/contact/)
