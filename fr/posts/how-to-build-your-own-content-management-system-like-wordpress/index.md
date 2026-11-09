---
title: Créer votre propre CMS comme WordPress | SawirStudio
url: https://sawirstudio.com/fr/posts/how-to-build-your-own-content-management-system-like-wordpress/
description: Créer un système de gestion de contenu comme WordPress avec Laravel et une architecture logicielle moderne.
---

# Comment créer votre propre système de gestion de contenu (comme WordPress)

Construire un système de gestion de contenu de A à Z a été une expérience vraiment passionnante.

WordPress est devenu la référence incontournable, en particulier dans le domaine des systèmes de gestion de contenu. Comme l'indique son site web...

> WordPress propulse plus de 43 % du web, et ce chiffre augmente chaque jour. Des sites web simples aux blogs, en passant par les portails complexes et les sites d'entreprise, et même des applications : tout est construit avec WordPress.

Pourtant, WordPress a été conçu à une époque où créer une application n'était pas aussi simple qu'aujourd'hui.

De nos jours, le développement PHP s'est accéléré, comme le montrent les dernières [évolutions de PHP](https://www.php.net/), et les pratiques de développement logiciel ont elles aussi évolué.

Par exemple, le développement avec des structures MVC, inspirées de Ruby on Rails, et certains principes d'architecture logicielle ont été adoptés par de nombreux frameworks.

Or, ce n'est pas le cas de WordPress, qui continue d'utiliser ses propres structures.

Cet article a donc pour but de vous aider à comprendre ce qu'il est possible de faire lorsqu'un système de gestion de contenu, comme WordPress, repose sur une architecture logicielle moderne.

Pour cela, nous allons utiliser [Laravel](https://laravel.com), un framework web PHP qui met en œuvre plusieurs bonnes pratiques d'ingénierie logicielle.

Voyons rapidement quelles fonctionnalités notre système de gestion de contenu proposera. À partir de celles de WordPress, nous pouvons en dégager quelques-unes essentielles :

- Publier en toute simplicité Si vous avez déjà créé un document, vous savez déjà créer du contenu avec WordPress. Vous pouvez créer des articles et des pages, les mettre en forme facilement, y insérer des médias, et en un clic votre contenu est en ligne.
- Outils de publication WordPress facilite la gestion de votre contenu. Créez des brouillons, programmez une publication, consultez les révisions de vos articles. Rendez votre contenu public ou privé, et protégez articles et pages par un mot de passe.
- Gestion des utilisateurs Tout le monde n'a pas besoin du même accès à votre site. Les administrateurs gèrent le site, les éditeurs travaillent sur le contenu, les auteurs et les contributeurs rédigent ce contenu, et les abonnés disposent d'un profil qu'ils peuvent gérer. Vous pouvez ainsi accueillir différents contributeurs sur votre site, et laisser les autres simplement faire partie de votre communauté.
- Gestion des médias On dit qu'une image vaut mille mots : il est donc important de pouvoir importer rapidement et facilement des images et des médias dans WordPress. Glissez-déposez vos fichiers dans l'outil d'import pour les ajouter à votre site. Ajoutez un texte alternatif et des légendes, puis insérez des images et des galeries dans votre contenu. Nous avons même ajouté quelques outils de retouche d'image avec lesquels vous pourrez vous amuser.
- Commentaires intégrés Votre blog est votre maison, et les commentaires offrent à vos proches et à vos abonnés un espace pour réagir à votre contenu. Les outils de commentaires de WordPress vous donnent tout ce qu'il faut pour animer la discussion et la modérer.
- Optimisé pour les moteurs de recherche WordPress est optimisé pour les moteurs de recherche dès l'installation. Pour un contrôle SEO plus fin, de nombreuses extensions SEO s'en chargent à votre place.

Alors, mettons-nous au travail !

Installation

```
composer create-project laravel/laravel wp-laravel
```

Essayons d'accéder à notre application

```
php artisan serve
```

Nous aurons besoin de configurer une base de données ; SQLite sera notre allié dans cet article.

```
touch database/database.sqlite
```

Nous devons mettre à jour notre fichier .env

```
DB_CONNECTION=sqlite
DB_DATABASE=database/database.sqlite
DB_FOREIGN_KEYS=true
```

[référence](https://laravel.com/docs/9.x/database#sqlite-configuration)

Lançons la migration initiale de notre base de données

```
php artisan migrate
```

Nous pouvons maintenant voir les modifications de notre base de données ainsi

```
php artisan db:show
```

Toutes les migrations ont été appliquées.

Nous voulons maintenant disposer de quelque chose comme les articles de WordPress. Si vous avez déjà inspecté les tables de WordPress, vous savez qu'il existe une table appelée wp\_posts. Créons-la

```
php artisan make:model Post
```

Cela va créer un modèle appelé Post. Créons maintenant la migration pour appliquer les modifications à la base de données.

```
php artisan make:migration "create posts table"
```

Ouvrons maintenant la migration et modifions notre code comme suit :

```
$table->string('title');
$table->string('slug')->nullable();
$table->text('excerpt')->nullable();
$table->longText('content')->nullable();
$table->dateTime('published_at')->nullable();
$table->foreignIdFor(\App\Models\User::class)->constrained();
```

**À lire aussi :** [Comment concevoir les restrictions de vos formules d'abonnement](https://sawirstudio.com/fr/posts/how-do-you-design-your-subscription-plan-restrictions/), ou découvrez les [extensions WordPress gratuites que nous créons](https://sawirstudio.com/fr/projects/wordpress-plugins/).
