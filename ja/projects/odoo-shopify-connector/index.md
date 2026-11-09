---
title: Odoo 向け Shopify コネクタ | SawirStudio
url: https://sawirstudio.com/ja/projects/odoo-shopify-connector/
description: SawirStudio が開発し、Odoo Community Association（OCA）で管理されている、Odoo 19 Community 向けのオープンソース Shopify コネクタです。
---

# Odoo 向け Shopify コネクタ

![GitHub 上の OCA connector-shopify リポジトリ](https://sawirstudio.com/assets/images/projects/odoo-shopify-connector.webp)

## プロジェクトについて

Odoo Community Association（OCA）は、Odoo Community のユーザーの多くが利用しているオープンソースモジュールを管理する非営利団体です。Connector Shopify は、Shopify ストアと Odoo 19 Community をつなぐ OCA のコネクタで、AGPL-3 ライセンスで公開されています。

## 制作したもの

このコネクタは SawirStudio が開発し、2026 年 7 月に OCA へ提供しました。対象は、コアとなる Shopify Connector、Shopify Connector Accounting、Shopify Connector POS です。現在も開発を続けており、直近では Shopify の 2026-01 Admin API に対応させました。ほかの OCA コントリビューターもこのコネクタをベースに開発を進めており、eCommerce の商品メディア用モジュールもその一つです。

- **コマース連携。** 商品、在庫、顧客、注文、返品、返金、出荷を同期します。1 つの Odoo データベースで複数の Shopify ストアを扱うことができ、会社ごとに管理できます。
- **確実に動く設計。** 署名付き Webhook で素早く受信し、OCA の queue\_job で処理を確実に実行します。さらに定期的な照合で、取りこぼしや遅延したイベントを補正します。
- **日々の運用ツール。** 初期設定のサポート、ヘルスチェック、同期ダッシュボード、エラーからの復旧、設定可能なデータ保持期間を備えています。
- **会計。** Shopify Payments の入金、残高トランザクション、手数料、チャージバック、ギフトカードの負債、複数通貨の照合に対応しています。
- **店舗販売（POS）。** Shopify POS の売上を販売オーダーとして取り込み、各店舗から Odoo の倉庫へ振り分けます。Odoo 標準の POS アプリは必要ありません。

## 進め方

ほとんどの場合は正しい、という連携では会計には使えません。そこでこのコネクタは、Shopify の GraphQL Admin API を利用し、Webhook が届いたものとは決めてかからず、定期的に Shopify と突き合わせて自己チェックを行います。仕訳は Shopify の正確な金額をもとに作成し、貸借が一致しなければ転記されません。数字が合わない場合は、調整用の行をこっそり追加するのではなく、わかりやすいエラーを表示します。

お客様のストア向けに導入や機能拡張をご希望ですか？ [Shopify 連携](https://sawirstudio.com/ja/shopify/)をご覧ください。Odoo そのものについては、[Odoo とビジネスシステム](https://sawirstudio.com/ja/odoo/)をご覧ください。

**使用技術** Python、Odoo 19、Shopify GraphQL Admin API

- **プロジェクト** Odoo Community Association（OCA）
- **ライセンス** AGPL-3
- **ソースコード** <https://github.com/OCA/connector-shopify>

## その他の実績

[![MyVivaldy](https://sawirstudio.com/assets/images/projects/myvivaldy.webp)

### MyVivaldy](https://sawirstudio.com/ja/projects/myvivaldy/)
[![Koperasi HOB アプリケーション](https://sawirstudio.com/assets/images/projects/koperasi-hob-application.webp)

### Koperasi HOB アプリケーション](https://sawirstudio.com/ja/projects/koperasi-hob-application/)
[![DBTV アプリケーション](https://sawirstudio.com/assets/images/projects/dbtv-application.webp)

### DBTV アプリケーション](https://sawirstudio.com/ja/projects/dbtv-application/)

私たちの話はここまでにして、あなたのお話を聞かせてください。

## どのような ご相談でしょうか？

[プロジェクトを始める](https://sawirstudio.com/ja/contact/)
