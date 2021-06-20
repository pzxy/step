package main

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Rectangle")

	rect := canvas.NewRectangle(color.White)
	w.SetContent(rect)

	w.Resize(fyne.NewSize(220, 500))
	w.ShowAndRun()
}

/**
canvas.Rectangle是Fyne中最简单的canvas对象。它显示指定颜色的块。您也可以使用FillColor字段设置颜色。

在此示例中，矩形填充了窗口，因为它是唯一的内容元素。

其他fyne.CanvaObject类型具有更多配置，接下来让我们看一下canvas.Text。
*/
