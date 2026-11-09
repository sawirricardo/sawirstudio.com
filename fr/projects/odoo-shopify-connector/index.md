---
title: Shopify Connector for Odoo | SawirStudio
url: https://sawirstudio.com/fr/projects/odoo-shopify-connector/
description: Un connecteur Shopify open source pour Odoo 19 Community, écrit par SawirStudio et maintenu au sein de l’Odoo Community Association (OCA).
---

# Shopify Connector for Odoo

![Le dépôt OCA connector-shopify sur GitHub](https://sawirstudio.com/assets/images/projects/odoo-shopify-connector.webp)

## Le projet

L’Odoo Community Association (OCA) est l’association à but non lucratif qui maintient les modules open source sur lesquels s’appuient la plupart des utilisateurs d’Odoo Community. Connector Shopify est son connecteur entre les boutiques Shopify et Odoo 19 Community, publié sous licence AGPL-3.

## Ce que nous avons réalisé

Nous avons écrit ce connecteur etl’avons confié à l’OCA en juillet 2026 : le Shopify Connector de base, Shopify Connector Accounting et Shopify Connector POS. Nous continuons de le faire évoluer, tout récemment pour l’aligner sur l’API Admin 2026-01 de Shopify, et d’autres contributeurs de l’OCA s’appuient désormais eux aussi dessus, notamment avec un module pour les médias produits de l’eCommerce.

- **Synchronisation commerciale.** Produits, stocks, clients, commandes, retours, remboursements et expéditions, pour autant de boutiques Shopify que vous le souhaitez par base Odoo, réparties par société.
- **Fiable par conception.** Des webhooks signés pour la rapidité, la queue\_job de l’OCA pour un traitement durable, et un rapprochement planifié qui rattrape les événements manqués ou retardés.
- **Outils du quotidien.** Prise en main, contrôles de santé, tableau de bord de synchronisation, reprise après erreur et durée de conservation des données configurable.
- **Comptabilité.** Versements Shopify Payments, transactions de solde, frais, litiges, passif des cartes cadeaux et rapprochement multidevise.
- **Point de vente.** Les ventes Shopify POS sont importées comme commandes de vente et acheminées depuis chaque point de vente vers un entrepôt Odoo, sans avoir besoin de l’application POS d’Odoo.

## Notre approche

Une intégration juste la plupart du temps est une intégration fausse en comptabilité. C’est pourquoi le connecteur s’appuie sur l’API GraphQL Admin de Shopify, ne présume jamais qu’un webhook est bien arrivé, et se vérifie lui-même auprès de Shopify à intervalles réguliers. Les écritures comptables sont préparées à partir des montants exacts de Shopify et doivent être équilibrées avant d’être validées : quand quelque chose ne colle pas, le connecteur affiche une erreur claire au lieu d’ajouter discrètement une ligne d’équilibrage.

Vous souhaitez le faire installer ou l’adapter à vos boutiques ? Découvrez nos [intégrations Shopify](https://sawirstudio.com/fr/shopify/), ou [Odoo et systèmes de gestion](https://sawirstudio.com/fr/odoo/) pour Odoo lui-même.

**Technologies** Python, Odoo 19 et l’API GraphQL Admin de Shopify

- **Projet** Odoo Community Association (OCA)
- **Licence** AGPL-3
- **Code source** <https://github.com/OCA/connector-shopify>

## Autres projets

[![MyVivaldy](https://sawirstudio.com/assets/images/projects/myvivaldy.webp)

### MyVivaldy](https://sawirstudio.com/fr/projects/myvivaldy/)
[![Koperasi HOB Application](https://sawirstudio.com/assets/images/projects/koperasi-hob-application.webp)

### Koperasi HOB Application](https://sawirstudio.com/fr/projects/koperasi-hob-application/)
[![DBTV Application](https://sawirstudio.com/assets/images/projects/dbtv-application.webp)

### DBTV Application](https://sawirstudio.com/fr/projects/dbtv-application/)

Assez parlé de nous.

## Comment pouvons-nous vous aider ?

[Démarrer un projet](https://sawirstudio.com/fr/contact/)
