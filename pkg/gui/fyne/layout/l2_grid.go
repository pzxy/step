package main

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Grid Layout")

	text1 := canvas.NewText("1", color.White)
	text2 := canvas.NewText("2", color.White)
	text3 := canvas.NewText("3", color.White)
	grid := fyne.NewContainerWithLayout(layout.NewGridLayout(2), // Grid
		text1, text2, text3)
	myWindow.SetContent(grid)
	myWindow.ShowAndRun()
}

/**
网格布局以具有固定列数的网格模式布置容器的元素。项目将填充一行，直到满足列数为止，此后将创建新行。垂直空间将在每一行对象之间平均分配。

您可以使用layout.NewGridLayout（cols）创建网格布局，其中cols是希望在每行中包含的项目（列）数。然后将此布局作为第一个参数传递给fyne.NewContainerWithLayout（...）。

如果您调整容器的大小，则每个单元的大小将相等，以共享可用空间。
*/
