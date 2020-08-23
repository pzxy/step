package main

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Box Layout")

	text1 := canvas.NewText("Hello", color.White)
	text2 := canvas.NewText("There", color.White)
	text3 := canvas.NewText("(right)", color.White)
	container := fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		text1, text2, layout.NewSpacer(), text3)

	text4 := canvas.NewText("centered", color.White)
	centered := fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		layout.NewSpacer(), text4, layout.NewSpacer())
	myWindow.SetContent(fyne.NewContainerWithLayout(layout.NewVBoxLayout(), container, centered))
	myWindow.ShowAndRun()
}

/**
如容器和布局中所述，可以使用布局来排列容器内的元素。本节探讨内置布局以及如何使用它们。

最常用的布局是layout.BoxLayout，它具有水平和垂直两个变体。箱形布局将所有元素排列在一行或一列中，并带有可选的空格以帮助对齐。

使用layout.NewHBoxLayout（）创建的水平框布局可在单行中创建项目的排列。框中的每个项目的宽度都将设置为其MinSize（）。Width，所有项目的高度都将相等，这是所有MinSize（）。Height值中的最大值。该布局可以在容器中使用，也可以使用框小部件widget.NewHBox（）。

垂直框的布局与此类似，但是它将项目排列在一个列中。每个项目的高度都将设置为最小，并且所有宽度都将相等，设置为最小宽度中的最大宽度。

为了在元素之间创建一个扩大的空间（例如，使某些元素保持对齐，而另一些元素保持右对齐），请添加layout.NewSpacer（）作为其中一项。垫片将扩展以填充所有可用空间。在垂直框布局的开头添加间隔符将使所有项目在底部对齐。您可以在水平排列的开头和结尾添加一个以创建中心对齐。
*/
