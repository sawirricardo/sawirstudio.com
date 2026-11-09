---
title: Shopify Script Tags 將於 2027 年 3 月停用 | SawirStudio
url: https://sawirstudio.com/zh-hant/posts/shopify-script-tags-stop-march-2027/
description: Shopify 將於 2027 年 3 月 1 日停止執行線上商店的 script tag。本文說明受影響的應用程式、如何檢查您的商店，以及有哪些替代方式。
---

# Shopify Script Tags 將於 2027 年 3 月 1 日停止運作，店家該怎麼做

2027 年 3 月 1 日起，Shopify 將不再把 script tag 注入線上商店的前台。凡是仍以這種方式載入的功能，例如評論小工具、聊天泡泡、追蹤像素，都會直接消失。不會有明顯的錯誤訊息，功能只是悄悄不見了。

Shopify 已在 [2026 年 8 月](https://shopify.dev/changelog/posts/online-store-script-tags-deprecation)公告這項異動。若您經營 Shopify 商店，以下是這件事對您的影響。

## 什麼是 script tag

script tag 是讓應用程式透過 API，在不修改佈景主題的情況下，把 JavaScript 加進商店前台的方式。應用程式開發者喜歡它，原因在此，問題也出在這裡：script tag 不會顯示在佈景主題編輯器中，所以您無法在那裡看到或關閉它們。已解除安裝的應用程式有時也會留下殘餘的 script tag。

## 時程

- **2026 年 10 月 1 日：** 應用程式無法再建立或更新 script tag，所有 API 版本皆然。既有的 script tag 仍會繼續執行。
- **2027 年 3 月 1 日：** Shopify 停止載入線上商店的 script tag。

訂單狀態頁面上的 script tag 已提前停用，並依另一套時程進行，所以如果那裡有功能已經不再顯示，原因就在這裡。

## 如何檢查您的商店

Shopify 沒有提供店家查看 script tag 清單的功能，所以必須自己查：

1. **列出會在前台顯示內容的應用程式。** 例如評論、聊天、彈出視窗、加購推薦、會員回饋、分析工具、廣告像素。只在後台運作的應用程式不受影響。
2. **查看佈景主題編輯器中的「應用程式嵌入」。** 如果應用程式出現在那裡且已啟用，代表它使用的是較新的方式，不會有問題。
3. **查看網頁原始碼。** 開啟您的商店，檢視原始碼並搜尋 `asyncLoad`。在大多數商店中，script tag 的網址都會列在該區塊內，每個網域通常對應一個應用程式。
4. **詢問應用程式開發者。**「您的應用程式有使用 script tag 嗎？2027 年 3 月 1 日之後還能正常運作嗎？」持續維護的應用程式，一句話就能回答這個問題。

Shopify 多年來一直推動應用程式改用佈景主題應用程式擴充功能，所以大多數維護良好的應用程式早已沒有問題。風險主要在老舊的應用程式、多年前請人為您開發的自訂應用程式，以及已不再使用的應用程式所留下的殘餘。

## 取代 script tag 的方式

- **應用程式嵌入區塊**，屬於佈景主題應用程式擴充功能的一部分，適用於所有會顯示在頁面上的內容。您可以在佈景主題編輯器中啟用，也就是說，您終於能看見它們了。
- **Web pixels**，適用於只負責收集分析或轉換資料的腳本。它們在沙箱中執行，應用程式支援之後，您不需要再做任何事。

## 如果應用程式不會再更新

請在 3 月前更換它，或請人重新開發其功能。若是自訂應用程式，通常的做法是把腳本移到佈景主題應用程式擴充功能中；如果只是做追蹤，則改用 web pixel。順便把已解除安裝的應用程式所留下的 script tag 刪除；`scriptTagDelete` mutation 目前仍可使用。

如果您有需要遷移的 Shopify 自訂應用程式，或不確定商店裡載入了哪些內容，歡迎[把商店網址寄給我們](https://sawirstudio.com/zh-hant/contact/)，我們會告訴您查到的結果。想進一步了解我們的 Shopify 服務，請看[這個頁面](https://sawirstudio.com/zh-hant/shopify/)。

**延伸閱讀：** [Shopify 2026-10 API：您的整合會有哪些地方出問題](https://sawirstudio.com/zh-hant/posts/shopify-2026-10-api-what-breaks/)，或[逐步教學：將 Shopify 與 Odoo 19 Community 串接](https://sawirstudio.com/zh-hant/posts/connect-shopify-to-odoo-19-community/)。
