package upload

//
//import (
//	"example.com/m/v2/common"
//	"example.com/m/v2/util"
//	"fmt"
//	"fyne.io/fyne/v2"
//	"fyne.io/fyne/v2/container"
//	"fyne.io/fyne/v2/dialog"
//	"fyne.io/fyne/v2/layout"
//	"fyne.io/fyne/v2/widget"
//	"regexp"
//	"runtime"
//)
//
//func (m *Manager) NewUploadLabelC() *fyne.Container {
//
//	loginData := widget.NewLabel(common.Password)
//	inLoginData := widget.NewEntry()
//	inLoginData.Password = true
//
//	loginType := widget.NewLabel(common.LoadType)
//	inLoginType := widget.NewEntry()
//	inLoginTypeSelect := widget.NewSelect([]string{common.PassWordType}, func(value string) {
//		inLoginType.Text = value
//		if common.KeyType == value {
//			loginData.Text = common.KeyPath
//			inLoginData.Password = false
//		} else {
//			loginData.Text = common.Password
//			inLoginData.Password = true
//		}
//	})
//	inLoginTypeSelect.SetSelectedIndex(0)
//
//	user := widget.NewLabel(common.User)
//	inUser := widget.NewEntry()
//	inUser.SetText("root")
//	host := widget.NewLabel(common.Host)
//	inHost := widget.NewEntry()
//
//	port := widget.NewLabel(common.Port)
//	inPort := widget.NewEntry()
//	inPort.SetText("22")
//
//	srcPath := widget.NewLabel(common.SrcPath)
//	inSrcPath := widget.NewEntry()
//	switch runtime.GOOS {
//	case common.Windows:
//		inSrcPath.Text = common.UploadWindowDefaultInFilePath
//	case common.Darwin, common.Linux:
//		inSrcPath.Text = common.UploadUnixDefaultInFilePath
//	default:
//		dialog.ShowError(fmt.Errorf("未知操作系统：%s", runtime.GOOS), m.W)
//	}
//	wb := widget.NewButton("选择文件", func() {
//		dialog.ShowFileOpen(func(closer fyne.URIReadCloser, err error) {
//			if err != nil {
//				dialog.ShowError(err, m.W)
//				return
//			}
//			if closer == nil {
//				return
//			}
//			inSrcPath.Text = closer.URI().Path()
//			inSrcPath.Refresh()
//		}, m.W)
//	})
//	c := container.NewHSplit(inSrcPath, wb)
//	c.SetOffset(10)
//	inUploadFileContainer := c
//
//	dstPath := widget.NewLabel(common.DstPath)
//	inDstPath := widget.NewEntry()
//	inDstPath.Text = "/tmp/"
//
//	m.UploadEntry = &common.UploadEntry{
//		LoginType: inLoginType,
//		User:      inUser,
//		LoginData: inLoginData,
//		Host:      inHost,
//		Port:      inPort,
//		SrcPath:   inSrcPath,
//		DstPath:   inDstPath,
//	}
//
//	return container.New(layout.NewFormLayout(),
//		loginType, inLoginTypeSelect,
//		user, inUser,
//		loginData, inLoginData,
//		host, inHost,
//		port, inPort,
//		srcPath, inUploadFileContainer,
//		dstPath, inDstPath,
//	)
//}
//
//func (m *Manager) NewUploadButtonC() *widget.Button {
//	buttonFunc := func() {
//		if err := checkoutUpload(m.UploadEntry); err != nil {
//			dialog.ShowInformation("错误", fmt.Sprintf("%s", err), m.W)
//			return
//		}
//	}
//	return widget.NewButton(common.ButtonUpload, buttonFunc)
//
//}
//
//func (m *Manager) NewPingButtonC() *widget.Button {
//	buttonFunc := func() {
//		if err := checkoutUpload(m.UploadEntry); err != nil {
//			dialog.ShowInformation("错误", fmt.Sprintf("%s", err), m.W)
//			return
//		}
//		if err := pingSSh(m.UploadEntry); err != nil {
//			dialog.ShowInformation("错误", fmt.Sprintf("%s", err), m.W)
//			return
//		}
//		dialog.ShowInformation("连接成功", "好嘞", m.W)
//	}
//
//	return widget.NewButton(common.ButtonPing, buttonFunc)
//}
//
//func checkoutUpload(data *common.UploadEntry) error {
//	if data == nil {
//		return fmt.Errorf("EncryptEntry is nil")
//	}
//	if data.User.Text == "" {
//		return fmt.Errorf("错误:%s为空", common.User)
//	}
//	if data.LoginData.Text == "" {
//		var loginStr string
//		if data.LoginType.Text == common.PassWordType {
//			loginStr = common.Password
//		} else if data.LoginType.Text == common.KeyType {
//			loginStr = common.KeyPath
//		} else {
//			loginStr = "data.LoginType.Text:" + data.LoginType.Text
//		}
//		return fmt.Errorf("错误:%s为空", loginStr)
//	}
//	if data.Host.Text == "" {
//		return fmt.Errorf("错误:%s为空", common.Host)
//	}
//	if data.Host.Text <= "0" || data.Host.Text > "65535" {
//		return fmt.Errorf("错误:%s大小超出范围", common.Host)
//	}
//	matched, err := regexp.MatchString(`^((2(5[0-5]|[0-4]\d))|[0-1]?\d{1,2})(\.((2(5[0-5]|[0-4]\d))|[0-1]?\d{1,2})){3}$`, data.Host.Text)
//	if err != nil {
//		return err
//	}
//	if !matched {
//		return fmt.Errorf("错误:%s格式不对", common.Host)
//	}
//	if data.Port.Text == "" {
//		return fmt.Errorf("错误:%s为空", common.Port)
//	}
//	if data.DstPath.Text == "" {
//		return fmt.Errorf("错误:%s为空", common.SrcPath)
//	}
//	return nil
//}
//
//func pingSSh(data *common.UploadEntry) error {
//	ls := "ls"
//
//	cmd := make(map[string]string, 0)
//	cmd[ls] = "ls"
//	order := []string{ls}
//	info := &common.SSHInfo{
//		LoginData: data.LoginData.Text,
//		Host:      data.Host.Text,
//		Port:      data.Port.Text,
//		User:      data.User.Text,
//		LoginType: data.LoginType.Text,
//		Cmd:       cmd,
//		CmdOrder:  order,
//	}
//	_, err := util.SSH(info)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
////cmd将文件copy过去
////func copyFileByCmd(data *common.UploadEntry)error{
////
////}
//
////ssh执行文件部署
//
//func uploadExecSsh(data *common.UploadEntry) error {
//	var kernel = "kernel"
//	var hwArch = "hwArch"
//
//	cmd := make(map[string]string, 0)
//	cmd[kernel] = "uname -s"
//	cmd[hwArch] = "uname -m"
//	order := []string{kernel, hwArch}
//	info := &common.SSHInfo{
//		LoginData: data.LoginData.Text,
//		Host:      data.Host.Text,
//		Port:      data.Port.Text,
//		User:      data.User.Text,
//		LoginType: common.PassWordType,
//		Cmd:       cmd,
//		CmdOrder:  order,
//	}
//	m, err := util.SSH(info)
//	if err != nil {
//		return err
//	}
//	kel := m[kernel]
//	arch := m[hwArch]
//	fileName := "edgex_v1.5.0" + "_" + kel + "_" + arch + ".tar.gz"
//	//	filePath := common.DownloadFilePrefix + fileName
//	//	url := common.DownloadUrlPrefix + version + "/" + fileName
//	util.DownloadFile(fileName, "")
//	return nil
//}
//
////func (t *DtTask) Md5Check(body []byte) bool {
////	flag := false
////	glog.V(log.Unc).Infof(t.StrTaskId + ", enter  Md5Check()...")
////	m := md5.Sum(body)
////
////	tmp := make([]byte, 0)
////	tmp = append(tmp, m[:]...)
////	mStr := hex.EncodeToString(tmp) //必须用这个
////
////	glog.V(log.Unc).Infof(t.StrTaskId+"md5校验 (原MD5:%s)(%s)", t.order.MD5, mStr)
////
////	if mStr == t.order.MD5 {
////		flag = true
////	}
////	glog.V(log.Unc).Infof(t.StrTaskId + ", leaves  Md5Check()...")
////	return flag
////}
//
//func (m *Manager) InitUploadCanvas() fyne.CanvasObject {
//	uploadLabel := m.NewUploadLabelC()
//	pingButton := m.NewPingButtonC()
//	uploadButton := m.NewUploadButtonC()
//	b := container.New(layout.NewAdaptiveGridLayout(2), pingButton, uploadButton)
//	return container.New(layout.NewVBoxLayout(), uploadLabel, b)
//}
