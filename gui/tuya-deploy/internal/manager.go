package internal

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func (m *Manager) Run() {
	m.InitEncryptCanvas()
	m.InitUploadCanvas()
}

func (m *Manager) InitEncryptCanvas() {
	encryptLabel := m.NewEncryptLabelC()
	encryptButton := m.NewEncryptButtonC()
	local := container.New(layout.NewVBoxLayout(), encryptLabel, encryptButton)
	other := container.New(layout.NewVBoxLayout(), encryptLabel, encryptButton)

	tabs := container.NewAppTabs(
		container.NewTabItem("本机", local),
		container.NewTabItem("远程主机", other),
	)
	tabs.SetTabLocation(container.TabLocationTop)

	m.EncryptCanvas = tabs
}

func (m *Manager) InitUploadCanvas() {
	uploadLabel := m.NewUploadLabelC()
	uploadButton := m.NewUploadButtonC()
	m.UploadCanvas = container.New(layout.NewVBoxLayout(), uploadLabel, uploadButton)
}

//func (m *Manager) InitEncryptCanvas4Form() {
//	inPid := widget.NewEntry()
//
//	inUuid := widget.NewEntry()
//
//	inDeviceId := widget.NewEntry()
//
//	inSecKey := widget.NewEntry()
//
//	inLocalKey := widget.NewEntry()
//
//	inLimit := widget.NewEntry()
//
//	inMacAddr := widget.NewEntry()
//	form := &widget.Form{
//		Items: []*widget.FormItem{ // we can specify items in the constructor
//			{Text: common.Pid, Widget: inPid}},
//		OnSubmit: func() { // optional, handle form submission
//			m.NewEncryptButtonC()
//		},
//	}
//	m.EncryptCanvas = container.New(layout.NewVBoxLayout(), encryptLabel, encryptButton)
//}
