---
title: Sawir Insights | SawirStudio
url: https://sawirstudio.com/fr/projects/wordpress-plugins/sawir-insights/
description: Cartes de chaleur, enregistrements de sessions, sondages, entonnoirs et analyse des formulaires. Tout reste sur votre serveur.
---

# Sawir Insights

![Bannière Sawir Insights](https://sawirstudio.com/assets/images/plugins/sawir-insights-banner.webp)

Sawir Insights vous montre comment les visiteurs utilisent votre site, avec des cartes de chaleur, des enregistrements de sessions, des sondages, des entonnoirs de conversion et une analyse des formulaires, le tout hébergé sur votre propre serveur.

[Télécharger gratuitement sur WordPress.org](https://wordpress.org/plugins/sawir-insights/)

## Ce qu'il fait

Sawir Insights est une extension d'analyse comportementale pour WordPress : cartes de chaleur, enregistrements de sessions, sondages, widget de feedback, analyse des formulaires, entonnoirs de conversion et un tableau de bord qui réunit le tout.

Sur les réseaux multisites, chaque site suit et configure ses statistiques de façon indépendante.

En bref : une alternative auto-hébergée à Hotjar ou Microsoft Clarity, où les données des visiteurs restent dans votre propre base de données.

La raison d'être de cette extension est simple. Les outils établis dans cette catégorie sont bons, mais ils fonctionnent en envoyant sur les serveurs d'un tiers une trace de tout ce que font vos visiteurs. Pour beaucoup de sites, surtout dans l'UE, c'est au mieux gênant, au pire inenvisageable. Je voulais le même type d'outils, mais avec des données qui ne quittent jamais le site qui les a collectées. C'est donc la contrainte autour de laquelle cette extension est construite : votre base de données, votre répertoire d'uploads, vos règles de conservation. Il n'y a aucun SaaS derrière, aucun compte à créer et aucune communication avec un serveur externe, quelle qu'elle soit.

### Ce qu'il fait

- Cartes de chaleur des clics, des mouvements et du défilement, affichées sur vos pages en ligne, avec une séparation entre ordinateur, tablette et mobile.
- Enregistrements de sessions avec lecture dans wp-admin. Les valeurs saisies sont toujours masquées ; l'enregistrement est désactivé tant que vous ne l'activez pas.
- Sondages en fenêtre contextuelle et onglet de feedback permanent, avec ciblage par URL, appareil, pourcentage de trafic, temps passé sur la page, profondeur de défilement ou intention de sortie.
- Analyse des formulaires : où les gens hésitent, ce qu'ils laissent vide, où ils abandonnent. Uniquement des sélecteurs et des durées, jamais ce que quelqu'un a saisi.
- Entonnoirs de conversion ordonnés avec rapports quotidiens.
- Un tableau de bord avec les indicateurs clés de trafic et leurs variations d'une période à l'autre, un graphique de tendance quotidienne, et des répartitions par appareil, navigateur, page, source de provenance, campagne, entonnoir et enregistrement.

### Ce que l'extension refuse volontairement de faire

Je pense que vous devez savoir ce qu'un outil refuse de faire avant de l'installer.

- Elle ne stocke aucune adresse IP. Ni hachée, ni tronquée : aucune.
- Elle n'enregistre jamais les valeurs des formulaires. La couche de capture n'envoie que des métadonnées, il n'y a donc rien à divulguer.
- Elle ne charge rien depuis un CDN et n'envoie rien vers un CDN. La bibliothèque de relecture rrweb est incluse dans l'extension.
- Elle ne conserve pas les données brutes indéfiniment. Les événements sont agrégés puis supprimés selon un calendrier de conservation que vous contrôlez, et les enregistrements ont leur propre durée de conservation ainsi qu'un plafond d'espace disque, avec suppression des plus anciens en premier, car une extension qui remplit discrètement le disque de votre hébergement est une mauvaise extension.

Le traceur côté frontend est un petit script en JavaScript natif, envoyé par lots via `sendBeacon`, avec des réglages d'échantillonnage. Les visiteurs qui ne sont pas sélectionnés pour l'enregistrement ne téléchargent jamais la bibliothèque de relecture.

### Respect de la vie privée

Le signal Do Not Track est respecté par défaut. Une API de gestion du consentement (`window.swinsConsent`) s'intègre au gestionnaire de consentement que vous utilisez déjà. Les sessions, les événements et les enregistrements sont reliés aux outils d'export et d'effacement des données personnelles de WordPress. La conformité de votre déploiement au RGPD dépend de vos mentions d'information et de votre base légale, et aucune extension ne peut s'en charger à votre place, mais le travail de minimisation des données est fait.

### Assistance

L'assistance passe par le forum wordpress.org de cette extension, ou par e-mail à ricardo@sawirstudio.com.

## Questions

### Comparaison avec les alternatives

Les services d'analyse hébergés comme Hotjar et Microsoft Clarity sont d'excellents outils, avec des offres gratuites généreuses, mais ils fonctionnent en envoyant une trace du comportement de vos visiteurs vers leurs serveurs. Sawir Insights renonce à cette infrastructure partagée entre sites au profit d'une seule contrainte : tout est collecté, stocké et consulté sur votre propre site. Choisissez un service hébergé si vous voulez ses fonctions supplémentaires et que le traitement par un tiers ne vous dérange pas ; choisissez Sawir Insights si le fait que les données des visiteurs quittent votre serveur est rédhibitoire. Les enregistrements et les cartes de chaleur sont stockés dans votre propre base de données et votre répertoire d'uploads : les réglages de conservation comptent donc davantage ici que dans un SaaS.

### Sawir Insights est-il conforme au RGPD ?

L'extension est conçue pour les déploiements soucieux de la vie privée : les données restent sur votre serveur, les adresses IP ne sont jamais stockées, la conservation est configurable, Do Not Track est respecté et les outils d'export et d'effacement de WordPress sont intégrés. La conformité elle-même dépend de votre configuration, de vos mentions d'information, de votre base légale et de la mise en œuvre du consentement : consultez un juriste qualifié pour votre site.

### En quoi est-ce différent de Hotjar ou de Microsoft Clarity ?

Ce sont des services hébergés : des outils performants, mais le comportement de vos visiteurs est enregistré sur leurs serveurs, selon leurs conditions. Sawir Insights couvre la même catégorie (cartes de chaleur, enregistrements de sessions, sondages, entonnoirs, analyse des formulaires) sous forme d'extension qui garde tout sur votre propre serveur, sans compte, sans service externe et sans partage de données.

### Les données des visiteurs quittent-elles mon serveur ?

Non. L'extension n'envoie aucune donnée comportementale à Sawir Studio ni à aucun fournisseur d'analyse. Les événements sont stockés dans votre base de données WordPress ; les fichiers d'enregistrement se trouvent dans votre répertoire d'uploads protégé.

## À associer avec

- [Sawir Pulse](https://sawirstudio.com/fr/projects/wordpress-plugins/sawir-pulse/) : montre à vos visiteurs ce qui se passe vraiment dans votre boutique (commandes récentes, personnes qui consultent un produit, nouveaux avis) sous forme de petites notifications dans un coin de l'écran.
- [Sawir Live Translate](https://sawirstudio.com/fr/projects/wordpress-plugins/sawir-live-translate/) : traduit vos pages directement là où elles se trouvent : choisissez une langue, traduisez automatiquement ce qui est affiché, enregistrez, puis proposez à vos visiteurs un sélecteur de langue.

Découvrez tous nos [plugins WordPress gratuits](https://sawirstudio.com/fr/projects/wordpress-plugins/).

- **Prix** Gratuit
- **Version** 1.0.0
- **Prérequis** : WordPress 6.2, PHP 7.4
- **Licence** GPLv2 ou ultérieure
- **WordPress.org** <https://wordpress.org/plugins/sawir-insights/>

Assez parlé de nous.

## Comment pouvons-nous vous aider ?

[Démarrer un projet](https://sawirstudio.com/fr/contact/)
