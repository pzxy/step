#!/bin/zsh

# 在mac下交叉编译打包windows版本
go get github.com/fyne-io/fyne-cross
fyne-cross windows -arch=amd64