---
title: واجهة WordPress Abilities API للإضافات | SawirStudio
url: https://sawirstudio.com/ar/posts/wordpress-abilities-api-for-plugins/
description: كيف تسجّلون القدرات في إضافة WordPress ليتمكن وكلاء الذكاء الاصطناعي والأدوات الأخرى من اكتشافها وتشغيلها بأمان: الفئات، والمخططات، والصلاحيات، والتوصيفات، ونقاط نهاية REST.
---

# اجعلوا إضافة WordPress قابلة للاستخدام من وكلاء الذكاء الاصطناعي عبر Abilities API

واجهة Abilities API، التي أُضيفت في WordPress 6.9، هي مكان واحد تصرّح فيه الإضافة بما تستطيع فعله: اسم ووصف، والمدخلات التي تقبلها، وما تُرجعه، ومن يحق له تنفيذها. وبعدها تستطيع وكلاء الذكاء الاصطناعي، وإضافة WordPress AI، ومحوّلات MCP، وحتى السكربتات العادية، اكتشاف هذه القدرات واستخدامها، دون أن تضطر كل أداة إلى تعلّم كل إضافة يدويًا.

إليكم طريقة عملها، مع مثال من إحدى [إضافات WordPress](https://sawirstudio.com/ar/projects/wordpress-plugins/) التي طوّرناها.

## المثال

[Sawir Restock Notifier](https://sawirstudio.com/ar/projects/wordpress-plugins/sawir-restock-notifier-for-woocommerce/) تتيح للمتسوّقين طلب إشعار عبر البريد الإلكتروني عند عودة المنتج إلى المخزون. ومن الأسئلة الطبيعية التي قد يجيب عنها الوكيل: «ما المنتجات النافدة التي ينتظرها أكبر عدد من الأشخاص؟». وهذه قدرة للقراءة فقط.

## 1. تسجيل فئة

تنتمي كل قدرة إلى فئة واحدة بالضبط، وتُسجَّل الفئات أولًا، على خطّاف (hook) خاص بها:

```
add_action( 'wp_abilities_api_categories_init', function () {
	wp_register_ability_category( 'store-insights', array(
		'label'       => __( 'Store insights', 'sawir-restock-notifier-for-woocommerce' ),
		'description' => __( 'Read-only questions about the store.', 'sawir-restock-notifier-for-woocommerce' ),
	) );
} );
```

## 2. تسجيل القدرة

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

دالة الاستدعاء (callback) هي كود إضافة عادي يُرجع بيانات تطابق مخطط المخرجات. لا شيء فيها خاص بالذكاء الاصطناعي.

## وظيفة كل جزء

- **الاسم**، بصيغة `namespace/ability-name`. استخدموا المعرّف المختصر (slug) لإضافتكم كمساحة اسم.
- **الوصف.** هذا ما يقرؤه الوكيل ليقرّر هل يستخدم القدرة أم لا، فاكتبوه لشخص لم يرَ إضافتكم من قبل.
- **مخططا المدخلات والمخرجات**، بصيغة JSON Schema. يجري التحقق من المدخلات وفقهما قبل تشغيل الكود، وتعرف الأدوات ما ستحصل عليه في المقابل.
- **دالة الصلاحيات.** هي الجزء الوحيد الذي يحمي فعليًا. تحققوا من الصلاحيات هنا كما تفعلون في شاشة الإدارة.
- **التوصيفات (Annotations).** تُعلم `readonly` و`destructive` و`idempotent` الأدوات بمدى الحذر المطلوب. وهي مجرد تلميحات وليست أمانًا؛ فالأمان في دالة الصلاحيات.
- **`show_in_rest`.** القدرات مخفية عن REST API ما لم تفعّلوا هذا الخيار.

## التشغيل عبر REST

تُدرج القدرات المعروضة في REST على المسار `/wp-json/wp-abilities/v1/abilities`، ولكل قدرة نقطة نهاية `/run`. تُشغَّل القدرة القرائية بطلب GET، ومدخلاتها بصيغة JSON مشفّرة في الرابط ضمن المعامل `input`. أما التي تُغيّر البيانات فتُشغَّل بطلب POST، والتي تحذف بطلب DELETE. يحتاج كل طلب إلى مستخدم موثَّق، وبعدها تحسم دالة الصلاحيات لديكم الأمر.

## وللوكلاء عبر MCP

يعرض MCP Adapter، الذي يُطوَّر بالتوازي مع إضافة WordPress AI، القدرات المسجّلة لعملاء MCP مثل Claude. لا تحتاجون إلى كتابة أي شيء إضافي له: القدرة الجيدة الوصف ذات المخططات الواضحة هي أصلًا أداة جيدة.

## ما الذي سنعرضه أولًا

ابدؤوا بالقدرات القرائية. فهي مفيدة فورًا («كم متسوّقًا ينتظر الوشاح الأزرق؟») ولا يمكن أن تُحدث ضررًا. وأجّلوا كل ما يرسل بريدًا إلكترونيًا أو يغيّر الأسعار أو يحذف بيانات إلى أن تتأكدوا من صحة فحوص الصلاحيات والتوصيفات.

تحتاج Abilities API إلى WordPress 6.9 أو أحدث. وتجدون المرجع الكامل في [الوثائق الرسمية](https://developer.wordpress.org/apis/abilities/). وإن أردتم إضافة هذه الميزة إلى إضافة تعتمدون عليها، [تواصلوا معنا](https://sawirstudio.com/ar/contact/).

**تابعوا القراءة:** [WordPress 7.2 يصدر في 8 ديسمبر: ما الذي نختبره في إضافاتنا](https://sawirstudio.com/ar/posts/wordpress-7-2-what-plugin-authors-should-test/)، أو [ما الذي تطلّبه جعل مواقعنا جاهزة للوكلاء](https://sawirstudio.com/ar/posts/making-our-sites-agent-ready/).
