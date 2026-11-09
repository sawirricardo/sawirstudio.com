---
title: Shopify 2026-10 API：哪些地方会出问题 | SawirStudio
url: https://sawirstudio.com/zh-hans/posts/shopify-2026-10-api-what-breaks/
description: Shopify Admin API 2026-10 版本中可能导致集成出问题的变更，以及应对方法：配送运费、草稿订单折扣、税费、条形码等。
---

# Shopify 2026-10 API：您的集成中哪些地方会出问题

Shopify 于 2026 年 10 月 1 日发布了 API 的 2026-10 版本。其中大部分是新功能，等需要时再看也不迟。但有少数变更，可能会悄悄搞坏一个在 2026-07 上运行正常的集成。在升级版本之前，值得花一个小时逐项检查这些变更。

这是我们用来检查自家集成的清单，其中包括 [Odoo 的 Shopify 连接器](https://sawirstudio.com/zh-hans/projects/odoo-shopify-connector/)。清单并未涵盖发布说明的每一行，只列出会改变行为的那些变更。

## 会悄悄造成问题的变更

### 新建的配送服务不会自动在结账时显示运费

使用 `carrierServiceCreate`（或 REST 中的 `POST carrier_services.json`）创建配送服务时，现在只会注册该服务，不再把它的运费加入通用配送方案（General shipping profile）。如果您的应用或初始化脚本创建了配送服务，并期望顾客能看到您的运费，那么在您亲自把该服务添加到某个配送方案之前，顾客是看不到的。

### 更改收货地址会重新计算税费

使用 `orderUpdate` 或 REST 订单更新接口修改未发货订单的收货地址时，现在会重新计算该订单的税费。这样做是对的，但如果您从 ERP 或客服工具中修改地址，订单总额可能会在您不知情的情况下发生变化。更新之后，请重新读取 `taxLines`、`totalTaxSet` 以及各项合计金额，并处理已付款或已退款金额上的差额。

### 草稿订单的折扣警告不再包含 `priceRule`

`DraftOrderDiscountNotAppliedWarning.priceRule` 和旧版 `PriceRule` 类型已被移除。如果您使用 `draftOrderCalculate`、`draftOrderCreate` 或 `draftOrderUpdate` 创建草稿订单，并读取折扣未生效的原因，请改为选取 `discountTitle` 和 `discountCode`。

### 错误的 metafield 筛选现在会明确报错

以前，如果产品查询按某个不可筛选的 metafield 过滤，系统会忽略这部分筛选条件并返回全部结果。现在则会返回错误。这样更好，但一个几个月来一直返回“所有产品”的同步任务，会在您切换版本的当天开始失败。请确认您用来筛选的 metafield 定义支持筛选，并且支持您所用的运算符。

### 一个变体可以有多个条形码

`ProductVariant.barcodes` 是新增的连接（connection），最多可包含 20 个带类型的条形码。旧的 `ProductVariant.barcode` 已被弃用，只会返回第一个条形码。如果您的仓库系统或销售终端（POS）通过条形码匹配产品，请读取完整列表。

### 一个库存错误码已被移除

`ITEM_NOT_STOCKED_AT_LOCATION` 已从库存 mutation 以及 `InventorySetQuantitiesUserErrorCode` 中移除。依赖它做分支判断的代码现在已成为无效代码，请将其删除，以免日后有人误以为它还有用。

## 值得在代码里搜索一遍的小变更

- **客户分群。** 用 `MATCHES` 和 `NOT MATCHES` 取代 `= true` 和 `= false`；`12_months_ago` 这类命名日期已被弃用，请改用 `-12m` 这样的偏移量。
- **短信营销授权。** 扁平的短信授权字段已被弃用，请读取 `CustomerPhoneNumber.smsMarketingConsent`。
- **分阶段发布中的折扣。** 当折扣属于某个分阶段发布（rollout）时，仅凭 `startsAt` 和 `endsAt` 无法判断它是否已生效。读取 `Discount.rollouts` 需要 `read_rollouts` 权限范围。
- **履约状态。** `FULFILLMENT_NOT_REQUIRED` 是新增的取值。如果您用 `switch` 映射所有状态，这个新值会落入默认分支。
- **Customer Account API。** `Customer.lastIncompleteCheckout` 和 `Checkout` 类型已被移除，没有替代方案。
- **POS UI 扩展。** 像 `applyCartDiscount` 这样的购物车 mutation，在 POS 未确认更改时现在会被拒绝；`session.currentSession.staffMemberId` 已被移除。

## 与 2026-10 无关的两个截止日期

如果您维护的是公开应用，则必须在 2027 年 1 月 1 日之前，把不会过期的离线访问令牌替换为会过期的令牌。自定义应用，以及商家在 Dev Dashboard 中创建的应用不受此限。此外，在线商店的 script tag 将于 2027 年 3 月 1 日起停止运行；关于[这对商店意味着什么](https://sawirstudio.com/zh-hans/posts/shopify-script-tags-stop-march-2027/)，我们另外写了一篇文章。

## 我们如何升级

我们把 API 版本固定在同一个地方，在开发店铺上用新版本运行集成，然后阅读按“Action required”筛选的 [2026-10 更新日志](https://shopify.dev/changelog?api_version=2026-10)。2026-10 版本至少会支持到 2027 年 10 月 1 日，所以时间还算充裕，但也不宜拖到收到警告邮件才动手。

如果您的店铺需要与 Odoo 或其他系统对接，又希望交给别人来做，这正是[我们的工作](https://sawirstudio.com/zh-hans/shopify/)。

**继续阅读：** [Shopify Script Tags 将于 2027 年 3 月 1 日停止运行，店主该怎么办](https://sawirstudio.com/zh-hans/posts/shopify-script-tags-stop-march-2027/)，或者[一步步将 Shopify 连接到 Odoo 19 Community](https://sawirstudio.com/zh-hans/posts/connect-shopify-to-odoo-19-community/)。
