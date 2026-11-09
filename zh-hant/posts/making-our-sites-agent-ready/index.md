---
title: 讓網站為 AI 代理做好準備 | SawirStudio
url: https://sawirstudio.com/zh-hant/posts/making-our-sites-agent-ready/
description: 我們如何讓產品網站可供 AI 代理閱讀：每個頁面的 Markdown 版本、Link 標頭、llms.txt、robots.txt 中的 Content Signals、API 目錄，以及 Agent Skills。
---

# 讓我們的網站為 AI 代理做好準備，需要做哪些事

越來越多的網站訪客並不是真人，而是 AI 代理：它們讀取頁面來回答某個人的問題，或是代替某個人使用產品。本週我們檢視了這個網站和我們的產品網站，包括 [SawirPost](https://sawirstudio.com/zh-hant/projects/sawirpost/)、[Sawirly](https://sawirstudio.com/zh-hant/projects/sawirly/)、[Rindraw](https://sawirstudio.com/zh-hant/projects/rindraw/)、[Math Daily](https://sawirstudio.com/zh-hant/projects/math-daily/) 等，讓代理更容易閱讀與使用。

我們以 [isitagentready.com](https://isitagentready.com) 作為檢查清單。它會將網站分成不同等級評分，以下是每個階段我們做了什麼。

## 每個頁面都有 Markdown 版本

代理在讀取 HTML 時，得先穿過導覽列、腳本和頁尾，才能找到幾段真正的內容。因此，這個網站的每個頁面現在旁邊都有一份 Markdown 版本，位於 `index.md`，只包含頁面的主要內容。這些檔案由一支小腳本從 HTML 產生，所以內容不會與原頁面不一致。

代理有兩種方式可以取得：在網址後加上 `index.md`，或是在請求中帶上 `Accept: text/markdown`。我們在靜態檔案前面放了一個小型的 Cloudflare Worker，它會檢查這個標頭，並回傳 Markdown 版本，同時在 `x-markdown-tokens` 標頭中附上大略的 token 數量。瀏覽器看到的仍然是一般頁面。

```
curl -H "Accept: text/markdown" https://sawirstudio.com/shopify/
```

## Link 標頭

現在每個 HTML 頁面都會透過 HTTP 的 `Link` 標頭（搭配 `rel="alternate"`）告知 Markdown 版本的位置。首頁還會指向 `llms.txt` 和網站地圖。代理只要發出一次請求，就能找到所有內容，不必猜測網址。

## llms.txt

`llms.txt` 是給語言模型看的純文字網站索引：一行簡介，加上每個頁面及其描述的清單，並連結到對應的 Markdown 副本。我們的 llms.txt 由同一支腳本根據網站地圖產生。

## Content Signals

`robots.txt` 只說明誰可以爬取，並沒有規範對方可以如何運用讀到的內容。[Content Signals](https://contentsignals.org/) 補上了這一點。我們只用一行，並允許全部三種用途，因為我們希望被找到、被引用：

```
Content-Signal: ai-train=yes, search=yes, ai-input=yes
```

光是這三個步驟，就讓所有網站達到了第 3 級。

## API 目錄與 Agent Skills

SawirPost 和 Sawirly 提供 API，所以我們做得更進一步。它們各自在 `/.well-known/api-catalog` 發布了 API 目錄（RFC 9727）：這是一份小型的 JSON 文件，指向 OpenAPI 說明和文件。代理一進到網站，不用讀任何頁面就能找到 API。

它們也發布了 Agent Skills，也就是簡短的指引，告訴代理如何使用該產品完成常見任務。這讓兩者都達到了第 4 級。

## 我們停在哪裡

第 5 級要求代理能夠透過 OAuth 代理註冊流程，自行註冊並完成驗證。光是有一個頁面說明如何用電子郵件和密碼註冊並不算數。這是一個需要真正開發的功能，也牽涉到實際的安全問題，所以我們選擇仔細思考，而不是倉促上線。

## 值得嗎？

對靜態網站來說，大部分的工作一個下午就能完成：一支腳本、一個小型 Worker，再加上 `robots.txt` 裡的一行。對訪客來說，網站沒有任何改變。我們認為未來會有更多問題由代理閱讀像我們這樣的網站來回答，而一份乾淨的 Markdown 頁面，正是讓內容被正確引用的最簡單方式。

我們也會在[上線與成長](https://sawirstudio.com/zh-hant/launch-and-growth/)服務中，為客戶做同樣的設定。

**繼續閱讀：** [用 Abilities API 讓您的 WordPress 外掛可供 AI 代理使用](https://sawirstudio.com/zh-hant/posts/wordpress-abilities-api-for-plugins/)，或是[我們首頁上可以拋擲的 App 圖示](https://sawirstudio.com/zh-hant/posts/throwable-app-icons-matter-js/)。
