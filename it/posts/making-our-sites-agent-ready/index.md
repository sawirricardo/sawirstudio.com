---
title: Rendere i nostri siti pronti per gli agenti | SawirStudio
url: https://sawirstudio.com/it/posts/making-our-sites-agent-ready/
description: Come abbiamo reso i siti dei nostri prodotti leggibili dagli agenti AI: copie in Markdown di ogni pagina, header Link, llms.txt, Content Signals in robots.txt, un catalogo delle API e Agent Skills.
---

# Cosa è servito per rendere i nostri siti pronti per gli agenti

Sempre più visite a un sito web non sono di persone, ma di agenti AI che leggono le pagine per rispondere alla domanda di qualcuno, o che cercano di usare un prodotto per conto di qualcuno. Questa settimana abbiamo passato in rassegna questo sito e i siti dei nostri prodotti, [SawirPost](https://sawirstudio.com/it/projects/sawirpost/), [Sawirly](https://sawirstudio.com/it/projects/sawirly/), [Rindraw](https://sawirstudio.com/it/projects/rindraw/), [Math Daily](https://sawirstudio.com/it/projects/math-daily/) e altri, per renderli più facili da leggere e da usare per gli agenti.

Come checklist abbiamo usato [isitagentready.com](https://isitagentready.com), che valuta un sito per livelli. Ecco cosa ha richiesto ogni passaggio.

## Copie in Markdown di ogni pagina

Un agente che legge l'HTML deve farsi strada tra navigazione, script e footer per trovare pochi paragrafi. Per questo ogni pagina di questo sito ha ora accanto una copia in Markdown, in `index.md`, con solo il contenuto principale della pagina. Le genera da HTML un piccolo script, così le due versioni non si disallineano mai.

Gli agenti le ottengono in due modi: aggiungendo `index.md` a un URL, oppure richiedendo `Accept: text/markdown`. Un piccolissimo Cloudflare Worker davanti ai file statici controlla quell'header e serve la copia in Markdown, con una stima approssimativa dei token in un header `x-markdown-tokens`. I browser continuano a ricevere la pagina normale.

```
curl -H "Accept: text/markdown" https://sawirstudio.com/shopify/
```

## Header Link

Ogni pagina HTML ora indica dove si trova la sua copia in Markdown, in un header HTTP `Link` con `rel="alternate"`. La home page rimanda anche a `llms.txt` e alla sitemap. Un agente può trovare tutto con una sola richiesta, senza dover indovinare gli URL.

## llms.txt

`llms.txt` è un indice del sito in testo semplice pensato per i modelli linguistici: un riassunto di una riga e un elenco di tutte le pagine con la relativa descrizione, con i link alle copie in Markdown. Il nostro viene generato dalla sitemap dallo stesso script.

## Content Signals

`robots.txt` dice chi può fare crawling, non cosa può fare con ciò che legge. [Content Signals](https://contentsignals.org/) aggiunge proprio questo. Il nostro è una sola riga e consente tutti e tre gli usi, perché vogliamo essere trovati e citati:

```
Content-Signal: ai-train=yes, search=yes, ai-input=yes
```

Questi tre passaggi hanno portato ogni sito al Livello 3.

## Un catalogo delle API e le Agent Skills

SawirPost e Sawirly hanno delle API, quindi sono andati oltre. Ciascuno pubblica un catalogo delle API in `/.well-known/api-catalog` (RFC 9727): un piccolo documento JSON che rimanda alla descrizione OpenAPI e alla documentazione. Un agente che arriva sul sito può trovare l'API senza leggere nemmeno una pagina.

Pubblicano anche delle Agent Skills, brevi istruzioni per gli agenti su come svolgere i compiti più comuni con il prodotto. Questo ha portato entrambi al Livello 4.

## Dove ci siamo fermati

Il Livello 5 richiede che gli agenti possano registrarsi e autenticarsi in autonomia, tramite un flusso di registrazione degli agenti OAuth. Una pagina che descrive come registrarsi con email e password non conta. È una funzionalità vera da sviluppare, con questioni di sicurezza altrettanto vere, quindi ci stiamo ragionando con calma invece di correre.

## Ne è valsa la pena?

Per un sito statico, quasi tutto si è risolto in un pomeriggio: uno script, un piccolo Worker e una riga in `robots.txt`. Per chi visita il sito non è cambiato nulla. Pensiamo che sempre più domande riceveranno risposta da agenti che leggono siti come il nostro, e una pagina Markdown pulita è il modo più semplice per essere citati correttamente.

Per i clienti configuriamo le stesse cose all'interno di [Lancio e crescita](https://sawirstudio.com/it/launch-and-growth/).

**Continuate a leggere:** [Rendete il vostro plugin WordPress utilizzabile dagli agenti AI con la Abilities API](https://sawirstudio.com/it/posts/wordpress-abilities-api-for-plugins/), oppure [Le icone delle app lanciabili nella nostra home page](https://sawirstudio.com/it/posts/throwable-app-icons-matter-js/).
