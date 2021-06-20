package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello")
	myWindow.SetContent(widget.NewLabel("Hello"))

	myWindow.Show()
	myApp.Run()
	tidyUp()
}

func tidyUp() {
	fmt.Println("Exited")
}

/**
应用程序和运行循环
要使 GUI 应用程序正常工作，它需要运行一个事件循环（有时称为 runloop）来处理用户交互和绘图事件。在 Fyne 中，这是使用App.Run() orWindow.ShowAndRun()函数开始的。必须从main()函数中设置代码的末尾调用其中之一。

一个应用程序应该只有一个 runloop，所以你应该只Run()在你的代码中调用一次。第二次调用它会导致错误。

对于桌面运行时，可以通过调用直接退出App.Quit() 应用程序（移动应用程序不支持此功能） - 通常在开发人员代码中不需要。一旦所有窗口都关闭，应用程序也将退出。另请参阅Run()在应用程序退出之前不会调用之后执行的函数。
*/
