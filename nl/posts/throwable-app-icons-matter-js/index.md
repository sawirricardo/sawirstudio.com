---
title: Gooibare app-iconen met Matter.js | SawirStudio
url: https://sawirstudio.com/nl/posts/throwable-app-icons-matter-js/
description: Hoe we de box met gooibare app-iconen op onze homepage bouwden: Matter.js voor de fysica, DOM-elementen in plaats van een canvas, een schaduw die blijft staan terwijl iconen draaien, en fallbacks.
---

# De gooibare app-iconen op onze homepage

Op onze [homepage](https://sawirstudio.com/nl/) staat een box met onze app-iconen. Ze vallen erin zodra je de pagina opent, en je kunt ze oppakken en gooien. Tik op een icoon zonder te slepen en de pagina van die app gaat open. Het kostte een middag, en een paar details waren interessanter dan we hadden verwacht.

## Waarom een box

Het eerste idee was om de iconen over de hele bovenkant van de pagina te laten vliegen. Maar daar staan onze kop en knoppen, en iconen die op “Start a project” landen zijn een slechte ruil voor wat plezier. Een box naast de intro, met een label op de rand, houdt het spel binnen de lijntjes. Het past ook in een gat dat er al was: de lege ruimte rechts van de intro en de knoppen op een breed scherm. Op telefoons staat de box onder de knoppen.

## Echte fysica, dankzij Matter.js

Cirkels die tegen elkaar botsen schrijf je makkelijk zelf. Maar app-iconen zijn afgeronde vierkanten, en een vierkant dat rolt als een munt ziet er meteen verkeerd uit. Voor goed tuimelen en stapelen heb je een echte physics-engine nodig, dus gebruikten we [Matter.js](https://brm.io/matter-js/). Elk icoon is een rechthoek met afgeronde hoeken (`chamfer`) die past bij de hoekradius van 22% van het icoon, binnen vier statische wanden.

De engine is 83KB groot, dus we laden hem alleen op de homepage, pas als de pagina staat, en alleen als de bezoeker niet om minder beweging heeft gevraagd.

## DOM-elementen, geen canvas

Matter.js kan tekenen op een canvas, maar wij gebruiken het alleen voor de berekeningen. De iconen zijn gewone links met afbeeldingen, die elk frame worden gepositioneerd met een CSS-transform:

```
st.setProperty("--x", (b.position.x - size / 2) + "px");
st.setProperty("--y", (b.position.y - size / 2) + "px");
st.setProperty("--a", b.angle + "rad");
```

Een tik is dus een echte link, de afbeeldingen blijven scherp op elk zoomniveau, en de iconen hebben dezelfde rand en stijl als overal elders op de site.

## De schaduw die blijft staan

In onze stijl hebben we een harde schaduw rechtsonder alles. Zet je die op een element dat draait, dan draait de schaduw mee, en de schaduw van een draaiend icoon zwaait rond als de wijzer van een klok. De oplossing is om de offset van de schaduw elk frame de andere kant op te draaien:

```
// keep the shadow at (3px, 3px) in screen space while the icon turns by a
st.setProperty("--shadow-x", 3 * (Math.cos(a) + Math.sin(a)) + "px");
st.setProperty("--shadow-y", 3 * (Math.cos(a) - Math.sin(a)) + "px");
```

## Vastpakken en gooien

Bij het indrukken koppelen we een veer (een Matter.js-constraint) tussen de aanwijzer en het exacte punt op het icoon dat je vastpakt, niet het midden. Zo zwaait het natuurlijk heen en weer als je het bij een hoek vasthoudt. Laat je los, dan behoudt het zijn snelheid, begrensd zodat een harde zwiep het niet door een wand kan sturen. Als de aanwijzer minder dan zes pixels bewoog, telt het als een tik en gaat de link open.

De bovenkant van de box is open terwijl de iconen van boven naar binnen vallen. Zodra ze er allemaal in zitten, gaat hij dicht, zodat je er geen over de bovenrand kunt gooien.

## Kleine dingen

- Bij een harde landing trilt er een rimpeling door de gestippelde achtergrond van de pagina.
- De simulatie slaapt als alles stilligt en pauzeert als de box buiten beeld is, dus hij kost niets terwijl je leest.
- De iconen vallen pas als de box in beeld is. Op een telefoon is dat wanneer je ernaartoe scrolt.
- Zonder JavaScript, of met beperkte beweging aan, liggen de iconen stil in een nette hoop, en het blijven gewoon links.

Het is iets kleins, maar het zegt meer over de apps die we maken dan een zin zou doen. Wil je zoiets voor je eigen site? [Laat het ons weten](https://sawirstudio.com/nl/contact/).

**Lees verder:** [What It Took to Make Our Sites Agent-Ready](https://sawirstudio.com/nl/posts/making-our-sites-agent-ready/), of [WordPress 7.2 Lands December 8: What We're Testing in Our Plugins](https://sawirstudio.com/nl/posts/wordpress-7-2-what-plugin-authors-should-test/).
