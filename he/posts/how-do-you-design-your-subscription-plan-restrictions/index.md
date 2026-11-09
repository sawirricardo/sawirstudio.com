---
title: מגבלות תוכניות מנוי ב-Laravel | SawirStudio
url: https://sawirstudio.com/he/posts/how-do-you-design-your-subscription-plan-restrictions/
description: איך להגביל את תוכניות המנוי ב-Laravel, מבדיקה פשוטה ב-controller ועד Gates ו-Policies.
---

# איך מעצבים הגבלות לתוכניות מנוי

בדרך כלל אני שומר את תוכנית המנוי של המשתמש בטבלת `users`, למשל בעמודה `plan_id` שמפנה לתוכניות המנוי.

בואו נכין לזה במה…

תוכניות המנוי שלנו.php

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

ואז אפשר להשתמש בזה ב-controller כך (בהנחה שרק משתמשים מחוברים יכולים לגשת ל-endpoint הזה)

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

זו דרך תקינה, ולדעתי ברוב היישומים היא עושה את העבודה כשצריך לאמת את מגבלת היצירה.

אבל, כפי שציינתם, יש גם גישה "ריכוזית" יותר.

שימוש ב-Gate וב-Policy.

בואו נעשה refactor לקוד הקודם ל-Gate

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

וב-controller אפשר לעשות refactor כך

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

זה כבר אמור לענות על הצרכים שלכם. אבל אולי יש דרך מחמירה יותר?

אתם צודקים. שימוש ב-Policy.

אומרים ש-Policy מתאים יותר כשמתעסקים ב"משאב" או במודלים בדומיין, כמו מודל List במקרה שלנו.

ו-Gate מתאים יותר כשמתעסקים ב"פעולה".

אבל לדעתי האישית כדאי להשתמש ב-Policy בכל פעם שאפשר, כי… חוץ מזה שזה נותן תירוץ ליצור קבצים נפרדים, זה גם לא מלכלך את `AuthServiceProvider`… (אגב, אני לא ממש אוהב לשנות את ה-service provider, אלא אם באמת צריך).

בואו נעשה refactor לקוד הקודם ל-Policy

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

שמתם לב שהמתודה מחקה את ה-resource controller? לדעתי זה מכוון מצד צוות Laravel. כל הכבוד להם.

ואפשר פשוט להשתמש בזה ב-controller כך

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

ה-`List::class` נמצא שם רק כדי לעזור ל-Laravel להסיק באיזו מחלקת Policy להשתמש.

בסך הכול, בדרך הזו אפשר להימנע מהבעיות שבהן… אפשר ממש ליצור את הרשימה מ-10 controllers שונים. פשוט שמים שם את `$this->authorize()` ומסיימים.

**להמשך קריאה:** [How To Build Your Own Content Management System (Like WordPress)](https://sawirstudio.com/he/posts/how-to-build-your-own-content-management-system-like-wordpress/), או הכירו את Sawir Renew, תוסף המנויים החינמי ל-WooCommerce, ב[תוספי WordPress שלנו](https://sawirstudio.com/he/projects/wordpress-plugins/).
