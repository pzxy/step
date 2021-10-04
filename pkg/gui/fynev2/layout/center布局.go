package main

/**
中心布局
layout.CenterLayout将其容器中的所有项目组织到可用空间的中心。对象将按照传递到容器的顺序绘制，最后一个绘制在最顶部。

中心布局使所有项目保持最小尺寸，如果您希望扩展项目以填充空间，请参阅 layout.MaxLayout。
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
	myWindow := myApp.NewWindow("Center Layout")

	img := canvas.NewImageFromResource(theme.FyneLogo())
	img.FillMode = canvas.ImageFillOriginal
	text := canvas.NewText("Overlay", color.Black)
	//顾名思义，放到中心
	content := container.New(layout.NewCenterLayout(), img, text)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
