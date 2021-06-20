package main

/**
入口
条目小部件用于用户输入简单文本内容。可以使用简单的widget.NewEntry() 构造函数创建条目。当您创建小部件时，请保留一个引用，以便您以后可以访问其Text字段。也可以使用OnChanged回调函数，在每次内容发生变化时得到通知。

条目小部件还可以验证输入到其中的文本输入。这可以通过将Validator 字段设置为 a来完成fyne.StringValidator。您还可以设置PlaceHolder 文本并将条目设置MultiLine为接受多行文本。

您还可以使用该NewPasswordEntry()功能创建密码条目（其中内容被隐藏）。
*/

import (
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Entry Widget")

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter text...")
	//好用
	content := container.NewVBox(input, widget.NewButton("Save", func() {
		log.Println("Content was:", input.Text)
	}))

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
