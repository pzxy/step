package internal

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func (m *Manager) Run() {
	encryptL := m.NewEncryptLabelC()
	encryptB := m.NewEncryptButtonC()
	uploadL := m.NewUploadLabelC()
	uploadB := m.NewUploadButtonC()
	m.EncryptCanvas = container.New(layout.NewVBoxLayout(), encryptL, encryptB)
	m.UploadCanvas = container.New(layout.NewVBoxLayout(), uploadL, uploadB)
}
