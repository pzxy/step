package main

/**
在前面的示例中，我们看到了如何将 a 设置为 aCanvasObject的内容Canvas，但仅显示一个视觉元素并不是很有用。要显示多个项目，我们使用Container类型。

由于fyne.Container也是 a fyne.CanvasObject，我们可以将其设置为 a 的内容fyne.Canvas。

在本例中，我们创建了 3 个文本对象，然后使用该container.NewWithoutLayout()函数将它们放置在一个容器中。
由于没有布局集，我们可以像使用text2.Move().

Afyne.Layout实现了一种在容器内组织项目的方法。
通过取消注释container.New()本示例中的行，您可以更改容器以使用具有 2 列的网格布局。
运行此代码并尝试调整窗口大小以查看布局如何自动配置窗口内容。
另请注意，text2布局代码会忽略的手动位置。

要了解更多信息，请跳转到本次游览的Layout和Container部分。
*/

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	//"fyne.io/fyne/v2/layout"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Container")

	text1 := canvas.NewText("Hello", color.White)
	text2 := canvas.NewText("There", color.White)
	text2.Move(fyne.NewPos(40, 40))
	text3 := canvas.NewText("World", color.White)
	//容器也是Canvas
	content := container.NewWithoutLayout(text1, text2, text3)
	// content := container.New(layout.NewGridLayout(2), text1, text2, text3)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
