---
title: Limits für Abo-Pläne in Laravel | SawirStudio
url: https://sawirstudio.com/de/posts/how-do-you-design-your-subscription-plan-restrictions/
description: So setzen Sie Limits für Abo-Pläne in Laravel durch, von der einfachen Prüfung im Controller bis zu Gates und Policies.
---

# Wie gestalten Sie die Beschränkungen Ihrer Abo-Pläne?

Den Abo-Plan eines Nutzers speichere ich meistens in der Tabelle `users`, zum Beispiel als `plan_id`, die auf die Abo-Pläne verweist.

Bauen wir dafür eine Grundlage …

unsere subscription plans.php

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

Dann können wir das im Controller so verwenden (wir gehen davon aus, dass nur angemeldete Nutzer diesen Endpunkt aufrufen können):

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

Das ist ein gültiger Weg, und für die meisten Anwendungen erfüllt er seinen Zweck, wenn es darum geht, das Limit beim Anlegen zu prüfen.

Wie Sie aber erwähnt haben, gibt es dafür noch einen „zentraleren“ Ansatz.

Gate und Policy verwenden.

Refaktorieren wir unseren bisherigen Code zu einem Gate

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

Und in Ihrem Controller können wir das so umbauen:

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

Damit sollten Ihre Anforderungen eigentlich schon erfüllt sein. Aber gibt es vielleicht einen klarer vorgegebenen Weg?

Sie haben recht: die Policy.

Es heißt, eine Policy eignet sich am besten, wenn man es mit einer „Ressource“ bzw. mit Modellen einer Domäne zu tun hat, in unserem Fall dem Modell List.

Ein Gate eignet sich dagegen besser für eine „Aktion“.

Meine persönliche Meinung ist aber, dass wir wann immer möglich eine Policy verwenden sollten. Sie gibt einem nämlich nicht nur einen Vorwand, eigene Dateien anzulegen, sondern sie vermüllt auch nicht Ihren `AuthServiceProvider` (ich ändere den Service Provider übrigens ungern, es sei denn, es ist wirklich nötig).

Refaktorieren wir unseren bisherigen Code zu einer Policy

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

Fällt Ihnen auf, dass die Methode dem Resource-Controller nachempfunden ist? Das ist meines Wissens so vom Laravel-Team gewollt. Respekt dafür.

und im Controller können wir das einfach so verwenden:

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

`List::class` ist nur dafür da, Laravel dabei zu helfen, die passende Policy-Klasse zu ermitteln.

Alles in allem vermeiden wir so das Problem, dass sich die Liste buchstäblich aus 10 verschiedenen Controllern anlegen lässt. Wir setzen einfach `$this->authorize()` dorthin, und die Sache ist erledigt.

**Weiterlesen:** [Wie Sie Ihr eigenes Content-Management-System bauen (wie WordPress)](https://sawirstudio.com/de/posts/how-to-build-your-own-content-management-system-like-wordpress/), oder sehen Sie sich Sawir Renew an, das kostenlose Abo-Plugin für WooCommerce, in unseren [WordPress-Plugins](https://sawirstudio.com/de/projects/wordpress-plugins/).
