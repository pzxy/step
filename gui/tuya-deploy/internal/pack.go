package internal

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"os"
	"step/gui/tuya-deploy/common"
	"step/gui/tuya-deploy/util"
	"strings"
)

func (m *Manager) InitPackCanvasForm() fyne.CanvasObject {
	inArch := widget.NewEntry()
	ArchSelect := widget.NewSelect([]string{"x86_84", "armv7", "arm64"}, func(value string) {
		inArch.Text = value
	})

	inOs := widget.NewEntry()
	inOsSelect := widget.NewSelect([]string{"Linux", "Darwin"}, func(value string) {
		inOs.Text = value
	})

	inGwVersion := widget.NewEntry()
	inGwVersion.SetPlaceHolder("格式举例 v1.4.0")

	inOutputDir := widget.NewEntry()
	inOutputDir.SetPlaceHolder("格式举例 /tmp/ ;默认文件名为：xxx.tar.gz")

	m.PackEntry = &common.PackEntry{
		Arch:      inArch,
		OS:        inOs,
		GwVersion: inGwVersion,
		OutputDir: inOutputDir,
	}

	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: common.HwArch, Widget: ArchSelect},
			{Text: common.Os, Widget: inOsSelect},
			{Text: common.PackVersion, Widget: inGwVersion},
			{Text: common.PackOutputDir, Widget: inOutputDir},
		},
		OnSubmit: func() { // optional, handle form submission
			m.encryptButtonFunc()
		},
		SubmitText: "打包",
	}
	return form
}

func (m *Manager) packButtonFunc() {
	if err := checkoutPackInfo(m.PackEntry); err != nil {
		dialog.ShowInformation("错误", fmt.Sprintf("%s", err), m.W)
		return
	}
	if err := assemblyPack(m.PackEntry); err != nil {
		dialog.ShowInformation("错误", fmt.Sprintf("打包失败:%s", err), m.W)
		return
	} else {
		dialog.ShowInformation("提示", "打包成功", m.W)
	}
}

func checkoutPackInfo(data *common.PackEntry) error {
	if data == nil {
		return fmt.Errorf("common.PackEntry is nil")
	}
	if data.Arch.Text == "" {
		return fmt.Errorf("错误:%s为空", common.HwArch)
	}
	if data.OS.Text == "" {
		return fmt.Errorf("错误:%s为空", common.Os)
	}
	if data.GwVersion.Text == "" {
		return fmt.Errorf("错误:%s为空", common.PackVersion)
	}
	if data.OutputDir.Text == "" {
		return fmt.Errorf("错误:%s为空", common.PackOutputDir)
	}
	fi, err := os.Stat(data.OutputDir.Text)
	if err != nil {
		return err
	}
	if !fi.IsDir() {
		return fmt.Errorf("错误:%s格式不对，这不是一个目录，或者没有此目录", common.OutputDir)
	}

	return nil
}

func assemblyPack(data *common.PackEntry, m *Manager) error {
	//edgex_v1.4.0_Darwin_x86_64.tar.gz
	info := []string{
		"edgex",
		data.GwVersion.Text,
		data.OS.Text,
		data.Arch.Text + ".tar.gz",
	}
	fileName := strings.Join(info, "_")
	filePath := data.OutputDir.Text + fileName

	url := data.OutputDir.Text + data.GwVersion.Text + "/" + fileName
	util.DownloadFile(filePath, url)

	return nil
}
