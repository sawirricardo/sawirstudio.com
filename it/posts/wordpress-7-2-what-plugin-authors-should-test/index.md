---
title: WordPress 7.2: cosa devono testare gli autori di plugin | SawirStudio
url: https://sawirstudio.com/it/posts/wordpress-7-2-what-plugin-authors-should-test/
description: WordPress 7.2 è previsto per l’8 dicembre 2026. Cosa cambia per gli autori di plugin, dalle Application Passwords alla proposta Secrets API fino al router dell’Interactivity API, e come testare fin da ora.
---

# WordPress 7.2 arriva l’8 dicembre: cosa stiamo testando nei nostri plugin

WordPress 7.2 è previsto per l’8 dicembre 2026. La prima release candidate è attesa per il 17 novembre: è il momento in cui l’elenco delle funzionalità viene congelato e esce il Field Guide, la raccolta di tutte le note per gli sviluppatori. Fino ad allora, le beta sono il modo più economico per scoprire se il vostro plugin ha un problema.

Gestiamo [12 plugin gratuiti su WordPress.org](https://sawirstudio.com/it/projects/wordpress-plugins/), quindi questa è la nostra checklist. Alcuni punti qui sotto erano ancora proposte quando abbiamo scritto l’articolo; fa fede il Field Guide.

## Come testare oggi

- **WP-CLI** su un sito di prova: `wp core update --version=7.2-beta4`, oppure la beta o RC più recente.
- **Plugin WordPress Beta Tester**, impostato su «Bleeding edge» e «Beta/RC Only».
- **WordPress Playground**, per provare tutto dal browser senza installare nulla.

Mai su un sito in produzione. Attivate `WP_DEBUG` e `WP_DEBUG_LOG` e tenete d’occhio gli avvisi di deprecazione: è lì che compaiono la maggior parte dei primi segnali.

## Cosa controlliamo

### Application Passwords

La 7.2 rende più rigorose le Application Passwords: migliore rilevamento degli ambienti locali e HTTPS, un’email quando ne viene aggiunta una e una gestione più sicura del ruolo predefinito. Se il vostro plugin o gli strumenti dei vostri clienti si autenticano alla REST API in questo modo, provate il flusso da cima a fondo.

### La Secrets API

Un sistema integrato in core per conservare le credenziali in forma cifrata, così i plugin smettono di salvare le chiavi API come semplici opzioni. Era stata proposta per la 7.2, con la schermata di amministrazione rimandata alla 7.3. Se il vostro plugin salva chiavi API, tenetela d’occhio nel Field Guide: una volta rilasciata, sarà il posto giusto dove conservarle. Nessuno dei nostri plugin gratuiti salva oggi chiavi di terze parti, quindi per noi riguarda soprattutto i lavori per i clienti.

### Sudo mode

In programma: chiedere a un utente connesso di confermare di nuovo la password prima di azioni con privilegi elevati. Se il vostro plugin esegue azioni di questo tipo dalle proprie schermate o via AJAX, provate cosa succede quando il core richiede una nuova verifica nel mezzo dell’operazione.

### Router dell’Interactivity API

Un nuovo router per osservare la navigazione lato client sostituisce la coppia deprecata `state.navigation.hasStarted` e `hasFinished`. C’è anche una direttiva `data-wp-html` per renderizzare markup attendibile.

### Icone e notifiche in admin

I Dashicons vengono sostituiti da icone SVG nella barra di amministrazione e nel menu, e le notifiche in admin ottengono annunci migliori per gli screen reader. Se il vostro plugin personalizza lo stile dell’icona del menu o usa le classi dei Dashicons nel CSS dell’admin, controllate come appare nella beta.

### DataViews e Site Editor

I plugin possono registrare le proprie schermate nel Site Editor e controllare dal server i campi e i layout di DataViews. Qui non si rompe nulla, ma se il vostro plugin ha una schermata con un elenco, vale la pena sapere dove si sta andando.

## Dopo i test

Aggiornate «Tested up to» nel readme solo dopo aver testato davvero, non prima. WordPress.org mostra agli utenti un avviso quando un plugin non è stato testato con le ultime versioni principali, e quell’avviso fa perdere installazioni. I nostri plugin dicono «Tested up to: 7.1» al momento; passeranno alla 7.2 una volta superati i test sulla RC.

Arriva anche il nuovo tema predefinito, Ipsum. Se il vostro plugin mostra qualcosa nel front end, controllatelo lì: è l’aspetto che avranno i nuovi siti.

Se avete un plugin da aggiornare per la 7.2, o uno su misura che nessuno mantiene più, [possiamo dargli un’occhiata](https://sawirstudio.com/it/contact/).

**Continuate a leggere:** [Rendete il vostro plugin WordPress utilizzabile dagli agenti AI con l’Abilities API](https://sawirstudio.com/it/posts/wordpress-abilities-api-for-plugins/), oppure [WooCommerce o Shopify per un’azienda indonesiana?](https://sawirstudio.com/it/posts/woocommerce-or-shopify-indonesia/).
