package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"image/color"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Gradient")
	//Horizontal adj. 水平的, 与地平线平行的
	// Gradient n. 梯度；倾斜度；坡度，adj. 倾斜的；步行的
	//	gradient := canvas.NewHorizontalGradient(color.White, color.Transparent)
	//Radial adj. 辐射状的;放射式的;径向的;星形的
	gradient := canvas.NewRadialGradient(color.White, color.Transparent)
	w.SetContent(gradient)

	w.Resize(fyne.NewSize(1000, 1000))
	w.ShowAndRun()
}

/**

最后一个画布基本类型是Gradient，可作为canvas.LinearGradient和canvas.RadialGradient使用，用于绘制从一种颜色到另一种颜色的各种图案的渐变。您可以使用NewHorizo​​ntalGradient（），NewVerticalGradient（）或NewRadialGradient（）创建渐变。

要创建渐变，您需要开始和结束颜色-中间的每种颜色都是由画布计算的。在此示例中，我们使用color.Transparent来显示渐变（或任何其他类型）如何使用alpha值对后面的内容进行半透明。

这就是我们在Fyne中的画布元素之旅。接下来，花一些时间来学习可用于布置界面元素的布局。

*/
