package main

import (
	"example.com/m/v2/lib"
	"fyne.io/fyne"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(&lib.MyTheme{})

	w := a.NewWindow("水调个头")
	w.Resize(fyne.Size{700, 460})
	w.ShowAndRun()
}
