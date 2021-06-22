package internal

import (
	"fyne.io/fyne/v2"
	"step/gui/tuya-deploy/common"
)

type Manager struct {
	W             fyne.Window
	EncryptCanvas fyne.CanvasObject
	EncryptEntry  *common.EncryptEntry

	UploadCanvas fyne.CanvasObject
	UploadEntry  *common.UploadEntry

	PackCanvas fyne.CanvasObject
	PackEntry  *common.PackEntry
}
