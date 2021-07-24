package encrypt

import (
	"example.com/m/v2/common"
	"example.com/m/v2/cryptool"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"io"
	"os"
	"path"
	"regexp"
	"runtime"
)

func (data *EncryptEntry) checkoutEncrypt() error {
	if data == nil {
		return fmt.Errorf("EncryptEntry is nil")
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

	if data.OutputDir.Text != "" {
		data.OutputDir.Text = path.Join(data.OutputDir.Text, EncryptFileName)
	}

	if data.DefaultOutputDir.Text == "" {
		data.DefaultOutputDir.Text = OutputFile
	} else {
		data.DefaultOutputDir.Text = path.Join(data.DefaultOutputDir.Text, EncryptFileName)
	}

	return nil
}

func (m *EncryptEntry) initEncryptForm() fyne.CanvasObject {

	m.MacAddr.SetPlaceHolder("格式 83:87:d3:61:e0:01")
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
	switch runtime.GOOS {
	case common.Windows:
		m.OutputDir.Text = common.PackWindowsDefaultOutputDir
	case common.Darwin, common.Linux:
		m.OutputDir.Text = common.PackUnixDefaultOutputDir
	default:
		dialog.ShowError(fmt.Errorf("未知操作系统：%s", runtime.GOOS), m.W)
	}
	wb := widget.NewButton("选择目录", func() {
		dialog.ShowFolderOpen(func(list fyne.ListableURI, err error) {
			if err != nil {
				dialog.ShowError(err, m.W)
				return
			}
			if list == nil {
				return
			}
			m.OutputDir.Text = list.Path()
			m.OutputDir.Refresh()
		}, m.W)
	})
	c := container.NewHSplit(m.OutputDir, wb)
	c.SetOffset(10)
	inOutputDirCont := c
	//inDefaultOutputDir.Validator = validation.NewRegexp(`^[A-Za-z0-9_-]+$`, "username can only contain letters, numbers, '_', and '-'")
	return &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: DeviceID, Widget: m.DeviceId},
			{Text: SecKey, Widget: m.SecKey},
			{Text: LocalKey, Widget: m.LocalKey},
			{Text: Uuid, Widget: m.Uuid},
			{Text: Pid, Widget: m.Pid},
			{Text: SubDeviceLimit, Widget: m.SubDeviceLimit},
			{Text: MacAddr, Widget: m.MacAddr},
			{Text: Region, Widget: inRegionSelect},
			{Text: Env, Widget: inEnvSelect},
			{Text: Saas, Widget: inSassSelect},
			{Text: OutputDir, Widget: inOutputDirCont},

			//{Text: common.DefaultOutputDir, Widget: inDefaultOutputDir},
		},
		OnSubmit: func() { // optional, handle form submission
			m.encryptButtonFunc()
		},
		SubmitText: "加密",
	}
}

func (m *EncryptEntry) encryptButtonFunc() {
	defer os.Remove(common.InputFile) //为了安全
	if err := m.checkoutEncrypt(); err != nil {
		dialog.ShowInformation("错误", fmt.Sprintf("%s", err), m.W)
		return
	}
	e := encryptEntry2EncryptInfo(m)
	if err := cryptool.CreateConfigFile(&e); err != nil {
		dialog.ShowInformation("错误", fmt.Sprintf("创建yaml文件失败:%s", err), m.W)
		return
	} else {
		go m.showEncryptInfo()
	}

	if err := cryptool.Encrypt(m.MacAddr.Text, false, common.InputFile, m.DefaultOutputDir.Text); err != nil {
		dialog.ShowInformation("错误", fmt.Sprintf("加密失败:%s", err), m.W)
		m.DefaultOutputDir.Text = "" //为了页面不显示这个值
		return
	} else {
		if _, err := copyFile(m.DefaultOutputDir.Text, m.OutputDir.Text); err != nil {
			dialog.ShowInformation("错误", fmt.Sprintf("文件导出失败:%s", err), m.W)
		}
		m.DefaultOutputDir.Text = "" //为了页面不显示这个值
		dialog.ShowInformation("提示", "加密成功", m.W)
	}
}

func copyFile(src, dst string) (int64, error) {
	if dst == "" {
		return 0, nil
	}
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}
	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}
	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func (m *EncryptEntry) showEncryptInfo() {
	b, err := os.ReadFile(common.InputFile)
	if err != nil {
		dialog.ShowError(err, m.W)
		return
	}

	w := fyne.CurrentApp().NewWindow("加密信息")
	w.SetContent(container.New(layout.NewCenterLayout(), widget.NewLabel(string(b))))
	w.Resize(fyne.Size{Height: 64, Width: 46})
	w.SetFixedSize(true)
	w.Show()
}

func (m *EncryptEntry) InitEncrypt() fyne.CanvasObject {
	return container.New(layout.NewVBoxLayout(), m.initEncryptForm())
}

func encryptEntry2EncryptInfo(e *EncryptEntry) cryptool.EncryptInfo {
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
