---
title: Sawir Social Scheduler | SawirStudio
url: https://sawirstudio.com/zh-hans/projects/wordpress-plugins/sawir-social-scheduler/
description: 只写一次，即可在 wp-admin 中定时发布到 27 个网络。支持按渠道修改、串帖、媒体和预览。您的令牌始终保存在您的服务器上。
---

# Sawir Social Scheduler

![Sawir Social Scheduler 横幅](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-banner.webp)

Sawir Social Scheduler 让您在 wp-admin 中只写一次帖子，就能定时发布到 27 个社交网络，并支持按渠道单独修改、串帖、媒体和预览。您的账号令牌始终保存在您自己的服务器上。

[在 WordPress.org 免费获取](https://wordpress.org/plugins/sawir-social-scheduler/)

## 功能介绍

Sawir Social Scheduler 把 WordPress 变成一个社交媒体定时发布工具。用各平台提供的令牌连接您的渠道，写好帖子，选择渠道，设定日期，然后发布。每次发送都在后台运行，并返回帖子链接或平台自己给出的错误信息。

中间没有任何托管服务。凭证使用您网站的 salt 加密，绝不会离开您的服务器。

### 渠道

渠道数量不限，每个平台也可以添加多个账号。每个渠道都有一个**测试连接** 按钮，会获取账号名称和头像，让您在正式使用前确认令牌确实有效。

**社交：** X (Twitter)、Facebook 主页、Instagram、Threads、LinkedIn 主页、Bluesky、Mastodon、TikTok、YouTube、Pinterest、Tumblr、Dribbble、Google 商家资料

**社区：** Telegram、Discord、Slack、Reddit、Lemmy、VK、Kick、Twitch、Whop

**博客和邮件通讯：** WordPress、Medium、Dev.to、Hashnode、Listmonk

可以按客户或品牌把渠道分组，支持嵌套分组，一键即可选中整个分组。

### 帖子

- **一份内容，多个渠道。** 用复选框选择渠道。未做改动的渠道会按原文发布。
- **按渠道自定义。** 在任意渠道上勾选“自定义”，即可为它设置单独的内容：文字、媒体和延迟。这些内容会预先填入帖子的原有内容，您修改的只是副本，其他渠道仍按原文发布。
- **串帖。** 添加多个部分，即可在 X、Threads、Bluesky、Mastodon 和 Telegram 上发布串帖，每个部分都可以带自己的媒体。不支持串帖的平台会收到合并后的内容。
- **媒体。** 图片和视频只需添加一次。每个平台会得到它所支持的形式：相册、轮播、Reels、可续传的视频上传。
- **定时。** 可以使用 WordPress 常规的“发布时间”日期，也可以为任意渠道单独设定时间。
- **发布时段。** 为渠道设置常用的发布时间。“下一个空闲时段”会填入下一个尚未使用的时段，日历会显示接下来的空闲时段，点击即可填写，API 也接受 `"date": "next"`。
- **预览。** 每个渠道都有实时预览，外观与真实平台一致，并已套用字数限制和媒体规则，让您确切看到将要发出的内容。
- **评论和延迟。** 在不支持串帖的平台上，多出的部分会作为评论发布在帖子下方。任何一个部分都可以在上一部分之后等待指定的分钟数再发布。
- **分享博客文章。** 在常规文章编辑器中有一个选项框，文章发布时会发布到您的渠道，并可使用单独的文字和媒体。
- **复制。** 把帖子连同其渠道、各部分和自定义内容复制为新的草稿。您可以把草稿留作模板。
- **提及。** 输入 `@handle`，网络会在 X、Instagram、Threads、Mastodon、Telegram、Reddit 和 Bluesky 上解析它，链接和用户名会变成真正的提及。LinkedIn 公司主页写作 `@[Name](org:12345)`，其中的 ID 取自主页网址；Facebook 主页写作 `@[12345]`。
- **重复。** 每隔 N 天或 N 周重新发送常青内容，可以持续到某个日期，或重复指定次数。每次重复都是独立的帖子，有各自的发送记录和统计数据。请在文字中加入 `{n}` 为每次重复编号，因为 X、Threads 和 Bluesky 会拒绝完全相同的文字。

### 日历

每个帖子会拆分为每个渠道一次发送，日历把每次发送显示为一张卡片，并按状态着色：已排期、已发送、失败、已暂停、草稿、进行中。拖动卡片即可改动该渠道的发布时间。点击卡片可以查看错误信息、重试按钮、实时链接和互动数据。点击空白时段，即可在该时间新建帖子。提供月、周、日和列表视图，可按分组和状态筛选。

### 数据分析

在每次发送后的 1 小时、1 天和 1 周，以及按需，插件会收集 16 个平台的点赞、评论、分享、浏览和曝光数据。每个渠道的关注者数量每天记录一次，并附带 30 天的走势小图。这些数据会显示在日历卡片、帖子的“数据分析”框和渠道列表中。“数据分析”页面会汇总最近 7 天、30 天或 90 天的情况，并与上一个周期对比：已发送的帖子、互动、覆盖人数和关注者及其变化，另有关注者图表、各渠道互动情况和表现最好的帖子。

### 客户审批

勾选“审批通过前暂停发送”，复制预览链接，发给您的客户。客户无需登录，就能看到帖子在每个渠道上的实际效果，并点击“批准”或附上备注“请求修改”。批准后才会放行发送。每一次决定都会保留为历史记录。

向 X、YouTube 和 TikTok 上传大视频时，会分成短小的片段，在多个后台任务之间接续进行，因此不受主机时间限制的影响，也不会把整个文件加载到内存中。

### REST API

后台能做的一切，都可以通过 `/wp-json/sawir-social/v1` 使用，方便用于自动化、移动应用和其他网站。使用应用密码进行身份验证；权限遵循插件的权限设置，渠道凭证绝不会返回。

- `GET /platforms`、`GET /channels`、`GET /channels/{id}`、`POST /channels/{id}/test`
- `GET|POST /posts`、`GET|PATCH|DELETE /posts/{id}`、`POST /posts/{id}/duplicate`、`GET /posts/{id}/analytics`
- `GET /sends`、`GET /sends/{id}`、`POST /sends/{id}/retry`、`POST /sends/{id}/stats`

用 `{"title","content","status":"future","date":"2026-12-25 09:00","media":[12],"parts":[{"text","media","delay"}],"channels":[{"id":38,"parts":[{"text","media","delay"}],"date":"…"}],"groups":["Client A"],"hold":false}` 创建帖子。渠道的 `parts` 会替换该渠道上帖子的原有内容；`text` 是单个部分的简写。

### AI 智能体

同样的操作也已注册到 WordPress Abilities API，名称为 `sawir-social/list-channels`、`create-post`、`update-post`、`list-posts`、`list-sends`、`retry-send`、`test-channel`、`get-analytics` 和 `list-platforms`，并带有类型化的输入和输出模式。任何通过 WordPress MCP 适配器连接的智能体都会把它们视为工具，也可以通过 `/wp-json/wp-abilities/v1/abilities/{name}/run` 调用。Abilities API 是 WordPress 6.9 及更新版本的一部分；对于旧版本，插件自带了它，因此只要插件在运行，这些功能就可以使用。

### 系统要求

如果您的网站有 Action Scheduler（WooCommerce 和许多其他插件都会提供），发送会通过它在后台运行，否则通过 WP-Cron 运行。使用 WordPress 内置的 cron 时，只有网站有访问量才会推进。如需及时发送，请添加系统 cron，日历页面会显示适合您网站的具体命令。

对于通过网址获取媒体的平台（Instagram、Threads、Facebook、Pinterest、Telegram），媒体必须能从互联网访问。本地开发站点无法向它们提供图片。

### 致谢

内置 FullCalendar (MIT)、WordPress Abilities API 库 (GPL)，以及来自 Simple Icons (CC0) 的品牌标志。当其他插件提供 Action Scheduler 时，会使用它。

## 截图

![Sawir Social Scheduler 截图 1](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-screenshot-1.webp)

![Sawir Social Scheduler 截图 2](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-screenshot-2.webp)

![Sawir Social Scheduler 截图 3](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-screenshot-3.webp)

![Sawir Social Scheduler 截图 4](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-screenshot-4.webp)

![Sawir Social Scheduler 截图 5](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-screenshot-5.webp)

![Sawir Social Scheduler 截图 6](https://sawirstudio.com/assets/images/plugins/sawir-social-scheduler-screenshot-6.webp)

## 常见问题

### 我从哪里获取令牌？

在各平台的开发者设置中。渠道表单会写明所需凭证的具体名称，例如 Facebook 主页访问令牌、Telegram 机器人令牌或 Bluesky 应用密码。最好使用长期有效的令牌；Meta 和 LinkedIn 的令牌大约 60 天后过期，需要重新粘贴。

### 我的 Telegram 测试显示“chat not found”

私有频道没有 @用户名。请把聊天 ID 留空，然后点击“测试连接”：插件会列出该机器人已加入的所有聊天，并给出可以粘贴的数字 ID。

### 为什么有一次发送一直停在“进行中…”？

后台队列还没有运行。请访问一下网站，或者设置日历页面上显示的系统 cron。

### 我能在本地站点向 Instagram 发布视频吗？

不能。Instagram 通过网址获取媒体，需要公开的地址。插件的其他功能在本地都可以正常使用。

## 推荐搭配

- [Sawir Insights](https://sawirstudio.com/zh-hans/projects/wordpress-plugins/sawir-insights/)：展示访客如何使用您的网站，提供热图、会话录制、问卷调查、转化漏斗和表单分析，并且所有数据都保存在您自己的服务器上。
- [Sawir Live Translate](https://sawirstudio.com/zh-hans/projects/wordpress-plugins/sawir-live-translate/)：就地翻译您的页面：选择语言，自动翻译可见内容，保存后还能为访客提供语言切换器。

查看我们所有的[免费 WordPress 插件](https://sawirstudio.com/zh-hans/projects/wordpress-plugins/)。

- **价格** 免费
- **版本** 0.1.0
- **要求** WordPress 5.8、PHP 7.4
- **许可证** GPL-2.0-or-later
- **WordPress.org** <https://wordpress.org/plugins/sawir-social-scheduler/>

我们先聊到这里。

## 我们能为您 做些什么？

[开始项目](https://sawirstudio.com/zh-hans/contact/)
