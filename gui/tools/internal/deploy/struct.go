package deploy

import (
	"example.com/m/v2/common"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type DeployEntry struct {
	W        fyne.Window
	Infinite *widget.ProgressBarInfinite

	OS         *widget.Entry
	Arch       *widget.Entry
	ScriptInfo *widget.Entry
	DockerUser *widget.Entry
	DockerPwd  *widget.Entry
	common.GwService
}

func NewDeploy(w fyne.Window) *DeployEntry {

	return &DeployEntry{
		W:          w,
		OS:         widget.NewEntry(),
		Arch:       widget.NewEntry(),
		ScriptInfo: widget.NewEntry(),
		DockerUser: widget.NewEntry(),
		DockerPwd:  widget.NewEntry(),
		GwService: common.GwService{
			DockerRegistry: widget.NewEntry(),

			GwVersion: widget.NewEntry(),
			TedgeUuid: widget.NewEntry(),
			EdgeHub:   widget.NewEntry(),
			Resource:  widget.NewEntry(),
			Agent:     widget.NewEntry(),
			Expert:    widget.NewEntry(),
			GateWay:   widget.NewEntry(),
			FileBeat:  widget.NewEntry(),
			WebClient: widget.NewEntry(),
		},
	}
}
