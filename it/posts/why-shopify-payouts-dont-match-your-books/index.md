---
title: Perché i payout di Shopify non coincidono con la vostra contabilità | SawirStudio
url: https://sawirstudio.com/it/posts/why-shopify-payouts-dont-match-your-books/
description: Perché il denaro che Shopify Payments accredita sulla vostra banca non coincide mai con il report delle vendite: commissioni, rimborsi, contestazioni, gift card, conversione di valuta e tempistiche, e come riconciliare tutto.
---

# Perché i payout di Shopify non coincidono mai con la vostra contabilità

Il report delle vendite dice una cifra. La banca ne dice un'altra. In mezzo ci sono commissioni, rimborsi, chargeback e qualche giorno di sfasamento, e ogni mese qualcuno passa un pomeriggio a separarli in un foglio di calcolo.

Abbiamo affrontato tutto questo scrivendo il modulo di contabilità del nostro [connettore Shopify per Odoo](https://sawirstudio.com/it/projects/odoo-shopify-connector/). Ecco perché le cifre non tornano e come farle quadrare.

## Un payout è un importo netto

Shopify Payments non vi accredita le vendite. Vi accredita ciò che resta dopo aver regolato tutto il resto: vendite, meno rimborsi, meno commissioni di elaborazione, meno contestazioni, più o meno rettifiche. Se registrate il payout come ricavo, i ricavi sono sbagliati e le commissioni diventano invisibili.

La soluzione è registrare le singole voci, non il totale. Le vendite vanno nei ricavi, le commissioni in un conto di costo, i rimborsi a storno delle vendite, e il payout in sé è solo denaro che passa da un conto transitorio alla banca.

## Un payout non copre un solo giorno

Un payout comprende le transazioni di più giorni, a seconda della pianificazione dei payout, e un ordine del venerdì può arrivare nel payout di martedì. Abbinare ordini e accrediti in base alla data funziona poco e per poco. Abbinateli invece per payout: Shopify registra quali transazioni di saldo compongono ciascun payout, e quell'elenco è l'unico collegamento affidabile tra un ordine e il denaro arrivato in banca.

## I rimborsi vengono detratti dai payout successivi

Quando rimborsate un ordine del mese scorso, il denaro esce da un payout di questo mese. Il report delle vendite attribuisce il rimborso all'ordine originale. La banca mostra oggi un accredito più basso. Hanno ragione entrambi.

## Le contestazioni tolgono il denaro due volte

Quando un cliente contesta un addebito, l'importo contestato e una commissione per il chargeback vengono prelevati dal vostro saldo. Se vincete la contestazione, l'importo torna indietro, di solito dopo settimane. Sono tre registrazioni distinte per un solo ordine, distribuite su due o tre payout.

## Le gift card non sono un ricavo quando le vendete

Vendere una gift card significa incassare denaro che dovrete restituire più avanti sotto forma di merce. Va registrato in un conto di debito finché la carta non viene utilizzata. Se registrate le vendite delle gift card come ricavo, contate la stessa vendita due volte: quando la carta viene acquistata e di nuovo quando viene spesa.

## Conversione di valuta

Se vendete in più valute, le conversioni compaiono come transazioni di saldo a sé stanti, e il cambio del giorno della vendita non è quello del giorno del payout. L'API Shopify 2026-10 ha aggiunto un tipo di transazione separato, `CURRENCY_CONVERSION`, così la conversione può essere registrata su una riga propria invece di nascondersi in una differenza di arrotondamento.

## Riserve e rettifiche

Shopify può trattenere una parte del vostro saldo come riserva, e le correzioni compaiono come rettifiche. Né l'una né l'altra è una vendita o una commissione, quindi servono conti dedicati.

## Come far quadrare i conti

Trattate ogni payout come una piccola scrittura contabile costruita a partire dalle sue transazioni di saldo: ogni vendita, rimborso, commissione, contestazione e rettifica su una riga propria, tutte confluite in un conto transitorio che l'accredito bancario poi svuota. Se le righe non tornano con il payout, qualcosa è associato male, ed è utile saperlo.

È così che funziona il nostro modulo di contabilità per Odoo. Pianifica ogni scrittura a partire dagli importi esatti di Shopify e, se la scrittura non quadra, si ferma con un errore chiaro. Non aggiunge mai una «riga di quadratura» per far sparire la differenza, perché quella riga è solo l'errore, scritto dove nessuno andrà a guardare.

Se i soldi di Shopify e la vostra contabilità non coincidono ogni mese, [possiamo aiutarvi](https://sawirstudio.com/it/shopify/).

**Continuate a leggere:** [Collegare Shopify a Odoo 19 Community, passo dopo passo](https://sawirstudio.com/it/posts/connect-shopify-to-odoo-19-community/), oppure [L'API Shopify 2026-10: cosa smette di funzionare nelle vostre integrazioni](https://sawirstudio.com/it/posts/shopify-2026-10-api-what-breaks/).
