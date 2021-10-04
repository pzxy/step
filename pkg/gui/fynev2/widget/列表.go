package main

/**
列表
该List部件是该工具包的集合窗口小部件之一。这些小部件旨在在呈现大量数据时帮助构建真正高性能的界面。您还可以查看具有类似 API的表格和Tree小部件。由于这种设计，它们使用起来有点复杂。

该List回调函数的应用在需要时索要数据。有 3 个主要回调Length，CreateItem和UpdateItem。Length 回调（首先传递）是最简单的，它返回要呈现的数据中有多少项。其他与模板有关 - 如何创建、缓存和重用图形元素。

该CreateItem回调函数返回一个新的模板对象。当显示小部件时，这将与真实数据一起使用。在MinSize此对象将影响List最小尺寸。最后UpdateItem被调用以将数据项应用到缓存模板。使用它来设置准备显示的内容。
*/

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

var data = []string{"a", "string", "list"}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("List Widget")

	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i])
		})

	myWindow.SetContent(list)
	myWindow.ShowAndRun()
}
