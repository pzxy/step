package internal

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"step/gui/tuya-deploy/common"
)

func (m *Manager) InitPackCanvasForm() fyne.CanvasObject {
	inArch := widget.NewEntry()
	//inPid.Validator = validation.NewRegexp(`^[A-Za-z0-9_-]+$`, "username can only contain letters, numbers, '_', and '-'")

	inOs := widget.NewEntry()
	//inUuid.Validator = validation.NewRegexp(`^[A-Za-z0-9_-]+$`, "username can only contain letters, numbers, '_', and '-'")

	inGwVersion := widget.NewEntry()
	inGwVersion.SetPlaceHolder("格式举例 v1.4.0")
	//inMacAddr.Validator = validation.NewRegexp(`^[A-Za-z0-9_-]+$`, "username can only contain letters, numbers, '_', and '-'")

	inOutputDir := widget.NewEntry()
	inOutputDir.SetPlaceHolder("格式举例 /tmp/encrypted; 注意：encrypted文件名")

	//inOutputDir.Validator = validation.NewRegexp(`^[A-Za-z0-9_-]+$`, "username can only contain letters, numbers, '_', and '-'")
	m.PackEntry = &common.PackEntry{
		Arch:      inArch,
		OS:        inOs,
		GwVersion: inGwVersion,
		OutputDir: inOutputDir,
	}

	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: common.HwArch, Widget: inArch},
			{Text: common.Os, Widget: inOs},
			{Text: common.PackVersion, Widget: inGwVersion},
			{Text: common.OutputDir, Widget: inOutputDir},
		},
		OnSubmit: func() { // optional, handle form submission
			m.encryptButtonFunc()
		},
		SubmitText: "加密",
	}
	return nil
}
