---
title: WordPress Abilities API לתוספים | SawirStudio
url: https://sawirstudio.com/he/posts/wordpress-abilities-api-for-plugins/
description: איך לרשום abilities בתוסף WordPress כדי שסוכני AI וכלים אחרים יוכלו לגלות ולהפעיל אותן בבטחה: קטגוריות, סכמות, הרשאות, annotations ונקודות הקצה של ה-REST.
---

# איך להפוך את התוסף שלכם ל-WordPress לשמיש עבור סוכני AI בעזרת ה-Abilities API

ה-Abilities API, שנוסף ב-WordPress 6.9, הוא מקום אחד שבו תוסף מצהיר מה הוא יודע לעשות: שם, תיאור, איזה קלט הוא מקבל, מה הוא מחזיר ומי מורשה להפעיל אותו. סוכני AI, תוסף ה-AI של WordPress, מתאמי MCP וסקריפטים רגילים יכולים אז למצוא את היכולות האלה ולהשתמש בהן, בלי שכל כלי יצטרך ללמוד כל תוסף ידנית.

הנה איך זה עובד, עם אחד מ[תוספי ה-WordPress](https://sawirstudio.com/he/projects/wordpress-plugins/) שלנו כדוגמה.

## הדוגמה

[Sawir Restock Notifier](https://sawirstudio.com/he/projects/wordpress-plugins/sawir-restock-notifier-for-woocommerce/) מאפשר ללקוחות לבקש הודעה במייל כשמוצר חוזר למלאי. שאלה מתבקשת שסוכן יכול לענות עליה היא "אילו מוצרים שאזלו מהמלאי יש הכי הרבה אנשים שמחכים להם?". זו יכולת לקריאה בלבד.

## 1. רישום קטגוריה

כל יכולת שייכת בדיוק לקטגוריה אחת, והקטגוריות נרשמות קודם, ב-hook נפרד משלהן:

```
add_action( 'wp_abilities_api_categories_init', function () {
	wp_register_ability_category( 'store-insights', array(
		'label'       => __( 'Store insights', 'sawir-restock-notifier-for-woocommerce' ),
		'description' => __( 'Read-only questions about the store.', 'sawir-restock-notifier-for-woocommerce' ),
	) );
} );
```

## 2. רישום היכולת

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

ה-callback הוא קוד רגיל של התוסף שמחזיר נתונים בהתאם ל-output schema. אין בו שום דבר ייחודי ל-AI.

## למה משמש כל חלק

- **שם**, בפורמט `namespace/ability-name`. השתמשו ב-slug של התוסף שלכם בתור ה-namespace.
- **תיאור.** זה מה שסוכן קורא כדי להחליט אם להשתמש ביכולת, אז כתבו אותו למי שמעולם לא ראה את התוסף שלכם.
- **סכמות קלט ופלט**, ב-JSON Schema. הקלט נבדק מולן לפני שהקוד שלכם רץ, והכלים יודעים מה הם יקבלו בחזרה.
- **Permission callback.** החלק היחיד שבאמת מגן על משהו. בדקו כאן הרשאות בדיוק כמו במסך ניהול.
- **Annotations.** הערכים `readonly`, `destructive` ו-`idempotent` אומרים לכלים עד כמה להיזהר. אלה רמזים ולא אבטחה; ה-permission callback הוא האבטחה.
- **`show_in_rest`.** יכולות מוסתרות מ-REST API אלא אם מפעילים את האפשרות הזו.

## הפעלה דרך REST

יכולות שמוצגות ב-REST מופיעות ברשימה בכתובת `/wp-json/wp-abilities/v1/abilities`, ולכל אחת יש endpoint בשם `/run`. יכולת לקריאה בלבד רצה עם GET, והקלט שלה מועבר כ-JSON מקודד ל-URL בפרמטר `input`. יכולת ששינה נתונים רצה עם POST, ויכולת הרסנית עם DELETE. כל בקשה דורשת משתמש מאומת, ואז ה-permission callback שלכם מחליט.

## ולסוכנים, דרך MCP

ה-MCP Adapter, שמפותח במקביל לתוסף ה-AI של WordPress, חושף את היכולות הרשומות ללקוחות MCP כמו Claude. אתם לא צריכים לכתוב בשבילו שום דבר נוסף: יכולת מתוארת היטב עם סכמות טובות היא כבר כלי טוב.

## מה היינו חושפים קודם

התחילו ביכולות לקריאה בלבד. הן שימושיות מיד ("כמה לקוחות מחכים לצעיף הכחול?") ואי אפשר לגרום בעזרתן נזק. השאירו לסוף כל דבר ששולח מייל, משנה מחירים או מוחק, עד שאתם בטוחים שבדיקות ההרשאות וה-annotations נכונים.

ה-Abilities API דורש WordPress 6.9 ומעלה. ב[תיעוד הרשמי](https://developer.wordpress.org/apis/abilities/) תמצאו את כל הפרטים. אם אתם רוצים שנוסיף את זה לתוסף שאתם מסתמכים עליו, [צרו איתנו קשר](https://sawirstudio.com/he/contact/).

**להמשך קריאה:** [WordPress 7.2 יוצא ב-8 בדצמבר: מה אנחנו בודקים בתוספים שלנו](https://sawirstudio.com/he/posts/wordpress-7-2-what-plugin-authors-should-test/), או [מה נדרש כדי להפוך את האתרים שלנו לידידותיים לסוכנים](https://sawirstudio.com/he/posts/making-our-sites-agent-ready/).
