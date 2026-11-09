---
title: Shopify mit Odoo 19 Community verbinden | SawirStudio
url: https://sawirstudio.com/de/posts/connect-shopify-to-odoo-19-community/
description: So verbinden Sie einen Shopify-Shop mit Odoo 19 Community über den quelloffenen OCA-Shopify-Connector: Dev-Dashboard-App, Scopes, queue_job, Webhooks und Onboarding.
---

# Shopify mit Odoo 19 Community verbinden: Schritt für Schritt

Wir haben den [Shopify-Connector für Odoo](https://sawirstudio.com/de/projects/odoo-shopify-connector/) selbst entwickelt und der Odoo Community Association (OCA) beigesteuert. Er ist kostenlos, quelloffen (AGPL-3) und läuft mit Odoo 19 Community. Hier die Einrichtung von Anfang bis Ende, so wie wir sie für Kunden durchführen.

Er synchronisiert Produkte, Bestände, Kunden, Bestellungen, Retouren, Erstattungen und Fulfillment, für einen oder mehrere Shops. Buchhaltung und Kassensystem sind separate Module, die Sie später ergänzen können.

## Bevor Sie beginnen

- Odoo 19 Community, erreichbar über öffentliches HTTPS. Shopify sendet Webhooks dorthin und lädt von dort Produktbilder herunter.
- Admin-Zugriff auf den Shopify-Shop und auf die Shopify-Organisation, der er gehört.
- Ein Shopify-Tarif, der geschützte Kundendaten für selbst erstellte Apps enthält. Der Tarif Basic enthält sie nicht. Ohne sie verweigert Shopify einige Bestell-Webhooks, und der Import von Bestellungen und Kunden schlägt fehl.

## 1. App im Shopify Dev Dashboard erstellen

Seit dem 1. Januar 2026 können im Shop-Admin keine alten „Custom Apps“ mehr erstellt werden. Erstellen Sie die App im **Dev Dashboard** der Organisation, der der Shop gehört. Das ist wichtig: Odoo erhält sein Zugriffstoken über den Client-Credentials-Grant, und der funktioniert nur, wenn App und Shop zur selben Organisation gehören.

Beim Ausfüllen der App-Version gilt:

- **App-URL:** `https://shopify.dev/apps/default-app-home`. Der Connector hat keine Seite innerhalb von Shopify.
- **App im Shopify-Admin einbetten:** aus.
- **Legacy-Installationsablauf verwenden:** aus, damit die Installation der App die ausgewählten Scopes erteilt.
- **Webhooks:** keine. Odoo legt seine eigenen beim Onboarding an.
- **Webhooks-API-Version:** `2026-01`, die vom Connector verwendete Version.
- **Scopes:** `read_customers`, `write_customers`, `read_draft_orders`, `read_fulfillments`, `read_inventory`, `write_inventory`, `read_locations`, `read_merchant_managed_fulfillment_orders`, `write_merchant_managed_fulfillment_orders`, `read_orders`, `read_products`, `write_products` und `read_returns`. Ergänzen Sie `read_all_orders`, wenn Sie Bestellungen importieren möchten, die älter als 60 Tage sind.

Geben Sie die Version frei, installieren Sie die App im Shop und kopieren Sie Client-ID und Client-Secret von der Seite „Einstellungen“ der App.

Sie haben bereits eine Legacy-Custom-App aus der Zeit vor 2026? Sie funktioniert weiterhin. Lassen Sie die Client-ID leer und verwenden Sie stattdessen das Admin-API-Zugriffstoken und den API-Secret-Key.

## 2. Module in Odoo installieren

Fügen Sie OCA `queue_job` und das Repository [connector-shopify](https://github.com/OCA/connector-shopify) zu Ihrem Addons-Pfad hinzu und installieren Sie dann **Shopify Connector**. Die Arbeit läuft als Hintergrundjobs, daher muss `queue_job` serverweit geladen sein, mit einer Kapazität für jeden Shopify-Kanal:

```
[options]
server_wide_modules = base,rpc,web,queue_job

[queue_job]
channels = root:4,root.shopify.import:2,root.shopify.export:1,root.shopify.webhook:1
```

Setzen Sie anschließend den Systemparameter `web.base.url` auf die öffentliche HTTPS-Adresse von Odoo und `web.base.url.freeze` auf `True`. Sonst ändert sich die Adresse, wenn sich ein Admin von einem anderen Ort aus anmeldet, und die Webhooks von Shopify landen am falschen Ziel.

## 3. Onboarding durchführen

Weisen Sie Ihrem Team die Gruppe **Shopify Manager** zu, öffnen Sie dann **Shopify > Configuration > Instances** und klicken Sie auf **Start Onboarding**:

1. Geben Sie die Domain `yourstore.myshopify.com`, die Client-ID und das Client-Secret ein und testen Sie die Verbindung. Odoo fordert das Token an und erneuert es alle 24 Stunden von selbst.
2. Prüfen Sie die Scopes. Fehlt einer, fügen Sie ihn in einer neuen App-Version hinzu, geben Sie sie frei und bestätigen Sie sie im Shop.
3. Ordnen Sie die Shopify-Standorte den Odoo-Lagern zu, prüfen Sie die von Odoo erzeugten Steuerzuordnungen und wählen Sie die Synchronisierungsrichtlinien.
4. Installieren Sie die Webhooks und stellen Sie die ersten Importe in die Warteschlange.

Bevor Sie Bestellungen und Fulfillments unbeaufsichtigt laufen lassen, prüfen Sie die Zuordnungen für Zahlungsarten und Versanddienstleister. Die meisten späteren Überraschungen entstehen durch eine Zuordnung, die sich niemand angesehen hat.

## 4. Behalten Sie es im Blick

**Shopify > Dashboard** zeigt Synchronisierungsstatus, den Job-Rückstau, aktuelle Fehler und den Zustand der Webhooks. Der Connector geht nicht davon aus, dass jeder Webhook ankommt: Geplante Jobs importieren alles erneut, was sich in Shopify geändert hat und in Odoo fehlt. Wenn etwas fehlschlägt, beheben Sie die Ursache und nutzen dann **Retry Failed Shopify Bindings**. Führen Sie **Check Webhooks** aus, nachdem Sie Domains oder Zugangsdaten geändert haben.

## Nächste Schritte

Ergänzen Sie **Shopify Connector Accounting**, um Auszahlungen, Gebühren und Streitfälle von Shopify Payments als Buchungssätze zu übernehmen ([warum sich Auszahlungen so schwer zuordnen lassen](https://sawirstudio.com/de/posts/why-shopify-payouts-dont-match-your-books/), erklären wir separat), und **Shopify Connector POS** für Verkäufe im Ladengeschäft.

Wenn Sie sich die Einrichtung abnehmen lassen, auf Ihre Abläufe zuschneiden oder erweitern möchten, [übernehmen wir das gern](https://sawirstudio.com/de/shopify/).

**Weiterlesen:** [Warum Ihre Shopify-Auszahlungen nie zur Buchhaltung passen](https://sawirstudio.com/de/posts/why-shopify-payouts-dont-match-your-books/) oder [Shopifys API 2026-10: Was in Ihren Integrationen kaputtgeht](https://sawirstudio.com/de/posts/shopify-2026-10-api-what-breaks/).
