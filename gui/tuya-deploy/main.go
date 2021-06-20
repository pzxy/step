package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"step/gui/tuya-deploy/internal"
	"step/gui/tuya-deploy/lib"
)

func main() {
	a := app.New()

	w := a.NewWindow("加密工具")
	a.Settings().SetTheme(&lib.MyTheme{})
	w.Resize(fyne.Size{450, 600})
	c := internal.InitCanvas(w)
	w.SetContent(c)
	w.ShowAndRun()
}
