---
title: Website agententauglich machen | SawirStudio
url: https://sawirstudio.com/de/posts/making-our-sites-agent-ready/
description: So haben wir unsere Produkt-Websites für KI-Agenten lesbar gemacht: Markdown-Kopien jeder Seite, Link-Header, llms.txt, Content Signals in der robots.txt, ein API-Katalog und Agent Skills.
---

# Was es brauchte, unsere Websites agententauglich zu machen

Immer mehr Besuche auf einer Website stammen nicht von Menschen. Es sind KI-Agenten, die Seiten lesen, um jemandes Frage zu beantworten, oder die im Auftrag einer Person ein Produkt nutzen wollen. Diese Woche haben wir diese Website und unsere Produkt-Websites durchgesehen, [SawirPost](https://sawirstudio.com/de/projects/sawirpost/), [Sawirly](https://sawirstudio.com/de/projects/sawirly/), [Rindraw](https://sawirstudio.com/de/projects/rindraw/), [Math Daily](https://sawirstudio.com/de/projects/math-daily/) und weitere, damit Agenten sie leichter lesen und nutzen können.

Als Checkliste diente uns [isitagentready.com](https://isitagentready.com). Die Seite bewertet eine Website in Stufen. Hier sehen Sie, was jede Stufe erfordert hat.

## Markdown-Kopien jeder Seite

Ein Agent, der HTML liest, muss sich durch Navigation, Skripte und Footer arbeiten, um ein paar Absätze zu finden. Deshalb hat jede Seite dieser Website jetzt eine danebenliegende Markdown-Kopie unter `index.md`, die nur den Hauptinhalt der Seite enthält. Ein kleines Skript erzeugt sie aus dem HTML, sodass beide Versionen nie auseinanderlaufen.

Agenten erhalten sie auf zwei Wegen: indem sie `index.md` an eine URL anhängen oder indem sie `Accept: text/markdown` anfordern. Ein winziger Cloudflare Worker vor den statischen Dateien prüft diesen Header und liefert die Markdown-Kopie aus, samt grober Token-Zahl in einem `x-markdown-tokens`-Header. Browser bekommen weiterhin die normale Seite.

```
curl -H "Accept: text/markdown" https://sawirstudio.com/shopify/
```

## Link-Header

Jede HTML-Seite gibt jetzt an, wo ihre Markdown-Kopie liegt, und zwar in einem HTTP-`Link`-Header mit `rel="alternate"`. Die Startseite verweist außerdem auf `llms.txt` und die Sitemap. Ein Agent findet so mit einer einzigen Anfrage alles, ohne URLs raten zu müssen.

## llms.txt

`llms.txt` ist ein Klartext-Index der Website für Sprachmodelle: eine einzeilige Zusammenfassung und eine Liste aller Seiten mit Beschreibung, verlinkt auf die Markdown-Kopien. Unsere wird vom selben Skript aus der Sitemap erzeugt.

## Content Signals

`robots.txt` legt fest, wer crawlen darf, aber nicht, was damit gemacht werden darf. [Content Signals](https://contentsignals.org/) ergänzt genau das. Unsere Angabe besteht aus einer Zeile und erlaubt alle drei Nutzungsarten, weil wir gefunden und zitiert werden möchten:

```
Content-Signal: ai-train=yes, search=yes, ai-input=yes
```

Mit diesen drei Schritten erreichten alle Websites Stufe 3.

## Ein API-Katalog und Agent Skills

SawirPost und Sawirly haben APIs und sind deshalb weitergegangen. Beide veröffentlichen unter `/.well-known/api-catalog` einen API-Katalog (RFC 9727): ein kleines JSON-Dokument, das auf die OpenAPI-Beschreibung und die Dokumentation verweist. Ein Agent, der auf der Website landet, findet die API, ohne eine einzige Seite zu lesen.

Außerdem veröffentlichen sie Agent Skills, kurze Anleitungen für Agenten, wie sich gängige Aufgaben mit dem Produkt erledigen lassen. Damit erreichten beide Stufe 4.

## Wo wir aufgehört haben

Für Stufe 5 müssen Agenten sich selbstständig registrieren und authentifizieren können, über einen OAuth-Registrierungsablauf für Agenten. Eine Seite, die beschreibt, wie man sich mit E-Mail und Passwort anmeldet, zählt nicht. Das ist eine echte Funktion mit echten Sicherheitsfragen, die wir uns in Ruhe überlegen, statt sie zu überstürzen.

## Hat es sich gelohnt?

Bei einer statischen Website war das meiste ein Nachmittag Arbeit: ein Skript, ein kleiner Worker und eine Zeile in `robots.txt`. Für Besucherinnen und Besucher hat sich nichts geändert. Wir gehen davon aus, dass künftig mehr Fragen von Agenten beantwortet werden, die Websites wie unsere lesen, und eine saubere Markdown-Seite ist der einfachste Weg, korrekt zitiert zu werden.

Für Kunden richten wir dasselbe im Rahmen von [Launch & growth](https://sawirstudio.com/de/launch-and-growth/) ein.

**Weiterlesen:** [WordPress-Plugin per Abilities API für KI-Agenten nutzbar machen](https://sawirstudio.com/de/posts/wordpress-abilities-api-for-plugins/) oder [Die werfbaren App-Icons auf unserer Startseite](https://sawirstudio.com/de/posts/throwable-app-icons-matter-js/).
