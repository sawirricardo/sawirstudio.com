---
title: 打造專屬於您的 CMS，就像 WordPress 一樣 | SawirStudio
url: https://sawirstudio.com/zh-hant/posts/how-to-build-your-own-content-management-system-like-wordpress/
description: 使用 Laravel 與現代軟體架構，打造類似 WordPress 的內容管理系統。
---

# 如何自行打造內容管理系統（類似 WordPress）

從零開始打造內容管理系統，是一段非常棒的經驗。

WordPress 一直是事實上的標準，尤其在內容管理系統領域。根據其官方網站……

> WordPress 驅動了超過 43% 的網站，而且這個數字每天都在成長。從簡單的網站、部落格，到複雜的入口網站、企業網站，甚至應用程式，都能用 WordPress 打造。

不過，我們可以看到，WordPress 誕生的年代，開發應用程式還沒有那麼容易。

現在，隨著近期 [PHP 更新](https://www.php.net/)的推出，PHP 開發已大幅加速，軟體開發的實務做法也不斷演進。

例如，受到 Ruby on Rails 啟發的 MVC 架構開發方式，以及一些軟體架構原則，已經被許多軟體框架採用。

然而，WordPress 並沒有採用這些做法，它仍然使用自己的一套結構。

因此，本文希望協助您了解，如果像 WordPress 這樣的內容管理系統採用現代軟體架構，可以做到什麼。

為了說明這一點，我們會使用 [Laravel](https://laravel.com)，這是一套 PHP 網頁框架，已經實踐了一些軟體工程的做法。

先簡單討論一下我們的內容管理系統會有哪些功能。根據 WordPress 的功能，我們可以歸納出一些主要功能，例如：

- 輕鬆發布 如果您曾經建立過文件，就已經很懂得如何用 WordPress 建立內容。您可以建立文章和頁面、輕鬆設定格式、插入媒體，只要按一下按鈕，內容就會上線。
- 發布工具 WordPress 讓您輕鬆管理內容。您可以建立草稿、排程發布、查看文章修訂紀錄。您可以將內容設為公開或私人，也可以用密碼保護文章和頁面。
- 使用者管理 並非每個人都需要相同的網站存取權限。管理員負責管理網站，編輯負責處理內容，作者和投稿者負責撰寫內容，訂閱者則有可自行管理的個人檔案。這讓您的網站可以有各種不同的貢獻者，也讓其他人單純成為社群的一份子。
- 媒體管理 俗話說一張圖勝過千言萬語，所以能快速、輕鬆地將圖片和媒體上傳到 WordPress 非常重要。只要將媒體拖放到上傳區，就能加入您的網站。您可以加上替代文字和說明文字，並將圖片和圖庫插入內容中。我們甚至加入了一些圖片編輯工具，讓您玩得開心。
- 內建留言功能 您的部落格就是您的家，而留言為朋友和追蹤者提供了與內容互動的空間。WordPress 的留言工具提供您所需的一切，讓您成為討論的平台，並管理這些討論。
- 搜尋引擎最佳化 WordPress 開箱即針對搜尋引擎進行最佳化。如果需要更細緻的 SEO 控制，也有許多 SEO 外掛可以代勞。

那麼，我們開始動手吧！

安裝

```
composer create-project laravel/laravel wp-laravel
```

來試著存取我們的應用程式

```
php artisan serve
```

我們需要設定資料庫，在本文中可以使用 SQLite。

```
touch database/database.sqlite
```

我們需要更新 .env

```
DB_CONNECTION=sqlite
DB_DATABASE=database/database.sqlite
DB_FOREIGN_KEYS=true
```

[參考資料](https://laravel.com/docs/9.x/database#sqlite-configuration)

接著執行初始的資料庫遷移

```
php artisan migrate
```

現在，我們可以看到資料庫的變更如下

```
php artisan db:show
```

所有遷移都已套用。

接下來，我們想要一個類似 WordPress 文章（Post）的東西。如果您曾經檢視過 WordPress 的資料表，就會知道有一個名為 wp\_posts 的資料表。我們來建立它

```
php artisan make:model Post
```

這會建立一個名為 Post 的模型。接著來建立遷移檔，以套用資料庫變更。

```
php artisan make:migration "create posts table"
```

現在，打開遷移檔，並將程式碼更新如下：

```
$table->string('title');
$table->string('slug')->nullable();
$table->text('excerpt')->nullable();
$table->longText('content')->nullable();
$table->dateTime('published_at')->nullable();
$table->foreignIdFor(\App\Models\User::class)->constrained();
```

**延伸閱讀：** [如何設計訂閱方案的限制](https://sawirstudio.com/zh-hant/posts/how-do-you-design-your-subscription-plan-restrictions/)，或參閱我們製作的[免費 WordPress 外掛](https://sawirstudio.com/zh-hant/projects/wordpress-plugins/)。
