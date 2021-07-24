package upload

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type UploadEntry struct {
	W                fyne.Window
	Pid              *widget.Entry
	Uuid             *widget.Entry
	DeviceId         *widget.Entry
	SecKey           *widget.Entry
	LocalKey         *widget.Entry
	SubDeviceLimit   *widget.Entry
	SubDeviceTotal   *widget.Entry
	Env              *widget.Entry
	Region           *widget.Entry
	GwStatus         *widget.Entry
	Saas             *widget.Entry
	MacAddr          *widget.Entry
	DefaultOutputDir *widget.Entry
	Infinite         *widget.ProgressBarInfinite
}

func NewUpload(w fyne.Window) *UploadEntry {

	return &UploadEntry{
		W:                w,
		Pid:              widget.NewEntry(),
		Uuid:             widget.NewEntry(),
		DeviceId:         widget.NewEntry(),
		SecKey:           widget.NewEntry(),
		LocalKey:         widget.NewEntry(),
		MacAddr:          widget.NewEntry(),
		SubDeviceLimit:   widget.NewEntry(),
		DefaultOutputDir: widget.NewEntry(),
		Env:              widget.NewEntry(),
		Region:           widget.NewEntry(),
		SubDeviceTotal:   widget.NewEntry(),
		GwStatus:         widget.NewEntry(),
		Saas:             widget.NewEntry(),
	}
}
