---
title: WordPress Abilities API voor plugins | SawirStudio
url: https://sawirstudio.com/nl/posts/wordpress-abilities-api-for-plugins/
description: Zo registreer je abilities in een WordPress-plugin, zodat AI-agents en andere tools ze veilig kunnen vinden en uitvoeren: categorieën, schema's, rechten, annotations en de REST-endpoints.
---

# Maak je WordPress-plugin bruikbaar voor AI-agents met de Abilities API

De Abilities API, toegevoegd in WordPress 6.9, is één plek waar een plugin vastlegt wat hij kan: een naam, een beschrijving, welke input hij nodig heeft, wat hij teruggeeft en wie hem mag uitvoeren. AI-agents, de WordPress AI-plugin, MCP-adapters en gewone scripts kunnen die abilities dan vinden en gebruiken, zonder dat elke tool elke plugin met de hand hoeft te leren kennen.

Zo werkt het, met een van onze [WordPress-plugins](https://sawirstudio.com/nl/projects/wordpress-plugins/) als voorbeeld.

## Het voorbeeld

[Sawir Restock Notifier](https://sawirstudio.com/nl/projects/wordpress-plugins/sawir-restock-notifier-for-woocommerce/) laat shoppers vragen om een e-mail te krijgen zodra een product weer op voorraad is. Een vraag die een agent dan logisch kan beantwoorden is: “welke producten die niet op voorraad zijn, hebben de meeste wachtenden?” Dat is een ability die alleen leest.

## 1. Registreer een categorie

Elke ability hoort bij precies één categorie. Categorieën registreer je eerst, op hun eigen hook:

```
add_action( 'wp_abilities_api_categories_init', function () {
	wp_register_ability_category( 'store-insights', array(
		'label'       => __( 'Store insights', 'sawir-restock-notifier-for-woocommerce' ),
		'description' => __( 'Read-only questions about the store.', 'sawir-restock-notifier-for-woocommerce' ),
	) );
} );
```

## 2. Registreer de ability

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

De callback is gewone plugincode die data teruggeeft die past bij het output-schema. Daar is niets AI-specifieks aan.

## Waar elk onderdeel voor dient

- **Naam**, in de vorm `namespace/ability-name`. Gebruik de slug van je plugin als namespace.
- **Beschrijving.** Hieraan beslist een agent of hij de ability gebruikt, dus schrijf hem voor iemand die je plugin nog nooit heeft gezien.
- **Input- en output-schema's**, in JSON Schema. Input wordt hiertegen gecontroleerd voordat jouw code draait, en tools weten wat ze terugkrijgen.
- **Permission callback.** Het enige onderdeel dat echt iets beschermt. Controleer hier de rechten zoals je dat ook voor een beheerscherm doet.
- **Annotations.** `readonly`, `destructive` en `idempotent` geven tools aan hoe voorzichtig ze moeten zijn. Het zijn hints, geen beveiliging; de permission callback is de beveiliging.
- **`show_in_rest`.** Abilities zijn verborgen in de REST API, tenzij je dit aanzet.

## Uitvoeren via REST

Abilities die in REST zichtbaar zijn, staan in een lijst op `/wp-json/wp-abilities/v1/abilities`, en elke ability heeft een `/run`-endpoint. Een ability die alleen leest, voer je uit met GET, met de input als URL-encoded JSON in een `input`-parameter. Een ability die iets wijzigt, voer je uit met POST en een destructieve met DELETE. Elke aanvraag heeft een ingelogde gebruiker nodig, en daarna beslist jouw permission callback.

## En naar agents, via MCP

De MCP Adapter, die samen met de WordPress AI-plugin wordt ontwikkeld, stelt geregistreerde abilities beschikbaar voor MCP-clients zoals Claude. Daar hoef je niets extra's voor te schrijven: een goed beschreven ability met goede schema's is al een goede tool.

## Wat we als eerste zouden aanbieden

Begin met abilities die alleen lezen. Ze zijn meteen nuttig (“hoeveel shoppers wachten op de blauwe sjaal?”) en kunnen geen schade aanrichten. Laat alles wat e-mail verstuurt, prijzen wijzigt of dingen verwijdert nog even liggen, tot je zeker weet dat je permission checks en annotations kloppen.

De Abilities API vraagt WordPress 6.9 of nieuwer. In de [officiële documentatie](https://developer.wordpress.org/apis/abilities/) vind je de volledige referentie. Wil je dit laten toevoegen aan een plugin waar je op vertrouwt? [Neem contact met ons op](https://sawirstudio.com/nl/contact/).

**Lees ook:** [WordPress 7.2 verschijnt op 8 december: dit testen we in onze plugins](https://sawirstudio.com/nl/posts/wordpress-7-2-what-plugin-authors-should-test/), of [Wat er nodig was om onze sites agent-ready te maken](https://sawirstudio.com/nl/posts/making-our-sites-agent-ready/).
