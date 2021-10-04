package main

/**
canvas.Rectangle是 Fyne 中最简单的画布对象。它显示指定颜色的块。您还可以使用该FillColor字段设置颜色。

在这个例子中，矩形填充了窗口，因为它是唯一的内容元素。

其他fyne.CanvaObject类型有更多的配置，让我们 接下来看看canvas.Text。

*/

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Rectangle")

	rect := canvas.NewRectangle(color.White)
	w.SetContent(rect)

	w.Resize(fyne.NewSize(150, 100))
	w.ShowAndRun()
}
