package main

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Max Layout")

	img := canvas.NewImageFromResource(theme.FyneLogo())
	text := canvas.NewText("Overlay", color.Black)
	content := fyne.NewContainerWithLayout(layout.NewMaxLayout(), //充满部件
		img, text)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

/**
layout.MaxLayout是最简单的布局，它将容器中的所有项目设置为与容器相同的大小。这在一般容器中通常不有用，但在组成小部件时可能很合适

最大布局会将容器扩展到至少最大项目最小尺寸的大小。将按照传递到容器的顺序绘制对象，最后一个绘制在最顶部。

现在，我们知道如何布置用户界面了，我们将继续使用可用于创建应用程序的小部件。
*/
