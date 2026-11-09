---
title: Sawir Renew – Facturation récurrente pour WooCommerce | SawirStudio
url: https://sawirstudio.com/fr/projects/wordpress-plugins/sawir-renew/
description: Des abonnements complets pour WooCommerce : facturation récurrente, périodes d'essai, frais d'inscription, pause et saut de renouvellement, changement de formule, relances de paiement et nouvelles tentatives intelligentes. Le tout gratuit.
---

# Sawir Renew

![Bannière de Sawir Renew](https://sawirstudio.com/assets/images/plugins/sawir-renew-banner.webp)

Sawir Renew transforme une boutique WooCommerce en activité d'abonnement : facturation récurrente, périodes d'essai, frais d'inscription et relances intelligentes, sans offre payante qui bride une seule fonctionnalité.

[Téléchargez-le gratuitement sur WordPress.org](https://wordpress.org/plugins/sawir-renew/)

## Ce qu'il fait

Vendez des abonnements avec WooCommerce sans payer d'extension. Toutes les fonctionnalités sont gratuites : il n'existe pas de version pro et rien n'est verrouillé.

### Produits et tarifs

- Types de produits d'abonnement simples et variables
- Facturation tous les N jours, semaines, mois ou ans, avec prise en charge des prix promotionnels
- Périodes d'essai gratuites, frais d'inscription et abonnements qui expirent après N paiements
- Renouvellements synchronisés (tout le monde renouvelle le 1er du mois, ou le jour de votre choix), avec première période au prorata ou gratuite en option
- Frais de livraison uniques, limite d'« un abonnement actif par client »

### Deux modes de paiement

- **Paiement à chaque renouvellement (compatible avec toutes les passerelles de paiement) :** lorsqu'un renouvellement arrive à échéance, le client reçoit une facture avec un lien de paiement : virement bancaire, portefeuilles électroniques, paiement à la livraison, tout ce que votre boutique accepte. Idéal pour les marchés où le prélèvement automatique par carte n'est pas la norme.
- **Facturation automatique via Stripe :** si vous utilisez la passerelle officielle WooCommerce Stripe Gateway, les cartes sont enregistrées lors de la commande et les renouvellements sont prélevés automatiquement, sans que le client soit présent. Les relances intelligentes sont incluses, ainsi qu'un lien de paiement de secours lorsqu'une carte exige une authentification. D'autres passerelles peuvent se brancher sur le même contrat d'adaptateur (voir la FAQ).

### Des renouvellements qui ne vous réveillent pas à 3 h du matin

- Toute la planification repose sur Action Scheduler, avec un contrôle quotidien d'auto-réparation
- Relances de paiement configurables : délais entre les tentatives, calendrier des rappels, périodes de grâce, annulation automatique
- La détection des sites de préproduction ou dupliqués suspend la facturation dès que l'URL de votre site change : fini les doubles prélèvements dus à une copie de préproduction oubliée
- E-mails d'avis avant renouvellement, envoyés avant les prélèvements automatiques

### Des clients autonomes

La section Mon compte → Abonnements permet aux clients d'annuler (immédiatement ou à la fin de la période, à vous de choisir), de mettre en pause et de reprendre, de sauter un renouvellement, de renouveler par anticipation, de modifier leur moyen de paiement ou leur adresse de livraison, et de passer d'une formule à une autre, avec un supplément calculé au prorata lors d'une montée en gamme.

### Codes promo récurrents

Quatre types de codes promo pour abonnements : remises récurrentes en pourcentage ou en montant fixe (éventuellement limitées aux N premiers paiements) et remises sur les frais d'inscription. Les codes promo standard continuent de s'appliquer uniquement à la commande initiale.

### Notifications

Treize e-mails activables ou désactivables (client et administrateur) pour chaque étape du cycle de vie d'un abonnement. Les développeurs peuvent répercuter n'importe quel événement vers d'autres canaux (WhatsApp, SMS, Slack, etc.) grâce au hook `swsb_notified`.

### Conçu pour le WooCommerce d'aujourd'hui

Compatible HPOS en natif (les abonnements sont stockés dans le système de commandes de WooCommerce, sans table personnalisée), compatible avec les blocs panier et commande, liste d'administration et éditeur complets reprenant les écrans de commande de WooCommerce, page de rapports mise en cache toutes les heures (abonnements actifs, MRR, taux de désabonnement, prochains renouvellements, revenus des renouvellements) et API REST.

## Questions

### Ce plugin est-il affilié à WooCommerce ou à Automattic ?

Non. « WooCommerce » est une marque d'Automattic ; ce plugin indépendant est conçu *pour* WooCommerce, dans le respect de leurs consignes de dénomination.

### Quelles passerelles prennent en charge les renouvellements automatiques ?

Dès l'installation : la passerelle officielle WooCommerce Stripe Gateway. Toutes les autres passerelles fonctionnent en mode paiement à chaque renouvellement : le client reçoit un lien de paiement à chaque période. Les développeurs peuvent ajouter la prise en charge automatique de n'importe quelle passerelle avec deux filtres :

`swsb_order_supports_auto` : renvoyez true lorsque la passerelle de la commande parente a enregistré un moyen de paiement réutilisable. `swsb_charge_renewal` : prélevez la commande de renouvellement sans que le client soit présent, puis renvoyez `true` ou une `WP_Error`.

### Peut-il fonctionner en même temps que WooCommerce Subscriptions ?

Non : faire tourner deux moteurs de facturation prélèverait deux fois les clients. Ce plugin désactive donc son moteur et affiche un avertissement si WooCommerce Subscriptions est actif.

### Que se passe-t-il sur un site de préproduction ?

Si l'URL du site change, la facturation et les e-mails clients sont automatiquement suspendus jusqu'à ce que vous confirmiez s'il s'agit d'une copie de préproduction ou du site en production. Si vous choisissez « production », tous les calendriers sont reconstruits.

## À associer avec

- [Sawir Referral](https://sawirstudio.com/fr/projects/wordpress-plugins/sawir-referral/) : permet à vos clients de faire venir leurs amis.
- [Sawir Pulse](https://sawirstudio.com/fr/projects/wordpress-plugins/sawir-pulse/) : montre à vos visiteurs ce qui se passe vraiment dans votre boutique (commandes récentes, personnes qui consultent un produit, nouveaux avis) sous forme de petites notifications dans un coin de l'écran.

Découvrez tous nos [plugins WordPress gratuits](https://sawirstudio.com/fr/projects/wordpress-plugins/).

- **Prix** Gratuit
- **Version** 1.0.0
- **Prérequis** : WordPress 6.0, PHP 7.4, WooCommerce
- **Licence** GPL-2.0-or-later
- **WordPress.org** <https://wordpress.org/plugins/sawir-renew/>

Assez parlé de nous.

## Comment pouvons-nous vous aider ?

[Démarrer un projet](https://sawirstudio.com/fr/contact/)
