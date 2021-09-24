package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("TabContainer Widget")

	tabs := widget.NewTabContainer(
		widget.NewTabItem("Tab 1", widget.NewLabel("Hello")),
		widget.NewTabItem("Tab 2", widget.NewLabel("World!")))

	tabs.Append(widget.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Home tab")))

	tabs.SetTabLocation(widget.TabLocationLeading)

	myWindow.SetContent(tabs)
	myWindow.ShowAndRun()
}

/**
选项卡容器窗口小部件用于允许用户在不同的内容面板之间切换。选项卡可以是文本，也可以是文本和图标。建议不要混用一些带有图标的标签，而有些则不要。使用widget.NewTabContainer（...）并传递widget.TabItem项（可以使用widget.NewTabItem（...）创建）来创建选项卡容器。

可以通过设置选项卡的位置来配置选项卡容器，这些选项卡是widget.TabLocationTop，widget.TabLocationBottom，widget.TabLocationLeading和widget.TabLocationTrailing之一。默认位置是top。

在移动设备上加载时，选项卡位置可能会被忽略。在纵向方向上，前导或尾随位置将更改为底部。当处于横向时，顶部或底部位置将移至最前。
*/
