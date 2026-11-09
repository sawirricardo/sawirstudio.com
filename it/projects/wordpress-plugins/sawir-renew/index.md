---
title: Sawir Renew – Fatturazione ricorrente per WooCommerce | SawirStudio
url: https://sawirstudio.com/it/projects/wordpress-plugins/sawir-renew/
description: Abbonamenti completi per WooCommerce: fatturazione ricorrente, prove gratuite, costi di attivazione, pausa/salto, cambio di piano, gestione dei pagamenti falliti, riaddebiti intelligenti. Tutto gratis.
---

# Sawir Renew

![Banner di Sawir Renew](https://sawirstudio.com/assets/images/plugins/sawir-renew-banner.webp)

Sawir Renew trasforma un negozio WooCommerce in un'attività a abbonamento: fatturazione ricorrente, prove gratuite, costi di attivazione e tentativi di riaddebito intelligenti, senza versioni a pagamento che blocchino alcuna funzione.

[Scaricatelo gratis da WordPress.org](https://wordpress.org/plugins/sawir-renew/)

## Cosa fa

Vendete abbonamenti con WooCommerce senza pagare per un'estensione. Tutte le funzioni sono gratuite: non esiste una versione pro e nulla è bloccato.

### Prodotti e prezzi

- Tipi di prodotto in abbonamento semplici e variabili
- Fatturazione ogni N giorni, settimane, mesi o anni, con supporto per i prezzi in offerta
- Prove gratuite, costi di attivazione e abbonamenti che scadono dopo N pagamenti
- Rinnovi sincronizzati (tutti rinnovano il 1° del mese, o nel giorno che scegliete voi) con primo periodo facoltativamente in proporzione o gratuito
- Spedizione una tantum, limite di «un solo abbonamento attivo per cliente»

### Due modalità di pagamento

- **Pagamento a ogni rinnovo (funziona con qualsiasi gateway):** quando un rinnovo è in scadenza, il cliente riceve una fattura con un link di pagamento: bonifico, e-wallet, contrassegno, qualsiasi metodo accettato dal vostro negozio. Ideale per i mercati in cui l'addebito automatico su carta non è la norma.
- **Addebito automatico tramite Stripe:** se usate il gateway ufficiale WooCommerce Stripe Gateway, le carte vengono salvate al checkout e i rinnovi vengono addebitati automaticamente senza che il cliente sia presente, con tentativi di riaddebito intelligenti e un link di pagamento alternativo quando la carta richiede l'autenticazione. Altri gateway possono collegarsi allo stesso adattatore (vedi le FAQ).

### Rinnovi che non vi svegliano alle 3 di notte

- Tutte le tempistiche passano da Action Scheduler, con un controllo giornaliero che si autoripara
- Gestione configurabile dei pagamenti falliti: ritardi dei tentativi, calendario dei promemoria, periodi di tolleranza, annullamento automatico
- Il rilevamento di siti di staging o duplicati mette in pausa la fatturazione non appena l'URL del vostro sito cambia: basta doppi addebiti causati da una copia di staging dimenticata
- E-mail di preavviso prima degli addebiti automatici

### I clienti possono fare da soli

Da Il mio account → Abbonamenti i clienti possono annullare (subito o a fine periodo, a vostra scelta), mettere in pausa e riprendere, saltare un rinnovo, rinnovare in anticipo, cambiare metodo di pagamento o indirizzo di spedizione, e passare a un piano superiore o inferiore con addebito proporzionale in caso di upgrade.

### Coupon ricorrenti

Quattro tipi di coupon per abbonamenti: sconti ricorrenti in percentuale o a importo fisso (anche limitati ai primi N pagamenti) e sconti sui costi di attivazione. I coupon standard continuano ad applicarsi solo all'ordine iniziale.

### Notifiche

Tredici e-mail attivabili o disattivabili (per clienti e amministratori) per ogni momento del ciclo di vita dell'abbonamento. Gli sviluppatori possono replicare qualsiasi evento su altri canali (WhatsApp, SMS, Slack…) tramite l'hook `swsb_notified`.

### Pensato per il WooCommerce di oggi

Nativo HPOS (gli abbonamenti risiedono nell'archiviazione ordini di WooCommerce, senza tabelle personalizzate), compatibile con i blocchi carrello/checkout, elenco e editor completi nell'area admin che usano le schermate ordini di WooCommerce, una pagina di report con cache oraria (attivi, MRR, abbandoni, prossimi rinnovi, ricavi da rinnovi) e un'API REST.

## Domande

### Ha qualche legame con WooCommerce o Automattic?

No. «WooCommerce» è un marchio di Automattic; questo è un plugin indipendente *per* WooCommerce, nel rispetto delle loro linee guida sui nomi.

### Quali gateway supportano i rinnovi automatici?

Di serie: il gateway ufficiale WooCommerce Stripe Gateway. Tutti gli altri gateway funzionano in modalità pagamento a ogni rinnovo: il cliente riceve un link di pagamento a ogni periodo. Gli sviluppatori possono aggiungere il supporto automatico per qualsiasi gateway con due filtri:

`swsb_order_supports_auto`: restituisce true quando il gateway dell'ordine principale ha salvato un metodo di pagamento riutilizzabile. `swsb_charge_renewal`: addebita l'ordine di rinnovo senza la presenza del cliente e restituisce `true` oppure un `WP_Error`.

### Può funzionare insieme a WooCommerce Subscriptions?

No: far girare due sistemi di fatturazione addebiterebbe due volte i clienti, quindi questo plugin disattiva il proprio motore e mostra un avviso se WooCommerce Subscriptions è attivo.

### Cosa succede su un sito di staging?

Se l'URL del sito cambia, la fatturazione e le e-mail ai clienti si fermano automaticamente finché non confermate se la copia è un sito di staging o il sito live. Scegliendo «live», tutte le pianificazioni vengono ricreate.

## Si abbina bene a

- [Sawir Referral](https://sawirstudio.com/it/projects/wordpress-plugins/sawir-referral/): permette ai vostri clienti di portare i loro amici.
- [Sawir Pulse](https://sawirstudio.com/it/projects/wordpress-plugins/sawir-pulse/): mostra ai clienti cosa succede davvero nel vostro negozio, come gli ordini recenti, le persone che stanno guardando un prodotto e le nuove recensioni, con piccole notifiche in un angolo della pagina.

Scoprite tutti i nostri [plugin WordPress gratuiti](https://sawirstudio.com/it/projects/wordpress-plugins/).

- **Prezzo** Gratuito
- **Versione** 1.0.0
- **Requisiti** WordPress 6.0, PHP 7.4, WooCommerce
- **Licenza** GPL-2.0-or-later
- **WordPress.org** <https://wordpress.org/plugins/sawir-renew/>

Basta parlare di noi.

## Come possiamo aiutarvi?

[Avviate un progetto](https://sawirstudio.com/it/contact/)
