---
title: Подключение Shopify к Odoo 19 Community | SawirStudio
url: https://sawirstudio.com/ru/posts/connect-shopify-to-odoo-19-community/
description: Как подключить магазин Shopify к Odoo 19 Community с помощью открытого коннектора Shopify от OCA: приложение в Dev Dashboard, scopes, queue_job, вебхуки и первичная настройка.
---

# Подключаем Shopify к Odoo 19 Community: пошаговая инструкция

Мы написали [коннектор Shopify для Odoo](https://sawirstudio.com/ru/projects/odoo-shopify-connector/) и передали его Odoo Community Association (OCA). Он бесплатный, с открытым исходным кодом (AGPL-3) и работает с Odoo 19 Community. Ниже описана настройка от начала до конца — так, как мы делаем её для клиентов.

Коннектор синхронизирует товары, остатки, клиентов, заказы, возвраты, возмещения и выполнение заказов — для одного магазина или для нескольких. Бухгалтерия и касса (POS) — отдельные модули, которые можно добавить позже.

## Перед началом

- Odoo 19 Community должна быть доступна по публичному HTTPS. Shopify отправляет на неё вебхуки и скачивает с неё изображения товаров.
- Доступ администратора к магазину Shopify и к организации Shopify, которой этот магазин принадлежит.
- Тариф Shopify, который включает защищённые данные клиентов для приложений, созданных вами самостоятельно. В тарифе Basic их нет, и без них Shopify отклоняет часть вебхуков по заказам, из-за чего импорт заказов и клиентов не работает.

## 1. Создайте приложение в Shopify Dev Dashboard

С 1 января 2026 года Shopify больше не позволяет создавать прежние «пользовательские приложения» (custom apps) из админки магазина. Создайте приложение в **Dev Dashboard** той организации, которой принадлежит магазин. Это важно: Odoo получает токен доступа через client credentials grant, а он работает только тогда, когда приложение и магазин принадлежат одной организации.

При заполнении версии приложения:

- **App URL:** `https://shopify.dev/apps/default-app-home`. У коннектора нет страницы внутри Shopify.
- **Embed app in Shopify admin:** выключено.
- **Use legacy install flow:** выключено, чтобы при установке приложения выдавались выбранные вами права доступа (scopes).
- **Webhooks:** не добавляйте. Odoo создаёт свои вебхуки сама при первичной настройке.
- **Webhooks API version:** `2026-01` — версия, которую использует коннектор.
- **Scopes:** `read_customers`, `write_customers`, `read_draft_orders`, `read_fulfillments`, `read_inventory`, `write_inventory`, `read_locations`, `read_merchant_managed_fulfillment_orders`, `write_merchant_managed_fulfillment_orders`, `read_orders`, `read_products`, `write_products` и `read_returns`. Добавьте `read_all_orders`, если нужны заказы старше 60 дней.

Опубликуйте версию, установите приложение в магазин и скопируйте client ID и client secret на странице Settings приложения.

Есть старое пользовательское приложение, созданное до 2026 года? Оно по-прежнему работает. Оставьте поле client ID пустым и укажите вместо него Admin API access token и API secret key этого приложения.

## 2. Установите модули в Odoo

Добавьте OCA `queue_job` и репозиторий [connector-shopify](https://github.com/OCA/connector-shopify) в addons path, затем установите **Shopify Connector**. Работа выполняется фоновыми задачами, поэтому `queue_job` нужно загружать на уровне всего сервера, указав мощность (capacity) для каждого канала Shopify:

```
[options]
server_wide_modules = base,rpc,web,queue_job

[queue_job]
channels = root:4,root.shopify.import:2,root.shopify.export:1,root.shopify.webhook:1
```

Затем задайте системному параметру `web.base.url` публичный HTTPS-адрес Odoo, а `web.base.url.freeze` — значение `True`. Иначе адрес меняется, когда администратор входит из другого места, и вебхуки Shopify уходят не туда.

## 3. Запустите первичную настройку

Добавьте сотрудников в группу **Shopify Manager**, затем откройте **Shopify > Configuration > Instances** и нажмите **Start Onboarding**:

1. Введите домен `yourstore.myshopify.com`, client ID и client secret и проверьте подключение. Odoo запрашивает токен и сама обновляет его каждые 24 часа.
2. Проверьте права доступа (scopes). Если какого-то не хватает, добавьте его в новую версию приложения, опубликуйте её и подтвердите в магазине.
3. Сопоставьте локации Shopify со складами Odoo, проверьте налоговые соответствия, которые создаёт Odoo, и выберите политики синхронизации.
4. Установите вебхуки и поставьте в очередь первые импорты.

Прежде чем позволить заказам и выполнению заказов работать без присмотра, проверьте соответствия платёжных шлюзов и служб доставки. Большинство неприятных сюрпризов потом возникает из-за соответствия, на которое никто не посмотрел.

## 4. Следите за работой

**Shopify > Dashboard** показывает состояние синхронизации, очередь задач, недавние ошибки и состояние вебхуков. Коннектор не рассчитывает, что каждый вебхук дойдёт: плановые задачи повторно импортируют всё, что изменилось в Shopify, но не попало в Odoo. Если что-то не сработало, устраните причину и нажмите **Retry Failed Shopify Bindings**. После смены домена или учётных данных запустите **Check Webhooks**.

## Что дальше

Добавьте **Shopify Connector Accounting**, чтобы выплаты Shopify Payments, комиссии и споры попадали в учёт в виде проводок ([почему выплаты так сложно сверить](https://sawirstudio.com/ru/posts/why-shopify-payouts-dont-match-your-books/), мы рассказываем отдельно), и **Shopify Connector POS** для продаж в офлайн-точках.

Если вы хотите, чтобы мы всё настроили, адаптировали под ваши бизнес-процессы или доработали, [мы этим занимаемся](https://sawirstudio.com/ru/shopify/).

**Читайте также:** [Почему выплаты Shopify никогда не совпадают с вашей бухгалтерией](https://sawirstudio.com/ru/posts/why-shopify-payouts-dont-match-your-books/) или [API Shopify 2026-10: что сломается в ваших интеграциях](https://sawirstudio.com/ru/posts/shopify-2026-10-api-what-breaks/).
