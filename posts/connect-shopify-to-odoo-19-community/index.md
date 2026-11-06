---
title: Connect Shopify to Odoo 19 Community | SawirStudio
url: https://sawirstudio.com/posts/connect-shopify-to-odoo-19-community/
description: How to connect a Shopify store to Odoo 19 Community with the open-source OCA Shopify connector: the Dev Dashboard app, scopes, queue_job, webhooks and onboarding.
---

# Connecting Shopify to Odoo 19 Community, Step by Step

We wrote the [Shopify connector for Odoo](https://sawirstudio.com/projects/odoo-shopify-connector/) and contributed it to the Odoo Community Association (OCA). It is free, open source (AGPL-3) and works with Odoo 19 Community. This is the setup, start to finish, the way we do it for clients.

It syncs products, stock, customers, orders, returns, refunds and fulfillment, for one store or several. Accounting and point of sale are separate modules you can add later.

## Before you start

- Odoo 19 Community reachable over public HTTPS. Shopify sends webhooks to it and downloads product images from it.
- Admin access to the Shopify store, and to the Shopify organization that owns it.
- A Shopify plan that includes protected customer data for apps you create yourself. The Basic plan doesn't, and without it Shopify refuses some order webhooks and the order and customer imports fail.

## 1. Create the app in the Shopify Dev Dashboard

Since January 1, 2026, Shopify no longer lets you create the old “custom apps” from the store admin. Create the app in the **Dev Dashboard** of the organization that owns the store. This matters: Odoo gets its access token with the client credentials grant, which only works when the app and the store belong to the same organization.

When you fill in the app version:

- **App URL:** `https://shopify.dev/apps/default-app-home`. The connector has no page inside Shopify.
- **Embed app in Shopify admin:** off.
- **Use legacy install flow:** off, so installing the app grants the scopes you pick.
- **Webhooks:** none. Odoo creates its own during onboarding.
- **Webhooks API version:** `2026-01`, the version the connector uses.
- **Scopes:** `read_customers`, `write_customers`, `read_draft_orders`, `read_fulfillments`, `read_inventory`, `write_inventory`, `read_locations`, `read_merchant_managed_fulfillment_orders`, `write_merchant_managed_fulfillment_orders`, `read_orders`, `read_products`, `write_products` and `read_returns`. Add `read_all_orders` if you want orders older than 60 days.

Release the version, install the app on the store, and copy the client ID and client secret from the app's Settings page.

Already have a legacy custom app from before 2026? It still works. Leave the client ID empty and use its Admin API access token and API secret key instead.

## 2. Install the modules in Odoo

Add OCA `queue_job` and the [connector-shopify](https://github.com/OCA/connector-shopify) repository to your addons path, then install **Shopify Connector**. The work runs as background jobs, so `queue_job` has to be loaded server-wide, with a capacity for each Shopify channel:

```
[options]
server_wide_modules = base,rpc,web,queue_job

[queue_job]
channels = root:4,root.shopify.import:2,root.shopify.export:1,root.shopify.webhook:1
```

Then set the `web.base.url` system parameter to Odoo's public HTTPS address and `web.base.url.freeze` to `True`. Otherwise the address changes when an admin logs in from somewhere else, and Shopify's webhooks go to the wrong place.

## 3. Run onboarding

Give your team the **Shopify Manager** group, then open **Shopify > Configuration > Instances** and click **Start Onboarding**:

1. Enter the `yourstore.myshopify.com` domain, client ID and client secret, and test the connection. Odoo requests the token and renews it every 24 hours on its own.
2. Check the scopes. If one is missing, add it to a new app version, release it and approve it on the store.
3. Map Shopify locations to Odoo warehouses, review the tax mappings Odoo generates, and choose the sync policies.
4. Install the webhooks and queue the first imports.

Before you let orders and fulfillments run unattended, review the payment gateway and carrier mappings. Most surprises later come from a mapping nobody looked at.

## 4. Keep an eye on it

**Shopify > Dashboard** shows sync states, the job backlog, recent errors and webhook health. The connector doesn't assume every webhook arrives: scheduled jobs import again anything Shopify changed that Odoo missed. When something fails, fix the cause, then use **Retry Failed Shopify Bindings**. Run **Check Webhooks** after changing domains or credentials.

## Next steps

Add **Shopify Connector Accounting** to bring in Shopify Payments payouts, fees and disputes as journal entries (we explain [why payouts are hard to match](https://sawirstudio.com/posts/why-shopify-payouts-dont-match-your-books/) separately), and **Shopify Connector POS** for in-store sales.

If you'd like it set up, configured for how your business works or extended, [we do that](https://sawirstudio.com/shopify/).

**Keep reading:** [Why Your Shopify Payouts Never Match Your Books](https://sawirstudio.com/posts/why-shopify-payouts-dont-match-your-books/), or [Shopify's 2026-10 API: What Breaks in Your Integrations](https://sawirstudio.com/posts/shopify-2026-10-api-what-breaks/).
