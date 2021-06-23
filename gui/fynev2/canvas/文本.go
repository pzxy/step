package main

/**
canvas.Text用于 Fyne 中的所有文本渲染。它是通过为文本指定文本和颜色来创建的。使用当前主题指定的默认字体呈现文本。

文本对象允许某些配置，如Alignment 和TextStyle字段。如此处的示例所示。要改用等宽字体，您可以指定 fyne.TextStyle{Monospace: true}.

可以通过指定FYNE_FONT 环境变量来使用替代字体。使用它来设置.ttf要使用的文件，而不是 Fyne 工具包或当前主题中提供的文件。
*/

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"step/gui/fynev2/theme/t3"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Text")
	w.Resize(fyne.Size{Height: 640, Width: 460})
	myApp.Settings().SetTheme(&t3.MyTheme{})
	//en := widget.NewEntry()

	intro := widget.NewLabel("使用说明")
	intro.Wrapping = fyne.TextTruncate
	intro.SetText("可以通过指定FYNE_FONT 环境变量来使用替代字体。\n 使用它来设置.ttf要使用的文件，而不是 Fyne 工具包或当前主题中提供的文件。")
	//text := canvas.NewText("可以通过指定FYNE_FONT 环境变量来使用替代字体。\\n"+
	//	"使用它来设置.ttf", color.White)
	//text.Alignment = fyne.TextAlignTrailing
	centered := container.New(layout.NewVBoxLayout(), intro, widget.NewSeparator(), intro, fyne.NewMenuItemSeparator())
	w.SetContent(centered)

	w.ShowAndRun()
}
