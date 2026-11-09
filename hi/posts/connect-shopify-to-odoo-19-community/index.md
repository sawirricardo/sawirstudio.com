---
title: Shopify को Odoo 19 Community से कनेक्ट करें | SawirStudio
url: https://sawirstudio.com/hi/posts/connect-shopify-to-odoo-19-community/
description: ओपन सोर्स OCA Shopify कनेक्टर की मदद से Shopify स्टोर को Odoo 19 Community से कैसे जोड़ें: Dev Dashboard ऐप, scopes, queue_job, वेबहुक और ऑनबोर्डिंग।
---

# Shopify को Odoo 19 Community से कनेक्ट करना, step by step

हमने [Odoo के लिए Shopify कनेक्टर](https://sawirstudio.com/hi/projects/odoo-shopify-connector/) बनाया और उसे Odoo Community Association (OCA) को योगदान के रूप में दिया। यह मुफ़्त और ओपन सोर्स (AGPL-3) है, और Odoo 19 Community के साथ काम करता है। यहाँ पूरा सेटअप शुरू से अंत तक दिया गया है, ठीक उसी तरह जैसे हम क्लाइंट्स के लिए करते हैं।

यह प्रोडक्ट, स्टॉक, कस्टमर, ऑर्डर, रिटर्न, रिफ़ंड और फ़ुलफ़िलमेंट सिंक करता है, चाहे एक स्टोर हो या कई। अकाउंटिंग और पॉइंट ऑफ़ सेल अलग मॉड्यूल हैं, जिन्हें आप बाद में जोड़ सकते हैं।

## शुरू करने से पहले

- Odoo 19 Community सार्वजनिक HTTPS पर उपलब्ध होना चाहिए। Shopify इसी पर वेबहुक भेजता है और यहीं से प्रोडक्ट की इमेज डाउनलोड करता है।
- Shopify स्टोर का एडमिन एक्सेस, और उस Shopify organization का भी, जिसके अंतर्गत वह स्टोर है।
- ऐसा Shopify प्लान जिसमें आपके खुद बनाए ऐप्स के लिए protected customer data शामिल हो। Basic प्लान में यह नहीं मिलता, और इसके बिना Shopify कुछ ऑर्डर वेबहुक को मना कर देता है, जिससे ऑर्डर और कस्टमर इम्पोर्ट फ़ेल हो जाते हैं।

## 1. Shopify Dev Dashboard में ऐप बनाएँ

1 जनवरी 2026 से Shopify स्टोर एडमिन से पुराने “custom apps” बनाने की सुविधा नहीं देता। ऐप को उस organization के **Dev Dashboard** में बनाएँ जिसके अंतर्गत स्टोर है। यह ज़रूरी है: Odoo अपना एक्सेस टोकन client credentials grant से लेता है, जो तभी काम करता है जब ऐप और स्टोर एक ही organization के हों।

ऐप वर्ज़न भरते समय:

- **App URL:** `https://shopify.dev/apps/default-app-home`। कनेक्टर का Shopify के अंदर कोई पेज नहीं है।
- **Embed app in Shopify admin:** बंद रखें।
- **Use legacy install flow:** बंद रखें, ताकि ऐप इंस्टॉल करने पर आपके चुने हुए scopes मिल जाएँ।
- **Webhooks:** कोई नहीं। Odoo ऑनबोर्डिंग के दौरान अपने वेबहुक खुद बना लेता है।
- **Webhooks API version:** `2026-01`, यानी वही वर्ज़न जो कनेक्टर इस्तेमाल करता है।
- **Scopes:** `read_customers`, `write_customers`, `read_draft_orders`, `read_fulfillments`, `read_inventory`, `write_inventory`, `read_locations`, `read_merchant_managed_fulfillment_orders`, `write_merchant_managed_fulfillment_orders`, `read_orders`, `read_products`, `write_products` और `read_returns`। अगर आपको 60 दिन से पुराने ऑर्डर चाहिए, तो `read_all_orders` भी जोड़ें।

वर्ज़न रिलीज़ करें, ऐप को स्टोर पर इंस्टॉल करें, और ऐप के Settings पेज से client ID और client secret कॉपी कर लें।

क्या आपके पास 2026 से पहले का कोई पुराना custom app है? वह अब भी काम करता है। client ID खाली छोड़ें और उसके बजाय उसका Admin API access token और API secret key इस्तेमाल करें।

## 2. Odoo में मॉड्यूल इंस्टॉल करें

OCA का `queue_job` और [connector-shopify](https://github.com/OCA/connector-shopify) रिपॉज़िटरी अपने addons path में जोड़ें, फिर **Shopify Connector** इंस्टॉल करें। काम बैकग्राउंड जॉब के रूप में चलता है, इसलिए `queue_job` को server-wide लोड करना होगा, और हर Shopify चैनल के लिए capacity तय करनी होगी:

```
[options]
server_wide_modules = base,rpc,web,queue_job

[queue_job]
channels = root:4,root.shopify.import:2,root.shopify.export:1,root.shopify.webhook:1
```

फिर `web.base.url` सिस्टम पैरामीटर को Odoo के सार्वजनिक HTTPS पते पर सेट करें और `web.base.url.freeze` को `True` करें। ऐसा न करने पर किसी एडमिन के दूसरी जगह से लॉग इन करते ही पता बदल जाता है, और Shopify के वेबहुक गलत जगह चले जाते हैं।

## 3. ऑनबोर्डिंग चलाएँ

अपनी टीम को **Shopify Manager** ग्रुप दें, फिर **Shopify > Configuration > Instances** खोलें और **Start Onboarding** पर क्लिक करें:

1. `yourstore.myshopify.com` डोमेन, client ID और client secret डालें, और कनेक्शन टेस्ट करें। Odoo टोकन माँगता है और हर 24 घंटे में उसे अपने आप रिन्यू कर लेता है।
2. Scopes जाँच लें। अगर कोई scope छूटा है, तो उसे नए ऐप वर्ज़न में जोड़ें, रिलीज़ करें और स्टोर पर उसे मंज़ूरी दें।
3. Shopify लोकेशन को Odoo वेयरहाउस से मैप करें, Odoo जो टैक्स मैपिंग बनाता है उन्हें देख लें, और सिंक पॉलिसी चुनें।
4. वेबहुक इंस्टॉल करें और पहले इम्पोर्ट कतार में लगाएँ।

ऑर्डर और फ़ुलफ़िलमेंट को बिना निगरानी के चलने देने से पहले, पेमेंट गेटवे और कैरियर की मैपिंग जाँच लें। आगे चलकर ज़्यादातर परेशानियाँ ऐसी मैपिंग से आती हैं जिसे किसी ने देखा ही नहीं था।

## 4. नज़र बनाए रखें

**Shopify > Dashboard** में सिंक की स्थिति, जॉब बैकलॉग, हाल की एरर और वेबहुक की सेहत दिखती है। कनेक्टर यह मानकर नहीं चलता कि हर वेबहुक पहुँच ही जाएगा: शेड्यूल्ड जॉब हर वह बदलाव दोबारा इम्पोर्ट कर लेते हैं जो Shopify में हुआ पर Odoo से छूट गया। कुछ फ़ेल होने पर पहले कारण ठीक करें, फिर **Retry Failed Shopify Bindings** इस्तेमाल करें। डोमेन या क्रेडेंशियल बदलने के बाद **Check Webhooks** चलाएँ।

## अगले कदम

Shopify Payments के पेआउट, फ़ीस और डिस्प्यूट को जर्नल एंट्री के रूप में लाने के लिए **Shopify Connector Accounting** जोड़ें (पेआउट मिलाना मुश्किल क्यों है, यह हम [अलग से समझाते हैं](https://sawirstudio.com/hi/posts/why-shopify-payouts-dont-match-your-books/)), और स्टोर में होने वाली बिक्री के लिए **Shopify Connector POS** जोड़ें।

अगर आप चाहते हैं कि इसे आपके लिए सेट अप किया जाए, आपके बिज़नेस के तरीके के हिसाब से कॉन्फ़िगर किया जाए या आगे बढ़ाया जाए, तो [हम यह करते हैं](https://sawirstudio.com/hi/shopify/)।

**आगे पढ़ें:** [आपके Shopify पेआउट आपकी बुक्स से क्यों नहीं मिलते](https://sawirstudio.com/hi/posts/why-shopify-payouts-dont-match-your-books/), या [Shopify का 2026-10 API: आपके इंटीग्रेशन में क्या टूटेगा](https://sawirstudio.com/hi/posts/shopify-2026-10-api-what-breaks/)।
