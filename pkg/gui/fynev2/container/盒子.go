package main

/**
盒子
框小部件是一个简单的水平或垂直容器，它使用框布局来布置子组件。您可以传递要包含在container.NewHBox()或container.NewVBox()构造函数中的对象。

还可以在使用Add()（在现有内容之后）创建框小部件后将项目添加到框小部件或使用 删除项目Remove()。
*/

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Entry Widget")

	content := container.NewVBox(
		widget.NewLabel("The top row of the VBox"),
		container.NewHBox(
			widget.NewLabel("Label 1"),
			widget.NewLabel("Label 2"),
		),
	)

	content.Add(widget.NewButton("Add more items", func() {
		content.Add(widget.NewLabel("Added"))
	}))

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
