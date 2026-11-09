---
title: חיבור Shopify ל-Odoo 19 Community | SawirStudio
url: https://sawirstudio.com/he/posts/connect-shopify-to-odoo-19-community/
description: איך מחברים חנות Shopify ל-Odoo 19 Community באמצעות מחבר ה-Shopify בקוד פתוח של OCA: אפליקציית Dev Dashboard, scopes, queue_job, webhooks ו-onboarding.
---

# חיבור Shopify ל-Odoo 19 Community, שלב אחר שלב

כתבנו את [מחבר Shopify ל-Odoo](https://sawirstudio.com/he/projects/odoo-shopify-connector/) ותרמנו אותו ל-Odoo Community Association (OCA). המחבר חינמי, בקוד פתוח (AGPL-3) ועובד עם Odoo 19 Community. זו ההגדרה מההתחלה ועד הסוף, כפי שאנחנו עושים אותה עבור לקוחות.

הוא מסנכרן מוצרים, מלאי, לקוחות, הזמנות, החזרות, זיכויים ומשלוחים, לחנות אחת או למספר חנויות. הנהלת חשבונות ונקודת מכירה הם מודולים נפרדים שאפשר להוסיף בהמשך.

## לפני שמתחילים

- Odoo 19 Community שנגיש דרך HTTPS ציבורי. Shopify שולחת אליו webhooks ומורידה ממנו תמונות מוצר.
- גישת מנהל לחנות ב-Shopify ולארגון ב-Shopify שמחזיק בה.
- תוכנית Shopify שכוללת גישה למידע מוגן של לקוחות עבור אפליקציות שאתם יוצרים בעצמכם. תוכנית Basic לא כוללת זאת, ובלעדיה Shopify דוחה חלק מה-webhooks של הזמנות, וייבוא ההזמנות והלקוחות נכשל.

## 1. יצירת האפליקציה ב-Shopify Dev Dashboard

מאז 1 בינואר 2026 Shopify כבר לא מאפשרת ליצור את ה-“custom apps” הישנות מתוך ממשק הניהול של החנות. יוצרים את האפליקציה ב-**Dev Dashboard** של הארגון שמחזיק בחנות. זה חשוב: Odoo מקבל את אסימון הגישה שלו באמצעות client credentials grant, וזה עובד רק כשהאפליקציה והחנות שייכות לאותו ארגון.

כשממלאים את פרטי גרסת האפליקציה:

- **App URL:** `https://shopify.dev/apps/default-app-home`. למחבר אין עמוד בתוך Shopify.
- **Embed app in Shopify admin:** כבוי.
- **Use legacy install flow:** כבוי, כך שהתקנת האפליקציה מעניקה את ההרשאות (scopes) שבחרתם.
- **Webhooks:** ללא. Odoo יוצר אותם בעצמו במהלך ה-onboarding.
- **Webhooks API version:** `2026-01`, הגרסה שבה המחבר משתמש.
- **Scopes:** `read_customers`, `write_customers`, `read_draft_orders`, `read_fulfillments`, `read_inventory`, `write_inventory`, `read_locations`, `read_merchant_managed_fulfillment_orders`, `write_merchant_managed_fulfillment_orders`, `read_orders`, `read_products`, `write_products` ו-`read_returns`. מוסיפים את `read_all_orders` אם רוצים הזמנות ישנות מ-60 יום.

משחררים את הגרסה, מתקינים את האפליקציה בחנות ומעתיקים את ה-client ID ואת ה-client secret מעמוד ה-Settings של האפליקציה.

כבר יש לכם custom app ישנה מלפני 2026? היא עדיין עובדת. משאירים את שדה ה-client ID ריק ומשתמשים במקום זאת באסימון הגישה ל-Admin API ובמפתח הסודי של ה-API שלה.

## 2. התקנת המודולים ב-Odoo

מוסיפים את OCA `queue_job` ואת המאגר [connector-shopify](https://github.com/OCA/connector-shopify) ל-addons path, ואז מתקינים את **Shopify Connector**. העבודה רצה כמשימות רקע, ולכן `queue_job` חייב להיטען ברמת השרת, עם קיבולת לכל ערוץ Shopify:

```
[options]
server_wide_modules = base,rpc,web,queue_job

[queue_job]
channels = root:4,root.shopify.import:2,root.shopify.export:1,root.shopify.webhook:1
```

לאחר מכן מגדירים את פרמטר המערכת `web.base.url` לכתובת ה-HTTPS הציבורית של Odoo ואת `web.base.url.freeze` ל-`True`. אחרת הכתובת משתנה כשמנהל מתחבר ממקום אחר, וה-webhooks של Shopify נשלחים ליעד הלא נכון.

## 3. הרצת ה-onboarding

נותנים לצוות שלכם את הקבוצה **Shopify Manager**, ואז נכנסים אל **Shopify > Configuration > Instances** ולוחצים על **Start Onboarding**:

1. מזינים את הדומיין `yourstore.myshopify.com`, את ה-client ID ואת ה-client secret, ובודקים את החיבור. Odoo מבקש את האסימון ומחדש אותו בעצמו כל 24 שעות.
2. בודקים את ההרשאות (scopes). אם חסרה אחת, מוסיפים אותה לגרסה חדשה של האפליקציה, משחררים אותה ומאשרים אותה בחנות.
3. ממפים את המיקומים ב-Shopify למחסנים ב-Odoo, בודקים את מיפויי המס ש-Odoo מייצר ובוחרים את מדיניות הסנכרון.
4. מתקינים את ה-webhooks ומכניסים לתור את הייבואים הראשונים.

לפני שמאפשרים להזמנות ולמשלוחים לרוץ ללא פיקוח, בודקים את המיפויים של שערי התשלום ושל חברות המשלוח. רוב ההפתעות בהמשך מגיעות ממיפוי שאף אחד לא בדק.

## 4. לשמור על העניינים בעין

**Shopify > Dashboard** מציג את מצבי הסנכרון, את צבר המשימות, שגיאות אחרונות ואת תקינות ה-webhooks. המחבר לא מניח שכל webhook מגיע: משימות מתוזמנות מייבאות שוב כל דבר ש-Shopify שינתה ו-Odoo פספס. כשמשהו נכשל, מתקנים את הסיבה ואז משתמשים ב-**Retry Failed Shopify Bindings**. מריצים **Check Webhooks** אחרי שינוי דומיינים או פרטי גישה.

## השלבים הבאים

מוסיפים את **Shopify Connector Accounting** כדי להביא תשלומים (payouts), עמלות ומחלוקות של Shopify Payments כרשומות יומן (ההסבר [למה קשה להתאים payouts](https://sawirstudio.com/he/posts/why-shopify-payouts-dont-match-your-books/) נמצא בנפרד), ואת **Shopify Connector POS** למכירות בחנות הפיזית.

אם תרצו שנתקין, נגדיר לפי אופן העבודה של העסק שלכם או נרחיב את המחבר, [אנחנו עושים את זה](https://sawirstudio.com/he/shopify/).

**להמשך קריאה:** [למה ה-payouts של Shopify אף פעם לא תואמים לספרים שלכם](https://sawirstudio.com/he/posts/why-shopify-payouts-dont-match-your-books/), או [ה-API של Shopify בגרסה 2026-10: מה יישבר באינטגרציות שלכם](https://sawirstudio.com/he/posts/shopify-2026-10-api-what-breaks/).
