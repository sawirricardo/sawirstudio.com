---
title: Creare un CMS come WordPress | SawirStudio
url: https://sawirstudio.com/it/posts/how-to-build-your-own-content-management-system-like-wordpress/
description: Come creare un content management system come WordPress con Laravel e un'architettura software moderna.
---

# Come creare il vostro content management system (come WordPress)

Costruire un content management system da zero è stata un'esperienza davvero fantastica.

WordPress è lo standard di fatto, soprattutto nel mondo dei content management system. Dal suo sito...

> WordPress alimenta più del 43% del web, e la cifra cresce ogni giorno. Tutto, dai semplici siti web ai blog, dai portali complessi ai siti aziendali, fino alle vere e proprie applicazioni, viene realizzato con WordPress.

Tuttavia, WordPress è nato in un'epoca in cui creare un'applicazione non era così semplice.

Oggi lo sviluppo in PHP ha preso velocità, come dimostrano i recenti [aggiornamenti di PHP](https://www.php.net/), e anche le pratiche di sviluppo software si sono evolute.

Ad esempio, lo sviluppo con strutture MVC, ispirate a Ruby on Rails, e alcuni principi di architettura software sono stati adottati da moltissimi framework.

Questo però non è avvenuto per WordPress, che continua a usare le proprie strutture.

Questo articolo vuole quindi aiutarvi a capire cosa si può fare se un content management system come WordPress viene realizzato con un'architettura software moderna.

Per farlo useremo [Laravel](https://laravel.com), un framework web PHP che applica alcune buone pratiche di ingegneria del software.

Vediamo brevemente quali funzionalità avrà il nostro content management system. Partendo da quelle di WordPress, possiamo individuare alcune funzionalità chiave:

- Pubblicare con facilità. Se avete mai creato un documento, siete già bravissimi a creare contenuti con WordPress. Potete creare articoli e pagine, formattarli facilmente, inserire contenuti multimediali e, con un clic, i vostri contenuti sono online.
- Strumenti di pubblicazione. WordPress rende semplice la gestione dei contenuti. Create bozze, programmate la pubblicazione e consultate le revisioni dei vostri articoli. Rendete i contenuti pubblici o privati e proteggete articoli e pagine con una password.
- Gestione degli utenti. Non tutti hanno bisogno degli stessi accessi al vostro sito. Gli amministratori gestiscono il sito, i redattori lavorano sui contenuti, autori e collaboratori scrivono quei contenuti, e gli iscritti hanno un profilo che possono gestire. In questo modo potete avere tanti collaboratori sul vostro sito e lasciare che altri facciano semplicemente parte della vostra community.
- Gestione dei media. Si dice che un'immagine valga più di mille parole, ed è per questo che è importante poter caricare immagini e contenuti multimediali su WordPress in modo rapido e semplice. Trascinate i vostri file nell'uploader per aggiungerli al sito. Aggiungete testo alternativo e didascalie, e inserite immagini e gallerie nei vostri contenuti. Abbiamo persino aggiunto qualche strumento di modifica delle immagini con cui divertirvi.
- Commenti integrati. Il vostro blog è casa vostra, e i commenti offrono ai vostri amici e follower uno spazio per interagire con i vostri contenuti. Gli strumenti per i commenti di WordPress vi danno tutto ciò che serve per creare uno spazio di discussione e moderarlo.
- Ottimizzato per i motori di ricerca. WordPress è ottimizzato per i motori di ricerca fin dall'installazione. Per un controllo SEO più preciso, esistono tanti plugin SEO che se ne occupano al posto vostro.

Bene, mettiamoci al lavoro!

Installazione

```
composer create-project laravel/laravel wp-laravel
```

Proviamo ad accedere alla nostra applicazione

```
php artisan serve
```

Dovremo configurare un database: in questo articolo SQLite sarà il nostro alleato.

```
touch database/database.sqlite
```

Dovremo aggiornare il nostro file .env

```
DB_CONNECTION=sqlite
DB_DATABASE=database/database.sqlite
DB_FOREIGN_KEYS=true
```

[riferimento](https://laravel.com/docs/9.x/database#sqlite-configuration)

Eseguiamo la migrazione delle modifiche iniziali al database

```
php artisan migrate
```

Ora possiamo vedere le modifiche al database in questo modo

```
php artisan db:show
```

Tutte le migrazioni sono state applicate.

Ora vogliamo qualcosa come i Post di WordPress. Se avete mai ispezionato le tabelle di WordPress, sapete che esiste una tabella chiamata wp\_posts. Creiamola

```
php artisan make:model Post
```

Questo creerà un modello chiamato Post. Creiamo ora la migrazione per applicare le modifiche al database.

```
php artisan make:migration "create posts table"
```

Ora apriamo la migrazione e aggiorniamo il codice come segue:

```
$table->string('title');
$table->string('slug')->nullable();
$table->text('excerpt')->nullable();
$table->longText('content')->nullable();
$table->dateTime('published_at')->nullable();
$table->foreignIdFor(\App\Models\User::class)->constrained();
```

**Continuate a leggere:** [Come progettare le restrizioni dei piani di abbonamento](https://sawirstudio.com/it/posts/how-do-you-design-your-subscription-plan-restrictions/), oppure scoprite i [plugin WordPress gratuiti che realizziamo](https://sawirstudio.com/it/projects/wordpress-plugins/).
