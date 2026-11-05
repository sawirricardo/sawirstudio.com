---
title: Sawir Renew – Recurring Billing for WooCommerce | SawirStudio
url: https://sawirstudio.com/projects/wordpress-plugins/sawir-renew/
description: Complete subscriptions for WooCommerce — recurring billing, trials, sign-up fees, pause/skip, plan switching, dunning, smart retries. All free.
---

# Sawir Renew

![Sawir Renew banner](https://sawirstudio.com/assets/images/plugins/sawir-renew-banner.webp)

Sawir Renew turns a WooCommerce shop into a subscription business: recurring billing, trials, sign-up fees and smart retries, with no paid tier holding any feature back.

[Get it free on WordPress.org](https://wordpress.org/plugins/sawir-renew/)

## What it does

Sell subscriptions with WooCommerce without paying for an extension. Every feature ships free — there is no pro version and nothing is locked.

### Products & pricing

- Simple and variable subscription product types
- Bill every N day(s), week(s), month(s) or year(s), with sale-price support
- Free trials, sign-up fees, and subscriptions that expire after N payments
- Synchronized renewals (everyone renews on the 1st, or any day you pick) with optional prorated or free first period
- One-time shipping, "one active subscription per customer" limits

### Two payment lanes

- **Pay-per-renewal (works with every gateway):** when a renewal is due, the customer receives an invoice with a pay link — bank transfer, e-wallets, COD, anything your store accepts. Perfect for markets where automatic card billing isn't the norm.
- **Automatic billing via Stripe:** if you use the official WooCommerce Stripe Gateway, cards are saved at checkout and renewals charge automatically off-session — including smart retries and a fallback pay link when a card demands authentication. Other gateways can plug into the same adapter contract (see FAQ).

### Renewals that don't wake you at 3am

- All timing runs on Action Scheduler with a daily self-healing sweep
- Configurable dunning: retry delays, reminder schedule, grace periods, auto-cancel
- Staging/duplicate-site detection pauses billing the moment your site URL changes — no more double charges from a forgotten staging copy
- Pre-renewal notice emails before automatic charges

### Customers can help themselves

My Account → Subscriptions lets customers cancel (immediately or at period end — your choice), pause and resume, skip a renewal, renew early, change their payment method or shipping address, and upgrade/downgrade between plans with prorated upgrade charges.

### Recurring coupons

Four subscription coupon types: recurring % / fixed discounts (optionally limited to the first N payments) and sign-up-fee discounts. Core coupons keep applying to the initial order only.

### Notifications

Thirteen toggleable emails (customer + admin) for every lifecycle moment. Developers can mirror any event to other channels (WhatsApp, SMS, Slack…) via the `swsb_notified` hook.

### Built for modern WooCommerce

HPOS-native (subscriptions live in WooCommerce order storage — no custom tables), cart/checkout blocks compatible, full admin list + editor using WooCommerce's own order screens, an hourly-cached reports page (active, MRR, churn, upcoming renewals, renewal revenue), and a REST API.

## Questions

### Is this affiliated with WooCommerce or Automattic?

No. "WooCommerce" is a trademark of Automattic; this is an independent plugin *for* WooCommerce, per their naming guidelines.

### Which gateways support automatic renewals?

Out of the box: the official WooCommerce Stripe Gateway. Every other gateway works in pay-per-renewal mode — customers get a pay link each period. Developers can add automatic support for any gateway with two filters:

`swsb_order_supports_auto` — return true when the parent order's gateway saved a reusable payment method. `swsb_charge_renewal` — charge the renewal order off-session, return `true` or a `WP_Error`.

### Can it run alongside WooCommerce Subscriptions?

No — running two billing engines would double-charge customers, so this plugin deactivates its engine and shows a notice if WooCommerce Subscriptions is active.

### What happens on a staging site?

If the site URL changes, billing and customer emails pause automatically until you confirm whether the copy is staging or the live site. Choose "live" and every schedule is rebuilt.

## Goes well with

- [Sawir Referral](https://sawirstudio.com/projects/wordpress-plugins/sawir-referral/): lets your customers bring their friends.
- [Sawir Pulse](https://sawirstudio.com/projects/wordpress-plugins/sawir-pulse/): shows shoppers what is really happening in your store, such as recent orders, people viewing a product and fresh reviews, as small notifications in the corner.

See all of our [free WordPress plugins](https://sawirstudio.com/projects/wordpress-plugins/).

- **Price** Free
- **Version** 1.0.0
- **Requires** WordPress 6.0, PHP 7.4, WooCommerce
- **License** GPL-2.0-or-later
- **WordPress.org** <https://wordpress.org/plugins/sawir-renew/>

Enough about us.

## How can we help you?

[Start a project](https://sawirstudio.com/contact/)
