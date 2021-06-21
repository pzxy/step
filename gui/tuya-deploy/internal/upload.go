package internal

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"io"
	"os/exec"
	"runtime"
	"step/gui/tuya-deploy/common"
)

func (m *Manager) NewUploadLabelC() *fyne.Container {
	user := widget.NewLabel(common.User)
	inUser := widget.NewEntry()

	host := widget.NewLabel(common.Host)
	inHost := widget.NewEntry()

	port := widget.NewLabel(common.Port)
	inPort := widget.NewEntry()

	path := widget.NewLabel(common.Path)
	inPath := widget.NewEntry()

	m.UploadEntry = &common.UploadEntry{
		User:    inUser,
		Host:    inHost,
		Port:    inPort,
		DstPath: inPath,
	}

	return container.New(layout.NewFormLayout(),
		user, inUser,
		host, inHost,
		port, inPort,
		path, inPath,
	)
}

func (m *Manager) NewUploadButtonC() *fyne.Container {
	buttonFunc := func() {
		if err := checkoutUpload(m.UploadEntry); err != nil {
			dialog.ShowInformation("错误", fmt.Sprintf("%s", err), m.W)
			return
		}

	}

	return container.New(layout.NewCenterLayout(), widget.NewButton(common.ButtonUpload, buttonFunc))

}

func checkoutUpload(data *common.UploadEntry) error {
	if data == nil {
		return fmt.Errorf("EncryptEntry is nil")
	}
	if data.User.Text == "" {
		return fmt.Errorf("错误:%s为空", common.User)
	}
	if data.Host.Text == "" {
		return fmt.Errorf("错误:%s为空", common.Host)
	}
	if data.Port.Text == "" {
		return fmt.Errorf("错误:%s为空", common.Port)
	}
	if data.DstPath.Text == "" {
		return fmt.Errorf("错误:%s为空", common.Path)
	}
	return nil
}

//ssh执行命令

//cmd将文件copy过去
//ssh执行文件部署

func uploadFile(u *common.UploadEntry) error {
	//scp /home/space/music/1.mp3 root@www.runoob.com:/home/root/others/music
	switch runtime.GOOS {
	case "darwin":
		if e := execute(nil, "bash", "-c", fmt.Sprintf("scp -P %v %v %v@%v:~/", u.Port, u.User, u.DstPath)); e != nil {
			return e
		}
	case "windows":
		//execute(nil, "cmd", "/c")
	case "linux":
		//execute(nil, "bash", "-c")
	}
	return nil
}

func execute(outPut io.Writer, command string, params ...string) error {
	cmd := exec.Command(command, params...)
	//cmd.Stdout = outPut
	//cmd.Stderr = outPut
	return cmd.Run()
}

/**
var sssString = []string{
		"uname -m",
		"uname -s",
		"ifconfig -a |\nawk '/^[a-z]/ { iface=$1; mac=$NF; next }\n    /inet addr:/ { print iface, mac }'",
	}


sshHost := "192.168.1.239"
	sshUser := "root"
	sshPassword := "root"
	sshType := "password" //password 或者 key
	sshKeyPath := ""      //ssh id_rsa.id 路径"
	sshPort := 22
*/
