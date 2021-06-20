package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Entry Widget")

	box := widget.NewVBox(widget.NewLabel("The top row of VBox"),
		widget.NewHBox(
			widget.NewLabel("Label 1"),
			widget.NewLabel("Label 2")))

	box.Append(widget.NewButton("Add more items", func() {
		box.Append(widget.NewLabel("Appended"))
	}))

	myWindow.SetContent(box)
	myWindow.ShowAndRun()
}

/**
Box小部件是一个简单的水平或垂直容器，使用Box布局来布置子组件。您可以传递对象以包含在widget.NewHBox（）或widget.NewVBox（）构造函数中。

也可以在使用Append（）（在现有内容之后添加）或Prepend（）（在内容之前添加）创建Box小部件之后将其添加到Box小部件中。
*/
