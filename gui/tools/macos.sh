#!/bin/sh
# mac下打包mac版本
go get fyne.io/fyne/v2/cmd/fyne
go build
fyne package -icon lib/logo/256.png -name Edgetools

