---
title: WordPress Abilities API for Plugins | SawirStudio
url: https://sawirstudio.com/zh-hans/posts/wordpress-abilities-api-for-plugins/
description: 如何在 WordPress 插件中注册 ability，让 AI 代理和其他工具能够安全地发现并运行它们：分类、schema、权限、注解和 REST 端点。
---

# 借助 Abilities API，让 AI 代理也能使用您的 WordPress 插件

Abilities API 自 WordPress 6.9 起加入，它为插件提供了一个统一的地方来声明自己能做什么：名称、说明、接受什么输入、返回什么结果，以及谁有权运行。这样，AI 代理、WordPress AI 插件、MCP 适配器乃至普通脚本都可以发现并使用这些功能，而无需每个工具都针对每个插件逐一手动适配。

下面以我们的一款 [WordPress 插件](https://sawirstudio.com/zh-hans/projects/wordpress-plugins/)为例，介绍它是如何运作的。

## 示例

[Sawir Restock Notifier](https://sawirstudio.com/zh-hans/projects/wordpress-plugins/sawir-restock-notifier-for-woocommerce/) 允许顾客申请在商品到货时收到电子邮件通知。代理很容易遇到的一个问题是：“哪些缺货商品等待到货的人最多？”这就是一个只读的 ability。

## 1. 注册分类

每个 ability 必须且只能属于一个分类，而分类需要先在各自的钩子上注册：

```
add_action( 'wp_abilities_api_categories_init', function () {
	wp_register_ability_category( 'store-insights', array(
		'label'       => __( 'Store insights', 'sawir-restock-notifier-for-woocommerce' ),
		'description' => __( 'Read-only questions about the store.', 'sawir-restock-notifier-for-woocommerce' ),
	) );
} );
```

## 2. 注册 ability

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

回调函数就是普通的插件代码，返回符合输出 schema 的数据。其中没有任何内容是专为 AI 设计的。

## 各部分的作用

- **名称**，格式为 `namespace/ability-name`。请使用您插件的 slug 作为命名空间。
- **说明。** 代理会根据它来判断是否使用这个 ability，所以请站在从未见过您插件的人的角度来写。
- **输入和输出 schema**，使用 JSON Schema 编写。输入会在您的代码运行之前按 schema 进行校验，各类工具也能提前知道会得到什么返回结果。
- **权限回调。** 真正起保护作用的只有这一部分。请像为管理后台页面做检查一样，在这里检查用户权限。
- **注解。** `readonly`、`destructive` 和 `idempotent` 告诉工具应当多谨慎。它们只是提示，不是安全措施；真正的安全靠权限回调。
- **`show_in_rest`。** 除非您开启它，否则 ability 不会在 REST API 中显示。

## 通过 REST 运行

在 REST 中显示的 ability 会列在 `/wp-json/wp-abilities/v1/abilities`，每个 ability 都有一个 `/run` 端点。只读的 ability 使用 GET 运行，输入以 URL 编码的 JSON 形式放在 `input` 参数中；会修改数据的使用 POST 运行，具有破坏性的使用 DELETE 运行。每个请求都需要已认证的用户，之后由您的权限回调来决定是否放行。

## 再通过 MCP 提供给代理

与 WordPress AI 插件一同开发的 MCP Adapter 会把已注册的 ability 开放给 Claude 等 MCP 客户端。您无需为此额外编写任何代码：说明清楚、schema 完善的 ability，本身就是一个好用的工具。

## 我们建议优先开放什么

先从只读的 ability 做起。它们马上就能派上用场（“有多少顾客在等那条蓝色围巾到货？”），而且不会造成破坏。凡是会发送邮件、修改价格或删除数据的功能，请等到您确认权限检查和注解都没有问题之后再开放。

Abilities API 需要 WordPress 6.9 或更高版本。完整的参考资料请见[官方文档](https://developer.wordpress.org/apis/abilities/)。如果您希望为自己依赖的某个插件加入这项功能，欢迎[联系我们](https://sawirstudio.com/zh-hans/contact/)。

**继续阅读：** [WordPress 7.2 将于 12 月 8 日发布：我们在自家插件中测试了什么](https://sawirstudio.com/zh-hans/posts/wordpress-7-2-what-plugin-authors-should-test/)，或者[让我们的网站做好迎接 AI 代理准备的全过程](https://sawirstudio.com/zh-hans/posts/making-our-sites-agent-ready/)。
