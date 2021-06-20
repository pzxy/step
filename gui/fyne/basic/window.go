package main

import (
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
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
	win.Hide()
}

/**
Windows是使用App.NewWindow（）创建的，需要使用Show（）函数显示。 fyne.Window上的辅助方法ShowAndRun（）允许您显示窗口并同时运行应用程序。

如果希望显示第二个窗口，则只能调用Show（）函数。在showAnother函数中对此进行了说明。

默认情况下，通过检查MinSize（）函数，窗口将具有合适的大小以显示其内容（在后面的示例中有更多介绍）。您可以通过调用Window.Resize（）函数来设置更大的尺寸。

请注意，桌面环境可能具有导致窗口小于要求的限制。
*/
