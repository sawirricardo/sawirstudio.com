---
title: Rendre nos sites accessibles aux agents IA | SawirStudio
url: https://sawirstudio.com/fr/posts/making-our-sites-agent-ready/
description: Comment nous avons rendu nos sites produits lisibles par les agents IA : copies Markdown de chaque page, en-têtes Link, llms.txt, Content Signals dans robots.txt, catalogue d’API et Agent Skills.
---

# Ce qu’il a fallu pour rendre nos sites accessibles aux agents IA

De plus en plus de visites sur un site web ne viennent pas de personnes, mais d’agents IA qui lisent des pages pour répondre à la question de quelqu’un, ou qui essaient d’utiliser un produit pour le compte d’un utilisateur. Cette semaine, nous avons passé en revue ce site et nos sites produits, [SawirPost](https://sawirstudio.com/fr/projects/sawirpost/), [Sawirly](https://sawirstudio.com/fr/projects/sawirly/), [Rindraw](https://sawirstudio.com/fr/projects/rindraw/), [Math Daily](https://sawirstudio.com/fr/projects/math-daily/) et d’autres, pour qu’ils soient plus faciles à lire et à utiliser pour les agents.

Nous avons utilisé [isitagentready.com](https://isitagentready.com) comme liste de contrôle. L’outil note un site par niveaux ; voici ce qu’a demandé chaque étape.

## Une copie Markdown de chaque page

Un agent qui lit du HTML doit traverser menus, scripts et pieds de page pour trouver quelques paragraphes. Chaque page de ce site a donc désormais sa copie Markdown, à côté, sous `index.md`, avec uniquement le contenu principal de la page. Un petit script les génère à partir du HTML, si bien qu’elles ne divergent jamais.

Les agents y accèdent de deux façons : en ajoutant `index.md` à une URL, ou en demandant `Accept: text/markdown`. Un minuscule Cloudflare Worker placé devant les fichiers statiques vérifie cet en-tête et renvoie la copie Markdown, avec une estimation du nombre de tokens dans un en-tête `x-markdown-tokens`. Les navigateurs, eux, reçoivent toujours la page habituelle.

```
curl -H "Accept: text/markdown" https://sawirstudio.com/shopify/
```

## En-têtes Link

Chaque page HTML indique désormais où se trouve sa copie Markdown, dans un en-tête HTTP `Link` avec `rel="alternate"`. La page d’accueil renvoie aussi vers `llms.txt` et le plan du site. Un agent peut tout trouver à partir d’une seule requête, sans deviner les URL.

## llms.txt

`llms.txt` est un index du site en texte brut destiné aux modèles de langage : un résumé en une ligne, puis la liste de toutes les pages avec leur description, qui renvoie vers les copies Markdown. Le nôtre est généré à partir du plan du site par le même script.

## Content Signals

`robots.txt` indique qui peut explorer le site, pas ce que les robots ont le droit de faire de ce qu’ils lisent. [Content Signals](https://contentsignals.org/) comble ce manque. Chez nous, il tient en une ligne et autorise les trois usages, car nous voulons être trouvés et cités :

```
Content-Signal: ai-train=yes, search=yes, ai-input=yes
```

Ces trois étapes ont suffi à amener tous nos sites au niveau 3.

## Un catalogue d’API et des Agent Skills

SawirPost et Sawirly disposent d’API, nous sommes donc allés plus loin. Chacun publie un catalogue d’API à l’adresse `/.well-known/api-catalog` (RFC 9727) : un petit document JSON qui renvoie vers la description OpenAPI et la documentation. Un agent qui arrive sur le site peut trouver l’API sans lire une seule page.

Ils publient aussi des Agent Skills, de courtes instructions qui expliquent aux agents comment réaliser les tâches courantes avec le produit. Cela a amené les deux sites au niveau 4.

## Là où nous nous sommes arrêtés

Le niveau 5 exige que les agents puissent s’inscrire et s’authentifier seuls, via un flux d’enregistrement d’agent OAuth. Une page qui explique comment s’inscrire avec une adresse e-mail et un mot de passe ne compte pas. C’est une vraie fonctionnalité à développer, avec de vraies questions de sécurité : nous y réfléchissons donc plutôt que de la précipiter.

## Cela en valait-il la peine ?

Pour un site statique, l’essentiel a pris un après-midi : un script, un petit Worker et une ligne dans `robots.txt`. Rien n’a changé pour les personnes qui visitent le site. Nous pensons que de plus en plus de questions trouveront leur réponse auprès d’agents qui lisent des sites comme le nôtre, et qu’une page Markdown propre est le meilleur moyen d’être cité correctement.

Nous mettons en place la même chose pour nos clients dans le cadre de [Launch & growth](https://sawirstudio.com/fr/launch-and-growth/).

**À lire aussi :** [Rendez votre plugin WordPress utilisable par les agents IA grâce à l’Abilities API](https://sawirstudio.com/fr/posts/wordpress-abilities-api-for-plugins/), ou [Les icônes d’application lançables de notre page d’accueil](https://sawirstudio.com/fr/posts/throwable-app-icons-matter-js/).
