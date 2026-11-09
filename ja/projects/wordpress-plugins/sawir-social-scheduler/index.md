---
title: Sawir Social Scheduler | SawirStudio
url: https://sawirstudio.com/ja/projects/wordpress-plugins/sawir-social-scheduler/
description: wp-admin から一度書くだけで、27 のネットワークに予約投稿。チャンネルごとの編集、スレッド、メディア、プレビューに対応。トークンはお客様のサーバー内に保管されます。
---

# Sawir Social Scheduler

![Sawir Social Scheduler バナー](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-banner.webp)

Sawir Social Scheduler なら、wp-admin で一度投稿を書くだけで、27 のソーシャルネットワークに予約投稿できます。チャンネルごとの編集、スレッド、メディア、プレビューにも対応しています。アカウントのトークンはお客様のサーバー内に保管されます。

[WordPress.org で無料入手](https://wordpress.org/plugins/sawir-social-scheduler/)

## 主な機能

Sawir Social Scheduler は、WordPress をソーシャルメディアのスケジューラーに変えるプラグインです。各プラットフォームが発行するトークンでチャンネルを接続し、投稿を書いて、チャンネルを選び、日時を設定して公開します。送信はすべてバックグラウンドで実行され、完了するとリンク、または各プラットフォームのエラーメッセージが表示されます。

間に介在するホスティングサービスはありません。認証情報はサイトの salt で暗号化され、サーバーの外に出ることはありません。

### チャンネル

チャンネルはいくつでも追加でき、1 つのプラットフォームに複数のアカウントを登録することもできます。各チャンネルには**接続テスト** ボタンがあり、アカウント名とアバターを取得して、トークンが実際に使えることを事前に確認できます。

**ソーシャル：** X (Twitter)、Facebook ページ、Instagram、Threads、LinkedIn ページ、Bluesky、Mastodon、TikTok、YouTube、Pinterest、Tumblr、Dribbble、Google ビジネス プロフィール

**コミュニティ：** Telegram、Discord、Slack、Reddit、Lemmy、VK、Kick、Twitch、Whop

**ブログ・ニュースレター：** WordPress、Medium、Dev.to、Hashnode、Listmonk

クライアントやブランドごとに、入れ子にできるグループでチャンネルをまとめられます。グループ全体もワンクリックで選択できます。

### 投稿

- **1 つの本文を、複数のチャンネルへ。** チェックボックスでチャンネルを選びます。変更していないチャンネルには、書いたとおりの内容が投稿されます。
- **チャンネルごとにカスタマイズ。** 任意のチャンネルで「カスタマイズ」にチェックを入れると、そのチャンネル専用のパート（テキスト、メディア、遅延時間）を設定できます。内容は投稿から自動で入力されるので、コピーを編集する形になります。ほかのチャンネルには、書いたとおりの内容が投稿されます。
- **スレッド。** パートを追加すると、X、Threads、Bluesky、Mastodon、Telegram でスレッド投稿ができ、各パートにそれぞれメディアを付けられます。スレッドのないプラットフォームには、パートを連結して送信します。
- **メディア。** 画像や動画は一度添付するだけです。アルバム、カルーセル、リール、再開可能な動画アップロードなど、各プラットフォームが対応する形式で送信されます。
- **予約。** WordPress 標準の「公開日時」を使うことも、チャンネルごとに個別の時刻を指定することもできます。
- **投稿時間帯。** チャンネルごとに、いつも投稿する時間帯を登録できます。「次の空き枠」を選ぶと、未使用の次の枠が自動で入ります。カレンダーには今後の空き枠が表示され、クリックして投稿を作成できます。API では `"date": "next"` も指定できます。
- **プレビュー。** 実際のプラットフォームに近い見た目のライブプレビューを、チャンネルごとに表示します。文字数制限やメディアのルールも反映されるため、投稿される内容を正確に確認できます。
- **コメントと遅延。** スレッドのないプラットフォームでは、追加のパートが投稿へのコメントとして投稿されます。どのパートにも、直前のパートから指定した分数だけ待つ設定ができます。
- **ブログ記事のシェア。** WordPress 標準の投稿エディターにあるボックスから、記事の公開時に、専用のテキストとメディアを付けて各チャンネルへ投稿できます。
- **複製。** 投稿を、チャンネル、パート、個別設定ごと新しい下書きとして複製できます。下書きをテンプレートとして残しておくこともできます。
- **メンション。** `@handle` と入力すると、X、Instagram、Threads、Mastodon、Telegram、Reddit、Bluesky でネットワーク側が解決します。Bluesky ではリンクとハンドルが本物のファセットになります。LinkedIn の会社ページは `@[Name](org:12345)` の形式で、ID にはページ URL の ID を使います。Facebook ページは `@[12345]` の形式です。
- **繰り返し。** 定番の投稿を N 日または N 週間ごとに再送信でき、終了日または回数も指定できます。実行ごとに独立した投稿として、送信と統計が記録されます。X、Threads、Bluesky は同一テキストの投稿を受け付けないため、テキストに `{n}` を入れて実行回数の番号を付けられます。

### カレンダー

1 件の投稿は、チャンネルごとに 1 件ずつの送信に分かれ、カレンダー上では状態（予約済み、送信済み、失敗、保留、下書き、処理中）ごとに色分けされたカードとして表示されます。カードをドラッグすると、そのチャンネルの時刻を変更できます。カードをクリックすると、エラー内容、再試行ボタン、公開リンク、エンゲージメントの数値が表示されます。空き枠をクリックすれば、その時刻に投稿を作成できます。月、週、日、リストの表示に対応し、グループやステータスで絞り込めます。

### 分析

いいね、コメント、シェア、閲覧数、インプレッション数は、各送信の 1 時間後、1 日後、1 週間後に自動で収集され、必要に応じて手動でも取得できます。対応しているのは 16 のプラットフォームです。フォロワー数はチャンネルごとに毎日記録され、30 日間のスパークラインで確認できます。数値は、カレンダーのカード、投稿の「分析」ボックス、チャンネル一覧に表示されます。「分析」ページでは、直近 7 日、30 日、90 日の実績を、その前の同じ期間と比較してまとめています。送信した投稿数、エンゲージメント、リーチ、フォロワー数とその増減、フォロワー数のグラフ、チャンネル別のエンゲージメント、人気の投稿を確認できます。

### クライアントの承認

「承認されるまで送信を保留」にチェックを入れ、プレビューのリンクをコピーして、クライアントに送ります。クライアントはログイン不要で、各チャンネルでの実際の見え方を確認し、「承認」または、メモ付きの「修正を依頼」をクリックします。承認されると送信が実行されます。判断はすべて履歴として残ります。

X、YouTube、TikTok への大容量動画のアップロードは、バックグラウンドジョブをまたいで再開できる短い単位に分けて実行されます。そのため、ホスティングの時間制限にかかりにくく、ファイル全体をメモリに読み込むこともありません。

### REST API

管理画面でできることはすべて、自動化、モバイルアプリ、ほかのサイト向けに `/wp-json/sawir-social/v1` から利用できます。認証にはアプリケーションパスワードを使います。権限はプラグインの権限設定に従い、チャンネルの認証情報が返されることはありません。

- `GET /platforms`, `GET /channels`, `GET /channels/{id}`, `POST /channels/{id}/test`
- `GET|POST /posts`, `GET|PATCH|DELETE /posts/{id}`, `POST /posts/{id}/duplicate`, `GET /posts/{id}/analytics`
- `GET /sends`, `GET /sends/{id}`, `POST /sends/{id}/retry`, `POST /sends/{id}/stats`

投稿の作成は次のように行います。`{"title","content","status":"future","date":"2026-12-25 09:00","media":[12],"parts":[{"text","media","delay"}],"channels":[{"id":38,"parts":[{"text","media","delay"}],"date":"…"}],"groups":["Client A"],"hold":false}`。チャンネルの `parts` は、そのチャンネルでは投稿の内容に置き換わります。`text` は、パートが 1 つの場合の省略形です。

### AI エージェント

同じ操作は WordPress Abilities API にも `sawir-social/list-channels`、`create-post`、`update-post`、`list-posts`、`list-sends`、`retry-send`、`test-channel`、`get-analytics`、`list-platforms` として登録されており、入力と出力のスキーマは型付けされています。WordPress MCP アダプター経由で接続したエージェントからはツールとして認識され、`/wp-json/wp-abilities/v1/abilities/{name}/run` から呼び出すこともできます。Abilities API は WordPress 6.9 以降に含まれています。それより古いバージョンではプラグインが同梱しているため、プラグインが動作する環境ならどこでも利用できます。

### 動作要件

送信は、サイトに Action Scheduler がある場合（WooCommerce などの多くのプラグインが提供しています）はそれを使い、ない場合は WP-Cron を使ってバックグラウンドで実行されます。WordPress 組み込みの cron では、サイトにアクセスがあったときにしか処理が進みません。すぐに送信したい場合は、システム cron を追加してください。ご利用のサイトに合ったコマンドが、カレンダー画面に表示されます。

URL でメディアを取得するプラットフォーム（Instagram、Threads、Facebook、Pinterest、Telegram）に送るには、メディアがインターネットから到達できる必要があります。ローカルの開発サイトからは、これらのプラットフォームに画像を渡せません。

### クレジット

FullCalendar（MIT）、WordPress Abilities API ライブラリ（GPL）、Simple Icons（CC0）のブランドマークを同梱しています。ほかのプラグインが Action Scheduler を提供している場合は、それを使用します。

## スクリーンショット

![Sawir Social Scheduler スクリーンショット 1](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-screenshot-1.webp)

![Sawir Social Scheduler スクリーンショット 2](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-screenshot-2.webp)

![Sawir Social Scheduler スクリーンショット 3](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-screenshot-3.webp)

![Sawir Social Scheduler スクリーンショット 4](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-screenshot-4.webp)

![Sawir Social Scheduler スクリーンショット 5](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-screenshot-5.webp)

![Sawir Social Scheduler スクリーンショット 6](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-screenshot-6.webp)

## よくある質問

### トークンはどこで取得できますか？

各プラットフォームの開発者設定で取得できます。チャンネルのフォームには、必要な認証情報の名称（Facebook ページのアクセストークン、Telegram のボットトークン、Bluesky のアプリパスワードなど）が明記されています。長期間有効なトークンがおすすめです。Meta と LinkedIn のトークンは約 60 日で期限切れになるため、再度貼り付ける必要があります。

### Telegram のテストで「chat not found」と表示されます

プライベートチャンネルには @username がありません。チャット ID を空欄のまま「接続テスト」をクリックしてください。ボットが追加されているすべてのチャットが、貼り付け用の数値 ID とともに一覧表示されます。

### 送信が「処理中…」のまま止まっているのはなぜですか？

バックグラウンドキューがまだ実行されていません。サイトにアクセスするか、カレンダー画面に表示されるシステム cron を設定してください。

### ローカルサイトから Instagram に動画を投稿できますか？

できません。Instagram は URL でメディアを取得するため、公開アドレスが必要です。それ以外の機能は、ローカル環境でもすべて動作します。

## あわせて使いたいプラグイン

- [Sawir Insights](https://sawirstudio.com/ja/projects/wordpress-plugins/sawir-insights/)：ヒートマップ、セッション録画、アンケート、ファネル、フォーム分析で訪問者がサイトをどう使っているかを確認でき、すべてのデータを自社のサーバー内で管理できます。
- [Sawir Live Translate](https://sawirstudio.com/ja/projects/wordpress-plugins/sawir-live-translate/)：ページをその場で翻訳します。言語を選んで表示中の内容を自動翻訳し、保存すれば、訪問者向けの言語切り替えも設置できます。

[無料の WordPress プラグイン](https://sawirstudio.com/ja/projects/wordpress-plugins/)の一覧はこちらです。

- **価格** 無料
- **バージョン** 0.1.0
- **動作要件** WordPress 5.8、PHP 7.4
- **ライセンス** GPL-2.0-or-later
- **WordPress.org** <https://wordpress.org/plugins/sawir-social-scheduler/>

私たちの話はここまでにして、あなたのお話を聞かせてください。

## どのような ご相談でしょうか？

[プロジェクトを始める](https://sawirstudio.com/ja/contact/)
