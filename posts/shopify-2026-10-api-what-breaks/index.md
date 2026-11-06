---
title: Shopify 2026-10 API: What Breaks | SawirStudio
url: https://sawirstudio.com/posts/shopify-2026-10-api-what-breaks/
description: The Shopify Admin API changes in version 2026-10 that can break an integration, and what to change: shipping rates, draft order discounts, taxes, barcodes and more.
---

# Shopify's 2026-10 API: What Breaks in Your Integrations

Shopify released version 2026-10 of its APIs on October 1, 2026. Most of it is new features you can ignore until you need them. A handful of changes can quietly break an integration that worked fine on 2026-07, and those are the ones worth an hour of your time before you bump the version.

This is the list we check against our own integrations, including the [Shopify connector for Odoo](https://sawirstudio.com/projects/odoo-shopify-connector/). It is not every line of the release notes, just the ones that change behavior.

## Changes that break things quietly

### New carrier services don't show rates at checkout on their own

Creating a carrier service with `carrierServiceCreate` (or `POST carrier_services.json` in REST) now only registers the service. It no longer adds its rates to the General shipping profile. If your app or onboarding script creates a carrier service and expects customers to see your rates, they won't, until you add the service to a shipping profile yourself.

### Changing a shipping address recalculates taxes

Updating the shipping address on an unfulfilled order with `orderUpdate`, or the REST order update endpoint, now recalculates its taxes. That is the right behavior, but if you change addresses from an ERP or a support tool, the order total can change under you. Read `taxLines`, `totalTaxSet` and the totals again after the update, and settle any difference in what was paid or refunded.

### Draft order discount warnings lost `priceRule`

`DraftOrderDiscountNotAppliedWarning.priceRule` and the legacy `PriceRule` types are gone. If you build draft orders with `draftOrderCalculate`, `draftOrderCreate` or `draftOrderUpdate` and read why a discount wasn't applied, select `discountTitle` and `discountCode` instead.

### Bad metafield filters now fail loudly

A product query that filters on a metafield that can't be filtered used to ignore that part of the filter and return everything. It now returns an error. That is better, but a sync that has been returning “all products” for months will start failing the day you switch versions. Check that the metafield definitions you filter on support filtering and the operators you use.

### A variant can have more than one barcode

`ProductVariant.barcodes` is a new connection with up to 20 typed barcodes. The old `ProductVariant.barcode` is deprecated and returns only the first one. If your warehouse or point of sale matches products by barcode, read the full list.

### One inventory error is gone

`ITEM_NOT_STOCKED_AT_LOCATION` is removed from the inventory mutations and from `InventorySetQuantitiesUserErrorCode`. Code that branches on it is now dead code; remove it so nobody relies on it later.

## Smaller ones worth a search of your code

- **Customer segments.** `MATCHES` and `NOT MATCHES` replace `= true` and `= false`, and named dates like `12_months_ago` are deprecated in favor of offsets like `-12m`.
- **SMS consent.** The flat SMS consent fields are deprecated. Read `CustomerPhoneNumber.smsMarketingConsent`.
- **Discounts in rollouts.** A discount's `startsAt` and `endsAt` don't tell you if it's live when it's part of a rollout. Reading `Discount.rollouts` needs the `read_rollouts` scope.
- **Fulfillment status.** `FULFILLMENT_NOT_REQUIRED` is a new value. A `switch` that maps every status will hit its default case.
- **Customer Account API.** `Customer.lastIncompleteCheckout` and the `Checkout` types are removed, with no replacement.
- **POS UI extensions.** Cart mutations like `applyCartDiscount` now reject when POS doesn't confirm the change, and `session.currentSession.staffMemberId` is removed.

## Two deadlines that aren't about 2026-10

If you maintain a public app, non-expiring offline access tokens have to be replaced with expiring ones by January 1, 2027. Custom apps and apps a merchant creates in the Dev Dashboard are excluded. And online store script tags stop running on March 1, 2027; we wrote up [what that means for stores](https://sawirstudio.com/posts/shopify-script-tags-stop-march-2027/) separately.

## How we upgrade

We pin the API version in one place, run the integration against a development store on the new version, and read the [2026-10 changelog](https://shopify.dev/changelog?api_version=2026-10) filtered to “Action required”. Version 2026-10 is supported until at least October 1, 2027, so there is time, but not so much that it should wait until the warning emails start.

If your store talks to Odoo or another system and you'd rather someone else did this, that is [what we do](https://sawirstudio.com/shopify/).

**Keep reading:** [Shopify Script Tags Stop Working on March 1, 2027. What Store Owners Should Do](https://sawirstudio.com/posts/shopify-script-tags-stop-march-2027/), or [Connecting Shopify to Odoo 19 Community, Step by Step](https://sawirstudio.com/posts/connect-shopify-to-odoo-19-community/).
