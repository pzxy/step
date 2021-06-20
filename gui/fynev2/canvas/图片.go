package main

/**
图像
Acanvas.Image代表 Fyne 中可扩展的图像资源。它可以从资源（如示例所示）、图像文件、包含图像的 URI 位置、内存中的io.Reader或 Go加载image.Image。

默认图像填充模式canvas.ImageFillStretch将导致它填充指定的空间（通过Resize()或布局）。或者，您可以使用canvas.ImageFillContain来确保保持纵横比并且图像在边界内。除此之外，您可以使用canvas.ImageFillOriginal（如此处示例中使用的那样），以确保它的最小尺寸也等于原始图像尺寸的最小尺寸。

图像可以基于位图（如 PNG 和 JPEG）或基于矢量（如 SVG）。在可能的情况下，我们推荐可缩放的图像，因为它们会随着尺寸的变化继续呈现良好的效果。使用原始图像尺寸时要小心，因为它们在不同用户界面比例下的表现可能不完全符合预期。由于 Fyne 允许整个用户界面缩放 25px 图像文件可能与 25 高度 fyne 对象的高度不同。
*/

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Image")

	image := canvas.NewImageFromResource(theme.FyneLogo())
	// image := canvas.NewImageFromURI(uri)
	// image := canvas.NewImageFromImage(src)
	// image := canvas.NewImageFromReader(reader, name)
	// image := canvas.NewImageFromFile(fileName)
	image.FillMode = canvas.ImageFillOriginal
	w.SetContent(image)

	w.ShowAndRun()
}
