---
title: Shopify Connector für Odoo | SawirStudio
url: https://sawirstudio.com/de/projects/odoo-shopify-connector/
description: Ein Open-Source-Shopify-Connector für Odoo 19 Community, geschrieben von SawirStudio und gepflegt in der Odoo Community Association (OCA).
---

# Shopify Connector for Odoo

![Das OCA-Repository connector-shopify auf GitHub](https://sawirstudio.com/assets/images/projects/odoo-shopify-connector.webp)

## Das Projekt

Die Odoo Community Association (OCA) ist die gemeinnützige Organisation, die die Open-Source-Module pflegt, auf die sich die meisten Odoo-Community-Nutzer verlassen. Connector Shopify ist ihr Connector zwischen Shopify-Shops und Odoo 19 Community und steht unter der AGPL-3-Lizenz.

## Was wir umgesetzt haben

Wir haben den Connector geschrieben und im Juli 2026 an die OCA übergeben: den Kern-Connector Shopify Connector sowie Shopify Connector Accounting und Shopify Connector POS. Wir entwickeln ihn weiter, zuletzt mit der Anpassung an Shopifys Admin API 2026-01. Auch andere OCA-Mitwirkende bauen inzwischen darauf auf, unter anderem mit einem Modul für eCommerce-Produktmedien.

- **Synchronisierung des Handels.** Produkte, Bestand, Kunden, Bestellungen, Retouren, Erstattungen und Versand, für beliebig viele Shopify-Shops pro Odoo-Datenbank, getrennt nach Unternehmen.
- **Auf Zuverlässigkeit ausgelegt.** Signierte Webhooks für Tempo, OCAs queue\_job für die dauerhafte Verarbeitung und ein zeitgesteuerter Abgleich, der verpasste oder verspätete Ereignisse nachholt.
- **Werkzeuge für den Alltag.** Einrichtung, Statusprüfungen, ein Sync-Dashboard, Fehlerbehebung und konfigurierbare Aufbewahrungsfristen für Daten.
- **Buchhaltung.** Auszahlungen über Shopify Payments, Saldobewegungen, Gebühren, Zahlungsstreitigkeiten, Verbindlichkeiten aus Geschenkgutscheinen und Abstimmung in mehreren Währungen.
- **Kassensystem.** Verkäufe aus Shopify POS werden als Verkaufsaufträge importiert und von jedem Standort einem Odoo-Lager zugeordnet, ganz ohne die eigene POS-App von Odoo.

## Unser Vorgehen

Eine Integration, die meistens stimmt, ist für die Buchhaltung falsch. Deshalb nutzt der Connector die GraphQL Admin API von Shopify, geht nie davon aus, dass ein Webhook angekommen ist, und gleicht sich regelmäßig mit Shopify ab. Buchungssätze werden aus den exakten Beträgen von Shopify geplant und müssen ausgeglichen sein, bevor sie gebucht werden: Wenn etwas nicht aufgeht, meldet der Connector einen klaren Fehler, anstatt stillschweigend eine Ausgleichszeile einzufügen.

Sie möchten ihn für Ihre Shops einrichten oder erweitern lassen? Sehen Sie sich unsere [Shopify-Integrationen](https://sawirstudio.com/de/shopify/) an, oder [Odoo & Unternehmenssysteme](https://sawirstudio.com/de/odoo/) für Odoo selbst.

**Entwickelt mit** Python, Odoo 19 und der Shopify GraphQL Admin API

- **Projekt** Odoo Community Association (OCA)
- **Lizenz** AGPL-3
- **Quellcode** <https://github.com/OCA/connector-shopify>

## Weitere Projekte

[![MyVivaldy](https://sawirstudio.com/assets/images/projects/myvivaldy.webp)

### MyVivaldy](https://sawirstudio.com/de/projects/myvivaldy/)
[![Koperasi HOB Application](https://sawirstudio.com/assets/images/projects/koperasi-hob-application.webp)

### Koperasi HOB Application](https://sawirstudio.com/de/projects/koperasi-hob-application/)
[![DBTV Application](https://sawirstudio.com/assets/images/projects/dbtv-application.webp)

### DBTV Application](https://sawirstudio.com/de/projects/dbtv-application/)

Genug über uns.

## Wie können wir Ihnen helfen?

[Projekt starten](https://sawirstudio.com/de/contact/)
