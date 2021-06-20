package main

/**
桌子
该Table部件是像List小部件（另一个工具包的集合小部件）与二维指数。像List这样的设计，以帮助建立真正的高性能接口时，大量的数据正在呈现。因此，小部件不是在所有数据都嵌入的情况下创建的，而是在需要时调用数据源。

该Table回调函数的应用在需要时索要数据。有 3 个主要回调Length，CreateCell和UpdateCell。Length 回调（首先传递）是最简单的，它返回要呈现的数据中有多少项，它返回的两个整数表示行数和列数。另外两个与内容模板有关。

该CreateItem回调函数返回一个新的模板对象，就像清单。不同之处在于MinSize将定义每个单元格的标准大小和表格的最小大小（它至少显示一个单元格）。如前所述，UpdateItem调用 将数据应用到单元格模板。传入的索引是相同的(row, col)int 对。

在下一个导览部分，我们将探讨数据绑定。
*/

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

var data2 = [][]string{[]string{"top left", "top right"},
	[]string{"bottom left", "bottom right"}}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Table Widget")

	list := widget.NewTable(
		func() (int, int) {
			return len(data2), len(data2[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data2[i.Row][i.Col])
		})

	myWindow.SetContent(list)
	myWindow.ShowAndRun()
}
