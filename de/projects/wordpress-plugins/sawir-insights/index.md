---
title: Sawir Insights | SawirStudio
url: https://sawirstudio.com/de/projects/wordpress-plugins/sawir-insights/
description: Heatmaps, Session-Aufzeichnungen, Umfragen, Funnels und Formular-Analysen. Alles bleibt auf Ihrem Server.
---

# Sawir Insights

![Sawir Insights Banner](https://sawirstudio.com/assets/images/plugins/sawir-insights-banner.webp)

Sawir Insights zeigt, wie Besucher Ihre Website nutzen: mit Heatmaps, Session-Aufzeichnungen, Umfragen, Funnels und Formular-Analysen. Alle Daten bleiben auf Ihrem eigenen Server.

[Kostenlos auf WordPress.org herunterladen](https://wordpress.org/plugins/sawir-insights/)

## Was es kann

Sawir Insights ist eine Verhaltensanalyse für WordPress: Heatmaps, Session-Aufzeichnungen, Umfragen, ein Feedback-Widget, Formular-Analysen, Conversion-Funnels und ein Dashboard, das alles zusammenführt.

In Multisite-Netzwerken erfasst und konfiguriert jede Website ihre Analysen unabhängig von den anderen.

Kurz gesagt: eine selbst gehostete Alternative zu Hotjar oder Microsoft Clarity, bei der die Besucherdaten in Ihrer eigenen Datenbank bleiben.

Der Grund dafür ist einfach. Die etablierten Tools dieser Kategorie sind gut, aber sie funktionieren, indem sie ein Protokoll von allem, was Ihre Besucher tun, an fremde Server senden. Für viele Websites, besonders in der EU, ist das irgendwo zwischen unangenehm und nicht praktikabel. Ich wollte dieselbe Art von Werkzeug, bei dem die Daten nie die Website verlassen, die sie erhoben hat. Genau darauf ist dieses Plugin ausgelegt: Ihre Datenbank, Ihr Uploads-Verzeichnis, Ihre Aufbewahrungsregeln. Dahinter steht kein SaaS, es ist kein Konto nötig, und es gibt keinerlei Verbindung zu externen Servern.

### Was es kann

- Klick-, Bewegungs- und Scroll-Heatmaps, dargestellt über Ihren Live-Seiten und getrennt nach Desktop, Tablet und Mobilgerät.
- Session-Aufzeichnungen mit Wiedergabe in wp-admin. Eingabewerte werden immer maskiert; die Aufzeichnung ist ausgeschaltet, bis Sie sie aktivieren.
- Popup-Umfragen und ein dauerhafter Feedback-Reiter, mit Targeting nach URL, Gerät, Traffic-Anteil, Verweildauer, Scrolltiefe oder Exit-Intent.
- Formular-Analysen: Wo Nutzer zögern, was sie leer lassen und wo sie abbrechen. Erfasst werden nur Selektoren und Zeitangaben, niemals, was jemand eingegeben hat.
- Geordnete Conversion-Funnels mit täglichen Berichten.
- Ein Dashboard mit Traffic-KPIs und Periodenvergleich, einem Diagramm zum täglichen Verlauf sowie Aufschlüsselungen nach Gerät, Browser, Seite, Referrer, Kampagne, Funnel und Aufzeichnung.

### Was das Plugin bewusst nicht tut

Ich finde, Sie sollten wissen, was ein Tool verweigert, bevor Sie es installieren.

- Es speichert keine IP-Adressen. Weder gehasht noch gekürzt, überhaupt nicht.
- Es zeichnet niemals Formularwerte auf. Die Erfassungsebene sendet nur Metadaten, es kann also nichts abfließen.
- Es lädt nichts von einem CDN und sendet auch nichts an eines. Die Replay-Bibliothek rrweb ist im Plugin enthalten.
- Es behält Rohdaten nicht ewig. Ereignisse werden aggregiert und nach einem Aufbewahrungsplan gelöscht, den Sie selbst festlegen. Für Aufzeichnungen gelten eigene Aufbewahrungsfristen und ein Speicherlimit, bei dem die ältesten zuerst gelöscht werden, denn ein Plugin, das heimlich den Speicherplatz Ihres Hostings füllt, ist ein schlechtes Plugin.

Der Frontend-Tracker ist kleines Vanilla-JavaScript, das Daten gebündelt über `sendBeacon` sendet und Sampling-Einstellungen bietet. Besucher, die nicht für die Aufzeichnung ausgewählt werden, laden die Replay-Bibliothek gar nicht erst herunter.

### Datenschutz

Do Not Track wird standardmäßig beachtet. Eine Consent-Mode-API (`window.swinsConsent`) lässt sich in jeden Consent-Manager einbinden, den Sie bereits einsetzen. Sessions, Ereignisse und Aufzeichnungen sind an die Export- und Löschfunktionen für personenbezogene Daten von WordPress angebunden. Ob Ihr konkreter Einsatz DSGVO-konform ist, hängt von Ihren Hinweisen und Ihrer Rechtsgrundlage ab, und das kann Ihnen kein Plugin abnehmen. Die Datensparsamkeit ist aber umgesetzt.

### Support

Support erhalten Sie über das wordpress.org-Forum zu diesem Plugin oder per E-Mail an ricardo@sawirstudio.com.

## Fragen

### Wie schneidet es im Vergleich zu Alternativen ab?

Gehostete Analysedienste wie Hotjar und Microsoft Clarity sind hervorragende Tools mit großzügigen kostenlosen Tarifen, aber sie funktionieren, indem sie ein Protokoll des Verhaltens Ihrer Besucher an ihre Server senden. Sawir Insights verzichtet auf diese übergreifende Infrastruktur und setzt auf eine einzige Vorgabe: Alles wird auf Ihrer eigenen Website erfasst, gespeichert und angezeigt. Wählen Sie einen gehosteten Dienst, wenn Sie dessen Zusatzfunktionen möchten und die Verarbeitung durch Dritte kein Problem für Sie ist. Wählen Sie Sawir Insights, wenn es für Sie ein Ausschlusskriterium ist, dass Besucherdaten Ihren Server verlassen. Aufzeichnungen und Heatmaps liegen in Ihrer eigenen Datenbank und Ihrem Uploads-Verzeichnis, daher sind die Einstellungen zur Aufbewahrung hier wichtiger als bei einem SaaS.

### Ist Sawir Insights DSGVO-konform?

Das Plugin ist für datenschutzbewusste Einsätze gebaut: Die Daten bleiben auf Ihrem Server, IP-Adressen werden nie gespeichert, die Aufbewahrung ist konfigurierbar, Do Not Track wird beachtet, und die Export- und Löschfunktionen von WordPress sind eingebunden. Die Konformität selbst hängt von Ihrer Konfiguration, Ihren Hinweisen, Ihrer Rechtsgrundlage und der Umsetzung der Einwilligung ab. Lassen Sie sich für Ihre Website von qualifizierten Rechtsberatern unterstützen.

### Wie unterscheidet sich das von Hotjar oder Microsoft Clarity?

Das sind gehostete Dienste: leistungsfähige Tools, aber das Verhalten Ihrer Besucher wird nach deren Bedingungen auf deren Servern aufgezeichnet. Sawir Insights deckt dieselbe Kategorie ab (Heatmaps, Session-Aufzeichnungen, Umfragen, Funnels, Formular-Analysen) als Plugin, das alles auf Ihrem eigenen Server hält, ohne Konto, ohne externen Dienst und ohne Datenweitergabe.

### Verlassen Besucherdaten meinen Server?

Nein. Das Plugin sendet keine Verhaltensdaten an Sawir Studio oder an einen Analyseanbieter. Ereignisse liegen in Ihrer WordPress-Datenbank, Aufzeichnungsdateien in Ihrem geschützten Uploads-Verzeichnis.

## Passt gut dazu

- [Sawir Pulse](https://sawirstudio.com/de/projects/wordpress-plugins/sawir-pulse/): zeigt Ihren Kunden als kleine Hinweise in der Ecke, was in Ihrem Shop wirklich passiert, etwa aktuelle Bestellungen, Besucher, die sich ein Produkt ansehen, und neue Bewertungen.
- [Sawir Live Translate](https://sawirstudio.com/de/projects/wordpress-plugins/sawir-live-translate/): übersetzt Ihre Seiten direkt dort, wo sie stehen: Sprache wählen, sichtbare Inhalte automatisch übersetzen, speichern und Besuchern einen Sprachumschalter anbieten.

Alle unsere [kostenlosen WordPress-Plugins](https://sawirstudio.com/de/projects/wordpress-plugins/) finden Sie hier.

- **Preis** Kostenlos
- **Version** 1.0.0
- **Voraussetzungen** WordPress 6.2, PHP 7.4
- **Lizenz** GPLv2 oder neuer
- **WordPress.org** <https://wordpress.org/plugins/sawir-insights/>

Genug über uns.

## Wie können wir Ihnen helfen?

[Projekt starten](https://sawirstudio.com/de/contact/)
