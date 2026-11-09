---
title: API Shopify 2026-10 : ce qui casse | SawirStudio
url: https://sawirstudio.com/fr/posts/shopify-2026-10-api-what-breaks/
description: Les changements de l'API Admin de Shopify dans la version 2026-10 qui peuvent casser une intégration, et ce qu'il faut modifier : tarifs de livraison, remises des commandes brouillon, taxes, codes-barres et plus encore.
---

# API Shopify 2026-10 : ce qui peut casser dans vos intégrations

Shopify a publié la version 2026-10 de ses API le 1er octobre 2026. L’essentiel, ce sont de nouvelles fonctionnalités que vous pouvez ignorer tant que vous n’en avez pas besoin. Quelques changements peuvent toutefois casser sans bruit une intégration qui fonctionnait très bien en 2026-07, et ce sont ceux-là qui méritent une heure de votre temps avant de changer de version.

Voici la liste que nous vérifions sur nos propres intégrations, notamment le [connecteur Shopify pour Odoo](https://sawirstudio.com/fr/projects/odoo-shopify-connector/). Elle ne reprend pas toutes les lignes des notes de version, seulement celles qui modifient le comportement.

## Les changements qui cassent sans prévenir

### Les nouveaux transporteurs n'affichent plus leurs tarifs à la commande d'eux-mêmes

Créer un service de transport avec `carrierServiceCreate` (ou `POST carrier_services.json` en REST) enregistre désormais uniquement le service. Ses tarifs ne sont plus ajoutés au profil d'expédition général. Si votre application ou votre script d'installation crée un service de transport et compte sur l'affichage de vos tarifs pour les clients, ils n'apparaîtront pas tant que vous n'aurez pas ajouté le service à un profil d'expédition vous-même.

### Modifier une adresse de livraison recalcule les taxes

Mettre à jour l'adresse de livraison d'une commande non expédiée avec `orderUpdate`, ou avec le point d'accès REST de mise à jour des commandes, recalcule désormais ses taxes. C'est le bon comportement, mais si vous modifiez les adresses depuis un ERP ou un outil de support, le total de la commande peut changer sans que vous l'ayez prévu. Relisez `taxLines`, `totalTaxSet` et les totaux après la mise à jour, puis réglez toute différence sur ce qui a été payé ou remboursé.

### Les avertissements de remise des commandes brouillon n'ont plus `priceRule`

`DraftOrderDiscountNotAppliedWarning.priceRule` et les anciens types `PriceRule` ont disparu. Si vous créez des commandes brouillon avec `draftOrderCalculate`, `draftOrderCreate` ou `draftOrderUpdate` et que vous lisez la raison pour laquelle une remise n'a pas été appliquée, sélectionnez plutôt `discountTitle` et `discountCode`.

### Les filtres de métachamps invalides échouent désormais explicitement

Auparavant, une requête produit filtrant sur un métachamp non filtrable ignorait cette partie du filtre et renvoyait tout. Elle renvoie maintenant une erreur. C'est mieux, mais une synchronisation qui renvoie « tous les produits » depuis des mois se mettra à échouer le jour où vous changerez de version. Vérifiez que les définitions de métachamps sur lesquelles vous filtrez prennent bien en charge le filtrage, ainsi que les opérateurs que vous utilisez.

### Une variante peut avoir plusieurs codes-barres

`ProductVariant.barcodes` est une nouvelle connexion qui peut contenir jusqu'à 20 codes-barres typés. L'ancien champ `ProductVariant.barcode` est obsolète et ne renvoie que le premier. Si votre entrepôt ou votre point de vente identifie les produits par code-barres, lisez la liste complète.

### Une erreur d'inventaire a disparu

`ITEM_NOT_STOCKED_AT_LOCATION` est supprimée des mutations d'inventaire et de `InventorySetQuantitiesUserErrorCode`. Le code qui s'appuie dessus est désormais du code mort : supprimez-le pour que personne ne s'y fie plus tard.

## Les changements mineurs à rechercher dans votre code

- **Segments de clients.** `MATCHES` et `NOT MATCHES` remplacent `= true` et `= false`, et les dates nommées comme `12_months_ago` sont obsolètes au profit de décalages comme `-12m`.
- **Consentement SMS.** Les champs de consentement SMS à plat sont obsolètes. Lisez `CustomerPhoneNumber.smsMarketingConsent`.
- **Remises dans les déploiements progressifs.** Les champs `startsAt` et `endsAt` d'une remise ne vous indiquent pas si elle est active lorsqu'elle fait partie d'un déploiement progressif. La lecture de `Discount.rollouts` nécessite le scope `read_rollouts`.
- **Statut d'exécution.** `FULFILLMENT_NOT_REQUIRED` est une nouvelle valeur. Un `switch` qui traite chaque statut tombera dans son cas par défaut.
- **Customer Account API.** `Customer.lastIncompleteCheckout` et les types `Checkout` sont supprimés, sans remplacement.
- **Extensions d'interface POS.** Les mutations du panier comme `applyCartDiscount` sont désormais rejetées lorsque le POS ne confirme pas la modification, et `session.currentSession.staffMemberId` est supprimé.

## Deux échéances qui ne concernent pas la version 2026-10

Si vous maintenez une application publique, les jetons d’accès hors ligne sans expiration doivent être remplacés par des jetons avec expiration avant le 1er janvier 2027. Les applications personnalisées et les applications qu’un marchand crée dans le Dev Dashboard sont exclues. Par ailleurs, les script tags de la boutique en ligne cessent de s’exécuter le 1er mars 2027 ; nous avons rédigé un article à part sur [ce que cela implique pour les boutiques](https://sawirstudio.com/fr/posts/shopify-script-tags-stop-march-2027/).

## Notre méthode de mise à niveau

Nous fixons la version de l’API à un seul endroit, nous testons l’intégration sur une boutique de développement avec la nouvelle version, puis nous lisons le [journal des modifications 2026-10](https://shopify.dev/changelog?api_version=2026-10) filtré sur « Action required ». La version 2026-10 est prise en charge jusqu’au 1er octobre 2027 au moins : vous avez donc du temps, mais pas au point d’attendre les premiers e-mails d’avertissement.

Si votre boutique communique avec Odoo ou un autre système et que vous préférez confier ce travail à quelqu'un d'autre, c'est [notre métier](https://sawirstudio.com/fr/shopify/).

**À lire aussi :** [Shopify Script Tags Stop Working on March 1, 2027. What Store Owners Should Do](https://sawirstudio.com/fr/posts/shopify-script-tags-stop-march-2027/), ou [Connecting Shopify to Odoo 19 Community, Step by Step](https://sawirstudio.com/fr/posts/connect-shopify-to-odoo-19-community/).
