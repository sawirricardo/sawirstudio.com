---
title: Sawir Social Scheduler | SawirStudio
url: https://sawirstudio.com/projects/wordpress-plugins/sawir-social-scheduler/
description: Write once, schedule to 27 networks from wp-admin. Per-channel edits, threads, media, previews. Your tokens stay on your server.
---

# Sawir Social Scheduler

![Sawir Social Scheduler banner](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-banner.webp)

Sawir Social Scheduler lets you write a post once in wp-admin and schedule it to 27 social networks, with per-channel edits, threads, media and previews. Your account tokens stay on your server.

[Get it free on WordPress.org](https://wordpress.org/plugins/sawir-social-scheduler/)

## What it does

Sawir Social Scheduler turns WordPress into a social media scheduler. Connect your channels with the tokens each platform gives you, write a post, pick channels, set a date, and publish. Every send runs in the background and reports back with a link or the platform's own error message.

There is no hosted service in between. Credentials are encrypted with your site's salts and never leave your server.

### Channels

Add as many channels as you like, several accounts per platform. Each channel has a **Test connection** button that fetches the account name and avatar and confirms the token really works before you rely on it.

**Social:** X (Twitter), Facebook Page, Instagram, Threads, LinkedIn Page, Bluesky, Mastodon, TikTok, YouTube, Pinterest, Tumblr, Dribbble, Google Business Profile

**Community:** Telegram, Discord, Slack, Reddit, Lemmy, VK, Kick, Twitch, Whop

**Blog and newsletter:** WordPress, Medium, Dev.to, Hashnode, Listmonk

Group channels by client or brand with nested groups, then select a whole group with one click.

### Posts

- **One body, many channels.** Pick channels with checkboxes. Untouched channels get the post as written.
- **Customize per channel.** Tick "Customize" on any channel to give it its own parts: text, media and delays, prefilled from the post so you edit a copy. Other channels keep the post as written.
- **Threads.** Add parts to post a thread on X, Threads, Bluesky, Mastodon and Telegram, each with its own media. Platforms without threads receive the parts joined.
- **Media.** Attach images and video once. Each platform gets what it supports: albums, carousels, Reels, resumable video uploads.
- **Schedule.** Use the normal WordPress "Publish on" date, or give any channel its own time.
- **Posting times.** Give a channel its usual times of day. "Next free slot" fills the next unused one, the calendar shows empty upcoming slots you can click to fill, and the API accepts `"date": "next"`.
- **Preview.** A live preview per channel that looks like the real platform, with character limits and media rules applied so you see exactly what will go out.
- **Comments and delays.** On platforms without threads, extra parts post as comments under the post. Any part can wait a set number of minutes after the previous one.
- **Share a blog post.** A box on the normal post editor posts to your channels when the post publishes, with its own text and media.
- **Duplicate.** Copy a post with its channels, parts and overrides as a new draft. Keep drafts as templates.
- **Mentions.** Type `@handle` and the network resolves it on X, Instagram, Threads, Mastodon, Telegram, Reddit and Bluesky, where links and handles become real facets. LinkedIn company pages are `@[Name](org:12345)` with the ID from the page URL; Facebook pages are `@[12345]`.
- **Repeat.** Resend evergreen posts every N days or weeks, until a date or a number of times. Each run is its own post with its own sends and stats. Put `{n}` in the text to number runs, since X, Threads and Bluesky reject identical text.

### Calendar

Every post fans out into one send per channel, and the calendar shows each as a card colored by state: scheduled, sent, failed, held, draft, in progress. Drag a card to move that channel's time. Click a card for the error, a Retry button, the live link, and engagement numbers. Click an empty slot to start a post at that time. Month, week, day, and list views, filterable by group and status.

### Analytics

Likes, comments, shares, views and impressions are collected 1 hour, 1 day and 1 week after each send, plus on demand, on 16 platforms. Follower counts are recorded daily per channel with a 30-day sparkline. Numbers appear on the calendar cards, in the post's Analytics box, and in the Channels list. The Analytics page sums it up for the last 7, 30 or 90 days against the period before: posts sent, engagement, reach and followers with their change, a followers chart, engagement by channel, and your top posts.

### Client approval

Tick "Hold sends until approved", copy the preview link, and send it to your client. They see the post as it will look on every channel, with no login, and click Approve or Request changes with a note. Approval releases the sends. Every decision is kept as history.

Large video uploads to X, YouTube and TikTok run in short slices that resume across background jobs, so they survive host time limits and never load a whole file into memory.

### REST API

Everything the admin does is available at `/wp-json/sawir-social/v1` for automations, mobile apps, and other sites. Authenticate with an application password; permissions follow the plugin's capabilities and channel credentials are never returned.

- `GET /platforms`, `GET /channels`, `GET /channels/{id}`, `POST /channels/{id}/test`
- `GET|POST /posts`, `GET|PATCH|DELETE /posts/{id}`, `POST /posts/{id}/duplicate`, `GET /posts/{id}/analytics`
- `GET /sends`, `GET /sends/{id}`, `POST /sends/{id}/retry`, `POST /sends/{id}/stats`

Create a post with `{"title","content","status":"future","date":"2026-12-25 09:00","media":[12],"parts":[{"text","media","delay"}],"channels":[{"id":38,"parts":[{"text","media","delay"}],"date":"…"}],"groups":["Client A"],"hold":false}`. A channel's `parts` replace the post for that channel; `text` is shorthand for a single part.

### AI agents

The same operations are registered with the WordPress Abilities API as `sawir-social/list-channels`, `create-post`, `update-post`, `list-posts`, `list-sends`, `retry-send`, `test-channel`, `get-analytics` and `list-platforms`, with typed input and output schemas. Any agent connected through the WordPress MCP adapter sees them as tools, and they are also callable at `/wp-json/wp-abilities/v1/abilities/{name}/run`. The Abilities API is part of WordPress 6.9 and newer; for older versions the plugin bundles it, so the abilities are available everywhere the plugin runs.

### Requirements

Sends run in the background through Action Scheduler when your site has it (WooCommerce and many other plugins provide it) and through WP-Cron otherwise. On WordPress's built-in cron they only progress when the site gets traffic. For prompt sends, add a system cron, the Calendar screen shows the exact command for your site.

Media must be reachable from the internet for platforms that fetch by URL (Instagram, Threads, Facebook, Pinterest, Telegram). A local development site cannot deliver images to them.

### Credits

Bundles FullCalendar (MIT), the WordPress Abilities API library (GPL), and brand marks from Simple Icons (CC0). Uses Action Scheduler when another plugin provides it.

## Screenshots

![Sawir Social Scheduler screenshot 1](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-screenshot-1.webp)

![Sawir Social Scheduler screenshot 2](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-screenshot-2.webp)

![Sawir Social Scheduler screenshot 3](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-screenshot-3.webp)

![Sawir Social Scheduler screenshot 4](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-screenshot-4.webp)

![Sawir Social Scheduler screenshot 5](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-screenshot-5.webp)

![Sawir Social Scheduler screenshot 6](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-screenshot-6.webp)

## Questions

### Where do I get the tokens?

Each platform's developer settings. The channel form names the exact credential it needs, for example a Facebook Page access token, a Telegram bot token, or a Bluesky app password. Long-lived tokens are best; Meta and LinkedIn tokens expire after about 60 days and need re-pasting.

### My Telegram test says "chat not found"

Private channels have no @username. Leave the Chat ID blank and click Test connection: the plugin lists every chat the bot has been added to, with the numeric ID to paste.

### Why is a send stuck on "In progress…"?

The background queue has not run yet. Visit the site, or set up the system cron shown on the Calendar screen.

### Can I post a video to Instagram from a local site?

No. Instagram fetches media by URL and needs a public address. Everything else about the plugin works locally.

## Goes well with

- [Sawir Insights](https://sawirstudio.com/projects/wordpress-plugins/sawir-insights/): shows how visitors use your site, with heatmaps, session recordings, surveys, funnels and form analytics, and keeps all of it on your own server..
- [Sawir Live Translate](https://sawirstudio.com/projects/wordpress-plugins/sawir-live-translate/): translates your pages where they stand: pick a language, auto-translate what is visible, save it, and give visitors a language switcher..

See all of our [free WordPress plugins](https://sawirstudio.com/projects/wordpress-plugins/).

- **Price** Free
- **Version** 0.1.0
- **Requires** WordPress 5.8, PHP 7.4
- **License** GPL-2.0-or-later
- **WordPress.org** <https://wordpress.org/plugins/sawir-social-scheduler/>

Enough about us.

## How can we help you?

[Start a project](https://sawirstudio.com/contact/)
