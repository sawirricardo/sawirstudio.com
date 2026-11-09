---
title: Shopify 2026-10 API：何が壊れるか | SawirStudio
url: https://sawirstudio.com/ja/posts/shopify-2026-10-api-what-breaks/
description: 連携を壊す可能性のある、バージョン 2026-10 の Shopify Admin API の変更点と対処方法：配送料金、ドラフト注文の割引、税金、バーコードなど。
---

# Shopify 2026-10 API：連携で壊れる可能性のある変更点

Shopify は 2026 年 10 月 1 日に、API のバージョン 2026-10 をリリースしました。内容の大半は新機能で、必要になるまで無視して構いません。ただし、2026-07 では問題なく動いていた連携が、気づかないうちに壊れてしまう変更がいくつかあります。バージョンを上げる前に、1 時間ほどかけて確認しておく価値があります。

ここでは、[Odoo 向け Shopify コネクタ](https://sawirstudio.com/ja/projects/odoo-shopify-connector/)を含む自社の連携に対して、私たちが実際に確認している項目をまとめました。リリースノートの全項目ではなく、動作が変わるものだけを挙げています。

## 気づかないうちに壊れる変更

### 新しい配送サービスは、そのままではチェックアウトに料金が表示されない

`carrierServiceCreate`（REST では `POST carrier_services.json`）で配送サービスを作成しても、サービスが登録されるだけになりました。これまでのように、料金が General 配送プロファイルへ自動で追加されることはありません。アプリやオンボーディング用スクリプトで配送サービスを作成し、お客様に料金が表示されることを前提にしている場合、配送プロファイルへ自分でサービスを追加しない限り、料金は表示されません。

### 配送先住所を変更すると税金が再計算される

未発送の注文の配送先住所を `orderUpdate` または REST の注文更新エンドポイントで更新すると、税金が再計算されるようになりました。動作としては正しいのですが、ERP やサポートツールから住所を変更している場合、注文の合計金額が意図せず変わることがあります。更新後に `taxLines`、`totalTaxSet` と各合計金額をもう一度読み取り、支払い済みまたは返金済みの金額に差が出た場合は精算してください。

### ドラフト注文の割引警告から `priceRule` がなくなった

`DraftOrderDiscountNotAppliedWarning.priceRule` と従来の `PriceRule` 型は廃止されました。`draftOrderCalculate`、`draftOrderCreate`、`draftOrderUpdate` でドラフト注文を作成し、割引が適用されなかった理由を読み取っている場合は、代わりに `discountTitle` と `discountCode` を選択してください。

### 不正なメタフィールドのフィルターは、はっきりエラーになる

フィルターできないメタフィールドで商品クエリを絞り込むと、これまではその条件が無視され、すべての商品が返ってきました。今後はエラーが返ります。この方が望ましい動作ですが、何か月も「全商品」を返し続けていた同期処理は、バージョンを切り替えた当日から失敗するようになります。絞り込みに使っているメタフィールド定義が、フィルターと使用している演算子に対応しているかを確認してください。

### 1 つのバリエーションに複数のバーコードを設定できる

`ProductVariant.barcodes` は新しいコネクションで、種類付きのバーコードを最大 20 件まで持てます。従来の `ProductVariant.barcode` は非推奨となり、最初の 1 件だけを返します。倉庫システムや POS がバーコードで商品を照合している場合は、全件を読み取るようにしてください。

### 在庫関連のエラーが 1 つなくなった

`ITEM_NOT_STOCKED_AT_LOCATION` は、在庫のミューテーションと `InventorySetQuantitiesUserErrorCode` から削除されました。このエラーで分岐しているコードは、今後は実行されない不要なコードになります。後で誰かが頼ってしまわないよう、削除しておいてください。

## コードを検索して確認したい、細かな変更

- **顧客セグメント。** `= true` と `= false` に代わって `MATCHES` と `NOT MATCHES` を使います。また、`12_months_ago` のような名前付きの日付は非推奨となり、`-12m` のようなオフセット指定が推奨されます。
- **SMS の同意。** フラットな SMS 同意フィールドは非推奨です。`CustomerPhoneNumber.smsMarketingConsent` を読み取ってください。
- **ロールアウト内の割引。** 割引がロールアウトの一部である場合、`startsAt` と `endsAt` だけでは有効かどうかを判断できません。`Discount.rollouts` を読み取るには `read_rollouts` スコープが必要です。
- **フルフィルメントのステータス。** `FULFILLMENT_NOT_REQUIRED` が新しい値として追加されました。すべてのステータスを対応付けている `switch` では、デフォルトのケースに入ることになります。
- **Customer Account API。** `Customer.lastIncompleteCheckout` と `Checkout` 型は、代替なしで削除されました。
- **POS UI 拡張機能。** `applyCartDiscount` などのカート操作は、POS が変更を承認しない場合に拒否されるようになりました。また、`session.currentSession.staffMemberId` は削除されました。

## 2026-10 とは関係のない 2 つの期限

公開アプリを運用している場合は、有効期限のないオフラインアクセストークンを、2027 年 1 月 1 日までに有効期限付きのものへ置き換える必要があります。カスタムアプリと、販売者が Dev Dashboard で作成したアプリは対象外です。また、オンラインストアのスクリプトタグは 2027 年 3 月 1 日に動作しなくなります。ストアへの影響については、[別の記事](https://sawirstudio.com/ja/posts/shopify-script-tags-stop-march-2027/)にまとめています。

## 私たちのアップグレード方法

API バージョンは 1 か所で固定し、新しいバージョンで開発ストアに対して連携を動かし、[2026-10 の変更履歴](https://shopify.dev/changelog?api_version=2026-10)を「Action required」で絞り込んで読みます。2026-10 は少なくとも 2027 年 10 月 1 日までサポートされるので時間はありますが、警告メールが届き始めるまで先延ばしにできるほどの余裕はありません。

ストアを Odoo などのシステムと連携していて、この作業を他の人に任せたい場合は、まさに[私たちの仕事](https://sawirstudio.com/ja/shopify/)です。

**あわせてお読みください：** [Shopify のスクリプトタグが 2027 年 3 月 1 日に動作しなくなります。ストアオーナーがすべきこと](https://sawirstudio.com/ja/posts/shopify-script-tags-stop-march-2027/)、または[Shopify と Odoo 19 Community の連携：手順ごとの解説](https://sawirstudio.com/ja/posts/connect-shopify-to-odoo-19-community/)。
