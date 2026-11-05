---
title: Sawir Min Max Quantities for WooCommerce | SawirStudio
url: https://sawirstudio.com/projects/wordpress-plugins/sawir-min-max-quantities-for-woocommerce/
description: Set minimum and maximum purchase quantity rules for WooCommerce products, product categories, and the whole order.
---

# Sawir Min Max Quantities

![Sawir Min Max Quantities banner](https://sawirstudio.com/assets/images/plugins/sawir-min-max-quantities-for-woocommerce-banner.webp)

Sawir Min Max Quantities sets how few or how many of something a customer can buy, for one product, a whole category or the entire order.

[Get it free on WordPress.org](https://wordpress.org/plugins/sawir-min-max-quantities-for-woocommerce/)

## What it does

Sawir Min Max Quantities for WooCommerce lets store owners control how many items customers can buy. Set minimum, maximum, and "group of" multiple quantities per product, per variation, per product category, and for the total number of items in an order. Rules are enforced on product page quantity inputs, when adding to cart, and again on the cart and checkout pages.

## Features

- Adds minimum, maximum, and group-of quantity fields to the product Inventory tab.
- Adds optional minimum, maximum, and group-of overrides to each variation.
- Adds minimum, maximum, and group-of quantity fields to product categories.
- Lets category rules apply to the combined cart quantity for that category instead of as a per-product fallback.
- Adds order-level minimum and maximum item limits under WooCommerce > Settings > Products > General.
- Applies a clear precedence: a variation rule wins; otherwise a product rule wins; otherwise the most restrictive unchecked category fallback applies.
- Variations inherit empty fields from their parent product or category fallback.
- Enforces rules on product page and variation quantity inputs, including step values for group-of multiples.
- Blocks add-to-cart actions that would exceed a product maximum.
- Validates all rules on the cart and checkout pages with clear error notices, including multiples such as "must be purchased in multiples of N."
- Leaves quantity inputs on the cart page free to reach zero so items can still be removed.
- Requires no configuration to install; rules are optional everywhere.
- Cleans up all plugin options and metadata when the plugin is uninstalled.

## Questions

### How does it compare to alternatives?

The official WooCommerce Min/Max Quantities extension and larger commercial suites cover this ground too, often alongside role-based and price-based rules. Sawir Min Max Quantities focuses on the common cases for free: minimum, maximum, and "group of" multiples per product, per variation, per category, and per order, enforced on the product page, at add-to-cart, and again at checkout. If your quantity rules depend on who the customer is — wholesale roles, customer groups — a commercial suite is the better fit.

### Does this require WooCommerce?

Yes. This plugin is built for WooCommerce products, cart, and checkout hooks.

### What happens when a product has both its own rule and a category rule?

The product rule wins. Category rules only apply to products that have no rule of their own.

### Can variations have their own rules?

Yes. Each variation can set its own minimum, maximum, and group-of values. Empty variation fields inherit from the parent product or category fallback.

## Goes well with

- [Sawir Restock Notifier](https://sawirstudio.com/projects/wordpress-plugins/sawir-restock-notifier-for-woocommerce/): lets customers leave their email on an out-of-stock product and tells them, once, when it is back..
- [Sawir Renew](https://sawirstudio.com/projects/wordpress-plugins/sawir-renew/): turns a WooCommerce shop into a subscription business: recurring billing, trials, sign-up fees and smart retries, with no paid tier holding any feature back..

See all of our [free WordPress plugins](https://sawirstudio.com/projects/wordpress-plugins/).

- **Price** Free
- **Version** 1.0.1
- **Requires** WordPress 5.8, PHP 7.4, WooCommerce
- **License** GPLv2 or later
- **WordPress.org** <https://wordpress.org/plugins/sawir-min-max-quantities-for-woocommerce/>

Enough about us.

## How can we help you?

[Start a project](https://sawirstudio.com/contact/)
