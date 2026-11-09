---
title: Sawir Social Scheduler | SawirStudio
url: https://sawirstudio.com/it/projects/wordpress-plugins/sawir-social-scheduler/
description: Scrivete una volta sola e programmate su 27 network da wp-admin. Modifiche per canale, thread, contenuti multimediali, anteprime. I vostri token restano sul vostro server.
---

# Sawir Social Scheduler

![Banner di Sawir Social Scheduler](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-banner.webp)

Sawir Social Scheduler vi permette di scrivere un post una sola volta in wp-admin e di programmarlo su 27 social network, con modifiche per singolo canale, thread, contenuti multimediali e anteprime. I token dei vostri account restano sul vostro server.

[Scaricatelo gratis da WordPress.org](https://wordpress.org/plugins/sawir-social-scheduler/)

## Cosa fa

Sawir Social Scheduler trasforma WordPress in uno strumento di programmazione per i social media. Collegate i vostri canali con i token che ogni piattaforma vi fornisce, scrivete un post, scegliete i canali, impostate una data e pubblicate. Ogni invio viene eseguito in background e vi restituisce un link oppure il messaggio di errore della piattaforma stessa.

Non c'è nessun servizio esterno in mezzo. Le credenziali sono cifrate con i salt del vostro sito e non lasciano mai il vostro server.

### Canali

Aggiungete tutti i canali che volete, anche più account per piattaforma. Ogni canale ha un pulsante **Test connection** che recupera il nome e l'avatar dell'account e conferma che il token funzioni davvero prima di affidarvi a esso.

**Social:** X (Twitter), Facebook Page, Instagram, Threads, LinkedIn Page, Bluesky, Mastodon, TikTok, YouTube, Pinterest, Tumblr, Dribbble, Google Business Profile

**Community:** Telegram, Discord, Slack, Reddit, Lemmy, VK, Kick, Twitch, Whop

**Blog e newsletter:** WordPress, Medium, Dev.to, Hashnode, Listmonk

Raggruppate i canali per cliente o per brand con gruppi annidati, poi selezionate un intero gruppo con un solo clic.

### Post

- **Un solo testo, tanti canali.** Scegliete i canali con le caselle di selezione. I canali che non modificate ricevono il post così come l'avete scritto.
- **Personalizzazione per canale.** Spuntate «Customize» su un canale qualsiasi per dargli parti tutte sue: testo, contenuti multimediali e ritardi, precompilati a partire dal post, così modificate una copia. Gli altri canali mantengono il post così com'è.
- **Thread.** Aggiungete delle parti per pubblicare un thread su X, Threads, Bluesky, Mastodon e Telegram, ciascuna con i propri contenuti multimediali. Le piattaforme senza thread ricevono le parti unite in un unico testo.
- **Contenuti multimediali.** Allegate immagini e video una volta sola. Ogni piattaforma riceve ciò che supporta: album, caroselli, Reels, caricamenti video ripristinabili.
- **Programmazione.** Usate la normale data «Pubblica il» di WordPress, oppure assegnate a ogni canale un orario tutto suo.
- **Orari di pubblicazione.** Assegnate a un canale i suoi orari abituali della giornata. «Next free slot» riempie il primo orario libero, il calendario mostra i prossimi slot vuoti su cui potete cliccare per riempirli, e l'API accetta `"date": "next"`.
- **Anteprima.** Un'anteprima in tempo reale per ogni canale, identica alla piattaforma vera, con limiti di caratteri e regole sui contenuti multimediali già applicati, così vedete esattamente ciò che verrà pubblicato.
- **Commenti e ritardi.** Sulle piattaforme senza thread, le parti aggiuntive vengono pubblicate come commenti sotto il post. Ogni parte può attendere un certo numero di minuti dopo la precedente.
- **Condividere un articolo del blog.** Un riquadro nel normale editor degli articoli pubblica sui vostri canali nel momento in cui l'articolo viene pubblicato, con testo e contenuti multimediali propri.
- **Duplicazione.** Copiate un post con i suoi canali, le sue parti e le sue personalizzazioni come nuova bozza. Conservate le bozze come modelli.
- **Menzioni.** Scrivete `@handle` e il network lo risolve su X, Instagram, Threads, Mastodon, Telegram, Reddit e Bluesky, dove link e handle diventano veri riferimenti. Le pagine aziendali di LinkedIn si scrivono `@[Name](org:12345)` con l'ID preso dall'URL della pagina; le pagine Facebook si scrivono `@[12345]`.
- **Ripetizione.** Rinviate i post evergreen ogni N giorni o settimane, fino a una certa data o per un certo numero di volte. Ogni ripetizione è un post a sé, con i propri invii e le proprie statistiche. Inserite `{n}` nel testo per numerare le ripetizioni, perché X, Threads e Bluesky rifiutano i testi identici.

### Calendario

Ogni post si divide in un invio per canale, e il calendario mostra ciascun invio come una scheda colorata in base allo stato: programmato, inviato, non riuscito, in attesa, bozza, in corso. Trascinate una scheda per spostare l'orario di quel canale. Cliccate su una scheda per vedere l'errore, un pulsante Retry, il link pubblicato e i dati di engagement. Cliccate su uno slot vuoto per iniziare un post a quell'ora. Viste per mese, settimana, giorno ed elenco, filtrabili per gruppo e stato.

### Analytics

Mi piace, commenti, condivisioni, visualizzazioni e impression vengono raccolti 1 ora, 1 giorno e 1 settimana dopo ogni invio, oltre che su richiesta, su 16 piattaforme. Il numero di follower viene registrato ogni giorno per canale, con uno sparkline di 30 giorni. I dati compaiono nelle schede del calendario, nel riquadro Analytics del post e nell'elenco dei canali. La pagina Analytics riassume gli ultimi 7, 30 o 90 giorni rispetto al periodo precedente: post inviati, engagement, copertura e follower con la relativa variazione, un grafico dei follower, l'engagement per canale e i vostri post migliori.

### Approvazione del cliente

Spuntate «Hold sends until approved», copiate il link dell'anteprima e inviatelo al vostro cliente. Vedrà il post esattamente come apparirà su ogni canale, senza bisogno di login, e potrà cliccare su Approve oppure su Request changes con una nota. L'approvazione sblocca gli invii. Ogni decisione viene conservata nello storico.

I caricamenti di video di grandi dimensioni su X, YouTube e TikTok avvengono in brevi porzioni che riprendono da un job in background all'altro, quindi superano i limiti di tempo dell'hosting e non caricano mai un intero file in memoria.

### REST API

Tutto ciò che si fa nell'admin è disponibile su `/wp-json/sawir-social/v1` per automazioni, app mobile e altri siti. L'autenticazione avviene con una application password; i permessi seguono le capability del plugin e le credenziali dei canali non vengono mai restituite.

- `GET /platforms`, `GET /channels`, `GET /channels/{id}`, `POST /channels/{id}/test`
- `GET|POST /posts`, `GET|PATCH|DELETE /posts/{id}`, `POST /posts/{id}/duplicate`, `GET /posts/{id}/analytics`
- `GET /sends`, `GET /sends/{id}`, `POST /sends/{id}/retry`, `POST /sends/{id}/stats`

Create a post with `{"title","content","status":"future","date":"2026-12-25 09:00","media":[12],"parts":[{"text","media","delay"}],"channels":[{"id":38,"parts":[{"text","media","delay"}],"date":"…"}],"groups":["Client A"],"hold":false}`. A channel's `parts` replace the post for that channel; `text` is shorthand for a single part.

### Agenti AI

Le stesse operazioni sono registrate nella WordPress Abilities API come `sawir-social/list-channels`, `create-post`, `update-post`, `list-posts`, `list-sends`, `retry-send`, `test-channel`, `get-analytics` e `list-platforms`, con schemi di input e output tipizzati. Qualsiasi agente collegato tramite il WordPress MCP adapter le vede come strumenti, e sono richiamabili anche su `/wp-json/wp-abilities/v1/abilities/{name}/run`. L'Abilities API fa parte di WordPress 6.9 e successivi; per le versioni precedenti il plugin la include, quindi le abilities sono disponibili ovunque giri il plugin.

### Requisiti

Gli invii vengono eseguiti in background tramite Action Scheduler se il vostro sito lo ha (lo forniscono WooCommerce e molti altri plugin), altrimenti tramite WP-Cron. Con il cron integrato di WordPress, gli invii avanzano solo quando il sito riceve traffico. Per invii puntuali, aggiungete un cron di sistema: la schermata Calendario mostra il comando esatto per il vostro sito.

Per le piattaforme che recuperano i file tramite URL (Instagram, Threads, Facebook, Pinterest, Telegram), i contenuti multimediali devono essere raggiungibili da internet. Un sito di sviluppo locale non può consegnare loro le immagini.

### Crediti

Include FullCalendar (MIT), la libreria WordPress Abilities API (GPL) e i marchi di Simple Icons (CC0). Usa Action Scheduler quando è fornito da un altro plugin.

## Screenshot

![Screenshot 1 di Sawir Social Scheduler](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-screenshot-1.webp)

![Screenshot 2 di Sawir Social Scheduler](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-screenshot-2.webp)

![Screenshot 3 di Sawir Social Scheduler](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-screenshot-3.webp)

![Screenshot 4 di Sawir Social Scheduler](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-screenshot-4.webp)

![Screenshot 5 di Sawir Social Scheduler](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-screenshot-5.webp)

![Screenshot 6 di Sawir Social Scheduler](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-screenshot-6.webp)

## Domande

### Dove trovo i token?

Nelle impostazioni sviluppatori di ciascuna piattaforma. Il modulo del canale indica la credenziale esatta di cui ha bisogno, ad esempio un access token di una Pagina Facebook, un bot token di Telegram o una app password di Bluesky. I token a lunga durata sono i migliori; i token di Meta e LinkedIn scadono dopo circa 60 giorni e vanno incollati di nuovo.

### Il test di Telegram mi dà «chat not found»

I canali privati non hanno un @username. Lasciate vuoto il Chat ID e cliccate su Test connection: il plugin elenca tutte le chat a cui è stato aggiunto il bot, con l'ID numerico da incollare.

### Perché un invio resta bloccato su «In progress…»?

La coda in background non è ancora stata eseguita. Visitate il sito, oppure configurate il cron di sistema indicato nella schermata Calendario.

### Posso pubblicare un video su Instagram da un sito locale?

No. Instagram recupera i file tramite URL e ha bisogno di un indirizzo pubblico. Tutto il resto del plugin funziona in locale.

## Si abbina bene a

- [Sawir Insights](https://sawirstudio.com/it/projects/wordpress-plugins/sawir-insights/): mostra come i visitatori usano il vostro sito, con heatmap, registrazioni delle sessioni, sondaggi, funnel e analisi dei form, e tiene tutto sul vostro server..
- [Sawir Live Translate](https://sawirstudio.com/it/projects/wordpress-plugins/sawir-live-translate/): traduce le vostre pagine direttamente dove si trovano: scegliete una lingua, traducete in automatico ciò che è visibile, salvate e offrite ai visitatori un selettore di lingua.

Scoprite tutti i nostri [plugin WordPress gratuiti](https://sawirstudio.com/it/projects/wordpress-plugins/).

- **Prezzo** Gratuito
- **Versione** 0.1.0
- **Richiede** WordPress 5.8, PHP 7.4
- **Licenza** GPL-2.0-or-later
- **WordPress.org** <https://wordpress.org/plugins/sawir-social-scheduler/>

Basta parlare di noi.

## Come possiamo aiutarvi?

[Avviate un progetto](https://sawirstudio.com/it/contact/)
