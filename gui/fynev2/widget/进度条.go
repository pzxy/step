package main

/**
进度条
进度条小部件有两种形式，标准进度条显示Value已到达的用户，从Min到 Max。默认最小值为0.0，最大值默认为1.0。要使用默认值，只需调用widget.NewProgressBar(). 创建后，您可以设置该Value字段。

要设置自定义范围，您可以手动设置Min和Max字段。标签将始终显示完成百分比。

另一种形式的进度小部件是无限进度条。此版本仅通过从左到右再向后移动条形图的一部分来简单地显示某些活动正在进行中。您使用widget.NewProgressBarInfinite()它创建它，它会在显示后立即开始动画。
*/

import (
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
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

	myWindow.SetContent(container.NewVBox(progress, infinite))
	myWindow.ShowAndRun()
}
