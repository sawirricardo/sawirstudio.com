---
title: הכנת האתרים שלנו לסוכני AI | SawirStudio
url: https://sawirstudio.com/he/posts/making-our-sites-agent-ready/
description: איך הפכנו את אתרי המוצרים שלנו לקריאים עבור סוכני AI: עותקי Markdown של כל דף, כותרות Link, llms.txt, Content Signals ב-robots.txt, קטלוג API ו-Agent Skills.
---

# מה נדרש כדי להכין את האתרים שלנו לסוכני AI

יותר ויותר מהביקורים באתר אינם של אנשים. אלה סוכני AI שקוראים דפים כדי לענות על שאלה של מישהו, או שמנסים להשתמש במוצר בשמו. השבוע עברנו על האתר הזה ועל אתרי המוצרים שלנו, [SawirPost](https://sawirstudio.com/he/projects/sawirpost/), [Sawirly](https://sawirstudio.com/he/projects/sawirly/), [Rindraw](https://sawirstudio.com/he/projects/rindraw/), [Math Daily](https://sawirstudio.com/he/projects/math-daily/) ועוד, כדי שיהיה קל יותר לסוכנים לקרוא אותם ולהשתמש בהם.

השתמשנו ב-[isitagentready.com](https://isitagentready.com) כרשימת בדיקה. האתר מדרג אתרים לפי רמות, וכאן מה שנדרש בכל שלב.

## עותק Markdown של כל דף

סוכן שקורא HTML צריך לעבור דרך תפריטים, סקריפטים ופוטרים כדי למצוא כמה פסקאות. לכן לכל דף באתר הזה יש עכשיו עותק Markdown לצידו, בכתובת `index.md`, עם התוכן הראשי של הדף בלבד. סקריפט קטן בונה אותם מתוך ה-HTML, כך שהם אף פעם לא יוצאים מסנכרון.

סוכנים מקבלים אותם בשתי דרכים: על ידי הוספת `index.md` לכתובת, או על ידי בקשה עם `Accept: text/markdown`. Cloudflare Worker קטנטן שעומד לפני הקבצים הסטטיים בודק את הכותרת הזאת ומחזיר את עותק ה-Markdown, יחד עם הערכה גסה של מספר הטוקנים בכותרת `x-markdown-tokens`. גולשים בדפדפן ממשיכים לקבל את הדף הרגיל.

```
curl -H "Accept: text/markdown" https://sawirstudio.com/shopify/
```

## כותרות Link

כל דף HTML מציין עכשיו איפה נמצא עותק ה-Markdown שלו, בכותרת HTTP מסוג `Link` עם `rel="alternate"`. דף הבית מצביע גם על `llms.txt` ועל מפת האתר. סוכן יכול למצוא הכול מבקשה אחת, בלי לנחש כתובות.

## llms.txt

`llms.txt` הוא אינדקס בטקסט פשוט של האתר עבור מודלי שפה: סיכום בשורה אחת ורשימה של כל הדפים עם התיאור שלהם, עם קישורים לעותקי ה-Markdown. אצלנו הוא נוצר ממפת האתר על ידי אותו סקריפט.

## Content Signals

`robots.txt` קובע מי רשאי לסרוק, לא מה מותר לעשות עם מה שנקרא. [Content Signals](https://contentsignals.org/) מוסיף את זה. אצלנו זו שורה אחת, והיא מתירה את שלושת השימושים, כי אנחנו רוצים שימצאו אותנו ויצטטו אותנו:

```
Content-Signal: ai-train=yes, search=yes, ai-input=yes
```

שלושת השלבים האלה הביאו את כל האתרים שלנו לרמה 3.

## קטלוג API ו-Agent Skills

ל-SawirPost ול-Sawirly יש API, אז הם הרחיקו לכת יותר. כל אחד מהם מפרסם קטלוג API בכתובת `/.well-known/api-catalog` (RFC 9727): מסמך JSON קטן שמצביע על תיאור ה-OpenAPI ועל התיעוד. סוכן שנוחת באתר יכול למצוא את ה-API בלי לקרוא אף דף.

הם מפרסמים גם Agent Skills, הוראות קצרות לסוכנים איך לבצע משימות נפוצות עם המוצר. זה הביא את שניהם לרמה 4.

## איפה עצרנו

רמה 5 דורשת שסוכנים יוכלו להירשם ולהזדהות בעצמם, באמצעות תהליך רישום סוכנים מבוסס OAuth. דף שמתאר איך נרשמים עם אימייל וסיסמה לא נחשב. זו יכולת אמיתית שצריך לבנות, עם שאלות אבטחה אמיתיות, ולכן אנחנו חושבים עליה בנחת במקום למהר.

## האם זה היה שווה את זה?

באתר סטטי, רוב העבודה לקחה אחר צהריים אחד: סקריפט אחד, Worker קטן אחד ושורה ב-`robots.txt`. שום דבר לא השתנה עבור אנשים שמבקרים באתר. אנחנו מאמינים שיותר ויותר שאלות ייענו על ידי סוכנים שקוראים אתרים כמו שלנו, ודף Markdown נקי הוא הדרך הקלה ביותר להיות מצוטטים נכון.

אנחנו מגדירים את אותו הדבר גם ללקוחות, כחלק מ-[Launch & growth](https://sawirstudio.com/he/launch-and-growth/).

**המשיכו לקרוא:** [איך להפוך את תוסף הוורדפרס שלכם לשמיש עבור סוכני AI עם ה-Abilities API](https://sawirstudio.com/he/posts/wordpress-abilities-api-for-plugins/), או [האייקונים הניתנים לזריקה בדף הבית שלנו](https://sawirstudio.com/he/posts/throwable-app-icons-matter-js/).
