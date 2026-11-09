---
title: Shopify 2026-10 API: wat er breekt | SawirStudio
url: https://sawirstudio.com/nl/posts/shopify-2026-10-api-what-breaks/
description: De wijzigingen in de Shopify Admin API versie 2026-10 die een integratie kunnen breken, en wat je moet aanpassen: verzendtarieven, kortingen op conceptbestellingen, belastingen, barcodes en meer.
---

# Shopify's 2026-10 API: wat er breekt in je integraties

Shopify heeft op 1 oktober 2026 versie 2026-10 van de API's uitgebracht. Het meeste daarvan zijn nieuwe functies die je gerust kunt negeren tot je ze nodig hebt. Een handvol wijzigingen kan een integratie die op 2026-07 prima werkte ongemerkt laten breken, en juist die zijn een uurtje van je tijd waard voordat je de versie omhoog zet.

Dit is de lijst die wij naloopen voor onze eigen integraties, waaronder de [Shopify-connector voor Odoo](https://sawirstudio.com/nl/projects/odoo-shopify-connector/). Het zijn niet alle regels uit de release notes, alleen de wijzigingen die het gedrag veranderen.

## Wijzigingen die ongemerkt dingen breken

### Nieuwe carrier services tonen niet vanzelf verzendtarieven bij het afrekenen

Een carrier service aanmaken met `carrierServiceCreate` (of `POST carrier_services.json` in REST) registreert de service nu alleen nog. De tarieven worden niet meer toegevoegd aan het algemene verzendprofiel (General shipping profile). Maakt jouw app of onboardingscript een carrier service aan en verwacht je dat klanten je tarieven zien? Dan zien ze die niet, totdat je de service zelf aan een verzendprofiel toevoegt.

### Een verzendadres wijzigen herberekent de belasting

Als je het verzendadres van een nog niet verzonden bestelling bijwerkt met `orderUpdate` of het REST-endpoint voor het bijwerken van bestellingen, worden de belastingen nu opnieuw berekend. Dat is het juiste gedrag, maar als je adressen wijzigt vanuit een ERP of een supporttool, kan het totaalbedrag van de bestelling ineens veranderen. Lees `taxLines`, `totalTaxSet` en de totalen na de update opnieuw uit, en regel een eventueel verschil in wat er is betaald of terugbetaald.

### Waarschuwingen over kortingen op conceptbestellingen zijn `priceRule` kwijt

`DraftOrderDiscountNotAppliedWarning.priceRule` en de oude `PriceRule`-types zijn verdwenen. Bouw je conceptbestellingen (draft orders) met `draftOrderCalculate`, `draftOrderCreate` of `draftOrderUpdate` en lees je uit waarom een korting niet is toegepast? Selecteer dan `discountTitle` en `discountCode`.

### Foute metafield-filters geven nu een duidelijke fout

Een productquery die filtert op een metafield dat niet filterbaar is, negeerde die filter vroeger gewoon en gaf alles terug. Nu krijg je een foutmelding. Dat is beter, maar een sync die al maanden “alle producten” teruggeeft, gaat de dag dat je van versie wisselt fouten geven. Controleer of de metafield-definities waarop je filtert filteren ondersteunen, en ook de operators die je gebruikt.

### Een variant kan meer dan één barcode hebben

`ProductVariant.barcodes` is een nieuwe connection met maximaal 20 barcodes met een type. De oude `ProductVariant.barcode` is deprecated en geeft alleen de eerste terug. Koppelt je magazijnsysteem of kassa (POS) producten op barcode? Lees dan de volledige lijst uit.

### Eén voorraadfout is verdwenen

`ITEM_NOT_STOCKED_AT_LOCATION` is verwijderd uit de voorraadmutaties en uit `InventorySetQuantitiesUserErrorCode`. Code die hierop vertakt is nu dode code. Haal die weg, zodat niemand er later op vertrouwt.

## Kleinere wijzigingen waarvoor een zoekopdracht in je code loont

- **Klantsegmenten.** `MATCHES` en `NOT MATCHES` vervangen `= true` en `= false`, en benoemde datums zoals `12_months_ago` zijn deprecated ten gunste van offsets zoals `-12m`.
- **Sms-toestemming.** De losse velden voor sms-toestemming zijn deprecated. Lees `CustomerPhoneNumber.smsMarketingConsent` uit.
- **Kortingen in rollouts.** De `startsAt` en `endsAt` van een korting zeggen niet of die live is als de korting onderdeel is van een rollout. Voor het uitlezen van `Discount.rollouts` heb je de scope `read_rollouts` nodig.
- **Fulfillment-status.** `FULFILLMENT_NOT_REQUIRED` is een nieuwe waarde. Een `switch` die elke status mapt, komt dan in het default-geval terecht.
- **Customer Account API.** `Customer.lastIncompleteCheckout` en de `Checkout`-types zijn verwijderd, zonder vervanging.
- **POS UI extensions.** Winkelwagenmutaties zoals `applyCartDiscount` worden nu afgewezen als POS de wijziging niet bevestigt, en `session.currentSession.staffMemberId` is verwijderd.

## Twee deadlines die niets met 2026-10 te maken hebben

Beheer je een publieke app? Dan moeten offline access tokens die niet verlopen uiterlijk 1 januari 2027 vervangen zijn door tokens die wel verlopen. Custom apps en apps die een merchant zelf aanmaakt in het Dev Dashboard zijn uitgezonderd. En script tags in de online winkel werken vanaf 1 maart 2027 niet meer. Wat dat voor webshops betekent, hebben we apart uitgewerkt in [dit artikel](https://sawirstudio.com/nl/posts/shopify-script-tags-stop-march-2027/).

## Zo upgraden wij

We leggen de API-versie op één plek vast, draaien de integratie op de nieuwe versie tegen een development store en lezen de [changelog van 2026-10](https://shopify.dev/changelog?api_version=2026-10), gefilterd op “Action required”. Versie 2026-10 wordt minstens tot 1 oktober 2027 ondersteund, dus je hebt tijd, maar niet zoveel dat je kunt wachten tot de waarschuwingsmails binnenkomen.

Praat je webshop met Odoo of een ander systeem en laat je dit liever door iemand anders doen? Dat is [wat wij doen](https://sawirstudio.com/nl/shopify/).

**Lees ook:** [Shopify-script tags werken niet meer vanaf 1 maart 2027. Wat webshopeigenaren moeten doen](https://sawirstudio.com/nl/posts/shopify-script-tags-stop-march-2027/), of [Shopify koppelen aan Odoo 19 Community, stap voor stap](https://sawirstudio.com/nl/posts/connect-shopify-to-odoo-19-community/).
