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
	myWindow := myApp.NewWindow("Border Layout")

	top := canvas.NewText("top bar", color.White)
	left := canvas.NewText("left", color.White)
	middle := canvas.NewText("content", color.White)
	//Border n. 边, 边缘   会虽适应，但是不会让布局换行。
	content := fyne.NewContainerWithLayout(layout.NewBorderLayout(top, nil, left, nil),
		top, left, middle)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

/**

边框布局可能是最广泛用于构建用户界面的布局，因为它允许围绕中心元素放置项目，该元素将扩展以填充空间。要创建边框布局，您需要将应放置在边框位置的fyne.CanvasObjects传递给布局（以及通常的容器）。此语法与其他布局略有不同，但基本上只是layout.NewBorderLayout（顶部，底部，左侧，右侧），如右侧示例中所示。

传递到容器的任何未出现在特定边界位置的物品都将放置在中央区域，并将展开以填充可用空间。您还可以将nil传递给您希望保留为空的border参数。

请注意，位于中心的所有项目都会展开以填充空间（就像它们位于layout.MaxLayout容器中一样）。要自己管理区域，您可以创建一个新的fyne.Container并使用所需的任何布局。
*/
