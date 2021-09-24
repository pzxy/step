package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Image")

	//image := canvas.NewImageFromResource(theme.FyneLogo())
	image := canvas.NewImageFromFile("./sdq.png")
	// image := canvas.NewImageFromImage(src)
	//image.FillMode = canvas.ImageFillOriginal
	image.SetMinSize(fyne.NewSize(20, 20))
	w.SetContent(image)

	w.ShowAndRun()
}

/**
canvas.Image表示Fyne中的可伸缩图像资源。它可以从资源（如示例中所示），图像文件或Go image.Memory中的图像加载。

默认的图像填充模式是canvas.ImageFillStretch，它将导致它填充指定的空间（通过Resize（）或布局）。或者，您可以使用canvas.ImageFillContain来确保保持宽高比并且图像在范围之内。除此之外，您还可以使用canvas.ImageFillOriginal（在此处的示例中使用），以确保它的最小大小也等于原始图像的大小。

图像可以基于位图（如PNG和JPEG）或基于矢量（如SVG）。在可能的情况下，我们建议您使用可缩放的图像，因为它们会随着尺寸的变化而继续呈现良好的效果。使用原始图像尺寸时要小心，因为它们在不同的用户界面比例下可能无法完全达到预期的效果。由于Fyne允许整个用户界面缩放25px图像文件的高度可能与25高度fyne对象的高度不同。
*/
