---
title: Pourquoi les versements Shopify ne correspondent pas à votre comptabilité | SawirStudio
url: https://sawirstudio.com/fr/posts/why-shopify-payouts-dont-match-your-books/
description: Pourquoi l'argent que Shopify Payments verse sur votre compte bancaire ne correspond jamais à votre rapport de ventes : frais, remboursements, litiges, cartes cadeaux, conversion de devises et décalages, et comment rapprocher le tout.
---

# Pourquoi vos versements Shopify ne correspondent jamais à votre comptabilité

Votre rapport de ventes indique un montant. Votre banque en affiche un autre. Entre les deux, il y a des frais, des remboursements, des rétrofacturations et quelques jours de décalage, et chaque mois, quelqu'un passe un après-midi à démêler tout ça dans un tableur.

Nous avons rencontré tous ces cas en développant le module comptable de notre [connecteur Shopify pour Odoo](https://sawirstudio.com/fr/projects/odoo-shopify-connector/). Voici pourquoi les chiffres diffèrent, et comment les faire concorder.

## Un versement est un montant net

Shopify Payments ne vous verse pas vos ventes. Il vous verse ce qu'il reste une fois tout le reste réglé : les ventes, moins les remboursements, moins les frais de traitement, moins les litiges, plus ou moins les ajustements. Si vous comptabilisez le versement comme du chiffre d'affaires, votre chiffre d'affaires est faux et vos frais deviennent invisibles.

La solution consiste à comptabiliser les composantes, pas le total. Les ventes vont en chiffre d'affaires, les frais dans un compte de charges, les remboursements en déduction des ventes, et le versement lui-même n'est qu'un mouvement d'argent entre un compte d'attente et la banque.

## Un versement ne couvre pas une seule journée

Un versement regroupe les transactions de plusieurs jours, selon votre calendrier de versement, et une commande passée le vendredi peut arriver dans le versement du mardi. Rapprocher les commandes et les dépôts par date ne fonctionne jamais longtemps. Rapprochez-les plutôt par versement : Shopify enregistre les transactions de solde qui composent chaque versement, et cette liste est le seul lien fiable entre une commande et l'argent reçu en banque.

## Les remboursements sont déduits des versements suivants

Lorsque vous remboursez une commande du mois dernier, l'argent est prélevé sur un versement de ce mois-ci. Votre rapport de ventes rattache le remboursement à la commande d'origine. Votre banque affiche aujourd'hui un dépôt plus faible. Les deux ont raison.

## Les litiges prélèvent de l'argent deux fois

Lorsqu'un client conteste un paiement, le montant contesté et des frais de rétrofacturation sont prélevés sur votre solde. Si vous obtenez gain de cause, le montant vous est restitué, généralement plusieurs semaines plus tard. Cela fait trois écritures distinctes pour une seule commande, réparties sur deux ou trois versements.

## Une carte cadeau n'est pas du chiffre d'affaires au moment de sa vente

Vendre une carte cadeau, c'est encaisser de l'argent que vous devrez rendre plus tard sous forme de marchandises. Cette somme relève d'un compte de passif jusqu'à l'utilisation de la carte. Si vous comptabilisez les ventes de cartes cadeaux en chiffre d'affaires, vous comptez la même vente deux fois : une fois à l'achat de la carte, une autre fois quand elle est utilisée.

## La conversion de devises

Si vous vendez dans plusieurs devises, les conversions apparaissent comme des transactions de solde à part entière, et le taux du jour de la vente n'est pas celui du jour du versement. L'API Shopify 2026-10 a ajouté un type de transaction `CURRENCY_CONVERSION` distinct, ce qui permet de comptabiliser la conversion sur sa propre ligne au lieu de la laisser se cacher dans un écart d'arrondi.

## Réserves et ajustements

Shopify peut retenir une partie de votre solde sous forme de réserve, et les corrections apparaissent sous forme d'ajustements. Ce ne sont ni des ventes ni des frais : chacun nécessite donc son propre compte.

## Comment équilibrer vos comptes

Traitez chaque versement comme une petite écriture comptable construite à partir de ses transactions de solde : chaque vente, remboursement, frais, litige et ajustement sur sa propre ligne, le tout aboutissant dans un compte d'attente que le dépôt bancaire vient ensuite solder. Si les lignes ne totalisent pas le montant du versement, c'est qu'un élément est mal affecté, et il vaut mieux le savoir.

C'est ainsi que fonctionne notre module comptable pour Odoo. Il prépare chaque écriture à partir des montants exacts de Shopify, et si l'écriture n'est pas équilibrée, il s'arrête avec un message d'erreur clair. Il n'ajoute jamais de « ligne d'équilibrage » pour faire disparaître l'écart, car cette ligne n'est que l'erreur, consignée là où personne n'ira la chercher.

Si votre argent Shopify et votre comptabilité ne concordent jamais d'un mois à l'autre, [nous pouvons vous aider](https://sawirstudio.com/fr/shopify/).

**À lire aussi :** [Connecter Shopify à Odoo 19 Community, pas à pas](https://sawirstudio.com/fr/posts/connect-shopify-to-odoo-19-community/), ou [API Shopify 2026-10 : ce qui risque de casser dans vos intégrations](https://sawirstudio.com/fr/posts/shopify-2026-10-api-what-breaks/).
