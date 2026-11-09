---
title: サイトをエージェント対応にするために行ったこと | SawirStudio
url: https://sawirstudio.com/ja/posts/making-our-sites-agent-ready/
description: 製品サイトを AI エージェントが読めるようにした方法：全ページの Markdown 版、Link ヘッダー、llms.txt、robots.txt の Content Signals、API カタログ、Agent Skills。
---

# サイトをエージェント対応にするために行ったこと

ウェブサイトへのアクセスのうち、人間ではないものが増えています。誰かの質問に答えるためにページを読んだり、誰かに代わって製品を使おうとしたりする AI エージェントです。今週、当サイトと各製品のサイト（[SawirPost](https://sawirstudio.com/ja/projects/sawirpost/)、[Sawirly](https://sawirstudio.com/ja/projects/sawirly/)、[Rindraw](https://sawirstudio.com/ja/projects/rindraw/)、[Math Daily](https://sawirstudio.com/ja/projects/math-daily/) など）を見直し、エージェントがより読みやすく、使いやすくなるようにしました。

チェックリストには [isitagentready.com](https://isitagentready.com) を使いました。サイトをレベルごとに評価してくれるサービスです。各ステップで何が必要だったかをご紹介します。

## すべてのページに Markdown 版を用意

エージェントが HTML を読むと、数段落の本文にたどり着くまでにナビゲーションやスクリプト、フッターを延々と読み進めることになります。そこで、当サイトのすべてのページに、`index.md` という Markdown 版を並べて置くようにしました。中身はそのページの主要なコンテンツだけです。HTML から小さなスクリプトで生成しているので、内容がずれることはありません。

エージェントは 2 通りの方法で取得できます。URL の末尾に `index.md` を付けるか、`Accept: text/markdown` を指定してリクエストする方法です。静的ファイルの前段に置いた小さな Cloudflare Worker がこのヘッダーを確認し、Markdown 版を返します。その際、おおよそのトークン数を `x-markdown-tokens` ヘッダーに載せています。ブラウザには、これまでどおり通常のページが表示されます。

```
curl -H "Accept: text/markdown" https://sawirstudio.com/shopify/
```

## Link ヘッダー

すべての HTML ページが、Markdown 版の場所を HTTP の `Link` ヘッダー（`rel="alternate"`）で示すようになりました。ホームページでは、`llms.txt` とサイトマップの場所も示しています。エージェントは URL を推測しなくても、1 回のリクエストですべてを見つけられます。

## llms.txt

`llms.txt` は、言語モデル向けに用意するサイトのプレーンテキストの目次です。1 行の概要と、各ページの説明付きの一覧で構成され、Markdown 版へのリンクが付いています。当サイトのものは、同じスクリプトでサイトマップから生成しています。

## Content Signals

`robots.txt` が示すのは、誰がクロールしてよいかであり、読み取った内容を何に使ってよいかではありません。それを補うのが [Content Signals](https://contentsignals.org/) です。当サイトの設定は 1 行で、3 つの用途すべてを許可しています。見つけてもらい、引用してもらいたいからです。

```
Content-Signal: ai-train=yes, search=yes, ai-input=yes
```

ここまでの 3 つのステップで、すべてのサイトがレベル 3 に到達しました。

## API カタログと Agent Skills

SawirPost と Sawirly には API があるため、さらに先へ進みました。それぞれ `/.well-known/api-catalog`（RFC 9727）で API カタログを公開しています。OpenAPI の定義とドキュメントを指し示す、小さな JSON ファイルです。サイトにたどり着いたエージェントは、ページを 1 つも読まなくても API を見つけられます。

あわせて Agent Skills も公開しています。製品でよくある作業の進め方を、エージェント向けに短くまとめた手順書です。これで両方ともレベル 4 に到達しました。

## ここで止めた理由

レベル 5 では、エージェントが OAuth のエージェント登録フローを通じて、自力でサインアップと認証を行えることが求められます。メールアドレスとパスワードでの登録方法を説明するページでは、条件を満たしません。これは本格的な機能開発であり、セキュリティ面の検討も欠かせません。そのため、急がずにじっくり考えているところです。

## やる価値はあったか

静的サイトであれば、大半の作業は半日ほどで終わります。スクリプト 1 本、小さな Worker 1 つ、そして `robots.txt` への 1 行です。サイトを訪れる人にとっては、何も変わりません。今後は、当サイトのようなサイトを読んだエージェントが、より多くの質問に答えるようになると考えています。整った Markdown ページを用意しておくことが、正しく引用してもらうための一番の近道です。

お客様向けにも、同様の対応を[ローンチ＆グロース](https://sawirstudio.com/ja/launch-and-growth/)の一環として行っています。

**あわせてお読みください：** [Abilities API で WordPress プラグインを AI エージェントから使えるようにする](https://sawirstudio.com/ja/posts/wordpress-abilities-api-for-plugins/)、[ホームページの「投げられるアプリアイコン」について](https://sawirstudio.com/ja/posts/throwable-app-icons-matter-js/)
