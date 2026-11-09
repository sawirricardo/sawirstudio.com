---
title: WordPress 7.2: Was Plugin-Entwickler testen sollten | SawirStudio
url: https://sawirstudio.com/de/posts/wordpress-7-2-what-plugin-authors-should-test/
description: WordPress 7.2 soll am 8. Dezember 2026 erscheinen. Was sich für Plugin-Entwickler ändert, von Application Passwords und der vorgeschlagenen Secrets API bis zum Interactivity-API-Router, und wie Sie schon jetzt testen.
---

# WordPress 7.2 erscheint am 8. Dezember: Das testen wir in unseren Plugins

WordPress 7.2 soll am 8. Dezember 2026 erscheinen. Der erste Release Candidate ist für den 17. November angesetzt. Dann wird die Feature-Liste eingefroren, und der Field Guide, die Übersicht aller Entwicklerhinweise, wird veröffentlicht. Bis dahin sind die Betas der günstigste Weg, um herauszufinden, ob Ihr Plugin ein Problem hat.

Wir pflegen [12 kostenlose Plugins auf WordPress.org](https://sawirstudio.com/de/projects/wordpress-plugins/), das hier ist also unsere eigene Checkliste. Einige Punkte waren zum Zeitpunkt des Schreibens noch Vorschläge; maßgeblich ist der Field Guide.

## So testen Sie schon heute

- **WP-CLI** auf einer Testseite: `wp core update --version=7.2-beta4` oder die jeweils neueste Beta bzw. den neuesten RC.
- **WordPress Beta Tester Plugin** mit den Einstellungen „Bleeding edge“ und „Beta/RC Only“.
- **WordPress Playground**, um alles im Browser durchzuklicken, ganz ohne Installation.

Niemals auf einer Live-Website. Aktivieren Sie `WP_DEBUG` und `WP_DEBUG_LOG` und achten Sie auf Deprecation-Hinweise; dort zeigen sich die meisten frühen Warnungen.

## Was wir prüfen

### Application Passwords

7.2 verschärft die Application Passwords: bessere Erkennung lokaler und HTTPS-Umgebungen, eine E-Mail, sobald eines hinzugefügt wird, und ein sichererer Umgang mit der Standardrolle. Wenn Ihr Plugin oder die Tools Ihrer Kunden sich auf diese Weise an der REST API anmelden, testen Sie diesen Ablauf von Anfang bis Ende.

### Die Secrets API

Eine Core-Funktion, um Zugangsdaten verschlüsselt zu speichern, damit Plugins API-Schlüssel nicht länger als Klartext-Optionen ablegen. Sie war für 7.2 vorgeschlagen, die Admin-Oberfläche dazu wurde auf 7.3 verschoben. Wenn Ihr Plugin API-Schlüssel speichert, behalten Sie den Field Guide im Blick: Sobald die Secrets API erscheint, sollten diese Schlüssel dort liegen. Keines unserer kostenlosen Plugins speichert derzeit Schlüssel von Drittanbietern, für uns betrifft das also vor allem die Kundenprojekte.

### Sudo-Modus

Geplant: Ein angemeldeter Nutzer soll vor besonders privilegierten Aktionen sein Passwort erneut bestätigen müssen. Wenn Ihr Plugin solche Aktionen auf eigenen Seiten oder per AJAX ausführt, testen Sie, was passiert, wenn der Core mittendrin eine erneute Bestätigung verlangt.

### Interactivity-API-Router

Ein neuer Router zur Beobachtung der clientseitigen Navigation ersetzt das veraltete Paar `state.navigation.hasStarted` und `hasFinished`. Außerdem gibt es eine Direktive `data-wp-html`, um vertrauenswürdiges Markup auszugeben.

### Admin-Icons und Hinweise

Die Dashicons in der Admin-Leiste und im Menü werden durch SVG-Icons ersetzt, und Admin-Hinweise werden von Screenreadern besser angesagt. Wenn Ihr Plugin sein Menü-Icon per CSS gestaltet oder in Admin-Stylesheets Dashicon-Klassen anspricht, sehen Sie es sich in der Beta an.

### DataViews und der Site Editor

Plugins können eigene Seiten im Site Editor registrieren und DataViews-Felder und -Layouts serverseitig steuern. Hier geht nichts kaputt, aber wenn Ihr Plugin eine Listenansicht hat, lohnt es sich zu wissen, wohin die Reise geht.

## Nach dem Testen

Erhöhen Sie „Tested up to“ in Ihrer readme erst, wenn Sie tatsächlich getestet haben, nicht vorher. WordPress.org zeigt Nutzern eine Warnung an, wenn ein Plugin nicht mit den letzten großen Versionen getestet wurde, und diese Warnung kostet Installationen. Unsere Plugins stehen aktuell auf „Tested up to: 7.1“; auf 7.2 setzen wir sie, sobald sie den RC bestanden haben.

Auch das neue Standard-Theme Ipsum wird mitgeliefert. Wenn Ihr Plugin etwas im Frontend ausgibt, prüfen Sie es dort; so werden neue Websites aussehen.

Wenn Sie ein Plugin haben, das für 7.2 angepasst werden muss, oder ein individuelles Plugin, das niemand mehr pflegt, [werfen wir gerne einen Blick darauf](https://sawirstudio.com/de/contact/).

**Weiterlesen:** [Machen Sie Ihr WordPress-Plugin mit der Abilities API für KI-Agenten nutzbar](https://sawirstudio.com/de/posts/wordpress-abilities-api-for-plugins/) oder [WooCommerce oder Shopify für ein indonesisches Unternehmen?](https://sawirstudio.com/de/posts/woocommerce-or-shopify-indonesia/).
