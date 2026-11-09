---
title: Sawir Social Scheduler | SawirStudio
url: https://sawirstudio.com/zh-hant/projects/wordpress-plugins/sawir-social-scheduler/
description: 撰寫一次，就能在 wp-admin 排程發佈到 27 個社群網路。支援各頻道個別編輯、串文、媒體與預覽。您的權杖只會保存在您自己的伺服器上。
---

# Sawir Social Scheduler

![Sawir Social Scheduler 橫幅圖](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-banner.webp)

Sawir Social Scheduler 讓您在 wp-admin 撰寫一次貼文，就能排程發佈到 27 個社群網路，並支援各頻道個別編輯、串文、媒體與預覽。您的帳號權杖只會保存在您自己的伺服器上。

[在 WordPress.org 免費取得](https://wordpress.org/plugins/sawir-social-scheduler/)

## 功能介紹

Sawir Social Scheduler 把 WordPress 變成社群媒體排程工具。使用各平台提供的權杖連接您的頻道，撰寫貼文、選擇頻道、設定日期，然後發佈。每一次發送都在背景執行，完成後會回報貼文連結，或是平台本身的錯誤訊息。

中間沒有任何代管服務。憑證會以您網站的 salt 加密，且不會離開您的伺服器。

### 頻道

頻道數量不限，同一個平台也可以新增多個帳號。每個頻道都有「**測試連線**」按鈕，會取得帳號名稱與大頭貼，在您正式使用前確認權杖確實有效。

**社群：** X (Twitter)、Facebook 粉絲專頁、Instagram、Threads、LinkedIn 專頁、Bluesky、Mastodon、TikTok、YouTube、Pinterest、Tumblr、Dribbble、Google 商家檔案

**社團與聊天：** Telegram、Discord、Slack、Reddit、Lemmy、VK、Kick、Twitch、Whop

**部落格與電子報：** WordPress、Medium、Dev.to、Hashnode、Listmonk

您可以用巢狀群組依客戶或品牌整理頻道，再一鍵選取整個群組。

### 貼文

- **一份內容，多個頻道。** 以核取方塊選擇頻道。未個別調整的頻道會直接使用原本撰寫的貼文。
- **各頻道個別調整。** 在任一頻道勾選「自訂」，就能為它設定專屬的內容段落：文字、媒體與延遲時間。內容會先從貼文預填，因此您編輯的是副本。其他頻道仍使用原本撰寫的貼文。
- **串文。** 新增多個段落，即可在 X、Threads、Bluesky、Mastodon 與 Telegram 發佈串文，每個段落都可以有自己的媒體。不支援串文的平台會收到合併後的內容。
- **媒體。** 圖片與影片只需附加一次。每個平台會取得它所支援的格式：相簿、輪播、Reels、可續傳的影片上傳。
- **排程。** 可使用 WordPress 一般的「發佈時間」日期，也可以為任一頻道單獨設定時間。
- **發文時段。** 為頻道設定常用的每日發文時間。「下一個空檔」會自動填入下一個尚未使用的時段，行事曆會顯示接下來的空時段，點一下就能填入內容，API 也接受 `"date": "next"`。
- **預覽。** 每個頻道都有即時預覽，外觀貼近真實平台，並套用字數限制與媒體規則，讓您清楚看到實際會發出的內容。
- **留言與延遲。** 在不支援串文的平台上，額外的段落會以留言形式發在貼文下方。任何一個段落都可以在前一段之後等待指定的分鐘數再發出。
- **分享部落格文章。** 在一般文章編輯器中加入一個設定區塊，文章發佈時就會發到您的頻道，並可使用專屬的文字與媒體。
- **複製。** 把貼文連同頻道、段落與個別設定複製成新的草稿。您可以把草稿當作範本保存。
- **標註。** 輸入 `@handle`，網路平台會在 X、Instagram、Threads、Mastodon、Telegram、Reddit 與 Bluesky 上解析，連結與帳號會成為真正的標註。LinkedIn 公司專頁的寫法是 `@[Name](org:12345)`，其中的 ID 來自專頁網址；Facebook 粉絲專頁的寫法是 `@[12345]`。
- **重複發佈。** 常青貼文可每隔 N 天或 N 週重新發送，直到指定日期或達到指定次數。每一次重發都是獨立的貼文，有各自的發送紀錄與統計。請在文字中加入 `{n}` 來為每次重發編號，因為 X、Threads 與 Bluesky 會拒絕內容完全相同的貼文。

### 行事曆

每則貼文會拆成每個頻道各一次發送，行事曆會把每次發送顯示為依狀態上色的卡片：已排程、已發送、失敗、保留中、草稿、進行中。拖曳卡片即可更改該頻道的發佈時間。點一下卡片可查看錯誤訊息、重試按鈕、貼文連結與互動數據。點一下空的時段，就能在該時間開始撰寫貼文。提供月、週、日與清單檢視，並可依群組與狀態篩選。

### 數據分析

按讚、留言、分享、觀看次數與曝光次數，會在每次發送後 1 小時、1 天與 1 週收集，也可以隨時手動更新，目前支援 16 個平台。每個頻道的追蹤人數每天記錄一次，並附有 30 天的走勢圖。這些數字會顯示在行事曆卡片、貼文的「數據分析」區塊與頻道清單中。「數據分析」頁面會彙整最近 7、30 或 90 天與前一段期間的比較：已發送貼文數、互動、觸及與追蹤人數及其變化、追蹤人數圖表、各頻道互動，以及表現最好的貼文。

### 客戶核准

勾選「核准前先保留發送」，複製預覽連結後傳給您的客戶。客戶不需登入，就能看到貼文在每個頻道上的呈現樣子，並點選「核准」或「要求修改」並附上備註。核准後才會放行發送。每一次決定都會保留為歷史紀錄。

上傳到 X、YouTube 與 TikTok 的大型影片會分成小段執行，並可在多個背景工作之間接續，因此不會受主機時間限制影響，也不會把整個檔案載入記憶體。

### REST API

後台能做的每件事，都可以透過 `/wp-json/sawir-social/v1` 使用，方便用於自動化、行動 App 與其他網站。請以應用程式密碼進行驗證；權限依外掛的權限設定而定，且不會回傳頻道憑證。

- `GET /platforms`、`GET /channels`、`GET /channels/{id}`、`POST /channels/{id}/test`
- `GET|POST /posts`、`GET|PATCH|DELETE /posts/{id}`、`POST /posts/{id}/duplicate`、`GET /posts/{id}/analytics`
- `GET /sends`、`GET /sends/{id}`、`POST /sends/{id}/retry`、`POST /sends/{id}/stats`

以 `{"title","content","status":"future","date":"2026-12-25 09:00","media":[12],"parts":[{"text","media","delay"}],"channels":[{"id":38,"parts":[{"text","media","delay"}],"date":"…"}],"groups":["Client A"],"hold":false}` 建立貼文。頻道的 `parts` 會取代該頻道原本的貼文內容；`text` 是單一段落的簡寫。

### AI 代理

相同的操作也已註冊到 WordPress Abilities API，包括 `sawir-social/list-channels`、`create-post`、`update-post`、`list-posts`、`list-sends`、`retry-send`、`test-channel`、`get-analytics` 與 `list-platforms`，並附有具型別的輸入與輸出結構。任何透過 WordPress MCP adapter 連接的代理，都會把它們視為工具，也可以在 `/wp-json/wp-abilities/v1/abilities/{name}/run` 呼叫。Abilities API 是 WordPress 6.9 以上版本的內建功能；舊版本則由外掛內建這個功能，因此只要外掛能運作的地方，這些功能就都能使用。

### 系統需求

若您的網站有 Action Scheduler（WooCommerce 與許多其他外掛都有提供），發送就會透過它在背景執行，否則會使用 WP-Cron。使用 WordPress 內建排程時，只有在網站有流量時才會往前執行。若要即時發送，請加入系統 cron，「行事曆」畫面會顯示適用於您網站的完整指令。

對於以網址擷取媒體的平台（Instagram、Threads、Facebook、Pinterest、Telegram），媒體必須能從網際網路存取。本機開發網站無法把圖片傳送給這些平台。

### 致謝

內含 FullCalendar (MIT)、WordPress Abilities API 函式庫 (GPL)，以及來自 Simple Icons (CC0) 的品牌標誌。當其他外掛提供 Action Scheduler 時，會使用該外掛的版本。

## 畫面截圖

![Sawir Social Scheduler 畫面截圖 1](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-screenshot-1.webp)

![Sawir Social Scheduler 畫面截圖 2](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-screenshot-2.webp)

![Sawir Social Scheduler 畫面截圖 3](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-screenshot-3.webp)

![Sawir Social Scheduler 畫面截圖 4](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-screenshot-4.webp)

![Sawir Social Scheduler 畫面截圖 5](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-screenshot-5.webp)

![Sawir Social Scheduler 畫面截圖 6](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-screenshot-6.webp)

## 常見問題

### 權杖要去哪裡取得？

請到各平台的開發者設定。頻道表單會註明所需的確切憑證，例如 Facebook 粉絲專頁存取權杖、Telegram 機器人權杖或 Bluesky 應用程式密碼。建議使用長效權杖；Meta 與 LinkedIn 的權杖約 60 天後會過期，需要重新貼上。

### 我的 Telegram 測試顯示「chat not found」

私人頻道沒有 @username。請將 Chat ID 留白，然後按下「測試連線」：外掛會列出機器人已加入的所有聊天室，以及可供貼上的數字 ID。

### 為什麼發送一直停在「進行中…」？

背景佇列還沒有執行。請造訪一下網站，或設定「行事曆」畫面中顯示的系統 cron。

### 我可以在本機網站上發佈影片到 Instagram 嗎？

不行。Instagram 是以網址擷取媒體，需要公開的網址。外掛的其他功能都可以在本機正常使用。

## 推薦搭配

- [Sawir Insights](https://sawirstudio.com/zh-hant/projects/wordpress-plugins/sawir-insights/)：呈現訪客如何使用您的網站，提供熱點圖、工作階段錄影、問卷、轉換漏斗與表單分析，且所有資料都保存在您自己的伺服器上。
- [Sawir Live Translate](https://sawirstudio.com/zh-hant/projects/wordpress-plugins/sawir-live-translate/)：直接在原頁面上翻譯您的內容：選擇語言、自動翻譯畫面上看得到的文字、儲存，並為訪客提供語言切換器。

歡迎瀏覽我們所有的[免費 WordPress 外掛](https://sawirstudio.com/zh-hant/projects/wordpress-plugins/)。

- **價格** 免費
- **版本** 0.1.0
- **系統需求** WordPress 5.8、PHP 7.4
- **授權** GPL-2.0-or-later
- **WordPress.org** <https://wordpress.org/plugins/sawir-social-scheduler/>

關於我們就先說到這裡。

## 有什麼我們 能幫上忙的嗎？

[開始專案](https://sawirstudio.com/zh-hant/contact/)
