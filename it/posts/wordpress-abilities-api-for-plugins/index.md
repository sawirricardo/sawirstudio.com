---
title: Abilities API di WordPress per i plugin | SawirStudio
url: https://sawirstudio.com/it/posts/wordpress-abilities-api-for-plugins/
description: Come registrare le abilities in un plugin WordPress, in modo che agenti AI e altri strumenti possano trovarle ed eseguirle in sicurezza: categorie, schemi, permessi, annotations ed endpoint REST.
---

# Rendete il vostro plugin WordPress utilizzabile dagli agenti AI con l'Abilities API

L'Abilities API, introdotta in WordPress 6.9, è un punto unico in cui un plugin dichiara cosa sa fare: un nome, una descrizione, quali dati richiede in ingresso, cosa restituisce e chi può eseguirlo. Agenti AI, il plugin WordPress AI, adattatori MCP e semplici script possono così trovare queste abilities e usarle, senza che ogni strumento debba imparare a mano ogni plugin.

Ecco come funziona, usando come esempio uno dei nostri [plugin WordPress](https://sawirstudio.com/it/projects/wordpress-plugins/).

## L'esempio

[Sawir Restock Notifier](https://sawirstudio.com/it/projects/wordpress-plugins/sawir-restock-notifier-for-woocommerce/) permette ai clienti di chiedere di essere avvisati via email quando un prodotto torna disponibile. Una domanda ovvia a cui un agente può rispondere è «quali prodotti esauriti hanno più persone in attesa?». Questa è un'ability di sola lettura.

## 1. Registrare una categoria

Ogni ability appartiene a una sola categoria, e le categorie vanno registrate per prime, con il loro hook:

```
add_action( 'wp_abilities_api_categories_init', function () {
	wp_register_ability_category( 'store-insights', array(
		'label'       => __( 'Store insights', 'sawir-restock-notifier-for-woocommerce' ),
		'description' => __( 'Read-only questions about the store.', 'sawir-restock-notifier-for-woocommerce' ),
	) );
} );
```

## 2. Registrare l'ability

```
add_action( 'wp_abilities_api_init', function () {
	wp_register_ability( 'sawir-restock/waiting-list', array(
		'label'       => __( 'Restock waiting list', 'sawir-restock-notifier-for-woocommerce' ),
		'description' => __( 'Out-of-stock products, with how many shoppers are waiting for each.', 'sawir-restock-notifier-for-woocommerce' ),
		'category'    => 'store-insights',
		'input_schema' => array(
			'type'       => 'object',
			'properties' => array(
				'limit' => array( 'type' => 'integer', 'minimum' => 1, 'maximum' => 50, 'default' => 10 ),
			),
		),
		'output_schema' => array(
			'type'  => 'array',
			'items' => array(
				'type'       => 'object',
				'properties' => array(
					'product_id' => array( 'type' => 'integer' ),
					'name'       => array( 'type' => 'string' ),
					'waiting'    => array( 'type' => 'integer' ),
				),
			),
		),
		'execute_callback'    => 'sawir_restock_waiting_list',
		'permission_callback' => function () {
			return current_user_can( 'manage_woocommerce' );
		},
		'meta' => array(
			'show_in_rest' => true,
			'annotations'  => array( 'readonly' => true ),
		),
	) );
} );
```

La callback è normale codice del plugin che restituisce dati conformi allo schema di output. Non ha nulla di specifico per l'AI.

## A cosa serve ogni parte

- **Nome**, nel formato `namespace/ability-name`. Usate lo slug del vostro plugin come namespace.
- **Descrizione.** È ciò che un agente legge per decidere se usare l'ability, quindi scrivetela pensando a qualcuno che non ha mai visto il vostro plugin.
- **Schemi di input e output**, in JSON Schema. Gli input vengono verificati rispetto agli schemi prima che il vostro codice venga eseguito, e gli strumenti sanno cosa riceveranno in risposta.
- **Permission callback.** L'unica parte che protegge davvero qualcosa. Qui controllate le capability come fareste per una schermata di amministrazione.
- **Annotations.** `readonly`, `destructive` e `idempotent` indicano agli strumenti quanta cautela usare. Sono suggerimenti, non sicurezza: la sicurezza è la permission callback.
- **`show_in_rest`.** Le abilities non sono visibili nella REST API a meno che non attiviate questa opzione.

## Eseguirla via REST

Le abilities esposte in REST sono elencate in `/wp-json/wp-abilities/v1/abilities`, e ciascuna ha un endpoint `/run`. Un'ability di sola lettura si esegue con GET, passando l'input come JSON con URL encoding nel parametro `input`. Una che modifica dei dati si esegue con POST, una distruttiva con DELETE. Ogni richiesta richiede un utente autenticato, e poi decide la vostra permission callback.

## E agli agenti, tramite MCP

L'MCP Adapter, sviluppato insieme al plugin WordPress AI, espone le abilities registrate ai client MCP come Claude. Non dovete scrivere nulla in più: un'ability ben descritta, con buoni schemi, è già un buon strumento.

## Cosa esporremmo per prima

Partite dalle abilities di sola lettura. Sono utili subito («quante persone aspettano la sciarpa blu?») e non possono fare danni. Lasciate da parte tutto ciò che invia email, cambia i prezzi o elimina dati finché non siete sicuri che i controlli dei permessi e le annotations siano corretti.

L'Abilities API richiede WordPress 6.9 o successivo. La [documentazione ufficiale](https://developer.wordpress.org/apis/abilities/) contiene il riferimento completo. Se volete che venga aggiunta a un plugin su cui fate affidamento, [contattateci](https://sawirstudio.com/it/contact/).

**Continuate a leggere:** [WordPress 7.2 arriva l'8 dicembre: cosa stiamo testando nei nostri plugin](https://sawirstudio.com/it/posts/wordpress-7-2-what-plugin-authors-should-test/), oppure [Cosa ci è voluto per rendere i nostri siti pronti per gli agenti](https://sawirstudio.com/it/posts/making-our-sites-agent-ready/).
