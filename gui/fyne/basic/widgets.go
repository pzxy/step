package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Widget")
	myWindow.SetContent(widget.NewEntry())
	myWindow.ShowAndRun()
}

/**
fyne.Widget是一种特殊的容器，具有与之关联的其他逻辑。在小部件中，逻辑与其外观不同（也称为WidgetRenderer）。

小部件也是CanvasObject的类型，因此我们可以将窗口的内容设置为单个小部件。请参阅我们如何创建新的小部件。在本示例中，将其输入并设置为窗口的内容。

正如我们在前面的示例中看到的，您还可以使用Container将多个对象添加到画布，并且可以使用一组小部件来完成相同操作，以开始建立图形应用程序界面。
*/
