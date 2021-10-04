package main

import "fyne.io/fyne/v2/container"

/**
盒子布局
如容器和布局中所述，容器内的元素可以使用布局进行排列。本节探讨内置布局以及如何使用它们。

最常用的布局是layout.BoxLayout，它有两种变体，水平和垂直。框布局将所有元素排列在一行或一列中，并带有可选的空格以帮助对齐。

创建的水平框布局layout.NewHBoxLayout()在单行中创建项目排列。框中的每个项目都将其宽度设置为它，MinSize().Width并且所有项目的高度将相等，所有MinSize().Height值中最大的一个。布局可以在容器中使用，也可以使用框小部件 widget.NewHBox()。

垂直框布局与此类似，但它将项目排列在一列中。每个项目的高度设置为最小值，所有宽度都相等，设置为最小宽度中的最大值。

要在元素之间创建一个扩展空间（例如，一些左对齐而另一些右对齐）添加 alayout.NewSpacer() 作为项目之一。垫片将扩展以填充所有可用空间。在垂直框布局的开头添加一个间隔将导致所有项目底部对齐。您可以在水平排列的开头和结尾添加一个以创建居中对齐
*/

//容器中的所有元素都可以用layout来布局。

import (
	"image/color"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/layout"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Box Layout")

	text1 := canvas.NewText("Hello", color.White)
	text2 := canvas.NewText("There", color.White)
	text3 := canvas.NewText("(right)", color.White)
	content := container.New(layout.NewHBoxLayout(), text1, text2, layout.NewSpacer(), text3)

	text4 := canvas.NewText("centered", color.White)
	centered := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), text4, layout.NewSpacer())
	myWindow.SetContent(container.New(layout.NewVBoxLayout(), content, centered))
	myWindow.ShowAndRun()
}
