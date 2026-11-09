---
title: Sawir Insights | SawirStudio
url: https://sawirstudio.com/zh-hans/projects/wordpress-plugins/sawir-insights/
description: 热力图、会话录制、问卷调查、漏斗和表单分析。所有数据都保存在您的服务器上。
---

# Sawir Insights

![Sawir Insights 横幅图](https://sawirstudio.com/assets/images/plugins/sawir-insights-banner.webp)

Sawir Insights 让您了解访客如何使用您的网站，提供热力图、会话录制、问卷调查、转化漏斗和表单分析，所有数据都保存在您自己的服务器上。

[在 WordPress.org 免费获取](https://wordpress.org/plugins/sawir-insights/)

## 功能介绍

Sawir Insights 是适用于 WordPress 的用户行为分析插件：包含热力图、会话录制、问卷调查、反馈小部件、表单分析、转化漏斗，以及把这些数据汇总在一起的仪表盘。

在多站点网络中，每个站点各自独立地进行追踪和配置分析。

简单来说：它是 Hotjar 或 Microsoft Clarity 的自托管替代方案，访客数据保存在您自己的数据库中。

做这个插件的原因很简单。这类工具中成熟的产品确实不错，但它们的工作方式，是把访客的一举一动记录下来，发送到别人的服务器上。对很多网站来说，尤其是欧盟地区的网站，这种做法轻则让人不安，重则根本行不通。我想要的是同类功能的工具，而且数据不会离开收集它的网站。所以这个插件围绕一个原则来构建：您的数据库、您的上传目录、您的保留规则。它背后没有 SaaS 服务，不需要注册账号，也不会有任何形式的回传连接。

### 功能介绍

- 点击、移动和滚动热力图，直接渲染在您的真实页面上，并按桌面端、平板和手机分别查看。
- 会话录制，可在 wp-admin 中回放。输入内容始终会被遮蔽；录制功能默认关闭，需要您手动开启。
- 弹出式问卷和常驻的反馈标签，可按网址、设备、流量比例、页面停留时间、滚动深度或离开意图进行定向。
- 表单分析：了解用户在哪里犹豫、哪些字段留空、在哪里放弃。只记录选择器和时间，绝不记录任何人输入的内容。
- 按顺序设置的转化漏斗，并提供每日报告。
- 仪表盘包含流量关键指标和周期对比、每日趋势图，以及设备、浏览器、页面、来源、推广活动、漏斗和录制的细分数据。

### 它刻意不做的事

我认为，在安装一个工具之前，您应该先知道它拒绝做什么。

- 它不会存储 IP 地址。无论是哈希还是截断处理，一概不存。
- 它绝不会记录表单内容。采集层只发送元数据，所以不存在泄露的可能。
- 它不会从 CDN 加载任何内容，也不会向 CDN 发送任何内容。rrweb 回放库已内置在插件中。
- 它不会永久保留原始数据。事件会按照您控制的保留计划进行汇总和清理；录制文件有单独的保留期限，并设有磁盘容量上限，超出时优先删除最旧的文件，因为悄悄占满主机磁盘的插件不是好插件。

前端追踪脚本是小巧的原生 JavaScript，通过 `sendBeacon` 批量发送，并提供采样控制。未被抽中参与录制的访客，根本不会下载回放库。

### 隐私立场

默认遵循“请勿跟踪”（Do Not Track）设置。同意模式 API（`window.swinsConsent`）可以与您已在使用的同意管理工具集成。会话、事件和录制都已接入 WordPress 自带的个人数据导出和删除工具。您的具体部署是否符合 GDPR，取决于您的告知内容和法律依据，这一部分任何插件都无法代劳，但数据最小化的工作已经完成。

### 支持

您可以通过该插件在 wordpress.org 上的论坛获取支持，也可以发送邮件至 ricardo@sawirstudio.com。

## 常见问题

### 与其他方案相比如何？

Hotjar 和 Microsoft Clarity 这样的托管分析服务都是很出色的工具，免费额度也很慷慨，但它们的工作方式是把您访客的行为记录发送到它们的服务器上。Sawir Insights 舍弃了这种跨站点的基础设施，只坚持一个原则：所有数据都在您自己的网站上采集、存储和查看。如果您想要托管服务的额外功能，并且不介意第三方处理数据，请选择托管服务；如果访客数据离开您的服务器是不可接受的，请选择 Sawir Insights。录制和热力图数据保存在您自己的数据库和上传目录中，因此这里的保留设置比在 SaaS 中更重要。

### Sawir Insights 符合 GDPR 吗？

这个插件专为注重隐私的部署场景而设计：数据保留在您的服务器上，从不存储 IP 地址，保留期限可配置，遵循“请勿跟踪”设置，并集成了 WordPress 的导出和删除工具。合规本身取决于您的配置、告知内容、法律依据和同意机制的实现，请就您的网站咨询专业的法律顾问。

### 它与 Hotjar 或 Microsoft Clarity 有什么不同？

那些都是托管服务：功能强大，但您访客的行为会按照它们的条款记录到它们的服务器上。Sawir Insights 覆盖同一类功能，包括热力图、会话录制、问卷调查、漏斗和表单分析，但它是一个把所有数据保留在您自己服务器上的插件，不需要账号，不依赖外部服务，也不共享数据。

### 访客数据会离开我的服务器吗？

不会。这个插件不会向 Sawir Studio 或任何分析服务商发送行为数据。事件保存在您的 WordPress 数据库中，录制文件保存在受保护的上传目录下。

## 推荐搭配

- [Sawir Pulse](https://sawirstudio.com/zh-hans/projects/wordpress-plugins/sawir-pulse/)：以角落里的小通知形式，向顾客展示您店铺里真实发生的动态，例如最新订单、正在浏览商品的人数和最新评价。
- [Sawir Live Translate](https://sawirstudio.com/zh-hans/projects/wordpress-plugins/sawir-live-translate/)：就地翻译您的页面：选择语言，自动翻译可见内容，保存后还能为访客提供语言切换器。

查看我们所有的[免费 WordPress 插件](https://sawirstudio.com/zh-hans/projects/wordpress-plugins/)。

- **价格** 免费
- **版本** 1.0.0
- **要求** WordPress 6.2、PHP 7.4
- **许可证** GPLv2 或更高版本
- **WordPress.org** <https://wordpress.org/plugins/sawir-insights/>

我们先聊到这里。

## 我们能为您 做些什么？

[开始项目](https://sawirstudio.com/zh-hans/contact/)
