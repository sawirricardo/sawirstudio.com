---
title: WordPress Abilities API for Plugins | SawirStudio
url: https://sawirstudio.com/posts/wordpress-abilities-api-for-plugins/
description: How to register abilities in a WordPress plugin so AI agents and other tools can discover and run them safely: categories, schemas, permissions, annotations and the REST endpoints.
---

# Make Your WordPress Plugin Usable by AI Agents with the Abilities API

The Abilities API, added in WordPress 6.9, is a single place where a plugin says what it can do: a name, a description, what input it takes, what it returns and who's allowed to run it. AI agents, the WordPress AI plugin, MCP adapters and plain old scripts can then find those abilities and use them, without each tool learning each plugin by hand.

Here's how it works, using one of our [WordPress plugins](https://sawirstudio.com/projects/wordpress-plugins/) as the example.

## The example

[Sawir Restock Notifier](https://sawirstudio.com/projects/wordpress-plugins/sawir-restock-notifier-for-woocommerce/) lets shoppers ask to be emailed when a product is back in stock. An obvious question for an agent to answer is “which out-of-stock products have the most people waiting?” That's a read-only ability.

## 1. Register a category

Every ability belongs to exactly one category, and categories are registered first, on their own hook:

```
add_action( 'wp_abilities_api_categories_init', function () {
	wp_register_ability_category( 'store-insights', array(
		'label'       => __( 'Store insights', 'sawir-restock-notifier-for-woocommerce' ),
		'description' => __( 'Read-only questions about the store.', 'sawir-restock-notifier-for-woocommerce' ),
	) );
} );
```

## 2. Register the ability

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

The callback is ordinary plugin code that returns data matching the output schema. Nothing about it is AI-specific.

## What each part is for

- **Name**, as `namespace/ability-name`. Use your plugin's slug as the namespace.
- **Description.** This is what an agent reads to decide whether to use the ability, so write it for someone who has never seen your plugin.
- **Input and output schemas**, in JSON Schema. Inputs are checked against them before your code runs, and tools know what they'll get back.
- **Permission callback.** The one part that actually protects anything. Check capabilities here as you would for an admin screen.
- **Annotations.** `readonly`, `destructive` and `idempotent` tell tools how careful to be. They're hints, not security; the permission callback is the security.
- **`show_in_rest`.** Abilities are hidden from the REST API unless you turn this on.

## Running it over REST

Abilities shown in REST are listed at `/wp-json/wp-abilities/v1/abilities`, and each one has a `/run` endpoint. A read-only ability runs with GET, with its input as URL-encoded JSON in an `input` parameter. One that changes things runs with POST, and a destructive one with DELETE. Every request needs an authenticated user, and then your permission callback decides.

## And to agents, through MCP

The MCP Adapter, developed alongside the WordPress AI plugin, exposes registered abilities to MCP clients such as Claude. You don't write anything extra for it: a well-described ability with good schemas is already a good tool.

## What we'd expose first

Start with read-only abilities. They're useful straight away (“how many shoppers are waiting for the blue scarf?”) and can't do damage. Leave anything that sends email, changes prices or deletes things until you're sure your permission checks and annotations are right.

The Abilities API needs WordPress 6.9 or later. The [official documentation](https://developer.wordpress.org/apis/abilities/) has the full reference. If you want this added to a plugin you rely on, [get in touch](https://sawirstudio.com/contact/).

**Keep reading:** [WordPress 7.2 Lands December 8: What We're Testing in Our Plugins](https://sawirstudio.com/posts/wordpress-7-2-what-plugin-authors-should-test/), or [What It Took to Make Our Sites Agent-Ready](https://sawirstudio.com/posts/making-our-sites-agent-ready/).
