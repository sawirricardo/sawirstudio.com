---
title: Shopify Connector for Odoo | SawirStudio
url: https://sawirstudio.com/projects/odoo-shopify-connector/
description: An open-source Shopify connector for Odoo 19 Community, written by SawirStudio and maintained in the Odoo Community Association (OCA).
---

# Shopify Connector for Odoo

![The OCA connector-shopify repository on GitHub](https://sawirstudio.com/assets/images/projects/odoo-shopify-connector.webp)

## The project

The Odoo Community Association (OCA) is the nonprofit that maintains the open-source modules most Odoo Community users rely on. Connector Shopify is its connector between Shopify stores and Odoo 19 Community, released under the AGPL-3 license.

## What we made

We wrote the connector and contributed it to the OCA in July 2026: the core Shopify Connector, Shopify Connector Accounting and Shopify Connector POS. We keep working on it, most recently aligning it with Shopify’s 2026-01 Admin API, and other OCA contributors now build on it too, including a module for eCommerce product media.

- **Commerce sync.** Products, inventory, customers, orders, returns, refunds and fulfillment, for any number of Shopify stores per Odoo database, scoped by company.
- **Reliable by design.** Signed webhooks for speed, OCA’s queue\_job for durable processing, and scheduled reconciliation that repairs missed or delayed events.
- **Day-to-day tools.** Onboarding, health checks, a sync dashboard, error recovery and configurable data retention.
- **Accounting.** Shopify Payments payouts, balance transactions, fees, disputes, gift-card liability and multi-currency reconciliation.
- **Point of sale.** Shopify POS sales imported as sale orders, routed from each retail location to an Odoo warehouse, without needing Odoo’s own POS app.

## How we approached it

An integration that is right most of the time is wrong for accounting. So the connector talks to Shopify’s GraphQL Admin API, never assumes a webhook arrived, and checks itself against Shopify on a schedule. Journal entries are planned from Shopify’s exact amounts and must balance before they post: when something doesn’t add up, the connector raises a clear error instead of quietly adding a balancing line.

Want it set up or extended for your stores? See [Shopify integrations](https://sawirstudio.com/shopify/), or [Odoo & business systems](https://sawirstudio.com/odoo/) for Odoo itself.

**Built with** Python, Odoo 19 and the Shopify GraphQL Admin API

- **Project** Odoo Community Association (OCA)
- **License** AGPL-3
- **Source code** <https://github.com/OCA/connector-shopify>

## More projects

[![MyVivaldy](https://sawirstudio.com/assets/images/projects/myvivaldy.webp)

### MyVivaldy](https://sawirstudio.com/projects/myvivaldy/)
[![Koperasi HOB Application](https://sawirstudio.com/assets/images/projects/koperasi-hob-application.webp)

### Koperasi HOB Application](https://sawirstudio.com/projects/koperasi-hob-application/)
[![DBTV Application](https://sawirstudio.com/assets/images/projects/dbtv-application.webp)

### DBTV Application](https://sawirstudio.com/projects/dbtv-application/)

Enough about us.

## How can we help you?

[Start a project](https://sawirstudio.com/contact/)
