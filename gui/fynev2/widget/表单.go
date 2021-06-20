package main

/**
形式
表单小部件用于布置许多带有标签的输入字段以及可选的取消和提交按钮。它以最简单的形式将标签对齐到每个输入小部件的左侧。通过设置 OnCancel 或 OnSubmit，表单将添加一个按钮栏，并在适当的时候调用指定的处理程序。

可以widget.NewForm(...)通过传递widget.FormItems列表或使用&widget.Form{}示例中说明的 语法来创建小部件。还有一个有用的Form.Append(label, widget)可用于替代语法。

在此示例中，我们创建了两个条目，其中一个是用于保存值的“多行”（如 HTML TextArea）。有一个 OnSubmit 处理程序，它在关闭窗口（以及应用程序）之前打印信息。
*/

import (
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Form Widget")

	entry := widget.NewEntry()
	textArea := widget.NewMultiLineEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "Entry", Widget: entry}},
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
