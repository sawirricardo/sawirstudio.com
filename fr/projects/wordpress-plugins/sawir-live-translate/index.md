---
title: Sawir Live Translate | SawirStudio
url: https://sawirstudio.com/fr/projects/wordpress-plugins/sawir-live-translate/
description: Sawir Live Translate traduit vos pages là où elles se trouvent : choisissez une langue, traduisez automatiquement ce qui est visible, enregistrez, et offrez à vos visiteurs un sélecteur de langue.
---

# Sawir Live Translate

![Bannière de Sawir Live Translate](https://sawirstudio.com/assets/images/plugins/sawir-live-translate-banner.webp)

Sawir Live Translate traduit vos pages là où elles se trouvent : choisissez une langue, traduisez automatiquement ce qui est visible, enregistrez, et offrez à vos visiteurs un sélecteur de langue.

[Téléchargez-le gratuitement sur WordPress.org](https://wordpress.org/plugins/sawir-live-translate/)

## Ce qu'il fait

Ouvrez n'importe quelle page de votre site, cliquez sur « Translate this page » dans la barre d'administration, choisissez une langue cible, puis traduisez automatiquement tous les textes visibles (via un service de traduction automatique gratuit) ou modifiez les traductions à la main en voyant la page se mettre à jour en direct.

Les traductions enregistrées sont intégrées côté serveur dans la réponse HTML (idéal pour le référencement : pas de flash de texte non traduit, et des alternates hreflang dans le head). Les visiteurs changent de langue grâce à un sélecteur à base de liens (en bas à droite).

Les URL utilisent un préfixe de langue dans le chemin, par exemple /es/about pour l'espagnol, et /about pour la langue source (qui reste canonique, sans préfixe). La langue est déterminée dans cet ordre : chemin /es/ → ?lang=es → cookie enregistré → Accept-Language du navigateur → langue source. La racine et les URL sans préfixe détectent automatiquement la langue du navigateur du visiteur sans redirection (un en-tête Vary est envoyé pour que les caches partagés se comportent correctement).

1. Réglages → Live Translate : définissez votre langue source et cochez les langues à proposer. 2. Ouvrez une page du site côté visiteur et cliquez sur « 🌐 Translate this page » dans la barre d'administration. 3. Choisissez une langue, cliquez sur « Auto-translate empty » (ou modifiez les textes), puis sur « Save ». 4. Les visiteurs utilisent le sélecteur (en bas à droite) ; chaque langue a sa propre URL /xx/.

Remarque : les préfixes de langue dans le chemin supposent une installation à la racine du domaine et l'absence de page de premier niveau dont le slug serait un code de langue à deux lettres.

## Questions

### Comparaison avec les alternatives

Les solutions multilingues comme WPML, Polylang et TranslatePress traitent le contenu traduit comme un élément à part entière de la structure du site, avec des flux de traduction par langue, des services de traduction et une intégration poussée aux constructeurs de pages et à WooCommerce. Sawir Live Translate est volontairement plus léger : ouvrez une page, traduisez ce qui est visible (automatiquement ou à la main) et servez-la depuis une URL préfixée par la langue, avec des balises hreflang. Si vous construisez un site entièrement multilingue avec des menus, des produits et des flux éditoriaux traduits, utilisez l'une de ces solutions. Si vous voulez que vos pages existantes soient lisibles dans une autre langue dès cet après-midi, c'est exactement ce que fait ce plugin.

### Je change de langue, mais rien ne change

Dans la quasi-totalité des cas, cela signifie qu'aucune traduction n'est encore enregistrée pour cette langue. Vérifiez Réglages → Live Translate → Saved translations : si la langue n'apparaît pas (ou affiche 0 texte), rouvrez l'éditeur de page, lancez la traduction automatique, puis cliquez sur Save. L'aperçu en direct dans l'éditeur n'est pas enregistré tant que vous ne le faites pas.

### Certains textes restent non traduits

L'éditeur intégré à la page ne voit que les textes de la page que vous avez ouverte. Les textes qui se trouvent sur d'autres pages nécessitent que ces pages soient traduites aussi, ou que vous lanciez l'analyse de tout le site. Les textes injectés par des carrousels ou des scripts après le chargement de la page peuvent demander une saisie manuelle via « Edit all site strings ».

### J'ai modifié mon contenu et sa traduction a disparu

Les traductions sont associées au texte d'origine exact. Si vous réécrivez une phrase, traduisez la nouvelle version une fois, puis enregistrez.

## À associer avec

- [Sawir Insights](https://sawirstudio.com/fr/projects/wordpress-plugins/sawir-insights/) : montre comment les visiteurs utilisent votre site, avec des cartes de chaleur, des enregistrements de sessions, des sondages, des tunnels de conversion et des statistiques de formulaires, et conserve le tout sur votre propre serveur.
- [Sawir Media Folders](https://sawirstudio.com/fr/projects/wordpress-plugins/sawir-media-folders/) : ajoute des dossiers à la médiathèque WordPress, autant que vous en avez besoin, avec glisser-déposer pour tout ranger.

Découvrez tous nos [plugins WordPress gratuits](https://sawirstudio.com/fr/projects/wordpress-plugins/).

- **Prix** Gratuit
- **Version** 1.1.0
- **Prérequis** WordPress 5.8, PHP 7.4
- **Licence** GPL-2.0-or-later
- **WordPress.org** <https://wordpress.org/plugins/sawir-live-translate/>

Assez parlé de nous.

## Comment pouvons-nous vous aider ?

[Démarrer un projet](https://sawirstudio.com/fr/contact/)
