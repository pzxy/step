package pack

import (
	"example.com/m/v2/common"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type PackEntry struct {
	W         fyne.Window
	Infinite  *widget.ProgressBarInfinite
	Arch      *widget.Entry
	OS        *widget.Entry
	OutputDir *widget.Entry

	ScriptInfo *widget.Entry
	common.GwService
}

func NewPackEntry(w fyne.Window) *PackEntry {
	p := &PackEntry{
		W:          w,
		Arch:       widget.NewEntry(),
		OS:         widget.NewEntry(),
		OutputDir:  widget.NewEntry(),
		ScriptInfo: widget.NewEntry(),

		GwService: common.GwService{
			DockerRegistry: widget.NewEntry(),
			GwVersion:      widget.NewEntry(),
			TedgeUuid:      widget.NewEntry(),
			EdgeHub:        widget.NewEntry(),
			Resource:       widget.NewEntry(),
			Agent:          widget.NewEntry(),
			Expert:         widget.NewEntry(),
			GateWay:        widget.NewEntry(),
			FileBeat:       widget.NewEntry(),
			WebClient:      widget.NewEntry(),
		},
	}
	return p
}
