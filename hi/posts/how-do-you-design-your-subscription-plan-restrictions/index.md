---
title: Laravel में सब्सक्रिप्शन प्लान की सीमाएँ | SawirStudio
url: https://sawirstudio.com/hi/posts/how-do-you-design-your-subscription-plan-restrictions/
description: Laravel में सब्सक्रिप्शन प्लान की सीमाएँ कैसे लागू करें, कंट्रोलर में एक साधारण जाँच से लेकर Gate और Policy तक।
---

# आप अपने सब्सक्रिप्शन प्लान की पाबंदियाँ कैसे डिज़ाइन करते हैं

मैं आमतौर पर यूज़र का सब्सक्रिप्शन प्लान `users` टेबल में रखता हूँ, जैसे `plan_id`, जो सब्सक्रिप्शन प्लान की ओर इशारा करता है।

चलिए, इसके लिए एक उदाहरण तैयार करते हैं…

हमारे subscription plans.php

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

फिर हम इसे कंट्रोलर में इस तरह इस्तेमाल कर सकते हैं (यहाँ मान लिया गया है कि इस एंडपॉइंट का इस्तेमाल सिर्फ़ लॉग-इन यूज़र ही कर सकते हैं)

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

यह एक सही तरीका है, और ज़्यादातर ऐप्लिकेशन के लिए मेरे हिसाब से यह क्रिएशन लिमिट जाँचने का काम अच्छे से कर देता है।

लेकिन, जैसा कि आपने कहा था, इसके लिए एक और 'ज़्यादा केंद्रीकृत' तरीका भी है।

Gate और Policy का इस्तेमाल।

चलिए, अपने पिछले कोड को Gate में बदलते हैं

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

और आपके कंट्रोलर में हम इसे इस तरह बदल सकते हैं

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

इससे आपकी ज़रूरत पूरी हो जानी चाहिए। लेकिन क्या इसका कोई और ज़्यादा तय ढाँचे वाला तरीका है?

आप सही कह रहे हैं। Policy का इस्तेमाल।

कहा जाता है कि जब आप किसी “रिसोर्स” या डोमेन के मॉडल के साथ काम कर रहे हों, तो Policy सबसे अच्छी रहती है। जैसे हमारे मामले में List मॉडल।

और Gate तब सबसे अच्छा रहता है जब आप किसी “एक्शन” के साथ काम कर रहे हों।

लेकिन मेरी निजी राय है कि जहाँ भी संभव हो, हमें Policy ही इस्तेमाल करनी चाहिए, क्योंकि… अलग फ़ाइलें बनाने का बहाना मिलने के अलावा, यह आपके `AuthServiceProvider` को भी गंदा नहीं करती… (वैसे मुझे सर्विस प्रोवाइडर में बदलाव करना ज़्यादा पसंद नहीं है, जब तक वाकई ज़रूरी न हो)।

चलिए, अपने पिछले कोड को Policy में बदलते हैं

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

अरे, देखा आपने कि मेथड रिसोर्स कंट्रोलर की नकल करते हैं? मेरे ख़याल से Laravel टीम ने यह जान-बूझकर ऐसा रखा है। उन्हें शाबाशी।

और कंट्रोलर में हम इसे बस इस तरह इस्तेमाल कर सकते हैं

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

`List::class` सिर्फ़ इसलिए है कि Laravel को पता चल सके कि कौन-सी Policy क्लास इस्तेमाल करनी है।

कुल मिलाकर, इस तरह हम उस समस्या से बच सकते हैं जहाँ… आप 10 अलग-अलग कंट्रोलर से लिस्ट बना सकते हैं। हम बस `$this->authorize()` वहाँ लगा दें और काम ख़त्म।

**आगे पढ़ें:** [अपना खुद का कंटेंट मैनेजमेंट सिस्टम कैसे बनाएँ (WordPress जैसा)](https://sawirstudio.com/hi/posts/how-to-build-your-own-content-management-system-like-wordpress/), या हमारे [WordPress प्लगइन](https://sawirstudio.com/hi/projects/wordpress-plugins/) में WooCommerce के लिए मुफ़्त सब्सक्रिप्शन प्लगइन Sawir Renew देखें।
