package main

import (
	"example.com/m/v2/internal"
	"example.com/m/v2/lib"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	a := app.New()

	w := a.NewWindow("部署工具")
	a.Settings().SetTheme(&lib.MyTheme{})
	w.Resize(fyne.Size{700, 460})

	m := &internal.Manager{}
	m.W = w
	m.Init()

	tabs := container.NewAppTabs(
		container.NewTabItem("使用说明", m.Instructions),
		container.NewTabItem("手动部署", m.ManualDeploy),
		container.NewTabItem("自动部署", m.AutoDeploy),
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	w.SetContent(tabs)
	w.ShowAndRun()
}
