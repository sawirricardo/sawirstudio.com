---
title: Shopify Connector voor Odoo | SawirStudio
url: https://sawirstudio.com/nl/projects/odoo-shopify-connector/
description: Een open-source Shopify-connector voor Odoo 19 Community, geschreven door SawirStudio en onderhouden binnen de Odoo Community Association (OCA).
---

# Shopify Connector for Odoo

![De OCA-repository connector-shopify op GitHub](https://sawirstudio.com/assets/images/projects/odoo-shopify-connector.webp)

## Het project

De Odoo Community Association (OCA) is de non-profitorganisatie die de open-sourcemodules onderhoudt waar de meeste Odoo Community-gebruikers op vertrouwen. Connector Shopify is de koppeling tussen Shopify-winkels en Odoo 19 Community, uitgebracht onder de AGPL-3-licentie.

## Wat we hebben gemaakt

Wij schreven de connector en droegen hem in juli 2026 bij aan de OCA: de kern van Shopify Connector, Shopify Connector Accounting en Shopify Connector POS. We blijven eraan werken, voor het laatst door hem af te stemmen op de Admin API 2026-01 van Shopify. Ook andere OCA-bijdragers bouwen er inmiddels op voort, onder meer met een module voor productmedia in de eCommerce.

- **Synchronisatie van je webshop.** Producten, voorraad, klanten, bestellingen, retouren, terugbetalingen en verzending, voor zoveel Shopify-winkels als je wilt per Odoo-database, per bedrijf afgebakend.
- **Betrouwbaar door ontwerp.** Ondertekende webhooks voor snelheid, de queue\_job van de OCA voor duurzame verwerking en geplande afstemming die gemiste of vertraagde events herstelt.
- **Tools voor elke dag.** Onboarding, gezondheidscontroles, een synchronisatiedashboard, foutherstel en instelbare bewaartermijnen voor gegevens.
- **Boekhouding.** Uitbetalingen van Shopify Payments, saldotransacties, kosten, betwistingen, verplichtingen uit cadeaubonnen en afstemming in meerdere valuta.
- **Kassasysteem.** Verkopen via Shopify POS worden als verkooporders geïmporteerd en vanaf elke winkellocatie naar een Odoo-magazijn gerouteerd, zonder dat je de eigen POS-app van Odoo nodig hebt.

## Onze aanpak

Een koppeling die meestal klopt, is voor de boekhouding fout. Daarom praat de connector met de GraphQL Admin API van Shopify, gaat hij er nooit van uit dat een webhook is aangekomen en controleert hij zichzelf volgens een schema tegen Shopify. Journaalposten worden opgebouwd uit de exacte bedragen van Shopify en moeten sluitend zijn voordat ze worden geboekt. Klopt er iets niet, dan geeft de connector een duidelijke foutmelding in plaats van stilletjes een correctieregel toe te voegen.

Wil je dat we het voor jouw winkels instellen of uitbreiden? Bekijk [Shopify-integraties](https://sawirstudio.com/nl/shopify/), of [Odoo & bedrijfssystemen](https://sawirstudio.com/nl/odoo/) voor Odoo zelf.

**Gebouwd met** Python, Odoo 19 en de Shopify GraphQL Admin API

- **Project** Odoo Community Association (OCA)
- **Licentie** AGPL-3
- **Broncode** <https://github.com/OCA/connector-shopify>

## Meer projecten

[![MyVivaldy](https://sawirstudio.com/assets/images/projects/myvivaldy.webp)

### MyVivaldy](https://sawirstudio.com/nl/projects/myvivaldy/)
[![Koperasi HOB Application](https://sawirstudio.com/assets/images/projects/koperasi-hob-application.webp)

### Koperasi HOB Application](https://sawirstudio.com/nl/projects/koperasi-hob-application/)
[![DBTV Application](https://sawirstudio.com/assets/images/projects/dbtv-application.webp)

### DBTV Application](https://sawirstudio.com/nl/projects/dbtv-application/)

Genoeg over ons.

## Waar kunnen we je mee helpen?

[Start een project](https://sawirstudio.com/nl/contact/)
