---
title: 將 Shopify 連接到 Odoo 19 Community | SawirStudio
url: https://sawirstudio.com/zh-hant/posts/connect-shopify-to-odoo-19-community/
description: 如何使用開源的 OCA Shopify 連接器，將 Shopify 商店連接到 Odoo 19 Community：Dev Dashboard 應用程式、scopes、queue_job、webhook 與初始設定。
---

# 將 Shopify 連接到 Odoo 19 Community 的逐步設定教學

我們開發了 [Odoo 的 Shopify 連接器](https://sawirstudio.com/zh-hant/projects/odoo-shopify-connector/)，並貢獻給 Odoo Community Association（OCA）。它免費、開源（AGPL-3），可搭配 Odoo 19 Community 使用。以下是我們為客戶設定時的完整流程，從頭到尾一步步說明。

它可以同步商品、庫存、客戶、訂單、退貨、退款與出貨，單一商店或多間商店都適用。會計與銷售點（POS）則是獨立的模組，之後可以再加裝。

## 開始之前

- Odoo 19 Community 必須能透過公開的 HTTPS 連線存取。Shopify 會將 webhook 傳送到這裡，也會從這裡下載商品圖片。
- 擁有 Shopify 商店的管理員權限，以及該商店所屬 Shopify 組織的管理員權限。
- 您的 Shopify 方案需包含自行建立之應用程式可使用的受保護客戶資料。Basic 方案不包含此項目，缺少的話，Shopify 會拒絕部分訂單 webhook，訂單與客戶的匯入也會失敗。

## 1. 在 Shopify Dev Dashboard 建立應用程式

自 2026 年 1 月 1 日起，Shopify 不再允許從商店後台建立舊版的「自訂應用程式」。請在商店所屬組織的 **Dev Dashboard** 中建立應用程式。這一點很重要：Odoo 是透過 client credentials grant 取得存取權杖，而這只有在應用程式與商店屬於同一個組織時才行得通。

填寫應用程式版本時：

- **App URL：** `https://shopify.dev/apps/default-app-home`。此連接器在 Shopify 內沒有任何頁面。
- **Embed app in Shopify admin：** 關閉。
- **Use legacy install flow：** 關閉，這樣安裝應用程式時就會授予您所選的權限範圍。
- **Webhooks：** 不用設定。Odoo 會在初始設定時自行建立。
- **Webhooks API version：** `2026-01`，也就是連接器使用的版本。
- **Scopes：** `read_customers`、`write_customers`、`read_draft_orders`、`read_fulfillments`、`read_inventory`、`write_inventory`、`read_locations`、`read_merchant_managed_fulfillment_orders`、`write_merchant_managed_fulfillment_orders`、`read_orders`、`read_products`、`write_products` 和 `read_returns`。如果需要 60 天以前的訂單，請再加上 `read_all_orders`。

發佈該版本，將應用程式安裝到商店，然後從應用程式的「Settings」頁面複製 client ID 與 client secret。

已經有 2026 年之前建立的舊版自訂應用程式嗎？它仍然可以使用。請將 client ID 留空，改用它的 Admin API access token 與 API secret key。

## 2. 在 Odoo 安裝模組

將 OCA 的 `queue_job` 與 [connector-shopify](https://github.com/OCA/connector-shopify) 儲存庫加入您的 addons 路徑，然後安裝 **Shopify Connector**。同步工作是以背景作業的方式執行，因此 `queue_job` 必須在伺服器層級載入，並為每個 Shopify 頻道設定處理容量：

```
[options]
server_wide_modules = base,rpc,web,queue_job

[queue_job]
channels = root:4,root.shopify.import:2,root.shopify.export:1,root.shopify.webhook:1
```

接著，將系統參數 `web.base.url` 設為 Odoo 的公開 HTTPS 網址，並將 `web.base.url.freeze` 設為 `True`。否則當管理員從其他地方登入時，網址就會改變，Shopify 的 webhook 也會被送到錯誤的位置。

## 3. 執行初始設定

將 **Shopify Manager** 群組授予您的團隊成員，然後開啟 **Shopify > Configuration > Instances**，點選 **Start Onboarding**：

1. 輸入 `yourstore.myshopify.com` 網域、client ID 與 client secret，並測試連線。Odoo 會自行取得權杖，並每 24 小時自動更新一次。
2. 檢查權限範圍。如果缺少任何一項，請將它加入新的應用程式版本，發佈後再到商店中核准。
3. 將 Shopify 的據點對應到 Odoo 的倉庫，檢查 Odoo 產生的稅務對應，並選擇同步原則。
4. 安裝 webhook，並將第一批匯入作業排入佇列。

在讓訂單與出貨自動無人值守地運行之前，請先檢查付款閘道與物流業者的對應。日後大多數的意外狀況，都來自沒有人檢查過的對應設定。

## 4. 持續留意運作狀況

**Shopify > Dashboard** 會顯示同步狀態、待處理作業數量、近期錯誤以及 webhook 的健康狀況。連接器並不假設每個 webhook 都會送達：排程作業會重新匯入 Shopify 有變更、但 Odoo 漏掉的任何資料。發生失敗時，請先排除原因，再使用 **Retry Failed Shopify Bindings**。更換網域或憑證之後，請執行 **Check Webhooks**。

## 後續步驟

加裝 **Shopify Connector Accounting**，即可將 Shopify Payments 的撥款、手續費與爭議款項匯入為日記帳分錄（我們另外說明了[為什麼撥款很難對帳](https://sawirstudio.com/zh-hant/posts/why-shopify-payouts-dont-match-your-books/)）；加裝 **Shopify Connector POS**，則可處理實體店面的銷售。

如果您希望我們協助安裝、依照貴公司的營運方式進行設定，或是擴充功能，[我們可以幫您處理](https://sawirstudio.com/zh-hant/shopify/)。

**延伸閱讀：** [為什麼您的 Shopify 撥款永遠對不上帳](https://sawirstudio.com/zh-hant/posts/why-shopify-payouts-dont-match-your-books/)，或是 [Shopify 2026-10 API：您的整合會有哪些地方壞掉](https://sawirstudio.com/zh-hant/posts/shopify-2026-10-api-what-breaks/)。
