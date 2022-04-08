# 目录

* [1. Introduction](正文/1.introduction/序言.md)
  * [1.1 常规约定](正文/1.introduction/1.1常规约定.md)
  * [1.2 准备工作](正文/1.introduction/1.2准备工作.md)
* [2. 新手入门](正文/2.新手入门/开始.md)
  * [2.1 项目配制和目录结构](正文/2.新手入门/2.1.项目配制和目录结构.md)
  * [2.2 HTTP服务基础](正文/2.新手入门/2.2.Http服务器基础.md)
  * [2.3 API接口和RESTful路由](正文/2.新手入门/2.3.API接口和RESTful路由.md)
* [3. 发送JSON响应](正文/3.发送JSON响应/章节介绍.md)
  * [3.1 格式固定的JSON](正文/3.发送JSON响应/3.1.格式固定的JSON.md)
  * [3.2 JSON编码](正文/3.发送JSON响应/3.2.JSON编码.md)
  * [3.3 结构体序列化](正文/3.发送JSON响应/3.3.结构体编码.md)
  * [3.4 格式化和封装响应](正文/3.发送JSON响应/3.4.格式化和封装响应.md)
  * [3.5 高级JSON定制化](正文/3.发送JSON响应/3.5.高级JSON定制化.md)
  * [3.6 发送错误响应](正文/3.发送JSON响应/3.6.发送错误响应.md)
* [4. 解析JSON请求](正文/4.解析JSON请求/章节介绍.md)
  * [4.1 JSON解码](正文/4.解析JSON请求/4.1.JSON解码.md)
  * [4.2 管理错误请求](正文/4.解析JSON请求/4.2.管理错误请求.md)
  * [4.3 限制HTTP请求body](正文/4.解析JSON请求/4.3.限制HTTP请求body.md)
  * [4.4 限制HTTP请求body](正文/4.解析JSON请求/4.4.自定义JSON解码.md)
  * [4.5 校验JSON输入](正文/4.解析JSON请求/4.5.校验JSON输入.md)
* [5. 数据库安装与配制](正文/5.数据库安装与配置/章节介绍.md)
  * [5.1 配置PostgreSQL](正文/5.数据库安装与配置/5.1.配置PostgreSQL数据库.md)
  * [5.2 连接PostgreSQL](正文/5.数据库安装与配置/5.2.连接PostgreSQL数据库.md)
  * [5.3 配置数据库连接池](正文/5.数据库安装与配置/5.3.配置数据库连接池.md)
* [6. SQL迁移](正文/6.SQL迁移/章节介绍.md)
  * [6.1 SQL迁移概述](正文/6.SQL迁移/6.1.SQL迁移概述.md)
  * [6.2 使用SQL迁移](正文/6.SQL迁移/6.2.使用SQL迁移.md)
* [7. 数据库CRUD操作](正文/7.CRUD操作/章节介绍.md)
  * [7.1 建立movie模型](正文/7.CRUD操作/7.1.建立movie模型.md)
  * [7.2 创建movie接口](正文/7.CRUD操作/7.2.创建新的movie接口.md)
  * [7.3 查询movie接口](正文/7.CRUD操作/7.3.查询movie接口.md)
  * [7.4 更新movie接口](正文/7.CRUD操作/7.4.更新movie接口.md)
  * [7.5 删除movie接口](正文/7.CRUD操作/7.5.删除movie接口.md)

##文章格式约定

在本书，代码块以Markdown格式显示，如下面的代码片段所示。如果代码块特别长，不相关的部分可以用省略号替换。为了便于理解，大多数代码块在顶部都有一个行文件名称，指示我们正在处理的文件。

> File: hello.go

```go
package main
... // Indicates that some existing code has been omitted.
func sayHello() { fmt.Println("Hello world!")
}
```

终端(命令行)指令以黑色背景显示，通常以美元符号开始。这些命令应该可以在任何基于unix的操作系统上执行，包括macOS和Linux。命令执行结果以白色显示在命令下面，如下所示:

```shell
 $ echo "Hello world!"
Hello world
```

如果你使用Windows，你应该用DOS等价的命令替换这些命令，或者通过正常的Windows GUI执行操作。

这本书的一些章节以附部分结束。这些部分包含的信息与我们的应用程序构建无关，但仍然很重要(有时只是有趣)。
## 背景知识

本书是接着Let's Go这本书来写的，将会沿用很多其中使用信息和代码模式。

如果你已经读过并喜欢《Let’s Go》，那么这本书应该很适合你，是你下一步学习的理想读物。如果你还没有读，那么我强烈建议你先从Let's Go开始——特别是如果你是一个新手。

您可以单独阅读这本书，但请注意它的难度更高——它没有详细解释基础知识，一些内容(如测试)根本没有出现，因为它们在前一本书中已经有大量介绍。但如果你已习惯使用Go，并且已经有了相当多的经验，那么这本书可能也很适合你。请随意直奔主题。

## GO 1.16

本书中的代码使用最新的Go 1.16版本验证过，如果您想要跟随构建应用程序的话，就应该安装该版本。

如果你已经按照来Go，你通过在终端中执行go version来检查版本。返回结果如下所示：

```shell
$ go version
go version go1.16 linux/amd64
```

