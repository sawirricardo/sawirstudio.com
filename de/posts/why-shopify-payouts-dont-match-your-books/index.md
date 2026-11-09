---
title: Warum Shopify-Auszahlungen nicht zu Ihrer Buchhaltung passen | SawirStudio
url: https://sawirstudio.com/de/posts/why-shopify-payouts-dont-match-your-books/
description: Warum das Geld, das Shopify Payments auf Ihr Bankkonto überweist, nie zu Ihrem Umsatzbericht passt: Gebühren, Rückerstattungen, Zahlungsstreitigkeiten, Geschenkgutscheine, Währungsumrechnung und Zeitversatz, und wie Sie alles abstimmen.
---

# Warum Ihre Shopify-Auszahlungen nie zu Ihrer Buchhaltung passen

Ihr Umsatzbericht nennt eine Zahl. Ihre Bank eine andere. Dazwischen liegen Gebühren, Rückerstattungen, Rückbuchungen und ein paar Tage Zeitversatz, und jeden Monat verbringt jemand einen Nachmittag damit, das in einer Tabelle auseinanderzuklauben.

Mit all dem haben wir uns beim Schreiben des Buchhaltungsmoduls unseres [Shopify-Connectors für Odoo](https://sawirstudio.com/de/projects/odoo-shopify-connector/) beschäftigt. Hier sind die Gründe, warum die Zahlen abweichen, und wie Sie sie zur Deckung bringen.

## Eine Auszahlung ist ein Nettobetrag

Shopify Payments überweist Ihnen nicht Ihre Umsätze. Es überweist, was übrig bleibt, nachdem alles andere verrechnet wurde: Umsätze, abzüglich Rückerstattungen, abzüglich Zahlungsgebühren, abzüglich Zahlungsstreitigkeiten, zuzüglich oder abzüglich Korrekturen. Wenn Sie die Auszahlung als Umsatz verbuchen, stimmt Ihr Umsatz nicht, und Ihre Gebühren sind unsichtbar.

Die Lösung: Buchen Sie die Einzelposten, nicht die Summe. Umsätze gehen auf das Erlöskonto, Gebühren auf ein Aufwandskonto, Rückerstattungen werden gegen die Umsätze gebucht, und die Auszahlung selbst ist nur eine Geldbewegung von einem Verrechnungskonto auf die Bank.

## Eine Auszahlung deckt nicht nur einen Tag ab

Eine Auszahlung enthält Transaktionen mehrerer Tage, je nach Ihrem Auszahlungsrhythmus, und eine Bestellung vom Freitag kann in der Auszahlung vom Dienstag landen. Bestellungen und Zahlungseingänge nach Datum abzugleichen funktioniert nie lange. Gleichen Sie stattdessen nach Auszahlung ab: Shopify erfasst, aus welchen Saldotransaktionen sich jede Auszahlung zusammensetzt, und diese Liste ist die einzige verlässliche Verbindung zwischen einer Bestellung und dem Geld auf dem Bankkonto.

## Rückerstattungen werden von späteren Auszahlungen abgezogen

Wenn Sie eine Bestellung vom Vormonat erstatten, wird das Geld von einer Auszahlung in diesem Monat abgezogen. Ihr Umsatzbericht ordnet die Rückerstattung der ursprünglichen Bestellung zu. Ihre Bank zeigt heute einen kleineren Zahlungseingang. Beides stimmt.

## Zahlungsstreitigkeiten kosten doppelt

Wenn ein Kunde eine Abbuchung anficht, werden der strittige Betrag und eine Rückbuchungsgebühr von Ihrem Guthaben abgezogen. Gewinnen Sie den Fall, kommt der Betrag zurück, meist erst nach Wochen. Das sind drei getrennte Einträge für eine Bestellung, verteilt auf zwei oder drei Auszahlungen.

## Geschenkgutscheine sind beim Verkauf kein Umsatz

Wer einen Geschenkgutschein verkauft, nimmt Geld ein, das er später als Ware zurückgeben muss. Das gehört auf ein Verbindlichkeitskonto, bis der Gutschein eingelöst wird. Verbuchen Sie Gutscheinverkäufe als Umsatz, zählen Sie denselben Verkauf doppelt: einmal beim Kauf des Gutscheins und noch einmal bei der Einlösung.

## Währungsumrechnung

Wenn Sie in mehreren Währungen verkaufen, erscheinen Umrechnungen als eigene Saldotransaktionen, und der Kurs am Verkaufstag ist nicht der Kurs am Auszahlungstag. Mit der Shopify-API-Version 2026-10 gibt es einen eigenen Transaktionstyp `CURRENCY_CONVERSION`, sodass die Umrechnung auf einer eigenen Zeile gebucht werden kann, statt sich in einer Rundungsdifferenz zu verstecken.

## Rücklagen und Korrekturen

Shopify kann einen Teil Ihres Guthabens als Rücklage einbehalten, und Korrekturen erscheinen als Anpassungen. Beides ist weder ein Umsatz noch eine Gebühr, daher braucht beides ein eigenes Konto.

## So geht die Rechnung auf

Behandeln Sie jede Auszahlung wie eine kleine Buchung, die aus ihren Saldotransaktionen aufgebaut wird: Jeder Verkauf, jede Rückerstattung, Gebühr, Zahlungsstreitigkeit und Korrektur steht auf einer eigenen Zeile, und alles landet auf einem Verrechnungskonto, das der Bankeingang anschließend ausgleicht. Wenn die Zeilen zusammen nicht die Auszahlung ergeben, ist etwas falsch zugeordnet, und das sollten Sie wissen.

So arbeitet unser Buchhaltungsmodul für Odoo. Es plant jede Buchung anhand der exakten Beträge von Shopify, und wenn die Buchung nicht aufgeht, bricht es mit einer klaren Fehlermeldung ab. Es fügt nie eine „Ausgleichszeile“ hinzu, um die Differenz verschwinden zu lassen, denn diese Zeile ist nur der Fehler, aufgeschrieben an einer Stelle, wo ihn niemand sucht.

Wenn Ihr Shopify-Geld und Ihre Buchhaltung jeden Monat auseinanderlaufen, [helfen wir Ihnen gern](https://sawirstudio.com/de/shopify/).

**Weiterlesen:** [Shopify mit Odoo 19 Community verbinden, Schritt für Schritt](https://sawirstudio.com/de/posts/connect-shopify-to-odoo-19-community/) oder [Shopifys 2026-10-API: Was in Ihren Integrationen kaputtgeht](https://sawirstudio.com/de/posts/shopify-2026-10-api-what-breaks/).
