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
	//encryptLabel := m.NewEncryptLabelC()
	//encryptButton := m.NewEncryptButtonC()
	form := m.InitEncryptCanvasForm()
	local := container.New(layout.NewVBoxLayout(), form)
	//other := container.New(layout.NewVBoxLayout(), encryptLabel, encryptButton)

	//tabs := container.NewAppTabs(
	//	container.NewTabItem("本机", local),
	//	container.NewTabItem("其他主机", other),
	//)
	//tabs.SetTabLocation(container.TabLocationTop)

	m.EncryptCanvas = local
}

func (m *Manager) InitUploadCanvas() {
	uploadLabel := m.NewUploadLabelC()
	uploadButton := m.NewUploadButtonC()
	m.UploadCanvas = container.New(layout.NewVBoxLayout(), uploadLabel, uploadButton)
}

func (m *Manager) InitPackCanvas() {
	form := m.InitPackCanvasForm()
	m.PackCanvas = container.New(layout.NewVBoxLayout(), form)
}
