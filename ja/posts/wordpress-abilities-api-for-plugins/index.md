---
title: プラグイン向け WordPress Abilities API | SawirStudio
url: https://sawirstudio.com/ja/posts/wordpress-abilities-api-for-plugins/
description: AI エージェントなどのツールが WordPress プラグインの機能を見つけて安全に実行できるようにするための、Ability の登録方法を解説します。カテゴリー、スキーマ、パーミッション、アノテーション、REST エンドポイントを取り上げます。
---

# Abilities API で WordPress プラグインを AI エージェントから使えるようにする

WordPress 6.9 で追加された Abilities API は、プラグインが「何ができるか」を一か所にまとめて宣言するための仕組みです。名前、説明、受け取る入力、返す結果、実行を許可する人を登録しておくと、AI エージェント、WordPress の AI プラグイン、MCP アダプター、従来のスクリプトがそれらの機能を見つけて利用できます。ツールごとにプラグインの使い方を手作業で覚えさせる必要はありません。

ここでは、当社の [WordPress プラグイン](https://sawirstudio.com/ja/projects/wordpress-plugins/)の一つを例に、仕組みをご紹介します。

## 例として使うプラグイン

[Sawir Restock Notifier](https://sawirstudio.com/ja/projects/wordpress-plugins/sawir-restock-notifier-for-woocommerce/) は、商品が再入荷したときにメールで知らせてほしいという購入者のリクエストを受け付けるプラグインです。エージェントに尋ねたい質問としてまず思い浮かぶのは、「在庫切れの商品のうち、待っている人が最も多いのはどれか」でしょう。これは読み取り専用の Ability になります。

## 1. カテゴリーを登録する

すべての Ability は必ず一つのカテゴリーに属します。カテゴリーは、専用のフックで先に登録します。

```
add_action( 'wp_abilities_api_categories_init', function () {
	wp_register_ability_category( 'store-insights', array(
		'label'       => __( 'Store insights', 'sawir-restock-notifier-for-woocommerce' ),
		'description' => __( 'Read-only questions about the store.', 'sawir-restock-notifier-for-woocommerce' ),
	) );
} );
```

## 2. Ability を登録する

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

コールバックは、出力スキーマに合ったデータを返す、ごく普通のプラグインのコードです。AI に特化した要素はありません。

## 各項目の役割

- **名前** は `namespace/ability-name` の形式にします。名前空間には、プラグインのスラッグを使います。
- **説明。** エージェントはこの説明を読んで、その Ability を使うかどうかを判断します。プラグインを一度も見たことがない人に向けて書いてください。
- **入力スキーマと出力スキーマ** は JSON Schema で記述します。入力はコードが実行される前にスキーマで検証され、ツール側も返ってくるデータの形を把握できます。
- **パーミッションコールバック。** 実際に保護の役割を果たすのは、この部分だけです。管理画面を作るときと同じように、ここで権限を確認してください。
- **アノテーション。** `readonly`、`destructive`、`idempotent` は、ツールにどの程度慎重に扱うべきかを伝えるものです。あくまでヒントであり、セキュリティ対策ではありません。セキュリティを担うのはパーミッションコールバックです。
- **`show_in_rest`。** これを有効にしない限り、Ability は REST API に公開されません。

## REST 経由で実行する

REST に公開された Ability は `/wp-json/wp-abilities/v1/abilities` で一覧表示され、それぞれに `/run` エンドポイントがあります。読み取り専用の Ability は GET で実行し、入力は URL エンコードした JSON として `input` パラメーターに渡します。データを変更する Ability は POST、破壊的な Ability は DELETE で実行します。どのリクエストも認証済みのユーザーが必要で、そのうえでパーミッションコールバックが可否を判断します。

## MCP を通じてエージェントから使う

WordPress の AI プラグインと並行して開発されている MCP Adapter は、登録済みの Ability を Claude などの MCP クライアントに公開します。そのために追加で書くものはありません。説明が丁寧で、スキーマがきちんと定義された Ability は、そのまま使いやすいツールになります。

## 最初に公開するもの

まずは読み取り専用の Ability から始めましょう。すぐに役立ちますし（「青いスカーフの再入荷を待っている購入者は何人か」など）、被害が出ることもありません。メールの送信、価格の変更、データの削除を伴うものは、パーミッションチェックとアノテーションが正しいと確信できるまで公開しないでください。

Abilities API には WordPress 6.9 以降が必要です。詳しいリファレンスは[公式ドキュメント](https://developer.wordpress.org/apis/abilities/)をご覧ください。お使いのプラグインに対応を追加したい場合は、[お問い合わせ](https://sawirstudio.com/ja/contact/)ください。

**あわせてお読みください：** [WordPress 7.2 は 12 月 8 日にリリース：当社プラグインで検証していること](https://sawirstudio.com/ja/posts/wordpress-7-2-what-plugin-authors-should-test/)、[当社サイトをエージェント対応にするまでにしたこと](https://sawirstudio.com/ja/posts/making-our-sites-agent-ready/)
