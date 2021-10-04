package main

/**
所述canvas.Line对象绘制从一个线Position1（默认为顶部，左侧）到Position2（默认为底部，右）。您可以指定其颜色并可以改变笔画宽度，否则默认为1。

可以使用Position1或Position2 字段或使用Move()和Resize()函数来操作线位置。例如，宽度为 0 的区域将显示一条垂直线，而高度为 0 的区域将显示为水平线。

线条通常用于自定义布局或手动控制。与文本不同的是，它们没有自然（最小）大小，但可以在复杂的布局中产生很好的效果。
*/

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Line")

	line := canvas.NewLine(color.White)
	line.StrokeWidth = 5
	w.SetContent(line)

	w.Resize(fyne.NewSize(100, 100))
	w.ShowAndRun()
}
