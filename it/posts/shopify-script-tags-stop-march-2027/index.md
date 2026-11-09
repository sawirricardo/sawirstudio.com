---
title: Gli script tag di Shopify finiscono a marzo 2027 | SawirStudio
url: https://sawirstudio.com/it/posts/shopify-script-tags-stop-march-2027/
description: Shopify smette di eseguire gli script tag dello store online il 1° marzo 2027. Quali app coinvolge, come controllare il vostro negozio e cosa li sostituisce.
---

# Gli script tag di Shopify smettono di funzionare il 1° marzo 2027. Cosa devono fare i titolari di negozi

Il 1° marzo 2027 Shopify smette di inserire gli script tag nelle vetrine online. Tutto ciò che nel vostro negozio viene ancora caricato in questo modo, un widget di recensioni, una chat, un pixel di tracciamento, semplicemente non comparirà più. Non si rompe nulla in modo evidente. Le cose, semplicemente, spariscono.

Shopify lo ha annunciato ad [agosto 2026](https://shopify.dev/changelog/posts/online-store-script-tags-deprecation). Ecco cosa significa se gestite un negozio.

## Cos'è uno script tag

Uno script tag è un modo per un'app di aggiungere JavaScript alla vostra vetrina tramite API, senza toccare il tema. Per questo le app lo hanno sempre usato volentieri, ma è anche il problema: gli script tag non compaiono nel theme editor, quindi non potete vederli né disattivarli da lì. A volte le app disinstallate hanno lasciato i propri.

## Le date da ricordare

- **1° ottobre 2026:** le app non possono più creare o aggiornare script tag, con nessuna versione dell'API. Quelli esistenti continuano a funzionare.
- **1° marzo 2027:** Shopify smette di caricare gli script tag dello store online.

Gli script tag nella pagina di stato dell'ordine sono stati ritirati prima, con un calendario separato: se qualcosa ha già smesso di comparire lì, il motivo è questo.

## Come controllare il vostro negozio

Shopify non offre ai merchant un elenco degli script tag, quindi bisogna andare a cercarli:

1. **Elencate le app che inseriscono qualcosa nella vostra vetrina.** Recensioni, chat, pop-up, upsell, fedeltà, analytics, pixel pubblicitari. Le app che funzionano solo nell'area di amministrazione non sono interessate.
2. **Guardate gli App embed nel theme editor.** Se un'app compare lì ed è attiva, usa il metodo più recente e non ci sono problemi.
3. **Guardate il codice sorgente della pagina.** Aprite il vostro negozio, visualizzate il sorgente e cercate `asyncLoad`. Nella maggior parte dei negozi gli URL degli script tag sono elencati in quel blocco. Ogni dominio che trovate di solito rimanda a un'app.
4. **Chiedete allo sviluppatore dell'app.** «La vostra app usa script tag e funzionerà dopo il 1° marzo 2027?» è una domanda a cui un'app mantenuta può rispondere in una riga.

Da anni Shopify spinge le app verso le theme app extension, quindi la maggior parte delle app ben mantenute è già a posto. Il rischio sta nelle app vecchie, nelle app custom sviluppate per voi anni fa e negli avanzi di app che non usate più.

## Cosa sostituisce gli script tag

- **App embed block**, parte di una theme app extension, per tutto ciò che compare nella pagina. Si attivano nel theme editor, il che vuol dire che finalmente potete vederli.
- **Web pixel** per gli script che raccolgono solo dati di analytics o di conversione. Funzionano in sandbox e, una volta che l'app li supporta, non richiedono nulla da parte vostra.

## Se un'app non verrà aggiornata

Sostituitela prima di marzo, oppure fate ricostruire la sua funzione. Per un'app custom, il lavoro consiste di solito nello spostare lo script in una theme app extension, oppure in un web pixel se serve solo per il tracciamento. Già che ci siete, eliminate gli script tag delle app che avete disinstallato: la mutation `scriptTagDelete` funziona ancora.

Se avete un'app Shopify custom da migrare, o non siete sicuri di cosa venga caricato nel vostro negozio, [mandateci l'indirizzo del vostro store](https://sawirstudio.com/it/contact/) e vi diremo cosa troviamo. Altre informazioni sul nostro lavoro con Shopify sono disponibili in [questa pagina](https://sawirstudio.com/it/shopify/).

**Continuate a leggere:** [L'API Shopify 2026-10: cosa si rompe nelle vostre integrazioni](https://sawirstudio.com/it/posts/shopify-2026-10-api-what-breaks/), oppure [Collegare Shopify a Odoo 19 Community, passo dopo passo](https://sawirstudio.com/it/posts/connect-shopify-to-odoo-19-community/).
