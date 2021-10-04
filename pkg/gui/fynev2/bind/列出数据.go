package main

/**
列出数据
为了演示如何连接更复杂的类型，我们将查看List小部件以及数据绑定如何使其更易于使用。首先我们创建一个StringList 数据绑定，它是一个String数据类型列表。一旦我们有了列表类型的数据，我们就可以将其连接到标准List小部件。为此，我们使用 widget.NewListWithData构造函数，就像其他小部件一样。

将此代码与列表游览进行比较 您将看到 2 个主要变化，第一个是我们将数据类型作为第一个参数而不是长度回调函数传递。第二个变化是最后一个参数，我们的UpdateItem回调。修订版采用一个binding.DataItem值而不是widget.ListIndexID。使用此回调结构时，我们应该Bind 调用模板标签小部件而不是调用SetText. 这意味着如果数据源中的任何字符串发生更改，表中每个受影响的行都将刷新。

在我们的演示代码中有一个“Append” Button，当点击它时，它将向数据源附加一个新值。这样做将自动触发数据更改处理程序并展开List小部件以显示新数据。
*/

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

//按了以后数据会出来，适合打印日志。
func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("List Data")

	data := binding.BindStringList(
		&[]string{"Item 1", "Item 2", "Item 3"},
	)

	list := widget.NewListWithData(data,
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		})

	add := widget.NewButton("Append", func() {
		val := fmt.Sprintf("Item %d", data.Length()+1)
		data.Append(val)
	})
	myWindow.SetContent(container.NewBorder(nil, add, nil, nil, list))
	myWindow.ShowAndRun()
}
