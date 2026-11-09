---
title: API Shopify 2026-10: cosa si rompe | SawirStudio
url: https://sawirstudio.com/it/posts/shopify-2026-10-api-what-breaks/
description: Le modifiche della versione 2026-10 della Shopify Admin API che possono rompere un'integrazione e cosa cambiare: tariffe di spedizione, sconti delle bozze d'ordine, imposte, codici a barre e altro.
---

# API 2026-10 di Shopify: cosa si rompe nelle vostre integrazioni

Il 1° ottobre 2026 Shopify ha rilasciato la versione 2026-10 delle sue API. Gran parte delle novità sono funzionalità che potete ignorare finché non servono. Alcune modifiche, però, possono rompere in silenzio un'integrazione che con la 2026-07 funzionava senza problemi: sono quelle per cui vale la pena dedicare un'ora prima di cambiare versione.

Questa è la lista che controlliamo sulle nostre integrazioni, compreso il [connettore Shopify per Odoo](https://sawirstudio.com/it/projects/odoo-shopify-connector/). Non comprende tutte le righe delle note di rilascio, ma solo quelle che cambiano il comportamento.

## Modifiche che rompono le cose in silenzio

### I nuovi carrier service non mostrano le tariffe al checkout da soli

Creare un carrier service con `carrierServiceCreate` (o con `POST carrier_services.json` in REST) ora si limita a registrare il servizio. Le sue tariffe non vengono più aggiunte al profilo di spedizione generale. Se la vostra app o il vostro script di onboarding crea un carrier service e si aspetta che i clienti vedano le vostre tariffe, non le vedranno finché non aggiungete voi il servizio a un profilo di spedizione.

### Cambiare l'indirizzo di spedizione ricalcola le imposte

Aggiornare l'indirizzo di spedizione di un ordine non ancora evaso con `orderUpdate`, o con l'endpoint REST di aggiornamento degli ordini, ora ne ricalcola le imposte. È il comportamento giusto, ma se modificate gli indirizzi da un ERP o da uno strumento di assistenza, il totale dell'ordine può cambiare senza che ve ne accorgiate. Dopo l'aggiornamento rileggete `taxLines`, `totalTaxSet` e i totali, e regolate l'eventuale differenza tra quanto pagato e quanto rimborsato.

### Gli avvisi sugli sconti delle bozze d'ordine hanno perso `priceRule`

`DraftOrderDiscountNotAppliedWarning.priceRule` e i tipi legacy `PriceRule` sono stati rimossi. Se create bozze d'ordine con `draftOrderCalculate`, `draftOrderCreate` o `draftOrderUpdate` e leggete il motivo per cui uno sconto non è stato applicato, selezionate invece `discountTitle` e `discountCode`.

### I filtri sui metafield errati ora falliscono in modo esplicito

Una query sui prodotti che filtrava su un metafield non filtrabile ignorava quella parte del filtro e restituiva tutto. Ora restituisce un errore. È meglio così, ma una sincronizzazione che da mesi restituisce «tutti i prodotti» inizierà a fallire il giorno in cui cambiate versione. Verificate che le definizioni dei metafield su cui filtrate supportino il filtro e gli operatori che usate.

### Una variante può avere più di un codice a barre

`ProductVariant.barcodes` è una nuova connection che contiene fino a 20 codici a barre tipizzati. Il vecchio `ProductVariant.barcode` è deprecato e restituisce solo il primo. Se il vostro magazzino o il vostro punto vendita associa i prodotti tramite codice a barre, leggete l'elenco completo.

### Un errore di inventario non esiste più

`ITEM_NOT_STOCKED_AT_LOCATION` è stato rimosso dalle mutation di inventario e da `InventorySetQuantitiesUserErrorCode`. Il codice che si basa su questo valore ora è codice morto: eliminatelo, così nessuno ci farà affidamento in futuro.

## Modifiche minori da cercare nel vostro codice

- **Segmenti di clienti.** `MATCHES` e `NOT MATCHES` sostituiscono `= true` e `= false`, e le date con nome come `12_months_ago` sono deprecate a favore di offset come `-12m`.
- **Consenso SMS.** I campi piatti per il consenso SMS sono deprecati. Leggete `CustomerPhoneNumber.smsMarketingConsent`.
- **Sconti nei rollout.** Quando uno sconto fa parte di un rollout, `startsAt` e `endsAt` non dicono se è attivo. Per leggere `Discount.rollouts` serve lo scope `read_rollouts`.
- **Stato di evasione.** `FULFILLMENT_NOT_REQUIRED` è un nuovo valore. Uno `switch` che gestisce ogni stato finirà nel caso predefinito.
- **Customer Account API.** `Customer.lastIncompleteCheckout` e i tipi `Checkout` sono stati rimossi, senza alternative.
- **Estensioni UI per POS.** Le mutation del carrello come `applyCartDiscount` ora vengono rifiutate se il POS non conferma la modifica, e `session.currentSession.staffMemberId` è stato rimosso.

## Due scadenze che non riguardano la 2026-10

Se gestite un'app pubblica, i token di accesso offline senza scadenza vanno sostituiti con token a scadenza entro il 1° gennaio 2027. Sono escluse le app personalizzate e le app create da un merchant nella Dev Dashboard. Inoltre, gli script tag dello store online smettono di funzionare il 1° marzo 2027: abbiamo spiegato a parte [cosa significa per i negozi](https://sawirstudio.com/it/posts/shopify-script-tags-stop-march-2027/).

## Come aggiorniamo

Fissiamo la versione dell'API in un unico punto, facciamo girare l'integrazione su un development store con la nuova versione e leggiamo il [changelog della 2026-10](https://shopify.dev/changelog?api_version=2026-10) filtrato su «Action required». La versione 2026-10 è supportata almeno fino al 1° ottobre 2027, quindi c'è tempo, ma non così tanto da aspettare l'arrivo delle email di avviso.

Se il vostro negozio dialoga con Odoo o con un altro sistema e preferite che se ne occupi qualcun altro, [è quello che facciamo](https://sawirstudio.com/it/shopify/).

**Continuate a leggere:** [Gli script tag di Shopify smettono di funzionare il 1° marzo 2027: cosa devono fare i titolari dei negozi](https://sawirstudio.com/it/posts/shopify-script-tags-stop-march-2027/), oppure [Collegare Shopify a Odoo 19 Community, passo dopo passo](https://sawirstudio.com/it/posts/connect-shopify-to-odoo-19-community/).
