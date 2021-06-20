package main

import "image/color"

/**
圆圈
canvas.Circle定义由指定颜色填充的圆形。您也可以设置 a StrokeWidth，因此 a 不同 StrokeColor，如本例所示。

圆圈将填充通过调用Resize()或由它控制的布局指定的空间。作为示例，将圆圈设置为窗口内容，它将在基本填充（由主题控制）内调整大小以填充窗口。

所有这些都是我们的驱动程序可以在没有附加信息的情况下呈现的基本类型。接下来，我们将看看更复杂的类型，以Image.
*/

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Circle")

	circle := canvas.NewCircle(color.White)
	circle.StrokeColor = color.Gray{0x99}
	circle.StrokeWidth = 5
	w.SetContent(circle)

	w.Resize(fyne.NewSize(100, 100))
	w.ShowAndRun()
}
