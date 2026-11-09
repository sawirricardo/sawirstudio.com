---
title: API Abilities de WordPress pour les extensions | SawirStudio
url: https://sawirstudio.com/fr/posts/wordpress-abilities-api-for-plugins/
description: Comment enregistrer des fonctionnalités (abilities) dans une extension WordPress pour que les agents IA et d'autres outils puissent les découvrir et les exécuter en toute sécurité : catégories, schémas, permissions, annotations et points de terminaison REST.
---

# Rendez votre extension WordPress utilisable par des agents IA avec l'API Abilities

L'API Abilities, introduite dans WordPress 6.9, est un point unique où une extension déclare ce qu'elle sait faire : un nom, une description, les données qu'elle attend en entrée, ce qu'elle renvoie et qui a le droit de l'exécuter. Les agents IA, l'extension WordPress AI, les adaptateurs MCP et de simples scripts peuvent alors trouver ces fonctionnalités et les utiliser, sans que chaque outil doive apprendre chaque extension à la main.

Voici comment cela fonctionne, avec l'une de nos [extensions WordPress](https://sawirstudio.com/fr/projects/wordpress-plugins/) comme exemple.

## L'exemple

[Sawir Restock Notifier](https://sawirstudio.com/fr/projects/wordpress-plugins/sawir-restock-notifier-for-woocommerce/) permet aux clients de demander à être prévenus par e-mail lorsqu'un produit est de nouveau en stock. Une question évidente à poser à un agent est « quels produits en rupture de stock ont le plus de personnes en attente ? » Il s'agit d'une fonctionnalité en lecture seule.

## 1. Enregistrer une catégorie

Chaque fonctionnalité appartient à une seule catégorie, et les catégories sont enregistrées en premier, sur leur propre hook :

```
add_action( 'wp_abilities_api_categories_init', function () {
	wp_register_ability_category( 'store-insights', array(
		'label'       => __( 'Store insights', 'sawir-restock-notifier-for-woocommerce' ),
		'description' => __( 'Read-only questions about the store.', 'sawir-restock-notifier-for-woocommerce' ),
	) );
} );
```

## 2. Enregistrer la fonctionnalité

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

La fonction de rappel (callback) est du code d'extension ordinaire qui renvoie des données conformes au schéma de sortie. Rien n'y est propre à l'IA.

## À quoi sert chaque élément

- **Nom**, sous la forme `namespace/ability-name`. Utilisez le slug de votre extension comme espace de noms.
- **Description.** C'est ce que lit un agent pour décider s'il doit utiliser la fonctionnalité : rédigez-la pour quelqu'un qui n'a jamais vu votre extension.
- **Schémas d'entrée et de sortie**, en JSON Schema. Les données d'entrée sont vérifiées avant l'exécution de votre code, et les outils savent ce qu'ils vont recevoir en retour.
- **Callback de permission.** La seule partie qui protège réellement quelque chose. Vérifiez-y les droits comme vous le feriez pour un écran d'administration.
- **Annotations.** `readonly`, `destructive` et `idempotent` indiquent aux outils le niveau de prudence à adopter. Ce sont des indications, pas de la sécurité : la sécurité, c'est le callback de permission.
- **`show_in_rest`.** Les fonctionnalités sont masquées dans l'API REST tant que vous n'activez pas cette option.

## L'exécuter via REST

Les fonctionnalités exposées dans REST sont listées à l'adresse `/wp-json/wp-abilities/v1/abilities`, et chacune dispose d'un point de terminaison `/run`. Une fonctionnalité en lecture seule s'exécute avec GET, ses données d'entrée étant transmises sous forme de JSON encodé dans l'URL, dans un paramètre `input`. Celle qui modifie des données s'exécute avec POST, et une fonctionnalité destructive avec DELETE. Chaque requête nécessite un utilisateur authentifié, puis c'est votre callback de permission qui tranche.

## Et pour les agents, via MCP

L'adaptateur MCP (MCP Adapter), développé en parallèle de l'extension WordPress AI, expose les fonctionnalités enregistrées aux clients MCP tels que Claude. Vous n'avez rien de plus à écrire : une fonctionnalité bien décrite, avec de bons schémas, est déjà un bon outil.

## Ce que nous exposerions en premier

Commencez par des fonctionnalités en lecture seule. Elles sont utiles tout de suite (« combien de clients attendent l'écharpe bleue ? ») et ne peuvent rien endommager. Gardez pour plus tard tout ce qui envoie des e-mails, modifie des prix ou supprime des données, le temps d'être certain que vos contrôles de permission et vos annotations sont corrects.

L'API Abilities nécessite WordPress 6.9 ou une version ultérieure. La [documentation officielle](https://developer.wordpress.org/apis/abilities/) contient la référence complète. Si vous souhaitez ajouter cela à une extension sur laquelle vous comptez, [contactez-nous](https://sawirstudio.com/fr/contact/).

**À lire aussi :** [WordPress 7.2 arrive le 8 décembre : ce que nous testons dans nos extensions](https://sawirstudio.com/fr/posts/wordpress-7-2-what-plugin-authors-should-test/), ou [Ce qu'il nous a fallu pour rendre nos sites prêts pour les agents](https://sawirstudio.com/fr/posts/making-our-sites-agent-ready/).
