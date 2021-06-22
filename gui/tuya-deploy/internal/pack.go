package internal

import (
	"fyne.io/fyne/v2"
)

func (m *Manager) InitPackCanvasForm() fyne.CanvasObject {
	//inArch := widget.NewEntry()
	////inPid.Validator = validation.NewRegexp(`^[A-Za-z0-9_-]+$`, "username can only contain letters, numbers, '_', and '-'")
	//
	//inOs := widget.NewEntry()
	////inUuid.Validator = validation.NewRegexp(`^[A-Za-z0-9_-]+$`, "username can only contain letters, numbers, '_', and '-'")
	//
	//inGwVersion := widget.NewEntry()
	//inGwVersion.SetPlaceHolder("格式举例 v1.4.0")
	////inMacAddr.Validator = validation.NewRegexp(`^[A-Za-z0-9_-]+$`, "username can only contain letters, numbers, '_', and '-'")
	//
	//inOutputDir := widget.NewEntry()
	//inOutputDir.SetPlaceHolder("格式举例 /tmp/encrypted; 注意：encrypted文件名")
	//
	////inOutputDir.Validator = validation.NewRegexp(`^[A-Za-z0-9_-]+$`, "username can only contain letters, numbers, '_', and '-'")
	//m.EncryptEntry = &common.EncryptEntry{
	//	Pid:            inPid,
	//	Uuid:           inUuid,
	//	DeviceId:       inDeviceId,
	//	SecKey:         inSecKey,
	//	LocalKey:       inLocalKey,
	//	MacAddr:        inMacAddr,
	//	SubDeviceLimit: inLimit,
	//	OutputDir:      inOutputDir,
	//}
	//form := &widget.Form{
	//	Items: []*widget.FormItem{ // we can specify items in the constructor
	//		{Text: common.Pid, Widget: inPid},
	//		{Text: common.Uuid, Widget: inUuid},
	//		{Text: common.DeviceID, Widget: inDeviceId},
	//		{Text: common.SecKey, Widget: inSecKey},
	//		{Text: common.LocalKey, Widget: inLocalKey},
	//		{Text: common.SubDeviceLimit, Widget: inLimit},
	//		{Text: common.MacAddr, Widget: inMacAddr},
	//		{Text: common.OutputDir, Widget: inOutputDir},
	//	},
	//	OnSubmit: func() { // optional, handle form submission
	//		m.encryptButtonFunc()
	//	},
	//	SubmitText: "加密",
	//}
	return nil
}
