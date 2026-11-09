---
title: WordPress 7.2 : ce que les auteurs de plugins doivent tester | SawirStudio
url: https://sawirstudio.com/fr/posts/wordpress-7-2-what-plugin-authors-should-test/
description: WordPress 7.2 est prévu pour le 8 décembre 2026. Ce qui change pour les auteurs de plugins, des Application Passwords à l'API Secrets proposée, en passant par le routeur de l'Interactivity API, et comment tester dès maintenant.
---

# WordPress 7.2 arrive le 8 décembre : ce que nous testons dans nos plugins

WordPress 7.2 est prévu pour le 8 décembre 2026. La première version candidate (RC) est attendue le 17 novembre : c'est à ce moment-là que la liste des fonctionnalités est figée et que paraît le Field Guide, qui recense toutes les notes pour les développeurs. D'ici là, les versions bêta sont le moyen le moins coûteux de découvrir que votre plugin pose problème.

Nous maintenons [12 plugins gratuits sur WordPress.org](https://sawirstudio.com/fr/projects/wordpress-plugins/) : voici donc notre propre liste de contrôle. Certains points ci-dessous n'étaient encore que des propositions au moment de la rédaction ; c'est le Field Guide qui fait foi.

## Comment tester dès maintenant

- **WP-CLI** sur un site de test : `wp core update --version=7.2-beta4`, ou la bêta ou RC la plus récente.
- **Plugin WordPress Beta Tester**, réglé sur « Bleeding edge » et « Beta/RC Only ».
- **WordPress Playground**, pour tout essayer dans un navigateur sans rien installer.

Jamais sur un site en production. Activez `WP_DEBUG` et `WP_DEBUG_LOG` et surveillez les avis de dépréciation : c'est là qu'apparaissent la plupart des premiers avertissements.

## Ce que nous vérifions

### Application Passwords

WordPress 7.2 renforce les Application Passwords : meilleure détection des environnements locaux et HTTPS, e-mail envoyé lors de l'ajout d'un mot de passe, et gestion plus sûre du rôle par défaut. Si votre plugin, ou les outils de vos clients, s'authentifie ainsi auprès de l'API REST, testez ce parcours de bout en bout.

### L'API Secrets

Une méthode native pour stocker des identifiants de façon chiffrée, afin que les plugins cessent de conserver leurs clés d'API en clair dans les options. Elle a été proposée pour la version 7.2, avec l'écran d'administration repoussé à la 7.3. Si votre plugin stocke des clés d'API, surveillez le Field Guide : une fois l'API disponible, c'est là que ces clés devront être conservées. Aucun de nos plugins gratuits ne stocke aujourd'hui de clés tierces ; pour nous, cela concerne donc surtout les projets clients.

### Mode sudo

Prévu : demander à un utilisateur connecté de confirmer à nouveau son mot de passe avant les actions à hauts privilèges. Si votre plugin effectue de telles actions depuis ses propres écrans ou via AJAX, testez ce qui se passe lorsque le cœur de WordPress demande une nouvelle vérification en cours de route.

### Routeur de l'Interactivity API

Un nouveau routeur pour observer la navigation côté client remplace le couple obsolète `state.navigation.hasStarted` et `hasFinished`. Une directive `data-wp-html` permet aussi d'afficher du balisage de confiance.

### Icônes et notifications de l'administration

Les Dashicons sont remplacées par des icônes SVG dans la barre d'administration et le menu, et les notifications de l'administration sont mieux annoncées par les lecteurs d'écran. Si votre plugin personnalise l'icône de son menu ou cible les classes Dashicon dans le CSS de l'administration, vérifiez le rendu dans la bêta.

### DataViews et l'éditeur de site

Les plugins peuvent enregistrer leurs propres écrans dans l'éditeur de site et piloter côté serveur les champs et les dispositions de DataViews. Rien ne casse de ce côté-là, mais si votre plugin comporte un écran de liste, il est utile de savoir où les choses vont.

## Après les tests

Mettez à jour la mention « Tested up to » de votre readme une fois que vous avez réellement testé, pas avant. WordPress.org affiche un avertissement aux utilisateurs lorsqu'un plugin n'a pas été testé avec les dernières versions majeures, et cet avertissement fait perdre des installations. Nos plugins indiquent actuellement « Tested up to: 7.1 » ; ils passeront à 7.2 une fois les tests de la RC réussis.

Le nouveau thème par défaut, Ipsum, est également livré. Si votre plugin affiche quoi que ce soit sur le site public, vérifiez-le avec ce thème : c'est l'apparence que prendront les nouveaux sites.

Si l'un de vos plugins doit être mis à jour pour WordPress 7.2, ou si vous avez un plugin sur mesure que plus personne ne maintient, [nous pouvons y jeter un œil](https://sawirstudio.com/fr/contact/).

**À lire aussi :** [Rendre votre plugin WordPress utilisable par des agents IA grâce à l'Abilities API](https://sawirstudio.com/fr/posts/wordpress-abilities-api-for-plugins/), ou [WooCommerce ou Shopify pour une entreprise indonésienne ?](https://sawirstudio.com/fr/posts/woocommerce-or-shopify-indonesia/).
