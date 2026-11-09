---
title: Sawir Renew – WooCommerce 定期扣款外掛 | SawirStudio
url: https://sawirstudio.com/zh-hant/projects/wordpress-plugins/sawir-renew/
description: WooCommerce 完整訂閱功能：定期扣款、試用期、註冊費、暫停／略過、方案切換、催繳、智慧重試，全部免費。
---

# Sawir Renew

![Sawir Renew 橫幅圖片](https://sawirstudio.com/assets/images/plugins/sawir-renew-banner.webp)

Sawir Renew 能讓 WooCommerce 商店轉型為訂閱制事業：定期扣款、試用期、註冊費與智慧重試一應俱全，沒有任何功能被付費方案鎖住。

[到 WordPress.org 免費取得](https://wordpress.org/plugins/sawir-renew/)

## 功能介紹

不必額外付費購買擴充套件，也能用 WooCommerce 販售訂閱。所有功能一律免費，沒有專業版，也沒有任何功能被鎖住。

### 商品與定價

- 支援單一與可變訂閱商品類型
- 可設定每 N 天、週、月或年扣款一次，並支援特價
- 免費試用、註冊費，以及付款 N 次後自動到期的訂閱
- 同步續訂（所有人都在每月 1 日續訂，或您指定的任一天），第一期可選擇按比例計費或免費
- 一次性運費，以及「每位顧客限一個有效訂閱」的限制

### 兩種付款方式

- **按期付款（適用所有金流）：** 續訂到期時，顧客會收到附有付款連結的帳單，可用銀行轉帳、電子錢包、貨到付款，或您商店支援的任何方式付款。非常適合不習慣信用卡自動扣款的市場。
- **透過 Stripe 自動扣款：** 如果您使用官方的 WooCommerce Stripe Gateway，顧客結帳時會儲存信用卡，續訂時自動在背景扣款，並包含智慧重試；當信用卡需要驗證時，也會提供備用的付款連結。其他金流也可以接上同一套轉接介面（請見常見問題）。

### 續訂不會半夜吵醒您

- 所有排程皆由 Action Scheduler 執行，並每天自動檢查修復
- 可自訂的催繳流程：重試間隔、提醒時程、寬限期、自動取消
- 測試站／複製站偵測：網站網址一變更，就立即暫停扣款，不會再因為被遺忘的測試站副本而重複扣款
- 自動扣款前會先寄送續訂通知信

### 讓顧客自己處理

顧客可在「我的帳號 → 訂閱」中取消訂閱（立即取消或期末取消，由您決定）、暫停與恢復、略過一次續訂、提前續訂、變更付款方式或配送地址，也能在不同方案之間升級或降級，升級時按比例收費。

### 定期優惠券

提供四種訂閱優惠券：定期百分比折扣、定期固定金額折扣（可限定只適用前 N 次付款），以及註冊費折扣。WooCommerce 內建的優惠券仍只套用於首筆訂單。

### 通知

共有十三種可開關的電子郵件（顧客與管理員），涵蓋訂閱生命週期的每個環節。開發人員也能透過 `swsb_notified` hook，將任何事件同步到其他管道（WhatsApp、簡訊、Slack 等）。

### 為新版 WooCommerce 打造

原生支援 HPOS（訂閱資料存放在 WooCommerce 訂單儲存區，不使用自訂資料表），相容購物車／結帳區塊，管理後台的清單與編輯頁面沿用 WooCommerce 內建的訂單畫面，另有每小時快取的報表頁面（有效訂閱、MRR、流失率、即將到來的續訂、續訂營收），以及 REST API。

## 常見問題

### 這個外掛與 WooCommerce 或 Automattic 有關係嗎？

不是。「WooCommerce」是 Automattic 的商標；依照其命名準則，這是一款獨立開發、*適用於* WooCommerce 的外掛。

### 哪些金流支援自動續訂？

開箱即用的是官方的 WooCommerce Stripe Gateway。其他金流都以按期付款模式運作，也就是每一期都會寄給顧客付款連結。開發人員可以透過兩個 filter，為任何金流加入自動扣款支援：

`swsb_order_supports_auto`：當母訂單的金流已儲存可重複使用的付款方式時，回傳 true。`swsb_charge_renewal`：在背景對續訂訂單扣款，並回傳 `true` 或 `WP_Error`。

### 可以和 WooCommerce Subscriptions 同時使用嗎？

不行。同時執行兩套扣款引擎會讓顧客被重複扣款，因此如果偵測到 WooCommerce Subscriptions 已啟用，本外掛會停用自己的引擎並顯示提示。

### 在測試站上會發生什麼事？

如果網站網址變更，扣款與顧客通知信會自動暫停，直到您確認這份副本是測試站還是正式站。選擇「正式站」後，所有排程都會重新建立。

## 推薦搭配

- [Sawir Referral](https://sawirstudio.com/zh-hant/projects/wordpress-plugins/sawir-referral/)：讓您的顧客幫您帶朋友來。
- [Sawir Pulse](https://sawirstudio.com/zh-hant/projects/wordpress-plugins/sawir-pulse/)：以角落的小通知，讓購物者看到您商店裡實際發生的事，例如最新訂單、正在瀏覽商品的人數和最新評價。

歡迎瀏覽我們所有的[免費 WordPress 外掛](https://sawirstudio.com/zh-hant/projects/wordpress-plugins/)。

- **價格** 免費
- **版本** 1.0.0
- **系統需求** WordPress 6.0、PHP 7.4、WooCommerce
- **授權** GPL-2.0-or-later
- **WordPress.org** <https://wordpress.org/plugins/sawir-renew/>

關於我們就先說到這裡。

## 有什麼我們 能幫上忙的嗎？

[開始專案](https://sawirstudio.com/zh-hant/contact/)
