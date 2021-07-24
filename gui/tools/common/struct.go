package common

import (
	"fyne.io/fyne/v2/widget"
)

type UploadEntry struct {
	Port      *widget.Entry
	User      *widget.Entry
	Host      *widget.Entry
	LoginData *widget.Entry
	SrcPath   *widget.Entry
	DstPath   *widget.Entry
	LoginType *widget.Entry
}

type SSHInfo struct {
	Host      string
	User      string
	Port      string
	LoginType string //password 或者 key,使用下拉框来做
	LoginData string
	Cmd       map[string]string
	CmdOrder  []string //命令执行顺序，从前往后执行。
}

type CMDInfo struct {
	Cmd      map[string]string
	CmdOrder []string //命令执行顺序，从前往后执行。
	PassWord string
}

type GwService struct {
	InitDirP       bool
	Clear          bool
	TedgeUuid      *widget.Entry
	GwVersion      *widget.Entry
	DockerRegistry *widget.Entry
	EdgeHub        *widget.Entry
	Resource       *widget.Entry
	Agent          *widget.Entry
	Expert         *widget.Entry
	GateWay        *widget.Entry
	FileBeat       *widget.Entry
	WebClient      *widget.Entry
}
