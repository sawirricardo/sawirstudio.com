---
title: Sawir Social Scheduler | SawirStudio
url: https://sawirstudio.com/nl/projects/wordpress-plugins/sawir-social-scheduler/
description: Schrijf één keer, plan in op 27 netwerken vanuit wp-admin. Aanpassingen per kanaal, threads, media, voorbeelden. Je tokens blijven op je eigen server.
---

# Sawir Social Scheduler

![Sawir Social Scheduler banner](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-banner.webp)

Met Sawir Social Scheduler schrijf je een bericht één keer in wp-admin en plan je het in op 27 sociale netwerken, met aanpassingen per kanaal, threads, media en voorbeeldweergaven. De tokens van je accounts blijven op je eigen server.

[Gratis downloaden via WordPress.org](https://wordpress.org/plugins/sawir-social-scheduler/)

## Wat het doet

Sawir Social Scheduler maakt van WordPress een planner voor social media. Koppel je kanalen met de tokens die elk platform je geeft, schrijf een bericht, kies kanalen, stel een datum in en publiceer. Elke verzending draait op de achtergrond en meldt terug met een link of de eigen foutmelding van het platform.

Er zit geen gehoste dienst tussen. Inloggegevens worden versleuteld met de salts van je site en verlaten je server nooit.

### Kanalen

Voeg zoveel kanalen toe als je wilt, ook meerdere accounts per platform. Elk kanaal heeft een knop **Verbinding testen** die de accountnaam en avatar ophaalt en bevestigt dat het token echt werkt voordat je erop vertrouwt.

**Social:** X (Twitter), Facebook-pagina, Instagram, Threads, LinkedIn-pagina, Bluesky, Mastodon, TikTok, YouTube, Pinterest, Tumblr, Dribbble, Google Bedrijfsprofiel

**Community:** Telegram, Discord, Slack, Reddit, Lemmy, VK, Kick, Twitch, Whop

**Blog en nieuwsbrief:** WordPress, Medium, Dev.to, Hashnode, Listmonk

Groepeer kanalen per klant of merk met geneste groepen en selecteer een hele groep met één klik.

### Berichten

- **Eén tekst, meerdere kanalen.** Kies kanalen met selectievakjes. Kanalen die je niet aanpast, krijgen het bericht zoals je het hebt geschreven.
- **Aanpassen per kanaal.** Vink "Aanpassen" aan bij een kanaal om het eigen onderdelen te geven: tekst, media en vertragingen, vooraf ingevuld vanuit het bericht, zodat je een kopie bewerkt. De andere kanalen houden het bericht zoals geschreven.
- **Threads.** Voeg onderdelen toe om een thread te plaatsen op X, Threads, Bluesky, Mastodon en Telegram, elk met eigen media. Platforms zonder threads krijgen de onderdelen samengevoegd.
- **Media.** Voeg afbeeldingen en video één keer toe. Elk platform krijgt wat het ondersteunt: albums, carrousels, Reels, hervatbare video-uploads.
- **Inplannen.** Gebruik de gewone WordPress-datum "Publiceren op", of geef elk kanaal een eigen tijdstip.
- **Posttijden.** Geef een kanaal de tijden van de dag waarop je meestal post. "Volgende vrije slot" vult het eerstvolgende ongebruikte slot in, de kalender toont lege komende slots die je kunt aanklikken om te vullen, en de API accepteert `"date": "next"`.
- **Voorbeeld.** Een live voorbeeld per kanaal dat eruitziet als het echte platform, met tekenlimieten en mediaregels toegepast, zodat je precies ziet wat er live gaat.
- **Reacties en vertragingen.** Op platforms zonder threads worden extra onderdelen als reacties onder het bericht geplaatst. Elk onderdeel kan een ingesteld aantal minuten na het vorige wachten.
- **Een blogbericht delen.** Een blok in de gewone berichteneditor plaatst op je kanalen zodra het bericht wordt gepubliceerd, met eigen tekst en media.
- **Dupliceren.** Kopieer een bericht met zijn kanalen, onderdelen en aanpassingen als nieuw concept. Bewaar concepten als sjablonen.
- **Vermeldingen.** Typ `@handle` en het netwerk herkent die op X, Instagram, Threads, Mastodon, Telegram, Reddit en Bluesky, waar links en handles echte facets worden. LinkedIn-bedrijfspagina's zijn `@[Name](org:12345)` met de ID uit de URL van de pagina; Facebook-pagina's zijn `@[12345]`.
- **Herhalen.** Verstuur evergreen berichten opnieuw om de N dagen of weken, tot een bepaalde datum of een aantal keer. Elke herhaling is een eigen bericht met eigen verzendingen en statistieken. Zet `{n}` in de tekst om herhalingen te nummeren, want X, Threads en Bluesky weigeren identieke tekst.

### Kalender

Elk bericht wordt uitgesplitst in één verzending per kanaal, en de kalender toont elke verzending als een kaart die is gekleurd naar status: ingepland, verzonden, mislukt, vastgehouden, concept, bezig. Sleep een kaart om het tijdstip van dat kanaal te verplaatsen. Klik op een kaart voor de foutmelding, een knop Opnieuw proberen, de live link en de engagementcijfers. Klik op een leeg slot om op dat tijdstip een bericht te starten. Weergave per maand, week, dag en als lijst, te filteren op groep en status.

### Analytics

Likes, reacties, shares, weergaven en vertoningen worden 1 uur, 1 dag en 1 week na elke verzending verzameld, en ook op aanvraag, op 16 platforms. Het aantal volgers wordt dagelijks per kanaal vastgelegd met een sparkline van 30 dagen. De cijfers staan op de kalenderkaarten, in het Analytics-blok van het bericht en in de kanalenlijst. De Analytics-pagina vat de laatste 7, 30 of 90 dagen samen en vergelijkt ze met de periode ervoor: verzonden berichten, engagement, bereik en volgers met hun verandering, een volgersgrafiek, engagement per kanaal en je beste berichten.

### Goedkeuring door de klant

Vink "Verzendingen vasthouden tot goedkeuring" aan, kopieer de voorbeeldlink en stuur die naar je klant. Die ziet het bericht zoals het er op elk kanaal uit komt te zien, zonder in te loggen, en klikt op Goedkeuren of Wijzigingen aanvragen met een toelichting. Na goedkeuring gaan de verzendingen live. Elke beslissing blijft bewaard als geschiedenis.

Grote video-uploads naar X, YouTube en TikTok draaien in korte delen die over achtergrondtaken heen worden hervat. Zo overleven ze de tijdslimieten van je host en wordt nooit een heel bestand in het geheugen geladen.

### REST API

Alles wat de beheeromgeving kan, is beschikbaar via `/wp-json/sawir-social/v1`, voor automatiseringen, mobiele apps en andere sites. Authenticeer met een applicatiewachtwoord; rechten volgen de capabilities van de plugin en inloggegevens van kanalen worden nooit teruggegeven.

- `GET /platforms`, `GET /channels`, `GET /channels/{id}`, `POST /channels/{id}/test`
- `GET|POST /posts`, `GET|PATCH|DELETE /posts/{id}`, `POST /posts/{id}/duplicate`, `GET /posts/{id}/analytics`
- `GET /sends`, `GET /sends/{id}`, `POST /sends/{id}/retry`, `POST /sends/{id}/stats`

Maak een bericht aan met `{"title","content","status":"future","date":"2026-12-25 09:00","media":[12],"parts":[{"text","media","delay"}],"channels":[{"id":38,"parts":[{"text","media","delay"}],"date":"…"}],"groups":["Client A"],"hold":false}`. De `parts` van een kanaal vervangen het bericht voor dat kanaal; `text` is een verkorte schrijfwijze voor één onderdeel.

### AI-agents

Dezelfde acties zijn geregistreerd bij de WordPress Abilities API als `sawir-social/list-channels`, `create-post`, `update-post`, `list-posts`, `list-sends`, `retry-send`, `test-channel`, `get-analytics` en `list-platforms`, met getypeerde invoer- en uitvoerschema's. Elke agent die via de WordPress MCP adapter is verbonden, ziet ze als tools, en ze zijn ook aan te roepen via `/wp-json/wp-abilities/v1/abilities/{name}/run`. De Abilities API maakt deel uit van WordPress 6.9 en nieuwer; voor oudere versies levert de plugin hem zelf mee, zodat de abilities overal beschikbaar zijn waar de plugin draait.

### Vereisten

Verzendingen draaien op de achtergrond via Action Scheduler als je site dat heeft (WooCommerce en veel andere plugins leveren het mee), en anders via WP-Cron. Op de ingebouwde cron van WordPress gaan ze alleen vooruit als de site bezoekers krijgt. Voor snelle verzending voeg je een systeem-cron toe; het Kalender-scherm toont het exacte commando voor jouw site.

Media moeten bereikbaar zijn vanaf het internet voor platforms die ze via een URL ophalen (Instagram, Threads, Facebook, Pinterest, Telegram). Een lokale ontwikkelsite kan geen afbeeldingen naar die platforms sturen.

### Credits

Bevat FullCalendar (MIT), de WordPress Abilities API-bibliotheek (GPL) en merklogo's van Simple Icons (CC0). Gebruikt Action Scheduler wanneer een andere plugin die levert.

## Schermafbeeldingen

![Sawir Social Scheduler screenshot 1](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-screenshot-1.webp)

![Sawir Social Scheduler screenshot 2](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-screenshot-2.webp)

![Sawir Social Scheduler screenshot 3](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-screenshot-3.webp)

![Sawir Social Scheduler screenshot 4](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-screenshot-4.webp)

![Sawir Social Scheduler screenshot 5](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-screenshot-5.webp)

![Sawir Social Scheduler screenshot 6](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-screenshot-6.webp)

## Vragen

### Waar haal ik de tokens vandaan?

Uit de ontwikkelaarsinstellingen van elk platform. Het kanaalformulier noemt precies welke inloggegevens nodig zijn, bijvoorbeeld een toegangstoken voor een Facebook-pagina, een bot-token voor Telegram of een app-wachtwoord voor Bluesky. Langlevende tokens zijn het best; tokens van Meta en LinkedIn verlopen na ongeveer 60 dagen en moet je opnieuw plakken.

### Mijn Telegram-test geeft "chat not found"

Privékanalen hebben geen @gebruikersnaam. Laat de Chat-ID leeg en klik op Verbinding testen: de plugin toont alle chats waaraan de bot is toegevoegd, met het numerieke ID dat je kunt plakken.

### Waarom blijft een verzending hangen op "Bezig…"?

De wachtrij op de achtergrond is nog niet gedraaid. Bezoek de site, of stel de systeem-cron in die op het Kalender-scherm staat.

### Kan ik vanaf een lokale site een video op Instagram plaatsen?

Nee. Instagram haalt media op via een URL en heeft een openbaar adres nodig. Al het andere in de plugin werkt gewoon lokaal.

## Past goed bij

- [Sawir Insights](https://sawirstudio.com/nl/projects/wordpress-plugins/sawir-insights/): laat zien hoe bezoekers je site gebruiken, met heatmaps, sessie-opnames, enquêtes, funnels en formulieranalyse, en houdt alles op je eigen server..
- [Sawir Live Translate](https://sawirstudio.com/nl/projects/wordpress-plugins/sawir-live-translate/): vertaalt je pagina's op de plek waar ze staan: kies een taal, vertaal automatisch wat zichtbaar is, sla het op en geef bezoekers een taalkeuzemenu.

Bekijk al onze [gratis WordPress-plugins](https://sawirstudio.com/nl/projects/wordpress-plugins/).

- **Prijs** Gratis
- **Versie** 0.1.0
- **Vereist** WordPress 5.8, PHP 7.4
- **Licentie** GPL-2.0-or-later
- **WordPress.org** <https://wordpress.org/plugins/sawir-social-scheduler/>

Genoeg over ons.

## Waar kunnen we je mee helpen?

[Start een project](https://sawirstudio.com/nl/contact/)
