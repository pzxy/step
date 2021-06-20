#!/bin/sh

#第1步：下载fyne工程模版 git@github.com:lusingander/fyne-font-example.git
#第2步：下载字体ShangShouJianSongXianXiTi-2.ttf到fyne工程目录（也可以放到子目录）
#第3步：进入字体下载目录，执行fyne bundle ShangShouJianSongXianXiTi-2.ttf >> bundle.go
#第4步：修改工程下的theme.go -> MyTheme结构体定义的字体变量
#第5步：main方法中引入新字体

go get fyne.io/fyne/v2/cmd/fyne
fyne bundle ShangShouJianSongXianXiTi-2.ttf >> bundle.go

#a := app.New()
#a.Settings().SetTheme(&lib.MyTheme{})
# 字体地址 https://www.fonts.net.cn/fonts-zh/tag-songti-1.html