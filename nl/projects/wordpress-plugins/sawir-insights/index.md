---
title: Sawir Insights | SawirStudio
url: https://sawirstudio.com/nl/projects/wordpress-plugins/sawir-insights/
description: Heatmaps, sessie-opnames, enquêtes, funnels en formulieranalyse. Alles blijft op je eigen server.
---

# Sawir Insights

![Sawir Insights-banner](https://sawirstudio.com/assets/images/plugins/sawir-insights-banner.webp)

Sawir Insights laat zien hoe bezoekers je site gebruiken, met heatmaps, sessie-opnames, enquêtes, funnels en formulieranalyse, en bewaart alles op je eigen server.

[Download gratis via WordPress.org](https://wordpress.org/plugins/sawir-insights/)

## Wat het doet

Sawir Insights is gedragsanalyse voor WordPress: heatmaps, sessie-opnames, enquêtes, een feedbackwidget, formulieranalyse, conversiefunnels en een dashboard dat alles samenbrengt.

Bij multisite-netwerken houdt elke site zijn eigen analytics bij en stel je die onafhankelijk van elkaar in.

Kortom: een zelfgehoste alternatief voor Hotjar of Microsoft Clarity, waarbij bezoekersdata in je eigen database blijft.

De reden dat het bestaat is simpel. De gevestigde tools in deze categorie zijn goed, maar ze werken door een overzicht van alles wat je bezoekers doen naar andermans servers te sturen. Voor veel sites, zeker in de EU, is dat ergens tussen onprettig en onwerkbaar in. Ik wilde dezelfde soort tools, maar dan met data die de site waar ze verzameld is nooit verlaat. Daarom is deze plugin om die ene voorwaarde heen gebouwd: jouw database, jouw uploadmap, jouw bewaartermijnen. Er zit geen SaaS achter, je hoeft geen account aan te maken en er wordt niets naar buiten verstuurd.

### Wat het doet

- Klik-, beweging- en scrollheatmaps, getoond over je live pagina's en uitgesplitst naar desktop, tablet en mobiel.
- Sessie-opnames die je afspeelt in wp-admin. Invoerwaarden worden altijd gemaskeerd en opnemen staat uit totdat jij het aanzet.
- Pop-upenquêtes en een vast feedbacktabblad, met targeting op URL, apparaat, percentage van het verkeer, tijd op de pagina, scrolldiepte of exit intent.
- Formulieranalyse: waar mensen twijfelen, wat ze leeg laten en waar ze afhaken. Alleen selectors en timing, nooit wat iemand heeft getypt.
- Conversiefunnels met vaste volgorde en dagelijkse rapporten.
- Een dashboard met verkeers-KPI's en vergelijking met de vorige periode, een dagelijkse trendgrafiek en uitsplitsingen naar apparaat, browser, pagina, verwijzer, campagne, funnel en opnames.

### Wat het bewust niet doet

Ik vind dat je moet weten wat een tool weigert te doen voordat je hem installeert.

- Het slaat geen IP-adressen op. Niet gehasht, niet ingekort, helemaal niet.
- Het legt nooit formulierwaarden vast. De opnamelaag stuurt alleen metadata, dus er valt niets te lekken.
- Het laadt niets van een CDN en stuurt er ook niets naartoe. De rrweb-replaybibliotheek zit in de plugin zelf.
- Het bewaart ruwe data niet eeuwig. Events worden samengevoegd en opgeschoond volgens een bewaarschema dat jij bepaalt. Opnames hebben hun eigen bewaartermijn en een schijflimiet, waarbij de oudste als eerste worden verwijderd, want een plugin die ongemerkt je hostingschijf volzet is een slechte plugin.

De frontend-tracker is klein, vanilla JavaScript, verstuurt in batches via `sendBeacon` en heeft instellingen voor sampling. Bezoekers die niet in de opnamesteekproef vallen, downloaden de replaybibliotheek helemaal niet.

### Privacy

Do Not Track wordt standaard gerespecteerd. Een consent-API (`window.swinsConsent`) werkt samen met de consentmanager die je al gebruikt. Sessies, events en opnames zijn gekoppeld aan de eigen export- en verwijdertools van WordPress voor persoonsgegevens. Of jouw specifieke situatie AVG-conform is, hangt af van je privacyverklaring en rechtsgrond. Dat kan geen enkele plugin voor je regelen, maar het werk om data te minimaliseren is gedaan.

### Support

Support loopt via het wordpress.org-forum van deze plugin, of mail naar ricardo@sawirstudio.com.

## Vragen

### Hoe verhoudt het zich tot alternatieven?

Gehoste analysediensten zoals Hotjar en Microsoft Clarity zijn uitstekende tools met royale gratis abonnementen, maar ze werken door een overzicht van het gedrag van je bezoekers naar hun servers te sturen. Sawir Insights ruilt die infrastructuur over meerdere sites in voor één voorwaarde: alles wordt verzameld, opgeslagen en bekeken op je eigen site. Kies een gehoste dienst als je de extra functies wilt en verwerking door derden geen bezwaar is. Kies Sawir Insights als het een dealbreaker is dat bezoekersdata je server verlaat. Opnames en heatmaps staan in je eigen database en uploadmap, dus de bewaarinstellingen zijn hier belangrijker dan bij een SaaS.

### Is Sawir Insights AVG-conform?

De plugin is gebouwd voor privacybewuste toepassingen: data blijft op je server, IP-adressen worden nooit opgeslagen, bewaartermijnen zijn instelbaar, Do Not Track wordt gerespecteerd en de export- en verwijdertools van WordPress zijn geïntegreerd. Of je zelf aan de AVG voldoet, hangt af van je configuratie, privacyverklaring, rechtsgrond en consentimplementatie. Raadpleeg voor je site een gekwalificeerde jurist.

### Wat is het verschil met Hotjar of Microsoft Clarity?

Dat zijn gehoste diensten: sterke tools, maar het gedrag van je bezoekers wordt op hun servers opgeslagen onder hun voorwaarden. Sawir Insights dekt dezelfde categorie, namelijk heatmaps, sessie-opnames, enquêtes, funnels en formulieranalyse, als een plugin die alles op je eigen server houdt, zonder account, zonder externe dienst en zonder data te delen.

### Verlaat bezoekersdata mijn server?

Nee. De plugin stuurt geen gedragsdata naar Sawir Studio of een andere analyticsaanbieder. Events staan in je WordPress-database en opnamebestanden staan in je beschermde uploadmap.

## Past goed bij

- [Sawir Pulse](https://sawirstudio.com/nl/projects/wordpress-plugins/sawir-pulse/): laat bezoekers zien wat er echt gebeurt in je winkel, zoals recente bestellingen, mensen die een product bekijken en nieuwe reviews, als kleine meldingen in de hoek van het scherm.
- [Sawir Live Translate](https://sawirstudio.com/nl/projects/wordpress-plugins/sawir-live-translate/): vertaalt je pagina's op de plek waar ze staan: kies een taal, vertaal automatisch wat zichtbaar is, sla het op en geef bezoekers een taalkeuzemenu.

Bekijk al onze [gratis WordPress-plugins](https://sawirstudio.com/nl/projects/wordpress-plugins/).

- **Prijs** Gratis
- **Versie** 1.0.0
- **Vereist** WordPress 6.2, PHP 7.4
- **Licentie** GPLv2 of nieuwer
- **WordPress.org** <https://wordpress.org/plugins/sawir-insights/>

Genoeg over ons.

## Waar kunnen we je mee helpen?

[Start een project](https://sawirstudio.com/nl/contact/)
