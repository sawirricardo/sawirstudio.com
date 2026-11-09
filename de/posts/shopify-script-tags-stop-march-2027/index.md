---
title: Shopify Script Tags enden im März 2027 | SawirStudio
url: https://sawirstudio.com/de/posts/shopify-script-tags-stop-march-2027/
description: Shopify stellt am 1. März 2027 die Ausführung von Script Tags im Online-Shop ein. Welche Apps betroffen sind, wie Sie Ihren Shop prüfen und was an ihre Stelle tritt.
---

# Shopify Script Tags funktionieren ab dem 1. März 2027 nicht mehr. Was Shop-Betreiber jetzt tun sollten

Am 1. März 2027 hört Shopify auf, Script Tags in Online-Shops einzubinden. Alles in Ihrem Shop, was noch auf diesem Weg geladen wird, etwa ein Bewertungs-Widget, eine Chat-Blase oder ein Tracking-Pixel, wird einfach nicht mehr angezeigt. Es gibt keine laute Fehlermeldung. Die Dinge fehlen schlicht.

Shopify hat dies im [August 2026](https://shopify.dev/changelog/posts/online-store-script-tags-deprecation) angekündigt. Das bedeutet es für Sie, wenn Sie einen Shop betreiben.

## Was ein Script Tag ist

Ein Script Tag ist eine Möglichkeit für eine App, über die API JavaScript in Ihren Shop einzubinden, ohne das Theme anzufassen. Genau deshalb waren Script Tags bei App-Entwicklern beliebt, und genau das ist auch das Problem: Script Tags erscheinen nicht im Theme-Editor, Sie können sie dort also weder sehen noch abschalten. Deinstallierte Apps haben ihre manchmal zurückgelassen.

## Der Zeitplan

- **1. Oktober 2026:** Apps können keine Script Tags mehr erstellen oder aktualisieren, in keiner API-Version. Bestehende laufen weiter.
- **1. März 2027:** Shopify lädt keine Script Tags des Online-Shops mehr.

Script Tags auf der Bestellstatusseite wurden früher und nach einem eigenen Zeitplan abgeschaltet. Falls dort bereits etwas nicht mehr angezeigt wird, liegt es daran.

## So prüfen Sie Ihren Shop

Shopify stellt Händlern keine Liste der Script Tags zur Verfügung, Sie müssen also selbst nachsehen:

1. **Listen Sie die Apps auf, die etwas in Ihrem Shop anzeigen.** Bewertungen, Chat, Pop-ups, Upsells, Treueprogramme, Analytics, Werbe-Pixel. Apps, die nur im Admin-Bereich arbeiten, sind nicht betroffen.
2. **Sehen Sie unter „App-Einbettungen“ im Theme-Editor nach.** Taucht dort eine App auf und ist eingeschaltet, nutzt sie die neuere Methode und ist in Ordnung.
3. **Sehen Sie in den Seitenquelltext.** Öffnen Sie Ihren Shop, rufen Sie den Quelltext auf und suchen Sie nach `asyncLoad`. In den meisten Shops stehen die URLs der Script Tags in diesem Block. Jede dortige Domain verweist meist auf eine App.
4. **Fragen Sie den App-Entwickler.** „Nutzt Ihre App Script Tags, und funktioniert sie nach dem 1. März 2027?“ ist eine Frage, die ein gepflegtes Produkt in einer Zeile beantworten kann.

Shopify drängt Apps seit Jahren zu Theme App Extensions, die meisten gut gepflegten Apps sind daher schon in Ordnung. Das Risiko liegt bei alten Apps, bei individuell entwickelten Apps, die vor Jahren jemand für Sie gebaut hat, und bei Überresten von Apps, die Sie nicht mehr nutzen.

## Was Script Tags ersetzt

- **App-Embed-Blöcke**, Teil einer Theme App Extension, für alles, was auf der Seite sichtbar ist. Sie schalten sie im Theme-Editor ein, sodass Sie sie endlich sehen können.
- **Web Pixels** für Skripte, die nur Analyse- oder Conversion-Daten erfassen. Sie laufen in einer Sandbox und brauchen von Ihnen nichts mehr, sobald die App sie unterstützt.

## Wenn eine App nicht mehr aktualisiert wird

Ersetzen Sie sie vor dem März oder lassen Sie ihre Funktion neu entwickeln. Bei einer individuellen App besteht die Arbeit meist darin, das Skript in eine Theme App Extension zu verschieben, oder in ein Web Pixel, falls es nur Daten erfasst. Löschen Sie bei dieser Gelegenheit auch die Script Tags von Apps, die Sie deinstalliert haben. Die Mutation `scriptTagDelete` funktioniert weiterhin.

Wenn Sie eine individuelle Shopify-App haben, die umgestellt werden muss, oder Sie nicht sicher sind, was in Ihrem Shop geladen wird, [schicken Sie uns Ihre Shop-Adresse](https://sawirstudio.com/de/contact/), und wir sagen Ihnen, was wir finden. Mehr zu unserer Arbeit mit Shopify finden Sie auf [dieser Seite](https://sawirstudio.com/de/shopify/).

**Weiterlesen:** [Shopifys API 2026-10: Was in Ihren Integrationen kaputtgeht](https://sawirstudio.com/de/posts/shopify-2026-10-api-what-breaks/) oder [Shopify mit Odoo 19 Community verbinden, Schritt für Schritt](https://sawirstudio.com/de/posts/connect-shopify-to-odoo-19-community/).
