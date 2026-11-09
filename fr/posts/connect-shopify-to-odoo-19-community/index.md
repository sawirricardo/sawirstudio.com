---
title: Connecter Shopify à Odoo 19 Community | SawirStudio
url: https://sawirstudio.com/fr/posts/connect-shopify-to-odoo-19-community/
description: Comment connecter une boutique Shopify à Odoo 19 Community avec le connecteur Shopify open source de l’OCA : application Dev Dashboard, scopes, queue_job, webhooks et configuration initiale.
---

# Connecter Shopify à Odoo 19 Community, étape par étape

Nous avons développé le [connecteur Shopify pour Odoo](https://sawirstudio.com/fr/projects/odoo-shopify-connector/) et l’avons confié à l’Odoo Community Association (OCA). Il est gratuit, open source (AGPL-3) et compatible avec Odoo 19 Community. Voici la mise en place, du début à la fin, telle que nous la réalisons pour nos clients.

Il synchronise les produits, les stocks, les clients, les commandes, les retours, les remboursements et les expéditions, pour une ou plusieurs boutiques. La comptabilité et le point de vente sont des modules distincts que vous pouvez ajouter plus tard.

## Avant de commencer

- Une instance Odoo 19 Community accessible en HTTPS depuis Internet. Shopify lui envoie des webhooks et y télécharge les images des produits.
- Un accès administrateur à la boutique Shopify, ainsi qu’à l’organisation Shopify à laquelle elle appartient.
- Une offre Shopify qui inclut l’accès aux données client protégées pour les applications que vous créez vous-même. Ce n’est pas le cas de l’offre Basic : sans cet accès, Shopify refuse certains webhooks de commande et les imports de commandes et de clients échouent.

## 1. Créer l’application dans le Dev Dashboard de Shopify

Depuis le 1er janvier 2026, Shopify ne permet plus de créer les anciennes « applications personnalisées » depuis l’administration de la boutique. Créez l’application dans le **Dev Dashboard** de l’organisation propriétaire de la boutique. C’est important : Odoo obtient son jeton d’accès via le flux « client credentials », qui ne fonctionne que si l’application et la boutique appartiennent à la même organisation.

Lorsque vous renseignez la version de l’application :

- **URL de l’application :** `https://shopify.dev/apps/default-app-home`. Le connecteur n’a pas de page dans Shopify.
- **Intégrer l’application dans l’administration Shopify :** désactivé.
- **Utiliser l’ancien flux d’installation :** désactivé, afin que l’installation de l’application accorde les autorisations (scopes) que vous avez choisies.
- **Webhooks :** aucun. Odoo crée les siens lors de la configuration initiale.
- **Version de l’API des webhooks :** `2026-01`, la version utilisée par le connecteur.
- **Scopes :** `read_customers`, `write_customers`, `read_draft_orders`, `read_fulfillments`, `read_inventory`, `write_inventory`, `read_locations`, `read_merchant_managed_fulfillment_orders`, `write_merchant_managed_fulfillment_orders`, `read_orders`, `read_products`, `write_products` et `read_returns`. Ajoutez `read_all_orders` si vous souhaitez récupérer les commandes de plus de 60 jours.

Publiez la version, installez l’application sur la boutique, puis copiez l’identifiant client (client ID) et la clé secrète (client secret) depuis la page Paramètres de l’application.

Vous avez déjà une ancienne application personnalisée créée avant 2026 ? Elle fonctionne toujours. Laissez le client ID vide et utilisez à la place son jeton d’accès à l’API Admin et sa clé secrète d’API.

## 2. Installer les modules dans Odoo

Ajoutez OCA `queue_job` et le dépôt [connector-shopify](https://github.com/OCA/connector-shopify) à votre chemin d’addons, puis installez **Shopify Connector**. Le travail s’exécute en tâches de fond : `queue_job` doit donc être chargé au niveau du serveur, avec une capacité pour chaque canal Shopify :

```
[options]
server_wide_modules = base,rpc,web,queue_job

[queue_job]
channels = root:4,root.shopify.import:2,root.shopify.export:1,root.shopify.webhook:1
```

Définissez ensuite le paramètre système `web.base.url` sur l’adresse HTTPS publique d’Odoo et `web.base.url.freeze` sur `True`. Sinon, l’adresse change lorsqu’un administrateur se connecte depuis un autre endroit, et les webhooks de Shopify partent au mauvais endroit.

## 3. Lancer la configuration initiale

Attribuez à votre équipe le groupe **Shopify Manager**, puis ouvrez **Shopify > Configuration > Instances** et cliquez sur **Start Onboarding** :

1. Saisissez le domaine `yourstore.myshopify.com`, le client ID et le client secret, puis testez la connexion. Odoo demande le jeton et le renouvelle seul toutes les 24 heures.
2. Vérifiez les scopes. S’il en manque un, ajoutez-le dans une nouvelle version de l’application, publiez-la et approuvez-la sur la boutique.
3. Associez les emplacements Shopify aux entrepôts Odoo, vérifiez les correspondances de taxes générées par Odoo et choisissez les politiques de synchronisation.
4. Installez les webhooks et lancez les premiers imports.

Avant de laisser les commandes et les expéditions tourner sans surveillance, vérifiez les correspondances des passerelles de paiement et des transporteurs. La plupart des mauvaises surprises viennent ensuite d’une correspondance que personne n’a regardée.

## 4. Garder un œil dessus

**Shopify > Dashboard** affiche l’état des synchronisations, la file des tâches en attente, les erreurs récentes et l’état des webhooks. Le connecteur ne suppose pas que tous les webhooks arrivent : des tâches planifiées réimportent tout ce que Shopify a modifié et qu’Odoo a manqué. Quand quelque chose échoue, corrigez la cause, puis utilisez **Retry Failed Shopify Bindings**. Lancez **Check Webhooks** après tout changement de domaine ou d’identifiants.

## Pour aller plus loin

Ajoutez **Shopify Connector Accounting** pour intégrer les virements, les frais et les litiges de Shopify Payments sous forme d’écritures comptables (nous expliquons séparément [pourquoi les virements sont difficiles à rapprocher](https://sawirstudio.com/fr/posts/why-shopify-payouts-dont-match-your-books/)), et **Shopify Connector POS** pour les ventes en magasin.

Si vous souhaitez que nous l’installions, l’adaptions au fonctionnement de votre entreprise ou le fassions évoluer, [c’est ce que nous faisons](https://sawirstudio.com/fr/shopify/).

**À lire aussi :** [Pourquoi vos virements Shopify ne correspondent jamais à votre comptabilité](https://sawirstudio.com/fr/posts/why-shopify-payouts-dont-match-your-books/), ou [L’API Shopify 2026-10 : ce qui risque de casser dans vos intégrations](https://sawirstudio.com/fr/posts/shopify-2026-10-api-what-breaks/).
