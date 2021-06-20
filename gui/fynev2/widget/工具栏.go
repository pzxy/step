package main

/**
工具栏
工具栏小部件创建一行操作按钮，使用图标来表示每个按钮。该widget.NewToolbar(...) 构造函数接受列表widget.ToolbarItem 参数。工具栏项目的内置类型是操作、分隔符和间隔符。

最常用的项目是使用该widget.NewToolbarItemAction(..)函数创建的操作 。一个动作有两个参数，第一个是要绘制的图标资源，第二个是func()点击时要调用的资源。这将创建一个标准的工具栏按钮。

您可以使用widget.NewToolbarSeparator()在工具栏中的项目之间创建一个小的分隔线（通常是一条细的垂直线）。最后，您可以使用widget.NewToolbarSpacer() 在元素之间创建灵活的空间。这对于右对齐在间隔符之后列出的工具栏项目最有用。

工具栏应始终位于内容区域的顶部，因此fyne.Container使用 layout.BorderLayout将其添加到 将其与其他内容对齐是正常的。
*/

import (
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
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
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {

		}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			log.Println("Display help")
		}),
	)

	content := container.NewBorder(toolbar, nil, nil, nil, widget.NewLabel("Content"))
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
