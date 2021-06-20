package internal

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"step/gui/tuya-deploy/common"
)

func (m *Manager) NewUploadLabelC() *fyne.Container {
	user := widget.NewLabel(common.User)
	inUser := widget.NewEntry()

	host := widget.NewLabel(common.Host)
	inHost := widget.NewEntry()

	port := widget.NewLabel(common.Port)
	inPort := widget.NewEntry()

	path := widget.NewLabel(common.Path)
	inPath := widget.NewEntry()

	m.UploadEntry = &common.UploadEntry{
		User: inUser,
		Host: inHost,
		Port: inPort,
		Path: inPath,
	}

	return container.New(layout.NewFormLayout(),
		user, inUser,
		host, inHost,
		port, inPort,
		path, inPath,
	)
}

func (m *Manager) NewUploadButtonC() *fyne.Container {
	buttonFunc := func() {
		if err := checkoutUpload(m.UploadEntry); err != nil {
			dialog.ShowInformation("错误", fmt.Sprintf("%s", err), m.W)
			return
		}
	}

	return container.New(layout.NewCenterLayout(), widget.NewButton(common.ButtonUpload, buttonFunc))

}

func checkoutUpload(data *common.UploadEntry) error {
	if data == nil {
		return fmt.Errorf("EncryptEntry is nil")
	}
	if data.User.Text == "" {
		return fmt.Errorf("错误:%s为空", common.User)
	}
	if data.Host.Text == "" {
		return fmt.Errorf("错误:%s为空", common.Host)
	}
	if data.Port.Text == "" {
		return fmt.Errorf("错误:%s为空", common.Port)
	}
	if data.Path.Text == "" {
		return fmt.Errorf("错误:%s为空", common.Path)
	}
	return nil
}
