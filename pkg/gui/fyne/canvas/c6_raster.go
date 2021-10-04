package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"image/color"
	"math/rand"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Raster")

	raster := canvas.NewRasterWithPixels( //n. 光栅，n. 像素；像素点
		func(_, _, w, h int) color.Color {
			return color.RGBA{uint8(rand.Intn(1)),
				uint8(rand.Intn(255)),
				uint8(rand.Intn(255)), 0xff}
		})
	// raster := canvas.NewRasterFromImage()
	w.SetContent(raster)
	w.Resize(fyne.NewSize(120, 100))
	w.ShowAndRun()
}

/**
canvas.Raster就像一幅图像，但是为屏幕上的每个像素精确地绘制了一个点。这意味着随着用户界面缩放或图像调整大小，将需要更多像素来填充空间。为此，我们使用本例中说明的Generator函数-将其用于返回每个像素的颜色。

生成器函数可以基于像素（如在本示例中，我们为每个像素生成新的随机颜色），也可以基于完整图像。生成完整图像（使用canvas.NewRaster（）效率更高，但有时直接控制像素更方便）。

如果像素数据存储在图像中，则可以通过NewRasterFromImage（）函数加载它，该函数将加载图像以在屏幕上完美显示像素。
*/
