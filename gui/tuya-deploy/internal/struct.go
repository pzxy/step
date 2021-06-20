package internal

import (
	"fyne.io/fyne/v2"
	"step/gui/tuya-deploy/common"
)

type Manager struct {
	W             fyne.Window
	EncryptCanvas fyne.CanvasObject
	UploadCanvas  fyne.CanvasObject
	EncryptEntry  *common.EncryptEntry
	UploadEntry   *common.UploadEntry
}
