---
title: موصل Shopify لـ Odoo | SawirStudio
url: https://sawirstudio.com/ar/projects/odoo-shopify-connector/
description: موصل Shopify مفتوح المصدر لـ Odoo 19 Community، كتبته SawirStudio وتتم صيانته ضمن Odoo Community Association (OCA).
---

# موصل Shopify لـ Odoo

![مستودع OCA connector-shopify على GitHub](https://sawirstudio.com/assets/images/projects/odoo-shopify-connector.webp)

## عن المشروع

جمعية Odoo Community Association (OCA) هي منظمة غير ربحية تتولى صيانة الوحدات مفتوحة المصدر التي يعتمد عليها معظم مستخدمي Odoo Community. و Connector Shopify هو الموصل الذي تقدمه الجمعية بين متاجر Shopify و Odoo 19 Community، ويُصدر بموجب رخصة AGPL-3.

## ماذا أنجزنا

كتبنا هذا الموصل وساهمنا به في OCA في يوليو 2026، ويشمل: Shopify Connector الأساسي، و Shopify Connector Accounting، و Shopify Connector POS. ونواصل العمل عليه، وكان آخر ما أنجزناه مواءمته مع Shopify Admin API إصدار 2026-01، كما يبني عليه الآن مساهمون آخرون في OCA، منهم وحدة خاصة بوسائط منتجات التجارة الإلكترونية.

- **مزامنة التجارة.** المنتجات والمخزون والعملاء والطلبات والمرتجعات والمبالغ المستردة والشحن، لأي عدد من متاجر Shopify في قاعدة بيانات Odoo الواحدة، مع فصلها حسب الشركة.
- **موثوقية في صلب التصميم.** webhooks موقّعة لسرعة الاستجابة، و queue\_job من OCA لمعالجة دائمة لا تضيع، ومطابقة مجدولة تصلح الأحداث الفائتة أو المتأخرة.
- **أدوات للعمل اليومي.** إعداد أولي، وفحوصات سلامة، ولوحة متابعة للمزامنة، واستعادة بعد الأخطاء، وإعدادات قابلة للضبط لمدة الاحتفاظ بالبيانات.
- **المحاسبة.** مدفوعات Shopify Payments، ومعاملات الرصيد، والرسوم، والنزاعات، والتزامات بطاقات الهدايا، والمطابقة متعددة العملات.
- **نقاط البيع.** استيراد مبيعات Shopify POS كطلبات بيع، وتوجيهها من كل موقع بيع بالتجزئة إلى مستودع في Odoo، دون الحاجة إلى تطبيق نقاط البيع الخاص بـ Odoo.

## كيف تناولنا المشروع

التكامل الذي يكون صحيحاً في أغلب الأحيان لا يصلح للمحاسبة. لذلك يتعامل الموصل مع Shopify GraphQL Admin API، ولا يفترض أبداً أن webhook قد وصل، ويراجع نفسه مقابل Shopify وفق جدول زمني. وتُخطَّط قيود اليومية بناءً على المبالغ الدقيقة القادمة من Shopify، ويجب أن تكون متوازنة قبل ترحيلها: وعندما لا تتطابق الأرقام، يُظهر الموصل خطأً واضحاً بدلاً من إضافة سطر موازنة بصمت.

هل ترغبون في إعداده أو تطويره لمتاجركم؟ اطلعوا على [تكاملات Shopify](https://sawirstudio.com/ar/shopify/)، أو [Odoo وأنظمة الأعمال](https://sawirstudio.com/ar/odoo/) للاطلاع على خدماتنا مع Odoo نفسه.

**بُني باستخدام** Python و Odoo 19 و Shopify GraphQL Admin API

- **المشروع** Odoo Community Association (OCA)
- **الرخصة** AGPL-3
- **الشيفرة المصدرية** <https://github.com/OCA/connector-shopify>

## مشاريع أخرى

[![MyVivaldy](https://sawirstudio.com/assets/images/projects/myvivaldy.webp)

### MyVivaldy](https://sawirstudio.com/ar/projects/myvivaldy/)
[![تطبيق Koperasi HOB](https://sawirstudio.com/assets/images/projects/koperasi-hob-application.webp)

### تطبيق Koperasi HOB](https://sawirstudio.com/ar/projects/koperasi-hob-application/)
[![تطبيق DBTV](https://sawirstudio.com/assets/images/projects/dbtv-application.webp)

### تطبيق DBTV](https://sawirstudio.com/ar/projects/dbtv-application/)

يكفي حديثًا عن أنفسنا.

## كيف يمكننا مساعدتكم؟

[ابدؤوا مشروعكم](https://sawirstudio.com/ar/contact/)
