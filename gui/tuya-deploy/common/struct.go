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
	Port *widget.Entry
	User *widget.Entry
	Host *widget.Entry
	Path *widget.Entry
}
