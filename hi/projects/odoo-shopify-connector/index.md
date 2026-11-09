---
title: Odoo के लिए Shopify Connector | SawirStudio
url: https://sawirstudio.com/hi/projects/odoo-shopify-connector/
description: Odoo 19 Community के लिए एक ओपन-सोर्स Shopify कनेक्टर, जिसे SawirStudio ने लिखा है और जो Odoo Community Association (OCA) में मेंटेन किया जाता है।
---

# Odoo के लिए Shopify Connector

![GitHub पर OCA की connector-shopify रिपॉज़िटरी](https://sawirstudio.com/assets/images/projects/odoo-shopify-connector.webp)

## प्रोजेक्ट के बारे में

Odoo Community Association (OCA) एक नॉनप्रॉफ़िट संस्था है, जो उन ओपन-सोर्स मॉड्यूल्स को मेंटेन करती है जिन पर ज़्यादातर Odoo Community यूज़र्स निर्भर करते हैं। Connector Shopify इसी का कनेक्टर है, जो Shopify स्टोर्स को Odoo 19 Community से जोड़ता है। यह AGPL-3 लाइसेंस के तहत जारी किया गया है।

## हमने क्या बनाया

यह कनेक्टर हमने लिखा और जुलाई 2026 में OCA को योगदान के रूप में दिया: मुख्य Shopify Connector, Shopify Connector Accounting और Shopify Connector POS। हम इस पर लगातार काम करते हैं; हाल ही में हमने इसे Shopify के 2026-01 Admin API के अनुरूप किया है। अब OCA के दूसरे कंट्रीब्यूटर भी इसी पर आगे काम कर रहे हैं, जिनमें eCommerce प्रोडक्ट मीडिया के लिए बना एक मॉड्यूल भी शामिल है।

- **कॉमर्स सिंक।** प्रोडक्ट, इन्वेंट्री, कस्टमर, ऑर्डर, रिटर्न, रिफ़ंड और फ़ुलफ़िलमेंट। एक Odoo डेटाबेस में कितने भी Shopify स्टोर जोड़े जा सकते हैं, और हर स्टोर कंपनी के हिसाब से अलग रहता है।
- **शुरू से भरोसेमंद।** तेज़ी के लिए साइन किए गए वेबहुक, भरोसेमंद प्रोसेसिंग के लिए OCA का queue\_job, और शेड्यूल्ड रिकॉन्सिलिएशन, जो छूटे हुए या देर से आए इवेंट ठीक कर देता है।
- **रोज़मर्रा के टूल।** ऑनबोर्डिंग, हेल्थ चेक, सिंक डैशबोर्ड, एरर रिकवरी और अपनी ज़रूरत के हिसाब से डेटा रिटेंशन।
- **अकाउंटिंग।** Shopify Payments के पेआउट, बैलेंस ट्रांज़ैक्शन, फ़ीस, डिस्प्यूट, गिफ़्ट-कार्ड देनदारी और मल्टी-करेंसी रिकॉन्सिलिएशन।
- **पॉइंट ऑफ़ सेल।** Shopify POS की बिक्री सेल ऑर्डर के रूप में इम्पोर्ट होती है और हर रिटेल लोकेशन से Odoo के वेयरहाउस तक पहुँचती है। इसके लिए Odoo का अपना POS ऐप ज़रूरी नहीं है।

## हमने इसे कैसे किया

जो इंटीग्रेशन ज़्यादातर समय सही चले, वह अकाउंटिंग के लिए ग़लत है। इसलिए यह कनेक्टर Shopify के GraphQL Admin API से बात करता है, यह कभी नहीं मानता कि वेबहुक पहुँच ही गया होगा, और तय समय पर Shopify से अपना डेटा मिलाकर जाँच करता है। जर्नल एंट्री Shopify की सटीक रकम से बनाई जाती हैं और पोस्ट होने से पहले उनका संतुलित होना ज़रूरी है: अगर कहीं हिसाब नहीं मिलता, तो कनेक्टर चुपचाप कोई बैलेंसिंग लाइन जोड़ने के बजाय साफ़ एरर दिखाता है।

क्या आप इसे अपने स्टोर्स के लिए सेट अप कराना या बढ़वाना चाहते हैं? [Shopify इंटीग्रेशन](https://sawirstudio.com/hi/shopify/) देखें, या Odoo के लिए [Odoo और बिज़नेस सिस्टम](https://sawirstudio.com/hi/odoo/) देखें।

**इनसे बना** Python, Odoo 19 और Shopify GraphQL Admin API

- **प्रोजेक्ट** Odoo Community Association (OCA)
- **लाइसेंस** AGPL-3
- **सोर्स कोड** <https://github.com/OCA/connector-shopify>

## और प्रोजेक्ट

[![MyVivaldy](https://sawirstudio.com/assets/images/projects/myvivaldy.webp)

### MyVivaldy](https://sawirstudio.com/hi/projects/myvivaldy/)
[![Koperasi HOB Application](https://sawirstudio.com/assets/images/projects/koperasi-hob-application.webp)

### Koperasi HOB Application](https://sawirstudio.com/hi/projects/koperasi-hob-application/)
[![DBTV ऐप्लिकेशन](https://sawirstudio.com/assets/images/projects/dbtv-application.webp)

### DBTV ऐप्लिकेशन](https://sawirstudio.com/hi/projects/dbtv-application/)

हमारी बात बहुत हुई।

## हम आपकी कैसे मदद करें?

[प्रोजेक्ट शुरू करें](https://sawirstudio.com/hi/contact/)
