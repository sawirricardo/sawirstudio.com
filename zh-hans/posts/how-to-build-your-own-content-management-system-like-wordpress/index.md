---
title: 自己动手搭建类似 WordPress 的 CMS | SawirStudio
url: https://sawirstudio.com/zh-hans/posts/how-to-build-your-own-content-management-system-like-wordpress/
description: 使用 Laravel 和现代软件架构，搭建类似 WordPress 的内容管理系统。
---

# 如何自己搭建内容管理系统（类似 WordPress）

从零开始搭建内容管理系统，是一段非常棒的经历。

WordPress 一直是事实上的标准，尤其是在内容管理系统领域。官网上是这样介绍的……

> WordPress 驱动着超过 43% 的网站，而且这个数字每天都在增长。从简单的网站、博客，到复杂的门户和企业网站，甚至各类应用，都可以用 WordPress 来搭建。

不过，WordPress 诞生的年代，开发应用还没有现在这么简单。

如今，随着近期 [PHP 更新](https://www.php.net/)的推出，PHP 开发已经大大提速，软件开发实践也在不断演进。

例如，受 Ruby on Rails 启发的 MVC 结构开发方式，以及一些软件架构原则，已经被许多软件框架采用。

但 WordPress 并没有采用这些做法，它依然使用自己的一套结构。

因此，本文希望帮助您了解：如果像 WordPress 这样的内容管理系统采用现代软件架构，可以做成什么样子。

为此，我们将使用 [Laravel](https://laravel.com)，这是一个应用了不少软件工程实践的 PHP Web 框架。

先简单说说我们的内容管理系统需要哪些功能。参考 WordPress 的功能，我们可以提炼出以下几个关键功能：

- 轻松发布 如果您写过文档，那么用 WordPress 创建内容对您来说就是小菜一碟。您可以创建文章和页面，轻松设置格式、插入媒体，点一下按钮，内容就会上线。
- 发布工具 WordPress 让内容管理变得简单。您可以创建草稿、定时发布、查看文章修订记录。可以将内容设为公开或私密，也可以用密码保护文章和页面。
- 用户管理 并不是每个人都需要相同的网站权限。管理员负责管理站点，编辑负责处理内容，作者和投稿者负责撰写内容，订阅者则拥有可自行管理的个人资料。这样既能让不同的贡献者参与您的网站，也能让其他人只是作为社区的一员。
- 媒体管理 俗话说一图胜千言，所以能够快速轻松地把图片和媒体上传到 WordPress 非常重要。把媒体文件拖放到上传器中，即可添加到您的网站。您可以添加替代文字和说明文字，把图片和图库插入内容中。我们甚至还加入了一些图片编辑工具，供您随意把玩。
- 内置评论 您的博客就是您的家，而评论为朋友和关注者提供了与您的内容互动的空间。WordPress 的评论工具能满足您打造讨论区、管理讨论内容所需的一切。
- 搜索引擎优化 WordPress 开箱即对搜索引擎友好。如果需要更精细的 SEO 控制，还有大量 SEO 插件可以帮您搞定。

好，那就开始动手吧！

安装

```
composer create-project laravel/laravel wp-laravel
```

我们来试着访问一下应用

```
php artisan serve
```

我们需要配置一个数据库，本文就用 SQLite 吧。

```
touch database/database.sqlite
```

我们需要更新 .env 文件

```
DB_CONNECTION=sqlite
DB_DATABASE=database/database.sqlite
DB_FOREIGN_KEYS=true
```

[参考文档](https://laravel.com/docs/9.x/database#sqlite-configuration)

来执行初始的数据库迁移

```
php artisan migrate
```

现在，我们可以看到数据库的变化，如下所示

```
php artisan db:show
```

所有迁移均已应用。

接下来，我们想要一个类似 WordPress 文章（Post）的东西。如果您看过 WordPress 的数据表，就会知道其中有一张叫 wp\_posts 的表。我们来创建它

```
php artisan make:model Post
```

这会创建一个名为 Post 的模型。接着我们来创建迁移文件，以应用数据库变更。

```
php artisan make:migration "create posts table"
```

现在，打开这个迁移文件，把代码更新为如下内容：

```
$table->string('title');
$table->string('slug')->nullable();
$table->text('excerpt')->nullable();
$table->longText('content')->nullable();
$table->dateTime('published_at')->nullable();
$table->foreignIdFor(\App\Models\User::class)->constrained();
```

**继续阅读：** [如何设计订阅套餐的限制](https://sawirstudio.com/zh-hans/posts/how-do-you-design-your-subscription-plan-restrictions/)，或者看看[我们开发的免费 WordPress 插件](https://sawirstudio.com/zh-hans/projects/wordpress-plugins/)。
