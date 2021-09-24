package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello")
	myWindow.SetContent(widget.NewLabel("Hello"))

	go showAnother(myApp)
	myWindow.ShowAndRun()
}

func showAnother(a fyne.App) {
	time.Sleep(time.Second * 5)

	win := a.NewWindow("Shown later")
	win.SetContent(widget.NewLabel("5 seconds later"))
	win.Resize(fyne.NewSize(200, 200))
	win.Show()

	time.Sleep(time.Second * 2)
	win.Close()
}

/**
窗口处理
窗口是使用创建的App.NewWindow()，需要使用该Show()函数来显示。帮助方法ShowAndRun()onfyne.Window 允许您显示窗口并同时运行应用程序。

如果您希望显示第二个窗口，您必须只调用该Show() 函数。这在showAnother函数中进行了说明。

默认情况下，通过检查MinSize()函数（在后面的示例中将详细介绍），窗口将具有合适的大小以显示其内容。您可以通过调用该Window.Resize()函数设置更大的尺寸。

请注意，桌面环境可能具有导致窗口小于请求的限制。
*/
