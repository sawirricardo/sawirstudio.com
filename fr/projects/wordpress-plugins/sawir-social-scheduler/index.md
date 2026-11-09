---
title: Sawir Social Scheduler | SawirStudio
url: https://sawirstudio.com/fr/projects/wordpress-plugins/sawir-social-scheduler/
description: Rédigez une fois, programmez sur 27 réseaux depuis wp-admin. Modifications par canal, fils, médias, aperçus. Vos jetons restent sur votre serveur.
---

# Sawir Social Scheduler

![Bannière de Sawir Social Scheduler](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-banner.webp)

Sawir Social Scheduler vous permet de rédiger un message une seule fois dans wp-admin et de le programmer sur 27 réseaux sociaux, avec des modifications par canal, des fils, des médias et des aperçus. Les jetons de vos comptes restent sur votre serveur.

[Téléchargez-le gratuitement sur WordPress.org](https://wordpress.org/plugins/sawir-social-scheduler/)

## Ce qu'il fait

Sawir Social Scheduler transforme WordPress en outil de programmation pour les réseaux sociaux. Connectez vos canaux avec les jetons fournis par chaque plateforme, rédigez un message, choisissez les canaux, fixez une date et publiez. Chaque envoi s'exécute en arrière-plan et vous renvoie un lien ou le message d'erreur de la plateforme elle-même.

Aucun service hébergé ne s'interpose. Les identifiants sont chiffrés avec les clés (salts) de votre site et ne quittent jamais votre serveur.

### Canaux

Ajoutez autant de canaux que vous le souhaitez, avec plusieurs comptes par plateforme. Chaque canal dispose d'un bouton **Tester la connexion** qui récupère le nom et l'avatar du compte et confirme que le jeton fonctionne vraiment avant que vous ne comptiez dessus.

**Réseaux sociaux :** X (Twitter), Page Facebook, Instagram, Threads, Page LinkedIn, Bluesky, Mastodon, TikTok, YouTube, Pinterest, Tumblr, Dribbble, Google Business Profile

**Communautés :** Telegram, Discord, Slack, Reddit, Lemmy, VK, Kick, Twitch, Whop

**Blog et newsletter :** WordPress, Medium, Dev.to, Hashnode, Listmonk

Regroupez les canaux par client ou par marque grâce à des groupes imbriqués, puis sélectionnez un groupe entier en un clic.

### Messages

- **Un seul texte, plusieurs canaux.** Choisissez les canaux avec des cases à cocher. Les canaux que vous ne modifiez pas reçoivent le message tel qu'il est rédigé.
- **Personnalisation par canal.** Cochez « Personnaliser » sur un canal pour lui donner ses propres parties : texte, médias et délais, préremplis à partir du message pour que vous modifiiez une copie. Les autres canaux gardent le message tel qu'il est rédigé.
- **Fils.** Ajoutez des parties pour publier un fil sur X, Threads, Bluesky, Mastodon et Telegram, chacune avec ses propres médias. Les plateformes sans fils reçoivent les parties réunies en un seul message.
- **Médias.** Joignez images et vidéos une seule fois. Chaque plateforme reçoit ce qu'elle prend en charge : albums, carrousels, Reels, envois de vidéos reprenables.
- **Programmation.** Utilisez la date « Publier le » habituelle de WordPress, ou donnez à chaque canal son propre horaire.
- **Horaires de publication.** Définissez pour un canal ses horaires habituels dans la journée. « Prochain créneau libre » remplit le prochain créneau inutilisé, le calendrier affiche les prochains créneaux vides sur lesquels vous pouvez cliquer pour les remplir, et l'API accepte `"date": "next"`.
- **Aperçu.** Un aperçu en direct par canal, fidèle à la vraie plateforme, avec les limites de caractères et les règles sur les médias appliquées, pour voir exactement ce qui sera publié.
- **Commentaires et délais.** Sur les plateformes sans fils, les parties supplémentaires sont publiées en commentaires sous le message. Chaque partie peut attendre un nombre de minutes défini après la précédente.
- **Partager un article de blog.** Un bloc dans l'éditeur d'articles habituel publie sur vos canaux au moment de la publication de l'article, avec son propre texte et ses propres médias.
- **Dupliquer.** Copiez un message avec ses canaux, ses parties et ses personnalisations sous forme de nouveau brouillon. Conservez des brouillons comme modèles.
- **Mentions.** Saisissez `@handle` et le réseau le résout sur X, Instagram, Threads, Mastodon, Telegram, Reddit et Bluesky, où les liens et les identifiants deviennent de vraies facettes. Les pages d'entreprise LinkedIn s'écrivent `@[Name](org:12345)` avec l'ID tiré de l'URL de la page ; les pages Facebook s'écrivent `@[12345]`.
- **Répétition.** Renvoyez des messages intemporels tous les N jours ou toutes les N semaines, jusqu'à une date donnée ou un certain nombre de fois. Chaque exécution est un message à part entière, avec ses propres envois et statistiques. Placez `{n}` dans le texte pour numéroter les exécutions, car X, Threads et Bluesky rejettent les textes identiques.

### Calendrier

Chaque message se décline en un envoi par canal, et le calendrier affiche chacun sous forme de carte colorée selon son état : programmé, envoyé, échoué, en attente, brouillon, en cours. Faites glisser une carte pour déplacer l'horaire de ce canal. Cliquez sur une carte pour voir l'erreur, un bouton Réessayer, le lien en ligne et les chiffres d'engagement. Cliquez sur un créneau vide pour créer un message à cette heure. Vues par mois, semaine, jour et liste, filtrables par groupe et par statut.

### Statistiques

Les mentions J'aime, commentaires, partages, vues et impressions sont collectés 1 heure, 1 jour et 1 semaine après chaque envoi, ainsi qu'à la demande, sur 16 plateformes. Le nombre d'abonnés est enregistré chaque jour par canal, avec une courbe sur 30 jours. Les chiffres apparaissent sur les cartes du calendrier, dans le bloc Statistiques du message et dans la liste des canaux. La page Statistiques fait la synthèse des 7, 30 ou 90 derniers jours par rapport à la période précédente : messages envoyés, engagement, portée et abonnés avec leur évolution, un graphique des abonnés, l'engagement par canal et vos meilleurs messages.

### Validation par le client

Cochez « Suspendre les envois jusqu'à validation », copiez le lien d'aperçu et envoyez-le à votre client. Il voit le message tel qu'il apparaîtra sur chaque canal, sans connexion, et clique sur Approuver ou sur Demander des modifications avec une note. L'approbation libère les envois. Chaque décision est conservée dans l'historique.

Les envois de grosses vidéos vers X, YouTube et TikTok se font par petits segments qui reprennent d'une tâche en arrière-plan à l'autre. Ils résistent ainsi aux limites de temps des hébergeurs et ne chargent jamais un fichier entier en mémoire.

### API REST

Tout ce que fait l'administration est disponible sur `/wp-json/sawir-social/v1` pour vos automatisations, applications mobiles et autres sites. L'authentification se fait avec un mot de passe d'application ; les permissions suivent les capacités du plugin et les identifiants des canaux ne sont jamais renvoyés.

- `GET /platforms`, `GET /channels`, `GET /channels/{id}`, `POST /channels/{id}/test`
- `GET|POST /posts`, `GET|PATCH|DELETE /posts/{id}`, `POST /posts/{id}/duplicate`, `GET /posts/{id}/analytics`
- `GET /sends`, `GET /sends/{id}`, `POST /sends/{id}/retry`, `POST /sends/{id}/stats`

Créez un message avec `{"title","content","status":"future","date":"2026-12-25 09:00","media":[12],"parts":[{"text","media","delay"}],"channels":[{"id":38,"parts":[{"text","media","delay"}],"date":"…"}],"groups":["Client A"],"hold":false}`. Les `parts` d'un canal remplacent celles du message pour ce canal ; `text` est un raccourci pour une partie unique.

### Agents IA

Les mêmes opérations sont enregistrées dans l'API Abilities de WordPress sous les noms `sawir-social/list-channels`, `create-post`, `update-post`, `list-posts`, `list-sends`, `retry-send`, `test-channel`, `get-analytics` et `list-platforms`, avec des schémas d'entrée et de sortie typés. Tout agent connecté via l'adaptateur WordPress MCP les voit comme des outils, et ils sont également appelables sur `/wp-json/wp-abilities/v1/abilities/{name}/run`. L'API Abilities fait partie de WordPress 6.9 et des versions ultérieures ; pour les versions plus anciennes, le plugin l'intègre, de sorte que les capacités sont disponibles partout où le plugin fonctionne.

### Prérequis

Les envois s'exécutent en arrière-plan via Action Scheduler lorsque votre site l'a (WooCommerce et de nombreux autres plugins le fournissent), et via WP-Cron sinon. Avec le cron intégré de WordPress, ils n'avancent que lorsque le site reçoit du trafic. Pour des envois rapides, ajoutez un cron système : l'écran Calendrier affiche la commande exacte pour votre site.

Les médias doivent être accessibles depuis Internet pour les plateformes qui les récupèrent par URL (Instagram, Threads, Facebook, Pinterest, Telegram). Un site de développement local ne peut pas leur transmettre d'images.

### Crédits

Intègre FullCalendar (MIT), la bibliothèque WordPress Abilities API (GPL) et des logos de marques issus de Simple Icons (CC0). Utilise Action Scheduler lorsqu'un autre plugin le fournit.

## Captures d'écran

![Capture d'écran 1 de Sawir Social Scheduler](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-screenshot-1.webp)

![Capture d'écran 2 de Sawir Social Scheduler](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-screenshot-2.webp)

![Capture d'écran 3 de Sawir Social Scheduler](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-screenshot-3.webp)

![Capture d'écran 4 de Sawir Social Scheduler](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-screenshot-4.webp)

![Capture d'écran 5 de Sawir Social Scheduler](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-screenshot-5.webp)

![Capture d'écran 6 de Sawir Social Scheduler](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-screenshot-6.webp)

## Questions

### Où trouver les jetons ?

Dans les paramètres développeur de chaque plateforme. Le formulaire du canal indique l'identifiant exact requis, par exemple un jeton d'accès de Page Facebook, un jeton de bot Telegram ou un mot de passe d'application Bluesky. Les jetons de longue durée sont préférables ; les jetons Meta et LinkedIn expirent après environ 60 jours et doivent être collés à nouveau.

### Mon test Telegram affiche « chat not found »

Les canaux privés n'ont pas de @username. Laissez l'ID du chat vide et cliquez sur Tester la connexion : le plugin liste tous les chats auxquels le bot a été ajouté, avec l'ID numérique à coller.

### Pourquoi un envoi reste-t-il bloqué sur « En cours… » ?

La file d'attente en arrière-plan ne s'est pas encore exécutée. Visitez le site, ou configurez le cron système indiqué sur l'écran Calendrier.

### Puis-je publier une vidéo sur Instagram depuis un site local ?

Non. Instagram récupère les médias par URL et a besoin d'une adresse publique. Tout le reste du plugin fonctionne en local.

## À associer avec

- [Sawir Insights](https://sawirstudio.com/fr/projects/wordpress-plugins/sawir-insights/) : montre comment les visiteurs utilisent votre site, avec des cartes de chaleur, des enregistrements de sessions, des sondages, des tunnels de conversion et des statistiques de formulaires, et conserve le tout sur votre propre serveur.
- [Sawir Live Translate](https://sawirstudio.com/fr/projects/wordpress-plugins/sawir-live-translate/) : traduit vos pages directement là où elles se trouvent : choisissez une langue, traduisez automatiquement ce qui est affiché, enregistrez, puis proposez à vos visiteurs un sélecteur de langue.

Découvrez tous nos [plugins WordPress gratuits](https://sawirstudio.com/fr/projects/wordpress-plugins/).

- **Prix** Gratuit
- **Version** 0.1.0
- **Prérequis** WordPress 5.8, PHP 7.4
- **Licence** GPL-2.0-or-later
- **WordPress.org** <https://wordpress.org/plugins/sawir-social-scheduler/>

Assez parlé de nous.

## Comment pouvons-nous vous aider ?

[Démarrer un projet](https://sawirstudio.com/fr/contact/)
