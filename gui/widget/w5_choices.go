package main

import (
	"log"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Choice Widgets")

	check := widget.NewCheck("Optional", func(value bool) {
		log.Println("Check set to", value)
	})
	radio := widget.NewRadio([]string{"Option 1", "Option 2"}, func(value string) {
		log.Println("Radio set to", value)
	})
	combo := widget.NewSelect([]string{"Option 1", "Option 2"}, func(value string) {
		log.Println("Select set to", value)
	})

	myWindow.SetContent(widget.NewVBox(check, radio, combo))
	myWindow.ShowAndRun()
}

/**

有各种小部件可用来为用户提供选择，其中包括一个复选框，单选组和选择弹出窗口。

widget.Check提供简单的是/否选择，并使用字符串标签创建。这些小部件中的每个小部件都还带有一个“已更改”的func（...），其中参数的类型正确。 widget.NewCheck（..）因此将标签的字符串参数和更改处理程序的func（bool）参数作为参数。您也可以使用Checked字段来获取布尔值。

单选小部件与此类似，但是第一个参数是代表每个选项的字符串切片。更改函数这次需要一个字符串参数来返回当前选择的值。调用widget.NewRadio（...）构造单选控件，您以后可以使用此引用来读取Selected字段，而不是使用change回调。

选择窗口小部件在构造函数签名中与单选窗口小部件相同。调用widget.NewSelect（...）时将显示一个按钮，当点击该按钮时会显示一个弹出窗口，用户可以从中进行选择。这对于较长的选项列表更好。
*/
