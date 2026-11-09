---
title: Laravelでのサブスクリプションプラン制限 | SawirStudio
url: https://sawirstudio.com/ja/posts/how-do-you-design-your-subscription-plan-restrictions/
description: Laravelでサブスクリプションプランの上限を実装する方法を、コントローラーでの簡単なチェックから、GateとPolicyまで解説します。
---

# サブスクリプションプランの制限をどう設計するか

私はたいてい、ユーザーのサブスクリプションプランを `users` テーブルに保存します。たとえば、サブスクリプションプランを参照する `plan_id` です。

まずは土台を作りましょう。

サブスクリプションプラン（plans.php）

```
return [

	cheap_plan ⇒ [

		‘maximum_list’ ⇒ 20,

	],

	expensive_plan ⇒ [

		‘maximum_list’ ⇒ 50,

	],

];
```

コントローラーでは次のように使えます（このエンドポイントを使えるのはログイン済みのユーザーだけとします）。

```
public function create()
{
	$postCount = Auth::user()->loadCount('posts')->posts_count;

	if ($postCount > Auth::user()->plan->maximum_list) {
		abort(403, "You've reached your maximum list. Please upgrade your plan.");
	}

	// proceed to create the list
	// $list = List::create(...)

	// return redirect()->route('lists.show', ['list' => $list]);
}
```

これは有効な方法で、多くのアプリケーションでは、作成数の上限をチェックする目的を十分に果たすと思います。

ただ、ご指摘のとおり、もう少し「一元化された」方法もあります。

GateとPolicyを使う方法です。

先ほどのコードをGateにリファクタリングしてみましょう。

```
// AuthServiceProvider.php

class AuthServiceProvider extends ServiceProvider
{
	//

	public function boot()
	{
		$this->registerPolicies();

		Gate::define('create-list', function($user) {
			$postCount = Auth::user()->loadCount('posts')->posts_count;

			if ($postCount > Auth::user()->plan->maximum_list) {
				return Response::deny("You've reached your maximum list. Please upgrade your plan.");
			}

			return Response::allow();
		});
	}
}
```

コントローラーは次のように書き換えられます。

```
public function create()
{
	$this->authorize('create-list');
	// or Gate::allows('create-list');
	// or Gate::inspect('create-list'); // if you want more customized response.

	// proceed to create the list
	// $list = List::create(...)

	// return redirect()->route('lists.show', ['list' => $list]);
}
```

これでもう十分かもしれません。とはいえ、もっと作法に沿った方法はないでしょうか。

そのとおり、Policyを使う方法です。

Policyは、ドメイン内の「リソース」やモデルを扱うときに向いていると言われます。今回の例でいえば、Listモデルです。

一方、Gateは「アクション」を扱うときに向いています。

ただ、私の個人的な意見としては、可能な限りPolicyを使うのがよいと思います。ファイルを分ける口実ができるうえ、`AuthServiceProvider` が散らかりません（そもそも、本当に必要な場合を除いて、サービスプロバイダーはあまり触りたくありません）。

先ほどのコードをPolicyにリファクタリングしてみましょう。

```
class ListPolicy
{
	use HandlesAuthorization;

	public function create(User $user)
	{
		$postCount = Auth::user()->loadCount('posts')->posts_count;

		if ($postCount > Auth::user()->plan->maximum_list) {
			return Response::deny("You've reached your maximum list. Please upgrade your plan.");
		}

		return Response::allow();
	}

}
```

メソッド名がリソースコントローラーに対応しているのにお気づきでしょうか。これはLaravelチームによる意図的な設計だと思います。さすがです。

コントローラーでは、次のようにシンプルに使えます。

```
public function create()
{
	$this->authorize('create', List::class);
	// or Gate::allows('create-list', List::class);
	// or Gate::inspect('create-list', List::class); // if you want more customized response.

	// proceed to create the list
	// $list = List::create(...)

	// return redirect()->route('lists.show', ['list' => $list]);
}
```

`List::class` は、どのPolicyクラスを使うかをLaravelが判断するための目印にすぎません。

結果として、この方法なら、たとえば10個もの別々のコントローラーからリストを作成できてしまう、といった問題を避けられます。ここに `$this->authorize()` を置くだけで済みます。

**あわせて読みたい：** [WordPressのようなコンテンツ管理システムを自作する方法](https://sawirstudio.com/ja/posts/how-to-build-your-own-content-management-system-like-wordpress/)。また、WooCommerce向けの無料サブスクリプションプラグイン「Sawir Renew」は、[WordPressプラグイン](https://sawirstudio.com/ja/projects/wordpress-plugins/)のページでご覧いただけます。
