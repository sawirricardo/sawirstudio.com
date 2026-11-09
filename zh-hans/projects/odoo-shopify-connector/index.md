---
title: Odoo 的 Shopify 连接器 | SawirStudio
url: https://sawirstudio.com/zh-hans/projects/odoo-shopify-connector/
description: 一个面向 Odoo 19 Community 的开源 Shopify 连接器，由 SawirStudio 编写，并在 Odoo Community Association（OCA）中维护。
---

# Odoo 的 Shopify 连接器

![GitHub 上的 OCA connector-shopify 代码仓库](https://sawirstudio.com/assets/images/projects/odoo-shopify-connector.webp)

## 项目简介

Odoo Community Association（OCA）是一家非营利组织，负责维护大多数 Odoo 社区版用户所依赖的开源模块。Connector Shopify 是该组织推出的连接器，用于连接 Shopify 店铺与 Odoo 19 Community，采用 AGPL-3 许可证发布。

## 我们做了什么

这个连接器由我们编写，并于 2026 年 7 月贡献给 OCA，包括核心的 Shopify Connector、Shopify Connector Accounting 和 Shopify Connector POS。我们仍在持续维护，最近一次是让它适配 Shopify 的 2026-01 Admin API，其他 OCA 贡献者也已在此基础上开发，其中包括一个电商产品媒体模块。

- **商务数据同步。** 支持商品、库存、客户、订单、退货、退款和发货同步，每个 Odoo 数据库可连接任意数量的 Shopify 店铺，并按公司划分范围。
- **可靠的设计。** 通过签名 Webhook 保证速度，借助 OCA 的 queue\_job 实现持久化处理，并通过定时对账修复遗漏或延迟的事件。
- **日常工具。** 包括初始设置引导、健康检查、同步仪表盘、错误恢复和可配置的数据保留期限。
- **会计核算。** 支持 Shopify Payments 的打款、余额交易、手续费、争议、礼品卡负债以及多币种对账。
- **销售点。** Shopify POS 的销售会作为销售订单导入，并从各个零售门店对应到 Odoo 仓库，无需使用 Odoo 自带的 POS 应用。

## 我们的做法

对会计来说，大部分时候正确的集成等于错误的集成。因此，这个连接器使用 Shopify 的 GraphQL Admin API，从不假定 Webhook 一定会送达，并会按计划定时与 Shopify 核对。日记账分录按 Shopify 的准确金额生成，过账前必须借贷平衡：如果数据对不上，连接器会给出明确的错误提示，而不是悄悄加一行来凑平。

需要为您的店铺进行设置或扩展？请查看 [Shopify 集成](https://sawirstudio.com/zh-hans/shopify/)；如果想了解 Odoo 本身，请查看 [Odoo 与业务系统](https://sawirstudio.com/zh-hans/odoo/)。

**技术栈** Python、Odoo 19 和 Shopify GraphQL Admin API

- **项目** Odoo Community Association（OCA）
- **许可证** AGPL-3
- **源代码** <https://github.com/OCA/connector-shopify>

## 更多项目

[![MyVivaldy](https://sawirstudio.com/assets/images/projects/myvivaldy.webp)

### MyVivaldy](https://sawirstudio.com/zh-hans/projects/myvivaldy/)
[![Koperasi HOB 应用](https://sawirstudio.com/assets/images/projects/koperasi-hob-application.webp)

### Koperasi HOB 应用](https://sawirstudio.com/zh-hans/projects/koperasi-hob-application/)
[![DBTV 应用](https://sawirstudio.com/assets/images/projects/dbtv-application.webp)

### DBTV 应用](https://sawirstudio.com/zh-hans/projects/dbtv-application/)

我们先聊到这里。

## 我们能为您 做些什么？

[开始项目](https://sawirstudio.com/zh-hans/contact/)
