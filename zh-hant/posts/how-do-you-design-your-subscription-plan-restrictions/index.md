---
title: Laravel 訂閱方案限制 | SawirStudio
url: https://sawirstudio.com/zh-hant/posts/how-do-you-design-your-subscription-plan-restrictions/
description: 如何在 Laravel 中落實訂閱方案的使用限制，從 controller 裡的簡單檢查，到 Gate 與 Policy。
---

# 如何設計訂閱方案的使用限制

我通常會把使用者的訂閱方案存放在 `users` 資料表中，例如用 `plan_id` 來對應訂閱方案。

我們先來搭個範例環境……

我們的訂閱方案 subscription plans.php

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

接著就可以在 controller 中這樣使用（這裡假設只有已登入的使用者才能呼叫這個 endpoint）

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

這是可行的做法，對大多數應用程式來說，用它來檢查建立數量上限已經夠用了。

不過，就像您提到的，還有另一種「更集中管理」的做法。

使用 Gate 與 Policy。

我們先把前面的程式碼重構成 Gate

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

接著在 controller 中可以改寫成這樣

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

這樣應該已經能滿足您的需求。不過，會不會還有更規範的做法？

您說得沒錯，就是使用 Policy。

一般的說法是，處理領域中的「資源」或 model 時，最適合用 Policy，就像我們這個例子中的 List model。

而 Gate 則比較適合處理「動作」。

不過，我個人的看法是，只要可以，就盡量用 Policy。除了可以順理成章地把程式拆成不同檔案之外，也不會把 `AuthServiceProvider` 弄得雜亂（順帶一提，除非真的有必要，否則我不太喜歡去改動 Service Provider）。

我們把前面的程式碼重構成 Policy

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

有注意到這些方法的命名和 resource controller 很相似嗎？我想這是 Laravel 團隊刻意的設計，值得稱讚。

然後在 controller 中就可以直接這樣使用

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

這裡的 `List::class` 只是用來幫助 Laravel 推斷要使用哪一個 Policy 類別。

總的來說，這樣做可以避免一個問題：您可能會在 10 個不同的 controller 裡建立清單。我們只要在各處加上 `$this->authorize()`，就大功告成了。

**延伸閱讀：** [如何打造自己的內容管理系統（像 WordPress 那樣）](https://sawirstudio.com/zh-hant/posts/how-to-build-your-own-content-management-system-like-wordpress/)，或是到我們的 [WordPress 外掛](https://sawirstudio.com/zh-hant/projects/wordpress-plugins/)看看 Sawir Renew，這是一款適用於 WooCommerce 的免費訂閱外掛。
