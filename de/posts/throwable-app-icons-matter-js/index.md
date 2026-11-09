---
title: Werfbare App-Icons mit Matter.js | SawirStudio
url: https://sawirstudio.com/de/posts/throwable-app-icons-matter-js/
description: Wie wir die Box mit den werfbaren App-Icons auf unserer Startseite gebaut haben: Matter.js für die Physik, DOM-Elemente statt Canvas, ein Schatten, der beim Drehen der Icons an Ort und Stelle bleibt, und Fallbacks.
---

# Die werfbaren App-Icons auf unserer Startseite

Auf unserer [Startseite](https://sawirstudio.com/de/) steht eine Box mit unseren App-Icons. Sie fallen hinein, sobald Sie die Seite öffnen, und Sie können sie greifen und werfen. Tippen Sie ein Icon an, ohne es zu ziehen, öffnet sich die Seite der jeweiligen App. Es hat einen Nachmittag gedauert, und einige Details waren interessanter, als wir erwartet hatten.

## Warum eine Box

Die erste Idee war, die Icons über den gesamten oberen Seitenbereich fliegen zu lassen. Dort stehen aber unsere Überschrift und die Buttons, und Icons, die auf „Start a project“ landen, sind ein schlechter Tausch für ein bisschen Spielerei. Eine Box neben dem Intro, mit einer Beschriftung am Rand, hält das Spiel in Grenzen. Sie füllt außerdem eine Lücke, die ohnehin da war: den leeren Platz rechts neben Intro und Buttons auf breiten Bildschirmen. Auf dem Smartphone sitzt sie unter den Buttons.

## Echte Physik dank Matter.js

Kreise, die gegeneinanderstoßen, sind leicht selbst zu programmieren. App-Icons sind aber abgerundete Quadrate, und ein Quadrat, das wie eine Münze rollt, sieht sofort falsch aus. Für sauberes Kippen und Stapeln braucht es eine richtige Physik-Engine, deshalb haben wir [Matter.js](https://brm.io/matter-js/) verwendet. Jedes Icon ist ein Rechteck mit abgerundeten Ecken (`chamfer`), passend zum Eckenradius des Icons von 22 %, innerhalb von vier statischen Wänden.

Die Engine ist 83 KB groß. Deshalb wird sie nur auf der Startseite geladen, erst nachdem die Seite steht, und nur, wenn der Besucher keine reduzierten Bewegungen eingestellt hat.

## DOM-Elemente statt Canvas

Matter.js kann auf ein Canvas zeichnen, wir nutzen es aber nur für die Berechnungen. Die Icons sind ganz normale Links mit Bildern, die in jedem Frame per CSS-Transform positioniert werden:

```
st.setProperty("--x", (b.position.x - size / 2) + "px");
st.setProperty("--y", (b.position.y - size / 2) + "px");
st.setProperty("--a", b.angle + "rad");
```

So ist ein Tippen ein echter Link, die Bilder sind bei jeder Zoomstufe scharf, und die Icons haben denselben Rahmen und Stil wie überall sonst auf der Website.

## Der Schatten, der an Ort und Stelle bleibt

Unser Stil setzt bei allem einen harten Schatten nach unten rechts. Legt man den auf ein Element, das sich dreht, dreht sich der Schatten mit, und der Schatten eines wirbelnden Icons schwingt dann wie ein Uhrzeiger herum. Die Lösung: Wir drehen den Versatz des Schattens in jedem Frame in die Gegenrichtung:

```
// keep the shadow at (3px, 3px) in screen space while the icon turns by a
st.setProperty("--shadow-x", 3 * (Math.cos(a) + Math.sin(a)) + "px");
st.setProperty("--shadow-y", 3 * (Math.cos(a) - Math.sin(a)) + "px");
```

## Greifen und Werfen

Beim Drücken befestigen wir eine Feder (ein Matter.js-Constraint) zwischen dem Zeiger und genau der Stelle am Icon, die Sie gegriffen haben, nicht an dessen Mittelpunkt. So schwingt es natürlich, wenn Sie es an einer Ecke halten. Lassen Sie los, behält es seine Geschwindigkeit, begrenzt, damit ein kräftiger Schwung es nicht durch eine Wand schleudert. Hat sich der Zeiger weniger als sechs Pixel bewegt, zählt das als Tippen, und der Link wird geöffnet.

Die Oberseite der Box ist offen, solange die Icons von oben hineinfallen, und schließt sich, sobald alle drin sind. So kann bei einem Wurf keines oben hinausfliegen.

## Kleine Details

- Ein harter Aufprall schickt eine Welle durch den gepunkteten Hintergrund der Seite.
- Die Simulation schläft, wenn alles still liegt, und pausiert, wenn die Box nicht im sichtbaren Bereich ist. Sie kostet also nichts, während Sie lesen.
- Die Icons fallen erst hinein, wenn die Box im sichtbaren Bereich ist. Auf dem Smartphone ist das der Fall, wenn Sie zu ihr scrollen.
- Ohne JavaScript oder bei aktivierter Bewegungsreduzierung liegen die Icons ruhig in einem ordentlichen Haufen, und sie sind weiterhin Links.

Es ist eine Kleinigkeit, aber sie sagt mehr über die Apps, die wir bauen, als ein Satz es könnte. Wenn Sie so etwas für Ihre eigene Website möchten, [sprechen Sie uns an](https://sawirstudio.com/de/contact/).

**Weiterlesen:** [Was nötig war, um unsere Websites agentenfähig zu machen](https://sawirstudio.com/de/posts/making-our-sites-agent-ready/) oder [WordPress 7.2 erscheint am 8. Dezember: Das testen wir in unseren Plugins](https://sawirstudio.com/de/posts/wordpress-7-2-what-plugin-authors-should-test/).
