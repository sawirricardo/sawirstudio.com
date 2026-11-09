---
title: WordPress 7.2：插件作者应该测试什么 | SawirStudio
url: https://sawirstudio.com/zh-hans/posts/wordpress-7-2-what-plugin-authors-should-test/
description: WordPress 7.2 将于 2026 年 12 月 8 日发布。从 Application Passwords、拟议中的 Secrets API 到 Interactivity API 路由器，插件作者需要关注哪些变化，以及现在如何开始测试。
---

# WordPress 7.2 将于 12 月 8 日发布：我们在自己的插件中测试什么

WordPress 7.2 计划于 2026 年 12 月 8 日发布。第一个候选版本（RC）预计在 11 月 17 日推出，届时功能列表将冻结，并发布 Field Guide（汇总所有开发者说明的清单）。在此之前，测试版是发现插件问题成本最低的地方。

我们在 WordPress.org 上维护着 [12 个免费插件](https://sawirstudio.com/zh-hans/projects/wordpress-plugins/)，所以这是我们自己的检查清单。下面有些内容在我们撰写时还只是提案，最终请以 Field Guide 为准。

## 现在如何测试

- **WP-CLI**：在测试站点上运行 `wp core update --version=7.2-beta4`，或使用当前最新的测试版或 RC。
- **WordPress Beta Tester 插件**：设置为“Bleeding edge”和“Beta/RC Only”。
- **WordPress Playground**：无需安装任何东西，直接在浏览器中点选体验。

切勿在线上站点测试。请开启 `WP_DEBUG` 和 `WP_DEBUG_LOG`，留意弃用提示；大多数早期警告都会出现在这里。

## 我们检查的内容

### Application Passwords

7.2 加强了 Application Passwords：能更好地识别本地环境和 HTTPS 配置，新增一个应用密码时会发送邮件，并更安全地处理默认角色。如果您的插件或客户的工具通过这种方式对 REST API 进行身份验证，请完整地跑一遍这个流程。

### Secrets API

这是核心提供的凭据加密存储方式，让插件不再把 API 密钥以明文选项保存。它曾被提议加入 7.2，管理界面则推迟到 7.3。如果您的插件会存储 API 密钥，请留意 Field Guide：它一旦发布，这些密钥就应该存放在那里。我们的免费插件目前都不存储第三方密钥，所以对我们来说，这主要和客户项目有关。

### Sudo 模式

计划中的功能：在执行高权限操作之前，要求已登录的用户再次确认密码。如果您的插件在自己的界面或通过 AJAX 执行这类操作，请测试核心在中途要求重新验证时会发生什么。

### Interactivity API 路由器

新的路由器用于监听客户端导航，取代已弃用的 `state.navigation.hasStarted` 和 `hasFinished` 这一对状态。另外还新增了 `data-wp-html` 指令，用于渲染可信的标记。

### 后台图标和通知

后台管理栏和菜单中的 Dashicons 正在被 SVG 图标取代，后台通知对屏幕阅读器的播报也有所改进。如果您的插件为菜单图标设置了样式，或在后台 CSS 中针对 Dashicon 类名，请在测试版中查看效果。

### DataViews 与站点编辑器

插件可以在站点编辑器中注册自己的界面，并从服务器端控制 DataViews 的字段和布局。这里不会有任何功能失效，但如果您的插件有列表界面，了解未来的方向会有帮助。

## 测试之后

请在真正测试过之后再提高 readme 中的“Tested up to”，不要提前。如果插件没有在最近的几个主要版本上测试过，WordPress.org 会向用户显示警告，而这个警告会减少安装量。我们的插件目前标注为“Tested up to: 7.1”，在 RC 上通过测试后会改为 7.2。

新的默认主题 Ipsum 也会随版本一起发布。如果您的插件会在前台输出任何内容，请在该主题下检查；新站点将是这个样子。

如果您有插件需要为 7.2 更新，或者有无人维护的定制插件，[我们可以帮您看看](https://sawirstudio.com/zh-hans/contact/)。

**继续阅读：** [用 Abilities API 让 AI 代理也能使用您的 WordPress 插件](https://sawirstudio.com/zh-hans/posts/wordpress-abilities-api-for-plugins/)，或者[印尼企业该选 WooCommerce 还是 Shopify？](https://sawirstudio.com/zh-hans/posts/woocommerce-or-shopify-indonesia/)
