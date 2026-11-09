---
title: WordPress 7.2：外掛作者該測試什麼 | SawirStudio
url: https://sawirstudio.com/zh-hant/posts/wordpress-7-2-what-plugin-authors-should-test/
description: WordPress 7.2 預定於 2026 年 12 月 8 日發布。本文說明外掛作者會遇到哪些變化，從 Application Passwords、提案中的 Secrets API 到 Interactivity API 路由器，以及現在該如何測試。
---

# WordPress 7.2 將於 12 月 8 日推出：我們正在測試外掛的哪些部分

WordPress 7.2 預定於 2026 年 12 月 8 日發布。第一個 Release Candidate（RC）預計在 11 月 17 日推出，屆時功能清單會凍結，並公布 Field Guide（收錄所有開發者說明的清單）。在那之前，Beta 版是最省成本、最早發現外掛問題的地方。

我們在 WordPress.org 上維護 [12 款免費外掛](https://sawirstudio.com/zh-hant/projects/wordpress-plugins/)，所以這份是我們自己的檢查清單。其中有些項目在我們撰寫時仍只是提案，請以 Field Guide 為準。

## 現在就能這樣測試

- 在測試站上使用 **WP-CLI**：`wp core update --version=7.2-beta4`，或改用最新的 Beta 或 RC 版本。
- **WordPress Beta Tester 外掛**，設定為「Bleeding edge」與「Beta/RC Only」。
- **WordPress Playground**，不用安裝任何東西，直接在瀏覽器中操作看看。

千萬不要在正式站上測試。請開啟 `WP_DEBUG` 與 `WP_DEBUG_LOG`，並留意棄用通知，大多數早期警告都會出現在那裡。

## 我們檢查的項目

### Application Passwords

7.2 強化了 Application Passwords：更準確地偵測本機與 HTTPS 環境、新增時會寄出電子郵件通知，並以更安全的方式處理預設角色。如果您的外掛，或您客戶的工具是以這種方式向 REST API 驗證身分，請完整跑過一遍這個流程。

### Secrets API

這是核心提供的加密憑證儲存方式，讓外掛不必再把 API 金鑰以純文字存在選項中。它原本提案納入 7.2，管理畫面則延到 7.3。如果您的外掛會儲存 API 金鑰，請留意 Field Guide：一旦正式推出，金鑰就該存放在這裡。我們的免費外掛目前都不會儲存第三方金鑰，所以對我們來說，這主要與客戶專案有關。

### Sudo mode

規劃中的功能：在執行高權限操作前，要求已登入的使用者再次確認密碼。如果您的外掛會在自己的畫面中或透過 AJAX 執行這類操作，請測試核心在過程中要求重新驗證時會發生什麼事。

### Interactivity API 路由器

新的路由器可用來監看用戶端導覽，取代已棄用的 `state.navigation.hasStarted` 與 `hasFinished` 這一組狀態。此外也新增了 `data-wp-html` 指令，用來渲染可信任的標記。

### 後台圖示與通知

管理列與選單中的 Dashicons 將改用 SVG 圖示，後台通知也改善了螢幕閱讀器的朗讀效果。如果您的外掛自訂了選單圖示的樣式，或在後台 CSS 中指定了 Dashicon 的 class，請在 Beta 版中檢查一下。

### DataViews 與網站編輯器

外掛可以在網站編輯器中註冊自己的畫面，並從伺服器端控制 DataViews 的欄位與版面配置。這部分不會造成相容性問題，但如果您的外掛有列表畫面，了解未來的方向會很有幫助。

## 測試之後

請等到真的測試過之後，再調整 readme 中的「Tested up to」，不要提前改。當外掛沒有針對最近幾個主要版本測試過時，WordPress.org 會向使用者顯示警告，而這個警告會讓安裝數下降。我們的外掛目前標示「Tested up to: 7.1」，等通過 RC 測試後才會改成 7.2。

新的預設佈景主題 Ipsum 也會一併推出。如果您的外掛會在前台輸出任何內容，請在這個佈景主題下檢查，因為新網站就會是這個樣子。

如果您有外掛需要更新以支援 7.2，或有一個已經沒人維護的客製外掛，[歡迎讓我們看看](https://sawirstudio.com/zh-hant/contact/)。

**延伸閱讀：** [用 Abilities API 讓 AI 代理程式也能使用您的 WordPress 外掛](https://sawirstudio.com/zh-hant/posts/wordpress-abilities-api-for-plugins/)，或是 [印尼的企業該選 WooCommerce 還是 Shopify？](https://sawirstudio.com/zh-hant/posts/woocommerce-or-shopify-indonesia/)
