package main

/**
双向绑定
到目前为止，我们已经将数据绑定视为使我们的用户界面元素保持最新的一种方式。然而，更常见的是需要从 UI 小部件更新值并保持数据在任何地方都是最新的。幸运的是，Fyne 中提供的绑定是“双向”的，这意味着可以将值推入其中也可以读出。数据更改将传达给所有连接的代码，无需任何附加代码。

要查看此操作，我们可以更新最后一个测试应用程序以显示绑定到相同值的aLabel和 an Entry。通过设置，您可以看到通过条目编辑值也会更新标签中的文本。这一切都是可能的，无需在我们的代码中调用 refresh 或引用小部件。

通过移动您的应用程序以使用数据绑定，您可以停止保存指向所有小部件的指针。通过将您的数据捕获为一组绑定值，您的用户界面可以是完全独立的代码。阅读更清晰，更易于管理。

接下来，我们将看看如何在我们的数据中添加转换。
*/

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Two Way")

	str := binding.NewString()
	str.Set("Hi!")

	w.SetContent(container.NewVBox(
		widget.NewLabelWithData(str),
		widget.NewEntryWithData(str),
	))

	w.ShowAndRun()
}
