package main

import (
	"log"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Entry Widget")

	input := widget.NewPasswordEntry() //NewEntry
	input.SetPlaceHolder("Enter text...")

	content := widget.NewVBox(input, widget.NewButton("Save", func() {
		log.Println("Content was:", input.Text)
	}))

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

/**
条目小部件用于用户输入简单文本内容。可以使用简单的widget.NewEntry（）构造函数创建一个条目。创建窗口小部件时，请保留一个引用，以便以后可以访问其“文本”字段。每当内容更改时，也可以使用OnChanged回调函数进行通知。

也可以使用ReadOnly字段或调用SetReadOnly（true）将条目窗口小部件设置为仅读取内容。您还可以设置PlaceHolder文本，也可以将条目设置为MultiLine以接受多行文本。

您还可以使用NewPasswordEntry（）函数创建一个密码条目（其中的内容被遮盖）。
*/
