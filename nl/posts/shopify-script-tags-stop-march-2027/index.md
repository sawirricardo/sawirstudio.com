---
title: Shopify Script Tags stoppen in maart 2027 | SawirStudio
url: https://sawirstudio.com/nl/posts/shopify-script-tags-stop-march-2027/
description: Shopify stopt op 1 maart 2027 met het uitvoeren van script tags in de online winkel. Welke apps dit raakt, hoe je je webshop controleert en wat ze vervangt.
---

# Shopify Script Tags werken niet meer vanaf 1 maart 2027. Wat webshopeigenaren nu moeten doen

Op 1 maart 2027 stopt Shopify met het injecteren van script tags in online storefronts. Alles in je webshop dat nog zo wordt geladen, zoals een reviewwidget, een chatbubbel of een trackingpixel, verdwijnt gewoon. Er gaat niets zichtbaar kapot. Dingen ontbreken ineens.

Shopify kondigde dit aan in [augustus 2026](https://shopify.dev/changelog/posts/online-store-script-tags-deprecation). Dit betekent het voor jou als je een webshop hebt.

## Wat een script tag is

Een script tag is een manier waarop een app via de API JavaScript aan je storefront toevoegt, zonder je thema aan te raken. Daarom gebruikten apps het graag, en dat is ook meteen het probleem: script tags verschijnen niet in de thema-editor, dus je kunt ze daar niet zien of uitzetten. Verwijderde apps lieten er soms nog een achter.

## De planning

- **1 oktober 2026:** apps kunnen geen script tags meer aanmaken of bijwerken, in welke API-versie dan ook. Bestaande script tags blijven werken.
- **1 maart 2027:** Shopify stopt met het laden van script tags in de online winkel.

Script tags op de orderstatuspagina zijn eerder en volgens een apart schema uitgefaseerd. Als daar al iets niet meer verscheen, ligt het daaraan.

## Zo controleer je je webshop

Shopify geeft verkopers geen lijst met script tags, dus je moet zelf zoeken:

1. **Maak een lijst van de apps die iets in je storefront zetten.** Reviews, chat, pop-ups, upsells, loyaliteitsprogramma's, analytics, advertentiepixels. Apps die alleen in het beheerpaneel werken, worden niet geraakt.
2. **Kijk bij de App-embeds in je thema-editor.** Staat een app daar en is hij ingeschakeld, dan gebruikt hij de nieuwere methode en zit je goed.
3. **Bekijk de paginabron.** Open je webshop, bekijk de broncode en zoek op `asyncLoad`. Op de meeste webshops staan de URL's van de script tags in dat blok. Elk domein daar verwijst meestal naar een app.
4. **Vraag het de app-ontwikkelaar.** “Gebruikt jullie app script tags, en werkt hij ook na 1 maart 2027?” is een vraag die een onderhouden app in één regel kan beantwoorden.

Shopify stuurt apps al jaren richting theme app extensions, dus de meeste goed onderhouden apps zijn al in orde. Het risico zit bij oude apps, maatwerk-apps die iemand jaren geleden voor je bouwde en restanten van apps die je niet meer gebruikt.

## Wat script tags vervangt

- **App embed blocks**, onderdeel van een theme app extension, voor alles wat op de pagina zichtbaar is. Je zet ze aan in de thema-editor, dus je kunt ze eindelijk zien.
- **Web pixels** voor scripts die alleen analytics- of conversiedata verzamelen. Ze draaien in een sandbox en vragen niets van jou zodra de app ze ondersteunt.

## Als een app niet wordt bijgewerkt

Vervang hem vóór maart, of laat de functie opnieuw bouwen. Bij een maatwerk-app komt het werk meestal neer op het verplaatsen van het script naar een theme app extension, of naar een web pixel als het alleen trackt. Verwijder meteen ook de script tags van apps die je hebt gedeïnstalleerd; de mutation `scriptTagDelete` werkt nog steeds.

Heb je een maatwerk-Shopify-app die moet worden aangepast, of weet je niet zeker wat er in je webshop wordt geladen? [Stuur ons het adres van je webshop](https://sawirstudio.com/nl/contact/) en we vertellen je wat we vinden. Meer over ons Shopify-werk vind je op [deze pagina](https://sawirstudio.com/nl/shopify/).

**Lees ook:** [Shopify's 2026-10 API: wat er in je integraties kapotgaat](https://sawirstudio.com/nl/posts/shopify-2026-10-api-what-breaks/), of [Shopify koppelen aan Odoo 19 Community, stap voor stap](https://sawirstudio.com/nl/posts/connect-shopify-to-odoo-19-community/).
