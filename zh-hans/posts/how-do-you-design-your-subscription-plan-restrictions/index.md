---
title: Laravel 订阅套餐限制 | SawirStudio
url: https://sawirstudio.com/zh-hans/posts/how-do-you-design-your-subscription-plan-restrictions/
description: 如何在 Laravel 中限制订阅套餐的使用额度，从控制器里的简单检查，到 Gate 和 Policy。
---

# 如何设计订阅套餐的限制

我通常会把用户的订阅套餐存放在 `users` 表中，比如用 `plan_id` 关联到订阅套餐。

先搭一个场景……

我们的 subscription plans.php

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

然后就可以在控制器里这样使用（这里假设只有已登录的用户才能访问这个接口）

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

这是一种可行的做法，我认为对大多数应用来说，用它来校验创建数量的上限已经足够了。

不过，正如您所说，还有一种“更集中”的做法。

使用 Gate 和 Policy。

我们把之前的代码重构为 Gate

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

在控制器中，可以重构成这样

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

这应该已经能满足您的需求了。但是，也许还有更规范的做法？

您说得对，那就是使用 Policy。

有一种说法是，Policy 最适合处理某个领域中的“资源”或模型，比如我们这个例子里的 List 模型。

而 Gate 最适合处理某个“操作”。

不过，我个人的看法是，只要可能就应该用 Policy。一来它给了您一个把代码拆分到不同文件的理由，二来它也不会让您的 `AuthServiceProvider` 变得杂乱（顺便说一句，我不太喜欢修改 Service Provider，除非确实有必要）。

我们把之前的代码重构为 Policy

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

注意到了吗，这个方法的命名和资源控制器是一致的。我相信这是 Laravel 团队有意为之的设计，值得点赞。

然后在控制器中直接这样使用即可

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

这里的 `List::class` 只是用来帮助 Laravel 推断应该使用哪个 Policy 类。

总的来说，这样就能避免一种问题：同一个列表可能在 10 个不同的控制器里被创建。我们只需要在这些地方加上 `$this->authorize()` 就可以了。

**继续阅读：** [如何打造您自己的内容管理系统（类似 WordPress）](https://sawirstudio.com/zh-hans/posts/how-to-build-your-own-content-management-system-like-wordpress/)，或者在我们的 [WordPress 插件](https://sawirstudio.com/zh-hans/projects/wordpress-plugins/)中了解 Sawir Renew，这是一款适用于 WooCommerce 的免费订阅插件。
