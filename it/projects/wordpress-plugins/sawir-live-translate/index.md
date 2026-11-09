---
title: Sawir Live Translate | SawirStudio
url: https://sawirstudio.com/it/projects/wordpress-plugins/sawir-live-translate/
description: Sawir Live Translate traduce le vostre pagine direttamente dove si trovano: scegliete una lingua, traducete in automatico ciò che è visibile, salvate e offrite ai visitatori un selettore di lingua.
---

# Sawir Live Translate

![Banner di Sawir Live Translate](https://sawirstudio.com/assets/images/plugins/sawir-live-translate-banner.webp)

Sawir Live Translate traduce le vostre pagine direttamente dove si trovano: scegliete una lingua, traducete in automatico ciò che è visibile, salvate e offrite ai visitatori un selettore di lingua.

[Scaricatelo gratis da WordPress.org](https://wordpress.org/plugins/sawir-live-translate/)

## Cosa fa

Aprite una qualsiasi pagina del vostro sito, cliccate su «Translate this page» nella barra di amministrazione, scegliete la lingua di destinazione e traducete in automatico ogni stringa visibile (con un servizio di traduzione automatica gratuito) oppure modificate le traduzioni a mano, vedendo la pagina aggiornarsi in tempo reale.

Le traduzioni salvate vengono inserite lato server nella risposta HTML (ottimo per la SEO: nessun lampo di testo non tradotto e alternative hreflang nell'head), e i visitatori cambiano lingua con un selettore basato su link (in basso a destra).

Gli URL usano un prefisso di lingua nel percorso, ad esempio /es/about per lo spagnolo e /about per la lingua di origine (che resta quella canonica, senza prefisso). La lingua viene stabilita in quest'ordine: percorso /es/ → ?lang=es → cookie salvato → Accept-Language del browser → lingua di origine. La radice e gli URL senza prefisso rilevano in automatico la lingua del browser del visitatore, senza fare redirect (viene inviato un header Vary, così le cache condivise funzionano correttamente).

1. Impostazioni → Live Translate: impostate la lingua di origine e spuntate le lingue da offrire. 2. Aprite una pagina del front-end e cliccate su «🌐 Translate this page» nella barra di amministrazione. 3. Scegliete una lingua, premete «Auto-translate empty» (oppure modificate le stringhe), poi «Save». 4. I visitatori usano il selettore (in basso a destra); ogni lingua ha il proprio URL /xx/.

Nota: i prefissi di lingua nel percorso presuppongono un'installazione nella radice del dominio e l'assenza di pagine di primo livello il cui slug sia un codice lingua di due lettere.

## Domande

### Come si confronta con le alternative?

Framework multilingua come WPML, Polylang e TranslatePress trattano i contenuti tradotti come parte strutturale del sito, con flussi di lavoro di modifica per ogni lingua, servizi di traduzione e un'integrazione profonda con i page builder e WooCommerce. Sawir Live Translate è volutamente più leggero: aprite una pagina, traducete ciò che è visibile (in automatico o a mano) e servitela da un URL con prefisso di lingua e tag hreflang. Se state costruendo un sito completamente multilingua, con menu, prodotti e flussi editoriali tradotti, usate uno di questi framework. Se invece volete che le pagine esistenti siano leggibili in un'altra lingua già nel pomeriggio, è esattamente quello che fa questo plugin.

### Cambio lingua ma non cambia nulla

Quasi sempre significa che per quella lingua non c'è ancora nessuna traduzione salvata. Controllate in Impostazioni → Live Translate → Saved translations: se la lingua non è nell'elenco (o mostra 0 stringhe), riaprite l'editor della pagina, traducete in automatico e cliccate su Save. L'anteprima live nell'editor non viene salvata finché non lo fate.

### Alcune stringhe restano non tradotte

L'editor nella pagina vede solo le stringhe della pagina che avete aperto. I testi presenti su altre pagine richiedono che anche quelle pagine vengano tradotte, oppure potete lanciare la scansione dell'intero sito. Le stringhe inserite da slider o script dopo il caricamento della pagina potrebbero richiedere un inserimento manuale tramite «Edit all site strings».

### Ho modificato il contenuto e la traduzione è sparita

Le traduzioni sono associate al testo originale esatto. Se riscrivete una frase, traducete la nuova versione una volta e salvate.

## Si abbina bene a

- [Sawir Insights](https://sawirstudio.com/it/projects/wordpress-plugins/sawir-insights/): mostra come i visitatori usano il vostro sito, con heatmap, registrazioni delle sessioni, sondaggi, funnel e analisi dei form, e tiene tutto sul vostro server..
- [Sawir Media Folders](https://sawirstudio.com/it/projects/wordpress-plugins/sawir-media-folders/): porta le cartelle nella libreria multimediale di WordPress, quante ne servono, con il drag and drop per mettere in ordine i file..

Scoprite tutti i nostri [plugin WordPress gratuiti](https://sawirstudio.com/it/projects/wordpress-plugins/).

- **Prezzo** Gratuito
- **Versione** 1.1.0
- **Richiede** WordPress 5.8, PHP 7.4
- **Licenza** GPL-2.0-or-later
- **WordPress.org** <https://wordpress.org/plugins/sawir-live-translate/>

Basta parlare di noi.

## Come possiamo aiutarvi?

[Avviate un progetto](https://sawirstudio.com/it/contact/)
