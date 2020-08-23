package main

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
)

//
func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Form Layout")

	label1 := canvas.NewText("Label 1", color.Black)
	value1 := canvas.NewText("Value", color.White)
	label2 := canvas.NewText("Label 2", color.Black)
	value2 := canvas.NewText("Something", color.White)
	grid := fyne.NewContainerWithLayout(layout.NewFormLayout(), //列网格布局，但为应用程序中的表单布局进行了调整
		label1, value1, label2, value2)
	myWindow.SetContent(grid)
	myWindow.ShowAndRun()
}

/**

layout.FormLayout类似于2列网格布局，但为应用程序中的表单布局进行了调整。每个项目的高度将是每行中两个最小高度中的较大者。左侧项目的宽度将是第一列中所有项目的最大最小宽度，而每行中的第二个项目将扩展以填充空间。

此布局通常在widget.Form中使用，但可以直接与传递给fyne.NewContainerWithLayout（...）第一个参数的layout.NewFormLayout（）一起使用。
*/
