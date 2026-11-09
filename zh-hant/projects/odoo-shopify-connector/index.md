---
title: Odoo 的 Shopify 連接器 | SawirStudio
url: https://sawirstudio.com/zh-hant/projects/odoo-shopify-connector/
description: 適用於 Odoo 19 Community 的開源 Shopify 連接器，由 SawirStudio 開發，並由 Odoo Community Association（OCA）維護。
---

# Shopify Connector for Odoo

![GitHub 上的 OCA connector-shopify 儲存庫](https://sawirstudio.com/assets/images/projects/odoo-shopify-connector.webp)

## 專案簡介

Odoo Community Association（OCA）是一個非營利組織，負責維護多數 Odoo Community 使用者所仰賴的開源模組。Connector Shopify 是 OCA 提供的連接器，可串接 Shopify 商店與 Odoo 19 Community，並以 AGPL-3 授權釋出。

## 我們的作品

這套連接器由我們撰寫，並於 2026 年 7 月貢獻給 OCA，包含核心的 Shopify Connector、Shopify Connector Accounting 與 Shopify Connector POS。我們持續維護，最近一次是配合 Shopify 的 2026-01 Admin API 進行調整，目前也有其他 OCA 貢獻者在此基礎上開發，其中包括電商商品媒體模組。

- **商務同步。** 支援商品、庫存、客戶、訂單、退貨、退款與出貨，每個 Odoo 資料庫可連接任意數量的 Shopify 商店，並依公司區分。
- **可靠的設計。** 以簽章 webhook 確保速度，以 OCA 的 queue\_job 進行穩定的背景處理，並透過排程對帳，修復遺漏或延遲的事件。
- **日常管理工具。** 包含初始設定引導、健康檢查、同步儀表板、錯誤復原，以及可自訂的資料保留設定。
- **會計。** 支援 Shopify Payments 撥款、餘額交易、手續費、爭議款項、禮品卡負債與多幣別對帳。
- **門市銷售（POS）。** 將 Shopify POS 的銷售匯入為銷售訂單，並依各零售據點對應到 Odoo 倉庫，不需要使用 Odoo 內建的 POS 應用程式。

## 我們的做法

對會計來說，多數時候正確的整合仍然是錯誤的。因此這套連接器透過 Shopify 的 GraphQL Admin API 溝通，不會假設 webhook 一定送達，並會定期與 Shopify 比對自我檢查。會計分錄依據 Shopify 的精確金額規劃，且必須借貸平衡才能過帳：若金額對不上，連接器會明確報錯，而不是悄悄補上一行平衡分錄。

需要為您的商店進行設定或擴充功能嗎？請參考 [Shopify 整合](https://sawirstudio.com/zh-hant/shopify/)，若是 Odoo 本身的需求，請看 [Odoo 與企業系統](https://sawirstudio.com/zh-hant/odoo/)。

**使用技術** Python、Odoo 19 與 Shopify GraphQL Admin API

- **專案** Odoo Community Association（OCA）
- **授權** AGPL-3
- **原始碼** <https://github.com/OCA/connector-shopify>

## 更多專案

[![MyVivaldy](https://sawirstudio.com/assets/images/projects/myvivaldy.webp)

### MyVivaldy](https://sawirstudio.com/zh-hant/projects/myvivaldy/)
[![Koperasi HOB Application](https://sawirstudio.com/assets/images/projects/koperasi-hob-application.webp)

### Koperasi HOB Application](https://sawirstudio.com/zh-hant/projects/koperasi-hob-application/)
[![DBTV 應用程式](https://sawirstudio.com/assets/images/projects/dbtv-application.webp)

### DBTV 應用程式](https://sawirstudio.com/zh-hant/projects/dbtv-application/)

關於我們就先說到這裡。

## 有什麼我們 能幫上忙的嗎？

[開始專案](https://sawirstudio.com/zh-hant/contact/)
