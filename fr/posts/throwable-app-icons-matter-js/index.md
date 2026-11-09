---
title: Icônes d'application lançables avec Matter.js | SawirStudio
url: https://sawirstudio.com/fr/posts/throwable-app-icons-matter-js/
description: Comment nous avons créé la boîte d'icônes d'application que vous pouvez lancer sur notre page d'accueil : Matter.js pour la physique, des éléments DOM plutôt qu'un canvas, une ombre qui reste en place pendant que les icônes tournent, et des solutions de repli.
---

# Les icônes d'application lançables de notre page d'accueil

Notre [page d'accueil](https://sawirstudio.com/fr/) contient une boîte remplie des icônes de nos applications. Elles tombent à votre arrivée, et vous pouvez les attraper et les lancer. Un simple toucher, sans glisser, ouvre la page de l'application correspondante. Cela nous a pris un après-midi, et certains détails se sont révélés plus intéressants que prévu.

## Pourquoi une boîte

Notre première idée était de laisser les icônes voler dans tout le haut de la page. Mais c'est là que se trouvent notre titre et nos boutons, et voir des icônes atterrir sur « Start a project » est un mauvais compromis pour un peu de plaisir. Une boîte à côté de l'introduction, avec une étiquette sur son bord, garde le jeu bien circonscrit. Elle comble aussi un vide déjà présent : l'espace libre à droite de l'introduction et des boutons sur un grand écran. Sur mobile, elle se place sous les boutons.

## Une vraie physique, grâce à Matter.js

Des cercles qui s'entrechoquent, c'est facile à coder soi-même. Mais les icônes d'application sont des carrés aux coins arrondis, et un carré qui roule comme une pièce de monnaie paraît tout de suite faux. Pour des culbutes et des empilements crédibles, il faut un vrai moteur physique : nous avons donc utilisé [Matter.js](https://brm.io/matter-js/). Chaque icône est un rectangle aux coins arrondis (`chamfer`) correspondant au rayon de 22 % de l'icône, placé entre quatre murs statiques.

Le moteur pèse 83 Ko : il ne se charge donc que sur la page d'accueil, une fois la page affichée, et seulement si le visiteur n'a pas demandé de réduire les animations.

## Des éléments DOM, pas un canvas

Matter.js peut dessiner dans un canvas, mais nous ne l'utilisons que pour les calculs. Les icônes sont de simples liens contenant des images, positionnés à chaque image avec une transformation CSS :

```
st.setProperty("--x", (b.position.x - size / 2) + "px");
st.setProperty("--y", (b.position.y - size / 2) + "px");
st.setProperty("--a", b.angle + "rad");
```

Ainsi, un toucher active un vrai lien, les images restent nettes à tous les niveaux de zoom, et les icônes gardent la même bordure et le même style que partout ailleurs sur le site.

## L'ombre qui reste en place

Notre style applique une ombre franche en bas à droite de chaque élément. Sur un élément qui pivote, l'ombre pivote avec lui, et celle d'une icône qui tourne se met alors à balayer l'espace comme l'aiguille d'une horloge. La solution consiste à faire pivoter le décalage de l'ombre en sens inverse à chaque image :

```
// keep the shadow at (3px, 3px) in screen space while the icon turns by a
st.setProperty("--shadow-x", 3 * (Math.cos(a) + Math.sin(a)) + "px");
st.setProperty("--shadow-y", 3 * (Math.cos(a) - Math.sin(a)) + "px");
```

## Attraper et lancer

Quand on appuie, nous attachons un ressort (une contrainte Matter.js) entre le pointeur et l'endroit exact de l'icône que vous avez saisi, et non son centre, de sorte qu'elle se balance naturellement quand vous la tenez par un coin. Quand vous relâchez, elle conserve sa vitesse, plafonnée pour qu'un coup sec ne puisse pas la projeter à travers un mur. Si le pointeur a bougé de moins de six pixels, cela compte comme un toucher et le lien s'ouvre.

Le haut de la boîte reste ouvert pendant que les icônes tombent d'en haut, puis se referme dès qu'elles sont toutes à l'intérieur, pour qu'un lancer ne puisse pas en faire sortir une par le haut.

## Les petits détails

- Un atterrissage brutal fait onduler le fond de points de la page.
- La simulation se met en veille quand tout est immobile et se met en pause quand la boîte est hors de l'écran : elle ne coûte donc rien pendant votre lecture.
- Les icônes ne tombent que lorsque la boîte est à l'écran. Sur mobile, c'est quand vous faites défiler la page jusqu'à elle.
- Sans JavaScript, ou avec la réduction des animations activée, les icônes restent immobiles en un tas bien rangé, et ce sont toujours des liens.

C'est un petit détail, mais il en dit plus sur les applications que nous concevons qu'une simple phrase. Si vous souhaitez quelque chose de similaire pour votre propre site, [parlez-nous-en](https://sawirstudio.com/fr/contact/).

**À lire aussi :** [Ce qu'il a fallu pour rendre nos sites prêts pour les agents](https://sawirstudio.com/fr/posts/making-our-sites-agent-ready/), ou [WordPress 7.2 arrive le 8 décembre : ce que nous testons dans nos extensions](https://sawirstudio.com/fr/posts/wordpress-7-2-what-plugin-authors-should-test/).
