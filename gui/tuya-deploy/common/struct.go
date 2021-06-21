package common

import "fyne.io/fyne/v2/widget"

type EncryptEntry struct {
	Pid            *widget.Entry
	Uuid           *widget.Entry
	DeviceId       *widget.Entry
	SecKey         *widget.Entry
	LocalKey       *widget.Entry
	SubDeviceLimit *widget.Entry
	MacAddr        *widget.Entry
}

type UploadEntry struct {
	Port     *widget.Entry
	User     *widget.Entry
	Host     *widget.Entry
	PassWord *widget.Entry
	DstPath  *widget.Entry
}

type SSHInfo struct {
	Host      string
	User      string
	Port      string
	PassWord  string
	LoginType LoginType //password 或者 key,使用下拉框来做
	KeyPath   string    //ssh id_rsa.id 路径"
	Cmd       map[CmdKey]string
	CmdOrder  []CmdKey //命令执行顺序，从前往后执行。
}

type CMDInfo struct {
	Cmd      map[CmdKey]string
	CmdOrder []CmdKey //命令执行顺序，从前往后执行。
	PassWord string
}
