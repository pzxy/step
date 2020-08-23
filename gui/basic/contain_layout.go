package main

import (
	"fyne.io/fyne/layout"
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	//"fyne.io/fyne/layout"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Container")

	text1 := canvas.NewText("Hello", color.White)
	text2 := canvas.NewText("There", color.White)
	text2.Move(fyne.NewPos(20, 0))
	text3 := canvas.NewText("World", color.White)
	//container := fyne.NewContainer(text1, text2, text3)
	container := fyne.NewContainerWithLayout(layout.NewGridLayout(3),
		text1, text2, text3) //表格布局

	myWindow.SetContent(container)
	myWindow.ShowAndRun()
}

/**
在前面的示例中，我们看到了如何将CanvasObject设置为Canvas的内容，但是仅显示一个视觉元素并不是很有用。为了显示多个项目，我们使用了容器类型。

由于fyne.Container也是fyne.CanvasObject，我们可以将其设置为fyne.Canvas的内容。在此示例中，我们创建3个文本对象，然后使用NewContainer（）函数将它们放置在容器中。由于没有布局设置，因此我们可以像使用text2.Move（）一样移动元素。

fyne.Layout实现了一种用于组织容器中项目的方法。通过取消注释此示例中的NewContainerWithLayout（）行，可以更改容器以使用具有2列的网格布局。运行此代码，然后尝试调整窗口大小，以查看布局如何自动配置窗口的内容。还要注意，布局代码将忽略text2的手动位置。

要了解更多信息，请跳至本教程的“布局”部分。
*/
