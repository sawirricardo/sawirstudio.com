---
title: Sawir Live Translate | SawirStudio
url: https://sawirstudio.com/zh-hans/projects/wordpress-plugins/sawir-live-translate/
description: Sawir Live Translate 就地翻译您的页面：选择语言，自动翻译可见内容，保存后再为访客提供语言切换器。
---

# Sawir Live Translate

![Sawir Live Translate 横幅图](https://sawirstudio.com/assets/images/plugins/sawir-live-translate-banner.webp)

Sawir Live Translate 就地翻译您的页面：选择语言，自动翻译可见内容，保存后再为访客提供语言切换器。

[在 WordPress.org 免费获取](https://wordpress.org/plugins/sawir-live-translate/)

## 功能介绍

打开您网站上的任意页面，点击管理栏中的“翻译此页面”，选择目标语言，然后可以自动翻译所有可见文字（通过免费的机器翻译服务），也可以手动编辑译文，同时实时看到页面的变化。

保存后的译文会在服务器端直接渲染到 HTML 响应中（对 SEO 友好：不会闪现未翻译的文字，head 中还有 hreflang 备用链接），访客则通过基于链接的切换器（位于右下角）切换语言。

URL 使用语言路径前缀，例如西班牙语是 /es/about，源语言则是 /about（仍为规范地址，不带前缀）。语言按以下顺序确定：/es/ 路径 → ?lang=es → 已保存的 cookie → 浏览器 Accept-Language → 源语言。根路径和不带前缀的 URL 会自动检测访客的浏览器语言，而不会进行重定向（并会发送 Vary 头，让共享缓存正常工作）。

1. 进入“设置 → Live Translate”：设置源语言，并勾选要提供的语言。2. 访问前台页面，点击管理栏中的“🌐 翻译此页面”。3. 选择语言，点击“自动翻译空白项”（或手动编辑文字），然后点击“保存”。4. 访客使用切换器（右下角）切换语言，每种语言都有各自的 /xx/ 网址。

注意：语言路径前缀要求安装在根域名下，并且没有任何以两位语言代码作为别名（slug）的顶级页面。

## 常见问题

### 与其他方案相比如何？

WPML、Polylang 和 TranslatePress 等多语言框架把译文内容视为网站结构的一等公民，提供按语言划分的编辑流程、翻译服务，并与页面构建器和 WooCommerce 深度集成。Sawir Live Translate 则刻意做得更轻量：打开一个页面，翻译可见的内容（自动或手动），再通过带语言前缀的 URL 和 hreflang 标签提供出来。如果您要搭建完整的多语言网站，包括翻译菜单、产品和编辑流程，请使用这些框架。如果您只想让现有页面今天下午就能以另一种语言阅读，这个插件就能做到。

### 切换了语言，但没有任何变化

这几乎总是意味着该语言还没有保存任何译文。请查看“设置 → Live Translate → 已保存的译文”：如果列表中没有该语言（或显示 0 条文字），请再次打开页面编辑器，自动翻译后点击“保存”。在您保存之前，编辑器中的实时预览不会被保存。

### 有些文字没有被翻译

页内编辑器只能看到您打开的那个页面上的文字。位于其他页面上的文字，需要把那些页面也翻译一遍，或者运行全站扫描。由滑块或脚本在页面加载后注入的文字，可能需要通过“编辑全站文字”手动添加。

### 我修改了内容，译文却消失了

译文是与原文的完整文字一一对应的。如果您改写了某个句子，只需把新的版本翻译一次并保存即可。

## 推荐搭配

- [Sawir Insights](https://sawirstudio.com/zh-hans/projects/wordpress-plugins/sawir-insights/)：展示访客如何使用您的网站，提供热图、会话录制、问卷调查、转化漏斗和表单分析，并且所有数据都保存在您自己的服务器上。
- [Sawir Media Folders](https://sawirstudio.com/zh-hans/projects/wordpress-plugins/sawir-media-folders/)：为 WordPress 媒体库带来文件夹功能，数量不限，通过拖放即可整理归类。

查看我们所有的[免费 WordPress 插件](https://sawirstudio.com/zh-hans/projects/wordpress-plugins/)。

- **价格** 免费
- **版本** 1.1.0
- **要求** WordPress 5.8、PHP 7.4
- **许可证** GPL-2.0-or-later
- **WordPress.org** <https://wordpress.org/plugins/sawir-live-translate/>

我们先聊到这里。

## 我们能为您 做些什么？

[开始项目](https://sawirstudio.com/zh-hans/contact/)
