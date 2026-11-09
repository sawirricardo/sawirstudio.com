---
title: 将 Shopify 连接到 Odoo 19 Community | SawirStudio
url: https://sawirstudio.com/zh-hans/posts/connect-shopify-to-odoo-19-community/
description: 如何使用开源的 OCA Shopify 连接器，将 Shopify 店铺连接到 Odoo 19 Community：Dev Dashboard 应用、权限范围、queue_job、webhook 和初始设置。
---

# 将 Shopify 连接到 Odoo 19 Community：分步指南

我们开发了 [Odoo 的 Shopify 连接器](https://sawirstudio.com/zh-hans/projects/odoo-shopify-connector/)，并将其贡献给了 Odoo Community Association（OCA）。它免费、开源（AGPL-3），支持 Odoo 19 Community。下面是我们为客户部署时从头到尾的设置流程。

它可以同步商品、库存、客户、订单、退货、退款和发货，支持单个或多个店铺。财务会计和销售终端（POS）是独立的模块，您可以之后再添加。

## 开始之前

- Odoo 19 Community 可通过公网 HTTPS 访问。Shopify 会向它发送 webhook，并从它下载商品图片。
- 拥有 Shopify 店铺的管理员权限，以及该店铺所属 Shopify 组织的管理员权限。
- 所用的 Shopify 套餐需支持自建应用访问受保护的客户数据。Basic 套餐不支持；没有这项权限，Shopify 会拒绝部分订单 webhook，订单和客户导入也会失败。

## 1. 在 Shopify Dev Dashboard 中创建应用

自 2026 年 1 月 1 日起，Shopify 不再允许在店铺后台创建旧版“自定义应用”。请在店铺所属组织的 **Dev Dashboard** 中创建应用。这一点很重要：Odoo 通过 client credentials grant 获取访问令牌，而这只有在应用和店铺属于同一个组织时才有效。

填写应用版本时：

- **App URL：** `https://shopify.dev/apps/default-app-home`。该连接器在 Shopify 内没有页面。
- **在 Shopify 后台中嵌入应用：** 关闭。
- **使用旧版安装流程：** 关闭，这样安装应用时就会授予您所选的权限范围。
- **Webhooks：** 不填。Odoo 会在初始设置过程中自行创建。
- **Webhooks API 版本：** `2026-01`，即连接器所使用的版本。
- **权限范围（Scopes）：** `read_customers`、`write_customers`、`read_draft_orders`、`read_fulfillments`、`read_inventory`、`write_inventory`、`read_locations`、`read_merchant_managed_fulfillment_orders`、`write_merchant_managed_fulfillment_orders`、`read_orders`、`read_products`、`write_products` 和 `read_returns`。如果需要导入 60 天以前的订单，请再添加 `read_all_orders`。

发布该版本，将应用安装到店铺，然后从应用的“设置”页面复制 client ID 和 client secret。

已有 2026 年之前创建的旧版自定义应用？它仍然可用。请将 client ID 留空，改用其 Admin API 访问令牌和 API 密钥（API secret key）。

## 2. 在 Odoo 中安装模块

将 OCA 的 `queue_job` 和 [connector-shopify](https://github.com/OCA/connector-shopify) 仓库添加到 addons 路径，然后安装 **Shopify Connector**。这些工作以后台任务的方式运行，因此 `queue_job` 必须以服务器级别加载，并为每个 Shopify 渠道配置相应的并发容量：

```
[options]
server_wide_modules = base,rpc,web,queue_job

[queue_job]
channels = root:4,root.shopify.import:2,root.shopify.export:1,root.shopify.webhook:1
```

然后将系统参数 `web.base.url` 设置为 Odoo 的公网 HTTPS 地址，并将 `web.base.url.freeze` 设置为 `True`。否则，当管理员从其他地方登录时，该地址会发生变化，Shopify 的 webhook 就会发到错误的地方。

## 3. 运行初始设置向导

为您的团队成员分配 **Shopify Manager** 用户组，然后打开 **Shopify > 配置 > 实例**，点击 **Start Onboarding**：

1. 输入 `yourstore.myshopify.com` 域名、client ID 和 client secret，并测试连接。Odoo 会自行请求令牌，并每 24 小时自动续期。
2. 检查权限范围。如果缺少某一项，请将其添加到新的应用版本中，发布后在店铺中批准。
3. 将 Shopify 的地点（location）映射到 Odoo 的仓库，检查 Odoo 生成的税项映射，并选择同步策略。
4. 安装 webhook，并将首次导入加入队列。

在让订单和发货无人值守地自动运行之前，请先检查支付网关和承运商的映射。后面出现的意外情况，大多是因为某个映射没有人检查过。

## 4. 持续关注运行情况

**Shopify > 仪表盘** 会显示同步状态、任务积压、近期错误和 webhook 健康状况。连接器不会假定每个 webhook 都能送达：定时任务会重新导入 Shopify 中已变更但 Odoo 遗漏的所有内容。出现失败时，先排除原因，再使用 **Retry Failed Shopify Bindings**。更换域名或凭据后，请运行 **Check Webhooks**。

## 后续步骤

添加 **Shopify Connector Accounting**，可将 Shopify Payments 的结算款、手续费和争议以日记账分录的形式导入（我们另外说明了[结算款为什么难以对账](https://sawirstudio.com/zh-hans/posts/why-shopify-payouts-dont-match-your-books/)）；添加 **Shopify Connector POS**，可处理门店销售。

如果您希望我们帮您完成部署，按您的业务方式进行配置或扩展，[我们可以做到](https://sawirstudio.com/zh-hans/shopify/)。

**继续阅读：** [为什么您的 Shopify 结算款总是与账目对不上](https://sawirstudio.com/zh-hans/posts/why-shopify-payouts-dont-match-your-books/)，或者 [Shopify 2026-10 API：您的集成会在哪些地方出问题](https://sawirstudio.com/zh-hans/posts/shopify-2026-10-api-what-breaks/)。
