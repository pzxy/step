package main

/**
网格布局
网格布局以具有固定列数的网格模式布置容器的元素。项目将填充单行，直到满足列数，之后将创建一个新行。垂直空间将在每一行对象之间平均分配。

您创建一个网格布局，layout.NewGridLayout(cols)其中 cols 是您希望在每行中拥有的项目（列）数。然后将此布局作为第一个参数传递给 container.New(...)。

如果您调整容器的大小，则每个单元格将平均调整大小以共享可用空间。
*/

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"image/color"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Grid Layout")

	text1 := canvas.NewText("1", color.White)
	text2 := canvas.NewText("2", color.White)
	text3 := canvas.NewText("3", color.White)
	//每行2个
	grid := container.New(layout.NewGridLayout(2), text1, text2, text3)
	myWindow.SetContent(grid)
	myWindow.ShowAndRun()
}
