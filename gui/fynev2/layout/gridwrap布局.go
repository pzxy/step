package main

/**
网格环绕布局
与之前的网格布局一样，网格环绕布局以网格模式创建元素排列。但是，此网格没有固定的列数，而是为每个单元格使用固定大小，然后将内容流动到显示项目所需的行数。

您可以使用layout.NewGridWrapLayout(size) 其中的 size 指定要应用于所有子元素的大小来创建网格环绕布局。然后将此布局作为第一个参数传递给 container.New(...)。列数和行数将根据容器的当前大小计算。

最初的网格环绕布局只有一列，如果您调整它的大小（如右侧的代码注释所示），它将重新排列子元素以填充空间。
*/

import (
	"image/color"

	"fyne.io/fyne/v2"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Grid Wrap Layout")

	text1 := canvas.NewText("1", color.White)
	text2 := canvas.NewText("2", color.White)
	text3 := canvas.NewText("3", color.White)
	//就是根据窗口大小，自动调整123的位置。
	grid := container.New(layout.NewGridWrapLayout(fyne.NewSize(50, 50)),
		text1, text2, text3)
	myWindow.SetContent(grid)

	// myWindow.Resize(fyne.NewSize(180, 75))
	myWindow.ShowAndRun()
}
