---
title: Shopify Connector per Odoo | SawirStudio
url: https://sawirstudio.com/it/projects/odoo-shopify-connector/
description: Un connettore Shopify open source per Odoo 19 Community, scritto da SawirStudio e mantenuto nell’Odoo Community Association (OCA).
---

# Shopify Connector for Odoo

![Il repository OCA connector-shopify su GitHub](https://sawirstudio.com/assets/images/projects/odoo-shopify-connector.webp)

## Il progetto

L’Odoo Community Association (OCA) è l’organizzazione non profit che mantiene i moduli open source su cui si appoggia la maggior parte degli utenti di Odoo Community. Connector Shopify è il suo connettore tra i negozi Shopify e Odoo 19 Community, rilasciato con licenza AGPL-3.

## Cosa abbiamo realizzato

Abbiamo scritto noi il connettore e lo abbiamo contribuito all’OCA a luglio 2026: il Shopify Connector di base, Shopify Connector Accounting e Shopify Connector POS. Continuiamo a lavorarci, da ultimo allineandolo alla Admin API 2026-01 di Shopify, e ormai anche altri collaboratori dell’OCA ci costruiscono sopra, tra cui un modulo per i media dei prodotti eCommerce.

- **Sincronizzazione commerciale.** Prodotti, giacenze, clienti, ordini, resi, rimborsi ed evasione degli ordini, per un numero qualsiasi di negozi Shopify per ogni database Odoo, suddivisi per azienda.
- **Affidabile per progetto.** Webhook firmati per la velocità, queue\_job dell’OCA per un’elaborazione duratura e una riconciliazione pianificata che ripara gli eventi persi o in ritardo.
- **Strumenti per l’uso quotidiano.** Onboarding, controlli di stato, una dashboard di sincronizzazione, ripristino in caso di errori e conservazione dei dati configurabile.
- **Contabilità.** Pagamenti di Shopify Payments, movimenti di saldo, commissioni, contestazioni, passività delle gift card e riconciliazione multivaluta.
- **Punto vendita.** Le vendite di Shopify POS vengono importate come ordini di vendita e instradate da ogni punto vendita fisico a un magazzino Odoo, senza bisogno dell’app POS di Odoo.

## Come ci siamo approcciati al progetto

Un’integrazione che funziona bene quasi sempre, per la contabilità, è sbagliata. Per questo il connettore dialoga con la GraphQL Admin API di Shopify, non dà mai per scontato che un webhook sia arrivato e si confronta con Shopify a intervalli regolari. Le scritture contabili vengono pianificate a partire dagli importi esatti di Shopify e devono quadrare prima di essere registrate: quando qualcosa non torna, il connettore segnala un errore chiaro invece di aggiungere in silenzio una riga di quadratura.

Volete che lo configuriamo o lo estendiamo per i vostri negozi? Consultate le [integrazioni Shopify](https://sawirstudio.com/it/shopify/), oppure [Odoo e sistemi gestionali](https://sawirstudio.com/it/odoo/) per Odoo in sé.

**Realizzato con** Python, Odoo 19 e la GraphQL Admin API di Shopify

- **Progetto** Odoo Community Association (OCA)
- **Licenza** AGPL-3
- **Codice sorgente** <https://github.com/OCA/connector-shopify>

## Altri progetti

[![MyVivaldy](https://sawirstudio.com/assets/images/projects/myvivaldy.webp)

### MyVivaldy](https://sawirstudio.com/it/projects/myvivaldy/)
[![Applicazione Koperasi HOB](https://sawirstudio.com/assets/images/projects/koperasi-hob-application.webp)

### Applicazione Koperasi HOB](https://sawirstudio.com/it/projects/koperasi-hob-application/)
[![Applicazione DBTV](https://sawirstudio.com/assets/images/projects/dbtv-application.webp)

### Applicazione DBTV](https://sawirstudio.com/it/projects/dbtv-application/)

Basta parlare di noi.

## Come possiamo aiutarvi?

[Avviate un progetto](https://sawirstudio.com/it/contact/)
