package internal

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func (m *Manager) Run() {
	encryptLabel := m.NewEncryptLabelC()
	encryptButton := m.NewEncryptButtonC()
	uploadLabel := m.NewUploadLabelC()
	uploadButton := m.NewUploadButtonC()
	m.EncryptCanvas = container.New(layout.NewVBoxLayout(), encryptLabel, encryptButton)
	m.UploadCanvas = container.New(layout.NewVBoxLayout(), uploadLabel, uploadButton)
}
