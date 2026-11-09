---
title: Limites des forfaits d’abonnement dans Laravel | SawirStudio
url: https://sawirstudio.com/fr/posts/how-do-you-design-your-subscription-plan-restrictions/
description: Comment appliquer les limites des forfaits d’abonnement dans Laravel, d’une simple vérification dans le contrôleur jusqu’aux Gates et aux Policies.
---

# Comment concevoir les restrictions de vos forfaits d’abonnement

Je stocke généralement le forfait d’abonnement de l’utilisateur dans la table `users`, par exemple avec un `plan_id` qui renvoie aux forfaits d’abonnement.

Posons le décor…

nos forfaits d’abonnement.php

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

On peut ensuite l’utiliser dans le contrôleur comme ceci (on suppose ici que seuls les utilisateurs connectés peuvent accéder à ce point d’accès)

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

C’est une méthode valable, et pour la plupart des applications, je pense qu’elle remplit son rôle pour contrôler la limite de création.

Cependant, comme vous l’avez mentionné, il existe une autre approche, « plus centralisée » :

Utiliser Gate et Policy.

Refactorons notre code précédent avec un Gate

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

Et dans votre contrôleur, on peut le refactoriser ainsi

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

Cela devrait déjà répondre à vos besoins. Mais existe-t-il peut-être une approche plus encadrée ?

Vous avez raison : il faut utiliser la policy.

On dit que la Policy convient mieux quand on manipule une « ressource » ou des modèles d’un domaine métier, comme le modèle List dans notre cas.

Et que le Gate convient mieux pour une « action ».

Mais personnellement, je pense qu’il vaut mieux utiliser une policy dès que c’est possible, car… outre le fait que cela donne une bonne excuse pour créer des fichiers séparés, cela évite aussi d’encombrer votre `AuthServiceProvider`… (je n’aime pas trop modifier le service provider, d’ailleurs, sauf si c’est vraiment nécessaire).

Refactorons notre code précédent avec une Policy

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

Vous avez remarqué que les méthodes reprennent celles du contrôleur de ressources ? C’est voulu, je crois, de la part de l’équipe Laravel. Bravo à elle.

et on peut simplement l’utiliser ainsi dans le contrôleur

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

le `List::class` est là uniquement pour aider Laravel à déterminer quelle classe Policy utiliser.

Au final, cette approche nous évite le problème où… on peut littéralement créer la liste depuis 10 contrôleurs différents. Il suffit d’y placer le `$this->authorize()` et l’affaire est réglée.

**À lire aussi :** [Comment créer votre propre système de gestion de contenu (comme WordPress)](https://sawirstudio.com/fr/posts/how-to-build-your-own-content-management-system-like-wordpress/), ou découvrez Sawir Renew, le plugin d’abonnements gratuit pour WooCommerce, dans nos [plugins WordPress](https://sawirstudio.com/fr/projects/wordpress-plugins/).
