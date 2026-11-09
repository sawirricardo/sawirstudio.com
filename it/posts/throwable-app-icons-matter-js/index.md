---
title: Icone di app da lanciare con Matter.js | SawirStudio
url: https://sawirstudio.com/it/posts/throwable-app-icons-matter-js/
description: Come abbiamo costruito il box di icone di app che potete lanciare sulla nostra home page: Matter.js per la fisica, elementi DOM invece di un canvas, un'ombra che resta ferma mentre le icone girano, e soluzioni di ripiego.
---

# Le icone delle app lanciabili sulla nostra home page

La nostra [home page](https://sawirstudio.com/it/) ha un box con dentro le icone delle nostre app. Cadono quando arrivate, e potete afferrarle e lanciarle. Toccatene una senza trascinarla e si apre la pagina di quell'app. Ci è voluto un pomeriggio, e alcuni dettagli si sono rivelati più interessanti del previsto.

## Perché un box

La prima idea era lasciare che le icone volassero per tutta la parte alta della pagina. Ma lì ci sono il titolo e i pulsanti, e le icone che atterrano sopra «Start a project» sono un brutto prezzo da pagare per un po' di divertimento. Un box accanto all'introduzione, con un'etichetta sul bordo, tiene il gioco circoscritto. Occupa anche uno spazio che c'era già: quello vuoto a destra dell'introduzione e dei pulsanti sugli schermi larghi. Sugli smartphone si trova sotto i pulsanti.

## Fisica vera, con Matter.js

Dei cerchi che si urtano sono facili da scrivere da soli. Ma le icone delle app sono quadrati arrotondati, e un quadrato che rotola come una moneta sembra subito sbagliato. Per far ruotare e impilare le icone in modo credibile serve un vero motore fisico, quindi abbiamo usato [Matter.js](https://brm.io/matter-js/). Ogni icona è un rettangolo con gli angoli arrotondati (`chamfer`), che corrisponde al raggio degli angoli del 22% dell'icona, dentro quattro pareti statiche.

Il motore pesa 83KB, quindi si carica solo sulla home page, solo dopo che la pagina è pronta, e solo se il visitatore non ha chiesto di ridurre le animazioni.

## Elementi DOM, non un canvas

Matter.js può disegnare su un canvas, ma noi lo usiamo solo per i calcoli. Le icone sono normali link con immagini, posizionati a ogni frame con una transform CSS:

```
st.setProperty("--x", (b.position.x - size / 2) + "px");
st.setProperty("--y", (b.position.y - size / 2) + "px");
st.setProperty("--a", b.angle + "rad");
```

Così un tocco è un vero link, le immagini restano nitide a qualsiasi zoom e le icone mantengono lo stesso bordo e lo stesso stile del resto del sito.

## L'ombra che resta ferma

Il nostro stile prevede un'ombra netta in basso a destra di ogni elemento. Se la applicate a un elemento che ruota, l'ombra ruota con lui, e l'ombra di un'icona che gira oscilla come la lancetta di un orologio. La soluzione è ruotare lo scostamento dell'ombra nel verso opposto a ogni frame:

```
// keep the shadow at (3px, 3px) in screen space while the icon turns by a
st.setProperty("--shadow-x", 3 * (Math.cos(a) + Math.sin(a)) + "px");
st.setProperty("--shadow-y", 3 * (Math.cos(a) - Math.sin(a)) + "px");
```

## Afferrare e lanciare

Quando si preme il puntatore colleghiamo una molla (un vincolo di Matter.js) tra il puntatore e il punto esatto dell'icona che avete afferrato, non il suo centro, così oscilla in modo naturale se la tenete per un angolo. Quando la lasciate, mantiene la sua velocità, con un limite perché un colpo troppo forte non possa farla attraversare una parete. Se il puntatore si è spostato di meno di sei pixel, conta come un tocco e il link si apre.

La parte superiore del box resta aperta mentre le icone cadono dall'alto, poi si chiude quando sono tutte dentro, così un lancio non ne fa uscire nessuna dall'alto.

## Piccoli dettagli

- Un atterraggio violento fa partire un'onda attraverso lo sfondo puntinato della pagina.
- La simulazione si ferma quando tutto è immobile e va in pausa quando il box è fuori dallo schermo, quindi non pesa mentre leggete.
- Le icone cadono solo quando il box è visibile sullo schermo. Su smartphone, succede quando scorrete fino a lì.
- Senza JavaScript, o con le animazioni ridotte attivate, le icone restano ferme in un mucchietto ordinato e sono comunque dei link.

È una piccola cosa, ma dice delle app che realizziamo più di quanto farebbe una frase. Se volete qualcosa di simile per il vostro sito, [scriveteci](https://sawirstudio.com/it/contact/).

**Continuate a leggere:** [Cosa ci è voluto per rendere i nostri siti pronti per gli agenti](https://sawirstudio.com/it/posts/making-our-sites-agent-ready/), oppure [WordPress 7.2 arriva l'8 dicembre: cosa stiamo testando nei nostri plugin](https://sawirstudio.com/it/posts/wordpress-7-2-what-plugin-authors-should-test/).
