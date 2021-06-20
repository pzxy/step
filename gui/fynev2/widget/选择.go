package main

/**
选择
有多种小部件可供用户选择，其中包括复选框、单选组和选择弹出窗口。

在widget.Check提供了一个简单的是/否的选择和使用的字符串标签创建。这些小部件中的每一个还需要“更改” func(...)，其中参数是适当的类型。widget.NewCheck(..)因此需要一个string标签参数和一个func(bool)更改处理程序参数。您还可以使用该Checked字段来获取布尔值。

无线电小部件类似，但第一个参数是string代表每个选项的s切片。这次 change 函数需要一个string参数来返回当前选择的值。调用widget.NewRadioGroup(...) 以构建无线电组小部件，您可以稍后使用此引用来读取Selected字段，而不是使用更改回调。

选择小部件在构造函数签名中与单选小部件相同。调用widget.NewSelect(...)将显示一个按钮，该按钮在点击时显示一个弹出窗口，用户可以从中进行选择。这对于长的选项列表更好。
*/

import (
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Choice Widgets")

	check := widget.NewCheck("Optional", func(value bool) {
		log.Println("Check set to", value)
	})
	radio := widget.NewRadioGroup([]string{"Option 1", "Option 2"}, func(value string) {
		log.Println("Radio set to", value)
	})
	combo := widget.NewSelect([]string{"Option 1", "Option 2"}, func(value string) {
		log.Println("Select set to", value)
	})

	myWindow.SetContent(container.NewVBox(check, radio, combo))
	myWindow.ShowAndRun()
}
