---
title: Sawir Insights | SawirStudio
url: https://sawirstudio.com/it/projects/wordpress-plugins/sawir-insights/
description: Heatmap, registrazioni delle sessioni, sondaggi, funnel e analisi dei moduli. Tutto resta sul vostro server.
---

# Sawir Insights

![Banner di Sawir Insights](https://sawirstudio.com/assets/images/plugins/sawir-insights-banner.webp)

Sawir Insights mostra come i visitatori usano il vostro sito, con heatmap, registrazioni delle sessioni, sondaggi, funnel e analisi dei moduli, e tiene tutto sul vostro server.

[Scaricatelo gratis da WordPress.org](https://wordpress.org/plugins/sawir-insights/)

## Cosa fa

Sawir Insights è uno strumento di analisi comportamentale per WordPress: heatmap, registrazioni delle sessioni, sondaggi, widget per i feedback, analisi dei moduli, funnel di conversione e una dashboard che mette tutto insieme.

Nelle reti multisite, ogni sito traccia e configura le proprie analytics in modo indipendente.

In breve: un'alternativa self-hosted a Hotjar o Microsoft Clarity, in cui i dati dei visitatori restano nel vostro database.

Il motivo per cui esiste è semplice. Gli strumenti più affermati di questa categoria sono validi, ma funzionano inviando a server di terzi una registrazione di tutto ciò che fanno i vostri visitatori. Per molti siti, soprattutto nell'UE, questo va da «poco comodo» a «inattuabile». Volevo lo stesso tipo di strumenti, ma con dati che non lasciano mai il sito che li ha raccolti. Ecco quindi il vincolo attorno a cui è costruito questo plugin: il vostro database, la vostra cartella degli upload, le vostre regole di conservazione. Non c'è nessun SaaS dietro, nessun account da creare e nessuna comunicazione verso l'esterno, di alcun tipo.

### Cosa fa

- Heatmap di clic, movimenti del mouse e scorrimento, mostrate sulle vostre pagine reali e suddivise per desktop, tablet e mobile.
- Registrazioni delle sessioni con riproduzione in wp-admin. I valori inseriti nei campi sono sempre mascherati; la registrazione resta disattivata finché non la attivate voi.
- Sondaggi popup e una scheda di feedback sempre visibile, con targeting per URL, dispositivo, percentuale di traffico, tempo sulla pagina, profondità di scorrimento o intento di uscita.
- Analisi dei moduli: dove le persone esitano, cosa lasciano vuoto, dove abbandonano. Solo selettori e tempi, mai ciò che qualcuno ha scritto.
- Funnel di conversione ordinati, con report giornalieri.
- Una dashboard con KPI di traffico e variazioni rispetto al periodo precedente, un grafico dell'andamento giornaliero e il dettaglio per dispositivo, browser, pagina, referrer, campagna, funnel e registrazioni.

### Cosa non fa, volutamente

Penso che sia giusto sapere cosa uno strumento si rifiuta di fare prima di installarlo.

- Non memorizza gli indirizzi IP. Né in forma di hash, né troncati: mai.
- Non registra mai i valori dei moduli. Il livello di acquisizione invia solo metadati, quindi non c'è nulla che possa essere divulgato.
- Non carica nulla da una CDN e non invia nulla a una CDN. La libreria di riproduzione rrweb è inclusa nel plugin.
- Non conserva i dati grezzi per sempre. Gli eventi vengono aggregati ed eliminati secondo una pianificazione di conservazione che decidete voi; le registrazioni hanno una propria conservazione e un limite di spazio su disco, con eliminazione a partire dalle più vecchie, perché un plugin che riempie in silenzio il disco del vostro hosting è un cattivo plugin.

Il tracker lato frontend è un piccolo script in JavaScript puro, che invia i dati a blocchi tramite `sendBeacon`, con controlli di campionamento. I visitatori che non vengono campionati per la registrazione non scaricano mai la libreria di riproduzione.

### Approccio alla privacy

Il Do Not Track è rispettato per impostazione predefinita. Un'API per la modalità consenso (`window.swinsConsent`) si integra con il gestore del consenso che usate già. Sessioni, eventi e registrazioni sono collegati agli strumenti nativi di WordPress per l'esportazione e la cancellazione dei dati personali. Se la vostra specifica installazione sia conforme al GDPR dipende dalle vostre informative e dalla base giuridica, e questo nessun plugin può farlo al posto vostro, ma il lavoro di minimizzazione dei dati è già fatto.

### Supporto

Il supporto passa dal forum di wordpress.org dedicato a questo plugin, oppure via email a ricardo@sawirstudio.com.

## Domande

### Come si confronta con le alternative?

I servizi di analytics in hosting come Hotjar e Microsoft Clarity sono strumenti eccellenti, con piani gratuiti generosi, ma funzionano inviando ai loro server una registrazione del comportamento dei vostri visitatori. Sawir Insights rinuncia a questa infrastruttura condivisa tra più siti in nome di un unico vincolo: tutto viene raccolto, archiviato e consultato sul vostro sito. Scegliete un servizio in hosting se volete le sue funzioni extra e non vi disturba il trattamento da parte di terzi; scegliete Sawir Insights se l'uscita dei dati dei visitatori dal vostro server è un ostacolo insormontabile. Registrazioni e heatmap risiedono nel vostro database e nella vostra cartella degli upload, quindi le impostazioni di conservazione contano qui più che in un SaaS.

### Sawir Insights è conforme al GDPR?

Il plugin è pensato per installazioni attente alla privacy: i dati restano sul vostro server, gli indirizzi IP non vengono mai memorizzati, la conservazione è configurabile, il Do Not Track è rispettato e gli strumenti di esportazione e cancellazione di WordPress sono integrati. La conformità vera e propria dipende dalla vostra configurazione, dalle informative, dalla base giuridica e dall'implementazione del consenso: rivolgetevi a un consulente legale qualificato per il vostro sito.

### In cosa si differenzia da Hotjar o Microsoft Clarity?

Quelli sono servizi in hosting: strumenti validi, ma il comportamento dei vostri visitatori viene registrato sui loro server, secondo i loro termini. Sawir Insights copre la stessa categoria (heatmap, registrazioni delle sessioni, sondaggi, funnel, analisi dei moduli) come plugin che tiene tutto sul vostro server, senza account, senza servizi esterni e senza condivisione dei dati.

### I dati dei visitatori lasciano il mio server?

No. Il plugin non invia dati comportamentali a Sawir Studio né ad alcun fornitore di analytics. Gli eventi risiedono nel vostro database WordPress; i file delle registrazioni si trovano nella vostra cartella degli upload protetta.

## Si abbina bene a

- [Sawir Pulse](https://sawirstudio.com/it/projects/wordpress-plugins/sawir-pulse/): mostra ai clienti cosa succede davvero nel vostro negozio, come gli ordini recenti, le persone che stanno guardando un prodotto e le nuove recensioni, con piccole notifiche in un angolo della pagina.
- [Sawir Live Translate](https://sawirstudio.com/it/projects/wordpress-plugins/sawir-live-translate/): traduce le vostre pagine direttamente dove si trovano: scegliete una lingua, traducete in automatico ciò che è visibile, salvate e offrite ai visitatori un selettore di lingua.

Scoprite tutti i nostri [plugin WordPress gratuiti](https://sawirstudio.com/it/projects/wordpress-plugins/).

- **Prezzo** Gratuito
- **Versione** 1.0.0
- **Requisiti** WordPress 6.2, PHP 7.4
- **Licenza** GPLv2 o successiva
- **WordPress.org** <https://wordpress.org/plugins/sawir-insights/>

Basta parlare di noi.

## Come possiamo aiutarvi?

[Avviate un progetto](https://sawirstudio.com/it/contact/)
