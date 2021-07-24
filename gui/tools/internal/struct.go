package internal

import (
	"fyne.io/fyne/v2"
)

type Manager struct {
	W            fyne.Window
	Instructions fyne.CanvasObject
	ManualDeploy fyne.CanvasObject
	AutoDeploy   fyne.CanvasObject
}
