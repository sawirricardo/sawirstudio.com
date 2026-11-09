---
title: प्लगइन के लिए WordPress Abilities API | SawirStudio
url: https://sawirstudio.com/hi/posts/wordpress-abilities-api-for-plugins/
description: WordPress प्लगइन में एबिलिटीज़ कैसे रजिस्टर करें, ताकि AI एजेंट्स और दूसरे टूल्स उन्हें खोज सकें और सुरक्षित तरीके से चला सकें: कैटेगरी, स्कीमा, परमिशन, एनोटेशन और REST एंडपॉइंट।
---

# Abilities API से अपने WordPress प्लगइन को AI एजेंट्स के लिए उपयोगी बनाएँ

WordPress 6.9 में जोड़ी गई Abilities API एक ऐसी जगह है जहाँ प्लगइन बताता है कि वह क्या-क्या कर सकता है: उसका नाम, विवरण, वह कौन-सा इनपुट लेता है, क्या लौटाता है और उसे चलाने की अनुमति किसे है। इसके बाद AI एजेंट्स, WordPress AI प्लगइन, MCP एडेप्टर और सामान्य स्क्रिप्ट इन एबिलिटीज़ को खोजकर इस्तेमाल कर सकते हैं, और हर टूल को हर प्लगइन अलग से हाथ से सीखने की ज़रूरत नहीं पड़ती।

