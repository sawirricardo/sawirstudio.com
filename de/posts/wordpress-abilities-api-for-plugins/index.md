---
title: WordPress Abilities API für Plugins | SawirStudio
url: https://sawirstudio.com/de/posts/wordpress-abilities-api-for-plugins/
description: So registrieren Sie Abilities in einem WordPress-Plugin, damit KI-Agenten und andere Tools sie finden und sicher ausführen können: Kategorien, Schemas, Berechtigungen, Annotations und die REST-Endpunkte.
---

# So machen Sie Ihr WordPress-Plugin mit der Abilities API für KI-Agenten nutzbar

Die Abilities API, die mit WordPress 6.9 eingeführt wurde, ist ein zentraler Ort, an dem ein Plugin festhält, was es kann: einen Namen, eine Beschreibung, welche Eingaben es erwartet, was es zurückgibt und wer es ausführen darf. KI-Agenten, das WordPress-AI-Plugin, MCP-Adapter und ganz normale Skripte können diese Abilities dann finden und nutzen, ohne dass jedes Tool jedes Plugin einzeln von Hand kennenlernen muss.

So funktioniert es, am Beispiel eines unserer [WordPress-Plugins](https://sawirstudio.com/de/projects/wordpress-plugins/).

## Das Beispiel

[Sawir Restock Notifier](https://sawirstudio.com/de/projects/wordpress-plugins/sawir-restock-notifier-for-woocommerce/) ermöglicht es Kunden, sich per E-Mail benachrichtigen zu lassen, sobald ein Produkt wieder auf Lager ist. Eine naheliegende Frage an einen Agenten lautet: „Bei welchen ausverkauften Produkten warten die meisten Leute auf eine Benachrichtigung?“ Das ist eine reine Leseabfrage (Read-only-Ability).

## 1. Eine Kategorie registrieren

Jede Ability gehört zu genau einer Kategorie. Kategorien werden zuerst registriert, und zwar über einen eigenen Hook:

```
add_action( 'wp_abilities_api_categories_init', function () {
	wp_register_ability_category( 'store-insights', array(
		'label'       => __( 'Store insights', 'sawir-restock-notifier-for-woocommerce' ),
		'description' => __( 'Read-only questions about the store.', 'sawir-restock-notifier-for-woocommerce' ),
	) );
} );
```

## 2. Die Ability registrieren

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

Der Callback ist ganz normaler Plugin-Code, der Daten passend zum Output-Schema zurückgibt. Daran ist nichts KI-spezifisch.

## Wofür die einzelnen Teile da sind

- **Name** in der Form `namespace/ability-name`. Verwenden Sie den Slug Ihres Plugins als Namespace.
- **Beschreibung.** Daran erkennt ein Agent, ob er die Ability nutzen soll. Schreiben Sie sie daher für jemanden, der Ihr Plugin noch nie gesehen hat.
- **Input- und Output-Schemas** in JSON Schema. Eingaben werden vor der Ausführung Ihres Codes dagegen geprüft, und Tools wissen, was sie zurückbekommen.
- **Permission-Callback.** Der einzige Teil, der tatsächlich etwas schützt. Prüfen Sie hier die Berechtigungen, so wie Sie es für eine Admin-Seite tun würden.
- **Annotations.** `readonly`, `destructive` und `idempotent` geben Tools einen Hinweis, wie vorsichtig sie sein sollten. Das sind nur Hinweise, keine Sicherheitsmaßnahme; die Sicherheit liefert der Permission-Callback.
- **`show_in_rest`.** Abilities sind in der REST API verborgen, solange Sie diese Option nicht aktivieren.

## Ausführung über REST

Abilities, die in REST sichtbar sind, werden unter `/wp-json/wp-abilities/v1/abilities` aufgelistet, und jede hat einen `/run`-Endpunkt. Eine Read-only-Ability wird per GET ausgeführt, wobei die Eingabe als URL-kodiertes JSON im Parameter `input` übergeben wird. Eine Ability, die etwas verändert, läuft per POST, eine destruktive per DELETE. Jede Anfrage erfordert einen authentifizierten Benutzer, danach entscheidet Ihr Permission-Callback.

## Und für Agenten über MCP

Der MCP Adapter, der zusammen mit dem WordPress-AI-Plugin entwickelt wird, stellt registrierte Abilities für MCP-Clients wie Claude bereit. Dafür müssen Sie nichts Zusätzliches schreiben: Eine gut beschriebene Ability mit sauberen Schemas ist bereits ein gutes Tool.

## Was wir zuerst bereitstellen würden

Beginnen Sie mit Read-only-Abilities. Sie sind sofort nützlich („Wie viele Kunden warten auf den blauen Schal?“) und können keinen Schaden anrichten. Abilities, die E-Mails versenden, Preise ändern oder etwas löschen, sollten Sie erst freigeben, wenn Sie sicher sind, dass Ihre Berechtigungsprüfungen und Annotations stimmen.

Die Abilities API benötigt WordPress 6.9 oder neuer. Die [offizielle Dokumentation](https://developer.wordpress.org/apis/abilities/) enthält die vollständige Referenz. Wenn Sie möchten, dass wir dies in ein Plugin einbauen, auf das Sie sich verlassen, [melden Sie sich bei uns](https://sawirstudio.com/de/contact/).

**Weiterlesen:** [WordPress 7.2 erscheint am 8. Dezember: Was wir in unseren Plugins testen](https://sawirstudio.com/de/posts/wordpress-7-2-what-plugin-authors-should-test/) oder [Was nötig war, um unsere Websites agentenfähig zu machen](https://sawirstudio.com/de/posts/making-our-sites-agent-ready/).
