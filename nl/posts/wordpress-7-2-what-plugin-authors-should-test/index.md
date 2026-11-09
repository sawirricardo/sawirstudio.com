---
title: WordPress 7.2: wat pluginontwikkelaars moeten testen | SawirStudio
url: https://sawirstudio.com/nl/posts/wordpress-7-2-what-plugin-authors-should-test/
description: WordPress 7.2 verschijnt op 8 december 2026. Wat verandert er voor pluginontwikkelaars, van Application Passwords en de voorgestelde Secrets API tot de Interactivity API-router, en hoe test je nu al.
---

# WordPress 7.2 verschijnt op 8 december: dit testen we in onze plugins

WordPress 7.2 staat gepland voor 8 december 2026. De eerste release candidate komt op 17 november. Dan wordt de featurelijst bevroren en verschijnt de Field Guide, het overzicht van alle developer notes. Tot die tijd zijn de bètaversies de goedkoopste plek om te ontdekken dat je plugin een probleem heeft.

Wij onderhouden [12 gratis plugins op WordPress.org](https://sawirstudio.com/nl/projects/wordpress-plugins/), dus dit is onze eigen checklist. Sommige punten hieronder waren nog voorstellen toen we dit schreven; de Field Guide heeft het laatste woord.

## Zo test je nu al

- **WP-CLI** op een testsite: `wp core update --version=7.2-beta4`, of welke bèta of RC dan ook het nieuwst is.
- **WordPress Beta Tester plugin**, ingesteld op “Bleeding edge” en “Beta/RC Only”.
- **WordPress Playground**, om in je browser door de omgeving te klikken zonder iets te installeren.

Nooit op een live site. Zet `WP_DEBUG` en `WP_DEBUG_LOG` aan en let op deprecation notices; daar zie je de meeste vroege waarschuwingen.

## Waar we op letten

### Application Passwords

7.2 trekt Application Passwords strakker aan: betere detectie van lokale en HTTPS-omgevingen, een e-mail zodra er een wordt toegevoegd en een veiligere omgang met de standaardrol. Als jouw plugin of de tools van je klanten op deze manier inloggen bij de REST API, test die flow dan van begin tot eind.

### De Secrets API

Een manier in core om inloggegevens versleuteld op te slaan, zodat plugins API-sleutels niet langer als gewone opties bewaren. Hij was voorgesteld voor 7.2, met het beheerscherm uitgesteld tot 7.3. Sla je API-sleutels op in je plugin? Houd de Field Guide dan in de gaten: zodra de Secrets API er is, horen die sleutels daar te staan. Geen van onze gratis plugins slaat nu sleutels van derden op, dus voor ons gaat het vooral om klantwerk.

### Sudo mode

Gepland: een ingelogde gebruiker vragen om het wachtwoord opnieuw te bevestigen vóór acties met hoge rechten. Voert jouw plugin zulke acties uit vanuit eigen schermen of via AJAX? Test dan wat er gebeurt als core halverwege om een nieuwe controle vraagt.

### Interactivity API-router

Een nieuwe router voor het volgen van navigatie aan de clientkant vervangt het verouderde duo `state.navigation.hasStarted` en `hasFinished`. Er is ook een `data-wp-html`-directive om vertrouwde markup te renderen.

### Admin-iconen en meldingen

Dashicons worden vervangen door SVG-iconen in de adminbalk en het menu, en adminmeldingen worden beter voorgelezen door schermlezers. Geef je plugin zijn menu-icoon een eigen stijl, of richt je admin-CSS zich op Dashicon-classes? Bekijk het dan in de bèta.

### DataViews en de Site Editor

Plugins kunnen hun eigen schermen in de Site Editor registreren en de velden en layouts van DataViews vanaf de server aansturen. Hier gaat niets stuk, maar heeft je plugin een lijstscherm, dan is het goed om te weten welke kant het opgaat.

## Na het testen

Verhoog “Tested up to” in je readme pas als je echt getest hebt, niet eerder. WordPress.org toont gebruikers een waarschuwing als een plugin niet getest is met de laatste grote releases, en die waarschuwing kost je installaties. Onze plugins zeggen nu “Tested up to: 7.1”; ze gaan naar 7.2 zodra ze op de RC slagen.

Het nieuwe standaardthema, Ipsum, wordt ook meegeleverd. Toont je plugin iets aan de voorkant van de site, check het dan daar; zo zien nieuwe sites er straks uit.

Heb je een plugin die bijgewerkt moet worden voor 7.2, of een maatwerkplugin die niemand meer onderhoudt? [We kijken er graag naar](https://sawirstudio.com/nl/contact/).

**Lees ook:** [Maak je WordPress-plugin bruikbaar voor AI-agents met de Abilities API](https://sawirstudio.com/nl/posts/wordpress-abilities-api-for-plugins/), of [WooCommerce of Shopify voor een Indonesisch bedrijf?](https://sawirstudio.com/nl/posts/woocommerce-or-shopify-indonesia/).
