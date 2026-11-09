---
title: Je sites klaar maken voor AI-agents | SawirStudio
url: https://sawirstudio.com/nl/posts/making-our-sites-agent-ready/
description: Zo maakten we onze productsites leesbaar voor AI-agents: Markdown-kopieën van elke pagina, Link-headers, llms.txt, Content Signals in robots.txt, een API-catalogus en Agent Skills.
---

# Wat er nodig was om onze sites klaar te maken voor AI-agents

Steeds meer bezoeken aan een website komen niet van mensen. Het zijn AI-agents die pagina's lezen om iemands vraag te beantwoorden, of die namens iemand een product proberen te gebruiken. Deze week zijn we door deze site en onze productsites gegaan, [SawirPost](https://sawirstudio.com/nl/projects/sawirpost/), [Sawirly](https://sawirstudio.com/nl/projects/sawirly/), [Rindraw](https://sawirstudio.com/nl/projects/rindraw/), [Math Daily](https://sawirstudio.com/nl/projects/math-daily/) en andere, om ze makkelijker leesbaar en bruikbaar te maken voor agents.

Als checklist gebruikten we [isitagentready.com](https://isitagentready.com). Die beoordeelt een site in niveaus. Dit hebben we per stap gedaan.

## Markdown-kopieën van elke pagina

Een agent die HTML leest, moet door navigatie, scripts en footers heen om een paar alinea's te vinden. Daarom heeft elke pagina op deze site nu een Markdown-kopie ernaast, op `index.md`, met alleen de hoofdinhoud van de pagina. Een klein script bouwt ze uit de HTML, zodat ze nooit uit de pas lopen.

Agents krijgen ze op twee manieren: door `index.md` achter een URL te zetten, of door `Accept: text/markdown` mee te sturen. Een kleine Cloudflare Worker voor de statische bestanden controleert die header en geeft de Markdown-kopie terug, met een grove tokentelling in een `x-markdown-tokens`-header. Browsers krijgen gewoon de normale pagina.

```
curl -H "Accept: text/markdown" https://sawirstudio.com/shopify/
```

## Link-headers

Elke HTML-pagina geeft nu aan waar de Markdown-kopie staat, in een HTTP `Link`-header met `rel="alternate"`. De homepage verwijst ook naar `llms.txt` en de sitemap. Een agent vindt alles met één request, zonder URL's te hoeven raden.

## llms.txt

`llms.txt` is een index van de site in platte tekst voor taalmodellen: een samenvatting in één regel en een lijst van alle pagina's met hun beschrijving, met links naar de Markdown-kopieën. De onze wordt door hetzelfde script gegenereerd uit de sitemap.

## Content Signals

`robots.txt` zegt wie mag crawlen, niet wat ze mogen doen met wat ze lezen. [Content Signals](https://contentsignals.org/) vult dat aan. Die van ons is één regel en staat alle drie de toepassingen toe, omdat we gevonden en geciteerd willen worden:

```
Content-Signal: ai-train=yes, search=yes, ai-input=yes
```

Met deze drie stappen kwam elke site op niveau 3.

## Een API-catalogus en Agent Skills

SawirPost en Sawirly hebben API's, dus die zijn verder gegaan. Beide publiceren een API-catalogus op `/.well-known/api-catalog` (RFC 9727): een klein JSON-document dat verwijst naar de OpenAPI-beschrijving en de documentatie. Een agent die op de site terechtkomt, vindt de API zonder één pagina te lezen.

Ze publiceren ook Agent Skills: korte instructies voor agents over hoe ze veelvoorkomende taken met het product uitvoeren. Daarmee kwamen beide op niveau 4.

## Waar we gestopt zijn

Voor niveau 5 moeten agents zich zelf kunnen aanmelden en authenticeren, via een OAuth-registratieflow voor agents. Een pagina die uitlegt hoe je je aanmeldt met e-mailadres en wachtwoord telt niet mee. Dat is een echte feature om te bouwen, met echte beveiligingsvragen, dus we denken er rustig over na in plaats van het te overhaasten.

## Was het de moeite waard?

Voor een statische site was het meeste een middagje werk: één script, één kleine Worker en een regel in `robots.txt`. Voor mensen die de site bezoeken veranderde er niets. We denken dat steeds meer vragen beantwoord worden door agents die sites als de onze lezen, en een nette Markdown-pagina is de makkelijkste manier om correct geciteerd te worden.

Voor klanten zetten we hetzelfde op als onderdeel van [Launch & growth](https://sawirstudio.com/nl/launch-and-growth/).

**Lees verder:** [Maak je WordPress-plugin bruikbaar voor AI-agents met de Abilities API](https://sawirstudio.com/nl/posts/wordpress-abilities-api-for-plugins/), of [De gooibare app-iconen op onze homepage](https://sawirstudio.com/nl/posts/throwable-app-icons-matter-js/).
