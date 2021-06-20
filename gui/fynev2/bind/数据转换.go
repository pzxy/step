package main

/**
数据转换
到目前为止，我们已经使用了数据类型与输出类型匹配的数据绑定（例如Stringand Label或Entry）。通常需要呈现格式不正确的数据。为此，该binding软件包提供了许多有用的转换功能。

最常见的是，这将用于将不同类型的数据转换为字符串以显示在Label 或Entry小部件中。在代码中查看我们如何将 a 转换Float为Stringusing binding.FloatToString。可以通过移动滑块来编辑原始值。每次数据更改时，它都会运行转换代码并更新任何连接的小部件。

还可以使用格式字符串为用户界面添加更自然的输出。您可以看到我们的short绑定也在转换为Float，String但是通过使用WithFormat帮助器，我们可以传递格式字符串（类似于fmt包）以提供自定义输出。

最后在本节中，我们将查看列表数据。
*/

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Conversion")

	f := binding.NewFloat()
	str := binding.FloatToString(f)
	short := binding.FloatToStringWithFormat(f, "%0.0f%%")
	f.Set(25.0)

	w.SetContent(container.NewVBox(
		widget.NewSliderWithData(0, 100.0, f),
		widget.NewLabelWithData(str),
		widget.NewLabelWithData(short),
	))

	w.ShowAndRun()
}
