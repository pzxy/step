package internal

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"os"
	"regexp"
	"step/gui/tuya-deploy/common"
	"step/gui/tuya-deploy/cryptool"
)

func (m *Manager) NewEncryptLabelC() *fyne.Container {
	pid := widget.NewLabel(common.Pid)
	inPid := widget.NewEntry()

	uuid := widget.NewLabel(common.Uuid)
	inUuid := widget.NewEntry()

	deviceId := widget.NewLabel(common.DeviceID)
	inDeviceId := widget.NewEntry()

	secKey := widget.NewLabel(common.SecKey)
	inSecKey := widget.NewEntry()

	localKey := widget.NewLabel(common.LocalKey)
	inLocalKey := widget.NewEntry()

	subDeviceLimit := widget.NewLabel(common.SubDeviceLimit)
	inLimit := widget.NewEntry()

	macAddr := widget.NewLabel(common.MacAddr)
	inMacAddr := widget.NewEntry()

	m.EncryptEntry = &common.EncryptEntry{
		Pid:            inPid,
		Uuid:           inUuid,
		DeviceId:       inDeviceId,
		SecKey:         inSecKey,
		LocalKey:       inLocalKey,
		MacAddr:        inMacAddr,
		SubDeviceLimit: inLimit,
	}

	return container.New(layout.NewFormLayout(),
		pid, inPid,
		uuid, inUuid,
		deviceId, inDeviceId,
		secKey, inSecKey,
		localKey, inLocalKey,
		subDeviceLimit, inLimit,
		macAddr, inMacAddr,
	)
}

func (m *Manager) NewEncryptButtonC() *fyne.Container {
	buttonFunc := func() {
		m.encryptButtonFunc()
	}

	return container.New(layout.NewCenterLayout(), widget.NewButton(common.ButtonEncrypt, buttonFunc))
}

func checkoutEncrypt(data *common.EncryptEntry) error {
	if data == nil {
		return fmt.Errorf("EncryptEntry is nil")
	}
	if data.Pid.Text == "" {
		return fmt.Errorf("错误:%s为空", common.Pid)
	}
	if data.Uuid.Text == "" {
		return fmt.Errorf("错误:%s为空", common.Uuid)
	}
	if data.DeviceId.Text == "" {
		return fmt.Errorf("错误:%s为空", common.DeviceID)
	}
	if data.SecKey.Text == "" {
		return fmt.Errorf("错误:%s为空", common.SecKey)
	}
	if data.LocalKey.Text == "" {
		return fmt.Errorf("错误:%s为空", common.LocalKey)
	}
	if data.SubDeviceLimit.Text == "" {
		return fmt.Errorf("错误:%s为空", common.SubDeviceLimit)
	}
	if data.MacAddr.Text == "" {
		return fmt.Errorf("错误:%s为空", common.MacAddr)
	}
	matched, err := regexp.MatchString(`^([0-9a-fA-F]{2})(([/\\s:][0-9a-fA-F]{2}){5})$`, data.MacAddr.Text)
	if err != nil {
		return err
	}
	if !matched {
		return fmt.Errorf("错误:%s格式不对", common.MacAddr)
	}

	if data.OutputDir.Text == "" {
		data.OutputDir.Text = common.OutputFile
	} else {
		fi, err := os.Stat(data.OutputDir.Text)
		if err != nil {
			return err
		}
		if !fi.IsDir() {
			return fmt.Errorf("错误:%s格式不对，这不是一个目录，或者没有此目录", common.OutputDir)
		}
	}

	return nil
}

func (m *Manager) InitEncryptCanvasForm() fyne.CanvasObject {
	inPid := widget.NewEntry()
	//inPid.Validator = validation.NewRegexp(`^[A-Za-z0-9_-]+$`, "username can only contain letters, numbers, '_', and '-'")

	inUuid := widget.NewEntry()
	//inUuid.Validator = validation.NewRegexp(`^[A-Za-z0-9_-]+$`, "username can only contain letters, numbers, '_', and '-'")

	inDeviceId := widget.NewEntry()
	//inDeviceId.Validator = validation.NewRegexp(`^[A-Za-z0-9_-]+$`, "username can only contain letters, numbers, '_', and '-'")

	inSecKey := widget.NewEntry()
	//inSecKey.Validator = validation.NewRegexp(`^[A-Za-z0-9_-]+$`, "username can only contain letters, numbers, '_', and '-'")

	inLocalKey := widget.NewEntry()
	//inLocalKey.Validator = validation.NewRegexp(`^[A-Za-z0-9_-]+$`, "username can only contain letters, numbers, '_', and '-'")

	inLimit := widget.NewEntry()
	//inLimit.Validator = validation.NewRegexp(`^[A-Za-z0-9_-]+$`, "username can only contain letters, numbers, '_', and '-'")

	inMacAddr := widget.NewEntry()
	inMacAddr.SetPlaceHolder("格式举例 83:87:d3:61:e0:01")
	//inMacAddr.Validator = validation.NewRegexp(`^[A-Za-z0-9_-]+$`, "username can only contain letters, numbers, '_', and '-'")

	inOutputDir := widget.NewEntry()
	inOutputDir.SetPlaceHolder("不查看可不填。格式举例 /tmp/ ;文件名默认encrypted")

	//inOutputDir.Validator = validation.NewRegexp(`^[A-Za-z0-9_-]+$`, "username can only contain letters, numbers, '_', and '-'")
	m.EncryptEntry = &common.EncryptEntry{
		Pid:            inPid,
		Uuid:           inUuid,
		DeviceId:       inDeviceId,
		SecKey:         inSecKey,
		LocalKey:       inLocalKey,
		MacAddr:        inMacAddr,
		SubDeviceLimit: inLimit,
		OutputDir:      inOutputDir,
	}
	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: common.Pid, Widget: inPid},
			{Text: common.Uuid, Widget: inUuid},
			{Text: common.DeviceID, Widget: inDeviceId},
			{Text: common.SecKey, Widget: inSecKey},
			{Text: common.LocalKey, Widget: inLocalKey},
			{Text: common.SubDeviceLimit, Widget: inLimit},
			{Text: common.MacAddr, Widget: inMacAddr},
			{Text: common.OutputDir, Widget: inOutputDir},
		},
		OnSubmit: func() { // optional, handle form submission
			m.encryptButtonFunc()
		},
		SubmitText: "加密",
	}
	return form
}

func (m *Manager) encryptButtonFunc() {
	if err := checkoutEncrypt(m.EncryptEntry); err != nil {
		dialog.ShowInformation("错误", fmt.Sprintf("%s", err), m.W)
		return
	}
	if err := cryptool.CreateConfigFile(m.EncryptEntry); err != nil {
		dialog.ShowInformation("错误", fmt.Sprintf("创建yaml文件失败:%s", err), m.W)
		return
	}

	if err := cryptool.Encrypt(m.EncryptEntry.MacAddr.Text, common.Opt, common.InputFile, m.EncryptEntry.OutputDir.Text); err != nil {
		dialog.ShowInformation("错误", fmt.Sprintf("加密失败:%s", err), m.W)
		return
	} else {
		dialog.ShowInformation("提示", "加密成功", m.W)
	}
}
