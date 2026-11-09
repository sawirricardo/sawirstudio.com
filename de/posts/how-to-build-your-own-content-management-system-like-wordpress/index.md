---
title: Eigenes CMS wie WordPress bauen | SawirStudio
url: https://sawirstudio.com/de/posts/how-to-build-your-own-content-management-system-like-wordpress/
description: Ein Content-Management-System wie WordPress mit Laravel und moderner Softwarearchitektur bauen.
---

# So bauen Sie Ihr eigenes Content-Management-System (wie WordPress)

Ein Content-Management-System von Grund auf zu entwickeln, war eine richtig spannende Erfahrung.

WordPress ist der De-facto-Standard, gerade bei Content-Management-Systemen. Von der eigenen Website:

> WordPress treibt mehr als 43 % des Webs an – und der Anteil wächst jeden Tag. Von einfachen Websites über Blogs bis hin zu komplexen Portalen, Unternehmenswebsites und sogar Anwendungen: Alles wird mit WordPress gebaut.

Allerdings entstand WordPress in einer Zeit, in der das Entwickeln von Anwendungen noch nicht so einfach war.

Heute hat sich die PHP-Entwicklung beschleunigt, wie die jüngsten [PHP-Updates](https://www.php.net/) zeigen, und auch die Praktiken der Softwareentwicklung haben sich weiterentwickelt.

So werden in vielen Software-Frameworks inzwischen MVC-Strukturen eingesetzt, die von Ruby on Rails inspiriert sind, sowie einige Prinzipien der Softwarearchitektur.

WordPress hat dies jedoch nicht übernommen und nutzt weiterhin seine eigenen Strukturen.

Dieser Beitrag soll Ihnen deshalb zeigen, was möglich ist, wenn ein Content-Management-System wie WordPress mit moderner Softwarearchitektur umgesetzt wird.

Dafür verwenden wir [Laravel](https://laravel.com), ein PHP-Web-Framework, das einige bewährte Praktiken der Softwareentwicklung umsetzt.

Besprechen wir kurz, welche Funktionen unser Content-Management-System haben soll. Anhand der Funktionen von WordPress lassen sich einige zentrale Features ableiten:

- Einfach veröffentlichen: Wenn Sie schon einmal ein Dokument erstellt haben, können Sie auch mit WordPress mühelos Inhalte anlegen. Sie erstellen Beiträge und Seiten, formatieren sie einfach, fügen Medien ein, und mit einem Klick ist Ihr Inhalt live im Web.
- Werkzeuge zum Veröffentlichen: WordPress macht die Verwaltung Ihrer Inhalte leicht. Legen Sie Entwürfe an, planen Sie Veröffentlichungen und sehen Sie sich frühere Versionen Ihrer Beiträge an. Machen Sie Ihre Inhalte öffentlich oder privat und schützen Sie Beiträge und Seiten mit einem Passwort.
- Benutzerverwaltung: Nicht jeder braucht denselben Zugriff auf Ihre Website. Administratoren verwalten die Website, Redakteure bearbeiten Inhalte, Autoren und Mitwirkende schreiben diese Inhalte, und Abonnenten haben ein Profil, das sie selbst pflegen können. So können Sie viele verschiedene Beitragende an Ihrer Website beteiligen und anderen einfach die Möglichkeit geben, Teil Ihrer Community zu sein.
- Medienverwaltung: Ein Bild sagt mehr als tausend Worte – deshalb ist es wichtig, Bilder und Medien schnell und einfach in WordPress hochladen zu können. Ziehen Sie Ihre Medien per Drag and drop in den Uploader, um sie Ihrer Website hinzuzufügen. Ergänzen Sie Alternativtexte und Bildunterschriften und fügen Sie Bilder und Galerien in Ihre Inhalte ein. Es gibt sogar ein paar Werkzeuge zur Bildbearbeitung, mit denen Sie experimentieren können.
- Integrierte Kommentare: Ihr Blog ist Ihr Zuhause, und Kommentare geben Ihren Freunden und Followern Raum, sich mit Ihren Inhalten auseinanderzusetzen. Die Kommentarfunktionen von WordPress bieten alles, was Sie brauchen, um ein Forum für Diskussionen zu schaffen und diese zu moderieren.
- Suchmaschinenoptimiert: WordPress ist von Haus aus für Suchmaschinen optimiert. Für eine feinere SEO-Steuerung gibt es zahlreiche SEO-Plugins, die das für Sie übernehmen.

Legen wir also los!

Installation

```
composer create-project laravel/laravel wp-laravel
```

Rufen wir unsere Anwendung einmal auf

```
php artisan serve
```

Wir brauchen eine Datenbank. In diesem Beitrag ist SQLite dafür eine gute Wahl.

```
touch database/database.sqlite
```

Wir müssen unsere .env anpassen

```
DB_CONNECTION=sqlite
DB_DATABASE=database/database.sqlite
DB_FOREIGN_KEYS=true
```

[Referenz](https://laravel.com/docs/9.x/database#sqlite-configuration)

Führen wir die erste Datenbankmigration aus

```
php artisan migrate
```

Jetzt sehen wir die Änderungen an unserer Datenbank so

```
php artisan db:show
```

Alle Migrationen wurden angewendet.

Nun möchten wir so etwas wie einen WordPress-Beitrag (Post) haben. Wenn Sie schon einmal die WordPress-Tabellen untersucht haben, wissen Sie, dass es eine Tabelle namens wp\_posts gibt. Legen wir sie an

```
php artisan make:model Post
```

Dadurch wird ein Model namens Post erstellt. Erstellen wir nun die Migration, um die Datenbankänderungen anzuwenden.

```
php artisan make:migration "create posts table"
```

Öffnen Sie jetzt die Migration und passen Sie den Code wie folgt an:

```
$table->string('title');
$table->string('slug')->nullable();
$table->text('excerpt')->nullable();
$table->longText('content')->nullable();
$table->dateTime('published_at')->nullable();
$table->foreignIdFor(\App\Models\User::class)->constrained();
```

**Weiterlesen:** [Wie gestalten Sie die Einschränkungen Ihrer Abo-Pläne?](https://sawirstudio.com/de/posts/how-do-you-design-your-subscription-plan-restrictions/) Oder sehen Sie sich die [kostenlosen WordPress-Plugins](https://sawirstudio.com/de/projects/wordpress-plugins/) an, die wir entwickeln.
