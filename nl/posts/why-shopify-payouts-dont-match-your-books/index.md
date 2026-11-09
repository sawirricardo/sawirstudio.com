---
title: Waarom Shopify-uitbetalingen nooit kloppen met je boekhouding | SawirStudio
url: https://sawirstudio.com/nl/posts/why-shopify-payouts-dont-match-your-books/
description: Waarom het geld dat Shopify Payments naar je bank stuurt nooit klopt met je verkooprapport: kosten, terugbetalingen, betwiste betalingen, cadeaubonnen, valutaomrekening en timing, en hoe je het sluitend krijgt.
---

# Waarom je Shopify-uitbetalingen nooit kloppen met je boekhouding

Je verkooprapport noemt het ene bedrag. Je bank een ander. Ergens tussenin zitten kosten, terugbetalingen, chargebacks en een paar dagen vertraging, en elke maand is er iemand die een middag in een spreadsheet bezig is om dat uit elkaar te trekken.

Dit kwamen we allemaal tegen bij het schrijven van de boekhoudmodule van onze [Shopify-connector voor Odoo](https://sawirstudio.com/nl/projects/odoo-shopify-connector/). Dit zijn de redenen waarom de bedragen verschillen, en hoe je ze sluitend krijgt.

## Een uitbetaling is een nettobedrag

Shopify Payments stuurt je niet je omzet. Het stuurt wat er overblijft nadat alles is verrekend: verkopen, min terugbetalingen, min transactiekosten, min betwiste betalingen, plus of min correcties. Boek je de uitbetaling als omzet, dan klopt je omzet niet en zie je je kosten niet.

De oplossing is om de onderdelen te boeken, niet het totaal. Verkopen gaan naar omzet, kosten naar een kostenrekening, terugbetalingen tegen de verkopen, en de uitbetaling zelf is alleen geld dat van een tussenrekening naar de bank gaat.

## Een uitbetaling beslaat niet één dag

Een uitbetaling bevat transacties van meerdere dagen, afhankelijk van je uitbetalingsschema, en een bestelling van vrijdag kan in de uitbetaling van dinsdag zitten. Bestellingen op datum aan stortingen koppelen werkt nooit lang. Koppel ze liever per uitbetaling: Shopify legt vast uit welke saldotransacties elke uitbetaling bestaat, en die lijst is de enige betrouwbare schakel tussen een bestelling en het geld op de bank.

## Terugbetalingen komen uit latere uitbetalingen

Als je een bestelling van vorige maand terugbetaalt, komt het geld uit een uitbetaling van deze maand. Je verkooprapport boekt de terugbetaling bij de oorspronkelijke bestelling. Je bank toont vandaag een lagere storting. Allebei kloppen.

## Betwiste betalingen kosten je dubbel

Als een klant een betaling betwist, worden het betwiste bedrag en een chargebackvergoeding van je saldo afgeschreven. Win je de zaak, dan komt het bedrag terug, meestal weken later. Dat zijn drie aparte boekingen voor één bestelling, verspreid over twee of drie uitbetalingen.

## Cadeaubonnen zijn bij verkoop nog geen omzet

Een cadeaubon verkopen betekent geld aannemen dat je later in de vorm van producten terug moet geven. Het hoort in een schuldenrekening tot de bon is gebruikt. Boek je cadeaubonverkopen als omzet, dan tel je dezelfde verkoop dubbel: één keer als de bon wordt gekocht en nog eens als hij wordt besteed.

## Valutaomrekening

Als je in meer dan één valuta verkoopt, verschijnen omrekeningen als eigen saldotransacties, en de koers op de verkoopdag is niet de koers op de dag van de uitbetaling. De 2026-10 API van Shopify heeft een apart transactietype `CURRENCY_CONVERSION` toegevoegd, zodat je de omrekening op een eigen regel kunt boeken in plaats van dat hij verdwijnt in een afrondingsverschil.

## Reserves en correcties

Shopify kan een deel van je saldo vasthouden als reserve, en correcties verschijnen als aanpassingen. Het zijn geen van beide verkopen of kosten, dus ze hebben allebei hun eigen rekeningen nodig.

## Zo krijg je het sluitend

Behandel elke uitbetaling als een kleine journaalpost die is opgebouwd uit de saldotransacties: elke verkoop, terugbetaling, kostenpost, betwiste betaling en correctie op een eigen regel, die allemaal op een tussenrekening landen die de bankstorting vervolgens leegmaakt. Als de regels niet optellen tot de uitbetaling, is er iets verkeerd gekoppeld, en dat wil je weten.

Zo werkt onze boekhoudmodule voor Odoo. Die stelt elke boeking op uit de exacte bedragen van Shopify, en als de boeking niet sluitend is, stopt hij met een duidelijke foutmelding. Hij voegt nooit een “sluitregel” toe om het verschil weg te werken, want zo'n regel is gewoon de fout, opgeschreven op een plek waar niemand kijkt.

Als je Shopify-geld en je boekhouding elke maand van elkaar verschillen, [kunnen we je helpen](https://sawirstudio.com/nl/shopify/).

**Lees ook:** [Shopify koppelen aan Odoo 19 Community, stap voor stap](https://sawirstudio.com/nl/posts/connect-shopify-to-odoo-19-community/), of [De 2026-10 API van Shopify: wat er stukgaat in je integraties](https://sawirstudio.com/nl/posts/shopify-2026-10-api-what-breaks/).
