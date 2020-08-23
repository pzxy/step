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
	myWindow := myApp.NewWindow("Grid Wrap Layout")

	text1 := canvas.NewText("1", color.White)
	text2 := canvas.NewText("2", color.White)
	text3 := canvas.NewText("3", color.White)
	//根据窗口大小自己适应布局
	grid := fyne.NewContainerWithLayout(layout.NewGridWrapLayout(fyne.NewSize(50, 50)), //  Wrap n. 弯曲，歪曲；偏见；乖戾
		text1, text2, text3)
	myWindow.SetContent(grid)

	// myWindow.Resize(fyne.NewSize(180, 75))
	myWindow.ShowAndRun()
}

/**
像以前的网格布局一样，网格环绕布局在网格图案中创建元素的排列。但是，此网格没有固定的列数，而是为每个单元格使用固定的大小，然后将内容流到显示项目所需的尽可能多的行中。

您可以使用layout.NewGridWrapLayout（size）创建网格环绕布局，其中size指定要应用于所有子元素的大小。然后将此布局作为第一个参数传递给fyne.NewContainerWithLayout（...）。列数和行数将根据容器的当前大小进行计算。

最初，网格环绕式布局只有一列，如果您调整其大小（如右侧的代码注释所示），它将重新排列子元素以填充空间。
*/
