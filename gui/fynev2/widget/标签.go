package main

/**
小部件是 Fyne 应用程序 GUI 的主要组件，它们可以用在任何基本的地方fyne.CanvasObject。他们管理用户交互，并将始终与当前主题相匹配。

Label 小部件是其中最简单的——它向用户显示文本。不像canvas.Text它可以处理一些简单的格式化（如\n）和包装（通过设置Wrapping字段）。您可以通过调用创建标签widget.NewLabel("some text")，结果可以分配给变量或直接传递到容器中。
*/

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Label Widget")

	content := widget.NewLabel("text")

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
