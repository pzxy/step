package main

/**
光栅
这canvas.Raster就像一张图像，但为屏幕上的每个像素精确绘制一个点。这意味着随着用户界面缩放或图像调整大小，将需要更多像素来填充空间。为此，我们使用Generator本示例中所示的函数 - 它将用于返回每个像素的颜色。

生成器函数可以是基于像素的（在这个例子中，我们为每个像素生成一个新的随机颜色）或基于完整图像。生成完整图像（使用canvas.NewRaster()）更有效，但有时直接控制像素更方便。

如果您的像素数据存储在图像中，您可以通过NewRasterFromImage()加载图像以在屏幕上完美显示像素的函数加载它 。
*/

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"image/color"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Gradient")

	gradient := canvas.NewHorizontalGradient(color.White, color.Transparent)
	//gradient := canvas.NewRadialGradient(color.White, color.Transparent)
	w.SetContent(gradient)

	w.Resize(fyne.NewSize(100, 100))
	w.ShowAndRun()
}