如果你需要升级Go，请马上开始。根据你的操作系统升级手册在[这里](https://golang.org/doc/manage-install#uninstalling)。

## 其他软件

如果你想要完全遵循本书，你应该确保你的电脑上有其他一些软件。它们是:

* curl工具：处理来自终端的HTTP请求和响应。在MacOS和Linux机器上，它应该是预先安装的，或者在您的软件仓库中可用。否则，您可以从[这里](https://curl.haxx.se/download.html)下载最新版本。
* hey：用于进行一些基本负载测试的工具。如果你的电脑上有Go 1.16，你可以使用Go install命令安装hey:

  ```shell
   $ go install github.com/rakyll/hey@latest
  ```
* git版本控制系统。所有操作系统的安装说明都可以在[这里](https://git-scm.com/downloads)找到。
* 浏览器：包含一些开发者工具。可以使用firefox、Chromium、Chrome或者Microsof Edge也可以。
* 文本编辑器😄
## 序言

本书我们将从头到尾构建一个名为Greenlight的应用程序-用于检索和管理有关电影的信息的JSONAPI。

你可以思考下它核心功能有点类似于[开放电影数据库API](http://www.omdbapi.com/)。

最后我们的Greenlight API将支持以下API和操作：


| Method | URL                           | 动作                           |
| -------- | ------------------------------- | -------------------------------- |
| GET    | /v1/healthcheck               | 显示应用程序运行状况和版本信息 |
| GET    | /v1/movies                    | 显示所有电影的详情             |
| POST   | /v1/movies                    | 添加新的电影                   |
| GET    | /v1/movies/:id                | 根据id查询特定电影             |
| PATCH  | /v1/movies/:id                | 更新特定电影                   |
| DELETE | /v1/movies/:id                | 删除特定电影                   |
| POST   | /v1/users                     | 注册用户                       |
| PUT    | /v1/users/activated           | 激活用户                       |
| PUT    | /v1/users/password            | 更新用户密码                   |
| POST   | /v1/tokens/authentication     | 生成新的身份验证token          |
| POST   | /v1/tokens/password-reset     | 生成一个新的密码重置令牌       |
| GET    | /debug/vars/ 显示应用性能参数 |                                |

为了能够从客户角度来看下API是怎样的，在本书末尾GET /v1/movies/:id端点将返回如下响应内容：

```shell
curl -H "Authorization: Bearer RIDBIAE3AMMK57T6IAEBUGA7ZQ" localhost:4000/v1/movies/1
```

```json
{
"movie": {
"id": 1,
"title": "Moana", "year": 2016, "runtime": "107 mins", "genres": [
"animation",
"adventure" ],
"version": 1 }
}
```

在后端我们将使用PostgreSQL作为数据库来持久化所有的数据。并在本书的结尾，我们将在一台Linux服务器中部署完成后的API应用程序。
## 项目配制和代码结构

首先创建一个greenlight目录，作为这个项目的家目录。我将在$HOME/Projects/greenlight上创建我的项目目录，但如果你愿意，可以自由选择不同的位置。

```shell
 $ mkdir -p $HOME/Projects/greenlight
```

然后切换到这个目录，使用go mod init命令为项目启用模块。

运行此命令时，您需要指定模块路径，这实际上是项目的唯一标识符。在本书中，我将使用greenlight.alexedwards.net作为我的模块路径，如果你跟随操作的话，你应该把它替换成你自己的路径。

```shell
$ cd $HOME/Projects/greenlight
$ go mod init greenlight.alexedwards.net
go: creating new go.mod: module greenlight.alexedwards.net
```

此时你将在项目的根目录下看到go.mod文件被创建了。如果你打开该文件，内容类似如下所示：

> File: go.mod

```shell
module greenlight.alexedwards.net

go 1.16
```

在第一本Let’s Go书中详细讨论了Go module，但作为快速复习，我们在这里回顾一下要点。

* 当在你项目的根路径下面有一个合法的go.mod文件的话，说明你的项目是一个模块。
* 当你在项目中工作时，使用go get去下载依赖包，然后对应特定版本的依赖包将在go.mod中记录。因为版本是确定的，所以很容易在其他的机器或环境中执行代码。
* 当你在项目中运行或构建代码，Go将使用go.mod中记录的特定依赖。如果需要的依赖并不在当前机器中，Go将自动为你下载，以及任何递归依赖。
* go.mod文件还定义了模块路径（就是前面说的greenlight.alexedwards.net)。这将用作项目中包导入的根路径标识符。
* 让模块根路径保持唯一性，是一个好的编码习惯。Go社区的常见约定是使用你的URL作为模块根路径。

> 如果你不确定go module的工作原理，可以查看Go module [wiki](https://github.com/golang/go/wiki/Modules)，包含很多FAQ-注意可能有些内容已经过时。

## 生成目录结构

现在project目录已经创建好了并包含go.mod文件，可以继续执行以下命令创建项目结构：

```shell
$ mkdir -p bin cmd/api internal migrations remote 
$ touch Makefile
$ touch cmd/api/main.go
```

此时项目的目录结构应该如下所示的样子：

```shell
.
|____cmd
| |____api
|      |____main.go
|____migrations
|____go.mod
|____bin
|____Makefile
|____internal
|____remote
```

我们花点时间来讨论这些文件和文件夹，并解释它们在项目中所起的作用。

* bin目录包含应用编译后的二进制文件，可直接部署到生成环境中。
* cmd/api目录包含Greenlight项目的API代码文件。包括运行HTTP服务器、读写HTTP请求以及认证管理等代码。
* internal目录包含API使用的各种辅助包。包括与数据库交互、数据校验、发送邮件等。基本上任何非应用特定但可复用的代码都放在这里。在cmd/api中的代码会导入internal目录的代码（但从来没有反过来导入的）
* imgrations目录包含数据库SQL迁移文件。
* remote目录包含配制文件和设置脚本。
* go.mod文件申明项目依赖，版本和module路径。
* Makefile文件包含常见自动执行管理任务的方法，例如：代码审计、构建二进制文件和执行数据库迁移。

要重点指出的是internal目录在Go中有特殊的含义：在这个目录中的任何包，只能被该目录的父目录导入使用。在我们的项目中，这意味着internal中的所有包只能被greenlight项目中代码所引用。或者，从另一个角度来说，这意味着internal的任何包都不能通过项目外部的代码导入。这是有用的，因为它可以防止其他代码库导入和依赖我们internal目录中的(可能未版本化和不支持的)包——即使项目代码在GitHub中公开可用。

## hello world

在继续下一步之前，我们快速检查下所有设置是否正确。用文本编辑器打开cmd/api/main.go文件，添加以下代码：

```go
package main

import "fmt"

func main() { fmt.Println("Hello world!")
}
```

保存文件，并在终端中使用go run命令编译和执行cmd/api包中的代码。如果一切正常，您将看到以下输出:

```shell
$ go run ./cmd/api
Hello world!
```
## 简单的HTTP服务器

现在项目的框架结构已经就绪，让我们将注意力集中在启动和运行HTTP服务器上。

首先，我们将服务器配置为只有一个路由:/v1/healthcheck。这条路由将返回有关API的一些基本信息，包括其当前版本号和操作环境(开发、进度、生产等)。


| URL模式         | Handler            | 操作         |
| ----------------- | -------------------- | -------------- |
| /v1/healthcheck | healthcheckHandler | 显示应用信息 |

如果你跟着本书操作的，下面打开cmd/api/main.go文件用以下代码替换’hello world‘应用：

File:main.go

```go
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// 声明一个包含应用程序版本号的字符串. 在本书的后面，我们将在构建时自动生成它
//但是现在我们只将版本号存储为一个硬编码的全局常量.
const version = "1.0.0"

// 定义一个配置结构体来保存应用程序的所有配置设置.
//目前，唯一的配置是服务器监听的端口和应用程序的环境名称 (开发, 预发, 生成等等)
//将从命令行参数中读取这些配制信息
type config struct {
	port int
	env  string
}

// 定义一个application结构体来保存HTTP处理程序的依赖项
//目前，这只包含配置实例的一个副本和日志对象，但随着项目深入会增加更多内容
type application struct {
	config config
	logger *log.Logger
}

func main() {
	// 声明一个配置结构体实例
	var cfg config

	// 从命令行参数中将port和env读取到配制结构体实例当中。
	//默认端口使用4000以及环境信息使用开发环境development
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	// 初始化日志对象，将消息写入到标准输出。 以当前日期和时间为前缀
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// 声明应用程序结构的实例, 包含配制对象实例和日志对象。
	app := &application{
		config: cfg,
		logger: logger}

	// 申明一个新的servemux并添加/v1/healthcheck路由，将请求路由到即将实现的handler处理程序中去
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)

	// 使用一些合理的超时设置声明HTTP服务器
	//使用配制对象中的端口好以及创建的servemux作为handler
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// 启动HTTP服务
	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}

```

## 创建healthcheck处理程序

接下来我们需要创建用于响应HTTP请求的healthcheckHandler方法。现在，我们将保持这个处理程序中的逻辑非常简单，只让它返回一个包含三段信息的纯文本响应:

* 一个固定的“status: available”字符串
* API版本从硬编码version常量获取
* 操作环境名从命令行参数读取存在env中

继续创建cmd/api/healthcheck.go文件：

```shell
$ touch cmd/api/healthcheck.go
```

添加如下代码到文件中：

File：cmd/api/healthcheck.go

```go
package main

import (
	"fmt"
	"net/http"
)

//申明一个handler返回应用程序状态，操作环境和版本
func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment %s\n", app.config.env)
	fmt.Fprintf(w, "version: %s\n", version)
}

```

这里需要指出的重要一点是，healthcheckHandler是作为application结构体上的一个方法实现的.

这是一种有效且惯用的方法，使我们的处理程序可以使用依赖项，而不需要借助全局变量或闭包——当我们在main()中初始化healthcheckHandler时，任何依赖项都可以简单地作为一个字段包含在应用程序结构中。

我们可以看到上面的代码中已经使用了这个模式，其中通过调用app.config.env从应用程序结构中检索操作环境名称。

## 演示

代码我们来试试接口，请确保你的所有更改都已保存，然后再次使用go run命令来执行cmd/api包中的代码。您应该会看到一条日志消息，确认HTTP服务器正在运行，类似如下:

```shell
 $ go run ./cmd/api
2021/11/15 19:42:50 starting development server on :4000
```

当服务器运行时，继续尝试在web浏览器中访问localhost:4000/v1/healthcheck。你应该从healthcheckHandler得到如下响应:
![](../../images/2.2/env1.png)

或者，您可以使用curl从您的终端发出请求:

```shell
$ curl -i localhost:4000/v1/healthcheck
HTTP/1.1 200 OK
Date: Mon, 05 Apr 2021 17:46:14 GMT Content-Length: 58
Content-Type: text/plain; charset=utf-8
status: available environment: development version: 1.0.0
```

> 注意:上面命令中的-i标志表示curl显示HTTP响应头和响应体。

如果你想验证在命令行参数是否正常工作，通过指定port和env值然后在启动服务器。你应该可以看到日志信息，如下：

```shell
 $ go run ./cmd/api -port=3030 -env=production
2021/04/05 19:48:34 starting production server on :3030
```

## 附加说明

### API版本

支持真实业务和用户的api经常需要随着时间的推移而更改其功能和API——有时是以一种向后不兼容的方式。因此，为了避免客户端出现问题和困惑，最好实现某种形式的API版本控制。

通常有两种方法来实现：

1、通过给所有url加上API版本的前缀，比如/v1/healthcheck或/v2/healthcheck。

2、通过在请求和响应中使用自定义的Accept和Content-Type头来传递API版本，比如Accept: application/vnd.greenlight-v1。

从HTTP语义的角度来看，使用头来传递API版本是一种“更纯粹”的方法。但从用户体验的角度来看，使用URL前缀可能更好。开发者一眼就能看出API的版本，也意味着可以使用常规的web浏览器来访问API（如果放在请求头就更不方便）。


在整本书中，我们将通过在所有URL路径前加上/v1/来对API进行版本化，就像我们在本章中对/v1/healthcheck接口所做的那样。
## API接口和RESTful路由

在本书接下来的几节中，我们将逐步构建我们的API，使接口看起来像这样:


| Method | URL             | 动作                           |
| -------- | ----------------- | -------------------------------- |
| GET    | /v1/healthcheck | 显示应用程序运行状况和版本信息 |
| GET    | /v1/movies      | 显示所有电影的详情             |
| POST   | /v1/movies      | 添加新的电影                   |
| GET    | /v1/movies/:id  | 根据id查询特定电影             |
| PATCH  | /v1/movies/:id  | 更新特定电影                   |
| DELETE | /v1/movies/:id  | 删除特定电影                   |

如果您以前用REST风格构建过APIs，那么上面的表可能对您非常熟悉，不需要太多解释。但如果你是新手，那么有几件重要的事情需要指出。

首先，具有相同URL模式的请求将基于HTTP请求方法路由到不同的处理程序。为了安全性和语义的正确性，我们为处理程序执行的操作使用适当的HTTP方法是很重要的。

总之：


| Method | 用途                                                                                             |
| -------- | -------------------------------------------------------------------------------------------------- |
| GET    | 用于只检索信息而不更改应用程序或任何数据状态的操作。                                             |
| POST   | 用于修改状态的非幂等操作。在REST API上下文中，POST通常用于创建新资源的操作。                     |
| PUT    | 用于修改特定URL上资源的状态的幂等操作。在REST API上下文中，PUT通常用于替换或更新现有资源的操作。 |
| PATCH  | 用于对特定URL上的资源部分更新操作。不管是幂等的还是非幂等的都是可以的。                          |
| DELETE | 用于删除特定URL上的资源的操作。                                                                  |

另一件重要的事情需要指出的是，我们的API接口将使用简洁URLs，在URL路径中插入参数。例如，要获取特定的电影信息客户端将发送请求为：/v1/movies/1，而不是添加电影ID到查询字符串参数中例如：GET /v1/movies?id=1。

## 选择路由

当你在Go中使用这种接口构建API时，你将遇到的第一个问题是http.ServeMux——Go标准库中的路由器——在功能上非常有限。特别是，它不允许基于请求方法(GET、POST等)将请求路由到不同的处理程序，也不支持带有url中插入参数值。

尽管您可以解决这些限制，或者实现您自己的路由器，但通常使用许多可用的第三方路由器会更容易。

在本章中，我们将httprouter包集成到我们的应用程序中。最重要的是httprouter稳定、经过良好测试，并提供了我们需要的功能——另外，由于使用了radix树来进行URL匹配，速度非常快。如果您正在构建一个供公众使用的REST API，那么httprouter是一个可靠的选择。

如果你正跟随本书操作，请使用httpprouter的1.3.0版本，如下所示:

```shell
$ go get github.com/julienschmidt/httprouter@v1.3.0
go: downloading github.com/julienschmidt/httprouter v1.3.0 go get: added github.com/julienschmidt/httprouter v1.3.0
```

> 注意：如果你电脑上已经其他项目使用httprouter v1.3.0包，那么可以使用已经存在的版本将不会看到go: downloading...这些信息

为了演示httprouter是如何工作的，我们首先将两个接口添加到代码库中，用于创建新电影并查询特定电影的详情。在本章结束时，我们的API接口将看起来像这样:


| Method | URL             | 动作                           |
| -------- | ----------------- | -------------------------------- |
| GET    | /v1/healthcheck | 显示应用程序运行状况和版本信息 |
| GET    | /v1/movies      | 显示所有电影的详情             |
| POST   | /v1/movies      | 添加新的电影                   |

## 封装API路由

为了防止main()函数随着API的增长而变得混乱，我们将所有的路由规则封装在一个新的cmd/api/routes.go文件中。

如果您按照前面的步骤操作，请创建这个新文件并添加以下代码:

```shell
$ touch cmd/api/routes.go
```

```go
package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
        //初始化一个新的httprouter路由器实例
	router := httprouter.New()

        //使用HandlerFunc()函数为接口注册相关方法，URL模式和处理函数。
        //注意http.MethodGet和http.MethodPost是常量，等价于字符串"GET"和"POST"
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)

        //返回httprouter实例
        return router
}

```

> 提示:httprouter包还提供了一个router.Handler()方法，当你想注册一个常规的http.handler(而不是处理程序函数，就像我们在上面的代码)。

以这种方式封装路由规则有几个好处。第一个好处是它使main()函数保持简洁，并确保所有路由都在一个地方定义。另一个好处是，现在我们可以通过初始化application实例并在其上调用routes()方法，轻松地在任何测试代码中访问路由器。j

接下来需要更新main()函数来删除http.ServeMux声明，并使用app.routes()返回的httprouter实例作为服务器处理程序。像这样:

File: cmd/api/main.go

---

```go
package main
...
func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)") flag.Parse()
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	app := &application{ config: cfg,
		logger: logger }

	// 使用http.router实例做服务器的处理程序handler
	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", cfg.port), Handler: app.routes(),
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second, WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}
```

## 添加新的handler函数

既然路由规则已经设置好了，我们需要为新增的接口创建createMovieHandler和showMovieHandler方法。这里的showMovieHandler比较特殊，因为作为URL的一部分，我们希望从URL中提取电影ID参数，并在HTTP响应中使用它。

继续并创建一个新的cmd/api/movies.go文件保存这两个新的处理程序:

```shell
 $ touch cmd/api/movies.go
```

添加如下代码：

```go
package main
import ( "fmt"
"net/http" "strconv"
"github.com/julienschmidt/httprouter"
)
// 为"POST /v1/movies"接口添加createMovieHandler方法，这里简单的返回文本
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
fmt.Fprintln(w, "create a new movie") }

// 为 "GET /v1/movies/:id" 接口添加showMovieHandler方法. 这里我们从请求url中解析出电影id，并在响应中返回。
func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
// 当httprouter解析请求时，任何内置的URL参数都将存储在请求上下文中。 
//可以使用ParamsFromContext()函数检索包含这些参数名称和值的切片。
params := httprouter.ParamsFromContext(r.Context())

// 然后我们可以使用ByName()方法从切片中获取“id”参数的值。在我们的项目中，所有的电影都有一个唯一的正整数ID
//但是ByName()返回的值总是一个字符串。所以我们尝试将它转换为一个以10为基数的整数(bit size为64)。
//如果参数不能被转换或者小于1，我们知道ID无效，所以我们使用http.NotFound()返回404 Not Found响应。
id, err := strconv.ParseInt(params.ByName("id"), 10, 64) if err != nil || id < 1 {
http.NotFound(w, r)
return
}

// 否则，在响应中插入电影ID
fmt.Fprintf(w, "show the details of movie %d\n", id) }
```

有了这些，我们现在可以试一下了!重新启动API应用程序…

```shell
 $ go run ./cmd/api
2021/04/06 08:57:25 starting development server on :4000
```

然后，在服务器运行时，打开第二个终端窗口，并使用curl向不同的接口发出请求。如果一切设置正确，你会看到一些类似这样的响应:

```shell
$ curl localhost:4000/v1/healthcheck
status: available environment: development version: 1.0.0

$ curl -X POST localhost:4000/v1/movies
create a new movie

$ curl localhost:4000/v1/movies/123
show the details of movie 123
```

注意，在最后一个例子中，电影id参数123的值是如何从URL中成功检索到并包含在响应中的?

您可能还想尝试使用不支持的HTTP方法对特定的URL发出一些请求。例如，我们尝试向/v1/healthcheck发送一个POST请求:

```shell
$ curl -i -X POST localhost:4000/v1/healthcheck
HTTP/1.1 405 Method Not Allowed
Allow: GET, OPTIONS
Content-Type: text/plain; charset=utf-8 X-Content-Type-Options: nosniff
Date: Tue, 06 Apr 2021 06:59:04 GMT Content-Length: 19

Method Not Allowed
```

看起来很不错。httprouter包自动发送了一个405 Method Not Allowed响应，包括一个Allow报头，它列出了端点支持的HTTP方法。

同样地，你可以向一个特定的URL发出一个OPTIONS请求，httprouter会返回一个带有Allow报头的响应，其中详细说明了所支持的HTTP方法。像这样:

```shell
$ curl -i -X OPTIONS localhost:4000/v1/healthcheck
HTTP/1.1 200 OK
Allow: GET, OPTIONS
Date: Tue, 06 Apr 2021 07:01:29 GMT Content-Length: 0
```

最后，您可能想尝试在URL中使用负数或非数字id值向GET /v1/movies/:id接口发出请求。这应该会得到 404 Not Found响应，类似于:

```shell
$ curl -i localhost:4000/v1/movies/abc
HTTP/1.1 404 Not Found
Content-Type: text/plain; charset=utf-8 X-Content-Type-Options: nosniff
Date: Tue, 06 Apr 2021 07:02:01 GMT Content-Length: 19

404 page not found
```

## 为读取请求参数ID创建辅助函数

从URL(如/v1/movies/:id)中提取id参数的代码是我们在应用程序中需要反复使用的，因此我们将此逻辑抽象为一个可重用的辅助方法。

创建一个新的文件：cmd/api/helper.go

```shell
$ touch cmd/api/helpers.go
```

并向application结构体添加一个新的readdparam()方法，如下所示:

File: cmd/api/helpers.go

---

```go
package main
import ( "errors"
"net/http" "strconv"

"github.com/julienschmidt/httprouter"
)

// 从当前请求上下文检索“id”URL参数, 然后将其转换为一个整数并返回。如果操作不成功，则返回0和一个错误. 
func (app *application) readIDParam(r *http.Request) (int64, error) {
    params := httprouter.ParamsFromContext(r.Context())

    id, err := strconv.ParseInt(params.ByName("id"), 10, 64) 
    if err != nil || id < 1 {
        return 0, errors.New("invalid id parameter") 
    }

    return id, nil 
}
```

> readadidparam()方法不使用来自application结构体的任何依赖项，因此它可能只是一个常规函数，而不是application上的方法。但一般来说，我建议设置所有特定于应用程序的处理程序和帮助函数，以便它们是application上的方法。它有助于保持代码结构的一致性，并且在那些处理程序和帮助函数稍后更改并且它们确实需要访问依赖项时代码不会过时。

有了这个辅助方法，我们的showMovieHandler中的代码现在可以变得简单得多:

File: cmd/api/movies.go

---

```go
package main
    import ( "fmt"
    "net/http"
)

...

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) { 
    id, err := app.readIDParam(r)
    if err != nil {
        http.NotFound(w, r) 
    return
    }

    fmt.Fprintf(w, "show the details of movie %d\n", id) 
}
```

## 附加说明

#### 路径冲突

要意识到httprouter不允许可能匹配相同请求的冲突路由。因此，你不能注册一个像GET /foo/new这样的路由和另一个带有与之冲突的参数的路由，比如GET /foo/:id。

如果你使用标准REST结构来设计API的话，就不会有这种问题出现。因为不允许有冲突的路由，所以不需要担心路由优先级规则，这就减少了应用程序中出现错误和意外行为的风险。但是如果你需要支持这种冲突路由的话（例如，您可能需要完全复制现有API的接口以实现向后兼容性）那么我建议你看看pat, chi或者Gorilla mux。所有这些都是很好的第三方路由器包，它们允许冲突的路由。

## 定制httprouter行为

httprouter包提供了一些配置选项，您可以使用这些选项进一步定制应用程序的行为，包括启用尾部斜杠重定向和启用自动URL路径清理。更多可设置的信息可以参考[这里](https://pkg.go.dev/github.com/julienschmidt/httprouter?tab=doc#Router)。
## 开始

在本书的第一部分，我们将建一个项目框架，并为构建我们的Greenlight API打下基础。我们将:

* 为项目创建一个框架目录结构，并从宏观上解释我们的Go代码和其他文件将如何组织。
* 建立HTTP服务器来监听传入的HTTP请求。
* 介绍一个合适的管理配置模式（通过命令行flags）并使用依赖项注入使依赖项在我们的处理程序可用。
* 使用httprouter包实现API路由，达到标准RESTful结构。
## 格式固定的JSON

我们从更新healthcheckHandler函数发送一个格式良好的JSON响应开始，如下所示:

```go
{"status": "available", "environment": "development", "version": "1.0.0"}
```

在这个阶段，我想强调的是JSON只是文本。当然，它有一些控制字符来赋予文本结构和意义，但从根本上说，它只是文本。

这意味着你可以从Go处理程序中编写JSON响应，就像你编写任何其他文本响应一样:使用w.Write()，io.WriteString()或fmt.Fprint函数中一个。

实际上，我们需要做的唯一特殊的事情是在响应上设置Content-Type: application/json头，以便客户端知道它正在接收json，并相应地解析它。

下面我就开始这么做。

打开cmd/api/healthcheck.go文件并按如下方式更新healthcheckHandler:

File: cmd/api/healthcheck.go

---

```go

package main

import (
    "fmt"
    "net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
    //从字符串中创建一个固定格式的JSON。注意我们是如何从字符串常量的，这样可以JSON可以包含双引号字符
    //使用占位符
    js := `{"status": "available", "environment": %q, "version": %q}`
    js = fmt.Sprintf(js, app.config.env, version)

    //设置响应头“ContentType: application/json”。如果你忘记设置，Go将默认发送
    //"Content-Type: text/plain; charset=utf-8"响应头
    w.Header().Set("ContentType", "application/json")

   //将JSON写入HTTP响应body中
   w.Write([]byte(js))
}
```

完成这些更改后，重新启动API，打开第二个终端，并使用curl向GET /v1/healthcheck接口发出请求。你现在应该得到一个像这样的响应:

```shell
$ curl -i localhost:4000/v1/healthcheck
HTTP/1.1 200 OK
Content-Type: application/json 
Date: Tue, 06 Apr 2021 08:38:12 GMT 
Content-Length: 73

{"status": "available", "environment": "development", "version": "1.0.0"}
```

您可能还想尝试在浏览器中访问localhost:4000/v1/healthcheck。

如果运行较新版本的Firefox，会发现浏览器知道响应包含JSON(由于Content-Type报头)，它将使用内置的JSON查看器显示响应。像这样:

![](../../images/3.1/固定JSON.png)

如果你点击Raw Data选项卡，你应该看到原始的未格式化JSON响应:
![](../../images/3.1/rawData.png)

当然，使用固定格式的字符串(就像我们在本章中看到的那样)是生成JSON响应的一种非常简单和低保真的方法。但值得记住的是，这是一个有效的选择。它对总是返回相同静态JSON的API接口很有用，或者作为一种快速而简单的方法来生成动态响应。

## 附加内容

#### JSON字符集

在你的编程生涯当中可能遇到过其他的JSON APIs，会发送带有ContentType: application/json; charset=utf-8的响应header。通常不需要包含这样的字符集参数。[JSON RFC](https://tools.ietf.org/html/rfc8259#section-8.1)中描述：

> 在不属于封闭生态系统之间交换的JSON文本必须使用utf-8编码。

这里关键词是“必须”。因为我们的API将是一个面向公众的应用程序，这意味着我们的JSON响应必须始终是UTF-8编码的。这也意味着客户端假设它得到的响应总是UTF-8编码是安全的。因此，包含charset=utf-8参数就多余了。

RFC还描述了application/json媒体类型没有charset参数定义，这意味着从技术上讲，包含一个字符是不正确的。

基本上，在我们的应用程序中，charset参数没有任何作用，在Content-Type: application/json头文件中不包含字符集参数是安全的(也是正确的)。
## JSON编码

我们继续看一些更令人兴奋的内容，看看如何将Go对象(如map、结构体和切片)编码为JSON。

从上层角度看，Go的encoding/json包提供了两个选择来将内容编码为json。你可以调用json.marshal()函数，或者你可以声明并使用json.Encoder类型。

我们将在本章中解释这两种方法是如何工作的，但是为了在HTTP响应中发送JSON,使用JSON.marshal()通常是更好的选择。所以我们从这里开始。

JSON.marshal()的工作方式在概念上非常简单——您将一个Go对象作为参数传递给它，它将以字节数组形式返回该对象的JSON表示。函数签名看起来像这样:

```go
func Marshal(v interface{}) ([]byte, error)
```

> 注意:上述方法中的v参数的类型为interface{}(称为空接口)。这实际上意味着我们能够将任何Go类型作为v参数传递给Marshal()。

我们开始并更新healthcheckHandler方法，以便它使用JSON.marshal()直接从Go map生成JSON响应——而不是像之前那样使用固定格式的字符串。像这样:

File: cmd/api/healthcheck.go

---

```go
package main

import(
    "encoding/json"
    "net/http"
)

func (app *application)healthcheckHandler(w http.ResponseWriter, r *http.Request) {
    //创建一个map包含我们想发送的响应内容
    data := map[string]string{
        "status": "available",
        "environment": app.config.env,
        "version": version,
    }
    //将map对象传给json.Marshal()函数。序列号成byte数组包含编码后的JSON内容。如果有错误就打印到日志
    //向客户端发送一个错误消息
    js, err := json.Marshal(data)
    if err != nil {
         app.logger.Println(err)
         http.Error(w, "the server encountered a problem and could not process request", http.StatusInternalServerError)
         return
    }
    //向JSON中添加换行符。这主要是有助于在终端上看起来方便
    js = append(js, '\n')
  
    //此时编码已经完成了，因此需要HTTP添加响应header，通知客户端接收JSON格式内容
    w.Header().Set("ContentType", "application/json")
  
    //使用w.Write()发送包含JSON内容的字节数组
    w.Write(js)
}
```

如果你重新启动API并在浏览器中访问localhost:4000/v1/healthcheck，你现在应该会得到类似这样的响应:

![](../../images/3.1/固定JSON.png)

这看起来很不错——我们可以看到map对象已经自动编码为JSON对象，map中的键/值对在JSON对象中按字母顺序排序的。

## 创建一个writeJSON帮助方法

随着API的增长，我们将发送大量JSON响应，因此有必要将其中一些逻辑转移到可重用的writeJSON()帮助方法中。

除了创建和发送JSON外，我们还希望通过这个帮助函数，在以后的成功响应中可以包含任意响应头信息，比如在系统中创建新电影后的Location头信息。

如果你跟随文章编码的话，打开cmd/api/helpers.go文件并创建以下writeJSON()方法:

File: cmd/api/helper.go

---

```go
 import (
    "encoding/json"
    "errors"
    "net/http"
    "strconv"

    "github.com/julienschmidt/httprouter"
)

    //定义writeJSON()帮组函数来发送响应。方法以http.ResponseWrite作为响应写入地方，
    //HTTP状态码status，需要编码为JSON的数据data和一个响应头map对象
    func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers http.Header) error {
    //将data编码为JSON，如果有错误就返回
    js, err := json.Marshal(data)
    if err != nil {
        return err
    }
    //附加一个换行符以使它更容易在终端应用程序中查看。
    js = append(js, '\n')

    //此时，我们知道在发送响应之前不会再遇到任何错误，所以添加任何我们想要包含的响应头都是安全的。
    //循环迭代header map将对应的header添加到http.ResponseWriter响应头。注意map是nil也不会报错
    for key, value := range headers {
        w.Header().Set(k) = value
    }
    //添加"ContentType"响应头，然后写入状态码和JSON内容
    w.Header().Set("ContentType", "application/json")
    w.WriteHeader(status)
    w.Write(js)

    return nil
}
```

现在writeJSON()帮助函数已经就位，我们可以显著地简化healthcheckHandler中的代码：

File: cmd/api/healthcheck.go

---

```go
package main

import (
    "net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
    data := map[string]string{
        "status": "available",
        "environment": app.config.env,
        "version": version,
    }
  
    err := app.writeJSON(w, http.StatusOK, data, nil)
    if err != nil {
        app.logger.Println(err)
        http.Error(w, "The server encountered a problem and could not process you request", http.StatusInternalServerError)
    }
}
```

如果现在再次运行应用程序，一切都编译正确，对GET /v1/healthcheck接口的请求应该会得到与之前相同的HTTP响应。

## 附加内容

#### 不同的Go类型是如何编码的

在本章中，我们已经将一个map[string]string类型编码为JSON，其键/值对中的值为JSON字符串。但是Go也支持编码其他基本类型。

下表总结了不同的Go类型在编码过程中是如何映射到JSON数据类型的:


| Go 类型                                           | jSON 类型              |
| --------------------------------------------------- | ------------------------ |
| bool                                              | JSON boolean           |
| string                                            | JSON string            |
| int*, uint*, float*, rune                        | JSON number            |
| array, slice                                      | JSON array             |
| struct, map                                       | JSON object            |
| nil pointers, interface values, slices, maps, etc | JSON null              |
| chan, func, complex*                              | 不支持                 |
| time.Time                                         | RRC3339格式字符串      |
| []byte                                            | Base64编码的JSON字符串 |

最后两个是特殊情况，需要更多的解释:

* time.Time值(这实际上是一个结构体)将根据RFC3339格式来编码成JSON字符串类似“2020-11-08T06:27:59+01:00”，而不是JSON对象。
* []byte字节数组将编码为base64类型的JSON字符串，而不是一个JSON数组。因此[]byte{'h','e','l','l','o'}将编码为"aGVsbG8="。base64编码使用填充和标准字符集。

其他需要指出的是：

* 支持嵌套对象的编码。例如，如果你有一个结构体切片Go将编码成JSON对象数组。
* 不能对channel、函数和复数类型进行编码。如果你想这么做， 你会得到一个json.UnsupportedTypeError。
* 任何指针值都将被编码为所指向的值。同样，interface{}值也会编码为接口中包含的值。

## 使用json.Encoder

在本章的开头，我提到过也可以使用Go的json.Encoder类型来编码。它支持将对象编码为JSON对象并将JSON写入到一个输出流中，而且两个操作一步到位。

例如，你可以在处理程序中这样使用：

```go
func (app *application) exampleHandler(w http.ResponseWrite, r *http.Request){
    data := map[string]string{
        "hello": "world",
    }
    //设置“ContentType”响应头
    w.Header().Set("ContentType", "application/json")

    //使用json.NewEncoder()函数初始化json.Encoder实例，将内容接入http.ResponseWriter。
    //然后调用Encode()方法，将希望编码为JSON的data传入，如果data可以成功编码为JSON，它将
    //编码后写入到http.ResponseWriter。
    err := json.NewEncoder(w).Encode(data)
    if err != nil {
        app.logger.Println(err)
        http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
    }
}
```

这种方式是可行的，非常整洁和优雅，但如果你仔细考虑它，你可能会注意到一个小问题……

当我们调用JSON.NewEncoder(w).Encode(data)时，在一行代码中JSON被创建并写入http.ResponseWriter中，这意味着没有机会根据Encode()方法是否返回错误来设置HTTP响应头。

假设您想在一个成功的响应上设置一个Cache-Control报头，但如果JSON编码失败，则不设置Cache-Control报头而必须返回一个错误响应。使用json.Encoder模式就比较难实现。

你可以设置Cache-Control响应头，如果出现错误可以从响应头删除，但这都是很老套的。

另一个选择是将JSON写入bytes.Buffer缓存，而不是直接写入http.ResponseWriter。在设置Cache-Control响应头之前，检查任何错误。然后将JSON内容从bytes.Buffer缓存拷贝到http.ResponseWriter。但你真正那么处理的话，相比较而已使用json.Marshal()方法更简单。

#### json.Encoder和json.Marshal性能

谈到速度，您可能想知道json.Encoder和json.Marshal()之间是否有差异。简单来讲是肯定的……但是差别很小，在大多数情况下你不需要担心。

下面的基准测试使用本[gist](https://gist.github.com/alexedwards/8f11d0dc57fc119cbc791def783d8492)中的代码演示了这两种方法的性能(注意每个基准测试重复三次):

![image.png](./assets/image.png)

在这些结果中，我们可以看到json.marshal()要比json.Encoder稍微多一点的内存(B/op)，并使用更多的堆内存分配(allocs/op)。

这两种方法在平均运行时(ns/op)上没有明显可观察到的差异。也许在更大的基准测试样本或更大的数据集，差异可能会变得明显，但它可能也是微秒级的，而不是更大。
## 结构体编码

在这一节中，我们将回到之前做的showMovieHandler方法，并更新它以返回一个JSON响应，表示系统中的单个电影。类似于:

```json
{
    "id": 123,
    "title": "Casablanca", 
    "runtime": 102, 
    "genres": [
        "drama", 
        "romance", 
        "war"
    ],
    "version": 1 
}
```

我们不使用map序列化来创建这个JSON对象(就像我们在上一章中所做的那样)，这次我们将编码一个自定义的Movie结构。

首先，需要定义一个自定义的Movie结构体。我们将在一个新internal/data包中完成此操作，该包稍后将扩展用来封装项目中所有自定义数据类型以及与数据库交互的逻辑。

如果您按照文章步骤操作，请继续创建一个新的internal/data目录，其中包含一个movies.go文件:

```shell
$ mkdir internal/data
$ touch internal/data/movies.go
```

在这个新文件中，定义Movie结构，像这样:

File: internal/data/movies.go

---

```go
package main

import (
    "time"
)

type Movie struct {
    ID             int64      //唯一整数ID
    CreatedAt      time.Time  //创建电影到数据库的时间
    Title          string     //电影标题
    Year           int32      //电影发布年份
    Runtime        int32      //电影时长
    Genres         []string   //电影类型(爱情片、喜剧片等)
    Version        int32      //版本号从1开始，每更新一次递增
}
```

> 这里需要指出的是，Movie结构体中的所有字段都是可导出的(即以大写字母开头)，这对于Go的encoding/json包可见是必要的。在将结构体编码为JSON时，不会包含任何未导出的字段。

现在结构体已经定义完成，让我们更新showMovieHandler处理程序来初始化一个Movie结构体实例，然后使用writeJSON()帮组函数将其作为JSON响应发送给客户端。

实现很简单：

File： cmd/api/movies.go

---

```go
package main

import (
    "fmt"
    "net/http"
    "time"

    "greenlight.alexedwards.net/internal/data"
)

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
    id, err := app.readIDParam(r)
    if err != nil {
        http.NotFound(w, r)
        return
    }

    //创建一个Move结构体实例，包含从请求URL中解析的ID虚构的数据。注意这里故意没有设置Year字段
    movie := date.Movie{
        ID: id,
        CreateAt: time.now(),
        Title: "Casablanca",
        Runtime: 102,
        Genres: []string{"drama", "romance", "war"},
        Version: 1,
    }

    //将结构体序列化为JSON并以HTTP响应发送给客户端
    err = app.writeJSON(w, http.StatusOK, movie, nil)
    if err != nil {
         app.logger.Println(err)
         http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
    }
}
```

ok，下面试试！

重启API，然后在浏览器中访问localhost:4000/v1/movies/123。你应该会看到一个类似这样的JSON响应:

![image.png](../../images/3.1/3.3结构体.png)

在这个回应中，有几件有趣的事情需要指出:

* Movie结构体被编码成一个JSON对象，字段名和值作为键/值对。
* 默认情况下，JSON对象中的键等于结构中的字段名(ID、CreatedAt、Title等等)。我们稍后将讨论如何自定义键值。
* 如果结构字段没有显式的值集，那么字段零值将被json编码。可以在上面的响应中看到一个例子——我们没有在Go代码中为Year字段设置值，但它仍然以0值出现在JSON输出中。

#### 更改JSON对象中的键

在Go中序列化结构体的一个好处是，您可以通过使用struct标签注释字段来定制JSON。

struct标签最常见的用途可能是更改JSON对象中出现的键名。当你的结构体字段名不适合面向公众展示，或者你想在JSON输出中使用另一种大小写样式时，这是很有用的。

为了说明如何实现，对Movies结构体字段打标签，使用蛇形格式：

File: internal/data/movies.go

---

```go
//使用标记对Movie结构进行注释，以控制json编码的key显示方式。
type Movie struct {
	ID       int64     `json:"id"`
	CreateAt time.Time `json:"created_at"`
	Title    string    `json:"title"`
	Year     int32     `json:"year"`
	Runtime  int32     `json:"runtime"`
	Genres   []string  `json:"genres"`
	Version  int32     `json:"version"`
}
```

如果你重启服务器并再次访问localhost:4000/v1/movies/123，应该会看到一个类似于这样的带有蛇形键的响应:

![image.png](../../images/3.1/movie.png)

## 在JSON对象中隐藏结构体字段

在定义结构体时候，通过使用omitempty可以控制对应字段在JSON中的可见性。当您不希望JSON输出中出现特定的结构体字段时，可以使用-(连字符)指令。这对包含和用户不相关的内部系统信息的字段或不想公开的敏感信息(如密码的散列)非常有用。

相反，当且仅当struct字段值为空时，omitempty指令会在JSON输出中隐藏字段，其中empty被定义为:

* 等于false，0或“”
* 空数组，切片或map
* nil指针或接口值为nil

为了演示如何使用这些指令，我们对Movie结构进行更多的改造。CreatedAt字段与我们的最终用户无关，所以我们使用-指令在输出中将其隐藏。我们还将使用omitempty指令在输出中隐藏Year、Runtime和types字段，当且仅当它们为空时情况下生效。

继续并像下面这样更新struct标签:

File：interface/data/movies.go

---

```go
package data

....

type Movie struct {
	ID       int64     `json:"id"`
	CreateAt time.Time `json:"-"`       //使用-指令
	Title    string    `json:"title"`
	Year     int32     `json:"year,omitempty"`            //添加omitempty
	Runtime  int32     `json:"runtime,omitempty"`         //添加omitempty
	Genres   []string  `json:"genres,omitempty"`          //添加omitempty
	Version  int32     `json:"version"`
}
```

> 如果你想使用omitempty而不改变键名，那么你可以在struct标签中保留它为空-如:json:"，omitempty"。注意，逗号是必要的。

现在，当你重新启动应用程序并刷新你的web浏览器时，你应该会看到如下响应:

![image.png](../../images/3.1/omitempty.png)

我们可以在这里看到，CreatedAt结构字段不再出现在JSON中，而且Year字段(值为0)也没有出现，这要感谢omitempty指令。其他字段使用了omitempty不受影响(例如Runtime和Genres)。

> 注意:您还可以通过简单地将结构体字段设置为不可导出来防止它出现在JSON序列化中。但使用json:“—“标记通常是一个更好的选择:明确告知阅读代码的人，你不希望该字段包含在json。
>
> 旧版本的go vet如果你试图在未导出的字段上使用struct标记会引发错误，但现在在go 1.16中已经修复了这个问题。

## 附加内容

#### 结构体标签string指令

最后一个不太常用的struct标记指令是string。可以使用这个标签明确表示字段值序列化成JSON字符串类型。例如，如果我们希望Runtime字段的值表示为一个JSON字符串 (而不是数字)我们可以像这样使用string指令:

```go
type Movie struct {
	ID       int64     `json:"id"`
	CreateAt time.Time `json:"-"`       //使用-指令
	Title    string    `json:"title"`
	Year     int32     `json:"year,omitempty"`   
	Runtime  Runtime   `json:"runtime,omitempty,string"` 
	Genres   []string  `json:"genres,omitempty"`       
	Version  int32     `json:"version"`
}
```

JSON序列化结果如下所示:

```shell
{
"id": 123,
"title": "Casablanca",
"runtime": "102",   //这是字符串
"genres": [
    "drama", 
    "romance", 
    "war"
    ],
"version": 1 
}
```

注意string指令只对int*， uint*， float*或bool类型的字段有效。对于任何其他类型的结构体字段没有作用。
## 格式化和封装响应

到目前为止，我们一直在使用Firefox向API发送请求，由于内置的JSON查看器提供了“输出格式化”，这使得JSON响应易于阅读。

但是如果您尝试使用curl发起请求，将看到实际的JSON响应数据都在一行中，没有空格。

```shell
$ curl localhost:4000/v1/healthcheck
{"environment":"development","status":"available","version":"1.0.0"}

$ curl localhost:4000/v1/movies/123
{"id":123,"title":"Casablanca","runtime":102,"genres":["drama","romance","war"],"version":1}
```

通过使用json.MarshalIndent()函数来编码响应数据，而不是使用常规的json.Marshal()函数，可以使这些内容更容易在终端中查看。自动将空格符添加到JSON输出中，每个元素放在单独的行，并在每个行前面加上可选的前缀和缩进字符。

我们更新writeJSON()助手来使用下面的代码:

File: cmd/api/helpers.go

---

```go
package main

...

func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers http.header) error {
    //使用json.MarshalIndent()函数，可以在JSON输出中添加空格。这里使用""作为前缀，"\t"格式化元素
    js, err := json.MarshalIndent(data, "", "\t")
    if err != nil {
        return err
    }
    js = append(js, '\n')

    for key, value := range headers {
        w.Header()[key] = value
    }
  
    w.Header.Set("ContentType", "application/json")
    w.WriteHeader(status)
    w.Write(js)

    return nil
}
```

如果你重新启动API服务器并尝试从你的终端再次发出相同的请求，将会收到一些类似于以下内容，包含格式化空格的JSON响应:

```shell
$ curl -i localhost:4000/v1/healthcheck
{
    "environment": "development",
    "status": "available",
    "version": "1.0.0",
}

$ curl localhost:4000/v1/movies/123
{
    "id": 123,
    "title": "casablanca",
    "runtime": 102,
    "genres": [
        "drama",
        "romance",
        "war"
    ],
    "version": 1
}
```

## 性能比价

虽然从可读性和用户体验的角度来看，使用json.MarshalIndent()是更适合，但不幸的是，它不是免费的。除了响应的总字节数稍微大一点之外，Go为添加空格所做的额外工作对性能有显著的影响。

下面的基准测试使用[gist](https://gist.github.com/alexedwards/809db1839a062c4f92a65f40359e18b7)中的代码演示json.Marshal()和json.MarshalIndent()的相对性能。
![img_1.png](assets/img_1.png)

在这些基准测试中，我们可以看到json.MarshalIndent()的运行时间比json.Marshal()长65%，使用的内存比json.Marshal()多30%，还多两次堆分配。这些数据会根据编码内容而变化，但根据我的经验，它们反映了性能影响。

但在大多数应用程序中这点性能不会造成担心。实际上，我们谈论的只是千分之一毫秒的时间——为提高响应的可读性可能值得这样做。但是，如果你的API在资源非常受限的环境中运行，或者需要管理大级别的流量，那么这是值得注意的，您可能更喜欢使用json.Marshal()。

> 注意:JSON.marshalindent()的工作原理是像调用JSON.marshal()一样，然后通过独立的JSON.indent()函数为JSON添加空格。还有一个反向函数JSON.compact()，您可以使用它从JSON中删除空格。

#### 封装响应

接下来，我们更新响应，使JSON数据始终包含在父JSON对象中。类似于:

```json
{
    "movie":
    {
        "id": 123,
        "title": "Casablanca",
        "runtime": 102,
        "genres":
        [
            "drama",
            "romance",
            "war"
        ],
        "version": 1
    }
}
```

注意到电影对象是如何嵌套在键“movie”下面的，而不是本身作为顶层JSON对象。

严格来说，这样封装响应数据是没有必要的，是否选择这样做在一定程度上取决于风格和品味。但也有一些切实的好处:

1、在JSON的顶层包含一个键名(如“movie”)有助于使响应更具文档化。对于任何从上下文中看到响应的人来说，理解响应数据与什么相关更容易一些。

2、它降低了客户端出错的风险，因为这样做会避免把一个响应当作其他内容来处理。为了获得数据，客户端必须通过“movie”键显式地引用它。

3、如果我们总是封装API返回的数据，那么我们就减轻了在旧浏览器中返回JSON数组作为响应可能出现的安全漏洞。

我们可以使用一些技术来封装API响应，但将保持事情简单，通过创建一个自定义的envelope map与底层类型map[string]interface{}来实现。

接下来将说明怎么封装响应：

从更新cmd/api/helper.go开始，如下所示:

File: cmd/api/helpers.go

---

```go
package main

...

//定义envelop类型
type envelope map[string]interface{}

//将data参数类型改为envelope
func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.header) error {
    js, err := json.MarshalIndent(data, "", "\t")
    if err != nil {
        return err
    }

    js = append(js, '\n')

    for key, value := range headers {
        w.Header()[key] = value
    }

    w.Header.Set("ContentType", "application/json")
    w.WriteHeader(status)
    w.write(js)

    return nil
}
```

然后，我们需要更新showMovieHandler，以创建包含电影数据的envelope实例，并将其传递给writeJSON()函数，而不是直接传递电影结构体实例，如下所示：

File: cmd/api/movies.go

---

```go
package main

...

func (app *application) showMovieHandler(w http.ResponseWriter, r http.Request) {
    id, err := app.readIDParam(r)
    if err != nil {
        http.NotFound(w, r)
        return
    }

    movie := data.Movie{
        ID:        id,
        CreateAt:  time.Now(),
        Title:     "Casablanca",
        Runtime:   102,
        Genres:    []string{"drama", "romance", "war"},
        Version:   1,
    }

    //创建envelope{"movie": movie}实例，然后传给writeJSON函数，而不是直接传movie结构体实例
    err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
    if err != nil {
        app.logger.Println(err)
        http.Error(w, "the server encountered a problem and could not process your request", http.StatusInternalServerError)
    }
}
```

还需要更新healthcheckHandler中的代码，以便它也将envelope类型传递给writeJSON()函数:

File: cmd/api/healthcheck.go

---

```go
package main

...

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
    //申明envelope类型包含响应数据。注意，我们构造它的方式意味着环境和版本数据现在将嵌套在JSON响应的system_info键下。
    env := envelope{
        "status": "available",
        "system_info": map[string]string{
            "environment": app.config.env,
            "version": version,
        },
    }
    err := app.writeJSON(w, http.StatusOK, env, nil)
    if err != nil {
        app.logger.Println(err)
        http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
    }
}
```

好了，我们来试试这些修改。重新启动服务器，然后使用curl向API接口发出请求。您现在应该得到如下格式的响应。

```shell
$ curl localhost:4000/v1/movies/123
{
    "movie":
    {
        "id": 123,
        "title": "Casablanca",
        "runtime": 102,
        "genres":
        [
            "drama",
            "romance",
            "war"
        ],
        "version": 1
    }
}

$ curl localhost:4000/v1/healthcheck
{
    "status": "available",
    "system_info":
    {
        "environment": "development",
        "version": "1.0.0"
    }
}
```

## 附加内容

#### HTTP响应结构

要强调下构造JSON响应没有唯一正确或错误方法。有一些流行的格式，如JSON:API和jsend，您可能希望遵循或使用它们来获得灵感，但这当然不是必需的，大多数API不遵循这些格式。

但是无论你做什么，考虑格式化和在不同的API接口之间保持清晰和一致的响应结构都是很有价值的——特别是当它们可供公众使用。
## 高级JSON定制化

通过使用结构体标签、添加空白和封装响应数据，我们已经能够为JSON响应添加大量定制信息。但是，当这些内容还不够时，您需要更自由地定制JSON时，会发生什么呢?

要回答这个问题，我们首先需要谈谈Go如何处理JSON序列化的一些理论。要理解的关键是:

Go是在什么时候将特殊类型序列化为JSON，它首先查看对应的类型是否实现了MarshalJSON()方法。如果实现了，GO将调用这个方法来决定JSON编码格式。

这么将有点模糊，我们更精确点。严格地说，当Go将特定类型编码为JSON时，它会查看该类型是否满足json.Marshaler接口，该接口如下所示:

```go
type Marshaler interface {
    MarshalJSON() ([]byte, error)
}
```

如果类型确实满足接口，那么Go将调用它的MarshalJSON()方法，并使用它返回的[]byte切片作为JSON编码的值。

如果该类型没有MarshalJSON()方法，那么Go将返回尝试根据自己的内部规则将其编码为JSON。

因此，如果我们想定制某些类型的编码方式，只需要在其上实现MarshalJSON()方法，该方法以[]byte类型返回自定义的JSON内容。

> 提示：如果您查看time.Time类型源代码，就可以看到这一点。time.Time实际上是一个结构体，但是它有一个MarshalJSON()方法，输出RFC3339格式JSON对象。当time.Time值被序列化为JSON对象时，就会调用MarshalJSON()方法。

## 定制Runtime字段JSON序列化

为了说明这一点，让我们看一下应用程序中的一个具体示例。

当我们的Movie结构被编码为JSON时，Runtime字段(它是一个int32类型)编码为JSON数字。现在我们来更改它，将其编码为"<<runtime>runtime> mins“的字符串。像这样:

```json
{
    "id": 123,
    "title": "Casablanca",
    "runtime": "102 mins",
    "genres":
    [
        "drama",
        "romance",
        "war"
    ],
    "version": 1
}
```

有几种方法可以实现这一点，但一种简单的方法是为Runtime字段创建一个自定义类型，并在这个类型上实现MarshalJSON()方法。

为了防止internal/data/movie.go文件不会太乱，我们创建一个新的文件来处理runtime类型序列化逻辑：

```shell
 $ touch internal/data/runtime.go
```

然后继续添加以下代码:

```go
package data

import (
    "fmt"
    "strconv"
)

//申明Runtime类型，其底层是int32类型（和movie中的字段一样）
type Runtime int32

//实现MarshalJSON()方法，这样就满足来json.Marshaler接口。
func (r Runtime) MarshalJSON() ([]byte, error) {
    //生成一个字符串包含电影时长
    jsonValue := fmt.Sprintf("%d mins", r)

    //使用strconv.Quote()函数封装双引号。为了在JSON中以字符串对象输出，需要用双引号。
    quotedJSONValue := strconv.Quote(jsonValue)
    //将字符串转为[]byte返回
    return []byte(quotedJSONValue)
}
```

这里我想强调两点:

* 如果您的MarshalJSON()方法像我们的方法一样返回一个JSON字符串值，那么您必须在返回字符串之前用双引号包装它。否则它将不会被解释为JSON字符串，你将收到类似于这样的运行时错误:
  ```shell
  json: error calling MarshalJSON for type data.Runtime: invalid character 'm' after top-level value
  ```
* 我们故意为MarshalJSON()方法使用值接收器，而不是指针接收器func (r *Runtime) MarshalJSON()。这给了我们更多的灵活性，因为这意味着定制JSON编码将对Runtime值对象和指针对象都有效。正如Effective Go提到的:

> 如果你不确定指针和值接收器之间的区别，那么这篇[博客](https://medium.com/globant/go-method-receiver-pointer-vs-value-ffc5ab7acdb)提供了一个很好的总结。

好的，现在有了自定义Runtime类型，打开internal/data/movies.go文件并更新Movie结构:

File: internal/data/movies.go

---

```go
package data

import (
    "time"
)

type Movie struct {
	ID       int64     `json:"id"`
	CreateAt time.Time `json:"-"`
	Title    string    `json:"title"`
	Year     int32     `json:"year,omitempty"`
        //使用Runtime类型取代int32，注意omitempty还是能生效的
	Runtime  Runtime   `json:"runtime,omitempty,string"`
	Genres   []string  `json:"genres,omitempty"`
	Version  int32     `json:"version"`
}
```

重启服务然后对GET /v1/movies/:id接口发起请求。你应该看到一个包含自定义运行时值的响应，格式为"xx mins"，类似如下:

```shell
$ curl localhost:4000/v1/movies/123
{
    "movie":
    {
        "id": 123,
        "title": "Casablanca",
        "runtime": "102 mins",
        "genres":
        [
            "drama",
            "romance",
            "war"
        ],
        "version": 1
    }
}
```

总之，这是生成定制JSON的一种很好的方法。我们的代码简洁明了，并且我们有一个自定义的Runtime类型，可以随时随地使用它。

但也有不利的一面。在将代码与其他包集成时，使用自定义类型有时会很尴尬，您可能需要执行类型转换，将自定义类型转换为其他包理解和可接受的值。
## 发送错误响应消息

此时，我们的API正在为成功的请求发送格式化良好的JSON响应，但如果客户端发出一个错误请求——或者应用程序中出现了问题——我们仍然会从http.Error()和http.NotFound()函数向他们发送纯文本的错误消息。

在本章中，我们将通过创建一些额外的帮助函数来管理错误并向客户端发送适当的JSON响应来解决这个问题。

如果你跟随本书的编码操作，下面创建一个新的文件cmd/api/errors.go：

```shell
$ touch cmd/api/errors.go
```

在文件中添加一些帮助函数：

```go
package main

import (
    "fmt"
    "net/http"
)

//logError方法是一个通用帮助函数用于记录错误信息。后面会升级该函数使用结构化的日志实例来记录HTTP请求额外信息。
func (app *application) logError(r *http.Request, err error) {
    app.logger.Println(err)
}

//errorResponse方法：向客户端发送给定错误码响应。使用interface{}类型表示要返回给客户端的消息，而不是string类型
//这样做更灵活
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}){
    env := envelope{"error": message}
    //使用writeJSON将错误信息写入http响应。如果函数执行错误，需要记录错误日志，并返回500错误给客户端
    err := app.writeJSON(w, status, env, nil)
    if err != nil {
        app.logError(r, err)
        w.WriteHeader(500)
    }
}

//serverErrorResponse()方法：当应用程序在运行时碰到非预期的问题使用。会记录详细错误日志，并使用errorResponse()函数
//发送500内部服务器错误状态码
func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
    app.logError(r, err)

    message := "the server encountered a problem and could not process your request"
    app.errorResponse(w, r, http.StatusInternalServerError, message)
}

//notFoundResponse()方法：发送404 not found状态码和JSON响应给客户端
func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
    message := "the requested resource could not be found"
    app.errorResponse(w, r, http.StatusNotFound, message)
}

//methodNotAllowedResponse()方法：发送405请求方法不允许状态码和JSON格式响应给客户端
func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
    message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
    app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}

```

现在这些帮助函数都就绪了，我们更新API处理程序，使用这些新的帮助函数而不是http.Error()和http.NotFound()函数。像这样:

File: cmd/api/healthcheck.go

---

```go
package main

...

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}

	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
                //使用新的serverErrorResponse()帮助函数
		app.serverErrorResponse(w, r, err)
	}
}
```

File: cmd/api/movies.go

---

```go
func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
                //使用notFoundResponse帮助函数
		app.notFoundResponse(w, r)
		return
	}
	movie := data.Movie{
		ID:       id,
		CreateAt: time.Now(),
		Title:    "Casablanca",
		Runtime:  102,
		Genres:   []string{"drama", "romance", "war"},
		Version:  1,
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
                //使用serverErrorResponse帮助函数
		app.serverErrorResponse(w, r, err)
	}
}

```

## 路由错误

我们自己的API处理程序发送的任何错误消息现在都将是格式良好的JSON响应。效果非常好!

但是当httprouter找不到匹配的路由时，它自动发送错误消息怎么处理呢？默认情况下，这些响应仍然是我们在本书前面看到的纯文本(非json)响应。

幸运的是，httprouter允许我们在初始化路由器时设置自己的自定义错误处理程序。这些自定义处理程序必须满足http.handler接口，这对我们来说是个好消息，因为这意味着我们可以轻松地重用我们刚刚创建的notFoundResponse()和methodnotallowdresponse()帮助函数。

打开cmd/api/routes.go文件并配置httprouter实例，如下所示:

```go
func (app *application) routes() http.Handler {
	router := httprouter.New()
   
        //http.HandlerFunc函数将notFoundResponse()帮助函数转为http.Handler，将自定义错误处理程序替换默认的
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
        //同理，methodNotAllowedResponse也同样转换并赋值
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)
	return app.recoverPanic(app.rateLimit(router))
}
```

我们测试一下这些更改。重新启动应用程序，然后尝试对不存在的接口或使用不支持的HTTP方法发起一些请求。现在你应该会得到一些类似的JSON错误响应:

```shell
$ curl -i localhost:4000/foo
HTTP/1.1 404 Not Found 
Content-Type: application/json 
Date: Tue, 06 Apr 2021 15:13:42 GMT 
Content-Length: 58

{
    "error": "the requested resource could not be found"
}

$ curl -i localhost:4000/v1/movies/abc
HTTP/1.1 404 Not Found 
Content-Type: application/json 
Date: Tue, 06 Apr 2021 15:14:01 GMT 
Content-Length: 58

{
    "error": "the requested resource could not be found"
}

$ curl -i -X PUT localhost:4000/v1/healthcheck
HTTP/1.1 405 Method Not Allowed 
Allow: GET, OPTIONS
Content-Type: application/json 
Date: Tue, 06 Apr 2021 15:14:21 GMT 
Content-Length: 66

{
    "error": "the PUT method is not supported for this resource"
}
```

在最后一个例子中，请注意httprouter仍然自动为我们设置正确的Allow头，即使它现在使用我们的自定义错误处理程序的响应。

## 附加内容

#### 系统生成的错误响应

当我们谈到错误的话题时，我想提一下在某些情况下Go的http.Server仍然可以自动生成和发送文本类型的HTTP响应。这些场景包括:

* HTTP请求指定不支持的HTTP协议版本
* HTTP请求缺少或包含无效的host请求头，或多个host请求头。
* HTTP请求包含无效的请求头name和value
* HTTP请求包含不支持的Transfer-Encoding头信息
* HTTP请求头大小超出服务器的最大设置MaxHeaderBytes
* 客户端想HTTPS服务器发送HTTP请求

例如，如果我们尝试发送一个带有无效Host头值的请求，我们将得到这样的响应:

```shell
$ curl -i -H "Host: こんにちは" http://localhost:4000/v1/healthcheck 
HTTP/1.1 400 Bad Request: malformed Host header
Content-Type: text/plain; charset=utf-8
Connection: close

400 Bad Request: malformed Host header
```

不幸的是，这些响应被硬编码到Go标准库中，我们无法对它们进行定制以使用JSON返回。

但是，虽然这是需要注意的事情，但我们不一定需要担心。在生产环境中，行为良好的、非恶意的客户端触发这些响应的可能性相对较小。
## 发送JSON响应

在本书的这一部分，我们将更新API处理程序，使它们返回JSON响应，而不是纯文本。

JSON (JavaScript Object Notation的首字母缩写)是一种人类可读的文本格式，可以用来表示结构化数据。例如，在我们的项目中，一个电影可以用以下JSON表示:

```json
{
    "id": 123,
    "title": "Casablanca",
    "runtime": 102,
    "genres":
    [
        "drama",
        "romance",
        "war"
    ],
    "version": 1
}
```

为了让您快速了解JSON语法……

* 括号{}定义一个JSON对象，由逗号分隔键/值对。
* 键必须是字符串，值可以是字符串、数字、布尔值、其他JSON对象或这些类型组成的数组（用方括号）
* 字符串必须使用双引号而不是单引号，布尔值不是true就是false。
* 空格符在json里面没有任何意义，对于我们来说，用下面的JSON来表示同一个电影是完全正确的:

```json
  {"id":123,"title":"Casablanca","runtime":102,"genres":["drama","romance","war"],"version":1}
```

如果你以前没使用JSON，这个[新手指南](https://towardsdatascience.com/an-introduction-to-json-c9acb464f43e)提供了完整的介绍，建议你在继续之前可以先阅读里面的内容。

本书的这部分，你将学习到以下内容：

* 如何从REST API发送JSON响应(包括错误响应)。
* 如何使用encoding/JSON包将原生Go对象编码为JSON。
* 不同的方法将Go对象编码为JSON-首先使用结构体标签，然后利用json.Marshaler接口。
* 如何创建发送JSON响应的可重用帮助函数，这将确保所有API响应具有合理和一致的结构。
## JSON解码（反序列化）

和JSON编码一样，有两种方式可以用于将JSON解码为Go对象：使用json.Decoder类型和json.Unmarshal()函数。

这两种方法各有优缺点，但为了从HTTP请求体解码JSON，使用JSON.Decoder通常是最好的选择。它比json.Unmarshal()更高效，需要更少的代码，并提供了一些有用的设置，您可以使用这些设置来调整其行为。

用代码而不是文字来说明json.Decoder是如何工作会更简单，所以让我们直接进入代码更新createMovieHandler处理函数:

File: cmd/api/movies.go

---

```go
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"

    "greenlight.alexedwards.net/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
    //申明一个匿名结构体来接收HTTP请求体中对JSON内容，注意结构体中字段和类型与之前创建movie结构体只包含部分字段
    //该结构体定义类型用于接收http请求，并解码为Go对象。
    var input struct {
        Title     string   `json:"title"`
        Year      int32    `json:"year"`
        Runtime   int32    `json:"runtime"`
        Genres    []string `json:"genres"`
    }

    //初始化json.Decoder实例，从http请求body中读取请求内容，然后使用Decode()方法将内容解析为input结构体。
    //注意Decoder函数接收对是指针类型，如果解析错误就调用errorResponse()帮助函数返回400错误给客户端。
    err := json.NewDecoder(r.Body).Decode(&input)
    if err != nil {
        app.errorResponse(w, r, http.StatusBadRequest, err.Error())
        return
    }

    //将解析后对input结构体写入HTTP响应，返回给客户端
    fmt.Fprintf(w, "%+v\n", input)
  
   ...
}
```

关于这段代码，有一些重要的地方需要指出:

* 当调用Decoder()必须传入一个非nil指针作为解析对存储目标。如果传入对不是指针，运行时将返回json.InvalidUnmarshalError错误。
* 如果传入对是结构体，想上面代码中对那样，结构体字段必须首字母大写。和编码一样，它们需要被导出，这样它们对encoding/json包是可见的。
* 当将JSON对象解码为结构体时，JSON中的键/值对将基于结构体标签名映射到结构体字段。如果没有匹配的结构体标签，Go将试图将对应对JSON值编码到匹配对结构体字段中，不区分大小写的匹配)。任何不能成功映射到结构体字段的JSON键/值对都将被忽略。
* 在r.Body被读取之后，没有必要关闭它。这将由Go的http.Server自动处理。

好吧，我们来试试。

启动应用程序，然后打开第二个终端窗口，向POST /v1/ moviesendpointwithavalidjsonrequestbody发出请求，其中包含一些movie数据。你应该会看到类似这样的响应:

```shell
#创建一BODY变量，包含要发送对JSON内容
$ BODY='{"title":"Moana","year":2016,"runtime":107, "genres":["animation","adventure"]}'

#使用-d命令行参数将BODY内容发送给服务端
$ curl -i -d "$BODY" localhost:4000/v1/movies
HTTP/1.1 200 OK
Date: Tue, 06 Apr 2021 17:13:46 GMT Content-Length: 65
Content-Type: text/plain; charset=utf-8

{Title:Moana Year:2016 Runtime:107 Genres:[animation adventure]}
```

太棒了!似乎很有效。从响应中数据可以看出，我们在请求体中提供的值已经被解码到input结构体的对应字段中。

## 零值

让我们快速看一下如果我们在JSON请求体中忽略特定的键/值对会发生什么。例如，在JSON中创建一个没有year字段的请求，如下所示:

```shell
$ BODY='{"title":"Moana","runtime":107, "genres":["animation","adventure"]}' 
$ curl -d "$BODY" localhost:4000/v1/movies
{Title:Moana Year:0 Runtime:107 Genres:[animation
```

正如您可能已经猜到的，当我们这样做时，输入结构中的Year字段将保留其零值(碰巧是0，因为Year字段是一个int32类型)。

这就引出了一个有趣的问题:如何区分客户端不提供键/值对和提供键/值对但故意将其设置为零的情况？例如:

```shell
$ BODY='{"title":"Moana","year":0,"runtime":107, "genres":["animation","adventure"]}' 
$ curl -d "$BODY" localhost:4000/v1/movies
{Title:Moana Year:0 Runtime:107 Genres:[animation adventure]}
```

尽管HTTP请求不同，但最终结果是相同的，并且如何区分这两种场景并不是很明显。我们将在本书的后面再回到这个话题，但现在，只需要了解这特殊情况为就够了。

## 附加内容

#### 支持对目标类型

值得一提的是，某些JSON类型只能成功解码为某些Go类型。例如，如果你有JSON字符串“foo”，它可以被解码成一个Go字符串，但试图将其解码成一个Go int或bool将导致运行时错误(我们将在下一章中演示)。

下表提供了不同JSON类型支持的目标解码GO类型:


| JSON 类型    | 支持的Go类型              |
| -------------- | --------------------------- |
| JSON boolean | bool                      |
| JSON string  | string                    |
| JSON number  | int*, uint*, float*, rune |
| JSON array   | array, slice              |
| JSON object  | struct, map               |

#### 使用json.Unmarshal函数

正如我们在本章开始时提到的，也可以使用json.Unmarshal()函数来解码HTTP请求体。

例如，你可以像这样在处理器中使用它:

```go
func (app *application) exampleHandler(w http.ResponseWriter, r *http.Request) {
    var input struct {
        Foo string `json:"foo"`
    }

    //使用io.ReadAll()读取整个HTTP请求body内容到[]byte切片中
    body, err := io.ReadAll(r.Body)
    if err != nil {
        app.serverErrorResponse(w, r, err)
        return
    }
  
    //使用json.Unmarshal()函数将切片中的JSON解码到input结构体。再次说明使用的参数是指针。
    err = json.Unmarshal(body, &input)
    if err != nil {
        app.errorResponse(w, r, http.StatusBadRequest, err.Error())
    }

    fmt.Fprintf(w, "%+v\n", input)

    ...
}
```

使用这种方法很简单。但没有我们之前提到的json.Decoder方法中的有点。

不仅代码稍微更冗长，而且效率也更低。如果我们对这个特定用例的相对性能进行基准测试，我们可以看到使用json. unmarshal()比json.Decoder多损耗80%的内存(B/op)。以及稍微慢一点(ns/op)。

![image.png](./assets/image.png)
## 管理错误请求

现在，当createMovieHandler接收到带有适当数据的有效JSON请求体时，它可以正常工作。但在这一点上你可能会想:

* 如果客户端发送的不是JSON，比如XML或一些随机字节，该怎么办?
* 如果JSON格式不正确或包含错误会发生什么?
* 如果JSON类型与我们试图解码的类型不匹配怎么办?
* 如果请求甚至不包含body呢?

下面我们就来看看！

```shell
# 发送一些XML作为请求体
$ curl -d '<?xml version="1.0" encoding="UTF-8"?><note><to>Alice</to></note>' localhost:4000/v1/movies
{
    "error": "invalid character '\u003c' looking for beginning of value"
}

# 发送一些格式错误的JSON(注意末尾的逗号)
$ curl -d '{"title": "Moana", }' localhost:4000/v1/movies
{
    "error": "invalid character '}' looking for beginning of object key string"
}

# 送一个JSON数组而不是一个对象
$ curl -d '["foo", "bar"]' localhost:4000/v1/movies
{
    "error": "json: cannot unmarshal array into Go value of type struct { Title string    \"json:\\\"title\\\"\"; Year int32 \"json:\\\"year\\\"\"; Runtime int32 \"json:\\ \"runtime\\\"\"; Genres []string \"json:\\\"genres\\\"\" }"
}

# 发送一个数值'title'值(而不是字符串)
$ curl -d '{"title": 123}' localhost:4000/v1/movies
{
    "error": "json: cannot unmarshal number into Go struct field .title of type string"
}

# 发送一个空的请求体
$ curl -X POST localhost:4000/v1/movies
{
    "error": "EOF"
}
```

在所有这些情况，我们可以看到createMovieHandler都在按正确请求处理。当它接收到一个无法解码到我们的input结构中的无效请求时，不会做进一步的处理，客户端会收到一个JSON响应，其中包含Decode()方法返回的错误消息。

对于一个私有API，它不会被公共的成员使用，那么这种返回没什么影响，你不需要做任何其他的事情。

但是对于面向公众的API，错误消息本身并不理想。有些过于详细，显示来有关底层API实现的信息。其他错误信息没有足够的描述性(如“EOF”)，还有一些难以理解错误。使用的格式或语言也不一致。

为了改进这一点，我们将解释如何对Decode()返回的错误进行分类，并将它们替换为更清晰、易于操作的错误消息，以帮助客户端准确地调试JSON的错误。

## 对解码错误进行分类

此时，在我们的应用程序构建中，Decode()方法可能会返回以下五种类型的错误:


| 错误类型                   | 原因                                                                                         |
| ---------------------------- | ---------------------------------------------------------------------------------------------- |
| json.SyntaxError           | 正在解码的JSON存在语法问题。                                                                 |
| io.ErrUnexpectedEOF        | 正在解码的JSON存在语法问题。                                                                 |
| json.InvalidUnmarshalError | 解码目标无效(通常是因为它不是指针)。这实际上是我们应用程序代码的问题，而不是JSON本身的问题。 |
| io.EOF                     | 正在解码的JSON是空的。                                                                       |

对这些潜在错误进行分类(我们可以使用Go的errors.is()和errors.as()函数来完成)将使createMovieHandler中的代码更长更复杂。这个逻辑我们也需要在整个项目的其他处理程序中重复处理。

所以，我们在cmd/api/helpers.go中创建一个新的readJSON()帮助函数。将从请求体中正常解码JSON，然后对错误进行分类，并在必要时用我们自己的定制消息替换它们。

如果你跟随本书操作，请继续并将以下代码添加到cmd/api/helper.go文件:

File: cmd/api/helpers.go

---

```go
package main

import (
    "encoding/json"
    "errrors"
    "fmt"
    "io"
    "net/http"
    "strconv"

    "github.com/julienschmidt/httprouter"
)

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {
    //将请求body解析到dst中
    err := json.NewDecoder(r.Body).Decode(dst)
    if err != nil {
        //如果在解析的过程中发生错误，开始分类...
        var syntaxError *json.SyntaxError
        var unmarshalTypeError * json.UnmarshalTypeError
        var invalidUnmarshalError *json.InvalidUnmarshalError

        switch {
        //使用errors.As()函数检查错误类型是否为*json.SyntaxError。如果是，返回错误
        case errors.As(err, &syntaxError):
            return fmt.Errorf("body contain badly-formed JSON (at charcter %d)", syntaxErr.Offset)
        //某些情况下，因为语法错误Decode()函数可能返回io.ErrUnexpectedEOF错误。
        case errors.Is(err, io.ErrUnexpectedEOF):
            return errors.New("body contain badly-formed JSON")
        //同样捕获*json.UnmarshalTypeError错误，这些错误是因为JSON值和接收对象不匹配。如果错误对应到特定到字段，
        //我们可以指出哪个字段错误方便客户端debug
        case errors.As(err, &unmarshalTypeError):
            if unmarshalTypeError.Field != "" {
                return fmt.Errorf(body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
            }
            return fmt.Errorf(body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)
         //如果解码到内容是空对象，会返回io.EOF。
         case errors.Is(err, io.EOF):
             return errors.New("body must not be empty")
          //如果decode()函数传入一个非空的指针，将返回json.InvalidUnmarshalError。
          case errors.As(err, &invalidUnmarshalError):
              panic(err)
          default:
              return err
        }
    }
    return nil
}
```

有了这个新的帮助函数，让我们回到cmd/api/movies.go文件并更新createMovieHandler来使用它。像这样:

File: cmd/api/movies.go

---

```go
package main

import (
    "fmt"
    "net/http"
    "time"

    "greenlight.alexedwards.net/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
    var input struct {
        Title      string   `json:"title"`
        Year       int32    `json:"year"`
        Runtime    int32    `json:"runtime"`
        Genres     []string `json:"genres"`
    }
    //使用新的readJSON()帮助函数解析请求体到input结构体实例。如果返回错误，我们将错误信息携带400错误码返回客户端
    err := app.readJSON(w, r, &input)
    if err != nil {
        app.errorResponse(w, r, http.StatusBadRequest, err.Error())
        return
    }
    fmt.Fprintf(w, "%+v\n", input)
}

...

```

重新启动API，然后通过重复我们在本章开始时发出的相同的错误请求。现在，你应该看到我们新的、自定义的错误消息，类似如下:

```shell
# 发送一些XML作为请求体
$ curl -d '<?xml version="1.0" encoding="UTF-8"?><note><to>Alice</to></note>' localhost:4000/v1/movies
{
    "error": "body contains badly-formed JSON (at character 1)"
}

# 发送一些格式错误的JSON(注意末尾的逗号)
$ curl -d '{"title": "Moana", }' localhost:4000/v1/movies
{
    "error": "body contains badly-formed JSON (at character 20)"
}

# 送一个JSON数组而不是一个对象
$ curl -d '["foo", "bar"]' localhost:4000/v1/movies
{
    "error": "body contains incorrect JSON type (at character 1)"
}

# 发送一个数值'title'值(而不是字符串)
$ curl -d '{"title": 123}' localhost:4000/v1/movies
{
    "error": "body contains incorrect JSON type for \"title\""
}

# 发送一个空的请求体
$ curl -X POST localhost:4000/v1/movies
{
    "error": "body must not be empty"
}
```

它们看起来真不错。错误消息现在在格式上更简单、更清晰和一致，而且它们不会暴露任何有关底层程序的不必要信息。

如果您愿意，可以随意使用它，并尝试发送不同的请求主体，以查看处理程序如何响应。

## 创建错误请求处理函数

在上面的createMovieHandler代码中，我们使用通用的app.errorResponse()帮助函数向客户端发送一个400 Bad Request响应和错误消息。

我们用一个专业的app.badRequestResponse()帮助函数代替它:

File: cmd/api/errors.go

---

```go
package main

...

func (app *application) bandRequestResponse(w http.ResponseWriter, r *http.Request) {
    app.errorResponse(w, r, http.StatusBadReqeust, err.Error())
}
```

File: cmd/api/movies.go

---

```go
package main

...

func(app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
    var input struct {
        Title      string   `json:"title"`
        Year       int32    `json:"year"`
        Runtime    int32    `json:"runtime"`
        Genres     []string `json:"genres"`
    }

    err := app.readJSON(w, r, &input)
    if err != nil {
        //使用新的badRequestResponse()帮助函数
        app.badRequestResponse(w, r, err)
        return 
    }
    fmt.Fprintf(w, "%+v\n", input)
}
```

这是一个很小的修改，但很有用。随着我们的应用程序变得越来越复杂，使用像这样的专业帮助函数来管理不同类型的错误将有助于确保我们的错误响应在所有接口上保持一致。

## 附加内容

#### panic对比返回错误

在之前错误分类中，碰到json.InvalidUnmarshalError错误时，我们在readJSON函数中直接panic。你可能觉得在go中错误一般都向上返回给调用者。但是在一些特殊情况下panic是没问题都。所以在必要都时候不需要坚持不panic的教条。当应用程序发生错误时，这里区分两类不同的错误是有用的。

第一类错误是在正常的操作过程中能预期到会发生的。预期错误的一些例子包括：数据库查询超时、网络资源不可用或用户输入错误。这些错误并不一定意味着您的程序本身有问题——事实上，它们通常是由程序控制之外的事情引起的。任何时候，返回这类错误并优雅地处理它们都是一种良好的实践。

另一类错误是非预期的错误。这些错误在正常操作中不应该发生，如果它们发生了，可能是开发人员错误或代码库中的逻辑错误造成的。这些错误确实是异常的，在这种情况下使用panic更被广泛接受。事实上，Go标准库经常会这么处理，当发生逻辑错误或试图以非预期的方式使用语言特性时——比如试图访问切片中的越界索引，或试图关闭已经关闭的通道。

但即便如此，我还是建议在大多数情况下尝试返回并优雅地处理意外错误。特殊情况是当返回错误时，会给代码库的其余部分增加不可接受的错误处理工作量。

回到readJSON()帮助函数，如果运行时发生json.InvalidUnmarshalError，因为开发人员给Decode()函数传递错误的值导致。这绝对是一个在正常操作下不应该看到的意外错误，而且应该在部署之前在开发和测试中发现。

如果我们返回错误而不是panic，我们需要在每个api处理程序中引入额外的代码来管理这些错误。对于不太可能在生产环境中看到的错误，这似乎不是一个很好的权衡。

[这篇文章](https://gobyexample.com/panic)很好的总结了panic。

> panic通常是指发生了非可预期错误。大多数情况下，正常操作中不应该发生的错误，我们使用它来快速失败，因为没有准备好优雅地处理这类错误。
## 限制HTPP请求body

在前一章中为处理无效JSON和其他错误请求所做的更改是朝着正确方向迈出的一大步。但我们仍然可以做一些事情来使JSON处理更加健壮。

其中对一件事就是处理未知字段。例如，你向下面对代码一样，向createMovieHandler接口发送一个包含未知字段rating的请求：

```shell
$ curl -i -d '{"title": "Moana", "rating":"PG"}' localhost:4000/v1/movies
HTTP/1.1 200 OK
Date: Tue, 06 Apr 2021 18:51:50 GMT Content-Length: 41
Content-Type: text/plain; charset=utf-8

{Title:Moana Year:0 Runtime:0 Genres:[]}
```

请注意这个请求是如何正常地工作的——没有通知客户端，应用程序碰到无法识别的rating字段。在某些场景中，默默忽略未知字段可能正是您想要的行为，但在这里，最好能够提醒客户端注意这个问题。

幸运的是json.Decoder提供了DisallowUnknownField()配制可以在碰到未知字段时报错。

另一个问题是json.Decoder支持JSON数据流。当我们在请求体上调用Decode()时，它实际上只从请求体读取第一个JSON值并解码。如果我们对Decode()进行第二次调用，它将读取并解码第二个JSON值，以此类推。

但是因为我们只在readJSON()中调用了一次Decode()，并且只调用了一次，所以在请求体中第一个JSON值之后的任何内容都会被忽略。这意味着您可以发送一个包含多个JSON值的请求体，或者在第一个JSON值之后发送垃圾内容，而我们的API处理程序不会引发错误。例如:

```shell
# Body包含多个JSON值
$ curl -i -d '{"title": "Moana"}{"title": "Top Gun"}' localhost:4000/v1/movies
HTTP/1.1 200 OK
Date: Tue, 06 Apr 2021 18:53:57 GMT Content-Length: 41
Content-Type: text/plain; charset=utf-8

{Title:Moana Year:0 Runtime:0 Genres:[]}

# Body在第一个JSON值之后包含垃圾内容
$ curl -i -d '{"title": "Moana"} :~()' localhost:4000/v1/movies
HTTP/1.1 200 OK
Date: Tue, 06 Apr 2021 18:54:15 GMT Content-Length: 41
Content-Type: text/plain; charset=utf-8

{Title:Moana Year:0 Runtime:0 Genres:[]}
```

同样，这种行为可能非常有用，但在解析api请求时它并不适合。我们希望对createMovieHandler处理程序的请求在请求体中只包含一个JSON对象，并只包含关于创建的电影的信息。

为了确保请求体中没有额外的JSON值(或任何其他内容)，我们需要在readJSON()函数中第二次调用Decode()，并检查它是否返回一个io.EOF(文件结束)错误。

最后，目前对我们接受的请求体的最大大小没有上限。这意味着对于任何企图对我们的API执行Dos攻击的恶意客户端来说，我们的createMovieHandler将是一个很好的目标。我们可以通过使用http.MaxBytesReader()函数来限制请求体的最大大小来解决这个问题。

让我们更新readJSON()函数来修复以下三个问题:

File: cmd/api/helper.go

---

```go

package main

...

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	maxBytes := 1_048_576
        //使用http.MaxBytesReader函数限制只读取1MB以内的请求内容
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))
	dec := json.NewDecoder(r.Body)
        //启用DisallowUnknownFields()函数，这意味着在解析JSON的过程中，如果有不可匹配的字段就返回错误。
	dec.DisallowUnknownFields()

	err := dec.Decode(dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError
		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formed JSON")
		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contain incorrect JSON type (at character %d)", unmarshalTypeError.Offset)
		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")
                //如果JSON中包含不能解析的字段，将返回“json: unknown field”错误，这里从错误内容中提取出字段名称
                //返回给客户端
                case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return fmt.Errorf("body contain unknown key %s", fieldName)
                //如果请求内容超过1MB将返回"http: request body too large"错误，返回给客户端
		case err.Error() == "http: request body too large":
			return fmt.Errorf("body must not be large than %d bytes", maxBytes)

		case errors.As(err, &invalidUnmarshalError):
			panic(err)
		default:
			return err
		}
	}
	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only contain a single JSON value")
	}
	return nil
}

```

修改完之后，我们再次尝试本章前面的请求:

```shell
$ curl -d '{"title": "Moana", "rating":"PG"}' localhost:4000/v1/movies
{
    "error": "body contains unknown key \"rating\""
}

$ curl -d '{"title": "Moana"}{"title": "Top Gun"}' localhost:4000/v1/movies
{
    "error": "body must only contain a single JSON value"
}

$ curl -d '{"title": "Moana"} :~()' localhost:4000/v1/movies
{
    "error": "body must only contain a single JSON value"
}
```

根据上面的结果发现请求处理被终止，客户端收到一个明确的错误消息，解释了错误原因。

下面可以发送一个包含大数据量的请求，看看响应结果。

为了演示，作者创建了一个1.5MB的JSON文件，您可以通过运行以下命令将其下载到/tmp目录:

```shell
$ wget -O /tmp/largefile.json https://www.alexedwards.net/static/largefile.json
```

如果你将这个文件内容作为请求体发送到POST /v1/movies接口，http.MaxByteReader()会检查并报错，你将收到如下响应信息：

```shell
$ curl -d @/tmp/largefile.json localhost:4000/v1/movies
{
    "error": "body must not be larger than 1048576 bytes"
}
```

完成以上修改，我们总算完成了readJSON()函数的业务处理。

我必须承认readJSON()里面的代码并不是最漂亮的…我们引入了许多错误处理和逻辑，用于对Decode()函数返回错误的处理。但现在已经写好了。您不需要再它，而且您可以轻松地将它复制和粘贴到其他项目中。
## 自定义JSON解码

在本书的前面，我们在API中添加了一些自定义的JSON编码处理，以便在JSON响应中以“(runtime) mins”的格式显示电影时长信息。

在本章，我们将从另一方面出发如何将“(runtime) mins”字段解析到int32类型的Go结构体字段中，也称为反序列化过处理。

如果你现在尝试用"(runtime) mins"这种格式发送一个请求，会得到一个400 Bad request响应(因为不可能将JSON字符串解码为int32类型)。像这样:

```shell
$ curl -d '{"title": "Moana", "runtime": "107 mins"}' localhost:4000/v1/movies
{
    "error": "body contains incorrect JSON type for \"runtime\""
}
```

为了使其正常工作，我们需要拦截JSON解码过程，并手动将“<runtime> mins”JSON字符串转换为int32。

因此我们该如何处理呢？

## json.Unmarshaler接口

这里的关键是了解Go的json.Unmarshaler接口，它看起来像这样:

```go
type Unmarshaler interface { 
    UnmarshalJSON([]byte) error
}
```

当Go解码JSON时，会检查目标字段的类型是否实现了json.Unmarshaler接口。如果实现了该接口，Go将调用UnmarshalJSON()函数来决定如何将JSON解码到目标类型。这刚好与json.Marshaler接口相反，我们之前使用它来定制JSON编码行为。

我们看看如何实现这个接口。

我们需要做的第一件事是更新createMovieHandler，以便输入结构体使用我们自定义的Runtime类型，而不是常规的int32。你应该记得，我们的runtime类型仍然有基础类型int32，但通过自定义类型，我们可以自由地让它上实现UnmarshalJSON()方法。

下面更新接口处理程序如下：

File：cmd/api/movies.go

---

```go
package main

...

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string       `json:"title"`
		Year    int32        `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Genres  []string     `json:"genres"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}

```

然后我们转到internal/data/runtime.go文件，为Runtime类型添加UnmarshalJSON方法。在这个方法里面需要解析“(runtime) mins”格式的JSON字符串，然后将runtime数值解析为int32，并将其赋值给Runtime本身。

这实际上有点复杂，有一些重要的细节，最好直接进入代码，并在后面用注释进行解释。

File：internal/data/runtime.go

---

```go
package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//如果UnmarshalJSON函数不能将JSON字符串解析为Runtime类型，就返回该错误
var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

type Runtime int32

//Runtime类型实现UnmarshalJSON()方法，这样就实现来json.Unmarshaler接口。注意，因为UnmarshalJSON()会改变状态
//因此需要使用指针接收者。否则就要返回一个实例的拷贝。
func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {
        //希望接收到的值是一个JSON字符串"（runtime） mins"，第一步需要去掉双引号。
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}
        //分割字符串以分离包含数字的部分
	parts := strings.Split(unquotedJSONValue, " ")
	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}
        //将数字字符串转为int32类型
	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidRuntimeFormat
	}
        //将int32强制转为Runtime类型，并赋值给接收者，注意是指针类型的接收者。
	*r = Runtime(i)
	return nil
}
```

完成以上代码后重启服务，然后使用"(runtime) min"格式的值作为JSON字段。可以看到请求会被正常处理。数值会从字符串中提取出来赋值给input结构体中的Runtime类型字段，如下所示：

```shell
 $ curl -d '{"title": "Moana", "runtime": "107 mins"}' localhost:4000/v1/movies
{Title:Moana Year:0 Runtime:107 Genres:[]}
```

然而，如果你使用JSON数字或任何其他格式发请求，你应该会得到一个ErrInvalidRuntimeFormat错误消息的响应，类似如下:

```shell
$ curl -d '{"title": "Moana", "runtime": 107}' localhost:4000/v1/movies
{
    "error": "invalid runtime format"
}

$ curl -d '{"title": "Moana", "runtime": "107 minutes"}' localhost:4000/v1/movies
{
    "error": "invalid runtime format"
}
```
## 校验JSON输入

在许多情况下，您需要对来自客户端的数据执行额外的验证检查，以确保它在处理之前满足特定的业务规则。在本章中，我们将通过更新createMovieHandler来演示如何在JSON API的上下文中做到这一点:

* 客户端提供的movie标题不为空，长度不超过500字节。
* movie的year字段不能是空的，而且是在1888年到今年之间。
* runtime字段不能空，而且是一个正数。
* genres字段包含1-5个不同的电影类型。

如果其中任何一个检查失败，我们希望向客户端发送一个422 Unprocessable Entity响应，以及清楚地描述验证失败的错误消息。

## 创建validator包

为了在整个项目中帮助我们进行验证，将创建一个internal/validator包，其中包含一些简单的可重用的帮助类型和函数。如果您正在跟随本的操作，请在您的机器上创建以下目录和文件:

```shell
$ mkdir internal/validator
$ touch internal/validator/validator.go
```

然后在文件internal/validator/validator.go中添加如下代码：

```go
package validator

import "regexp"

var (
        //申明一个正则表达式用于检查email的格式，如果你有兴趣该正则表达式来自于“https://html.spec.whatwg.org/#valid-e-mail-address”网站。
	EmailRx = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\. [a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

//定义一个新的Validator类型，其中包含验证错误的map。
type Validator struct {
	Errors map[string]string
}

//New是一个构造函数，用于创建Validator实例
func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
}

//valid 返回true如果map中没有错误
func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

//AddError 向map中添加错误(map中不存在对应key的错误)
func (v *Validator) AddError(key, message string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

//Check 向map中添加错误消息，如果校验失败即ok为false
func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

//In 如果list切片中存在value字符串返回true
func In(value string, list ...string) bool {
	for i := range list {
		if value == list[i] {
			return true
		}
	}
	return false
}

//Match 如果字符串满足正则表达式就返回true
func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

//如果切片中的字符串都不同返回true
func Unique(values []string) bool {
	uniqueValues := make(map[string]bool)
	for _, value := range values {
		uniqueValues[value] = true
	}
	return len(values) == len(uniqueValues)
}

```

总结：

在上面的代码中定义了Validator类型，包含一个错误信息map字段。Validator提供了Check()方法，根据校验结果向map中添加错误信息，而Valid()方法返回map是否包含错误信息。还添加了In(), Matches()和Unique()方法来帮助我们执行特定字段的检查。

从概念上讲，这个Validator类型是非常基本的，但这并不是一件坏事。正如我们将在本书中看到的，它在开发中功能强大，为我们提供了很多灵活性的字段检查。

## 执行字段检查

下面我们把validator类型使用起来。

我们需要做的第一件事是更新cmd/api/errors.go文件，添加一个新的failedValidationResponse()帮组函数，它将写入一个422 Unprocessable Entity错误码，并将来自新Validator类型的错误内容映射为JSON响应体。

File: cmd/api/errors.go

---

```go
package main

...

//注意errors参数是一个map类型，和validator类型包含map一致
func (app *application) failedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	app.errorResponse(w, r, http.StatusUnprocessableEntity, errors)
}
```

完成之后，返回到createMovieHandler并更新它，以对input结构体各个字段进行必要的检查。像这样:

File：cmd/api/movies.go

---

```go
package main

import (
	"fmt"
	"net/http"
	"time"

	"greenlight.alexedwards.net/internal/data"
	"greenlight.alexedwards.net/internal/validator"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string       `json:"title"`
		Year    int32        `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Genres  []string     `json:"genres"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	movie := &data.Movie{
		Title:   input.Title,
		Year:    input.Year,
		Runtime: input.Runtime,
		Genres:  input.Genres,
	}
        //初始化一个新的Validator实例
	v := validator.New()
  
        //使用Check()方法执行字段校验。如果校验失败就会向map中添加错误信息。例如下面第一行检查title不能为空，然后再检查长度不能超过500字节等等。
        v.Check(movie.Title != "", "title", "must be provide")
	v.Check(len(movie.Title) <= 500, "title", "must not be more than 500 bytes long")

	v.Check(movie.Year != 0, "year", "must be provided")
	v.Check(movie.Year >= 1888, "year", "must be greater than 1888")
	v.Check(movie.Year <= int32(time.Now().Year()), "year", "must not be in the future")

	v.Check(movie.Runtime != 0, "runtime", "must be provided")
	v.Check(movie.Runtime > 0, "runtime", "must be a positive integer")

	v.Check(movie != nil, "genres", "must be provided")
	v.Check(len(movie.Genres) >= 1, "genres", "must contain at least 1 genres")
	v.Check(len(movie.Genres) <= 5, "genres", "must not contain more than 5 genres")
        //使用Unique()方法，检查input.Genres每个字段是否唯一。
	v.Check(validator.Unique(movie.Genres), "genres", "must not contain duplicate values")

        //使用Valid()方法确认检查是否通过。如果有错误就使用failedValidationResponse()帮助函数返回错误信息给客户端。
        if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	fmt.Fprintf(w, "%+v\n", input)
}
```

做完这个之后，我们就可以试一下了。重新启动服务，然后向post /v1/movie接口发送请求，其中包含一些不合法的字段信息，类似下面:

```shell
$ BODY='{"title":"","year":1000,"runtime":"-123 mins","genres":["sci-fi","sci-fi"]}' $ curl -i -d "$BODY" localhost:4000/v1/movies
HTTP/1.1 422 Unprocessable Entity
Content-Type: application/json
Date: Wed, 07 Apr 2021 10:33:57 GMT 
Content-Length: 180

{
    "error":
    {
        "genres": "must not contain duplicate values",
        "runtime": "must be a positive integer",
        "title": "must be provided",
        "year": "must be greater than 1888"
    }
}
```

看起来不错。我们的检查功能生效了，并阻止请求被执行—甚至更好的是，向客户端返回一个格式良好的JSON响应，其中包含针对每个检验错误的详细信息。

你也可以发送正常的请求体，你会发现请求被正常处理，input内容在响应中返回给客户端：

```shell
$ BODY='{"title":"Moana","year":2016,"runtime":"107 mins","genres":["animation","adventure"]}'
$ curl -i -d "$BODY" localhost:4000/v1/movies
HTTP/1.1 200 OK
Date: Tue, 23 Nov 2021 12:33:45 GMT
Content-Length: 65
Content-Type: text/plain; charset=utf-8

{Title:Moana Year:2016 Runtime:107 Genres:[animation adventure]}

```

## 使校验规则可重用

在大型项目中，很多个接口需要重复这种校验的过程，因此将上面的校验规则抽象成方法供其他地方使用。比如客户端要更新movie也会传一些新的字段内容，也需要校验。

避免重复，我们可以将movie的校验整合到一个单独的ValidateMovie()函数中去。理论上，这个函数可以放在任意位置。但就个人而言，我喜欢将验证检查放在internal/data包中的相关领域类型附近。

如果按照下面的步骤操作，请重新打开内部/data/movies.go然后添加一个ValidateMovie()函数，其中包含如下检查:

File: internal/data/movies.go

---

```go
package data

import (
	"encoding/json"
	"fmt"
	"greenlight.alexedwards.net/internal/validator"
	"time"
)

type Movie struct {
	ID       int64     `json:"id"`
	CreateAt time.Time `json:"-"`
	Title    string    `json:"title"`
	Year     int32     `json:"year,omitempty"`
	Runtime  Runtime   `json:"runtime,omitempty,string"`
	Genres   []string  `json:"genres,omitempty"`
	Version  int32     `json:"version"`
}

func ValidateMovie(v *validator.Validator, movie *Movie) {
	v.Check(movie.Title != "", "title", "must be provide")
	v.Check(len(movie.Title) <= 500, "title", "must not be more than 500 bytes long")

	v.Check(movie.Year != 0, "year", "must be provided")
	v.Check(movie.Year >= 1888, "year", "must be greater than 1888")
	v.Check(movie.Year <= int32(time.Now().Year()), "year", "must not be in the future")

	v.Check(movie.Runtime != 0, "runtime", "must be provided")
	v.Check(movie.Runtime > 0, "runtime", "must be a positive integer")

	v.Check(movie != nil, "genres", "must be provided")
	v.Check(len(movie.Genres) >= 1, "genres", "must contain at least 1 genres")
	v.Check(len(movie.Genres) <= 5, "genres", "must not contain more than 5 genres")
	v.Check(validator.Unique(movie.Genres), "genres", "must not contain duplicate values")
}
```

> 重要提示:现在检查是对一个movie结构体实例各个字段进行的，而不是对input结构体。

完成上面的改造之后，我们需要返回createMovieHandler并更新代码，通过初始化一个新的Movie结构体，从input结构体复制数据到movie结构体中，然后调用这个新的验证函数。像这样:

File：cmd/api/movies.go

---

```go
package main

...

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string       `json:"title"`
		Year    int32        `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Genres  []string     `json:"genres"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	movie := &data.Movie{
		Title:   input.Title,
		Year:    input.Year,
		Runtime: input.Runtime,
		Genres:  input.Genres,
	}

        //初始化Validator实例
	v := validator.New()

        调用ValidateMovie()函数，如果有错误就返回给客户端。
	if data.ValidateMovie(v, movie); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	fmt.Fprintf(w, "%+v\n", input)
}
```

当您查看这些代码时，您的脑海中可能会有几个问题。

首先，您可能想知道为什么我们在处理程序中初始化Validator实例并将其传递给ValidateMovie()函数——而不是在ValidateMovie()中初始化它并将其作为返回值传递回来。

这是因为随着应用程序变得越来越复杂，我们将需要从处理程序调用多个校验帮助函数，而不是像上面所示的一个。因此，在处理程序中初始化Validator，然后传递给帮助函数，这给了我们更多的灵活性。

您可能还想知道，为什么我们要将JSON请求解码为input结构体类型，然后复制数据，而不是直接解码为Movie结构体实例。

因为movie里面有些字段例如ID和Version是不需要客户端提供的，如果使用movie的话，客户端提供ID和Verison字段也会被解码到movie结构体中，这就需要多余的检查工作。

但是将客户端的请求内容解析到一个临时的结构体中，会更灵活，简洁而且代码更健壮。

有了这些解释，您应该能够再次启动应用程序，并且从客户端的角度来看，效果应该与之前的一样。如果你发起一个无效的请求，你应该会得到一个包含类似这样的错误消息的响应:

```shell
$ BODY='{"title":"","year":1000,"runtime":"-123 mins","genres":["sci-fi","sci-fi"]}' 
$ curl -i -d "$BODY" localhost:4000/v1/movies
HTTP/1.1 422 Unprocessable Entity
Content-Type: application/json
Date: Wed, 07 Apr 2021 10:33:57 GMT 
Content-Length: 180

{
    "error":
    {
        "genres": "must not contain duplicate values",
        "runtime": "must be a positive integer",
        "title": "must be provided",
        "year": "must be greater than 1888"
    }
}
```

您可以随意测试，并尝试在JSON中发送不同的值，直到所有的校验都按预期工作为止。
## 解析JSON请求

到目前为止，我们一直在研究如何从我们的API中创建和发送JSON响应，但在本书的下一节中，我们将从另一方面探索，并讨论如何读取和解析来自客户端的JSON请求。

为了帮助说明这一点，我们将从POST /v1/movies接口和之前设置的createMovieHandler上开始工作。


| Method | URL             | Handler            | 动作               |
| -------- | ----------------- | -------------------- | -------------------- |
| GET    | /v1/healthcheck | healthcheckHanlder | 查询应用程程序信息 |
| POST   | /v1/movies      | createMovieHandler | 创建新的电影       |
| GET    | /v1/movies/:id  | showMovieHandler   | 查询特定电影详情   |

当客户端调用这个接口时，我们希望它们提供一个JSON请求体，其中包含想要在我们的系统中创建的电影的数据。例如，如果客户端想要为电影Moana添加一条记录到我们的API中，会发送一个类似于这样的请求体:

```json
{
    "title": "Moana",
    "year": 2016,
    "runtime": 107,
    "genres":
    [
        "animation",
        "adventure"
    ]
}
```

现在，我们只关注处理这个JSON请求体的读取、解析和验证。你将学习:

* 如何使用encoding/json包读取请求体并将其反序列化为本地Go对象。
* 如何处理来自客户端的错误请求和无效的JSON，并返回清晰的、可操作的错误消息。
* 如何创建可重用的帮助程序来验证数据，以确保数据符合业务规则。
* 控制和定制JSON解码方式的不同技术。
## 配置PostgreSQL数据库

#### 安装PostgreSQL

如果您跟随本书步骤操作，那么此时您需要在计算机上安装PostgreSQL。官方的PostgreSQL文档包含了所有类型操作系统的下载和安装说明，但如果你使用的是macOS，你应该能够使用brew安装:

```shell
$ brew install postgresql
```

如果您使用的是Linux发行版，您应该能够通过包管理器来安装它。例如，如果你的操作系统支持apt包管理器(像Debian和Ubuntu一样)，你可以这样安装:

```shell
$ sudo apt install postgresql
```

在Windows机器上，你可以使用Chocolatey包管理器安装PostgreSQL:

```shell
> choco install postgresql
```

#### 命令行连接PostgreSQL

当PostgreSQL被安装的时候，你的计算机上应该已经创建了一个psql二进制文件。它包含一个基于终端的命令行工具，用于使用PostgreSQL。

你可以通过在你的终端上运行psql --version命令来检查它是否可用，如下所示:

```shell
$ psql --version
psql (PostgreSQL) 13.4
```

如果您还不熟悉PostgreSQL，那么第一次使用psql连接到它的过程可能有点不直观。让我们花点时间来解释一下。

当PostgreSQL刚安装时，它只有一个用户用户：一个叫做postgres的超级用户。在第一次使用时，我们需要以该超级用户身份连接到PostgreSQL来做任何操作-此时我们可以执行任何需要的设置，如创建数据库和创建其他用户。

在安装期间，已经在您的机器上创建一个名为postgres的操作系统用户。在基于unix的系统上，你可以检查/etc/passwd文件来确认这一点，像这样:

```shell
$ cat /etc/passwd|grep 'postgres'
_postgres:*:216:216:PostgreSQL Server:/var/empty:/usr/bin/false
```

这很重要，因为默认情况下，PostgreSQL使用对等认证(peer authentication)的身份验证方案，对来自本地机器的任何连接进行验证。对等认证是指如果当前操作系统用户的用户名与一个有效的PostgreSQL用户名匹配，可以作为该用户登录PostgreSQL，而无需进一步认证。不涉及密码。

因此，如果我们切换到名为postgres的操作系统用户，我们应该能够使用psql连接到PostgreSQL，而不需要任何进一步的身份验证。事实上，你可以用下面的命令一步完成这两件事:

```shell
$ sudo -u postgres psql
psql (13.4)
Type "help" for help.

postgres=# 
```

因此，为了确认一下，我们在这里使用sudo命令(superuser do)以操作系统用户postgres的身份运行psql命令。这将在在终端打开一个会话，以PostgreSQL超级用户(称为postgres)进行身份验证。

你可以通过运行“SELECT current_user”来确认你当前是哪个PostgreSQL用户:

```shell
postgres=# SELECT current_user;
current_user 
--------------
postgres 
(1 row)
```

#### 创建数据库，用户和扩展

当我们以postgres超级用户的身份连接时，可以为我们的项目创建一个名为greenlight的新数据库，然后使用\c命令连接到数据库，如下所示:

```shell
postgres=# create database greenlight;
CREATE DATABASE
postgres=# \c greenlight
You are now connected to database "greenlight" as user "postgres".
```

> 在PostgreSQL中，\字符表示元命令。其他一些有用的元命令有\l用于列出所有数据库，\dt用于列出表，以及\du用于列出用户。你还可以运行\?查看可用元命令的完整列表。

现在我们的greenlight数据库已经存在并连接到它，有两个任务需要完成。

第一个任务是创建一个没有超级用户权限的新用户greenlight，我们可以使用该用户执行SQL迁移并从Go应用程序连接到数据库。我们希望将这个新用户设置为使用基于密码的身份验证，而不是对等身份验证。

PostgreSQL也有扩展的概念，它在标准功能的基础上增加额外的功能。PostgreSQL附带的扩展列表可以在[这里](https://www.postgresql.org/docs/current/contrib.html)找到，还有一些[其他扩展](https://www.postgresql.org/download/products/6-postgresql-extensions/)你可以单独下载。

在这个项目中，我们将使用[citext](https://www.postgresql.org/docs/9.5/citext.html)扩展。该扩展给PostgreSQL添加了区分大小写的字符串类型，我们将在本书后面使用它来存储用户电子邮件地址。需要注意的是，扩展只能由超级用户添加到特定的数据库中。

继续，并运行以下命令来创建一个新的greenlight用户与特定的密码，并添加citext扩展到我们的数据库:

```shell
greenlight=# CREATE ROLE greenlight WITH LOGIN PASSWORD 'pa55word';
CREATE ROLE
greenlight=# CREATE EXTENSION IF NOT EXISTS citext;
CREATE EXTENSION
```

> 重要提示:如果您按照此方法进行操作，请务必记住您为greenlight用户设置的密码。在接下来的步骤中您将需要它。

一旦成功完成以上，您可以键入exit或\q来关闭终端与数据库连接，并恢复为您的正常操作系统用户。

```shell
 greenlight=# exit
```

#### 以新用户连接数据库

在进行进一步操作之前，让我们证明之前的操作都设置正确，并尝试以greenlight用户连接到greenlight数据库。出现提示时，输入在上面步骤中设置的密码。

```shell
$ psql --host=localhost --dbname=greenlight --username=greenlight
Password for user greenlight:
psql (13.4)
SSL connection (protocol: TLSv1.3, cipher: TLS_AES_256_GCM_SHA384, bits: 256, compression: off) Type "help" for help.
greenlight=> SELECT current_user;
current_user --------------
greenlight (1 row)
greenlight=> exit
```

很好！这确认了我们的数据库和带有密码验证的新greenlight用户可以正常使用，并且我们能够以该用户执行SQL语句。

## 附加内容

#### 优化PostgreSQL设置

PostgreSQL的默认设置是相当保守的，通常可以通过调整PostgreSQL.conf文件中的值来提高数据库的性能。

你可以通过下面的SQL查询来检查postgresql.conf文件的位置:

```shell
$ sudo -u postgres psql -c 'SHOW config_file;'
                config_file   
--------------------------------------------
 /opt/homebrew/var/postgres/postgresql.conf
(1 row)
```

[这篇文章](https://www.enterprisedb.com/postgres-tutorials/how-tune-postgresql-memory)地介绍了一些最重要的PostgreSQL配置，并指导哪些值可以作为合理的其实配置。如果你对优化PostgreSQL感兴趣，我建议你读一读这篇文章。

或者，您可以使用这个基于web的[工具](https://pgtune.leopard.in.ua)来根据你的系统硬件生成推荐配置。这个工具的特点是输出"ALTER SYSTEM" SQL语句，你可以运行它来改变你的数据库设置，而不是手动修改postgresql.conf文件。
## 连接PostgreSQL数据库

好了，现在我们的新greenlight数据库已经创建好了，来看看如何从Go应用程序连接到它。

要连接SQL数据库，我们需要使用一个数据库驱动程序，作为Go和数据库之间的“中间人”。你可以在[Go wiki](https://github.com/golang/go/wiki/SQLDrivers)中找到一个可用的PostgreSQL驱动程序列表，但对于我们的项目，我们将选择pq包，它的特点是常用、可靠的和功能完善。

如果你跟随本书操作，那么使用go get下载pq的v1.10.0版本，就像这样:

```shell
$ go get github.com/lib/pq@v1.10.0
go: downloading github.com/lib/pq v1.10.0
```

为了连接到数据库，我们还需要一个数据源名称(DSN)，它是一个包含必要连接参数的字符串。DSN格式将取决于您使用的数据库驱动程序(应该在驱动程序文档中有描述)，但当使用pq时，以greenlight用户结合DSN，您应该能够连接到您的本地greenlight数据库:

```shell
 postgres://greenlight:pa55word@localhost/greenlight
```

#### 创建数据库连接池

我们在Go应用程序连接到greenlight数据库的代码几乎与mysql中的代码完全相同。我就不细讲了，希望大家都很熟悉。

包含以下几点：

* 我们希望DSN在运行时可配置，因此我们将使用命令行参数将它传递给应用程序，而不是硬编码它。在开发过程中，为了简单起见，我们将使用上面的DSN作为参数的默认值。
* 在cmd/api/main.go文件中我们创建一个新的openDB()函数。在该函数中，使用sql.Open()来建立sql.DB连接池，因为数据库连接是在第一次需要时惰性地建立的。因此，我们还需要使用db.PingContext()方法来创建一个实际的连接，并验证数据库设置正确。

我们回到cmd/api/main.go文件更新代码如下：

File: cmd/api/main.go

---

```go
package main

import (
	"context" 
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
        //导入pq驱动程序，通过包database/sql来注册数据库驱动程序
	_ "github.com/lib/pq"
)

const version = "1.0.0"

//添加一个db结构字段来保存数据库连接的配置
type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
}
type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	// 将DSN值从db-dsn命令行参数中读入配置，设置默认值
	flag.StringVar(&cfg.db.dsn, "db-dsn", "postgres://greenlight:pa55word@localhost/greenlight", "PostgreSQL DSN")
	flag.Parse()
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
  
	//调用openDB()帮助函数(参见下面)来创建连接池
	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}
  
	// defer调用， 以便main()函数退出之前关闭连接池。
	defer db.Close()
  
	//打印连接数据库成功日志
	logger.Printf("database connection pool established")
	app := &application{config: cfg,
		logger: logger}
	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", cfg.port), Handler: app.routes(),
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second, WriteTimeout: 30 * time.Second,
	}
	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err = srv.ListenAndServe()
	logger.Fatal(err)
}

//openDB()函数返回一个sql.DB连接池。
func openDB(cfg config) (*sql.DB, error) {
	//使用sql.Open()创建一个空连接池
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}
 
	//创建一个具有5秒超时期限的上下文。
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
 
	//使用PingContext()建立到数据库的新连接，并传入上下文信息，连接超时就返回
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	// 返回sql.DB连接池
	return db, nil
}
```

> 重要提示：我们将在本书的后面讨论如何使用上下文来正确地管理超时，所以现在不要过多地担心这个问题。现在，只需要知道PingContext()调用不能在5秒内成功完成，将返回一个超时错误。

接下来启动应用程序，你将看到条关于连接数据库成功的日志，如下所示：

```shell
$ go run ./cmd/api
2021/04/07 15:56:54 database connection pool established 
2021/04/07 15:56:54 starting development server on :4000
```

#### 解耦DSN

目前，DSN的命令行参数默认值是在cmd/api/main.go文件中写死的。尽管DSN中的用户名和密码只是用于您自己机器上的开发数据库，但最好不要将这些敏感信息硬编码到项目文件中(这些文件可能在未来被共享或分发)。

因此，我们采取一些步骤将DSN从我们的项目代码中解耦，并将其存储为本地机器上的一个环境变量。

如果你按照本书操作，先创建GREENLIGHT_DB_DSN环境变量，在$HOME/.profile文件中添加如下内容：

File: $HOME/.profile

```shell
...
export GREENLIGHT_DB_DSN='postgres://greenlight:pa55word@localhost/greenlight'
```

或者在$HOME/.bashrc文件中添加以上内容也可以。

完成这些之后，需要重新启动您的计算机，如果现在不方便重启的话，可以在刚刚编辑的文件上运行source命令来实现环境变量生效。例如:

```shell
$ source $HOME/.profile
```

> 注意:运行source命令只会在当前的终端窗口生效。因此，如果您切换到另一个终端，您需要再次运行source生效。

不管怎样，一旦您重启机器或运行source，通过运行echo命令就可以在终端中看到GREENLIGHT_DB_DSN环境变量的值。像这样:

```shell
 $ echo $GREENLIGHT_DB_DSN
postgres://greenlight:pa55word@localhost/greenlight
```

下面更新cmd/api/main.go文件，可以通过访问环境变量来获取DSN值。Go中读取环境变量值可以调用os.Getenv()函数，然后将DSN命令行参数默认值设置为环境变量值。

File：cmd/api/main.go

---

```go
package main

...

func main() {
var cfg config

flag.IntVar(&cfg.port, "port", 4000, "API server port")
flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

// 使用GREENLIGHT_DB_DSN环境变量，作为DSN命令行参数默认值
flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("GREENLIGHT_DB_DSN"), "PostgreSQL DSN")
flag.Parse()
...
}
```

如果您现在再次重新启动应用程序，您应该会发现可以正常编译，并像之前一样工作。

您还可以尝试在运行应用程序时指定-help标志。应该会输出三个命令行参数的描述内容和默认值，包括从环境变量中提取的DSN值。类似于:

```go
$ go run ./cmd/api -help
Usage of /tmp/go-build417842398/b001/exe/api: 
  -db-dsn string
    PostgreSQL DSN (default "postgres://greenlight:pa55word@localhost/greenlight") 
  -env string
    Environment (development|staging|production) (default "development") 
  -port int
    API server port (default 4000)
```

## 附加内容

#### psql使用DSN连接数据库

将DSN存储在环境变量中的一个很好处是：能够以greenlight用户轻松地连接到greenlight数据库，而不是在运行psql时手动指定所有连接配置。像这样:

```shell
$ psql $GREENLIGHT_DB_DSN
psql (13.4)
SSL connection (protocol: TLSv1.3, cipher: TLS_AES_256_GCM_SHA384, bits: 256, compression: off) 
Type "help" for help.

greenlight=>
```
## 配置数据库连接池

本节内容我们将解释连接池背后是如何工作的，并探索如何配置数据库能改变或优化其性能。

> 注意：本章包含了相当多的理论，虽然它很有趣，但对应用程序的构建并不重要。如果你觉得它太难了，可以先浏览一下，然后再回头看。

那么sql.DB连接池是如何工作的呢?

需要理解的最重要一点是，sql.DB池包含两种类型的连接——“正在使用”连接和“空闲”连接。当您使用连接执行数据库任务(例如执行SQL语句或查询行)时，该连接被标记为正在使用，任务完成后，该连接被标记为空闲。

当您使用Go执行数据库操作时，它将首先检查池中是否有可用的空闲连接。如果有可用的连接，那么Go将重用这个现有连接，并在任务期间将其标记为正在使用。如果在您需要空闲连接时池中没有空闲连接，那么Go将创建一个新的连接。

当Go重用池中的空闲连接时，与该连接有关的任何问题都会被优雅地处理。异常连接将在放弃之前自动重试两次，这时Go将从池中删除异常连接并创建一个新的连接来执行该任务。

#### 配置连接池

连接池有四个方法，我们可以使用它们来配置连接池的行为。让我们一个一个地来讨论。

##### SetMaxOpenConns方法

SetMaxOpenConns()方法允许您设置池中“打开”连接(使用中+空闲连接)数量的上限。默认情况下，打开的连接数是无限的。

> 注意“打开”连接等于“正在使用”加上“空闲”连接，不仅仅是“正在使用”连接。

一般来说，MaxOpenConns设置得越大，可以并发执行的数据库查询就越多，连接池本身成为应用程序中的瓶颈的风险就越低。

但让它无限并不是最好的选择。默认情况下，PostgreSQL最多100个打开连接的硬限制，如果达到这个限制的话，它将导致pq驱动返回"sorry, too many clients already"错误。

> 注意：最大打开连接数限制可以在postgresql.conf文件中对max_connections设置来更改。

为了避免这个错误，将池中打开的连接数量限制在100以下是有意义的，可以为其他需要使用PostgreSQL的应用程序或会话留下足够的空间。

设置MaxOpenConns限制的另一个好处是，它充当一个非常基本的限流器，防止数据库同时被大量任务压垮。

但设定上限有一个重要的警告。如果达到MaxOpenConns限制，并且所有连接都在使用中，那么任何新的数据库任务将被迫等待，直到有连接空闲。在我们的API上下文中，用户的HTTP请求可能在等待空闲连接时无限期地“挂起”。因此，为了缓解这种情况，使用上下文为数据库任务设置超时是很重要的。我们将在书的后面解释如何处理。

##### SetMaxIdleConns方法

SetMaxIdleConns()方法的作用是：设置池中空闲连接数的上限。缺省情况下，最大空闲连接数为2。

理论上，在池中允许更多的空闲连接将增加性能。因为它减少了从头建立新连接发生概率—，因此有助于节省资源。

但要意识到保持空闲连接是有代价的。它占用了本来可以用于应用程序和数据库的内存，而且如果一个连接空闲时间过长，它也可能变得不可用。例如，默认情况下MySQL会自动关闭任何8小时未使用的连接。

因此，与使用更小的空闲连接池相比，将MaxIdleConns设置得过高可能会导致更多的连接变得不可用，浪费资源。因此保持适量的空闲连接是必要的。理想情况下，你只希望保持一个连接空闲，可以快速使用。

另一件要指出的事情是MaxIdleConns值应该总是小于或等于MaxOpenConns。Go会强制保证这点，并在必要时自动减少MaxIdleConns值。

##### SetConnMaxLifetime方法

SetConnMaxLifetime()方法用于设置ConnMaxLifetime的极限值，表示一个连接保持可用的最长时间。默认连接的存活时间没有限制，永久可用。

如果设置ConnMaxLifetime的值为1小时，意味着所有的连接在创建后，经过一个小时就会被标记为失效连接，标志后就不可复用。但需要注意：

* 这并不能保证一个连接将在池中存在一整个小时；有可能某个连接由于某种原因变得不可用，并在此之前自动关闭。
* 连接在创建后一个多小时内仍然可以被使用—只是在这个时间之后它不能被重用。
* 这不是一个空闲超时。连接将在创建后一小时过期，而不是在空闲后一小时过期。
* Go每秒运行一次后台清理操作，从池中删除过期的连接。

理论上，ConnMaxLifetime为无限大（或设置为很长生命周期）将提升性能，因为这样可以减少新建连接。但是在某些情况下，设置短期存活时间有用。比如：

* 如果SQL数据库对连接强制设置最大存活时间，这时将ConnMaxLifetime设置稍短时间更合理。
* 有助于数据库替换

如果您决定对连接池设置ConnMaxLifetime，那么一定要记住连接过期(然后重新创建)的频率。例如，如果连接池中有100个打开的连接，而ConnMaxLifetime为1分钟，那么您的应用程序平均每秒可以杀死并重新创建多达1.67个连接。您不希望频率太大而最终影响性能吧。

##### SetConnMaxIdleTime方法

SetConnMaxIdleTime()方法在Go 1.15版本引入对ConnMaxIdleTime进行配置。其效果和ConnMaxLifeTime类似，但这里设置的是：在被标记为失效之前一个连接最长空闲时间。例如，如果我们将ConnMaxIdleTime设置为1小时，那么自上次使用以后在池中空闲了1小时的任何连接都将被标记为过期并被后台清理操作删除。

这个配置非常有用，因为它意味着我们可以对池中空闲连接的数量设置相对较高的限制，但可以通过删除不再真正使用的空闲连接来周期性地释放资源。

#### 实操一波

所以有很多信息要吸收。这在实践中意味着什么？我们把以上所有的内容总结成一些可行的要点。

1、根据经验，您应该显式地设置MaxOpenConns值。这个值应该低于数据库和操作系统对连接数量的硬性限制，您还可以考虑将其保持在相当低的水平，以充当基本的限流作用。

对于本书中的项目，我们将MaxOpenConns限制为25个连接。我发现这对于小型到中型的web应用程序和API来说是一个合理的初始值，但理想情况下，您应该根据基准测试和压测结果调整这个值。

2、通常，更大的MaxOpenConns和MaxIdleConns值会带来更好的性能。但是，效果是逐渐降低的，而且您应该注意，太多的空闲连接(连接没有被复用)实际上会导致性能下降和不必要的资源消耗。

因为MaxIdleConns应该总是小于或等于MaxOpenConns，所以对于这个项目，我们还将MaxIdleConns限制为25个连接。

3、为了降低上面第2点的风险，通常应该设置ConnMaxIdleTime值来删除长时间未使用的空闲连接。在这个项目中，我们将设置ConnMaxIdleTime持续时间为15分钟。

4、ConnMaxLifetime默认设置为无限大是可以的，除非您的数据库对连接生命周期施加了硬限制，或者您需要它协助一些操作，比如优雅地交换数据库。这些都不适用于本项目，所以我们将保留这个默认的无限制配置。

#### 配置连接池

与其硬编码这些配置，不如更新cmd/api/main.go文件通过命令行参数读取配置。

ConnMaxIdleTime值比较有意思，因为我们希望它传递一段时间，最终需要将其转换为Go的time.Duration类型。这里有几个选择:

1、我们可以使用一个整数来表示秒(或分钟)的数量，并将其转换为time.Duration。

2、我们可以使用一个表示持续时间的字符串——比如“5s”(5秒)或“10m”(10分钟)——然后使用time.ParseDuration()函数解析它。

3、两种方法都可以很好地工作，但是在这个项目中我们将使用选项2。继续并更新cmd/api/main.go文件如下:

File: cmd/api/main.go

---

```go
package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

//添加maxOpenConns, maxIdleConns和maxIdleTime字段来存放连接池配置
type config struct {
	port int
	env  string
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  int
	}
}
type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.db.dsn, "db-dsn", "postgres://greenlight:pa55word@localhost/greenlight", "PostgreSQL DSN")

	//从命令参数中读取连接池配置到config结构体中
	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.StringVar(&cfg.db.maxIdleTime, "db-max-idle-time", "15m", "PostgreSQL max connection idle time")
	flag.Parse()
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	//调用openDB()帮助函数(参见下面)来创建连接池
	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}

	// defer调用， 以便main()函数退出之前关闭连接池。
	defer db.Close()

	//打印连接数据库成功日志
	logger.Printf("database connection pool established")
	app := &application{config: cfg,
		logger: logger}
	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", cfg.port), Handler: app.routes(),
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second, WriteTimeout: 30 * time.Second,
	}
	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err = srv.ListenAndServe()
	logger.Fatal(err)
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	//设置最大开放连接数，注意该值为小于0或等于0指的是无限制连接数
	db.SetMaxOpenConns(cfg.db.maxOpenConns)

	//设置空闲连接数，小于等于0表示无限制
	db.SetMaxIdleConns(cfg.db.maxIdleConns)

	//将空闲时间字符串解析为time.Duration类型
	duration, err := time.ParseDuration(cfg.db.maxIdleTime)
	if err != nil {
		return nil, err
	}

	//设置最大空闲超时
	db.SetConnMaxIdleTime(duration)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}
```

如果再次运行该应用程序，那么一切应该都能正常工作。您不会发现有任何更改，而且此时我们无法演示这些配置的作用。但是在本书的后面，我们将进行一些负载测试，并解释如何使用db.Stats()方法实时监视连接池的状态。到那时，你就能看到我们在这一章中讨论的一些设置。
## 数据库安装与配置

在本书的下一部分中，我们将继续我们的项目，构建并设置一个SQL数据库来持久存储我们的movie数据。

这里我们将使用PostgreSQL数据库。它是开源的，非常可靠，并且具有一些有用的新的特性——包括支持数组和JSON数据类型、全文搜索和地理空间查询。在项目构建过程中，我们将使用这些新的PostgreSQL特性。

本章你将学习：

* 如何在自己机器上安装和设置PostgreSQL。
* 如何使用psql交互式工具创建数据库，PostgreSQL扩展和用户帐户。
* 如何在Go中初始化数据库连接池并修改配置以提高性能和稳定性。
## SQL迁移概述

如果你不熟悉SQL迁移的概念，作为一个非常高级的概念它是这样的工作的:

1、对数据库模式进行的每一次更改(如创建表、添加列或删除未使用的索引)，都要创建一对迁移文件。一个文件是“up”迁移，其中包含实现更改所需的SQL语句，另一个文件是“down”迁移，其中包含逆转(或回滚)更改所需的SQL语句。

2、每对迁移文件按顺序编号，通常是0001、0002、0003…或者使用Unix时间戳，表示将迁移应用到数据库的顺序。

3、您可以使用某种工具或脚本，根据迁移文件中的SQL语句执行或回滚数据库。该工具跟踪哪些迁移已经生效，以便执行必要的SQL语句。

使用迁移来管理你的数据库模式，而不是自己手动执行SQL语句，有几个好处:

* 数据库模式(及其演变和更改)完全由“up”和“down”SQL迁移文件描述。由于这些文件只是包含一些SQL语句，所以可以在版本控制系统中将它们与其他代码一起跟踪。
* 通过运行“up”迁移，可以在另一台机器上精确地复制当前数据库模式。当您需要在不同的环境(开发、测试、生产等)中管理和同步数据库模式时，这是一个很大的帮助。
* 如果有必要，可以通过应用适当的“down”迁移回滚数据库模式更改。

#### 按照迁移工具

为了在这个项目中管理SQL迁移，我们将使用migrate命令行工具(它本身是用Go编写的)。

不同操作系统的详细安装说明可以在[这里](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)找到，但在macOS上，你应该能够使用以下命令进行安装:

```shell
brew install golang-migrate
```

在Linux和Windows上，最简单的方法是下载二进制文件并将其拷贝到系统路径上的某个位置。例如，在Linux上:

```shell
$ curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz 
$ mv migrate.linux-amd64 $GOPATH/bin/migrate
```

在继续之前，请检查它是否可用，在您的机器上运行，尝试使用-version命令行参数执行migrate二进制文件。它输出当前版本号，如下所示:

```shell
$ migrate -version
4.14.1
```
## 使用SQL迁移

现在已经安装了迁移工具，让我们通过在数据库中创建一个新的movies表来演示如何使用它。

首先使用migrate create创建一对迁移文件。跟随本书在终端中运行以下命令:

```shell
$ migrate create -seq -ext=.sql -dir=./migrations create_movies_table
/Users/wangmingjun/Projects/greenlight/migrations/000001_create_movies_table.up.sql
/Users/wangmingjun/Projects/greenlight/migrations/000001_create_movies_table.down.sql
```

在这个命令中：

* -seq命令行参数表示：我们想要使用顺序编号为迁移文件命名，比如0001,0002...(而不是Unix时间戳，这是默认值)。
* -ext命令行参数表示：迁移文件使用.sql扩展名。
* -dir表示：创建的迁移文件存放目录。（如果目录不存在就会自动创建）
* create_movies_table是一个描述性的标签，便于识别迁移文件的内容。

现在进入migrations目录，你将看到一对新建对"up"和“down”迁移文件：

```shell
./migrations/
├── 000001_create_movies_table.down.sql 
└── 000001_create_movies_table.up.sql
```

此时这两个文件还是空的。下面我们编辑'up'迁移文件，添加CREATE TABLE语句来创建movies表。如下所示：

File: imgrations/000001_create_movies_table.up.sql

---

```sql
CREATE TABLE IF NOT EXISTS movies (
id bigserial PRIMARY KEY,
created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(), title text NOT NULL,
year integer NOT NULL,
runtime integer NOT NULL,
genres text[] NOT NULL,
version integer NOT NULL DEFAULT 1
);
```

注意这个表中的字段和类型与我们之前创建的Movie结构中的字段和类型相似吗？这很重要，因为这意味着我们可以轻松地将movies表中的每一行数据映射到Go代码中的单个Movie结构体。

如果您对上面SQL语句中的不同PostgreSQL数据类型不熟悉，可以查看[官方文档](https://www.postgresql.org/docs/current/datatype.html)，上面提供了全面的概述。但需要指出的最重要的事情是:

* id列类似为[bigserial](https://www.postgresql.org/docs/12/datatype-numeric.html#DATATYPE-NUMERIC-TABLE)它是一个64位自增整数，从1开始是表的主键。
* genres列类型为text[]，是一个包含零个或多个text值数组。需要注意的是，PostgreSQL中的数组本身是可查询和可索引的，这一点我们将在本书的后面进行演示。
* 在可能的情况下，在每个列上设置NOT NULL约束和适当的默认值是最简单的——就像我们上面所做的那样。
* 对于存储字符串，我们使用文本类型，而不是varchar或varchar(n)类型。如果你感兴趣，这篇[博文](https://www.depesz.com/2010/03/02/charx-vs-varcharx-vs-varchar-vs-text/)解释了为什么文本通常是PostgreSQL中使用的最佳字符类型。

好了，让我们继续“down”迁移，并添加对应的SQL语句来回退我们刚刚编写的“up”迁移。

File: migrations/000001_create_movies_table.down.sql

---

```sql
DROP TABLE IF EXISTS movies;
```

在PostgreSQL中，DROP TABLE命令是删除表中存在的所有索引和约束，所以这一条语句就足以逆转上面对“up”操作。

ok，前面是我们准备好的第一对迁移文件!

同时，让我们创建第二个对含[CHECK](https://www.postgresql.org/docs/9.4/ddl-constraints.html)约束的迁移文件，在数据库级别上强制执行一些业务规则。具体来说，我们希望确保runtime值总是大于零，year值在1888和当前年份之间，并且genres数组总是包含1到5个条目。

执行以下命令来创建第二对迁移文件：

```shell
$ migrate create -seq -ext=.sql -dir=./migrations add_movies_check_constrains

/Users/wangmingjun/Projects/greenlight/migrations/000002_add_movies_check_constrains.up.sql
/Users/wangmingjun/Projects/greenlight/migrations/000002_add_movies_check_constrains.down.sql

```

然后添加以下SQL语句，分别添加和删除CHECK约束:

File: migrations/000002_add_movies_check_constraints.up.sql

---

```sql
ALTER TABLE movies ADD CONSTRAINT movies_runtime_check CHECK (runtime >= 0);
ALTER TABLE movies ADD CONSTRAINT movies_year_check CHECK (year BETWEEN 1888 AND date_part('year', now()));
ALTER TABLE movies ADD CONSTRAINT genres_length_check CHECK (array_length(genres, 1) BETWEEN 1 AND 5);
```

File: File: migrations/000002_add_movies_check_constraints.down.sql

---

```sql
ALTER TABLE movies DROP CONSTRAINT IF EXISTS movies_runtime_check;
ALTER TABLE movies DROP CONSTRAINT IF EXISTS movies_year_check;
ALTER TABLE movies DROP CONSTRAINT IF EXISTS genres_check;
```

当我们在movies表中插入或更新数据时，如果这些检查失败，我们的数据库驱动程序将返回类似这样的错误:

```shell
pq: new row for relation "movies" violates check constraint "movies_year_check"
```

> 注意:单个迁移文件包含多个SQL语句是完全可以的，正如我们在上面的两个文件中看到的那样。事实上，我们本可以在第一对迁移文件中包含CHECK约束和CREATE TABLE语句，但是出于本书的目的，将它们放在单独的第二次迁移中可以帮助我们说明迁移工具是如何工作的。

#### 执行迁移

现在我们准备在greenlight数据库上运行两个“up”迁移。

如果您跟随前面的步骤进行操作，那么使用下面的命令执行迁移，并从您的环境变量传入数据库DSN。如果一切都配置正确，您应该看到一些输出，确认迁移已经成功执行。类似于:

```shell
$ migrate -path=./migrations -database=$GREENLIGHT_DB_DSN up
1/u create_movies_table (48.079834ms)
2/u add_movies_check_constrains (71.984417ms)
```

此时，你使用psql连接到数据库，并使用\dt命令列出刚创建对表:

```shell
$ psql $GREENLIGHT_DB_DSN
psql (13.4)
Type "help" for help.

greenlight=> \dt
                 List of relations
 Schema |       Name        |   Type   |   Owner  
--------+-------------------+----------+------------
 public | movies            | table    | greenlight
 public | schema_migrations | table    | greenlight
(2 rows)
```

您应该看到已经创建了movies表和schema_migrations表，这两个表都由greenlight用户拥有。

schema_migrations表是由迁移工具自动生成的，用于跟踪哪些迁移已经生效。让我们快速浏览一下它的表内容:

```shell
greenlight=> select * from schema_migrations;
 version | dirty 
---------+-------
       2 | f
(1 row)
```

这里的version列表明已经在数据库上执行了(包括)序列文件中的第2个迁移文件。dirty列的值为false，这表明迁移文件被干净地执行，没有任何错误，其中包含的SQL语句全部成功地执行。

如果您愿意，还可以在movies表上运行\d命令查看表结构，并确认CHECK约束已正确创建。像这样:

```shell
greenlight=> \d movies;
                                        Table "public.movies"
  Column   |            Type             | Collation | Nullable |              Default       
-----------+-----------------------------+-----------+----------+------------------------------------
 id        | bigint                      |           | not null | nextval('movies_id_seq'::regclass)
 create_at | timestamp(0) with time zone |           | not null | now()
 title     | text                        |           | not null | 
 year      | integer                     |           | not null | 
 runtime   | integer                     |           | not null | 
 genres    | text[]                      |           | not null | 
 version   | integer                     |           | not null | 1
Indexes:
    "movies_pkey" PRIMARY KEY, btree (id)
Check constraints:
    "genres_length_check" CHECK (array_length(genres, 1) >= 1 AND array_length(genres, 1) <= 5)
    "movies_runtime_check" CHECK (runtime >= 0)
    "movies_year_check" CHECK (year >= 1888 AND year::double precision <= date_part('year'::text, now()))
```

#### 附加内容

##### 迁移到指定版本

查看schema_migrations表的另一种方法：如果您想查看您的数据库当前处于哪个迁移版本，您可以运行迁移工具的version命令，如下所示:

```shell
$ migrate -path=./migrations -database=$GREENLIGHT_DB_DSN version
2
```

你也可以使用goto命令向上或向下迁移到一个特定的版本:

```shell
$ greenlight migrate -path=./migrations -database=$GREENLIGHT_DB_DSN goto 1 
2/d add_movies_check_constrains (41.570125ms)
$ greenlight migrate -path=./migrations -database=$GREENLIGHT_DB_DSN version
1
```

##### 执行回退

可以使用down命令回滚特定版本的迁移。例如，要回滚最近的迁移，你需要运行:

```shell
 $ migrate -path=./migrations -database =$EXAMPLE_DSN down 1
```

就我个人而言，我通常更喜欢使用goto命令来执行回滚(因为它对目标版本更明确)，并保留使用down命令来回滚所有迁移，如下所示:

```shell
$ migrate -path=./migrations -database=$GREENLIGHT_DB_DSN down
Are you sure you want to apply all down migrations? [y/N]
y
Applying all down migrations
2/d add_movies_check_constrains (27.182625ms)
1/d create_movies_table (37.182834ms)
```

另一种方法是使用drop命令，它将从数据库中删除所有表，包括schema_migrations表——但是数据库本身以及已经创建的序列和枚举等其他内容将保留下来。因此，使用drop可能会使数据库处于混乱和未知的状态，如果您想要回滚所有内容，通常最好使用down命令。

##### 修改SQL迁移错误

理解在SQL迁移文件中出现语法错误时会发生什么是很重要的，因为迁移工具可能会导致一些令人费解对结果。

当运行有错误的迁移时，该错误后续所有SQL语句都会影响，然后迁移工具将退出并打印错误信息。类似于:

```shell
$ migrate -path=./migrations -database=$EXAMPLE_DSN up
1/u create_foo_table (36.6328ms)
2/u create_bar_table (71.835442ms)
error: migration failed: syntax error at end of input in line 0: CREATE TABLE (details: pq: syntax error at end of input)
```

如果迁移的失败文件包含多个SQL语句，那么可能是在遇到错误之前部分SQL操作已经生效。反过来，就迁移工具而言，这意味着数据库处于未知状态。

因此，schema_migrations字段中的version字段将包含失败迁移的编号，dirty字段将被设置为true。此时，如果你运行另一个迁移(即使是“向下”迁移)，你将会看到类似这样的错误消息:

```shell
 Dirty database version {X}. Fix and force version.
```

您需要做的是检查错误原因，并确定失败的迁移文件是否部分已经生效。如果是，那么您需要手动回滚部分生效的迁移。

完成之后，还必须将schema_migrations表中的版本号“强制”设置为正确的值。例如，要设置数据库版本号为1，你应该像这样使用force命令:

```shell
 $ migrate -path=./migrations -database=$EXAMPLE_DSN force 1
```

一旦您强制设置版本，数据库被认为是“干净的”，您应该能够再次运行迁移而没有任何问题。

#### 远程迁移文件

migrate工具还支持从远程源(包括Amazon S3和GitHub存储库)读取迁移文件。例如:

```shell
$ migrate -source="s3://<bucket>/<path>" -database=$EXAMPLE_DSN up
$ migrate -source="github://owner/repo/path#ref" -database=$EXAMPLE_DSN up
$ migrate -source="github://user:personal-access-token@owner/repo/path#ref" -database=$EXAMPLE_DSN up
```

关于此功能的更多信息和支持的远程sources的完整列表可以在[这里](https://github.com/golang-migrate/migrate#migration-sources)找到。

#### 应用启动时执行迁移

如果您愿意，也可以使用[golang-migrate/migrate](https://github.com/golang-migrate/migrate) Go包(不是命令行工具)在应用程序启动时自动执行数据库迁移。

在本书中我们不使用这种方法，所以如果您遵循本文，请不要更改您的代码。但大致上代码是这样的:

```go
package main

...

func main() {

   ...

    db, err := openDB(cfg)
    if err != nil {
        logger.Fatal(err, nil)
    }
    defer db.Close()

    logger.Info("database connection pool established", nil)

    migrationDriver, err := postgres.WithInstance(db, &postgres.Config{})
    if err != nil {
        logger.Fatal(err, nil)
    }

    migrator, err := migrate.NewWithDatabaseInstance("file:///path/to/your/migrations", "postgres", migrationDriver)
    if err != nil {
        logger.Fatal(err, nil)
    }

    err = migrator.up()
    if err != nil && err != migrator.ErrNoChange {
        logger.Fatal(err, nil)
    }

    logger.Printf("database migrations apply")
}
```

虽然这是可行的(看起来很有吸引力)，但是从长远来看，迁移的执行与应用程序源代码的紧密耦合可能会产生限制和问题。

[将数据库迁移与服务器启动分离](https://pythonspeed.com/articles/schema-migrations-server-startup/)这篇文章对此进行了很好的讨论，如果您感兴趣，我建议您阅读这篇文章。它以Python为中心，但不要因此而却步——同样的原则也适用于Go应用程序。
## SQL迁移

本书的下一部分我们将回到一些更具体的代码实践中，并一步步在greenlight数据库中创建一个movie表。

为此，我们只需使用psql工具然后执行必要的CREATE TABLE语句。但是，我们将探讨如何使用SQL迁移来创建表(更一般地说，管理整个项目中数据库模式更改)。

你将学习到：

* SQL迁移背后的原则以及它们为什么有用。
* 如何使用命令行工具以编程方式管理数据库模式的更改。
## 建立movie模型

在本章中，我们将为我们的movie数据库模型设置代码框架。

如果您不喜欢模型这个术语，那么您可以将其视为数据访问或存储层。但不管你喜欢怎么称呼，原理都是一样的——它封装了所有从PostgreSQL数据库中读取和写入movie数据的代码。

我们回到internal/data/movies.go文件，创建MovieModel结构体类型以及一些对movie数据库表执行CURD（create，read，update和delete）操作。

File：internal/data/movies.go

---

```go
package data

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"greenlight.alexedwards.net/internal/validator"
)

...

//定义MovieModel结构体，包含一个sql.DB连接池字段
type MovieModel struct {
	DB *sql.DB
}

//添加一个占位符方法，用于在movies表中插入一条新记录。
func (m MovieModel)Insert(movie *Movie) error {
	return nil
}

//Get 查询特定movie方法
func (m MovieModel)Get(id int64) (*Movie, error) {
	return nil, nil
}

//Update 更新数据库特定movie
func (m MovieModel)Update(movie *Movie) error {
	return nil
}

//Delete 删除数据库中特定movie
func (m MovieModel)Delete(id int64) error {
	return nil
}
```

另外，我们将MovieModel包装在一个Models结构中。这么做完全是可选的，但它的好处是可以为项目中所有需要存储数据库的模型集中管理。

创建internal/data/models.go文件，添加如下代码：

```shell
$ touch internal/data/models.go
```

File: internal/data/models.go

---

```go
package data

import (
	"database/sql"
	"errors"
)

//自定义错误ErrRecordNotFound常量。在Get()方法中如果查询的moive数据库中不存在就返回该错误。
var (
	ErrRecordNotFound = errors.New("record not found")
)

// Models结构体包装MovieModel。后面会将所有其他models都添加到这个结构体中。
type Models struct {
	Movies MovieModel
}

//为了便于使用，我们添加一个New()方法，它返回一个Models结构体，其中包含初始化的MovieModel。
func NewModels(db *sql.DB) Models {
	return Models{
		Movies: MovieModel{DB: db},
	}
}
```

接下来，编辑cmd/api/main.go文件，在main函数中初始化Models结构体实例，作为依赖项传递给我们的处理程序。像这样:

```go
package main

import (
	"context"
	"database/sql"
	"flag"
	"greenlight.alexedwards.net/internal/data"
	"os"
	"time"

	_ "github.com/lib/pq"
	"greenlight.alexedwards.net/internal/jsonlog"
)

...

//添加models字段
type application struct {
	config config
	logger *jsonlog.Logger
	models data.Models
}

/ 添加models字段，引入数据库模型为接口处理程序所用
app := &application {
	config: cfg,
	logger: logger,
	models: data.NewModels(db),
}

...
```

现在你可以重启应用程序。以上代码可以正常编译和运行。

```shell
$ go run ./cmd/api
2021/11/28 21:01:22 database connection pool established 
2021/11/28 21:01:22 starting development server on :4000
```

这个模式的优点之一是，从API处理程序的角度来看，在movie表上执行操作的代码非常清晰且可读。例如，我们可以简单地通过以下代码来执行Insert()方法:

```go
app.models.Movie.Insert(...)
```

通用结构体也很容易扩展。以后创建更多的数据库模型时，所要做的就是将它们添加到models结构体中，将自动在API处理程序中可用。

## 附加内容

#### Mocking models

通过对该模式进行小调整，还可以为单元测试而mock数据库模型。但是作为一个简单的例子，你可以创建一个类似于下面这样的MockMovieModel结构提:

File: internal/data/movies.go

---

```go
package main

...

type MockMovieModel struct {}

func (m MockMovieModel)Insert(movie *Movie) error {
	//Mock操作
	return nil
}

func (m MockMovieModel)Update(movie *Movie) error {
	//Mock操作
	return nil
}

func (m MockMovieModel)Delete(id int64) error {
	//Mock操作
	return nil
}

```

然后更新internal/data/models.go文件，如下所示：

File：internal/data/models.go

---

```go
package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
        //将Movies字段设置为接口类型，支持model和mockmodel赋值
	Movies interface {
		Insert(movie *Movie) error
		Get(id int64) (*Movie, error)
		Update(movie *Movie) error
		Delete(id int64) error
	}
}

...

//创建一个new函数，返回一个包含mock Models实例
func NewMockModels() Models {
	return Models{Movies: MockMovieModel{}}
}
```

只要在单元测试中用NewMockModels()来代替“真正的”NewModels()函数，就可以调用它。
## 创建新的movie接口

从数据库模型的Insert()方法开始，更新该方法为在movies表中创建一条新记录。具体来说，我们希望它执行以下SQL查询:

```sql
INSERT INTO movies (title, year, runtime, genres) 
VALUES ($1, $2, $3, $4)
RETURNING id, created_at, version
```

关于这个查询，有一些事情值得解释一下。

* 使用$N符号作为占位符，表示我们想要插入movies表中的数据参数。每次客户端向SQL数据库传递不可信的输入数据时，使用占位符参数来帮助防止SQL注入攻击，这非常重要，除非您有非常明确的理由不使用它们。
* 只要插入title，year，runtime和genres的值。movies表中的其余列将在插入时用系统生成的值填充——id将是一个自动递增的整数，created_at和version分别以当前时间和1来插入表中。
* 在查询的末尾，有一个returns子句。这是一个特定于postgresql的子句(它不是SQL标准)，您可以使用它从INSERT、UPDATE或DELETE语句所操作的任何记录中返回对应列的值。在这个查询中，我们使用它来返回系统生成的id、created_at和version值。

#### 执行SQL查询

在整个项目中，我们将坚持使用Go的[database/sql](https://golang.org/pkg/database/sql/)包来执行我们的数据库查询，而不是使用第三方ORM或其他工具。

通常使用Go的Exec()方法对数据库表执行INSERT语句。但是由于我们的SQL查询返回单行数据(多亏了returning子句)，我们在这里需要使用QueryRow()方法。

回到internal/data/movies.go文件，更新代码如下：

File：internal/data/movies.go

---

```go
package data

...

//添加一个占位符方法，用于在movies表中插入一条新记录。
func (m MovieModel) Insert(movie *Movie) error {
	//定义SQL查询，插入一条新的movie数据，并返回系统生成数据
	query := `
    	INSERT INTO movies (title, year, runtime, genres)
        VALUES ($1, $2, $3, $4)
		RETURNING id, create_at, version`

	//创建args切片包含占位符参数值。声明切片使得传入SQL中的值看起来更清晰
	args := []interface{}{movie.Title, movie.Year, movie.Runtime, pq.Array(movie.Genres)}

	//在连接池中使用QueryRow()方法执行SQL查询，传入args切片可变参数，并扫描生成的Id，create_at和version值到movie结构体中
	return m.DB.QueryRow(query, args...).Scan(&movie.ID, &movie.CreateAt, &movie.Version)
}
```

该代码非常简洁，但是有一些重要的地方需要说明。

因为Insert()方法签名以*Movie指针作为参数，所以当我们调用Scan()来读取系统生成的数据时，我们正在更新参数所指向的位置的值。实际上，我们的Insert()方法改变了传递给它的Movie结构体，并将系统生成的值添加到对应字段。

接下来要讨论的是我们在args切片中声明的占位符参数输入，如下所示:

```go
 args := []interface{}{movie.Title, movie.Year, movie.Runtime, pq.Array(movie.Genres)}
```

将参数存放在切片中不是必须的，正如我们在代码注释说明的，能够使代码更简洁。就个人而言，我通常对具有三个以上占位符参数的SQL查询都这样做。

另外，注意到切片中的最终值了吗？为了存储movie.Genres值(它是一个字符串切片)，在执行SQL查询之前，我们需要将它传递给pq.Array()适配器函数。这个函数作用就是对切片进行强制类型转换为可以直接插入数据库类型。

这里pq.Array()函数将[]string切片转换为pq.StringArray类型。因为pq.StringArray类型实现了driver.Valuer和sql.Scanner接口，它可以将原生字符串切片转换为PostgreSQL数据库可以理解的值并存储在text[]列中的值。

> 提示：你也可以在代码中使用pq.Array()函数对[]bool, []byte, []int32, []int64, []float32和[]float64进行转换。

#### 连接到API处理程序

下面激动人心时刻到了，我们把Insert()方法连接到createMovieHandler接口处理程序中，因此POST /v1/movies接口可以实现向数据库添加movie数据。如下所示：

File: cmd/api/movies.go

---

```go
package main

...

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string       `json:"title"`
		Year    int32        `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Genres  []string     `json:"genres"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	movie := &data.Movie{
		Title:   input.Title,
		Year:    input.Year,
		Runtime: input.Runtime,
		Genres:  input.Genres,
	}
	v := validator.New()
	if data.ValidateMovie(v, movie); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	//调用movie模型的Insert()方法，传入校验后的movie指针结构体。
	//将在数据库中创建movie条目，并用系统生成对信息更新结构体
	err = app.models.Movies.Insert(movie)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	//当发送HTTP响应时，我们希望包含Location头信息，使客户端知道哪个URL可以查询新建对资源。
	//可以创建一个空对http.header map然后用Set()方法添加Location头。
	//在URL中插入系统为movie生成的ID。
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/movies/%d", movie.ID))

	//使用201返回码以及movie数据作为响应body
	err = app.writeJSON(w, http.StatusCreated, envelope{"movie": movie}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
```

ok，我们来试试！

重启服务，然后打开第二个终端窗口，向POST /v1/movies端点发出如下请求:

```shell
$ BODY='{"title":"Moana","year":2016,"runtime":"107 mins", "genres":["animation","adventure"]}'
$ curl -i -d "$BODY" localhost:4000/v1/movies
HTTP/1.1 201 Created
Content-Type: application/json
Location: /v1/movies/1
Date: Sun, 28 Nov 2021 09:36:54 GMT
Content-Length: 140

{
        "movie": {
                "id": 1,
                "title": "Moana",
                "runtime": "107 mins",
                "genres": [
                        "animation",
                        "adventure"
                ],
                "Version": 1
        }
}
```

看起来很完美。我们可以看到，JSON响应包含movie的所有信息，包括系统生成的ID和版本号。响应还包括Location: /v1/movies/1头，指向URL可以查询到新添加到数据库的movie。

我们去数据库查看下，movie数据是否插入成功：

```shell
$  psql $GREENLIGHT_DB_DSN                         

psql (13.4)
Type "help" for help.

greenlight=> select * from movies;
 id |       create_at        | title | year | runtime |        genres         | version 
----+------------------------+-------+------+---------+-----------------------+---------
  1 | 2021-11-28 17:36:54+08 | Moana | 2016 |     107 | {animation,adventure} |       1
(1 row)
```

#### 创建其他数据库条目

让我们在系统中创建更多的记录，以帮助我们在构建过程中演示不同的功能。如果你跟随本书操作，请执行以下命令在数据库中创建更多的movie数据：

```shell
$ BODY='{"title":"Black Panther","year":2018,"runtime":"134 mins","genres":["action","adventure"]}'
$ curl -d "$BODY" localhost:4000/v1/movies
{
        "movie": {
                "id": 2,
                "title": "Black Panther",
                "runtime": "134 mins",
                "genres": [
                        "action",
                        "adventure"
                ],
                "Version": 1
        }
}
```

```shell
$ BODY='{"title":"Deadpool","year":2016, "runtime":"108 mins","genres":["action","comedy"]}'
$ curl -d "$BODY" localhost:4000/v1/movies                                      
{
        "movie": {
                "id": 3,
                "title": "Deadpool",
                "runtime": "108 mins",
                "genres": [
                        "action",
                        "comedy"
                ],
                "Version": 1
        }
}
```

```shell
$ BODY='{"title":"The Breakfast Club","year":1986, "runtime":"96 mins","genres":["drama"]}'
$ curl -d "$BODY" localhost:4000/v1/movies                                           
{
        "movie": {
                "id": 4,
                "title": "The Breakfast Club",
                "runtime": "96 mins",
                "genres": [
                        "drama"
                ],
                "Version": 1
        }
}
```

此时，您可能还需要查看PostgreSQL，以确认记录是否已正确创建。您应该看到movies表的内容现在看起来与此类似(在数组中包括适当的电影类型)。

```shell
$ psql $GREENLIGHT_DB_DSN

psql (13.4)
Type "help" for help.

greenlight=> select * from movies;
 id |       create_at        |       title        | year | runtime |        genres         | version 
----+------------------------+--------------------+------+---------+-----------------------+---------
  1 | 2021-11-28 17:36:54+08 | Moana              | 2016 |     107 | {animation,adventure} |       1
  2 | 2021-11-28 18:34:07+08 | Black Panther      | 2018 |     134 | {action,adventure}    |       1
  3 | 2021-11-28 18:37:13+08 | Deadpool           | 2016 |     108 | {action,comedy}       |       1
  4 | 2021-11-28 18:38:42+08 | The Breakfast Club | 1986 |      96 | {drama}               |       1
(4 rows)
```

## 附加内容

#### $N符号

PostgreSQL占位符参数$N符号的一个很好的特性是：您可以在SQL语句的多个位置使用相同的参数值。例如，像这样写代码是完全可以的:

```go
// 这个SQL语句两次使用$1参数，`123`这个值将在$1占位符中使用两次。
stmt := "UPDATE foo SET bar = $1 + $2 WHERE bar = $1"
err := db.Exec(stmt, 123, 456)
if err != nil { ...
}
```

#### 执行多条SQL语句

偶尔你可能会发现自己在一个数据库调用中需要执行多个SQL语句，就像这样:

```go
stmt := `
    UPDATE foo SET bar = true; 
    UPDATE foo SET baz = false;`
err := db.Exec(stmt) if err != nil {
...
}
```

pq驱动支持在一个调用包含多条SQL语句，只要不包含任何占位符即可。如果确实包含占位符参数，那么你将在运行时收到以下错误消息:

```shell
pq: cannot insert multiple commands into a prepared statement
```

要解决这个问题，您需要将语句拆分为单独的数据库调用，或者如果不可能，您可以在PostgreSQL中创建一个[自定义函数](https://www.postgresql.org/docs/current/xfunc-sql.html)，作为要运行的多个SQL语句的包装器。
## 查询Movie接口

现在我们来看看用于获取和显示特定movie数据的代码。
同样，我们将从数据库模型开始，首先将Get()方法更新为执行以下SQL查询:

```sql
SELECT id, create_at, title, year, runtime, genres, version
FROM moives
where id = $1
```

因为movies表使用id列作为它的主键，所以这个查询将只返回一个数据库行(或者根本不返回任何行)。因此，我们应该再次使用Go的QueryRow()方法来执行这个查询。
打开internal/data/movies.go文件，更新以下代码：

```go
func (m MovieModel) Get(id int64) (*Movie, error) {
	//PostgreSQL的bigserial类型用于movie的ID，将自增1，因此我们知道id不会小于1。
	//为了避免不必要的查询，这里加个判断。
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	//定义SQL查询语句用于查询movie
	query := `
		SELECT id, create_at, title, year, runtime, genres, version
        FROM movies
        where id = $1`
	var movie Movie
	err := m.DB.QueryRow(query, id).Scan(
		&movie.ID,
		&movie.CreateAt,
		&movie.Title,
		&movie.Year,
		&movie.Runtime,
		pq.Array(&movie.Genres),
		&movie.Version,
		)
	//错误处理，如果没找到对应的movie，Scan()将返回sql.ErrNoRows错误。
	//检查错误类型并返回自定义ErrRecordNotFound
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err

		}
	}
	//否则，返回Movie指针
	return &movie, nil
}
```

上面代码唯一值得注意的是，当扫描来自PostgreSQL text[]数组的类型数据时，我们需要再次使用pq.Array()函数。如果我们不使用这个适配器函数，将在运行时得到以下错误:

```shell
sql: Scan error on column index 5, name "genres": unsupported Scan, storing driver.Value type []uint8 into type *[]string
```

#### 更新对应的API处理程序

接下来要做的就是更新showMovieHandler可以调用movie模型的Get()方法。处理程序需要查看Get()函数是否返回ErrRecordNotFound错误，如果返回该错误，向客户端发送404 Not Found响应。否则，我们可以继续在JSON响应中呈现返回的Movie结构体。如下所示：

File：cmd/api/movies.go

---

```go
package main

...

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	//调用Get()方法获取数据库中对应id的movie信息。并使用errors.Is()方法检查返回错误类型
	//如果是未找到对应id的movie就返回404 Not Found响应
	movie, err := app.models.Movies.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
```

代码很简洁，多亏了前面准备的帮助函数。您可以通过重新启动API服务，并查找已经在数据库中创建的movie数据。例如:

```shell
$ curl -i localhost:4000/v1/movies/2
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 28 Nov 2021 12:44:06 GMT
Content-Length: 145

{
        "movie": {
                "id": 2,
                "title": "Black Panther",
                "runtime": "134 mins",
                "genres": [
                        "action",
                        "adventure"
                ],
                "Version": 1
        }
}

```

同样地，你也可以尝试用一个在数据库中不存在的ID发启请求。在那种情况下，你应该收到404 Not Found响应，像这样:

```shell
$ curl -i localhost:4000/v1/movies/42
HTTP/1.1 404 Not Found
Content-Type: application/json
Date: Sun, 28 Nov 2021 12:45:35 GMT
Content-Length: 59

{
        "error": "the requested  resource could not be found"
}
```

## 附加内容

#### 为什么不使用无符号整数作为电影ID?

在Get()方法的开头，我们用以下代码来检查id参数是否小于1:

```go
func (m MovieModel) Get(id int64) (*Movie, error) { 
	if id < 1 {
		return nil, ErrRecordNotFound 
	}
    ...
}
```

这可能想知道：如果电影ID不为负，在Go代码中为什么我们不使用unsigned uint64类型来存储ID，而使用int64?

有两个原因：

* 第一个原因是PostgreSQL没有无符号整数。为了避免溢出或其他兼容性问题，将Go和数据库整数类型对齐是明智的，因为PostgreSQL没有unsigned integer，这意味着我们应该避免在Go代码中使用uint*类型来读取或写入PostgreSQL的任何值。

相反，最好根据下表对齐整型:


| PostgreSQL 类型       | Go 类型                                             |
| ----------------------- | ----------------------------------------------------- |
| smallint, smallserial | int16(-32768 到 32767)                              |
| integer, serial       | int32(-2147483648 到 2147483647)                    |
| bigint, bigserial     | int64 (-9223372036854775808 to 9223372036854775807) |

还有另一个原因。Go的database/sql包实际上不支持任何大于9223372036854775807 (int64的最大值)的整型值。uint64可能大于这个值，这会导致Go产生类似这样的运行时错误:

```shell
sql: converting argument $1 type: uint64 values with high bit set are not supported
```

通过在Go代码中使用int64，我们消除了遇到这个错误的风险。
## 更新Movie

在本节，我们将继续构建我们的应用程序，并添加一个全新的API接口，允许客户端更新特定mvoie数据。


| Method  | URL Pattern        | Handler                | 操作                  |
| --------- | -------------------- | ------------------------ | ----------------------- |
| **PUT** | **/v1/movies/:id** | **updateMovieHandler** | **更新特定movie信息** |

更准确地说，我们将添加接口，以便客户端可以更新数据库中movie的title、year、runtime和genres内容。在我们的项目中，一旦创建id和created_at值，它们就不应该改变，并且版本值不应该由客户端控制，所以不允许编辑这些字段。

现在，我们将配置这个接口，使它对movie的值进行替换。这意味着客户端需要在其JSON请求体中为所有可编辑字段提供值……即使他们只想改变其中一个字段。

例如，如果一个客户端想要在我们的数据库中添加科幻电影《黑豹》，需要发送一个JSON请求体，如下所示:

```json
{
    "title": "Black Panther", 
    "year": 2018,
    "runtime": "134 mins", 
    "genres": [
        "action",
        "adventure",
        "sci-fi"
    ] 
}
```

#### 执行SQL查询

让我们再次开始数据库模型处理，并编辑Update()方法来执行下面SQL语句：

```sql
UPDATE movies
SET title = $1, year = $2, runtime = $3, genres = $4, version = version + 1 
WHERE id = $5
RETURNING version
```

注意到这里我们将版本值作为查询的一部分进行递增？最后我们用return子句返回这个新的加1的版本值。

和前面一样这个query返回一条数据，因此我们需要使用Go的QueryRow()方法执行。如果你跟随本书操作，返回到

internal/data/movies.go文件，然后在Update方法中添加如下代码：

File: internal/data/movies.go

---

```go
package data

...

func (m MovieModel) Update(movie *Movie) error {
	//声明SQL query更新记录并返回最新版本号
	query := `
		UPDATE movies
		set title = $1, year = $2, runtime = $3, genres = $4, version = version + 1
		WHERE id = $5
		RETURNING version`
	//创建args切片包含所有占位符参数值
	args := []interface{}{
		movie.Title,
		movie.Year,
		movie.Runtime,
		pq.Array(movie.Genres),
		movie.ID,
	}
	//使用QueryRow()方法执行，并以可变参数传入args切片，读取最新version值到movie结构体
	return m.DB.QueryRow(query, args...).Scan(&movie.Version)
}
```

需要强调的是：就像我们的Insert()方法一样Update()方法接受一个指向Movie结构体的指针作为参数，并再次原地修改它——这一次只使用新版本号更新。

#### 创建API处理程序（handler）

现在，我们在cmd/api/movies.go文件中，添加updateMovieHandler方法。


| Method  | URL Pattern        | Handler                | 操作                  |
| --------- | -------------------- | ------------------------ | ----------------------- |
| **PUT** | **/v1/movies/:id** | **updateMovieHandler** | **更新特定movie信息** |

这个处理程序的好处在于，我们已经为它打了所有的基础。这里的工作主要是将代码和已经编写的帮助函数串起来，以处理请求。

具体来说，我们需要:

1、 使用app.readIDParam()帮助函数从URL中提取电影ID。

2、使用我们在前一章中创建的Get()方法从数据库中获取相应的movie记录。

3、将包含更新movie数据的JSON请求体读入一个input结构。

4、将数据从input结构体复制到movie记录。

5、使用data.ValidateMovie()函数检查更新的movie记录各个字段是否有效。

6、调用Update()方法将新的movie信息存储到数据库中。

7、使用app.writeJSON()帮助函数将更新的movie数据写入JSON响应中。

下面开始写代码：
File: cmd/api/movies.go

---

```go
package main

...

func (app *application) updateMovieHandler(w http.ResponseWriter, r *http.Request) {
	//从URL中读取要更新的movie ID
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	//根据ID从数据库中读取旧movie信息，如果不存在就返回404 Not Found
	movie, err := app.models.Movies.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	//声明input结构体存放客户端发送来的数据
	var input struct {
		Title   string       `json:"title"`
		Year    int32        `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Genres  []string     `json:"genres"`
	}
	//读取JSON请求体到input结构体中
	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	//从请求体中将值拷贝到数据库movie记录对应字段
	movie.Title = input.Title
	movie.Year = input.Year
	movie.Runtime = input.Runtime
	movie.Genres = input.Genres
	//校验更新后到movie字段，如果校验失败返回422 Unprocessable Entity响应给客户端
	v := validator.New()
	if data.ValidateMovie(v, movie); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	//将检验后到movie传给Update()方法
	err = app.models.Movies.Update(movie)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	//将更新后到movie返回给客户端
	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
```

最后，为了完成这个任务，我们还需要更新应用程序路由以包含更新movie到API。像这样:
File：cmd/api/routers.go最后，为了完成这个任务，我们还需要更新应用程序路由以包含更新movie到API。像这样:
File：cmd/api/routers.go

---

```go
package main

...

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)
	//为"/v1/movies/:id"接口添加路由
	router.HandlerFunc(http.MethodPut, "/v1/movies/:id", app.updateMovieHandler)
	return app.recoverPanic(app.rateLimit(router))
}
```

#### 使用新的接口

现在，我们可以试试更新movie接口。

为了演示，让我们继续本章开始时给出的例子，并更新我们的记录，使《黑豹》包含科幻题材。提醒一下，目前的记录是这样的:

```shell
$  curl -i localhost:4000/v1/movies/2 
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 28 Nov 2021 13:48:43 GMT
Content-Length: 145

{
        "movie": {
                "id": 2,
                "title": "Black Panther",
                "runtime": "134 mins",
                "genres": [
                        "action",
                        "adventure"
                ],
                "Version": 1
        }
}
```

为了更新genres字段，我们可以执行以下API调用:

```shell
$ BODY='{"title":"Black Panther","year":2018,"runtime":"134 mins","genres":["sci-fi","action","adventure"]}'
$ curl -X PUT -d "$BODY" localhost:4000/v1/movies/2
{
        "movie": {
                "id": 2,
                "title": "Black Panther",
                "runtime": "134 mins",
                "genres": [
                        "sci-fi",
                        "action",
                        "adventure"
                ],
                "Version": 2
        }
}
```

这看起来很棒，我们可以从响应中看到，电影genres已经更新到包括“sci-fi”，版本号已经像我们预期的那样增加到2。

你也能够通过GET /v1/movies/2请求来验证更改是否被持久化，就像这样:

```shell
curl -i localhost:4000/v1/movies/2                                                                      
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 28 Nov 2021 13:52:52 GMT
Content-Length: 158

{
        "movie": {
                "id": 2,
                "title": "Black Panther",
                "runtime": "134 mins",
                "genres": [
                        "sci-fi",
                        "action",
                        "adventure"
                ],
                "Version": 2
        }
}
```
## 删除Movie接口

在本章中，我们将添加最终的CRUD接口，以便客户端可以从系统中删除特定的电影。


| Method                                              | URL                | Handler                | 动作                           |
| ----------------------------------------------------- | -------------------- | ------------------------ | -------------------------------- |
| GET                                                 | /v1/healthcheck    | createMovieHandler     | 显示应用程序运行状况和版本信息 |
| POST                                                | /v1/movies         | createMovieHandler     | 添加新的电影                   |
| GET                                                 | /v1/movies/:id     | showMovieHandler       | 根据id查询特定电影             |
| PUT                                                 | /v1/movies/:id     | updateMovieHandler     | 更新特定电影                   |
| **DELETE**                                          | **/v1/movies/:id** | **deleteMovieHandler** | **删除特定电影**               |
| 与API中的其他接口相比，Delete要实现的操作非常简单。 |                    |                        |                                |

* 如果数据库中存在一个具有URL中提供的id的movie记录，删除相应的记录并向客户端返回一条成功消息。
* 如果movie id不存在，向客户端返回一个404 Not Found响应。

在数据库中删除记录的SQL查询也很简单:

```sql
DELETE FROM movies
WHERE id = $1
```

在这种情况下，SQL查询不返回任何行，因此我们可以使用Go的Exec()方法来执行它。

Exec()的一个好处是它返回一个sql.Result对象，其中包含查询受影响的行数的信息。在我们的场景中，这是非常有用的信息。

* 如果受影响的行数是1，那么我们知道要删除的movie存在于表中，并且现在已经被删除了...所以我们可以向客户端发送成功消息。
* 相反，如果受影响的行数为0，我们知道对应id的movie不存在，向客户端返回404 Not Found响应。

#### 添加Delete接口

让我们继续更新数据库模型中的Delete()方法。本质上，我们希望它执行上面的SQL语句，并在受影响行数为0时返回ErrRecordNotFound错误。像这样:

File：internal/data/movies.go

---

```go
package data

...

func (m MovieModel) Delete(id int64) error {
	//如果movie ID小于1返回ErrRecordNotFound
	if id < 1 {
		return ErrRecordNotFound
	}
	//构造SQL query删除movie
	query := `
		DELETE FROM movies
		WHERE id = $1`
	//使用Exec()方法执行SQL语句，传入id变量作为占位符参数值。Exec()方法返回sql.Result对象
	result, err := m.DB.Exec(query, id)
	if err != nil {
		return err
	}
	//调用sql.Result对象的RowsAffected()方法，获取查询影响行数
	rowAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	//如果没有影响行数，可以推出数据库不存在对应id的记录，返回ErrRecordNotFound
	if rowAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}
```

完成之后，我们转到cmd/api/movies.go文件，添加一个新的deleteMovieHandler方法。在这个过程中，我们需要从请求URL中读取movie ID，调用刚刚创建的Delete()方法，然后根据Delete()的返回值向客户机返回适当的响应。

File: cmd/api/movies.go

---

```go
package main

...

func (app *application) deleteMovieHandler(w http.ResponseWriter, r *http.Request)  {
	//从URL中提取movie ID
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	//从数据库中删除movie，如果没有对应id的movie就返回404
	err = app.models.Movies.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	//返回200 Ok以及成功删除
	err = app.writeJSON(w, http.StatusOK, envelope{"message": "movie successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
```

> 注意：在这里您可能更喜欢发送一个空响应体和一个204 No Content状态码，而不是一个“电影已成功删除”消息。这取决于你的客户端是谁——如果他们是人，那么发送类似于上面的信息是一个很好的；如果它们是机器，那么一个204 No Content响应可能就足够了。

最后，需要在cmd/api/routes.go文件中添加删除movie的路由。

File：cmd/api/routes.go

---

```go
package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)
	router.HandlerFunc(http.MethodPut, "/v1/movies/:id", app.updateMovieHandler)
	//添加Delete /v1/movies/:id 接口路由
	router.HandlerFunc(http.MethodDelete, "/v1/movies/:id", app.deleteMovieHandler)
	return app.recoverPanic(app.rateLimit(router))
}
```

好了，重新启动API服务，通过从movie数据库中删除Deadpool来尝试下删除接口(如果您一直跟随本书操作，它的ID应该是3)。删除操作应该没有任何问题，您应该收到如下的确认消息:

```shell
$  curl -X DELETE http://localhost:4000/v1/movies/3
{
        "message": "movie successfully deleted"
}
```

如果您重复相同的请求来删除已经不存在的movie，应该会得到一个404 Not Found响应和错误消息。类似于:

```shell
$ curl -X DELETE http://localhost:4000/v1/movies/3
{
        "error": "the requested  resource could not be found"
}
```
## 数据库CRUD操作

本章节我们将重点介绍如何在我们的项目中创建、查询、更新和删除movies。

我们将在接下来的几节中取得重大进展，在本节结束时，我们将完成以下API接口并使其生效:


| Method                 | URL             | 动作                           |
| ------------------------ | ----------------- | -------------------------------- |
| GET                    | /v1/healthcheck | 显示应用程序运行状况和版本信息 |
| POST                   | /v1/movies      | 添加新的电影                   |
| GET                    | /v1/movies/:id  | 根据id查询特定电影             |
| PUT                    | /v1/movies/:id  | 更新特定电影                   |
| DELETE                 | /v1/movies/:id  | 删除特定电影                   |
| 本章将学习到以下内容： |                 |                                |

* 如何创建一个数据库模型，将所有执行SQL查询的逻辑与数据库隔离。
* 如何在API上下文中对特定资源实现四个基本CRUD(创建、查询、更新和删除)操作。
