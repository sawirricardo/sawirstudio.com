---
title: Sawir Renew – WooCommerce 周期性计费插件 | SawirStudio
url: https://sawirstudio.com/zh-hans/projects/wordpress-plugins/sawir-renew/
description: WooCommerce 完整订阅方案：周期性计费、试用、开通费、暂停/跳过、套餐切换、催款、智能重试，全部免费。
---

# Sawir Renew

![Sawir Renew 横幅图](https://sawirstudio.com/assets/images/plugins/sawir-renew-banner.webp)

Sawir Renew 能把 WooCommerce 商店变成订阅制生意：周期性计费、试用、开通费和智能重试一应俱全，没有任何功能被付费版本锁住。

[在 WordPress.org 免费获取](https://wordpress.org/plugins/sawir-renew/)

## 功能介绍

用 WooCommerce 销售订阅，无需另外购买扩展。所有功能均免费提供，没有专业版，也没有任何功能被锁定。

### 产品与定价

- 支持简单订阅商品和可变订阅商品两种类型
- 可按每 N 天、周、月或年计费，并支持促销价
- 免费试用、开通费，以及在 N 次付款后自动到期的订阅
- 统一续订日期（所有人都在每月 1 日续订，或在您指定的任意日期续订），首期可选按比例计费或免费
- 一次性运费，以及「每位客户仅限一个有效订阅」的限制

### 两种付款方式

- **按次续订付款（适用于所有支付网关）：** 续订到期时，客户会收到带有付款链接的账单，可通过银行转账、电子钱包、货到付款等您店铺支持的任何方式付款。非常适合不习惯自动扣款的市场。
- **通过 Stripe 自动扣款：** 如果您使用官方的 WooCommerce Stripe Gateway，客户在结账时即可保存银行卡，续订时系统会自动在后台扣款，包括智能重试，以及银行卡需要验证时的备用付款链接。其他支付网关也可以接入同一个适配器接口（见常见问题）。

### 续订不再半夜打扰您

- 所有定时任务都基于 Action Scheduler 运行，并每天进行一次自我修复检查
- 可配置的催款流程：重试间隔、提醒时间表、宽限期、自动取消
- 预发布/复制站点检测：一旦网站 URL 发生变化，就会立即暂停计费，不会再因为遗忘的预发布副本而重复扣款
- 自动扣款前发送续订提醒邮件

### 客户可以自助管理

在「我的账户 → 订阅」中，客户可以取消订阅（立即取消或在当期结束时取消，由您决定）、暂停和恢复、跳过一次续订、提前续订、更换付款方式或收货地址，还可以在不同套餐之间升级或降级，升级时按比例收取差价。

### 周期性优惠券

提供四种订阅优惠券：周期性百分比/固定金额折扣（可限定仅适用于前 N 次付款），以及开通费折扣。WooCommerce 自带的优惠券仍然只适用于首笔订单。

### 通知

共有十三种可开关的邮件（面向客户和管理员），覆盖订阅生命周期中的每个节点。开发者可以通过 `swsb_notified` 钩子，把任意事件同步到其他渠道（WhatsApp、短信、Slack 等）。

### 为新版 WooCommerce 而生

原生支持 HPOS（订阅直接存放在 WooCommerce 订单存储中，没有自定义数据表），兼容购物车/结账区块，后台列表和编辑页完全使用 WooCommerce 自带的订单界面，另有每小时缓存更新的报表页面（活跃订阅、MRR、流失率、即将到来的续订、续订收入）以及 REST API。

## 常见问题

### 这个插件与 WooCommerce 或 Automattic 有关联吗？

没有。「WooCommerce」是 Automattic 的商标；按照其命名规范，这是一个独立开发的、*面向* WooCommerce 的插件。

### 哪些支付网关支持自动续订？

开箱即用的是官方的 WooCommerce Stripe Gateway。其他所有网关都以按次续订付款的模式运行，即每个周期客户都会收到一条付款链接。开发者可以通过两个过滤器，为任意网关添加自动扣款支持：

`swsb_order_supports_auto`：当父订单所用的网关保存了可重复使用的付款方式时，返回 true。`swsb_charge_renewal`：在后台对续订订单扣款，返回 `true` 或 `WP_Error`。

### 可以和 WooCommerce Subscriptions 同时运行吗？

不可以。同时运行两个计费引擎会让客户被重复扣款，所以如果检测到 WooCommerce Subscriptions 处于启用状态，本插件会停用自己的计费引擎并显示提示。

### 在预发布站点上会发生什么？

如果网站 URL 发生变化，计费和客户邮件会自动暂停，直到您确认这个副本是预发布站点还是正式站点。选择「正式站点」后，所有计划任务都会重新创建。

## 推荐搭配

- [Sawir Referral](https://sawirstudio.com/zh-hans/projects/wordpress-plugins/sawir-referral/)：让您的客户带来他们的朋友。
- [Sawir Pulse](https://sawirstudio.com/zh-hans/projects/wordpress-plugins/sawir-pulse/)：以角落里的小通知形式，向顾客展示您店铺里真实发生的动态，例如最新订单、正在浏览商品的人数和最新评价。

查看我们所有的[免费 WordPress 插件](https://sawirstudio.com/zh-hans/projects/wordpress-plugins/)。

- **价格** 免费
- **版本** 1.0.0
- **运行要求** WordPress 6.0、PHP 7.4、WooCommerce
- **许可证** GPL-2.0-or-later
- **WordPress.org** <https://wordpress.org/plugins/sawir-renew/>

我们先聊到这里。

## 我们能为您 做些什么？

[开始项目](https://sawirstudio.com/zh-hans/contact/)
