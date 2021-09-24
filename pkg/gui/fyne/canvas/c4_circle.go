package main

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
)

//圆
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

/**
canvas.Circle定义由指定颜色填充的圆形。您还可以设置一个StrokeWidth，从而设置一个不同的StrokeColor，如本示例所示。

该圆将填充通过调用Resize（）或由其控制的布局指定的空间。由于该示例将圆圈设置为窗口内容，它将在基本填充（由主题控制）内调整大小以填充窗口。

所有这些都是基本类型，可以由我们的驱动程序呈现，而没有其他信息。接下来，我们将从Image开始看更复杂的类型。
*/
