---
title: Shopify 2026-10 API：哪些地方會出問題 | SawirStudio
url: https://sawirstudio.com/zh-hant/posts/shopify-2026-10-api-what-breaks/
description: Shopify Admin API 2026-10 版本中可能讓整合出問題的變更，以及該如何調整：運費、草稿訂單折扣、稅額、條碼等。
---

# Shopify 2026-10 API：您的整合有哪些地方會出問題

Shopify 於 2026 年 10 月 1 日釋出了 API 的 2026-10 版本。其中大部分是新功能，等到需要時再處理即可。但有少數幾項變更，可能會讓原本在 2026-07 運作正常的整合悄悄出錯，建議您在升級版本之前，花一個小時檢查這幾項。

這是我們用來檢查自家整合的清單，其中包括 [Odoo 的 Shopify 連接器](https://sawirstudio.com/zh-hant/projects/odoo-shopify-connector/)。清單並未涵蓋發行說明的每一行，只列出會改變行為的項目。

## 會悄悄出錯的變更

### 新的物流服務商不會自動在結帳時顯示運費

使用 `carrierServiceCreate`（或 REST 的 `POST carrier_services.json`）建立物流服務商，現在只會註冊該服務，不會再把它的運費加入「一般運送方案」。如果您的應用程式或初始設定腳本建立了物流服務商，並預期顧客會看到您的運費，那麼在您自行把該服務加入運送方案之前，顧客是看不到的。

### 變更收件地址會重新計算稅額

使用 `orderUpdate` 或 REST 的訂單更新端點，更新尚未出貨訂單的收件地址時，現在會重新計算該訂單的稅額。這樣的行為是正確的，但如果您是從 ERP 或客服工具修改地址，訂單總額可能會在您不知情的情況下改變。更新後請重新讀取 `taxLines`、`totalTaxSet` 以及各項總額，並處理已付款或已退款金額的差額。

### 草稿訂單的折扣警告不再包含 `priceRule`

`DraftOrderDiscountNotAppliedWarning.priceRule` 與舊版的 `PriceRule` 類型已經移除。如果您使用 `draftOrderCalculate`、`draftOrderCreate` 或 `draftOrderUpdate` 建立草稿訂單，並且會讀取折扣未套用的原因，請改為選取 `discountTitle` 與 `discountCode`。

### 錯誤的 metafield 篩選條件現在會直接報錯

以前，如果商品查詢使用了無法篩選的 metafield，系統會忽略那部分條件並回傳所有商品。現在則會回傳錯誤。這樣比較好，但如果您的同步作業幾個月來一直回傳「所有商品」，在切換版本的那一天就會開始失敗。請確認您用來篩選的 metafield 定義支援篩選，也支援您所使用的運算子。

### 一個商品款式可以有多個條碼

`ProductVariant.barcodes` 是新的 connection，最多可包含 20 個具類型的條碼。舊的 `ProductVariant.barcode` 已被棄用，只會回傳第一個條碼。如果您的倉儲或 POS 系統是以條碼比對商品，請讀取完整的條碼清單。

### 移除了一個庫存錯誤

`ITEM_NOT_STOCKED_AT_LOCATION` 已從庫存 mutation 與 `InventorySetQuantitiesUserErrorCode` 中移除。依據它進行判斷的程式碼現在已成為無效程式碼，請把它刪除，以免日後有人誤以為仍可依賴。

## 較小的變更，值得在程式碼中搜尋一下

- **顧客分群。** `MATCHES` 與 `NOT MATCHES` 取代了 `= true` 與 `= false`，而像 `12_months_ago` 這類具名日期已被棄用，改用 `-12m` 這類位移寫法。
- **簡訊同意。** 扁平的簡訊同意欄位已被棄用，請改為讀取 `CustomerPhoneNumber.smsMarketingConsent`。
- **分階段推出中的折扣。** 當折扣屬於分階段推出（rollout）的一部分時，光看 `startsAt` 與 `endsAt` 無法判斷折扣是否已生效。讀取 `Discount.rollouts` 需要 `read_rollouts` 權限範圍。
- **出貨狀態。** `FULFILLMENT_NOT_REQUIRED` 是新增的值。如果您的 `switch` 對應了每一種狀態，這個新值會落入 default 分支。
- **Customer Account API。** `Customer.lastIncompleteCheckout` 與 `Checkout` 類型已被移除，且沒有替代方案。
- **POS UI 擴充功能。** 像 `applyCartDiscount` 這類購物車 mutation，在 POS 未確認變更時現在會被拒絕，而且 `session.currentSession.staffMemberId` 已被移除。

## 兩個與 2026-10 無關的截止日期

如果您維護的是公開應用程式，必須在 2027 年 1 月 1 日前，把不會過期的離線存取權杖換成會過期的權杖。自訂應用程式，以及商家在 Dev Dashboard 中建立的應用程式不在此限。另外，線上商店的 script tag 將於 2027 年 3 月 1 日停止執行；我們另外寫了一篇文章，說明[這對商店有什麼影響](https://sawirstudio.com/zh-hant/posts/shopify-script-tags-stop-march-2027/)。

## 我們如何升級

我們會把 API 版本固定在單一位置，在開發商店上以新版本執行整合，並閱讀 [2026-10 更新日誌](https://shopify.dev/changelog?api_version=2026-10)中篩選為「Action required」的項目。2026-10 版本至少會支援到 2027 年 10 月 1 日，所以時間還算充裕，但也不宜拖到警告郵件開始寄來才處理。

如果您的商店需要與 Odoo 或其他系統串接，而且您希望由別人來處理這件事，這正是[我們的專長](https://sawirstudio.com/zh-hant/shopify/)。

**延伸閱讀：** [Shopify Script Tags 將於 2027 年 3 月 1 日停止運作，店家該怎麼做](https://sawirstudio.com/zh-hant/posts/shopify-script-tags-stop-march-2027/)，或[將 Shopify 連接到 Odoo 19 Community：逐步教學](https://sawirstudio.com/zh-hant/posts/connect-shopify-to-odoo-19-community/)。
