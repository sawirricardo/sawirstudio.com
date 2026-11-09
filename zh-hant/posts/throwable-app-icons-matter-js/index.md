---
title: 使用 Matter.js 製作可拋擲的 App 圖示 | SawirStudio
url: https://sawirstudio.com/zh-hant/posts/throwable-app-icons-matter-js/
description: 我們如何製作首頁上可以拋擲 App 圖示的盒子：用 Matter.js 做物理效果、以 DOM 元素取代 canvas、讓圖示旋轉時陰影保持不動，以及備援方案。
---

# 我們首頁上可以拋擲的 App 圖示

我們的[首頁](https://sawirstudio.com/zh-hant/)有一個裝著 App 圖示的盒子。您一進來，圖示就會掉進盒子裡，您可以拿起它們、丟出去。輕點圖示而不拖曳，就會開啟該 App 的頁面。這個功能花了一個下午完成，其中有幾個細節比我們預期的還有意思。

## 為什麼是一個盒子

我們最初的想法，是讓圖示在整個頁面上方到處飛。但標題和按鈕就在那裡，圖示掉到「Start a project」上面，為了一點趣味而犧牲可用性，並不划算。所以我們在介紹文字旁放了一個盒子，邊緣加上標籤，把玩樂的範圍限制在裡面。這也剛好填補了原本就空著的位置：寬螢幕上，介紹文字和按鈕右側的空白處。在手機上，盒子則位於按鈕下方。

## Matter.js 帶來的真實物理效果

圓形互相碰撞，自己寫並不難。但 App 圖示是圓角方形，方形像硬幣一樣滾動，一看就不對勁。要做出自然的翻滾和堆疊，需要真正的物理引擎，所以我們選用了 [Matter.js](https://brm.io/matter-js/)。每個圖示都是一個圓角矩形（`chamfer`），圓角大小與圖示 22% 的圓角半徑一致，並放在四面固定的牆壁之內。

這個引擎大小為 83KB，所以只在首頁載入，而且要等頁面準備好之後才載入，若訪客設定了減少動態效果，則完全不會載入。

## 使用 DOM 元素，而不是 canvas

Matter.js 可以繪製到 canvas 上，但我們只拿它來做運算。圖示是一般的連結加圖片，每一幀都用 CSS transform 來定位：

```
st.setProperty("--x", (b.position.x - size / 2) + "px");
st.setProperty("--y", (b.position.y - size / 2) + "px");
st.setProperty("--a", b.angle + "rad");
```

因此，輕點就是真正的連結，圖片在任何縮放比例下都清晰，圖示的邊框和樣式也與網站其他地方一致。

## 固定不動的陰影

我們的風格是在所有元素的右下方加上硬邊陰影。如果把它加在會旋轉的元素上，陰影也會跟著轉，旋轉中的圖示，其陰影就會像時鐘指針一樣甩來甩去。解決方法是在每一幀把陰影的偏移量朝相反方向旋轉：

```
// keep the shadow at (3px, 3px) in screen space while the icon turns by a
st.setProperty("--shadow-x", 3 * (Math.cos(a) + Math.sin(a)) + "px");
st.setProperty("--shadow-y", 3 * (Math.cos(a) - Math.sin(a)) + "px");
```

## 抓取與拋擲

指標按下時，我們會在指標與您抓住的圖示位置之間加上一條彈簧（Matter.js 的 constraint），這個位置是您實際抓取的點，而不是圖示中心，所以抓著邊角時，圖示會自然地擺動。放開後圖示會保留原本的速度，但設有上限，避免用力一甩就穿牆而出。如果指標移動不到六個像素，就視為輕點，並開啟連結。

圖示從上方掉落時，盒子的頂部是開著的，等到全部進入後才會關上，所以拋擲時不會有圖示從上方飛出去。

## 小細節

- 用力落地時，頁面的點狀背景會泛起一圈漣漪。
- 當所有圖示都靜止時，模擬會進入休眠；盒子移出畫面時則會暫停，所以您閱讀時不會耗費任何資源。
- 只有盒子進入畫面後，圖示才會掉落。在手機上，就是當您捲動到盒子的位置時。
- 如果沒有 JavaScript，或是開啟了減少動態效果，圖示會整齊地堆在一起靜止不動，而且依然是可以點擊的連結。

這只是一件小事，但比起一句介紹，更能說明我們做的 App 是什麼樣子。如果您也想在自己的網站上加入類似的效果，歡迎[與我們聯絡](https://sawirstudio.com/zh-hant/contact/)。

**繼續閱讀：** [我們如何讓網站為 AI 代理做好準備](https://sawirstudio.com/zh-hant/posts/making-our-sites-agent-ready/)，或是[WordPress 7.2 將於 12 月 8 日推出：我們在外掛中要測試的項目](https://sawirstudio.com/zh-hant/posts/wordpress-7-2-what-plugin-authors-should-test/)。
