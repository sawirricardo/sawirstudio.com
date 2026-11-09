---
title: Shopify Connector for Odoo | SawirStudio
url: https://sawirstudio.com/he/projects/odoo-shopify-connector/
description: מחבר Shopify בקוד פתוח עבור Odoo 19 Community, שנכתב על ידי SawirStudio ומתוחזק ב-Odoo Community Association (OCA).
---

# Shopify Connector for Odoo

![מאגר OCA connector-shopify ב-GitHub](https://sawirstudio.com/assets/images/projects/odoo-shopify-connector.webp)

## הפרויקט

Odoo Community Association (OCA) הוא הארגון ללא מטרות רווח שמתחזק את המודולים בקוד פתוח שרוב משתמשי Odoo Community נשענים עליהם. Connector Shopify הוא המחבר שלו בין חנויות Shopify לבין Odoo 19 Community, ומשוחרר ברישיון AGPL-3.

## מה בנינו

אנחנו כתבנו את המחבר ותרמנו אותו ל-OCA ביולי 2026: את ליבת Shopify Connector, את Shopify Connector Accounting ואת Shopify Connector POS. אנחנו ממשיכים לעבוד עליו, ולאחרונה התאמנו אותו ל-Admin API של Shopify בגרסת 2026-01. גם תורמים אחרים ב-OCA בונים עליו כיום, בהם מודול למדיה של מוצרים ב-eCommerce.

- **סנכרון מסחר.** מוצרים, מלאי, לקוחות, הזמנות, החזרות, זיכויים ומשלוחים, עבור כל מספר של חנויות Shopify בכל מסד נתונים של Odoo, בהפרדה לפי חברה.
- **אמינות מהיסוד.** Webhooks חתומים למהירות, queue\_job של OCA לעיבוד עמיד, והתאמה מתוזמנת שמתקנת אירועים שהוחמצו או התעכבו.
- **כלים ליום-יום.** הטמעה, בדיקות תקינות, לוח בקרה לסנכרון, התאוששות משגיאות ושמירת נתונים הניתנת להגדרה.
- **הנהלת חשבונות.** תשלומי Shopify Payments, תנועות יתרה, עמלות, מחלוקות תשלום, התחייבויות בגין כרטיסי מתנה והתאמות רב-מטבעיות.
- **נקודות מכירה.** מכירות Shopify POS מיובאות כהזמנות מכירה ומנותבות מכל חנות פיזית למחסן ב-Odoo, בלי צורך באפליקציית ה-POS של Odoo עצמה.

## איך ניגשנו לזה

אינטגרציה שנכונה רוב הזמן היא שגויה מבחינת הנהלת חשבונות. לכן המחבר פונה ל-GraphQL Admin API של Shopify, אף פעם לא מניח ש-webhook הגיע, ובודק את עצמו מול Shopify בתזמון קבוע. רשומות היומן מתוכננות לפי הסכומים המדויקים של Shopify וחייבות להתאזן לפני הרישום: כשמשהו לא מסתדר, המחבר מציג שגיאה ברורה ולא מוסיף בשקט שורת איזון.

רוצים שנגדיר או נרחיב אותו עבור החנויות שלכם? ראו [אינטגרציות Shopify](https://sawirstudio.com/he/shopify/), או [Odoo ומערכות עסקיות](https://sawirstudio.com/he/odoo/) לעבודה על Odoo עצמו.

**נבנה עם** Python, Odoo 19 ו-Shopify GraphQL Admin API

- **פרויקט** Odoo Community Association (OCA)
- **רישיון** AGPL-3
- **קוד מקור** <https://github.com/OCA/connector-shopify>

## עוד פרויקטים

[![MyVivaldy](https://sawirstudio.com/assets/images/projects/myvivaldy.webp)

### MyVivaldy](https://sawirstudio.com/he/projects/myvivaldy/)
[![אפליקציית Koperasi HOB](https://sawirstudio.com/assets/images/projects/koperasi-hob-application.webp)

### אפליקציית Koperasi HOB](https://sawirstudio.com/he/projects/koperasi-hob-application/)
[![אפליקציית DBTV](https://sawirstudio.com/assets/images/projects/dbtv-application.webp)

### אפליקציית DBTV](https://sawirstudio.com/he/projects/dbtv-application/)

מספיק עלינו.

## איך נוכל לעזור לכם?

[בואו נתחיל פרויקט](https://sawirstudio.com/he/contact/)
