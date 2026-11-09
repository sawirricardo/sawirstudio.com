---
title: Sawir Insights | SawirStudio
url: https://sawirstudio.com/zh-hant/projects/wordpress-plugins/sawir-insights/
description: 熱點圖、工作階段錄影、問卷調查、轉換漏斗與表單分析，所有資料都留在您的伺服器上。
---

# Sawir Insights

![Sawir Insights 橫幅圖片](https://sawirstudio.com/assets/images/plugins/sawir-insights-banner.webp)

Sawir Insights 讓您看見訪客如何使用您的網站，提供熱點圖、工作階段錄影、問卷調查、轉換漏斗與表單分析，而且所有資料都保存在您自己的伺服器上。

[在 WordPress.org 免費取得](https://wordpress.org/plugins/sawir-insights/)

## 功能介紹

Sawir Insights 是 WordPress 的行為分析外掛，包含熱點圖、工作階段錄影、問卷調查、意見回饋小工具、表單分析、轉換漏斗，以及整合這些資料的儀表板。

在多站網路（multisite）中，每個網站都會獨立追蹤並設定自己的分析資料。

簡單來說：這是可自行架設的 Hotjar 或 Microsoft Clarity 替代方案，訪客資料會留在您自己的資料庫裡。

它誕生的原因很單純。這類工具中的老牌產品確實很好用，但運作方式是把訪客的一舉一動記錄下來，傳送到別人的伺服器上。對許多網站來說，尤其是歐盟地區的網站，這種做法從令人不安到根本行不通都有可能。我想要的是同類型的工具，但資料絕不離開收集它的網站。所以這個外掛就是圍繞這個前提打造的：您的資料庫、您的上傳目錄、您的保存規則。背後沒有任何 SaaS 服務，不需要建立帳號，也不會有任何形式的回傳連線。

### 功能介紹

- 點擊、滑鼠移動與捲動熱點圖，直接呈現在您的線上頁面上，並區分桌機、平板與手機。
- 可在 wp-admin 內播放工作階段錄影。輸入欄位的內容一律會被遮蔽；錄影功能預設關閉，需要您自行開啟。
- 彈出式問卷與常駐的意見回饋標籤，可依網址、裝置、流量比例、停留時間、捲動深度或離開意圖來設定顯示對象。
- 表單分析：看出使用者在哪裡猶豫、哪些欄位留白、在哪裡放棄。只記錄選擇器與時間，絕不記錄任何人輸入的內容。
- 有順序的轉換漏斗，並提供每日報表。
- 儀表板提供流量 KPI 與各期間的變化幅度、每日趨勢圖，以及裝置、瀏覽器、頁面、來源網站、行銷活動、漏斗與錄影的細項分析。

### 它刻意不做的事

我認為，在安裝工具之前，您應該先知道它拒絕做什麼。

- 它不會儲存 IP 位址。無論是雜湊或截斷後的版本，一律不存。
- 它絕不記錄表單內容。擷取層只會傳送中繼資料，因此不會有外洩的問題。
- 它不會從 CDN 載入任何內容，也不會將任何資料傳送到 CDN。rrweb 重播函式庫已內含在外掛中。
- 它不會永遠保留原始資料。事件會依照您控制的保存期程進行彙總與清除，錄影檔則有獨立的保存期限，以及會優先刪除最舊檔案的磁碟用量上限，因為悄悄把主機硬碟塞滿的外掛，不是好外掛。

前端追蹤程式是小巧的原生 JavaScript，透過 `sendBeacon` 批次傳送，並提供取樣控制。沒有被抽樣進錄影的訪客，完全不會下載重播函式庫。

### 隱私立場

預設會遵守「請勿追蹤」（Do Not Track）設定。同意模式 API（`window.swinsConsent`）可與您現有的同意管理工具整合。工作階段、事件與錄影都已串接 WordPress 內建的個人資料匯出與清除工具。您的實際部署是否符合 GDPR，取決於您的告知內容與法律依據，這部分沒有任何外掛能代勞，但資料最小化的工作已經完成。

### 支援

請透過此外掛在 wordpress.org 的論壇取得支援，或寄信至 ricardo@sawirstudio.com。

## 常見問題

### 與其他方案相比如何？

Hotjar 與 Microsoft Clarity 這類託管式分析服務都是優秀的工具，免費方案也相當大方，但它們的運作方式，是把您訪客的行為記錄傳送到它們的伺服器。Sawir Insights 捨棄了這種跨網站的基礎架構，只堅持一個前提：所有資料都在您自己的網站上收集、儲存與檢視。如果您想要託管服務的額外功能，也不介意由第三方處理資料，請選擇託管服務；如果訪客資料離開您的伺服器是無法接受的，請選擇 Sawir Insights。錄影與熱點圖都存放在您自己的資料庫與上傳目錄中，因此保存設定在這裡比在 SaaS 服務上更重要。

### Sawir Insights 符合 GDPR 嗎？

這個外掛是為重視隱私的部署環境而打造：資料留在您的伺服器上、絕不儲存 IP 位址、保存期限可自行設定、遵守「請勿追蹤」，並整合了 WordPress 的匯出與清除工具。至於是否合規，取決於您的設定、告知內容、法律依據與同意機制的實作，請就您的網站諮詢合格的法律顧問。

### 這和 Hotjar 或 Microsoft Clarity 有什麼不同？

它們是託管式服務：功能強大，但您訪客的行為會依照它們的條款被記錄到它們的伺服器上。Sawir Insights 涵蓋同一類功能，包括熱點圖、工作階段錄影、問卷調查、轉換漏斗與表單分析，但它是一個把所有資料都留在您自己伺服器上的外掛，不需要帳號、不使用外部服務，也不會分享資料。

### 訪客資料會離開我的伺服器嗎？

不會。這個外掛不會將任何行為資料傳送給 Sawir Studio 或任何分析服務供應商。事件資料存放在您的 WordPress 資料庫中，錄影檔則存放在受保護的上傳目錄底下。

## 推薦搭配

- [Sawir Pulse](https://sawirstudio.com/zh-hant/projects/wordpress-plugins/sawir-pulse/)：以角落的小通知，讓購物者看到您商店裡實際發生的事，例如最新訂單、正在瀏覽商品的人數和最新評價。
- [Sawir Live Translate](https://sawirstudio.com/zh-hant/projects/wordpress-plugins/sawir-live-translate/)：直接在原頁面上翻譯您的內容：選擇語言、自動翻譯畫面上看得到的文字、儲存，並為訪客提供語言切換器。

歡迎瀏覽我們所有的[免費 WordPress 外掛](https://sawirstudio.com/zh-hant/projects/wordpress-plugins/)。

- **價格** 免費
- **版本** 1.0.0
- **需求** WordPress 6.2、PHP 7.4
- **授權** GPLv2 或更新版本
- **WordPress.org** <https://wordpress.org/plugins/sawir-insights/>

關於我們就先說到這裡。

## 有什麼我們 能幫上忙的嗎？

[開始專案](https://sawirstudio.com/zh-hant/contact/)
