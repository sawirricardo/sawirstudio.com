---
title: Shopify と Odoo 19 Community を連携する | SawirStudio
url: https://sawirstudio.com/ja/posts/connect-shopify-to-odoo-19-community/
description: オープンソースの OCA Shopify コネクタを使って、Shopify ストアを Odoo 19 Community に連携する方法を解説します。Dev Dashboard のアプリ、スコープ、queue_job、Webhook、オンボーディングまで扱います。
---

# Shopify と Odoo 19 Community を連携する手順

私たちは [Odoo 用の Shopify コネクタ](https://sawirstudio.com/ja/projects/odoo-shopify-connector/)を開発し、Odoo Community Association（OCA）に提供しました。無料のオープンソース（AGPL-3）で、Odoo 19 Community に対応しています。ここでは、お客様向けに実際に行っているセットアップの流れを、最初から最後までご紹介します。

商品、在庫、顧客、注文、返品、返金、出荷の情報を、1 店舗でも複数店舗でも同期できます。会計とポイントオブセールス（POS）は別のモジュールで、後から追加できます。

## 始める前に

- 公開 HTTPS でアクセスできる Odoo 19 Community。Shopify がここへ Webhook を送信し、商品画像もここからダウンロードします。
- Shopify ストア、およびそのストアを所有する Shopify 組織の管理者権限。
- 自社で作成するアプリで保護された顧客データを利用できる Shopify プラン。Basic プランでは利用できず、この権限がないと Shopify が一部の注文 Webhook を拒否し、注文と顧客のインポートが失敗します。

## 1. Shopify Dev Dashboard でアプリを作成する

2026 年 1 月 1 日以降、Shopify ではストアの管理画面から従来の「カスタムアプリ」を作成できなくなりました。ストアを所有する組織の **Dev Dashboard** でアプリを作成してください。これは重要なポイントです。Odoo はクライアント認証情報グラントでアクセストークンを取得しますが、これはアプリとストアが同じ組織に属している場合にのみ機能します。

アプリのバージョンを入力するときは、次のように設定します。

- **アプリ URL：** `https://shopify.dev/apps/default-app-home`。このコネクタには Shopify 内のページはありません。
- **Shopify 管理画面にアプリを埋め込む：** オフ。
- **従来のインストールフローを使用する：** オフ。これにより、アプリをインストールすると選択したスコープが付与されます。
- **Webhook：** 設定しません。Odoo がオンボーディング中に自動で作成します。
- **Webhook API バージョン：** `2026-01`。コネクタが使用するバージョンです。
- **スコープ：** `read_customers`、`write_customers`、`read_draft_orders`、`read_fulfillments`、`read_inventory`、`write_inventory`、`read_locations`、`read_merchant_managed_fulfillment_orders`、`write_merchant_managed_fulfillment_orders`、`read_orders`、`read_products`、`write_products`、`read_returns`。60 日より前の注文も取り込みたい場合は、`read_all_orders` を追加してください。

バージョンをリリースし、ストアにアプリをインストールしたら、アプリの設定ページからクライアント ID とクライアントシークレットをコピーします。

2026 年より前に作成した従来のカスタムアプリをお持ちの場合も、そのまま使えます。クライアント ID は空欄のままにして、代わりに Admin API アクセストークンと API シークレットキーを使用してください。

## 2. Odoo にモジュールをインストールする

OCA の `queue_job` と [connector-shopify](https://github.com/OCA/connector-shopify) リポジトリをアドオンパスに追加し、**Shopify Connector** をインストールします。処理はバックグラウンドジョブとして実行されるため、`queue_job` をサーバー全体で読み込み、Shopify のチャネルごとに処理容量を設定する必要があります。

```
[options]
server_wide_modules = base,rpc,web,queue_job

[queue_job]
channels = root:4,root.shopify.import:2,root.shopify.export:1,root.shopify.webhook:1
```

次に、システムパラメータ `web.base.url` に Odoo の公開 HTTPS アドレスを設定し、`web.base.url.freeze` を `True` にします。こうしないと、管理者が別の場所からログインしたときにアドレスが変わり、Shopify の Webhook が誤った宛先に送られてしまいます。

## 3. オンボーディングを実行する

チームのメンバーに **Shopify Manager** グループを付与してから、**Shopify > Configuration > Instances** を開き、**Start Onboarding** をクリックします。

1. `yourstore.myshopify.com` のドメイン、クライアント ID、クライアントシークレットを入力し、接続をテストします。Odoo がトークンを取得し、24 時間ごとに自動で更新します。
2. スコープを確認します。不足しているものがあれば、新しいアプリバージョンに追加してリリースし、ストア側で承認してください。
3. Shopify のロケーションを Odoo の倉庫に対応付け、Odoo が生成する税のマッピングを確認して、同期ポリシーを選択します。
4. Webhook をインストールし、最初のインポートをキューに登録します。

注文や出荷を自動で動かす前に、決済ゲートウェイと配送業者のマッピングを確認してください。後から起こる想定外のトラブルの多くは、誰も見ていなかったマッピングが原因です。

## 4. 運用状況を確認する

**Shopify > Dashboard** では、同期の状態、ジョブの滞留状況、直近のエラー、Webhook の稼働状況を確認できます。このコネクタは、すべての Webhook が届くことを前提にしていません。Shopify 側で変更があったのに Odoo が受け取れなかったものは、スケジュールされたジョブが改めて取り込みます。エラーが発生したら、原因を修正してから **Retry Failed Shopify Bindings** を実行してください。ドメインや認証情報を変更した後は、**Check Webhooks** を実行します。

## 次のステップ

**Shopify Connector Accounting** を追加すると、Shopify Payments の入金、手数料、ディスピュートを仕訳として取り込めます（入金の突合が難しい理由は、[別の記事](https://sawirstudio.com/ja/posts/why-shopify-payouts-dont-match-your-books/)で解説しています）。店舗での販売には **Shopify Connector POS** をご利用ください。

セットアップや、お客様の業務に合わせた設定、機能の拡張をご希望でしたら、[私たちが対応します](https://sawirstudio.com/ja/shopify/)。

**あわせてお読みください：** [Shopify の入金が帳簿と合わない理由](https://sawirstudio.com/ja/posts/why-shopify-payouts-dont-match-your-books/)、または [Shopify の 2026-10 API：連携で何が動かなくなるか](https://sawirstudio.com/ja/posts/shopify-2026-10-api-what-breaks/)
