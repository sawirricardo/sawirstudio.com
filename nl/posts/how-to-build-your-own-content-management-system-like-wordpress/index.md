---
title: Bouw je eigen CMS zoals WordPress | SawirStudio
url: https://sawirstudio.com/nl/posts/how-to-build-your-own-content-management-system-like-wordpress/
description: Een contentmanagementsysteem zoals WordPress bouwen met Laravel en moderne softwarearchitectuur.
---

# Zo bouw je je eigen contentmanagementsysteem (zoals WordPress)

Een contentmanagementsysteem helemaal from scratch bouwen was een geweldige ervaring.

WordPress is de de-factostandaard geworden, zeker in de wereld van contentmanagementsystemen. Op de eigen website staat:

> WordPress draait op meer dan 43% van het web, en dat percentage groeit elke dag. Van simpele websites tot blogs, van complexe portalen tot bedrijfswebsites, en zelfs apps: alles wordt met WordPress gebouwd.

Toch zie je dat WordPress is ontwikkeld in een tijd waarin het maken van applicaties nog niet zo eenvoudig was.

Tegenwoordig is PHP-ontwikkeling flink versneld, zoals te zien is aan de recente [PHP-updates](https://www.php.net/), en ook de manier waarop we software bouwen is veranderd.

Denk aan ontwikkelen met MVC-structuren, geïnspireerd door Ruby on Rails, en aan architectuurprincipes die in veel softwareframeworks zijn toegepast.

Bij WordPress is dat niet gebeurd: het gebruikt nog steeds zijn eigen structuur.

Dit artikel laat daarom zien wat er mogelijk is als je een contentmanagementsysteem zoals WordPress bouwt met moderne softwarearchitectuur.

Daarvoor gebruiken we [Laravel](https://laravel.com), een PHP-webframework dat een aantal bewezen principes uit de softwareontwikkeling toepast.

Laten we eerst kort bekijken welke functies ons contentmanagementsysteem krijgt. Op basis van de functies van WordPress komen we uit op een aantal kernfuncties:

- Eenvoudig publiceren Als je ooit een document hebt gemaakt, kun je ook content maken met WordPress. Je maakt berichten en pagina's aan, past de opmaak eenvoudig aan, voegt media toe en met één klik staat je content live op het web.
- Publicatietools WordPress maakt het makkelijk om je content te beheren. Maak concepten, plan publicaties in en bekijk eerdere versies van je berichten. Maak je content openbaar of privé, en beveilig berichten en pagina's met een wachtwoord.
- Gebruikersbeheer Niet iedereen heeft dezelfde toegang tot je website nodig. Beheerders beheren de site, redacteuren werken met content, auteurs en medewerkers schrijven die content, en abonnees hebben een profiel dat ze zelf kunnen beheren. Zo kun je veel verschillende mensen aan je website laten bijdragen en anderen gewoon onderdeel van je community laten zijn.
- Mediabeheer Een plaatje zegt meer dan duizend woorden, en daarom is het belangrijk dat je snel en eenvoudig afbeeldingen en media in WordPress kunt uploaden. Sleep je media naar de uploader om ze aan je website toe te voegen. Voeg alt-tekst en bijschriften toe en plaats afbeeldingen en galerijen in je content. We hebben zelfs een paar tools voor beeldbewerking toegevoegd waar je mee kunt spelen.
- Ingebouwde reacties Je blog is je thuis, en reacties geven vrienden en volgers een plek om met je content in gesprek te gaan. De reactietools van WordPress bieden alles wat je nodig hebt om een plek voor discussie te zijn en die discussie te modereren.
- Geoptimaliseerd voor zoekmachines WordPress is meteen na installatie geoptimaliseerd voor zoekmachines. Voor meer gedetailleerde SEO-controle zijn er genoeg SEO-plugins die dat voor je regelen.

Laten we beginnen!

Installatie

```
composer create-project laravel/laravel wp-laravel
```

Laten we proberen onze applicatie te openen

```
php artisan serve
```

We hebben een database nodig. In dit artikel is SQLite onze vriend.

```
touch database/database.sqlite
```

We moeten ons .env-bestand aanpassen

```
DB_CONNECTION=sqlite
DB_DATABASE=database/database.sqlite
DB_FOREIGN_KEYS=true
```

[bron](https://laravel.com/docs/9.x/database#sqlite-configuration)

Laten we onze eerste databasewijziging migreren

```
php artisan migrate
```

Nu kunnen we onze databasewijzigingen zo bekijken

```
php artisan db:show
```

Alle migraties zijn toegepast.

Nu willen we iets als een WordPress-bericht. Als je ooit de tabellen van WordPress hebt bekeken, weet je dat er een tabel wp\_posts bestaat. Die gaan we aanmaken

```
php artisan make:model Post
```

Dit maakt een model met de naam Post aan. Laten we de migratie maken om de databasewijzigingen toe te passen.

```
php artisan make:migration "create posts table"
```

Open nu de migratie en pas de code aan zodat die er zo uitziet:

```
$table->string('title');
$table->string('slug')->nullable();
$table->text('excerpt')->nullable();
$table->longText('content')->nullable();
$table->dateTime('published_at')->nullable();
$table->foreignIdFor(\App\Models\User::class)->constrained();
```

**Lees verder:** [Hoe ontwerp je de beperkingen van abonnementen](https://sawirstudio.com/nl/posts/how-do-you-design-your-subscription-plan-restrictions/), of bekijk de [gratis WordPress-plugins die wij maken](https://sawirstudio.com/nl/projects/wordpress-plugins/).
