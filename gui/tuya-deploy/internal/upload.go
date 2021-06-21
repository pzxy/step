package internal

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"io"
	"net/http"
	"os"
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
		if e := execute(nil, "bash", "-c", fmt.Sprintf("scp -P %v %v %v@:~/", u.Port, u.User, u.DstPath)); e != nil {
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

func uploadExecSsh(data *common.UploadEntry) error {
	cmd := make(map[common.CmdKey]string, 0)
	cmd[common.Kernel] = "uname -s"
	cmd[common.HwArch] = "uname -m"
	order := []common.CmdKey{common.Kernel, common.HwArch}
	info := &common.SSHInfo{
		PassWord:  data.PassWord.Text,
		Host:      data.Host.Text,
		Port:      data.Port.Text,
		User:      data.User.Text,
		LoginType: common.PassWord,
		Cmd:       cmd,
		CmdOrder:  order,
	}
	m, err := common.SSH(info)
	if err != nil {
		return err
	}
	kernel := m[common.Kernel]
	hwArch := m[common.HwArch]
	fileName := "edgex_v1.5.0" + "_" + kernel + "_" + hwArch + ".tar.gz"
	downloadFile(fileName, "")
	return nil
}

func downloadFile(fileName string, version string) error {
	filePath := common.DownloadFilePrefix + fileName
	url := common.DownloadUrlPrefix + version + "/" + fileName
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

//func (t *DtTask) Md5Check(body []byte) bool {
//	flag := false
//	glog.V(log.Unc).Infof(t.StrTaskId + ", enter  Md5Check()...")
//	m := md5.Sum(body)
//
//	tmp := make([]byte, 0)
//	tmp = append(tmp, m[:]...)
//	mStr := hex.EncodeToString(tmp) //必须用这个
//
//	glog.V(log.Unc).Infof(t.StrTaskId+"md5校验 (原MD5:%s)(%s)", t.order.MD5, mStr)
//
//	if mStr == t.order.MD5 {
//		flag = true
//	}
//	glog.V(log.Unc).Infof(t.StrTaskId + ", leaves  Md5Check()...")
//	return flag
//}
