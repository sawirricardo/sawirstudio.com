---
title: Limiti dei piani di abbonamento in Laravel | SawirStudio
url: https://sawirstudio.com/it/posts/how-do-you-design-your-subscription-plan-restrictions/
description: Come applicare i limiti dei piani di abbonamento in Laravel, da un semplice controllo nel controller fino a Gate e Policy.
---

# Come progettare le restrizioni dei piani di abbonamento

Di solito salvo il piano di abbonamento dell’utente nella tabella `users`, ad esempio con il campo `plan_id`, che fa riferimento ai piani di abbonamento.

Prepariamo il terreno…

i nostri piani di abbonamento.php

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

Poi possiamo usarlo nel controller in questo modo (si presume già che solo gli utenti autenticati possano usare questo endpoint)

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

È un approccio valido e, per la maggior parte delle applicazioni, credo che abbia fatto il suo dovere per verificare il limite di creazione.

Tuttavia, come dicevi, c’è anche un altro approccio, «più centralizzato».

Usare Gate e Policy.

Rifattorizziamo il codice di prima con un Gate

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

E nel controller possiamo rifattorizzare così

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

Questo dovrebbe già bastare. Ma forse esiste un modo più strutturato?

Hai ragione. Usare la Policy.

Si dice che la Policy sia l’ideale quando si ha a che fare con una «risorsa» o con i modelli di un dominio. Come nel nostro caso, il modello List.

Il Gate, invece, è l’ideale quando si tratta di un’«azione».

Però, secondo me, conviene usare la Policy ogni volta che è possibile, perché… oltre a darvi una scusa per creare file separati, non riempie nemmeno l’`AuthServiceProvider` (a me, tra l’altro, non piace molto modificare il service provider, a meno che non sia davvero necessario).

Rifattorizziamo il codice di prima con una Policy

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

Notate che il metodo imita il resource controller? Credo sia voluto dal team di Laravel. Complimenti a loro.

e nel controller possiamo semplicemente usarla così

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

`List::class` serve solo ad aiutare Laravel a capire quale classe Policy usare.

Insomma, in questo modo evitiamo i problemi per cui… si può creare una lista da 10 controller diversi. Basta inserire `$this->authorize()` e il gioco è fatto.

**Continuate a leggere:** [Come costruire il vostro content management system (come WordPress)](https://sawirstudio.com/it/posts/how-to-build-your-own-content-management-system-like-wordpress/), oppure scoprite Sawir Renew, il plugin gratuito per gli abbonamenti su WooCommerce, tra i nostri [plugin WordPress](https://sawirstudio.com/it/projects/wordpress-plugins/).
