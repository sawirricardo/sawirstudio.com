---
title: Limieten van abonnementen in Laravel | SawirStudio
url: https://sawirstudio.com/nl/posts/how-do-you-design-your-subscription-plan-restrictions/
description: Zo dwing je limieten van abonnementen af in Laravel, van een simpele check in de controller tot Gates en Policies.
---

# Hoe ontwerp je de beperkingen van je abonnementen

Ik sla het abonnement van de gebruiker meestal op in de tabel `users`, bijvoorbeeld als `plan_id`, die verwijst naar de abonnementen.

Laten we hier een voorbeeld van maken…

onze subscription plans.php

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

Daarna kunnen we het zo in de controller gebruiken (ervan uitgaand dat alleen ingelogde gebruikers deze endpoint kunnen gebruiken)

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

Dit is een prima aanpak en voor de meeste applicaties volstaat dit om de limiet op het aanmaken te controleren.

Maar zoals je al zei, is er ook een meer gecentraliseerde aanpak.

Gate en Policy gebruiken.

Laten we onze vorige code ombouwen naar een Gate

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

En in je controller kunnen we dit zo aanpassen

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

Dit zou al genoeg moeten zijn voor wat je nodig hebt. Maar is er misschien een meer eigenzinnige manier?

Je hebt gelijk: de Policy.

Ze zeggen dat een Policy het beste past bij een "resource" of model uit je domein, zoals in ons geval het model List.

En een Gate past het beste bij een "actie".

Maar mijn persoonlijke mening is dat we waar mogelijk een Policy moeten gebruiken, want… behalve dat het je een excuus geeft om aparte bestanden aan te maken, vervuilt het ook je `AuthServiceProvider` niet… (ik pas de service provider trouwens liever niet aan, tenzij het echt nodig is).

Laten we onze vorige code ombouwen naar een Policy

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

Valt je op dat de methodes die van een resource controller nabootsen? Dat is volgens mij bewust zo bedacht door het Laravel-team. Petje af voor hen.

en in de controller kunnen we het gewoon zo gebruiken

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

de `List::class` staat er alleen om Laravel te helpen afleiden welke Policy-klasse het moet gebruiken.

Al met al voorkomen we zo de problemen waarbij je de lijst letterlijk vanuit 10 verschillende controllers kunt aanmaken. We zetten gewoon `$this->authorize()` erbij en klaar.

**Lees ook:** [Zo bouw je zelf een contentmanagementsysteem (zoals WordPress)](https://sawirstudio.com/nl/posts/how-to-build-your-own-content-management-system-like-wordpress/), of bekijk Sawir Renew, de gratis abonnementenplugin voor WooCommerce, bij onze [WordPress plugins](https://sawirstudio.com/nl/projects/wordpress-plugins/).
