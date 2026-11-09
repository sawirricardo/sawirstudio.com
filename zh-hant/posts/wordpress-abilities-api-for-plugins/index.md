---
title: 適用於外掛的 WordPress Abilities API | SawirStudio
url: https://sawirstudio.com/zh-hant/posts/wordpress-abilities-api-for-plugins/
description: 如何在 WordPress 外掛中註冊 abilities，讓 AI 代理與其他工具能安全地探索並執行這些功能：分類、schema、權限、註記與 REST 端點。
---

# 用 Abilities API 讓 AI 代理也能使用您的 WordPress 外掛

WordPress 6.9 新增的 Abilities API，讓外掛可以在同一個地方說明自己能做什麼：名稱、說明、接受的輸入、回傳的內容，以及誰有權限執行。AI 代理、WordPress AI 外掛、MCP 轉接器，甚至一般的腳本，都能找到這些功能並加以使用，不必為每個外掛逐一手動適配。

以下以我們的一款 [WordPress 外掛](https://sawirstudio.com/zh-hant/projects/wordpress-plugins/)為例，說明運作方式。

## 範例

[Sawir Restock Notifier](https://sawirstudio.com/zh-hant/projects/wordpress-plugins/sawir-restock-notifier-for-woocommerce/) 讓顧客可以登記，在商品重新到貨時收到電子郵件通知。代理很容易會遇到這樣的問題：「哪些缺貨商品等待的人最多？」這就是一個唯讀的 ability。

## 1. 註冊分類

每個 ability 都必須屬於且僅屬於一個分類，而分類要先註冊，使用它自己的 hook：

```
add_action( 'wp_abilities_api_categories_init', function () {
	wp_register_ability_category( 'store-insights', array(
		'label'       => __( 'Store insights', 'sawir-restock-notifier-for-woocommerce' ),
		'description' => __( 'Read-only questions about the store.', 'sawir-restock-notifier-for-woocommerce' ),
	) );
} );
```

## 2. 註冊 ability

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

callback 就是一般的外掛程式碼，回傳符合輸出 schema 的資料，完全不需要針對 AI 做任何特別處理。

## 各部分的用途

- **名稱**，格式為 `namespace/ability-name`。請用外掛的 slug 作為 namespace。
- **說明。** 代理會根據這段文字判斷是否要使用該 ability，所以請以從未看過您外掛的人為對象來撰寫。
- **輸入與輸出 schema**，使用 JSON Schema。輸入會在您的程式碼執行前先依 schema 驗證，工具也能事先知道會拿到什麼資料。
- **權限 callback。** 這是唯一真正發揮保護作用的部分。請像處理後台管理畫面一樣，在這裡檢查使用者權限。
- **註記（annotations）。** `readonly`、`destructive` 與 `idempotent` 用來告訴工具該多謹慎。它們只是提示，並非安全機制；真正的安全防線是權限 callback。
- **`show_in_rest`。** 除非您開啟這個選項，否則 ability 不會顯示在 REST API 中。

## 透過 REST 執行

顯示於 REST 的 abilities 會列在 `/wp-json/wp-abilities/v1/abilities`，每個 ability 都有自己的 `/run` 端點。唯讀的 ability 以 GET 執行，輸入內容以 URL 編碼的 JSON 放在 `input` 參數中；會變更資料的以 POST 執行，具破壞性的則以 DELETE 執行。每個請求都必須是已通過驗證的使用者，接著由您的權限 callback 決定是否放行。

## 也能透過 MCP 交給代理使用

與 WordPress AI 外掛一同開發的 MCP Adapter，會把已註冊的 abilities 提供給 Claude 等 MCP 用戶端。您不需要為此額外撰寫任何程式：說明清楚、schema 完整的 ability，本身就是一個好用的工具。

## 我們會優先開放哪些功能

建議從唯讀的 abilities 開始。它們馬上就能派上用場（例如「有多少顧客在等藍色圍巾到貨？」），而且不會造成損害。凡是會寄送電子郵件、變更價格或刪除資料的功能，請等到確定權限檢查與註記都設定正確之後再開放。

Abilities API 需要 WordPress 6.9 或更新版本。完整的參考資料請見[官方文件](https://developer.wordpress.org/apis/abilities/)。如果您希望在自己倚賴的外掛中加入這項功能，歡迎[與我們聯絡](https://sawirstudio.com/zh-hant/contact/)。

**延伸閱讀：** [WordPress 7.2 將於 12 月 8 日推出：我們在自家外掛中測試的項目](https://sawirstudio.com/zh-hant/posts/wordpress-7-2-what-plugin-authors-should-test/)，或[讓我們的網站準備好迎接 AI 代理，我們做了哪些事](https://sawirstudio.com/zh-hant/posts/making-our-sites-agent-ready/)。
