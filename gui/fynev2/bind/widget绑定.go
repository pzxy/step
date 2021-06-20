package main

/**
绑定简单的小部件
绑定小部件的最简单方法是将绑定项作为值而不是静态值传递给它。许多小部件提供了一个WithData构造函数，该构造函数将接受类型化数据绑定项。要设置绑定，您需要做的就是传入正确的类型。

尽管这在初始代码中看起来并没有多大好处，但您可以看到它如何确保显示的内容始终与数据源保持同步。你会发现，我们怎么也没需要调用Refresh() 的Label小部件，甚至保持对它的引用，但它相应地更新。

在下一步中，我们将了解如何设置通过双向绑定来编辑值的小部件。
*/

import (
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Simple")

	str := binding.NewString()
	str.Set("Initial value")

	text := widget.NewLabelWithData(str)
	w.SetContent(text)

	go func() {
		time.Sleep(time.Second * 4)
		str.Set("A new string")
	}()

	w.ShowAndRun()
}
