---
title: Sawir Renew – Terugkerende betalingen voor WooCommerce | SawirStudio
url: https://sawirstudio.com/nl/projects/wordpress-plugins/sawir-renew/
description: Complete abonnementen voor WooCommerce: terugkerende betalingen, proefperiodes, aanmeldkosten, pauzeren/overslaan, van abonnement wisselen, betalingsherinneringen en slimme nieuwe pogingen. Allemaal gratis.
---

# Sawir Renew

![Sawir Renew banner](https://sawirstudio.com/assets/images/plugins/sawir-renew-banner.webp)

Sawir Renew maakt van je WooCommerce-webshop een abonnementsbedrijf: terugkerende betalingen, proefperiodes, aanmeldkosten en slimme nieuwe pogingen, zonder dat een betaalde versie functies achterhoudt.

[Download gratis via WordPress.org](https://wordpress.org/plugins/sawir-renew/)

## Wat het doet

Verkoop abonnementen met WooCommerce zonder te betalen voor een extensie. Alle functies zijn gratis: er is geen pro-versie en er zit niets op slot.

### Producten & prijzen

- Simpele en variabele abonnementsproducten
- Factureer elke N dag(en), week/weken, maand(en) of jaar/jaren, met ondersteuning voor actieprijzen
- Gratis proefperiodes, aanmeldkosten en abonnementen die na N betalingen aflopen
- Gesynchroniseerde verlengingen (iedereen verlengt op de 1e, of op een dag naar keuze) met optioneel een naar verhouding berekende of gratis eerste periode
- Eenmalige verzendkosten en de limiet "één actief abonnement per klant"

### Twee manieren van betalen

- **Betalen per verlenging (werkt met elke betaalmethode):** als een verlenging aan de beurt is, krijgt de klant een factuur met een betaallink: bankoverschrijving, e-wallets, rembours, alles wat jouw webshop accepteert. Ideaal voor markten waar automatisch afschrijven van creditcards niet gebruikelijk is.
- **Automatisch betalen via Stripe:** als je de officiële WooCommerce Stripe Gateway gebruikt, worden kaarten bij het afrekenen opgeslagen en worden verlengingen automatisch afgeschreven, zonder dat de klant erbij is. Inclusief slimme nieuwe pogingen en een betaallink als terugvaloptie wanneer een kaart om authenticatie vraagt. Andere betaalproviders kunnen aansluiten op hetzelfde adaptercontract (zie de FAQ).

### Verlengingen waar je 's nachts niet wakker van ligt

- Alle timing draait op Action Scheduler, met een dagelijkse zelfherstellende controle
- Instelbare betalingsherinneringen (dunning): wachttijd tussen nieuwe pogingen, herinneringsschema, respijtperiodes en automatisch opzeggen
- Detectie van staging- en gedupliceerde sites pauzeert de betalingen zodra de URL van je site verandert. Zo voorkom je dubbele afschrijvingen door een vergeten staging-kopie
- E-mails met een melding vóór de verlenging, voordat er automatisch wordt afgeschreven

### Klanten kunnen zichzelf helpen

Via Mijn account → Abonnementen kunnen klanten opzeggen (direct of aan het einde van de periode, dat bepaal jij), pauzeren en hervatten, een verlenging overslaan, eerder verlengen, hun betaalmethode of verzendadres wijzigen en upgraden of downgraden tussen abonnementen, met een naar verhouding berekende bijbetaling bij een upgrade.

### Terugkerende kortingsbonnen

Vier soorten abonnementskortingsbonnen: terugkerende kortingen in % of een vast bedrag (eventueel beperkt tot de eerste N betalingen) en kortingen op aanmeldkosten. Standaardkortingsbonnen blijven alleen gelden voor de eerste bestelling.

### Meldingen

Dertien e-mails die je aan en uit kunt zetten (voor klant en beheerder) voor elk moment in de levenscyclus van een abonnement. Ontwikkelaars kunnen elke gebeurtenis doorsturen naar andere kanalen (WhatsApp, sms, Slack…) via de hook `swsb_notified`.

### Gebouwd voor modern WooCommerce

HPOS-native (abonnementen staan in de bestelopslag van WooCommerce, dus geen eigen tabellen), compatibel met de winkelwagen- en afrekenblokken, een volledige beheerlijst en editor met de eigen bestelschermen van WooCommerce, een rapportagepagina die elk uur wordt gecachet (actief, MRR, churn, aankomende verlengingen, omzet uit verlengingen) en een REST API.

## Vragen

### Is dit gelieerd aan WooCommerce of Automattic?

Nee. "WooCommerce" is een handelsmerk van Automattic. Dit is een onafhankelijke plugin *voor* WooCommerce, volgens hun richtlijnen voor naamgeving.

### Welke betaalproviders ondersteunen automatische verlengingen?

Standaard: de officiële WooCommerce Stripe Gateway. Elke andere betaalprovider werkt in de modus betalen per verlenging: klanten krijgen elke periode een betaallink. Ontwikkelaars kunnen met twee filters automatische ondersteuning toevoegen voor elke betaalprovider:

`swsb_order_supports_auto`: geef true terug als de betaalprovider van de oorspronkelijke bestelling een herbruikbare betaalmethode heeft opgeslagen. `swsb_charge_renewal`: schrijf de verlengingsbestelling af zonder dat de klant erbij is en geef `true` of een `WP_Error` terug.

### Kan het naast WooCommerce Subscriptions draaien?

Nee. Twee betaalsystemen naast elkaar zouden klanten dubbel laten betalen. Daarom schakelt deze plugin zijn eigen systeem uit en toont een melding als WooCommerce Subscriptions actief is.

### Wat gebeurt er op een staging-site?

Als de URL van de site verandert, worden betalingen en e-mails naar klanten automatisch gepauzeerd totdat je bevestigt of de kopie een staging-site of de live site is. Kies je "live", dan worden alle schema's opnieuw opgebouwd.

## Past goed bij

- [Sawir Referral](https://sawirstudio.com/nl/projects/wordpress-plugins/sawir-referral/): laat je klanten hun vrienden meenemen.
- [Sawir Pulse](https://sawirstudio.com/nl/projects/wordpress-plugins/sawir-pulse/): laat bezoekers zien wat er echt gebeurt in je winkel, zoals recente bestellingen, mensen die een product bekijken en nieuwe reviews, als kleine meldingen in de hoek van het scherm.

Bekijk al onze [gratis WordPress-plugins](https://sawirstudio.com/nl/projects/wordpress-plugins/).

- **Prijs** Gratis
- **Versie** 1.0.0
- **Vereist** WordPress 6.0, PHP 7.4, WooCommerce
- **Licentie** GPL-2.0-or-later
- **WordPress.org** <https://wordpress.org/plugins/sawir-renew/>

Genoeg over ons.

## Waar kunnen we je mee helpen?

[Start een project](https://sawirstudio.com/nl/contact/)
