package main

import (
	"log"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Toolbar Widget")

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			log.Println("New document")
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentCutIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			log.Println("Display help")
		}),
	)

	content := fyne.NewContainerWithLayout(layout.NewBorderLayout(toolbar, nil, nil, nil),
		toolbar, widget.NewLabel("Content"))
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

/**
工具栏小部件使用图标表示一行操作按钮。 widget.NewToolbar（...）构造函数使用一个widget.ToolbarItem参数列表。工具栏项目的内置类型是动作，分隔符和分隔符。

最常用的项目是使用widget.NewToolbarItemAction（..）函数创建的操作。一个动作有两个参数，第一个是要绘制的图标资源，第二个是在点击时调用的func（）。这将创建一个标准的工具栏按钮。

您可以使用widget.NewToolbarSeparator（）在工具栏中的项目之间创建一个小的分隔线（通常是一条细的垂直线）。最后，您可以使用widget.NewToolbarSpacer（）在元素之间创建一个灵活的空间。这对于右对齐分隔符后列出的工具栏项最有用。

工具栏应始终位于内容区域的顶部，因此通常将其添加到fyne.Container中，使用layout.BorderLayout将其与其他内容对齐。
*/