यह कैसे काम करता है, इसे हम अपने एक [WordPress प्लगइन](https://sawirstudio.com/hi/projects/wordpress-plugins/) के उदाहरण से समझाते हैं।

## उदाहरण

[Sawir Restock Notifier](https://sawirstudio.com/hi/projects/wordpress-plugins/sawir-restock-notifier-for-woocommerce/) की मदद से ग्राहक यह अनुरोध कर सकते हैं कि कोई प्रोडक्ट दोबारा स्टॉक में आने पर उन्हें ईमेल मिले। एजेंट के लिए एक स्वाभाविक सवाल होगा: “किन आउट-ऑफ़-स्टॉक प्रोडक्ट्स का इंतज़ार सबसे ज़्यादा लोग कर रहे हैं?” यह एक रीड-ओनली एबिलिटी है।

## 1. एक कैटेगरी रजिस्टर करें

हर एबिलिटी ठीक एक कैटेगरी में आती है, और कैटेगरी पहले, अपने अलग हुक पर रजिस्टर की जाती हैं:

```
add_action( 'wp_abilities_api_categories_init', function () {
	wp_register_ability_category( 'store-insights', array(
		'label'       => __( 'Store insights', 'sawir-restock-notifier-for-woocommerce' ),
		'description' => __( 'Read-only questions about the store.', 'sawir-restock-notifier-for-woocommerce' ),
	) );
} );
```

## 2. एबिलिटी रजिस्टर करें

```
add_action( 'wp_abilities_api_init', function () {
	wp_register_ability( 'sawir-restock/waiting-list', array(
		'label'       => __( 'Restock waiting list', 'sawir-restock-notifier-for-woocommerce' ),
		'description' => __( 'Out-of-stock products, with how many shoppers are waiting for each.', 'sawir-restock-notifier-for-woocommerce' ),
		'category'    => 'store-insights',
		'input_schema' => array(
			'type'       => 'object',
			'properties' => array(
				'limit' => array( 'type' => 'integer', 'minimum' => 1, 'maximum' => 50, 'default' => 10 ),
			),
		),
		'output_schema' => array(
			'type'  => 'array',
			'items' => array(
				'type'       => 'object',
				'properties' => array(
					'product_id' => array( 'type' => 'integer' ),
					'name'       => array( 'type' => 'string' ),
					'waiting'    => array( 'type' => 'integer' ),
				),
			),
		),
		'execute_callback'    => 'sawir_restock_waiting_list',
		'permission_callback' => function () {
			return current_user_can( 'manage_woocommerce' );
		},
		'meta' => array(
			'show_in_rest' => true,
			'annotations'  => array( 'readonly' => true ),
		),
	) );
} );
```

कॉलबैक सामान्य प्लगइन कोड ही है, जो आउटपुट स्कीमा से मेल खाता डेटा लौटाता है। इसमें AI से जुड़ा कुछ भी अलग नहीं है।

## हर हिस्से का काम

- **नाम**, `namespace/ability-name` के रूप में। नेमस्पेस के लिए अपने प्लगइन का स्लग इस्तेमाल करें।
- **विवरण।** एजेंट इसे पढ़कर तय करता है कि इस एबिलिटी को इस्तेमाल करना है या नहीं, इसलिए इसे ऐसे व्यक्ति के लिए लिखें जिसने आपका प्लगइन कभी देखा ही नहीं।
- **इनपुट और आउटपुट स्कीमा**, JSON Schema में। आपका कोड चलने से पहले इनपुट को इन्हीं के आधार पर जाँचा जाता है, और टूल्स को पहले से पता होता है कि बदले में उन्हें क्या मिलेगा।
- **परमिशन कॉलबैक।** असल में सुरक्षा यही एक हिस्सा करता है। यहाँ क्षमताओं (capabilities) की जाँच वैसे ही करें जैसे किसी एडमिन स्क्रीन के लिए करते हैं।
- **एनोटेशन।** `readonly`, `destructive` और `idempotent` टूल्स को बताते हैं कि उन्हें कितनी सावधानी बरतनी है। ये सिर्फ़ संकेत हैं, सुरक्षा नहीं; सुरक्षा परमिशन कॉलबैक से ही आती है।
- **`show_in_rest`।** जब तक आप इसे चालू नहीं करते, एबिलिटीज़ REST API में दिखाई नहीं देतीं।

## REST के ज़रिए चलाना

REST में दिखाई गई एबिलिटीज़ की सूची `/wp-json/wp-abilities/v1/abilities` पर मिलती है, और हर एबिलिटी का अपना एक `/run` एंडपॉइंट होता है। रीड-ओनली एबिलिटी GET से चलती है, जिसका इनपुट `input` पैरामीटर में URL-एनकोडेड JSON के रूप में जाता है। जो एबिलिटी कुछ बदलती है वह POST से चलती है, और डिस्ट्रक्टिव एबिलिटी DELETE से। हर रिक्वेस्ट के लिए एक ऑथेंटिकेटेड यूज़र ज़रूरी है, और उसके बाद फ़ैसला आपका परमिशन कॉलबैक करता है।

## और MCP के ज़रिए एजेंट्स तक

WordPress AI प्लगइन के साथ विकसित किया जा रहा MCP Adapter, रजिस्टर की गई एबिलिटीज़ को Claude जैसे MCP क्लाइंट्स के सामने उपलब्ध कराता है। इसके लिए आपको अलग से कुछ नहीं लिखना पड़ता: अच्छे विवरण और सही स्कीमा वाली एबिलिटी अपने आप में एक अच्छा टूल है।

## हम सबसे पहले क्या उपलब्ध कराएँगे

रीड-ओनली एबिलिटीज़ से शुरुआत करें। ये तुरंत काम आती हैं (“नीले स्कार्फ़ का इंतज़ार कितने ग्राहक कर रहे हैं?”) और इनसे कोई नुकसान नहीं हो सकता। जो भी एबिलिटी ईमेल भेजती है, कीमतें बदलती है या कुछ डिलीट करती है, उसे तब तक छोड़ दें जब तक आपको भरोसा न हो जाए कि आपकी परमिशन जाँच और एनोटेशन सही हैं।

Abilities API के लिए WordPress 6.9 या उसके बाद का वर्शन ज़रूरी है। पूरा रेफ़रेंस [आधिकारिक डॉक्यूमेंटेशन](https://developer.wordpress.org/apis/abilities/) में उपलब्ध है। अगर आप चाहते हैं कि यह आपके किसी भरोसेमंद प्लगइन में जोड़ा जाए, तो [हमसे संपर्क करें](https://sawirstudio.com/hi/contact/)।

**यह भी पढ़ें:** [WordPress 7.2 8 दिसंबर को आ रहा है: हम अपने प्लगइन्स में क्या टेस्ट कर रहे हैं](https://sawirstudio.com/hi/posts/wordpress-7-2-what-plugin-authors-should-test/), या [हमारी साइट्स को एजेंट-रेडी बनाने में क्या-क्या लगा](https://sawirstudio.com/hi/posts/making-our-sites-agent-ready/)।
