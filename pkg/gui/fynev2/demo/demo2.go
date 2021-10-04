package main

/**
与基于 Web 或命令行的应用程序相比，创建图形应用程序通常更复杂。Fyne 通过利用 Go 的出色设计改变了这一点，使构建漂亮的图形应用程序变得简单快捷。

为了让图形应用程序工作，我们需要创建一个窗口并要求应用程序运行。这样做将启动响应用户输入的事件处理代码，并在我们的代码运行时更新屏幕。

此示例创建一个新应用程序，其中包含一个标题为“Hello”的窗口。在这个窗口中，我们放置了一个包含文本“Hello Fyne！”的标签。

设置窗口内容后，我们将显示它并运行应用程序（Window.ShowAndRun()是Window.Show()&&的快捷方式App.Run()）。在调用 Run() 或 ShowAndRun() 之后，我们的应用程序将运行，该函数将在窗口关闭后返回。
*/

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello")

	myWindow.SetContent(widget.NewLabel("Hello Fyne!"))

	myWindow.ShowAndRun()
}

/**
Fyne 项目分为许多包，每个包提供不同类型的功能。它们如下：

fyne.io/fyne/v2
此导入提供所有 Fyne 代码通用的基本定义
包括数据类型和接口。
fyne.io/fyne/v2/app
应用程序包提供启动新应用程序的 API。
通常你只需要app.New()或app.NewWithID()。
fyne.io/fyne/v2/canvas
canvas 包提供了 Fyne 中的所有绘图 API。
完整的 Fyne 工具包由这些原始图形类型组成。
fyne.io/fyne/v2/container
容器包提供了用于布局和组织应用程序的容器。
fyne.io/fyne/v2/data/binding
绑定包包含将数据源绑定到小部件的方法。
fyne.io/fyne/v2/data/validation
验证包为验证小部件内的数据提供收费。
fyne.io/fyne/v2/dialog
对话框包包含确认、错误和文件保存/打开等对话框。
fyne.io/fyne/v2/layout
布局包提供了各种布局实现供使用
使用容器（在后面的教程中讨论）。
fyne.io/fyne/v2/storage
存储包提供存储访问和管理功能。
fyne.io/fyne/v2/test
使用测试中的工具可以更轻松地测试应用程序
包裹。
fyne.io/fyne/v2/widget
大多数图形应用程序是使用一组小部件创建的。
Fyne 中的所有小部件和交互元素都在这个包中。
*/
