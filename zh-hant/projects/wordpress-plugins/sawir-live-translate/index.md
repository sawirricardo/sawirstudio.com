---
title: Sawir Live Translate | SawirStudio
url: https://sawirstudio.com/zh-hant/projects/wordpress-plugins/sawir-live-translate/
description: Sawir Live Translate 讓您直接在原頁面上完成翻譯：選擇語言、自動翻譯畫面上的文字、儲存，再為訪客提供語言切換器。
---

# Sawir Live Translate

![Sawir Live Translate 橫幅圖片](https://sawirstudio.com/assets/images/plugins/sawir-live-translate-banner.webp)

Sawir Live Translate 讓您直接在原頁面上完成翻譯：選擇語言、自動翻譯畫面上的文字、儲存，再為訪客提供語言切換器。

[在 WordPress.org 免費取得](https://wordpress.org/plugins/sawir-live-translate/)

## 功能介紹

開啟網站上的任一頁面，點選管理列中的「Translate this page」，選擇目標語言，然後您可以自動翻譯所有畫面上的文字（透過免費的機器翻譯服務），也可以一邊看著頁面即時更新，一邊手動編輯翻譯。

儲存後的翻譯會由伺服器端直接輸出到 HTML 回應中（對 SEO 友善：不會先閃現未翻譯的文字，且 head 中會加入 hreflang 替代連結），訪客則透過連結式的切換器（位於右下角）切換語言。

網址使用語言路徑前綴，例如西班牙文為 /es/about，原始語言則為 /about（維持為標準網址，不加前綴）。語言的判定順序如下：/es/ 路徑 → ?lang=es → 已儲存的 Cookie → 瀏覽器的 Accept-Language → 原始語言。根網址與不含前綴的網址會自動偵測訪客的瀏覽器語言，且不會進行重新導向（外掛會送出 Vary 標頭，讓共用快取能正確運作）。

1. 前往「設定」→「Live Translate」：設定原始語言，並勾選要提供的語言。2. 開啟前台頁面，點選管理列中的「🌐 Translate this page」。3. 選擇語言，按下「Auto-translate empty」（或自行編輯字串），再按「Save」。4. 訪客使用切換器（位於右下角）切換語言，每種語言都有各自的 /xx/ 網址。

注意：語言路徑前綴的前提是安裝在根網域，且沒有任何頂層頁面的網址代稱（slug）是兩個字母的語言代碼。

## 常見問題

### 與其他方案相比如何？

WPML、Polylang、TranslatePress 這類多語系框架，會把翻譯內容視為網站結構的核心，提供各語言獨立的編輯流程、翻譯服務，以及與頁面編輯器和 WooCommerce 的深度整合。Sawir Live Translate 刻意走輕量路線：開啟頁面，翻譯畫面上看得到的內容（自動或手動皆可），再以帶有語言前綴的網址搭配 hreflang 標籤提供給訪客。如果您要建置完整的多語系網站，包含翻譯選單、商品和編輯流程，請使用上述框架；如果您只是想讓現有頁面今天下午就能以另一種語言閱讀，這個外掛就能做到。

### 我切換了語言，但內容沒有任何變化

幾乎都是因為該語言還沒有儲存任何翻譯。請檢查「設定」→「Live Translate」→「Saved translations」，如果清單中沒有該語言（或顯示 0 個字串），請再次開啟頁面編輯器，執行自動翻譯並按下「Save」。在您儲存之前，編輯器中的即時預覽並不會被保存。

### 部分字串仍然沒有翻譯

頁面內的編輯器只能看到您所開啟頁面上的字串。位於其他頁面的文字，需要把那些頁面也翻譯好，或是執行全站掃描。由滑桿或指令碼在頁面載入後才插入的字串，可能需要透過「Edit all site strings」手動新增。

### 我修改了內容，翻譯就不見了

翻譯是依據完整的原文來比對的。如果您改寫了某個句子，只要將新版本翻譯一次並儲存即可。

## 推薦搭配

- [Sawir Insights](https://sawirstudio.com/zh-hant/projects/wordpress-plugins/sawir-insights/)：呈現訪客如何使用您的網站，提供熱點圖、工作階段錄影、問卷、轉換漏斗與表單分析，且所有資料都保存在您自己的伺服器上。
- [Sawir Media Folders](https://sawirstudio.com/zh-hant/projects/wordpress-plugins/sawir-media-folders/)：為 WordPress 媒體庫加入資料夾功能，數量不限，並可透過拖放來整理檔案。

歡迎瀏覽我們所有的[免費 WordPress 外掛](https://sawirstudio.com/zh-hant/projects/wordpress-plugins/)。

- **價格** 免費
- **版本** 1.1.0
- **系統需求** WordPress 5.8、PHP 7.4
- **授權** GPL-2.0-or-later
- **WordPress.org** <https://wordpress.org/plugins/sawir-live-translate/>

關於我們就先說到這裡。

## 有什麼我們 能幫上忙的嗎？

[開始專案](https://sawirstudio.com/zh-hant/contact/)
