---
title: 让我们的网站对智能体友好 | SawirStudio
url: https://sawirstudio.com/zh-hans/posts/making-our-sites-agent-ready/
description: 我们如何让产品网站可供 AI 智能体读取：为每个页面提供 Markdown 副本、Link 响应头、llms.txt、robots.txt 中的 Content Signals、API 目录和 Agent Skills。
---

# 让我们的网站对智能体友好，我们做了什么

如今，访问网站的不只是人，越来越多的是 AI 智能体：它们读取页面来回答某个人的问题，或者代表某个人去使用一款产品。本周，我们检查了本站和旗下的产品网站，包括 [SawirPost](https://sawirstudio.com/zh-hans/projects/sawirpost/)、[Sawirly](https://sawirstudio.com/zh-hans/projects/sawirly/)、[Rindraw](https://sawirstudio.com/zh-hans/projects/rindraw/)、[Math Daily](https://sawirstudio.com/zh-hans/projects/math-daily/) 等，让智能体更容易阅读和使用它们。

我们以 [isitagentready.com](https://isitagentready.com) 作为检查清单。它会按等级为网站评分，下面是每一步我们做了什么。

## 为每个页面提供 Markdown 副本

智能体读取 HTML 时，要穿过导航、脚本和页脚，才能找到寥寥几段正文。因此，本站的每个页面现在都在旁边附带一份 Markdown 副本，位于 `index.md`，只包含页面的主要内容。副本由一个小脚本根据 HTML 生成，所以内容不会与原页面脱节。

智能体可以通过两种方式获取：在 URL 后加上 `index.md`，或者发送 `Accept: text/markdown` 请求头。我们在静态文件前面放了一个很小的 Cloudflare Worker，它会检查这个请求头，并返回 Markdown 副本，同时在 `x-markdown-tokens` 响应头中给出大致的 token 数量。浏览器看到的仍然是正常的页面。

```
curl -H "Accept: text/markdown" https://sawirstudio.com/shopify/
```

## Link 响应头

现在，每个 HTML 页面都会通过带有 `rel="alternate"` 的 HTTP `Link` 响应头，说明它的 Markdown 副本在哪里。首页还会指向 `llms.txt` 和站点地图。智能体只需一次请求就能找到所有内容，不必猜测 URL。

## llms.txt

`llms.txt` 是面向语言模型的纯文本网站索引：一行简介，加上每个页面及其描述的列表，并链接到对应的 Markdown 副本。我们的 llms.txt 由同一个脚本根据站点地图生成。

## Content Signals

`robots.txt` 只说明谁可以抓取，没有说明抓取后可以如何使用内容。[Content Signals](https://contentsignals.org/) 补上了这一点。我们的只有一行，并且允许全部三种用途，因为我们希望被找到、被引用：

```
Content-Signal: ai-train=yes, search=yes, ai-input=yes
```

这三步让所有网站都达到了第 3 级。

## API 目录和 Agent Skills

SawirPost 和 Sawirly 提供 API，所以我们做得更进一步。两者都在 `/.well-known/api-catalog` 发布了 API 目录（RFC 9727）：这是一份很小的 JSON 文档，指向 OpenAPI 描述和文档。智能体进入网站后，不用读任何页面就能找到 API。

它们还发布了 Agent Skills，也就是告诉智能体如何用这款产品完成常见任务的简短说明。这让两者都达到了第 4 级。

## 我们停在了哪里

第 5 级要求智能体能够通过 OAuth 智能体注册流程，自行注册并完成身份验证。仅有一个页面说明如何用电子邮件和密码注册，是不算数的。这是一项需要真正开发的功能，还涉及实际的安全问题，所以我们打算认真考虑清楚，而不是仓促上线。

## 值得吗？

对于静态网站来说，这些工作大部分一个下午就能完成：一个脚本、一个小小的 Worker，再加上 `robots.txt` 里的一行。对访问网站的人来说，什么都没有改变。我们认为，会有越来越多的问题由智能体通过阅读像我们这样的网站来回答，而一份干净的 Markdown 页面，是被正确引用的最简单方式。

我们也会把同样的配置作为[上线与增长](https://sawirstudio.com/zh-hans/launch-and-growth/)服务的一部分，为客户完成。

**继续阅读：** [借助 Abilities API，让您的 WordPress 插件可供 AI 智能体使用](https://sawirstudio.com/zh-hans/posts/wordpress-abilities-api-for-plugins/)，或者[我们首页上可以抛掷的应用图标](https://sawirstudio.com/zh-hans/posts/throwable-app-icons-matter-js/)。
