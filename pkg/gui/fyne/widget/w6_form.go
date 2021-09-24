package main

import (
	"log"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Form Widget")

	entry := widget.NewEntry()
	textArea := widget.NewMultiLineEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{"Entry", entry}},
		OnSubmit: func() { // optional, handle form submission
			log.Println("Form submitted:", entry.Text)
			log.Println("multiline:", textArea.Text)
			myWindow.Close()
		},
	}

	// we can also append items
	form.Append("Text", textArea)

	myWindow.SetContent(form)
	myWindow.ShowAndRun()
}

/**
表单小部件用于布置许多带有标签以及可选的取消和提交按钮的输入字段。它以最裸露的形式将标签对齐到每个输入小部件的左侧。通过设置OnCancel或OnSubmit，表单将添加一个按钮栏，并在适当时调用指定的处理程序。

可以使用widget.NewForm（...）传递widget.FormItems列表或使用示例中所示的＆widget.Form {}语法来创建窗口小部件。还有一个有用的Form.Append（label，widget）可用于替代语法。

在此示例中，我们创建两个条目，其中一个是用于保存值的“多行”（例如HTML TextArea）。有一个OnSubmit处理程序，该处理程序在关闭窗口并因此关闭应用程序之前先打印信息。
*/
