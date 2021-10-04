package main

/**
表单布局
这layout.FormLayout就像一个 2 列网格布局， 但经过调整以在应用程序中布局表单。每个项目的高度将是每行中两个最小高度中的较大者。左侧项目的宽度将是第一列中所有项目的最大最小宽度，而每行中的第二个项目将扩展以填充空间。

这种布局更常用于widget.Form（用于验证、提交和取消按钮等），但它也可以直接用于layout.NewFormLayout()传递给container.New(...).
*/

import (
	"image/color"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Form Layout")

	label1 := canvas.NewText("Label 1", color.Black)
	value1 := canvas.NewText("Value", color.White)
	label2 := canvas.NewText("Label 2", color.Black)
	value2 := canvas.NewText("Something", color.White)
	//都不会动，不会对聚在一起
	grid := container.New(layout.NewFormLayout(), label1, value1, label2, value2)
	myWindow.SetContent(grid)
	myWindow.ShowAndRun()
}
