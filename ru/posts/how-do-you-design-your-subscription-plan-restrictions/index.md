---
title: Ограничения тарифных планов в Laravel | SawirStudio
url: https://sawirstudio.com/ru/posts/how-do-you-design-your-subscription-plan-restrictions/
description: Как внедрить ограничения тарифных планов в Laravel: от простой проверки в контроллере до Gate и Policy.
---

# Как спроектировать ограничения тарифных планов подписки

Обычно я храню тарифный план пользователя в таблице `users`. Допустим, в поле `plan_id`, которое ссылается на тарифные планы подписки.

Подготовим почву…

наши тарифные планы.php

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

Затем мы можем использовать это в контроллере вот так (предполагается, что этим эндпоинтом могут пользоваться только авторизованные пользователи)

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

Это рабочий способ, и для большинства приложений, как мне кажется, он вполне справляется с задачей проверки лимита на создание.

Однако, как вы заметили, есть и другой, «более централизованный» подход.

Использование Gate и Policy.

Давайте перепишем наш предыдущий код на Gate

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

А в контроллере мы можем переписать так

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

Этого уже должно хватить для ваших задач. Но, возможно, есть более строгий способ?

Вы правы. Это Policy.

Говорят, что Policy лучше всего подходит для работы с «ресурсом», то есть с моделями предметной области. Как в нашем случае, с моделью List.

А Gate лучше подходит для «действий».

Но, на мой взгляд, по возможности стоит использовать Policy: помимо того, что это даёт повод создать отдельные файлы, она ещё и не засоряет ваш `AuthServiceProvider`… (кстати, я не очень люблю менять сервис-провайдер без крайней необходимости).

Давайте перепишем наш предыдущий код на Policy

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

Заметили, что методы повторяют методы ресурсного контроллера? Думаю, это сделано намеренно командой Laravel. Хвала им за это.

и в контроллере мы можем просто использовать это вот так

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

`List::class` нужен лишь для того, чтобы Laravel мог определить, какой класс Policy использовать.

В итоге так можно избежать проблем, когда список буквально можно создать из 10 разных контроллеров. Достаточно поставить туда `$this->authorize()`, и дело сделано.

**Читайте также:** [Как создать собственную систему управления контентом (как WordPress)](https://sawirstudio.com/ru/posts/how-to-build-your-own-content-management-system-like-wordpress/) или познакомьтесь с Sawir Renew, бесплатным плагином подписок для WooCommerce, в разделе [плагины для WordPress](https://sawirstudio.com/ru/projects/wordpress-plugins/).
