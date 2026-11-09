---
title: 使用 Matter.js 制作可抛掷的应用图标 | SawirStudio
url: https://sawirstudio.com/zh-hans/posts/throwable-app-icons-matter-js/
description: 我们如何制作首页上可以抛掷的应用图标盒子：用 Matter.js 实现物理效果，用 DOM 元素而不是 canvas，图标旋转时阴影保持不动，以及相应的降级方案。
---

# 我们首页上可抛掷的应用图标

我们的[首页](https://sawirstudio.com/zh-hans/)上有一个装着应用图标的盒子。您一打开页面，图标就会掉进盒子里，您可以拿起它们，再扔出去。轻点一个图标而不拖动，就会打开对应的应用页面。这个效果只花了一个下午，但其中有几处细节比我们预想的更有意思。

## 为什么是一个盒子

最初的想法是让图标在整个页面顶部飞来飞去。但标题和按钮就在那里，图标落在“Start a project”上面，为了一点乐趣而影响使用，不划算。把盒子放在介绍内容旁边，边缘加一个标签，就能把玩法限制在一定范围内。它也正好填上了原本就有的空白：宽屏上介绍文字和按钮右侧的空位。在手机上，它位于按钮下方。

## 真实的物理效果，来自 Matter.js

圆形互相碰撞，自己写很容易。但应用图标是圆角方形，方形像硬币一样滚动，一眼就会觉得不对劲。要做出真实的翻滚和堆叠，需要一个真正的物理引擎，所以我们用了 [Matter.js](https://brm.io/matter-js/)。每个图标都是一个圆角矩形（`chamfer`），圆角与图标 22% 的圆角半径一致，放在四面固定的墙之内。

这个引擎有 83KB，所以只在首页加载，只在页面准备好之后加载，而且只在访客没有要求减少动态效果时加载。

## 使用 DOM 元素，而不是 canvas

Matter.js 可以绘制到 canvas 上，但我们只用它来做计算。图标是带图片的普通链接，每一帧都通过 CSS transform 来定位：

```
st.setProperty("--x", (b.position.x - size / 2) + "px");
st.setProperty("--y", (b.position.y - size / 2) + "px");
st.setProperty("--a", b.angle + "rad");
```

这样，轻点就是一次真正的链接跳转，图片在任何缩放比例下都清晰，图标的边框和样式也与网站其他地方保持一致。

## 始终不动的阴影

我们的风格是给所有元素加一个朝右下方的硬阴影。如果把它加在会旋转的元素上，阴影也会跟着转，旋转图标的阴影就会像时钟指针一样甩来甩去。解决办法是每一帧都把阴影的偏移量朝相反方向旋转：

```
// keep the shadow at (3px, 3px) in screen space while the icon turns by a
st.setProperty("--shadow-x", 3 * (Math.cos(a) + Math.sin(a)) + "px");
st.setProperty("--shadow-y", 3 * (Math.cos(a) - Math.sin(a)) + "px");
```

## 抓取与抛掷

按下指针时，我们会在指针和您抓住的图标位置之间连一根弹簧（Matter.js 的约束），连接点是您抓住的那个确切位置，而不是图标中心，所以抓住一个角时，图标会自然地摆动。松手后图标会保持当时的速度，但设有上限，免得用力一甩让它穿墙而出。如果指针移动不到六像素，就算作轻点，链接随之打开。

图标从上方掉落时，盒子顶部是敞开的，等它们全部进入后才会关上，这样抛掷时就不会有图标从顶部飞出去。

## 一些小细节

- 图标重重落地时，页面的点状背景会泛起一圈涟漪。
- 当所有图标静止时，模拟会进入休眠；盒子移出屏幕时会暂停，所以您阅读页面时它不占用任何资源。
- 只有当盒子出现在屏幕上时，图标才会掉落。在手机上，就是您滚动到盒子的时候。
- 如果没有 JavaScript，或者开启了减少动态效果，图标会整齐地堆在一起静止不动，而且依然是链接。

这是一件小事，但它比一句话更能说明我们做的应用是什么样的。如果您也想在自己的网站上加上类似的效果，[告诉我们](https://sawirstudio.com/zh-hans/contact/)。

**继续阅读：** [我们如何让网站为智能体做好准备](https://sawirstudio.com/zh-hans/posts/making-our-sites-agent-ready/)，或者 [WordPress 7.2 将于 12 月 8 日发布：我们在插件中要测试什么](https://sawirstudio.com/zh-hans/posts/wordpress-7-2-what-plugin-authors-should-test/)。
