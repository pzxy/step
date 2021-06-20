package main

/**
按钮
按钮小部件可以包含文本、图标或两者，构造函数是widget.NewButton()和widget.NewButtonWithIcon()。要创建文本按钮，只有 2 个参数，string内容和func()点击按钮时将调用的 0 参数。请参阅示例以了解如何创建它。

带有图标的按钮构造函数包含一个附加参数，该参数fyne.Resource包含图标数据。theme包中的内置图标都适当地适应了主题的变化。如果您将自己的图像作为资源加载，则可以传入自己的图像 - 诸如此类的fyne.LoadResourceFromPath()帮助程序可能会提供帮助，但建议在可能的情况下捆绑资源。

要创建一个只有图标的按钮，您应该将“”作为标签参数传递给widget.NewButtonWithIcon()。
*/

import (
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	//"fyne.io/fyne/v2/theme"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Button Widget")

	content := widget.NewButton("click me", func() {
		log.Println("tapped")
	})

	//content := widget.NewButtonWithIcon("Home", theme.HomeIcon(), func() {
	//	log.Println("tapped home")
	//})

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
