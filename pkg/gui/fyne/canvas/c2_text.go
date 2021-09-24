package main

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Text")

	text := canvas.NewText("Text Object", color.White)
	text.Alignment = fyne.TextAlignTrailing
	//text.TextStyle = fyne.TextStyle{Italic: true}
	text.TextStyle = fyne.TextStyle{Bold: true, Italic: true, Monospace: true}

	w.SetContent(text)

	w.ShowAndRun()
}

/**
canvas.Text用于Fyne中的所有文本呈现。通过指定文本和文本的颜色来创建它。文本使用当前主题指定的默认字体呈现。

文本对象允许某些配置，例如Alignment和TextStyle字段。如此处示例所示。要改用等宽字体，可以指定fyne.TextStyle {Monospace：true}。

通过指定FYNE_FONT环境变量，可以使用其他字体。使用它来设置要使用的.ttf文件，而不是Fyne工具箱或当前主题中提供的文件。
*/
