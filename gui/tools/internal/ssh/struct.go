package ssh

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

var s *SSHEntry

type SSHEntry struct {
	W         fyne.Window
	Port      *widget.Entry
	User      *widget.Entry
	Host      *widget.Entry
	LoginData *widget.Entry
	LoginType *widget.Entry
	MacAddr   map[string]string
}

func NewSSHEntry(w fyne.Window) *SSHEntry {
	s = &SSHEntry{
		W:         w,
		Port:      widget.NewEntry(),
		User:      widget.NewEntry(),
		Host:      widget.NewEntry(),
		LoginData: widget.NewEntry(),
		LoginType: widget.NewEntry(),
	}
	return s
}
