---
title: Sawir Social Scheduler | SawirStudio
url: https://sawirstudio.com/de/projects/wordpress-plugins/sawir-social-scheduler/
description: Einmal schreiben, aus wp-admin auf 27 Netzwerke planen. Anpassungen pro Kanal, Threads, Medien, Vorschau. Ihre Tokens bleiben auf Ihrem Server.
---

# Sawir Social Scheduler

![Sawir Social Scheduler Banner](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-banner.webp)

Mit Sawir Social Scheduler schreiben Sie einen Beitrag einmal in wp-admin und planen ihn für 27 soziale Netzwerke, mit Anpassungen pro Kanal, Threads, Medien und Vorschau. Die Tokens Ihrer Konten bleiben auf Ihrem Server.

[Kostenlos auf WordPress.org herunterladen](https://wordpress.org/plugins/sawir-social-scheduler/)

## Was es kann

Sawir Social Scheduler macht WordPress zu einem Planungstool für Social Media. Verbinden Sie Ihre Kanäle mit den Tokens, die jede Plattform Ihnen gibt, schreiben Sie einen Beitrag, wählen Sie die Kanäle, legen Sie ein Datum fest und veröffentlichen Sie. Jeder Versand läuft im Hintergrund und meldet sich mit einem Link oder der Fehlermeldung der jeweiligen Plattform zurück.

Es gibt keinen gehosteten Dienst dazwischen. Zugangsdaten werden mit den Salts Ihrer Website verschlüsselt und verlassen Ihren Server nie.

### Kanäle

Fügen Sie beliebig viele Kanäle hinzu, auch mehrere Konten pro Plattform. Jeder Kanal hat eine Schaltfläche **Verbindung testen**, die Kontonamen und Avatar abruft und bestätigt, dass der Token tatsächlich funktioniert, bevor Sie sich darauf verlassen.

**Social:** X (Twitter), Facebook-Seite, Instagram, Threads, LinkedIn-Seite, Bluesky, Mastodon, TikTok, YouTube, Pinterest, Tumblr, Dribbble, Google Business Profile

**Community:** Telegram, Discord, Slack, Reddit, Lemmy, VK, Kick, Twitch, Whop

**Blog und Newsletter:** WordPress, Medium, Dev.to, Hashnode, Listmonk

Gruppieren Sie Kanäle nach Kunde oder Marke mit verschachtelten Gruppen und wählen Sie dann eine ganze Gruppe mit einem Klick aus.

### Beiträge

- **Ein Text, viele Kanäle.** Wählen Sie Kanäle per Kontrollkästchen aus. Unveränderte Kanäle erhalten den Beitrag so, wie er geschrieben ist.
- **Pro Kanal anpassen.** Aktivieren Sie bei einem Kanal „Anpassen“, um ihm eigene Teile zu geben: Text, Medien und Verzögerungen. Sie sind mit dem Beitrag vorausgefüllt, sodass Sie eine Kopie bearbeiten. Die anderen Kanäle behalten den Beitrag unverändert.
- **Threads.** Fügen Sie Teile hinzu, um einen Thread auf X, Threads, Bluesky, Mastodon und Telegram zu veröffentlichen, jeweils mit eigenen Medien. Plattformen ohne Threads erhalten die Teile zusammengefügt.
- **Medien.** Hängen Sie Bilder und Videos nur einmal an. Jede Plattform erhält, was sie unterstützt: Alben, Karussells, Reels, fortsetzbare Video-Uploads.
- **Planen.** Nutzen Sie das übliche WordPress-Datum „Veröffentlichen am“ oder legen Sie für jeden Kanal eine eigene Zeit fest.
- **Posting-Zeiten.** Legen Sie für einen Kanal die üblichen Tageszeiten fest. „Nächster freier Slot“ belegt den nächsten unbenutzten, der Kalender zeigt leere kommende Slots, die Sie per Klick füllen können, und die API akzeptiert `"date": "next"`.
- **Vorschau.** Eine Live-Vorschau pro Kanal, die wie die echte Plattform aussieht. Zeichenlimits und Medienregeln werden angewendet, sodass Sie genau sehen, was veröffentlicht wird.
- **Kommentare und Verzögerungen.** Auf Plattformen ohne Threads werden zusätzliche Teile als Kommentare unter dem Beitrag veröffentlicht. Jeder Teil kann eine festgelegte Anzahl von Minuten nach dem vorherigen warten.
- **Blogbeitrag teilen.** Ein Feld im normalen Beitragseditor veröffentlicht auf Ihren Kanälen, sobald der Beitrag erscheint, mit eigenem Text und eigenen Medien.
- **Duplizieren.** Kopieren Sie einen Beitrag samt Kanälen, Teilen und Anpassungen als neuen Entwurf. Entwürfe lassen sich als Vorlagen aufbewahren.
- **Erwähnungen.** Tippen Sie `@handle`, und das Netzwerk löst es auf X, Instagram, Threads, Mastodon, Telegram, Reddit und Bluesky auf, wo Links und Handles zu echten Facets werden. LinkedIn-Unternehmensseiten schreiben Sie als `@[Name](org:12345)` mit der ID aus der Seiten-URL, Facebook-Seiten als `@[12345]`.
- **Wiederholen.** Senden Sie Evergreen-Beiträge alle N Tage oder Wochen erneut, bis zu einem Datum oder eine bestimmte Anzahl von Malen. Jeder Durchlauf ist ein eigener Beitrag mit eigenen Sendungen und Statistiken. Setzen Sie `{n}` in den Text, um die Durchläufe zu nummerieren, da X, Threads und Bluesky identischen Text ablehnen.

### Kalender

Jeder Beitrag verzweigt sich in einen Versand pro Kanal, und der Kalender zeigt jeden als Karte, eingefärbt nach Status: geplant, gesendet, fehlgeschlagen, angehalten, Entwurf, in Bearbeitung. Ziehen Sie eine Karte, um die Zeit dieses Kanals zu verschieben. Klicken Sie auf eine Karte für den Fehler, eine Schaltfläche „Erneut versuchen“, den Live-Link und Engagement-Zahlen. Klicken Sie auf einen leeren Slot, um zu dieser Zeit einen Beitrag zu beginnen. Monats-, Wochen-, Tages- und Listenansicht, filterbar nach Gruppe und Status.

### Analytics

Likes, Kommentare, Shares, Aufrufe und Impressionen werden auf 16 Plattformen 1 Stunde, 1 Tag und 1 Woche nach jedem Versand sowie auf Abruf erfasst. Follower-Zahlen werden täglich pro Kanal aufgezeichnet, mit einer 30-Tage-Sparkline. Die Zahlen erscheinen auf den Kalenderkarten, im Analytics-Feld des Beitrags und in der Kanalliste. Die Analytics-Seite fasst die letzten 7, 30 oder 90 Tage im Vergleich zum vorherigen Zeitraum zusammen: gesendete Beiträge, Engagement, Reichweite und Follower samt Veränderung, ein Follower-Diagramm, Engagement nach Kanal und Ihre besten Beiträge.

### Freigabe durch Kunden

Aktivieren Sie „Versand bis zur Freigabe anhalten“, kopieren Sie den Vorschau-Link und schicken Sie ihn Ihrem Kunden. Er sieht den Beitrag so, wie er auf jedem Kanal aussehen wird, ohne Login, und klickt auf „Freigeben“ oder „Änderungen anfordern“ mit einer Notiz. Mit der Freigabe werden die Sendungen losgeschickt. Jede Entscheidung bleibt als Verlauf erhalten.

Große Video-Uploads zu X, YouTube und TikTok laufen in kurzen Abschnitten, die über Hintergrundjobs hinweg fortgesetzt werden. So überstehen sie die Zeitlimits des Hostings und laden nie eine ganze Datei in den Arbeitsspeicher.

### REST-API

Alles, was das Backend kann, steht unter `/wp-json/sawir-social/v1` für Automatisierungen, mobile Apps und andere Websites bereit. Die Authentifizierung erfolgt mit einem Anwendungspasswort; die Berechtigungen folgen den Capabilities des Plugins, und Zugangsdaten der Kanäle werden nie zurückgegeben.

- `GET /platforms`, `GET /channels`, `GET /channels/{id}`, `POST /channels/{id}/test`
- `GET|POST /posts`, `GET|PATCH|DELETE /posts/{id}`, `POST /posts/{id}/duplicate`, `GET /posts/{id}/analytics`
- `GET /sends`, `GET /sends/{id}`, `POST /sends/{id}/retry`, `POST /sends/{id}/stats`

Einen Beitrag erstellen Sie mit `{"title","content","status":"future","date":"2026-12-25 09:00","media":[12],"parts":[{"text","media","delay"}],"channels":[{"id":38,"parts":[{"text","media","delay"}],"date":"…"}],"groups":["Client A"],"hold":false}`. Die `parts` eines Kanals ersetzen den Beitrag für diesen Kanal; `text` ist die Kurzform für einen einzelnen Teil.

### KI-Agenten

Dieselben Operationen sind bei der WordPress Abilities API registriert als `sawir-social/list-channels`, `create-post`, `update-post`, `list-posts`, `list-sends`, `retry-send`, `test-channel`, `get-analytics` und `list-platforms`, mit typisierten Ein- und Ausgabeschemata. Jeder Agent, der über den WordPress MCP Adapter verbunden ist, sieht sie als Tools, und sie lassen sich auch unter `/wp-json/wp-abilities/v1/abilities/{name}/run` aufrufen. Die Abilities API ist Teil von WordPress 6.9 und neuer; für ältere Versionen bringt das Plugin sie mit, sodass die Abilities überall verfügbar sind, wo das Plugin läuft.

### Voraussetzungen

Sendungen laufen im Hintergrund über Action Scheduler, wenn Ihre Website es hat (WooCommerce und viele andere Plugins bringen es mit), sonst über WP-Cron. Beim integrierten Cron von WordPress kommen sie nur voran, wenn die Website Besuch bekommt. Für zügigen Versand richten Sie einen System-Cron ein; der Kalender-Bildschirm zeigt den genauen Befehl für Ihre Website.

Medien müssen für Plattformen, die sie per URL abrufen (Instagram, Threads, Facebook, Pinterest, Telegram), aus dem Internet erreichbar sein. Eine lokale Entwicklungsumgebung kann ihnen keine Bilder liefern.

### Credits

Enthält FullCalendar (MIT), die WordPress-Abilities-API-Bibliothek (GPL) und Markenzeichen von Simple Icons (CC0). Nutzt Action Scheduler, wenn ein anderes Plugin es bereitstellt.

## Screenshots

![Sawir Social Scheduler Screenshot 1](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-screenshot-1.webp)

![Sawir Social Scheduler Screenshot 2](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-screenshot-2.webp)

![Sawir Social Scheduler Screenshot 3](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-screenshot-3.webp)

![Sawir Social Scheduler Screenshot 4](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-screenshot-4.webp)

![Sawir Social Scheduler Screenshot 5](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-screenshot-5.webp)

![Sawir Social Scheduler Screenshot 6](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-screenshot-6.webp)

## Fragen

### Woher bekomme ich die Tokens?

In den Entwicklereinstellungen der jeweiligen Plattform. Das Kanalformular nennt die genau benötigten Zugangsdaten, zum Beispiel ein Facebook-Seiten-Zugriffstoken, ein Telegram-Bot-Token oder ein Bluesky-App-Passwort. Langlebige Tokens sind am besten; Tokens von Meta und LinkedIn laufen nach etwa 60 Tagen ab und müssen neu eingefügt werden.

### Mein Telegram-Test meldet „chat not found“

Private Kanäle haben keinen @Benutzernamen. Lassen Sie die Chat-ID leer und klicken Sie auf „Verbindung testen“: Das Plugin listet alle Chats auf, zu denen der Bot hinzugefügt wurde, mit der numerischen ID zum Einfügen.

### Warum bleibt ein Versand bei „In Bearbeitung…“ hängen?

Die Hintergrundwarteschlange ist noch nicht gelaufen. Rufen Sie die Website auf oder richten Sie den System-Cron ein, der im Kalender-Bildschirm angezeigt wird.

### Kann ich von einer lokalen Website aus ein Video auf Instagram veröffentlichen?

Nein. Instagram ruft Medien per URL ab und braucht eine öffentliche Adresse. Alles andere am Plugin funktioniert lokal.

## Passt gut dazu

- [Sawir Insights](https://sawirstudio.com/de/projects/wordpress-plugins/sawir-insights/): zeigt, wie Besucher Ihre Website nutzen, mit Heatmaps, Sitzungsaufzeichnungen, Umfragen, Funnels und Formular-Analysen, und behält alle Daten auf Ihrem eigenen Server..
- [Sawir Live Translate](https://sawirstudio.com/de/projects/wordpress-plugins/sawir-live-translate/): übersetzt Ihre Seiten direkt dort, wo sie stehen: Sprache wählen, sichtbare Inhalte automatisch übersetzen, speichern und Besuchern einen Sprachumschalter anbieten.

Alle unsere [kostenlosen WordPress-Plugins](https://sawirstudio.com/de/projects/wordpress-plugins/) finden Sie hier.

- **Preis** Kostenlos
- **Version** 0.1.0
- **Voraussetzungen** WordPress 5.8, PHP 7.4
- **Lizenz** GPL-2.0-or-later
- **WordPress.org** <https://wordpress.org/plugins/sawir-social-scheduler/>

Genug über uns.

## Wie können wir Ihnen helfen?

[Projekt starten](https://sawirstudio.com/de/contact/)
