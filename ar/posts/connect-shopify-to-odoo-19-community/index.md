---
title: ربط Shopify بـ Odoo 19 Community | SawirStudio
url: https://sawirstudio.com/ar/posts/connect-shopify-to-odoo-19-community/
description: كيف تربطون متجر Shopify بـ Odoo 19 Community عبر موصل Shopify المفتوح المصدر من OCA: تطبيق Dev Dashboard، والصلاحيات (scopes)، وqueue_job، والـ webhooks، والإعداد الأولي.
---

# ربط Shopify بـ Odoo 19 Community خطوة بخطوة

كتبنا [موصل Shopify لـ Odoo](https://sawirstudio.com/ar/projects/odoo-shopify-connector/) وأسهمنا به في Odoo Community Association (OCA). الموصل مجاني ومفتوح المصدر (AGPL-3) ويعمل مع Odoo 19 Community. فيما يلي خطوات الإعداد من البداية إلى النهاية، كما ننفذها لعملائنا.

يزامن الموصل المنتجات والمخزون والعملاء والطلبات والمرتجعات والمبالغ المستردة والشحن، لمتجر واحد أو لعدة متاجر. أما المحاسبة ونقاط البيع فهما وحدتان منفصلتان يمكنكم إضافتهما لاحقًا.

## قبل أن تبدأوا

- نسخة Odoo 19 Community يمكن الوصول إليها عبر HTTPS علني. فـ Shopify يرسل إليها الـ webhooks ويحمّل منها صور المنتجات.
- صلاحيات مسؤول على متجر Shopify، وعلى مؤسسة Shopify (organization) التي تملكه.
- خطة Shopify تتيح الوصول إلى بيانات العملاء المحمية للتطبيقات التي تنشئونها بأنفسكم. الخطة Basic لا تتيح ذلك، ومن دونه يرفض Shopify بعض webhooks الطلبات فيفشل استيراد الطلبات والعملاء.

## 1. إنشاء التطبيق في Shopify Dev Dashboard

منذ 1 يناير 2026 لم يعد Shopify يسمح بإنشاء «التطبيقات المخصصة» القديمة (custom apps) من لوحة إدارة المتجر. أنشئوا التطبيق في **Dev Dashboard** الخاصة بالمؤسسة التي تملك المتجر. وهذا مهم: فـ Odoo يحصل على رمز الوصول عبر client credentials grant، وهو لا يعمل إلا إذا كان التطبيق والمتجر تابعين للمؤسسة نفسها.

عند تعبئة بيانات إصدار التطبيق:

- **App URL:** `https://shopify.dev/apps/default-app-home`. فليس للموصل أي صفحة داخل Shopify.
- **Embed app in Shopify admin:** معطّل.
- **Use legacy install flow:** معطّل، حتى يمنح تثبيت التطبيق الصلاحيات (scopes) التي تختارونها.
- **Webhooks:** بلا أي webhook. فـ Odoo ينشئ webhooks الخاصة به أثناء الإعداد الأولي.
- **Webhooks API version:** `2026-01`، وهو الإصدار الذي يستخدمه الموصل.
- **Scopes:** `read_customers`, `write_customers`, `read_draft_orders`, `read_fulfillments`, `read_inventory`, `write_inventory`, `read_locations`, `read_merchant_managed_fulfillment_orders`, `write_merchant_managed_fulfillment_orders`, `read_orders`, `read_products`, `write_products` و`read_returns`. أضيفوا `read_all_orders` إذا أردتم استيراد الطلبات التي مضى عليها أكثر من 60 يومًا.

انشروا الإصدار (Release)، ثم ثبّتوا التطبيق على المتجر، وانسخوا client ID وclient secret من صفحة Settings في التطبيق.

لديكم تطبيق مخصص قديم أُنشئ قبل 2026؟ ما زال يعمل. اتركوا client ID فارغًا واستخدموا بدلًا منه Admin API access token وAPI secret key الخاصين به.

## 2. تثبيت الوحدات في Odoo

أضيفوا `queue_job` من OCA ومستودع [connector-shopify](https://github.com/OCA/connector-shopify) إلى مسار addons، ثم ثبّتوا **Shopify Connector**. يعمل الموصل عبر مهام في الخلفية، لذا يجب تحميل `queue_job` على مستوى الخادم كله، مع تحديد سعة لكل قناة Shopify:

```
[options]
server_wide_modules = base,rpc,web,queue_job

[queue_job]
channels = root:4,root.shopify.import:2,root.shopify.export:1,root.shopify.webhook:1
```

ثم اضبطوا معامل النظام `web.base.url` على عنوان Odoo العلني عبر HTTPS، واضبطوا `web.base.url.freeze` على `True`. وإلا تغيّر العنوان عندما يسجّل مسؤول الدخول من مكان آخر، فتذهب webhooks الخاصة بـ Shopify إلى الوجهة الخطأ.

## 3. تشغيل الإعداد الأولي

امنحوا فريقكم مجموعة **Shopify Manager**، ثم افتحوا **Shopify > Configuration > Instances** وانقروا **Start Onboarding**:

1. أدخلوا النطاق `yourstore.myshopify.com` وclient ID وclient secret، ثم اختبروا الاتصال. يطلب Odoo الرمز ويجدده تلقائيًا كل 24 ساعة.
2. راجعوا الصلاحيات (scopes). إذا كانت إحداها ناقصة، فأضيفوها إلى إصدار جديد من التطبيق وانشروه ووافقوا عليه في المتجر.
3. اربطوا مواقع Shopify بمستودعات Odoo، وراجعوا ربط الضرائب الذي ينشئه Odoo، ثم اختاروا سياسات المزامنة.
4. ثبّتوا الـ webhooks وضعوا الاستيرادات الأولى في قائمة الانتظار.

قبل أن تتركوا الطلبات والشحنات تعمل دون إشراف، راجعوا ربط بوابات الدفع وشركات الشحن. فمعظم المفاجآت لاحقًا تأتي من ربط لم ينظر فيه أحد.

## 4. تابعوا الأداء

تعرض لوحة **Shopify > Dashboard** حالات المزامنة، والمهام المتراكمة، وآخر الأخطاء، وسلامة الـ webhooks. لا يفترض الموصل أن كل webhook سيصل: فالمهام المجدولة تعيد استيراد أي تغيير حدث في Shopify وفاته Odoo. وعندما يفشل شيء، عالجوا السبب ثم استخدموا **Retry Failed Shopify Bindings**. وشغّلوا **Check Webhooks** بعد تغيير النطاقات أو بيانات الاعتماد.

## الخطوات التالية

أضيفوا **Shopify Connector Accounting** لإدخال دفعات Shopify Payments ورسومها ونزاعاتها كقيود يومية (نشرح [سبب صعوبة مطابقة الدفعات](https://sawirstudio.com/ar/posts/why-shopify-payouts-dont-match-your-books/) في مقال منفصل)، وأضيفوا **Shopify Connector POS** لمبيعات المتاجر الفعلية.

إذا أردتم أن نتولى الإعداد، أو نضبطه بما يناسب طريقة عمل شركتكم، أو نوسّعه، فـ[هذا ما نقدمه](https://sawirstudio.com/ar/shopify/).

**تابعوا القراءة:** [لماذا لا تتطابق دفعات Shopify مع دفاتركم أبدًا](https://sawirstudio.com/ar/posts/why-shopify-payouts-dont-match-your-books/)، أو [واجهة Shopify 2026-10: ما الذي سيتعطل في تكاملاتكم](https://sawirstudio.com/ar/posts/shopify-2026-10-api-what-breaks/).
