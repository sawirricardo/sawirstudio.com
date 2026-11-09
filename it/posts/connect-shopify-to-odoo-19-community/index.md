---
title: Collegare Shopify a Odoo 19 Community | SawirStudio
url: https://sawirstudio.com/it/posts/connect-shopify-to-odoo-19-community/
description: Come collegare un negozio Shopify a Odoo 19 Community con il connettore Shopify open source di OCA: app nella Dev Dashboard, scope, queue_job, webhook e onboarding.
---

# Collegare Shopify a Odoo 19 Community, passo dopo passo

Abbiamo scritto il [connettore Shopify per Odoo](https://sawirstudio.com/it/projects/odoo-shopify-connector/) e lo abbiamo donato alla Odoo Community Association (OCA). È gratuito, open source (AGPL-3) e funziona con Odoo 19 Community. Questa è la configurazione, dall'inizio alla fine, come la facciamo per i nostri clienti.

Sincronizza prodotti, giacenze, clienti, ordini, resi, rimborsi ed evasione degli ordini, per un solo negozio o per più negozi. Contabilità e punto vendita sono moduli separati che potete aggiungere in seguito.

## Prima di iniziare

- Odoo 19 Community raggiungibile tramite HTTPS pubblico. Shopify invia i webhook a questo indirizzo e da qui scarica le immagini dei prodotti.
- Accesso come amministratore al negozio Shopify e all'organizzazione Shopify a cui appartiene.
- Un piano Shopify che includa i dati protetti dei clienti per le app create da voi. Il piano Basic non li include e, senza, Shopify rifiuta alcuni webhook degli ordini e l'importazione di ordini e clienti non funziona.

## 1. Creare l'app nella Dev Dashboard di Shopify

Dal 1° gennaio 2026 Shopify non permette più di creare le vecchie «app personalizzate» dall'area amministrativa del negozio. Create l'app nella **Dev Dashboard** dell'organizzazione a cui appartiene il negozio. È importante: Odoo ottiene il suo token di accesso con il client credentials grant, che funziona solo se l'app e il negozio appartengono alla stessa organizzazione.

Quando compilate la versione dell'app:

- **App URL:** `https://shopify.dev/apps/default-app-home`. Il connettore non ha nessuna pagina dentro Shopify.
- **Embed app in Shopify admin:** disattivato.
- **Use legacy install flow:** disattivato, così l'installazione dell'app concede gli scope che avete scelto.
- **Webhooks:** nessuno. Odoo crea i propri durante l'onboarding.
- **Webhooks API version:** `2026-01`, la versione usata dal connettore.
- **Scopes:** `read_customers`, `write_customers`, `read_draft_orders`, `read_fulfillments`, `read_inventory`, `write_inventory`, `read_locations`, `read_merchant_managed_fulfillment_orders`, `write_merchant_managed_fulfillment_orders`, `read_orders`, `read_products`, `write_products` e `read_returns`. Aggiungete `read_all_orders` se volete gli ordini più vecchi di 60 giorni.

Rilasciate la versione, installate l'app sul negozio e copiate client ID e client secret dalla pagina Impostazioni dell'app.

Avete già una vecchia custom app creata prima del 2026? Funziona ancora. Lasciate vuoto il client ID e usate il suo Admin API access token e la sua API secret key.

## 2. Installare i moduli in Odoo

Aggiungete OCA `queue_job` e il repository [connector-shopify](https://github.com/OCA/connector-shopify) al percorso degli addons, poi installate **Shopify Connector**. Il lavoro viene eseguito come job in background, quindi `queue_job` deve essere caricato a livello di server, con una capacità per ogni canale Shopify:

```
[options]
server_wide_modules = base,rpc,web,queue_job

[queue_job]
channels = root:4,root.shopify.import:2,root.shopify.export:1,root.shopify.webhook:1
```

Poi impostate il parametro di sistema `web.base.url` sull'indirizzo HTTPS pubblico di Odoo e `web.base.url.freeze` su `True`. Altrimenti l'indirizzo cambia quando un amministratore accede da un altro posto e i webhook di Shopify finiscono nel posto sbagliato.

## 3. Eseguire l'onboarding

Assegnate al vostro team il gruppo **Shopify Manager**, poi aprite **Shopify > Configuration > Instances** e fate clic su **Start Onboarding**:

1. Inserite il dominio `yourstore.myshopify.com`, il client ID e il client secret, e verificate la connessione. Odoo richiede il token e lo rinnova da solo ogni 24 ore.
2. Controllate gli scope. Se ne manca uno, aggiungetelo a una nuova versione dell'app, rilasciatela e approvatela sul negozio.
3. Associate le sedi Shopify ai magazzini Odoo, controllate le mappature delle imposte generate da Odoo e scegliete le politiche di sincronizzazione.
4. Installate i webhook e mettete in coda le prime importazioni.

Prima di lasciare ordini ed evasioni in funzione senza supervisione, controllate le mappature dei gateway di pagamento e dei corrieri. La maggior parte delle sorprese successive nasce da una mappatura che nessuno ha guardato.

## 4. Tenere tutto sotto controllo

**Shopify > Dashboard** mostra lo stato delle sincronizzazioni, i job in coda, gli errori recenti e lo stato dei webhook. Il connettore non dà per scontato che ogni webhook arrivi: dei job pianificati importano di nuovo tutto ciò che Shopify ha modificato e Odoo si è perso. Quando qualcosa non funziona, risolvete la causa, poi usate **Retry Failed Shopify Bindings**. Eseguite **Check Webhooks** dopo aver cambiato dominio o credenziali.

## Prossimi passi

Aggiungete **Shopify Connector Accounting** per portare in Odoo i payout, le commissioni e le contestazioni di Shopify Payments come scritture contabili (spieghiamo a parte [perché è difficile riconciliare i payout](https://sawirstudio.com/it/posts/why-shopify-payouts-dont-match-your-books/)), e **Shopify Connector POS** per le vendite in negozio.

Se volete che lo installiamo, lo configuriamo in base a come lavora la vostra azienda o lo estendiamo, [ce ne occupiamo noi](https://sawirstudio.com/it/shopify/).

**Continuate a leggere:** [Perché i payout di Shopify non coincidono mai con la vostra contabilità](https://sawirstudio.com/it/posts/why-shopify-payouts-dont-match-your-books/), oppure [L'API Shopify 2026-10: cosa si rompe nelle vostre integrazioni](https://sawirstudio.com/it/posts/shopify-2026-10-api-what-breaks/).
