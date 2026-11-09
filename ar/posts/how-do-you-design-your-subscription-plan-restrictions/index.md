---
title: حدود خطط الاشتراك في Laravel | SawirStudio
url: https://sawirstudio.com/ar/posts/how-do-you-design-your-subscription-plan-restrictions/
description: كيف تفرضون حدود خطط الاشتراك في Laravel، بدءًا من فحص بسيط في الـ controller وصولًا إلى Gates وPolicies.
---

# كيف تصمّمون قيود خطط الاشتراك

عادةً أخزّن خطة اشتراك المستخدم في جدول `users`، مثلًا عبر `plan_id` الذي يشير إلى خطط الاشتراك.

لنُجهّز المثال…

ملف خطط الاشتراك subscription plans.php

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

بعدها يمكننا استخدامها في الـ controller بهذا الشكل (مع افتراض أن نقطة الوصول هذه متاحة للمستخدمين المسجّلين فقط)

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

هذه طريقة صحيحة، وفي معظم التطبيقات أراها كافية للتحقق من حد الإنشاء.

لكن، كما ذكرتم، هناك أسلوب آخر أكثر مركزية.

استخدام Gate وPolicy.

لنُعد كتابة الكود السابق باستخدام Gate

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

وفي الـ controller يمكننا تبسيط الكود هكذا

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

هذا يفي بالغرض غالبًا. لكن، هل هناك أسلوب أكثر اتباعًا للأعراف المعتمدة؟

أنتم محقون. إنه استخدام Policy.

يُقال إن Policy هي الأنسب عند التعامل مع «مورد» أو نماذج (models) داخل نطاق العمل، كما في حالتنا مع النموذج List.

أما Gate فهي الأنسب عند التعامل مع «إجراء» محدد.

لكن رأيي الشخصي أن نستخدم Policy كلما أمكن، لأنها، إلى جانب أنها تمنحكم مبررًا لإنشاء ملفات منفصلة، لا تُثقل ملف `AuthServiceProvider` (وأنا في الحقيقة لا أحب تعديل الـ service provider إلا عند الضرورة).

لنُعد كتابة الكود السابق باستخدام Policy

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

هل لاحظتم أن أسماء الدوال تحاكي أسماء دوال الـ resource controller؟ أظن أن هذا مقصود من فريق Laravel، وهو يستحق الثناء.

ويمكننا استخدامها ببساطة في الـ controller هكذا

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

الصنف `List::class` موجود فقط ليساعد Laravel على تحديد صنف الـ Policy الذي يجب استخدامه.

في المحصلة، بهذه الطريقة نتجنب المشكلات التي قد تجعلكم تُنشئون القائمة من 10 controllers مختلفة. يكفي أن نضع `$this->authorize()` هناك وينتهي الأمر.

**تابعوا القراءة:** [كيف تبنون نظام إدارة المحتوى الخاص بكم (مثل WordPress)](https://sawirstudio.com/ar/posts/how-to-build-your-own-content-management-system-like-wordpress/)، أو اطّلعوا على Sawir Renew، إضافة الاشتراكات المجانية لـ WooCommerce ضمن [إضافات WordPress](https://sawirstudio.com/ar/projects/wordpress-plugins/) لدينا.
