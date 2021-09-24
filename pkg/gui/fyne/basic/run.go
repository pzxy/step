package main

import (
	"fmt"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello")
	myWindow.SetContent(widget.NewLabel("Hello"))
	myWindow.Show()
	myApp.Run()
	tidyUp()
}

func tidyUp() { //不会输出
	fmt.Println("Exited")
}

/**
为了使GUI应用程序正常工作，它需要运行一个事件循环（有时称为运行循环），以处理用户交互和绘制事件。在Fyne中，这是使用App.Run（）或Window.ShowAndRun（）函数启动的。其中之一必须在main（）函数的安装代码结尾处调用。

一个应用程序应该只有一个运行循环，因此您应该只在代码中调用一次Run（）。再次调用它会导致错误。

可以通过调用App.Quit（）直接退出应用程序，但这仅应响应用户事件而调用，以避免出现意外情况。另请参见在应用程序退出之前，不会调用Run（）之后执行的函数。
*/
