---
title: WordPressのようなCMSを自作する方法 | SawirStudio
url: https://sawirstudio.com/ja/posts/how-to-build-your-own-content-management-system-like-wordpress/
description: Laravelと現代的なソフトウェアアーキテクチャを使って、WordPressのようなコンテンツ管理システムを構築します。
---

# コンテンツ管理システム（WordPress風）を自作する方法

コンテンツ管理システムをゼロから作る経験は、とても面白いものでした。

WordPressは、特にコンテンツ管理システムの分野で事実上の標準となっています。公式サイトには次のように書かれています。

> WordPressはウェブ全体の43%以上で使われており、その数字は日々増え続けています。シンプルなウェブサイトからブログ、複雑なポータルサイトや企業サイト、さらにはアプリケーションまで、あらゆるものがWordPressで作られています。

ただし、WordPressはアプリケーションを作ることが今ほど簡単ではなかった時代に開発されたものです。

現在では、最近の[PHPのアップデート](https://www.php.net/)に見られるようにPHP開発が加速し、ソフトウェア開発の手法も進化しています。

たとえば、Ruby on Railsに影響を受けたMVC構造での開発や、いくつかのソフトウェアアーキテクチャの原則が、多くのソフトウェアフレームワークに取り入れられています。

しかし、WordPressにはこれが当てはまらず、今も独自の構造を使い続けています。

そこでこの記事では、WordPressのようなコンテンツ管理システムに現代的なソフトウェアアーキテクチャを適用すると何ができるのかを、理解していただくことを目指します。

そのために、ソフトウェアエンジニアリングの実践手法を取り入れているPHPのウェブフレームワーク、[Laravel](https://laravel.com)を使います。

まず、今回作るコンテンツ管理システムにどんな機能を持たせるのかを簡単に整理しましょう。WordPressの機能をもとにすると、次のような主な機能を抜き出せます。

- 簡単に公開 文書を作成したことがあれば、WordPressでのコンテンツ作成もすぐに使いこなせます。投稿や固定ページを作成し、簡単に書式を整え、メディアを挿入して、ボタンをクリックするだけでコンテンツをウェブ上に公開できます。
- 公開ツール WordPressなら、コンテンツの管理も簡単です。下書きの作成、公開の予約、投稿のリビジョンの確認ができます。コンテンツを公開または非公開に設定したり、投稿や固定ページをパスワードで保護したりすることもできます。
- ユーザー管理 すべての人が同じ権限を必要とするわけではありません。管理者はサイトを管理し、編集者はコンテンツを扱い、投稿者と寄稿者はコンテンツを執筆し、購読者は自分で管理できるプロフィールを持ちます。これにより、さまざまな寄稿者をウェブサイトに迎えつつ、それ以外の方にはコミュニティの一員として参加してもらうことができます。
- メディア管理 「百聞は一見にしかず」と言うように、画像やメディアをすばやく簡単にWordPressへアップロードできることは大切です。メディアをアップローダーにドラッグ＆ドロップするだけで、ウェブサイトに追加できます。代替テキストやキャプションを設定したり、コンテンツに画像やギャラリーを挿入したりできます。気軽に試せる画像編集ツールも用意されています。
- コメント機能を標準搭載 ブログはいわば自分の家であり、コメントは友人やフォロワーがコンテンツに関わるための場所です。WordPressのコメント機能は、議論の場を作り、そのやり取りを管理するために必要なものをすべて備えています。
- 検索エンジン最適化済み WordPressは、導入した状態ですでに検索エンジン向けに最適化されています。さらに細かくSEOを調整したい場合は、多数のSEOプラグインが対応してくれます。

それでは、さっそく作り始めましょう。

インストール

```
composer create-project laravel/laravel wp-laravel
```

アプリケーションにアクセスしてみましょう

```
php artisan serve
```

データベースの設定が必要です。この記事ではSQLiteを使います。

```
touch database/database.sqlite
```

.env を更新する必要があります

```
DB_CONNECTION=sqlite
DB_DATABASE=database/database.sqlite
DB_FOREIGN_KEYS=true
```

[参考](https://laravel.com/docs/9.x/database#sqlite-configuration)

最初のデータベース変更をマイグレーションしましょう

```
php artisan migrate
```

これで、データベースの変更を次のように確認できます

```
php artisan db:show
```

すべてのマイグレーションが適用されました。

次に、WordPressの投稿のようなものを用意します。WordPressのテーブルを見たことがあれば、wp\_postsというテーブルがあることをご存じでしょう。これを作ってみましょう

```
php artisan make:model Post
```

これにより、Postというモデルが作成されます。次に、データベースの変更を適用するマイグレーションを作成しましょう。

```
php artisan make:migration "create posts table"
```

次に、マイグレーションファイルを開いて、コードを次のように更新します。

```
$table->string('title');
$table->string('slug')->nullable();
$table->text('excerpt')->nullable();
$table->longText('content')->nullable();
$table->dateTime('published_at')->nullable();
$table->foreignIdFor(\App\Models\User::class)->constrained();
```

**あわせて読みたい：** [サブスクリプションプランの制限をどう設計するか](https://sawirstudio.com/ja/posts/how-do-you-design-your-subscription-plan-restrictions/)、または[SawirStudioが提供する無料のWordPressプラグイン](https://sawirstudio.com/ja/projects/wordpress-plugins/)もご覧ください。
