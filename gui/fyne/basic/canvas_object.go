package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/theme"
	"image/color"
	"time"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Canvas")
	myCanvas := myWindow.Canvas()

	text := canvas.NewText("Text", color.White)
	text.TextStyle.Bold = true
	myCanvas.SetContent(text)
	go changeContent(myCanvas)

	myWindow.Resize(fyne.NewSize(100, 100))
	myWindow.ShowAndRun()
}

func changeContent(c fyne.Canvas) {
	time.Sleep(time.Second * 2)

	c.SetContent(canvas.NewRectangle(color.Black))

	time.Sleep(time.Second * 2)
	c.SetContent(canvas.NewLine(color.Gray{0x66}))

	time.Sleep(time.Second * 2)
	circle := canvas.NewCircle(color.White)
	circle.StrokeWidth = 4
	circle.StrokeColor = color.RGBA{0xff, 0x33, 0x33, 0xff}
	c.SetContent(circle)

	time.Sleep(time.Second * 2)
	c.SetContent(canvas.NewImageFromResource(theme.FyneLogo()))
}

/**
在Fyne中，“画布”是在其中绘制应用程序的区域。每个窗口都有一个画布，您可以使用Window.Canvas（）访问它，但是通常您会在Window上找到避免访问画布的函数。

Fyne中可以绘制的所有内容都是CanvasObject的一种。此处的示例打开一个新窗口，然后通过设置窗口画布的内容来显示不同类型的原始图形元素。可以通过多种方式自定义每种类型的对象，如文本和圆圈示例所示。

除了使用Canvas.SetContent（）更改显示的内容外，还可以更改显示的内容。例如，如果更改矩形的颜色，则可以使用canvas.Refresh（rect）请求刷新此现有组件。
*/
