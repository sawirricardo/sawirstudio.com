---
title: Shopify koppelen aan Odoo 19 Community | SawirStudio
url: https://sawirstudio.com/nl/posts/connect-shopify-to-odoo-19-community/
description: Zo koppel je een Shopify-winkel aan Odoo 19 Community met de open source OCA Shopify-connector: de app in het Dev Dashboard, scopes, queue_job, webhooks en onboarding.
---

# Shopify koppelen aan Odoo 19 Community, stap voor stap

We schreven de [Shopify-connector voor Odoo](https://sawirstudio.com/nl/projects/odoo-shopify-connector/) en droegen die bij aan de Odoo Community Association (OCA). Hij is gratis, open source (AGPL-3) en werkt met Odoo 19 Community. Dit is de installatie van begin tot eind, zoals wij die voor klanten doen.

Hij synchroniseert producten, voorraad, klanten, bestellingen, retouren, terugbetalingen en fulfillment, voor één winkel of meerdere. Boekhouding en kassa (point of sale) zijn aparte modules die je later kunt toevoegen.

## Voordat je begint

- Odoo 19 Community die bereikbaar is via openbare HTTPS. Shopify stuurt er webhooks naartoe en haalt er productafbeeldingen op.
- Beheerderstoegang tot de Shopify-winkel, en tot de Shopify-organisatie waar die winkel onder valt.
- Een Shopify-abonnement dat beschermde klantgegevens bevat voor apps die je zelf maakt. Het Basic-abonnement doet dat niet, en zonder die toegang weigert Shopify een aantal bestelwebhooks en mislukken de import van bestellingen en klanten.

## 1. Maak de app aan in het Shopify Dev Dashboard

Sinds 1 januari 2026 kun je de oude "custom apps" niet meer aanmaken vanuit de winkelbeheerder. Maak de app aan in het **Dev Dashboard** van de organisatie waar de winkel onder valt. Dat is belangrijk: Odoo krijgt zijn toegangstoken via de client credentials grant, en die werkt alleen als de app en de winkel bij dezelfde organisatie horen.

Bij het invullen van de appversie:

- **App URL:** `https://shopify.dev/apps/default-app-home`. De connector heeft geen pagina binnen Shopify.
- **Embed app in Shopify admin:** uit.
- **Use legacy install flow:** uit, zodat bij het installeren van de app de scopes worden toegekend die je kiest.
- **Webhooks:** geen. Odoo maakt tijdens de onboarding zelf webhooks aan.
- **Webhooks API version:** `2026-01`, de versie die de connector gebruikt.
- **Scopes:** `read_customers`, `write_customers`, `read_draft_orders`, `read_fulfillments`, `read_inventory`, `write_inventory`, `read_locations`, `read_merchant_managed_fulfillment_orders`, `write_merchant_managed_fulfillment_orders`, `read_orders`, `read_products`, `write_products` en `read_returns`. Voeg `read_all_orders` toe als je bestellingen ouder dan 60 dagen wilt.

Publiceer de versie, installeer de app in de winkel en kopieer de client ID en het client secret van de instellingenpagina van de app.

Heb je al een legacy custom app van vóór 2026? Die werkt nog steeds. Laat de client ID leeg en gebruik in plaats daarvan het Admin API-toegangstoken en de API secret key.

## 2. Installeer de modules in Odoo

Voeg OCA `queue_job` en de [connector-shopify](https://github.com/OCA/connector-shopify)-repository toe aan je addons-pad en installeer daarna **Shopify Connector**. Het werk draait als achtergrondtaken, dus `queue_job` moet serverbreed geladen worden, met een capaciteit per Shopify-kanaal:

```
[options]
server_wide_modules = base,rpc,web,queue_job

[queue_job]
channels = root:4,root.shopify.import:2,root.shopify.export:1,root.shopify.webhook:1
```

Stel daarna de systeemparameter `web.base.url` in op het openbare HTTPS-adres van Odoo en `web.base.url.freeze` op `True`. Anders verandert het adres zodra een beheerder vanaf een andere plek inlogt, en gaan de webhooks van Shopify naar de verkeerde plek.

## 3. Doorloop de onboarding

Geef je team de groep **Shopify Manager**, open daarna **Shopify > Configuration > Instances** en klik op **Start Onboarding**:

1. Vul het domein `yourstore.myshopify.com`, de client ID en het client secret in en test de verbinding. Odoo vraagt het token op en vernieuwt het zelf elke 24 uur.
2. Controleer de scopes. Ontbreekt er een, voeg die dan toe aan een nieuwe appversie, publiceer die en keur hem goed in de winkel.
3. Koppel Shopify-locaties aan Odoo-magazijnen, bekijk de belastingkoppelingen die Odoo genereert en kies het synchronisatiebeleid.
4. Installeer de webhooks en zet de eerste imports in de wachtrij.

Controleer de koppelingen van betaalmethodes en vervoerders voordat je bestellingen en fulfillments zonder toezicht laat draaien. De meeste verrassingen later komen door een koppeling waar niemand naar heeft gekeken.

## 4. Houd het in de gaten

**Shopify > Dashboard** toont de synchronisatiestatus, de achterstand aan taken, recente fouten en de gezondheid van de webhooks. De connector gaat er niet van uit dat elke webhook aankomt: geplande taken importeren alles opnieuw wat in Shopify is gewijzigd en wat Odoo heeft gemist. Gaat er iets mis, los dan de oorzaak op en gebruik daarna **Retry Failed Shopify Bindings**. Voer **Check Webhooks** uit na het wijzigen van domeinen of inloggegevens.

## Volgende stappen

Voeg **Shopify Connector Accounting** toe om uitbetalingen, kosten en geschillen van Shopify Payments als journaalposten binnen te halen (waarom [uitbetalingen lastig te matchen zijn](https://sawirstudio.com/nl/posts/why-shopify-payouts-dont-match-your-books/), leggen we apart uit), en **Shopify Connector POS** voor verkopen in de winkel.

Wil je dat wij het voor je installeren, afstemmen op hoe jouw bedrijf werkt of uitbreiden? [Dat doen we](https://sawirstudio.com/nl/shopify/).

**Lees verder:** [Waarom je Shopify-uitbetalingen nooit kloppen met je boekhouding](https://sawirstudio.com/nl/posts/why-shopify-payouts-dont-match-your-books/), of [De Shopify 2026-10 API: wat er in je integraties kapotgaat](https://sawirstudio.com/nl/posts/shopify-2026-10-api-what-breaks/).
