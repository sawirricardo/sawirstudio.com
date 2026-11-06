---
title: Why Shopify Payouts Don't Match Your Books | SawirStudio
url: https://sawirstudio.com/posts/why-shopify-payouts-dont-match-your-books/
description: Why the money Shopify Payments sends to your bank never matches your sales report: fees, refunds, disputes, gift cards, currency conversion and timing, and how to reconcile it.
---

# Why Your Shopify Payouts Never Match Your Books

Your sales report says one number. Your bank says another. Somewhere in between are fees, refunds, chargebacks and a few days of timing, and every month someone spends an afternoon pulling them apart in a spreadsheet.

We dealt with all of this while writing the accounting module of our [Shopify connector for Odoo](https://sawirstudio.com/projects/odoo-shopify-connector/). These are the reasons the numbers differ, and how to make them line up.

## A payout is a net number

Shopify Payments doesn't send you your sales. It sends what's left after it settles everything else: sales, minus refunds, minus processing fees, minus disputes, plus or minus adjustments. If you book the payout as revenue, your revenue is wrong and your fees are invisible.

The fix is to book the parts, not the total. Sales go to revenue, fees to an expense account, refunds against sales, and the payout itself is just money moving from a clearing account to the bank.

## A payout doesn't cover one day

A payout includes transactions from several days, depending on your payout schedule, and a Friday order can arrive in Tuesday's payout. Matching orders to deposits by date never works for long. Match them by payout instead: Shopify records which balance transactions make up each payout, and that list is the only reliable link between an order and the money in the bank.

## Refunds come out of later payouts

When you refund an order from last month, the money comes out of a payout this month. Your sales report puts the refund against the original order. Your bank shows a smaller deposit today. Both are right.

## Disputes take money twice

When a customer disputes a charge, the disputed amount and a chargeback fee are taken from your balance. If you win, the amount comes back, usually weeks later. That's three separate entries for one order, spread over two or three payouts.

## Gift cards aren't revenue when you sell them

Selling a gift card is taking money you owe back as goods later. It belongs in a liability account until the card is used. Book gift card sales as revenue and you count the same sale twice: once when the card is bought, again when it's spent.

## Currency conversion

If you sell in more than one currency, conversions show up as their own balance transactions, and the rate on the day of the sale isn't the rate on the day of the payout. Shopify's 2026-10 API added a separate `CURRENCY_CONVERSION` transaction type, so the conversion can be booked on its own line instead of hiding inside a rounding difference.

## Reserves and adjustments

Shopify can hold part of your balance as a reserve, and corrections show up as adjustments. Neither is a sale or a fee, so both need their own accounts.

## How to make it balance

Treat every payout as a small journal entry built from its balance transactions: each sale, refund, fee, dispute and adjustment on its own line, all landing in a clearing account that the bank deposit then empties. If the lines don't add up to the payout, something is mapped wrong, and that's worth knowing.

That's how our accounting module for Odoo works. It plans each entry from Shopify's exact amounts, and if the entry doesn't balance it stops with a clear error. It never adds a “balancing line” to make the difference go away, because that line is just the mistake, written down where nobody will look.

If your Shopify money and your books disagree every month, [we can help](https://sawirstudio.com/shopify/).

**Keep reading:** [Connecting Shopify to Odoo 19 Community, Step by Step](https://sawirstudio.com/posts/connect-shopify-to-odoo-19-community/), or [Shopify's 2026-10 API: What Breaks in Your Integrations](https://sawirstudio.com/posts/shopify-2026-10-api-what-breaks/).
