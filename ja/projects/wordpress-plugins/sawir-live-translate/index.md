---
title: Sawir Live Translate | SawirStudio
url: https://sawirstudio.com/ja/projects/wordpress-plugins/sawir-live-translate/
description: Sawir Live Translate は、ページをそのままの場所で翻訳します。言語を選び、表示されている文章を自動翻訳して保存すれば、訪問者向けの言語切り替えも用意できます。
---

# Sawir Live Translate

![Sawir Live Translate のバナー](https://sawirstudio.com/assets/images/plugins/sawir-live-translate-banner.webp)

Sawir Live Translate は、ページをそのままの場所で翻訳します。言語を選び、表示されている文章を自動翻訳して保存すれば、訪問者向けの言語切り替えも用意できます。

[WordPress.org で無料入手](https://wordpress.org/plugins/sawir-live-translate/)

## 主な機能

サイトのどのページでも開き、管理バーの「Translate this page」をクリックして翻訳先の言語を選びます。表示されているすべての文字列を（無料の機械翻訳サービスで）自動翻訳することも、ページが更新される様子を見ながら手作業で翻訳を編集することもできます。

保存した翻訳はサーバー側でHTMLレスポンスに反映されます（SEOにも配慮：未翻訳の文章が一瞬表示されることがなく、head には hreflang の代替タグも出力されます）。訪問者は、右下のリンク式の切り替えから言語を変更できます。

URLには言語のパスプレフィックスを使います。たとえばスペイン語なら /es/about、元の言語なら /about です（元の言語はプレフィックスなしのまま正規URLになります）。言語は次の順で判定されます。/es/ のパス → ?lang=es → 保存済みのクッキー → ブラウザの Accept-Language → 元の言語。ルートURLやプレフィックスのないURLでは、リダイレクトせずに訪問者のブラウザの言語を自動判定します（共有キャッシュが正しく動作するよう、Vary ヘッダーを送信します）。

1. 「設定」→「Live Translate」で、元の言語と提供する言語にチェックを入れます。 2. フロントエンドのページを開き、管理バーの「🌐 Translate this page」をクリックします。 3. 言語を選んで「Auto-translate empty」を押し（または文字列を編集して）、「Save」を押します。 4. 訪問者は右下の切り替えを使います。各言語にはそれぞれ /xx/ のURLがあります。

注意：言語のパスプレフィックスは、ルートドメインへのインストールを前提としており、スラッグが2文字の言語コードになっているトップレベルのページがないことが条件です。

## よくある質問

### 他の選択肢と比べてどうですか？

WPML、Polylang、TranslatePress などの多言語フレームワークは、翻訳したコンテンツをサイト構造の中心に据え、言語ごとの編集ワークフロー、翻訳サービス、ページビルダーや WooCommerce との深い連携を備えています。Sawir Live Translate は、あえてもっと軽く作ってあります。ページを開き、表示されている文章を（自動または手作業で）翻訳し、言語プレフィックス付きのURLと hreflang タグで配信します。メニューや商品、編集ワークフローまで翻訳した本格的な多言語サイトを作るなら、フレームワークをお使いください。今日の午後のうちに既存のページを別の言語で読めるようにしたい、という場合には、このプラグインが向いています。

### 言語を切り替えても何も変わらない

ほとんどの場合、その言語の翻訳がまだ保存されていないことが原因です。「設定」→「Live Translate」→「Saved translations」を確認してください。その言語が一覧にない（または 0 strings と表示されている）場合は、もう一度ページエディターを開いて自動翻訳を行い、「Save」をクリックします。エディター内のライブプレビューは、保存するまで反映されません。

### 翻訳されない文字列がある

ページ内エディターが認識できるのは、開いているページ上の文字列だけです。ほかのページにある文章は、そのページも翻訳する必要があります。サイト全体のスキャンを実行する方法もあります。スライダーやスクリプトがページ読み込み後に挿入する文字列は、「Edit all site strings」から手動で入力が必要になる場合があります。

### コンテンツを編集したら翻訳が消えた

翻訳は、元の文章と完全に一致するものに対応づけられます。文章を書き換えた場合は、新しい文章を一度翻訳して保存してください。

## あわせて使いたいプラグイン

- [Sawir Insights](https://sawirstudio.com/ja/projects/wordpress-plugins/sawir-insights/)：ヒートマップ、セッション録画、アンケート、ファネル、フォーム分析で訪問者がサイトをどう使っているかを確認でき、すべてのデータを自社のサーバー内で管理できます。
- [Sawir Media Folders](https://sawirstudio.com/ja/projects/wordpress-plugins/sawir-media-folders/)：WordPress のメディアライブラリにフォルダーを追加します。必要な数だけ作成でき、ドラッグ＆ドロップで整理できます。

[無料の WordPress プラグイン](https://sawirstudio.com/ja/projects/wordpress-plugins/)の一覧はこちらです。

- **価格** 無料
- **バージョン** 1.1.0
- **動作要件** WordPress 5.8、PHP 7.4
- **ライセンス** GPL-2.0-or-later
- **WordPress.org** <https://wordpress.org/plugins/sawir-live-translate/>

私たちの話はここまでにして、あなたのお話を聞かせてください。

## どのような ご相談でしょうか？

[プロジェクトを始める](https://sawirstudio.com/ja/contact/)
