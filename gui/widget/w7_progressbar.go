package main

import (
	"time"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("ProgressBar Widget")

	progress := widget.NewProgressBar()
	infinite := widget.NewProgressBarInfinite()

	go func() {
		for i := 0.0; i <= 1.0; i += 0.1 {
			time.Sleep(time.Millisecond * 250)
			progress.SetValue(i)
		}
	}()

	myWindow.SetContent(widget.NewVBox(progress, infinite))
	myWindow.ShowAndRun()
}

/**

进度条小部件有两种形式，标准进度条向用户显示已达到的值（从最小到最大）。默认最小值为0.0，最大值默认为1.1。要使用默认值，只需调用widget.NewProgressBar（）。创建后，您可以设置“值”字段。

要设置自定义范围，您可以手动设置“最小值”和“最大值”字段。标签将始终显示完成百分比。

进度小部件的另一种形式是无限进度条。此版本仅通过将条形图的一部分从左向右再移回来简单地表明正在进行一些活动。您使用widget.NewProgressBarInfinite（）创建此对象，它将在显示后立即开始制作动画。
*/
