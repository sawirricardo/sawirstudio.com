---
title: Sawir Live Translate | SawirStudio
url: https://sawirstudio.com/de/projects/wordpress-plugins/sawir-live-translate/
description: Sawir Live Translate übersetzt Ihre Seiten direkt dort, wo sie stehen: Sprache wählen, sichtbare Texte automatisch übersetzen, speichern und Besuchern einen Sprachumschalter anbieten.
---

# Sawir Live Translate

![Sawir Live Translate Banner](https://sawirstudio.com/assets/images/plugins/sawir-live-translate-banner.webp)

Sawir Live Translate übersetzt Ihre Seiten direkt dort, wo sie stehen: Sprache wählen, sichtbare Texte automatisch übersetzen, speichern und Besuchern einen Sprachumschalter anbieten.

[Kostenlos auf WordPress.org herunterladen](https://wordpress.org/plugins/sawir-live-translate/)

## Was es kann

Öffnen Sie eine beliebige Seite Ihrer Website, klicken Sie in der Admin-Leiste auf „Translate this page“ und wählen Sie eine Zielsprache. Dann übersetzen Sie entweder alle sichtbaren Texte automatisch (über einen kostenlosen maschinellen Übersetzungsdienst) oder bearbeiten die Übersetzungen von Hand und sehen dabei live, wie sich die Seite verändert.

Gespeicherte Übersetzungen werden serverseitig direkt in die HTML-Antwort eingebaut (SEO-freundlich: kein kurzes Aufblitzen unübersetzter Texte, dazu hreflang-Alternativen im Head). Besucher wechseln die Sprache über einen Link-Umschalter (unten rechts).

Die URLs verwenden ein Sprachpräfix im Pfad, z. B. /es/about für Spanisch und /about für die Ausgangssprache (sie bleibt kanonisch und ohne Präfix). Die Sprache wird in dieser Reihenfolge bestimmt: /es/-Pfad → ?lang=es → gespeichertes Cookie → Accept-Language des Browsers → Ausgangssprache. Bei der Root-URL und bei URLs ohne Präfix wird die Browsersprache des Besuchers ohne Weiterleitung erkannt (es wird ein Vary-Header gesendet, damit gemeinsam genutzte Caches korrekt arbeiten).

1. Unter Einstellungen → Live Translate legen Sie Ihre Ausgangssprache fest und wählen die Sprachen aus, die angeboten werden sollen. 2. Rufen Sie eine Seite im Frontend auf und klicken Sie in der Admin-Leiste auf „🌐 Translate this page“. 3. Wählen Sie eine Sprache, klicken Sie auf „Auto-translate empty“ (oder bearbeiten Sie die Texte) und dann auf „Save“. 4. Besucher nutzen den Umschalter (unten rechts); jede Sprache hat ihre eigene /xx/-URL.

Hinweis: Die Sprachpräfixe im Pfad setzen eine Installation in der Root-Domain voraus und dass keine Seite auf oberster Ebene einen Slug aus zwei Buchstaben als Sprachcode verwendet.

## Fragen

### Wie schneidet es im Vergleich zu Alternativen ab?

Mehrsprachige Frameworks wie WPML, Polylang und TranslatePress behandeln übersetzte Inhalte als festen Teil der Website-Struktur, mit eigenen Bearbeitungsabläufen pro Sprache, Übersetzungsdiensten und tiefer Integration in Page Builder und WooCommerce. Sawir Live Translate ist bewusst schlanker: Sie öffnen eine Seite, übersetzen, was sichtbar ist (automatisch oder von Hand), und liefern sie über eine URL mit Sprachpräfix samt hreflang-Tags aus. Wenn Sie eine vollständig mehrsprachige Website mit übersetzten Menüs, Produkten und redaktionellen Abläufen aufbauen, greifen Sie zu einem der Frameworks. Wenn Sie bestehende Seiten noch heute Nachmittag in einer anderen Sprache lesbar machen möchten, dann ist genau das die Aufgabe dieses Plugins.

### Ich wechsle die Sprache, aber nichts ändert sich

Meistens bedeutet das, dass für diese Sprache noch keine Übersetzungen gespeichert sind. Prüfen Sie unter Einstellungen → Live Translate → Saved translations: Wird die Sprache dort nicht aufgeführt (oder mit 0 Strings angezeigt), öffnen Sie den Seiteneditor erneut, übersetzen Sie automatisch und klicken Sie auf „Save“. Die Live-Vorschau im Editor wird erst gespeichert, wenn Sie das tun.

### Einige Texte bleiben unübersetzt

Der Editor auf der Seite sieht nur die Texte der Seite, die Sie geöffnet haben. Texte auf anderen Seiten müssen ebenfalls übersetzt werden, oder Sie starten den Scan der gesamten Website. Texte, die Slider oder Skripte erst nach dem Laden der Seite einfügen, müssen Sie unter Umständen manuell über „Edit all site strings“ ergänzen.

### Ich habe meinen Inhalt geändert und die Übersetzung ist verschwunden

Übersetzungen werden dem exakten Originaltext zugeordnet. Wenn Sie einen Satz umformulieren, übersetzen Sie die neue Version einmal und speichern Sie sie.

## Passt gut dazu

- [Sawir Insights](https://sawirstudio.com/de/projects/wordpress-plugins/sawir-insights/): zeigt, wie Besucher Ihre Website nutzen, mit Heatmaps, Sitzungsaufzeichnungen, Umfragen, Funnels und Formular-Analysen, und behält alle Daten auf Ihrem eigenen Server..
- [Sawir Media Folders](https://sawirstudio.com/de/projects/wordpress-plugins/sawir-media-folders/): bringt Ordner in die WordPress-Mediathek, so viele Sie brauchen, und Dateien lassen sich per Drag-and-drop einsortieren..

Alle unsere [kostenlosen WordPress-Plugins](https://sawirstudio.com/de/projects/wordpress-plugins/) finden Sie hier.

- **Preis** Kostenlos
- **Version** 1.1.0
- **Voraussetzungen** WordPress 5.8, PHP 7.4
- **Lizenz** GPL-2.0-or-later
- **WordPress.org** <https://wordpress.org/plugins/sawir-live-translate/>

Genug über uns.

## Wie können wir Ihnen helfen?

[Projekt starten](https://sawirstudio.com/de/contact/)
