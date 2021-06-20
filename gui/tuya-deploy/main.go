package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"step/gui/tuya-deploy/internal"
	"step/gui/tuya-deploy/lib"
)

func main() {
	a := app.New()

	w := a.NewWindow("加密工具")
	a.Settings().SetTheme(&lib.MyTheme{})
	w.Resize(fyne.Size{600, 400})
	m := &internal.Manager{
		W: w,
	}
	m.Run()

	tabs := container.NewAppTabs(
		//这个第二个参数是画布对象
		container.NewTabItem("加密工具", m.EncryptCanvas),
		container.NewTabItem("上传文件", m.UploadCanvas),
	)

	//tabs.Append(container.NewTabItemWithIcon("Home", theme.FileVideoIcon(), widget.NewLabel("Home tab")))

	tabs.SetTabLocation(container.TabLocationLeading)

	w.SetContent(tabs)
	w.ShowAndRun()
}
