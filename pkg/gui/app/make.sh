#!/bin/sh
# shellcheck disable=SC2034
APP_NAME=${1:-"stitch"}
set -e
go get -d fyne.io/fyne/v2/cmd/fyne
go mod vendor
# mac下打包mac版本
go build
fyne package -icon lib/deploy_logo/256.png -name $APP_NAME
zip -r -o  ./$APP_NAME-mac.zip ./$APP_NAME.app

# 交叉编译打包windows版本
#go get github.com/fyne-io/fyne-cross
fyne-cross windows -arch=amd64 -name $APP_NAME.exe -icon lib/deploy_logo/256.png
cp ./fyne-cross/dist/windows-amd64/$APP_NAME.exe.zip  ./$APP_NAME-win.zip