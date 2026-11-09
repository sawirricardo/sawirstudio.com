---
title: Fin des script tags Shopify en mars 2027 | SawirStudio
url: https://sawirstudio.com/fr/posts/shopify-script-tags-stop-march-2027/
description: Shopify cesse d'exécuter les script tags des boutiques en ligne le 1er mars 2027. Les applications concernées, comment vérifier votre boutique et ce qui les remplace.
---

# Les script tags Shopify cesseront de fonctionner le 1er mars 2027. Ce que les propriétaires de boutique doivent faire

Le 1er mars 2027, Shopify cessera d'injecter des script tags dans les boutiques en ligne. Tout ce qui, sur votre boutique, se charge encore de cette façon (widget d'avis, bulle de chat, pixel de suivi) disparaîtra tout simplement. Rien ne plantera bruyamment. Des éléments vont simplement manquer.

Shopify l'a annoncé en [août 2026](https://shopify.dev/changelog/posts/online-store-script-tags-deprecation). Voici ce que cela signifie si vous gérez une boutique.

## Qu'est-ce qu'un script tag

Un script tag permet à une application d'ajouter du JavaScript à votre boutique via l'API, sans toucher à votre thème. C'est ce qui plaisait aux développeurs d'applications, et c'est aussi le problème : les script tags n'apparaissent pas dans l'éditeur de thème, vous ne pouvez donc ni les voir ni les désactiver à cet endroit. Des applications désinstallées ont parfois laissé les leurs derrière elles.

## Le calendrier

- **1er octobre 2026 :** les applications ne peuvent plus créer ni modifier de script tags, quelle que soit la version de l'API. Ceux qui existent continuent de fonctionner.
- **1er mars 2027 :** Shopify cesse de charger les script tags de la boutique en ligne.

Les script tags de la page d'état de la commande ont été supprimés plus tôt, selon un calendrier distinct. Si un élément a déjà disparu de cette page, c'est pour cette raison.

## Comment vérifier votre boutique

Shopify ne fournit pas aux marchands la liste des script tags, il faut donc aller voir soi-même :

1. **Listez les applications qui affichent quelque chose sur votre boutique.** Avis clients, chat, pop-ups, ventes additionnelles, fidélité, statistiques, pixels publicitaires. Les applications qui ne fonctionnent que dans l'interface d'administration ne sont pas concernées.
2. **Consultez les « App embeds » de l'éditeur de thème.** Si une application y figure et qu'elle est activée, elle utilise la nouvelle méthode et n'a rien à craindre.
3. **Regardez le code source de la page.** Ouvrez votre boutique, affichez le code source et recherchez `asyncLoad`. Sur la plupart des boutiques, les URL des script tags sont listées dans ce bloc. Chaque domaine renvoie généralement à une application.
4. **Interrogez le développeur de l'application.** « Votre application utilise-t-elle des script tags, et fonctionnera-t-elle après le 1er mars 2027 ? » est une question à laquelle une application maintenue répond en une ligne.

Depuis des années, Shopify pousse les applications vers les extensions d'application de thème (theme app extensions) : la plupart des applications bien maintenues sont donc déjà en règle. Le risque se situe du côté des anciennes applications, des applications sur mesure développées pour vous il y a longtemps, et des restes d'applications que vous n'utilisez plus.

## Ce qui remplace les script tags

- **Les blocs App embed**, qui font partie d'une extension d'application de thème, pour tout ce qui s'affiche sur la page. Vous les activez dans l'éditeur de thème, ce qui signifie que vous pouvez enfin les voir.
- **Les web pixels** pour les scripts qui ne servent qu'à collecter des données d'analyse ou de conversion. Ils s'exécutent dans un environnement isolé (sandbox) et ne demandent rien de votre part dès lors que l'application les prend en charge.

## Si une application n'est pas mise à jour

Remplacez-la avant mars, ou faites reconstruire sa fonctionnalité. Pour une application sur mesure, le travail consiste généralement à déplacer le script vers une extension d'application de thème, ou vers un web pixel s'il ne fait que du suivi. Profitez-en pour supprimer les script tags des applications que vous avez désinstallées : la mutation `scriptTagDelete` fonctionne toujours.

Si vous avez une application Shopify sur mesure à migrer, ou si vous ne savez pas ce qui se charge sur votre boutique, [envoyez-nous l'adresse de votre boutique](https://sawirstudio.com/fr/contact/) et nous vous dirons ce que nous trouvons. Vous en saurez plus sur notre travail autour de Shopify sur [cette page](https://sawirstudio.com/fr/shopify/).

**À lire aussi :** [API Shopify 2026-10 : ce qui risque de casser dans vos intégrations](https://sawirstudio.com/fr/posts/shopify-2026-10-api-what-breaks/), ou [Connecter Shopify à Odoo 19 Community, pas à pas](https://sawirstudio.com/fr/posts/connect-shopify-to-odoo-19-community/).
