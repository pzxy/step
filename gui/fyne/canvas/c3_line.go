package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"image/color"
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

/**
canvas.Line对象从Position1（默认为左上）到Position2（默认为右下）绘制一条线。您可以指定其颜色，并可以更改笔划宽度，否则默认为1。

可以使用Position1或Position2字段或使用Move（）和Resize（）函数来操纵行位置。例如，宽度为0的区域将显示一条垂直线，而高度为0的区域将显示水平线。

线通常用于自定义布局或手动控制。与文本不同，它们没有自然的（最小）大小，但可以在复杂的布局中发挥很大的作用。
*/
