---
title: Shopify Connector for Odoo | SawirStudio
url: https://sawirstudio.com/ru/projects/odoo-shopify-connector/
description: Open-source-коннектор Shopify для Odoo 19 Community, написанный SawirStudio и поддерживаемый в Odoo Community Association (OCA).
---

# Shopify Connector for Odoo

![Репозиторий OCA connector-shopify на GitHub](https://sawirstudio.com/assets/images/projects/odoo-shopify-connector.webp)

## О проекте

Odoo Community Association (OCA) — некоммерческая организация, которая поддерживает open-source-модули, на которые опирается большинство пользователей Odoo Community. Connector Shopify — её коннектор между магазинами Shopify и Odoo 19 Community, распространяемый по лицензии AGPL-3.

## Что мы сделали

Мы написали этот коннектор и передали его в OCA в июле 2026 года: основной Shopify Connector, Shopify Connector Accounting и Shopify Connector POS. Мы продолжаем работать над ним, недавно привели его в соответствие с Admin API Shopify версии 2026-01, а другие участники OCA уже строят на его основе свои решения, в том числе модуль для медиафайлов товаров в eCommerce.

- **Синхронизация торговых данных.** Товары, остатки, клиенты, заказы, возвраты, возмещения и выполнение заказов — для любого числа магазинов Shopify в одной базе Odoo с разделением по компаниям.
- **Надёжность по замыслу.** Подписанные webhook для скорости, queue\_job от OCA для надёжной обработки и регулярная сверка, которая восстанавливает пропущенные или задержавшиеся события.
- **Инструменты на каждый день.** Подключение, проверки состояния, панель синхронизации, восстановление после ошибок и настраиваемый срок хранения данных.
- **Бухгалтерия.** Выплаты Shopify Payments, балансовые транзакции, комиссии, споры, обязательства по подарочным картам и сверка в нескольких валютах.
- **Розничные продажи.** Продажи через Shopify POS импортируются как заказы на продажу и направляются с каждой торговой точки на нужный склад Odoo — приложение Odoo POS для этого не требуется.

## Как мы подошли к задаче

Интеграция, которая работает правильно в большинстве случаев, для бухгалтерии не подходит. Поэтому коннектор обращается к GraphQL Admin API Shopify, никогда не считает, что webhook точно пришёл, и по расписанию сверяет себя с Shopify. Проводки формируются по точным суммам из Shopify и должны сходиться до отправки в учёт: если что-то не сходится, коннектор выдаёт понятную ошибку, а не добавляет втихую балансирующую строку.

Хотите, чтобы мы настроили или доработали коннектор для ваших магазинов? Смотрите [интеграции с Shopify](https://sawirstudio.com/ru/shopify/) или [Odoo и бизнес-системы](https://sawirstudio.com/ru/odoo/), если вас интересует сам Odoo.

**Технологии** Python, Odoo 19 и Shopify GraphQL Admin API

- **Проект** Odoo Community Association (OCA)
- **Лицензия** AGPL-3
- **Исходный код** <https://github.com/OCA/connector-shopify>

## Другие проекты

[![MyVivaldy](https://sawirstudio.com/assets/images/projects/myvivaldy.webp)

### MyVivaldy](https://sawirstudio.com/ru/projects/myvivaldy/)
[![Приложение Koperasi HOB](https://sawirstudio.com/assets/images/projects/koperasi-hob-application.webp)

### Приложение Koperasi HOB](https://sawirstudio.com/ru/projects/koperasi-hob-application/)
[![Приложение DBTV](https://sawirstudio.com/assets/images/projects/dbtv-application.webp)

### Приложение DBTV](https://sawirstudio.com/ru/projects/dbtv-application/)

Хватит о нас.

## Чем мы можем Вам помочь?

[Начать проект](https://sawirstudio.com/ru/contact/)
