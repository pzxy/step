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
	myWindow := myApp.NewWindow("Center Layout")

	img := canvas.NewImageFromResource(theme.FyneLogo())
	img.FillMode = canvas.ImageFillOriginal
	text := canvas.NewText("Overlay", color.Black)
	content := fyne.NewContainerWithLayout(layout.NewCenterLayout(), //居于中心
		img, text)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

/**

layout.CenterLayout将其容器中的所有项目组织在可用空间的中心。将按照传递到容器的顺序绘制对象，最后一个绘制在最顶部。

中心布局使所有项目保持最小尺寸，如果您希望扩展项目以填充空间，请参见layout.MaxLayout。
*/
