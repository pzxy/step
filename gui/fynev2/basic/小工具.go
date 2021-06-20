package main

/**
Afyne.Widget是一种特殊类型的容器，具有与之关联的附加逻辑。在小部件中，逻辑与其外观（也称为WidgetRenderer）是分开的。

小部件也是 的类型，CanvasObject因此我们可以将窗口的内容设置为单个小部件。看看我们如何widget.Entry在这个例子中创建一个新的 并将其设置为窗口的内容。

正如我们在前面的示例中看到的，您还可以使用 a 将多个对象添加到画布，Container并且可以使用一组小部件来开始构建图形应用程序界面。

接下来花一些时间了解其他“画布对象”。
*/

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Widget")

	myWindow.SetContent(widget.NewEntry())
	myWindow.ShowAndRun()
}
