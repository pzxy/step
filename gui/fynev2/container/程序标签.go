package main

/**
应用程序标签
AppTabs 容器用于允许用户在不同的内容面板之间切换。标签要么只是文本，要么是文本和图标。建议不要混合一些有图标的标签和一些没有图标的标签。使用container.NewAppTabs(...)和传递 container.TabItem项目（可以使用 来创建 container.NewTabItem(...)）创建选项卡容器。

选项卡容器可以通过设置选项卡的位置container.TabLocationTop、container.TabLocationBottom、 container.TabLocationLeading和之一来配置container.TabLocationTrailing。默认位置是顶部。

在移动设备上加载时，标签位置可能会被忽略。在纵向中，前导或尾随位置将更改为底部。在横向时，顶部或底部位置将移动到前导。
*/

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	//"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("TabContainer Widget")

	tabs := container.NewAppTabs(
		//这个第二个参数是画布对象
		container.NewTabItem("Tab 1", widget.NewLabel("Hello")),
		container.NewTabItem("Tab 2", widget.NewLabel("World!")),
	)

	//tabs.Append(container.NewTabItemWithIcon("Home", theme.FileVideoIcon(), widget.NewLabel("Home tab")))

	tabs.SetTabLocation(container.TabLocationTrailing)

	myWindow.SetContent(tabs)
	myWindow.ShowAndRun()
}
