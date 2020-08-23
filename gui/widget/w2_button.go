package main

import (
	"fyne.io/fyne/theme"
	"log"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	//"fyne.io/fyne/theme"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Button Widget")

	//content := widget.NewButton("click me", func() {
	//	log.Println("tapped")
	//})
	content := widget.NewButtonWithIcon("Home", theme.CancelIcon(), func() {
		log.Println("tapped home")
	})

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

/**

按钮小部件可以包含文本，图标或两者都包含，构造函数为widget.NewButton（）和widget.NewButtonWithIcon（）。要创建文本按钮，只有两个参数，即字符串内容和一个0参数func（），在点击按钮时将调用它们。请参阅示例以了解如何创建该示例。

带图标的按钮构造函数包括一个附加参数，即fyne.Resource，其中包含图标数据。主题包中的内置图标都可以适当地适应主题的更改。如果将图像作为资源加载，则可以传递自己的图像-可以使用fyne.LoadResourceFromPath（）之类的帮助程序，尽管建议尽可能捆绑资源。

要创建仅包含图标的按钮，应将“”作为标签参数传递给widget.NewButtonWithIcon（）。
*/
