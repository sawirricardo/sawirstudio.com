---
title: Sawir Renew – Wiederkehrende Abrechnung für WooCommerce | SawirStudio
url: https://sawirstudio.com/de/projects/wordpress-plugins/sawir-renew/
description: Komplette Abonnements für WooCommerce – wiederkehrende Abrechnung, Testphasen, Einrichtungsgebühren, Pausieren/Überspringen, Tarifwechsel, Mahnwesen, intelligente Wiederholungen. Alles kostenlos.
---

# Sawir Renew

![Sawir Renew Banner](https://sawirstudio.com/assets/images/plugins/sawir-renew-banner.webp)

Sawir Renew macht aus einem WooCommerce-Shop ein Abo-Geschäft: wiederkehrende Abrechnung, Testphasen, Einrichtungsgebühren und intelligente Zahlungswiederholungen – ohne kostenpflichtige Version, die Funktionen zurückhält.

[Kostenlos auf WordPress.org herunterladen](https://wordpress.org/plugins/sawir-renew/)

## Was es kann

Verkaufen Sie Abonnements mit WooCommerce, ohne für eine Erweiterung zu bezahlen. Alle Funktionen sind kostenlos enthalten – es gibt keine Pro-Version und nichts ist gesperrt.

### Produkte & Preise

- Produkttypen für einfache und variable Abonnements
- Abrechnung alle N Tage, Wochen, Monate oder Jahre, mit Unterstützung für Angebotspreise
- Kostenlose Testphasen, Einrichtungsgebühren und Abonnements, die nach N Zahlungen enden
- Synchronisierte Verlängerungen (alle verlängern am 1. oder an einem Tag Ihrer Wahl) mit optional anteiliger oder kostenloser erster Laufzeit
- Einmalige Versandkosten, Begrenzung auf „ein aktives Abonnement pro Kunde“

### Zwei Zahlungswege

- **Zahlung pro Verlängerung (funktioniert mit jedem Zahlungsanbieter):** Wird eine Verlängerung fällig, erhält der Kunde eine Rechnung mit Zahlungslink – per Überweisung, E-Wallet, Nachnahme oder was immer Ihr Shop akzeptiert. Ideal für Märkte, in denen automatische Kartenabrechnung nicht üblich ist.
- **Automatische Abrechnung über Stripe:** Wenn Sie das offizielle WooCommerce Stripe Gateway nutzen, werden Karten beim Checkout gespeichert und Verlängerungen automatisch im Hintergrund abgebucht – inklusive intelligenter Wiederholungen und einem Ersatz-Zahlungslink, falls eine Karte eine Authentifizierung verlangt. Andere Zahlungsanbieter lassen sich über dieselbe Adapter-Schnittstelle anbinden (siehe FAQ).

### Verlängerungen, die Sie nachts nicht wecken

- Das gesamte Timing läuft über Action Scheduler, mit einer täglichen selbstheilenden Prüfung
- Konfigurierbares Mahnwesen: Wiederholungsabstände, Erinnerungsplan, Kulanzfristen, automatische Kündigung
- Erkennung von Staging- und Duplikat-Websites pausiert die Abrechnung, sobald sich die URL Ihrer Website ändert – keine doppelten Abbuchungen mehr durch eine vergessene Staging-Kopie
- E-Mails zur Vorabankündigung vor automatischen Abbuchungen

### Kunden helfen sich selbst

Unter „Mein Konto → Abonnements“ können Kunden kündigen (sofort oder zum Laufzeitende – Sie entscheiden), pausieren und fortsetzen, eine Verlängerung überspringen, vorzeitig verlängern, ihre Zahlungsmethode oder Lieferadresse ändern sowie zwischen Tarifen wechseln (Upgrade/Downgrade), mit anteiliger Berechnung bei Upgrades.

### Wiederkehrende Gutscheine

Vier Gutscheintypen für Abonnements: wiederkehrende prozentuale oder feste Rabatte (optional auf die ersten N Zahlungen begrenzt) sowie Rabatte auf die Einrichtungsgebühr. Standard-Gutscheine gelten weiterhin nur für die erste Bestellung.

### Benachrichtigungen

Dreizehn einzeln aktivierbare E-Mails (an Kunden und Admin) für jede Phase im Abo-Verlauf. Entwickler können jedes Ereignis über den Hook `swsb_notified` auch an andere Kanäle weiterleiten (WhatsApp, SMS, Slack …).

### Für modernes WooCommerce gebaut

HPOS-nativ (Abonnements liegen im WooCommerce-Bestellspeicher – keine eigenen Tabellen), kompatibel mit Warenkorb- und Checkout-Blöcken, vollständige Admin-Liste und Editor in den eigenen Bestellansichten von WooCommerce, eine stündlich zwischengespeicherte Berichtsseite (aktive Abos, MRR, Abwanderung, anstehende Verlängerungen, Umsatz aus Verlängerungen) und eine REST-API.

## Fragen

### Besteht eine Verbindung zu WooCommerce oder Automattic?

Nein. „WooCommerce“ ist eine Marke von Automattic; dies ist ein unabhängiges Plugin *für* WooCommerce, gemäß deren Namensrichtlinien.

### Welche Zahlungsanbieter unterstützen automatische Verlängerungen?

Von Haus aus: das offizielle WooCommerce Stripe Gateway. Alle anderen Zahlungsanbieter funktionieren im Modus „Zahlung pro Verlängerung“ – Kunden erhalten in jeder Periode einen Zahlungslink. Entwickler können mit zwei Filtern automatische Unterstützung für jeden Zahlungsanbieter ergänzen:

`swsb_order_supports_auto` – gibt true zurück, wenn der Zahlungsanbieter der Ursprungsbestellung eine wiederverwendbare Zahlungsmethode gespeichert hat. `swsb_charge_renewal` – bucht die Verlängerungsbestellung im Hintergrund ab und gibt `true` oder einen `WP_Error` zurück.

### Kann es parallel zu WooCommerce Subscriptions laufen?

Nein – zwei Abrechnungssysteme würden Kunden doppelt belasten. Deshalb deaktiviert dieses Plugin seine Abrechnungs-Engine und zeigt einen Hinweis an, wenn WooCommerce Subscriptions aktiv ist.

### Was passiert auf einer Staging-Website?

Ändert sich die Website-URL, werden Abrechnung und Kunden-E-Mails automatisch pausiert, bis Sie bestätigen, ob es sich um eine Staging-Kopie oder die Live-Website handelt. Wählen Sie „Live“, werden alle Zeitpläne neu aufgebaut.

## Passt gut dazu

- [Sawir Referral](https://sawirstudio.com/de/projects/wordpress-plugins/sawir-referral/): Ihre Kunden werben ihre Freunde.
- [Sawir Pulse](https://sawirstudio.com/de/projects/wordpress-plugins/sawir-pulse/): zeigt Ihren Kunden als kleine Hinweise in der Ecke, was in Ihrem Shop wirklich passiert, etwa aktuelle Bestellungen, Besucher, die sich ein Produkt ansehen, und neue Bewertungen.

Alle unsere [kostenlosen WordPress-Plugins](https://sawirstudio.com/de/projects/wordpress-plugins/) finden Sie hier.

- **Preis** Kostenlos
- **Version** 1.0.0
- **Voraussetzungen** WordPress 6.0, PHP 7.4, WooCommerce
- **Lizenz** GPL-2.0-or-later
- **WordPress.org** <https://wordpress.org/plugins/sawir-renew/>

Genug über uns.

## Wie können wir Ihnen helfen?

[Projekt starten](https://sawirstudio.com/de/contact/)
