package internal

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"step/gui/tuya-deploy/const"
	"step/gui/tuya-deploy/cryptool"
)

func InitCanvas(w fyne.Window) fyne.CanvasObject {
	pid := widget.NewLabel(_const.Pid)
	inPid := widget.NewEntry()

	uuid := widget.NewLabel(_const.Uuid)
	inUuid := widget.NewEntry()

	deviceId := widget.NewLabel(_const.DeviceID)
	inDeviceId := widget.NewEntry()

	localKey := widget.NewLabel(_const.LocalKey)
	inLocalKey := widget.NewEntry()

	subDeviceLimit := widget.NewLabel(_const.SubDeviceLimit)
	inLimit := widget.NewEntry()

	macAddr := widget.NewLabel(_const.MacAddr)
	inMacAddr := widget.NewEntry()
	form := container.New(layout.NewFormLayout(),
		pid, inPid,
		uuid, inUuid,
		deviceId, inDeviceId,
		localKey, inLocalKey,
		subDeviceLimit, inLimit,
		macAddr, inMacAddr,
	)

	button := widget.NewButton(_const.ButtonEncrypt, func() {
		ret := CheckoutInput(inPid, inUuid, inDeviceId, inLocalKey, inLimit, inMacAddr)
		if ret != "" {
			dialog.ShowInformation("错误", ret, w)
			return
		}
		if err := cryptool.CreateConfigFile(inPid.Text, inUuid.Text, inDeviceId.Text, inLocalKey.Text, inLimit.Text); err != nil {
			dialog.ShowInformation("错误", fmt.Sprintf("创建yaml文件失败:%s", err), w)
			return
		}
		if err := cryptool.Encrypt(inMacAddr.Text, _const.Opt, _const.InputFile, _const.OutputFile); err != nil {
			dialog.ShowInformation("错误", fmt.Sprintf("加密失败:%s", err), w)
			return
		} else {
			dialog.ShowInformation("提示", "加密成功", w)
		}
	})

	center := container.New(layout.NewCenterLayout(), button)
	return container.New(layout.NewVBoxLayout(), form, center)
}

func CheckoutInput(inPid, inUuid, inDeviceId, inLocalKey, inLimit, inMacAddr *widget.Entry) string {
	if inPid.Text == "" {
		return fmt.Sprintf("错误:%s为空", _const.Pid)
	}
	if inUuid.Text == "" {
		return fmt.Sprintf("错误:%s为空", _const.Uuid)
	}
	if inDeviceId.Text == "" {
		return fmt.Sprintf("错误:%s为空", _const.DeviceID)
	}
	if inLocalKey.Text == "" {
		return fmt.Sprintf("错误:%s为空", _const.LocalKey)
	}
	if inLimit.Text == "" {
		return fmt.Sprintf("错误:%s为空", _const.SubDeviceLimit)
	}
	if inMacAddr.Text == "" {
		return fmt.Sprintf("错误:%s为空", _const.MacAddr)
	}

	return ""
}

func EncryptButton(inPid, inUuid, inDeviceId, inLocalKey, inLimit, inMacAddr *widget.Entry) {

}
