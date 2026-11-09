---
title: Sawir Live Translate | SawirStudio
url: https://sawirstudio.com/nl/projects/wordpress-plugins/sawir-live-translate/
description: Sawir Live Translate vertaalt je pagina's op de plek waar ze staan: kies een taal, vertaal automatisch wat zichtbaar is, sla het op en geef bezoekers een taalkeuze-menu.
---

# Sawir Live Translate

![Sawir Live Translate banner](https://sawirstudio.com/assets/images/plugins/sawir-live-translate-banner.webp)

Sawir Live Translate vertaalt je pagina's op de plek waar ze staan: kies een taal, vertaal automatisch wat zichtbaar is, sla het op en geef bezoekers een taalkeuze-menu.

[Download gratis via WordPress.org](https://wordpress.org/plugins/sawir-live-translate/)

## Wat het doet

Open een willekeurige pagina op je site, klik in de adminbalk op "Translate this page", kies een doeltaal en vertaal alle zichtbare teksten automatisch (via een gratis machinevertaaldienst) of pas de vertalingen met de hand aan terwijl je de pagina live ziet meeveranderen.

Opgeslagen vertalingen worden server-side in de HTML-respons gerenderd (goed voor SEO: geen flits van onvertaalde tekst, en hreflang-alternatieven in de head). Bezoekers wisselen van taal via een linkgebaseerde taalkiezer (rechtsonder).

URL's gebruiken een taalprefix in het pad, bijvoorbeeld /es/about voor Spaans en /about voor de brontaal (die canoniek blijft, zonder prefix). De taal wordt in deze volgorde bepaald: pad met /es/ → ?lang=es → opgeslagen cookie → Accept-Language van de browser → brontaal. Root-URL's en URL's zonder prefix herkennen de browsertaal van de bezoeker automatisch, zonder redirect (er wordt een Vary-header meegestuurd, zodat gedeelde caches goed werken).

1. Ga naar Instellingen → Live Translate: stel je brontaal in en vink de talen aan die je wilt aanbieden. 2. Open een pagina aan de voorkant en klik in de adminbalk op "🌐 Translate this page". 3. Kies een taal, klik op "Auto-translate empty" (of pas teksten zelf aan) en daarna op "Save". 4. Bezoekers gebruiken de taalkiezer (rechtsonder); elke taal heeft een eigen /xx/-URL.

Let op: taalprefixen in het pad gaan uit van een installatie op het hoofddomein en van een site zonder hoofdpagina met een slug die uit een tweeletterige taalcode bestaat.

## Vragen

### Hoe verhoudt het zich tot alternatieven?

Meertalige frameworks zoals WPML, Polylang en TranslatePress behandelen vertaalde content als volwaardig onderdeel van de sitestructuur, met bewerkworkflows per taal, vertaaldiensten en diepe integratie met page builders en WooCommerce. Sawir Live Translate is bewust lichter: open een pagina, vertaal wat zichtbaar is (automatisch of met de hand) en lever het uit via een URL met taalprefix en hreflang-tags. Bouw je een volledig meertalige site met vertaalde menu's, producten en redactionele workflows? Gebruik dan een van die frameworks. Wil je dat bestaande pagina's vanmiddag al in een andere taal te lezen zijn? Dan is dit de plugin die dat doet.

### Ik wissel van taal, maar er verandert niets

Meestal betekent dit dat er voor die taal nog geen vertalingen zijn opgeslagen. Kijk bij Instellingen → Live Translate → Saved translations. Staat de taal er niet tussen (of staat er 0 strings), open dan de pagina-editor opnieuw, vertaal automatisch en klik op Save. De live preview in de editor wordt pas bewaard als je dat doet.

### Sommige teksten blijven onvertaald

De editor op de pagina ziet alleen de teksten van de pagina die je hebt geopend. Teksten die op andere pagina's staan, vragen om vertaling van die pagina's, of om een scan van de hele site. Teksten die na het laden van de pagina door sliders of scripts worden toegevoegd, moet je mogelijk handmatig toevoegen via "Edit all site strings".

### Ik heb mijn content aangepast en de vertaling is verdwenen

Vertalingen horen bij de exacte originele tekst. Herschrijf je een zin, vertaal dan de nieuwe versie één keer en sla die op.

## Past goed bij

- [Sawir Insights](https://sawirstudio.com/nl/projects/wordpress-plugins/sawir-insights/): laat zien hoe bezoekers je site gebruiken, met heatmaps, sessie-opnames, enquêtes, funnels en formulieranalyse, en houdt alles op je eigen server..
- [Sawir Media Folders](https://sawirstudio.com/nl/projects/wordpress-plugins/sawir-media-folders/): brengt mappen naar de WordPress-mediabibliotheek, zoveel als je nodig hebt, met slepen en neerzetten om alles op te bergen..

Bekijk al onze [gratis WordPress-plugins](https://sawirstudio.com/nl/projects/wordpress-plugins/).

- **Prijs** Gratis
- **Versie** 1.1.0
- **Vereist** WordPress 5.8, PHP 7.4
- **Licentie** GPL-2.0-or-later
- **WordPress.org** <https://wordpress.org/plugins/sawir-live-translate/>

Genoeg over ons.

## Waar kunnen we je mee helpen?

[Start een project](https://sawirstudio.com/nl/contact/)
