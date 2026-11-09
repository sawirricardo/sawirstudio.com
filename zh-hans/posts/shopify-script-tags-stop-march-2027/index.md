---
title: Shopify Script Tags 将于 2027 年 3 月停用 | SawirStudio
url: https://sawirstudio.com/zh-hans/posts/shopify-script-tags-stop-march-2027/
description: Shopify 将于 2027 年 3 月 1 日停止运行在线商店的 script tag。本文介绍受影响的应用、如何检查您的店铺，以及用什么来替代。
---

# Shopify Script Tags 将于 2027 年 3 月 1 日停止工作：店主该怎么做

2027 年 3 月 1 日起，Shopify 将不再向在线商店前台注入 script tag。店铺里仍以这种方式加载的内容，比如评价插件、在线客服气泡、跟踪像素，都会直接不再显示。不会有明显的报错，只是东西悄悄不见了。

Shopify 已在[2026 年 8 月](https://shopify.dev/changelog/posts/online-store-script-tags-deprecation)宣布了这项调整。如果您经营着一家店铺，下面是它对您的影响。

## 什么是 script tag

script tag 是应用通过 API 向店铺前台添加 JavaScript 的一种方式，不需要改动您的主题。应用开发者喜欢它正是因为这一点，问题也出在这里：script tag 不会出现在主题编辑器中，所以您无法在那里看到或关闭它们。已卸载的应用有时还会留下自己的 script tag。

## 时间表

- **2026 年 10 月 1 日：** 无论使用哪个 API 版本，应用都不能再创建或更新 script tag。已有的仍会继续运行。
- **2027 年 3 月 1 日：** Shopify 停止加载在线商店的 script tag。

订单状态页上的 script tag 已更早按另一套时间表停用，所以如果那里已经有内容不再显示，原因就在这里。

## 如何检查您的店铺

Shopify 不会向商家提供 script tag 的清单，所以需要自己动手查：

1. **列出在店铺前台添加了内容的应用。** 包括评价、聊天、弹窗、追加销售、会员积分、数据分析和广告像素。只在后台管理中运行的应用不受影响。
2. **查看主题编辑器中的“应用嵌入”。** 如果某个应用出现在那里并且已开启，说明它使用的是较新的方式，没有问题。
3. **查看页面源代码。** 打开您的店铺，查看源代码并搜索 `asyncLoad`。在大多数店铺中，script tag 的网址都列在那一段代码里，其中每个域名通常对应一个应用。
4. **询问应用开发者。**“您的应用是否使用 script tag？2027 年 3 月 1 日之后还能正常使用吗？”对于一个持续维护的应用，一句话就能回答这个问题。

多年来，Shopify 一直在推动应用改用主题应用扩展，所以大多数维护良好的应用已经没有问题。风险主要在于老旧的应用、多年前有人为您定制的应用，以及早已不用的应用留下的残留。

## 用什么替代 script tag

- **应用嵌入块**（App embed blocks）属于主题应用扩展的一部分，适用于所有显示在页面上的内容。您可以在主题编辑器中开启它们，这意味着您终于能看到它们了。
- **Web pixels** 适用于只收集分析数据或转化数据的脚本。它们在沙盒中运行，只要应用支持，就不需要您做任何操作。

## 如果某个应用不会再更新

请在 3 月之前替换掉它，或者重新开发它的功能。对于定制应用，通常的做法是把脚本迁移到主题应用扩展中；如果只是做跟踪，则迁移到 web pixel。顺便也把已卸载应用留下的 script tag 删除掉，`scriptTagDelete` mutation 仍然可用。

如果您有需要迁移的 Shopify 定制应用，或者不确定店铺里到底加载了什么，请[把您的店铺地址发给我们](https://sawirstudio.com/zh-hans/contact/)，我们会告诉您查到的结果。关于我们的 Shopify 业务，更多内容请看[这个页面](https://sawirstudio.com/zh-hans/shopify/)。

**继续阅读：** [Shopify 2026-10 API：您的集成中哪些会出问题](https://sawirstudio.com/zh-hans/posts/shopify-2026-10-api-what-breaks/)，或者[将 Shopify 连接到 Odoo 19 Community 的分步指南](https://sawirstudio.com/zh-hans/posts/connect-shopify-to-odoo-19-community/)。
