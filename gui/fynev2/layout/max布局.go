package main

/**
最大布局
这layout.MaxLayout是最简单的布局，它将容器中的所有项目设置为与容器相同的大小。这在一般容器中并不常用，但在组合小部件时可能适用。

最大布局将容器扩展到至少最大项目的最小尺寸的大小。对象将按照传递到容器的顺序绘制，最后一个绘制在最顶部。

既然我们知道如何布置用户界面，我们将继续介绍Containers包，它可以简化布局并让您以更多方式布置对象。
*/

import (
	"image/color"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Max Layout")

	img := canvas.NewImageFromResource(theme.FyneLogo())
	text := canvas.NewText("Overlay", color.Black)
	content := container.New(layout.NewMaxLayout(), img, text)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
