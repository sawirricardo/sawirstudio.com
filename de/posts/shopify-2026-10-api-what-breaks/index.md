---
title: Shopify 2026-10 API: Was kaputtgeht | SawirStudio
url: https://sawirstudio.com/de/posts/shopify-2026-10-api-what-breaks/
description: Die Änderungen der Shopify Admin API in Version 2026-10, die eine Integration kaputtmachen können, und was Sie anpassen müssen: Versandraten, Rabatte in Draft Orders, Steuern, Barcodes und mehr.
---

# Shopify 2026-10 API: Was in Ihren Integrationen kaputtgehen kann

Shopify hat die Version 2026-10 seiner APIs am 1. Oktober 2026 veröffentlicht. Das meiste sind neue Funktionen, die Sie ignorieren können, bis Sie sie brauchen. Eine Handvoll Änderungen kann jedoch eine Integration, die mit 2026-07 problemlos lief, unbemerkt zum Absturz bringen. Für diese lohnt sich eine Stunde Ihrer Zeit, bevor Sie die Version hochsetzen.

Das ist die Liste, die wir bei unseren eigenen Integrationen abgleichen, darunter dem [Shopify-Connector für Odoo](https://sawirstudio.com/de/projects/odoo-shopify-connector/). Es sind nicht alle Zeilen der Release Notes, sondern nur die, die das Verhalten ändern.

## Änderungen, die unbemerkt etwas kaputtmachen

### Neue Carrier Services zeigen ihre Versandraten nicht mehr von selbst im Checkout an

Wenn Sie einen Carrier Service mit `carrierServiceCreate` (oder `POST carrier_services.json` in REST) anlegen, wird der Service jetzt nur noch registriert. Seine Versandraten werden nicht mehr automatisch zum allgemeinen Versandprofil hinzugefügt. Wenn Ihre App oder Ihr Onboarding-Skript einen Carrier Service anlegt und erwartet, dass Kunden Ihre Versandraten sehen, passiert das nicht, solange Sie den Service nicht selbst einem Versandprofil hinzufügen.

### Eine geänderte Lieferadresse berechnet die Steuern neu

Wenn Sie die Lieferadresse einer noch nicht erfüllten Bestellung mit `orderUpdate` oder dem REST-Endpunkt zum Aktualisieren von Bestellungen ändern, werden jetzt deren Steuern neu berechnet. Das ist das richtige Verhalten. Wenn Sie Adressen aber aus einem ERP oder einem Support-Tool ändern, kann sich der Bestellbetrag unter Ihren Händen verändern. Lesen Sie nach dem Update `taxLines`, `totalTaxSet` und die Summen erneut aus und gleichen Sie jede Differenz bei bezahlten oder erstatteten Beträgen aus.

### Warnungen zu Rabatten in Draft Orders haben `priceRule` verloren

`DraftOrderDiscountNotAppliedWarning.priceRule` und die älteren `PriceRule`-Typen gibt es nicht mehr. Wenn Sie Draft Orders mit `draftOrderCalculate`, `draftOrderCreate` oder `draftOrderUpdate` erstellen und auslesen, warum ein Rabatt nicht angewendet wurde, wählen Sie stattdessen `discountTitle` und `discountCode` aus.

### Fehlerhafte Metafield-Filter schlagen jetzt deutlich fehl

Eine Produktabfrage, die nach einem nicht filterbaren Metafield filtert, hat diesen Teil des Filters früher ignoriert und alles zurückgegeben. Jetzt gibt sie einen Fehler zurück. Das ist besser, aber ein Sync, der seit Monaten „alle Produkte“ zurückgibt, schlägt an dem Tag fehl, an dem Sie die Version wechseln. Prüfen Sie, ob die Metafield-Definitionen, nach denen Sie filtern, das Filtern und die von Ihnen verwendeten Operatoren unterstützen.

### Eine Variante kann mehr als einen Barcode haben

`ProductVariant.barcodes` ist eine neue Connection mit bis zu 20 typisierten Barcodes. Das alte `ProductVariant.barcode` ist veraltet und gibt nur den ersten zurück. Wenn Ihr Lager oder Ihr Kassensystem Produkte anhand des Barcodes zuordnet, lesen Sie die vollständige Liste aus.

### Ein Inventar-Fehler ist entfallen

`ITEM_NOT_STOCKED_AT_LOCATION` wurde aus den Inventar-Mutationen und aus `InventorySetQuantitiesUserErrorCode` entfernt. Code, der darauf verzweigt, ist jetzt toter Code. Entfernen Sie ihn, damit sich später niemand darauf verlässt.

## Kleinere Änderungen, nach denen sich eine Suche in Ihrem Code lohnt

- **Kundensegmente.** `MATCHES` und `NOT MATCHES` ersetzen `= true` und `= false`. Benannte Datumsangaben wie `12_months_ago` sind veraltet, stattdessen gibt es Offsets wie `-12m`.
- **SMS-Einwilligung.** Die einfachen SMS-Einwilligungsfelder sind veraltet. Lesen Sie `CustomerPhoneNumber.smsMarketingConsent` aus.
- **Rabatte in Rollouts.** `startsAt` und `endsAt` eines Rabatts verraten Ihnen nicht, ob er aktiv ist, wenn er Teil eines Rollouts ist. Für das Auslesen von `Discount.rollouts` benötigen Sie den Scope `read_rollouts`.
- **Fulfillment-Status.** `FULFILLMENT_NOT_REQUIRED` ist ein neuer Wert. Ein `switch`, der jeden Status abbildet, landet im Default-Fall.
- **Customer Account API.** `Customer.lastIncompleteCheckout` und die `Checkout`-Typen wurden ohne Ersatz entfernt.
- **POS-UI-Extensions.** Warenkorb-Mutationen wie `applyCartDiscount` werden jetzt abgelehnt, wenn POS die Änderung nicht bestätigt, und `session.currentSession.staffMemberId` wurde entfernt.

## Zwei Fristen, die nichts mit 2026-10 zu tun haben

Wenn Sie eine öffentliche App betreiben, müssen nicht ablaufende Offline-Zugriffstoken bis zum 1. Januar 2027 durch ablaufende ersetzt werden. Custom Apps und Apps, die ein Händler im Dev Dashboard erstellt, sind ausgenommen. Außerdem werden Script Tags im Online-Shop ab dem 1. März 2027 nicht mehr ausgeführt. Was das für Shops bedeutet, haben wir [separat beschrieben](https://sawirstudio.com/de/posts/shopify-script-tags-stop-march-2027/).

## So führen wir Upgrades durch

Wir legen die API-Version an einer einzigen Stelle fest, lassen die Integration in einem Development Store mit der neuen Version laufen und lesen das [Changelog zu 2026-10](https://shopify.dev/changelog?api_version=2026-10), gefiltert nach „Action required“. Version 2026-10 wird mindestens bis zum 1. Oktober 2027 unterstützt. Es bleibt also Zeit, aber nicht so viel, dass Sie warten sollten, bis die Warn-E-Mails eintreffen.

Wenn Ihr Shop mit Odoo oder einem anderen System verbunden ist und Sie das lieber jemand anderem überlassen möchten, ist das [genau unsere Arbeit](https://sawirstudio.com/de/shopify/).

**Weiterlesen:** [Shopify Script Tags funktionieren ab dem 1. März 2027 nicht mehr: Was Shop-Betreiber tun sollten](https://sawirstudio.com/de/posts/shopify-script-tags-stop-march-2027/) oder [Shopify mit Odoo 19 Community verbinden, Schritt für Schritt](https://sawirstudio.com/de/posts/connect-shopify-to-odoo-19-community/).
