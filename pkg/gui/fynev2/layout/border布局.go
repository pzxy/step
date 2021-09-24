package main

import (
	"image/color"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Border Layout")

	top := canvas.NewText("top bar", color.White)
	left := canvas.NewText("left", color.White)
	middle := canvas.NewText("content", color.White)
	//边框布局，无论窗口如何调整，都在窗口的一边。
	content := container.New(layout.NewBorderLayout(top, nil, left, nil),
		top, left, middle)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
