package upload

import (
	"example.com/m/v2/common"
	"example.com/m/v2/cryptool"
	"example.com/m/v2/internal/ssh"
	"example.com/m/v2/util"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"os"
	"regexp"
)

func (data *UploadEntry) checkoutEncrypt() error {
	if data == nil {
		return fmt.Errorf("UploadEntry is nil")
	}

	if data.DeviceId.Text == "" || data.SecKey.Text == "" || data.LocalKey.Text == "" {
		if data.DeviceId.Text != "" || data.SecKey.Text != "" || data.LocalKey.Text != "" {
			return fmt.Errorf("错误:%s,%s,%s,必须同时存在或为空", DeviceID, SecKey, LocalKey)
		}
	}
	if data.DeviceId.Text != "" && data.SecKey.Text != "" && data.LocalKey.Text != "" {
		data.GwStatus.Text = "true"
	} else {
		data.GwStatus.Text = "false"
	}
	if data.Uuid.Text == "" {
		return fmt.Errorf("错误:%s为空", Uuid)
	}

	if data.Pid.Text == "" {
		return fmt.Errorf("错误:%s为空", Pid)
	}

	if data.SubDeviceLimit.Text == "" {
		return fmt.Errorf("错误:%s为空", SubDeviceLimit)
	}
	if data.MacAddr.Text == "" {
		return fmt.Errorf("错误:%s为空", MacAddr)
	}

	matched, err := regexp.MatchString(`^([0-9a-fA-F]{2})(([/\\s:][0-9a-fA-F]{2}){5})$`, data.MacAddr.Text)
	if err != nil {
		return err
	}
	if !matched {
		return fmt.Errorf("错误:%s格式不对", MacAddr)
	}

	if data.GwStatus.Text == "" {
		return fmt.Errorf("错误:%s为空", GwStatus)
	}
	if data.Region.Text == "" {
		return fmt.Errorf("错误:%s为空", Region)
	}
	if data.Env.Text == "" {
		return fmt.Errorf("错误:%s为空", Env)
	}
	if data.Saas.Text == "" {
		return fmt.Errorf("错误:%s为空", Saas)
	}

	data.DefaultOutputDir.Text = OutputFile

	return nil
}

func (m *UploadEntry) initUploadForm() fyne.CanvasObject {
	m.Infinite = widget.NewProgressBarInfinite()
	m.Infinite.Hide()

	inRegionSelect := widget.NewSelect(RegionList, func(value string) {
		region := GetRegion(value)
		if region == "" {
			dialog.ShowError(fmt.Errorf("未知区域类型"), m.W)
			return
		}
		m.Region.Text = region
	})
	inRegionSelect.SetSelectedIndex(0)

	inEnvSelect := widget.NewSelect(EevList, func(value string) {
		env := GetEnv(value)
		if env == "" {
			dialog.ShowError(fmt.Errorf("未知环境类型"), m.W)
			return
		}
		m.Env.Text = env
	})
	inEnvSelect.SetSelectedIndex(1)

	inSassSelect := widget.NewSelect([]string{"true", "false"}, func(value string) {
		m.Saas.Text = value
	})
	inSassSelect.SetSelectedIndex(0)

	return &widget.Form{
		Items: []*widget.FormItem{
			{Text: DeviceID, Widget: m.DeviceId},
			{Text: SecKey, Widget: m.SecKey},
			{Text: LocalKey, Widget: m.LocalKey},
			{Text: Uuid, Widget: m.Uuid},
			{Text: Pid, Widget: m.Pid},
			{Text: SubDeviceLimit, Widget: m.SubDeviceLimit},
			{Text: Region, Widget: inRegionSelect},
			{Text: Env, Widget: inEnvSelect},
			{Text: Saas, Widget: inSassSelect},
		},
		OnSubmit: func() {
			m.uploadButtonFunc()
		},
		SubmitText: "上传加密",
	}
}

func (m *UploadEntry) setMacAddr() error {
	sshInfo := ssh.GetSSHLoginInfo()
	mapAddrs, err := util.GetOsMacAddr(&sshInfo)
	if err != nil {
		return err
	}
	for i, v := range mapAddrs {
		//一般第一个网卡都是docker0，所以不用这个。
		if len(mapAddrs) > 1 && i == 0 {
			continue
		}
		m.MacAddr.Text = v
		fmt.Println(m.MacAddr.Text)
		break
	}
	if m.MacAddr.Text == "" {
		return fmt.Errorf("无法获取mac地址：%v", mapAddrs)
	}
	return nil
}
func (m *UploadEntry) uploadButtonFunc() {
	m.Infinite.Show()
	defer m.Infinite.Hide()
	defer os.Remove(common.InputFile) //为了安全
	if err := m.setMacAddr(); err != nil {
		dialog.ShowInformation("错误", fmt.Sprintf("%s", err), m.W)
		return
	}

	if err := m.checkoutEncrypt(); err != nil {
		dialog.ShowInformation("错误", fmt.Sprintf("%s", err), m.W)
		return
	}

	e := uploadEntry2EncryptInfo(m)
	if err := cryptool.CreateConfigFile(&e); err != nil {
		dialog.ShowInformation("错误", fmt.Sprintf("创建yaml文件失败:%s", err), m.W)
		return
	} else {
		//go m.showEncryptInfo()
	}

	if err := cryptool.Encrypt(m.MacAddr.Text, false, common.InputFile, m.DefaultOutputDir.Text); err != nil {
		dialog.ShowInformation("错误", fmt.Sprintf("加密失败:%s", err), m.W)
		return
	}

	if err := m.uploadFile(); err != nil {
		dialog.ShowInformation("错误", fmt.Sprintf("上传文件:%s", err), m.W)
		return
	} else {
		dialog.ShowInformation("提示", fmt.Sprintf("上传成功"), m.W)
	}
}

func (m *UploadEntry) showEncryptInfo() {
	b, err := os.ReadFile(common.InputFile)
	if err != nil {
		dialog.ShowError(err, m.W)
		return
	}

	w := fyne.CurrentApp().NewWindow("加密信息:")
	w.SetContent(container.New(layout.NewCenterLayout(), widget.NewLabel(string(b))))
	w.Resize(fyne.Size{Height: 64, Width: 46})
	w.SetFixedSize(true)
	w.Show()
}

func (m *UploadEntry) showResultInfo(ret string) {
	w := fyne.CurrentApp().NewWindow("部署结果:")
	w.SetContent(container.New(layout.NewCenterLayout(), widget.NewLabel(ret)))
	w.SetFixedSize(true)
	w.Show()
}

func (m *UploadEntry) InitUpload() fyne.CanvasObject {
	return container.New(layout.NewVBoxLayout(), m.initUploadForm(), m.Infinite)
}

func uploadEntry2EncryptInfo(e *UploadEntry) cryptool.EncryptInfo {
	return cryptool.EncryptInfo{
		Pid:            e.Pid.Text,
		Uuid:           e.Uuid.Text,
		DeviceId:       e.DeviceId.Text,
		SecKey:         e.SecKey.Text,
		LocalKey:       e.LocalKey.Text,
		SubDeviceLimit: e.SubDeviceLimit.Text,
		SubDeviceTotal: e.SubDeviceTotal.Text,
		Env:            e.Env.Text,
		Region:         e.Region.Text,
		GwStatus:       e.GwStatus.Text,
		Saas:           e.Saas.Text,
	}
}

func (m *UploadEntry) uploadFile() error {
	info := ssh.GetSSHLoginInfo()
	src, dst := TarSrcPath, TarDstPath
	return util.ScpFile(src, dst, &info)
}
