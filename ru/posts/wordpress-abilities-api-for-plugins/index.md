---
title: WordPress Abilities API для плагинов | SawirStudio
url: https://sawirstudio.com/ru/posts/wordpress-abilities-api-for-plugins/
description: Как зарегистрировать abilities в плагине для WordPress, чтобы ИИ-агенты и другие инструменты могли безопасно находить и запускать их: категории, схемы, права доступа, аннотации и REST endpoints.
---

# Как сделать ваш плагин для WordPress доступным для ИИ-агентов с помощью Abilities API

Abilities API, появившийся в WordPress 6.9, — это единое место, где плагин описывает, что он умеет: название, описание, какие данные принимает на вход, что возвращает и кому разрешено это запускать. ИИ-агенты, плагин WordPress AI, адаптеры MCP и обычные скрипты могут находить такие возможности и пользоваться ими, и каждому инструменту не нужно вручную разбираться в каждом плагине.

Покажем, как это работает, на примере одного из наших [плагинов для WordPress](https://sawirstudio.com/ru/projects/wordpress-plugins/).

## Пример

[Sawir Restock Notifier](https://sawirstudio.com/ru/projects/wordpress-plugins/sawir-restock-notifier-for-woocommerce/) позволяет покупателям попросить письмо, когда товар снова появится в наличии. Очевидный вопрос, на который может ответить агент: «у каких товаров, которых нет в наличии, больше всего ожидающих?» Это ability только для чтения.

## 1. Зарегистрируйте категорию

Каждая ability относится ровно к одной категории, а категории регистрируются первыми, на отдельном хуке:

```
add_action( 'wp_abilities_api_categories_init', function () {
	wp_register_ability_category( 'store-insights', array(
		'label'       => __( 'Store insights', 'sawir-restock-notifier-for-woocommerce' ),
		'description' => __( 'Read-only questions about the store.', 'sawir-restock-notifier-for-woocommerce' ),
	) );
} );
```

## 2. Зарегистрируйте ability

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

Callback — это обычный код плагина, который возвращает данные в соответствии со схемой вывода. Ничего специфичного для ИИ в нём нет.

## Для чего нужна каждая часть

- **Название** в формате `namespace/ability-name`. В качестве пространства имён используйте slug вашего плагина.
- **Описание.** Именно его читает агент, чтобы решить, использовать ли ability, поэтому пишите так, будто человек никогда не видел ваш плагин.
- **Схемы входных и выходных данных** в формате JSON Schema. Входные данные проверяются по ним до запуска вашего кода, а инструменты заранее знают, что получат в ответ.
- **Permission callback.** Единственная часть, которая реально что-то защищает. Проверяйте права здесь так же, как для экрана в админке.
- **Аннотации.** `readonly`, `destructive` и `idempotent` подсказывают инструментам, насколько осторожно действовать. Это только подсказки, а не защита: защитой служит permission callback.
- **`show_in_rest`.** Пока вы не включите этот параметр, abilities скрыты от REST API.

## Запуск через REST

Abilities, доступные в REST, перечислены по адресу `/wp-json/wp-abilities/v1/abilities`, и у каждой есть endpoint `/run`. Ability только для чтения запускается методом GET, а входные данные передаются в параметре `input` в виде JSON, закодированного для URL. Та, что что-то изменяет, запускается методом POST, а разрушающая — методом DELETE. Каждый запрос требует аутентифицированного пользователя, а дальше решение принимает ваш permission callback.

## И для агентов, через MCP

MCP Adapter, который разрабатывается вместе с плагином WordPress AI, предоставляет зарегистрированные abilities клиентам MCP, например Claude. Ничего дополнительного для этого писать не нужно: хорошо описанная ability с продуманными схемами — это уже хороший инструмент.

## Что мы открыли бы в первую очередь

Начните с abilities только для чтения. Они полезны сразу («сколько покупателей ждут синий шарф?») и не могут ничего испортить. Всё, что отправляет письма, меняет цены или удаляет данные, оставьте на потом, пока не убедитесь, что проверки прав и аннотации настроены правильно.

Для Abilities API нужен WordPress 6.9 или новее. Полный справочник есть в [официальной документации](https://developer.wordpress.org/apis/abilities/). Если хотите, чтобы это добавили в плагин, на который вы полагаетесь, [напишите нам](https://sawirstudio.com/ru/contact/).

**Читайте также:** [WordPress 7.2 выйдет 8 декабря: что мы тестируем в наших плагинах](https://sawirstudio.com/ru/posts/wordpress-7-2-what-plugin-authors-should-test/) или [Что потребовалось, чтобы подготовить наши сайты для работы с агентами](https://sawirstudio.com/ru/posts/making-our-sites-agent-ready/).
