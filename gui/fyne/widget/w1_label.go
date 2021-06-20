package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Label Widget")

	content := widget.NewLabel("text") //标签

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

/**
小部件是Fyne应用程序GUI的主要组件，可以在基本的fyne.CanvasObject可以使用的地方使用它们。他们管理用户交互，但也将始终与当前主题匹配。

Label小部件是其中最简单的-它向用户显示文本。与canvas.Text不同，它可以处理一些简单的格式（例如\ n）。您可以通过调用widget.NewLabel（“ some text”）创建标签，结果可以分配给变量或直接传递到容器中。
*/
