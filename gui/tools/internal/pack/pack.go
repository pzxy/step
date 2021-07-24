package pack

import (
	"example.com/m/v2/common"
	"example.com/m/v2/util"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"log"
	"os"
	"path"
	"strings"
)

func (m *PackEntry) initPackForm() fyne.CanvasObject {
	m.Infinite = widget.NewProgressBarInfinite()
	m.Infinite.Hide()

	ArchSelect := widget.NewSelect([]string{"x86_64", "armv7", "arm64"}, func(value string) {
		m.Arch.Text = value
	})
	ArchSelect.SetSelectedIndex(0)
	inOsSelect := widget.NewSelect([]string{"Linux", "Darwin"}, func(value string) {
		m.OS.Text = value
	})
	inOsSelect.SetSelectedIndex(0)
	inActiveSelect := widget.NewSelect([]string{"true", "false"}, func(s string) {
		if s == "true" {
			m.InitDirP = true
		} else {
			m.InitDirP = false
		}
	})
	inActiveSelect.SetSelectedIndex(0)
	inGwVersionSelect := widget.NewSelect(common.GwVersionList, func(value string) {
		v := common.GetGwVersion(value)
		m.GwVersion.SetText(v)
	})
	m.WebClient.SetPlaceHolder("全部部署为true时,可不填;默认部署网关版本,如网关为1.6 则部署1.6.0")
	m.EdgeHub.SetPlaceHolder("全部部署为true时,可不填;默认部署网关版本,如网关为1.6 则部署1.6.0")
	m.Resource.SetPlaceHolder("全部部署为true时,可不填;默认部署网关版本,如网关为1.6 则部署1.6.0")
	m.Expert.SetPlaceHolder("全部部署为true时,可不填;默认部署网关版本,如网关为1.6 则部署1.6.0")
	m.Agent.SetPlaceHolder("全部部署为true时,可不填;默认部署网关版本,如网关为1.6 则部署1.6.0")
	m.FileBeat.SetPlaceHolder("全部部署为true时,可不填;默认部署网关版本,如网关为1.6 则部署1.6.0")
	m.GateWay.SetPlaceHolder("全部部署为true时,可不填;默认部署网关版本,如网关为1.6 则部署1.6.0")
	m.TedgeUuid.SetPlaceHolder("当全部部署为true时,或tedge-filebeat已填时，此处需填写。")
	m.DockerRegistry.SetText(common.AliyunDockerAddr)
	inGwVersionSelect.SetSelectedIndex(2)
	inDockerAddrSelect := widget.NewSelect(common.DockerAddrList, func(s string) {
		m.DockerRegistry.SetText(s)
	})
	inDockerAddrSelect.SetSelectedIndex(0)

	return &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: HwArch, Widget: ArchSelect},
			{Text: Os, Widget: inOsSelect},
			{Text: common.DockerRegistry, Widget: inDockerAddrSelect},
			{Text: common.GwVersion, Widget: inGwVersionSelect},
			{Text: common.DeployAll, Widget: inActiveSelect},
			{Text: common.TedgeUuid, Widget: m.TedgeUuid},
			{Text: common.EdgeHub, Widget: m.EdgeHub},
			{Text: common.Resource, Widget: m.Resource},
			{Text: common.GateWay, Widget: m.GateWay},
			{Text: common.Expert, Widget: m.Expert},
			{Text: common.Agent, Widget: m.Agent},
			{Text: common.FileBeat, Widget: m.FileBeat},
			{Text: common.WebClient, Widget: m.WebClient},
		},
		OnSubmit: func() { // optional, handle form submission
			m.buildScriptButtonFunc()
		},
		SubmitText: "生成命令",
	}
}
func fileSaved(f fyne.URIWriteCloser, w fyne.Window) {
	defer f.Close()
	_, err := f.Write([]byte("Written by Fyne demo\n"))
	if err != nil {
		dialog.ShowError(err, w)
	}
	log.Println("Saved to...", f.URI())
}

func (m *PackEntry) buildScriptButtonFunc() {
	if m.InitDirP == true && strings.TrimSpace(m.TedgeUuid.Text) == "" && strings.TrimSpace(m.GwVersion.Text) >= "1.6.0" {
		dialog.ShowInformation("错误", fmt.Sprintf("错误：%s为true时，%s不能为空", common.DeployAll, common.TedgeUuid), m.W)
		return
	}
	if strings.TrimSpace(m.FileBeat.Text) != "" && strings.TrimSpace(m.TedgeUuid.Text) == "" && strings.TrimSpace(m.GwVersion.Text) >= "1.6.0" {
		dialog.ShowInformation("错误", fmt.Sprintf("错误：%strue时，%s不能为空", common.FileBeat, common.TedgeUuid), m.W)
		return
	}
	if err := m.buildScript(); err != nil {
		dialog.ShowInformation("错误", fmt.Sprintf("%s", err), m.W)
		return
	}
}

func (m *PackEntry) buildScript() error {
	var err error

	//edgex_v1.4.0_Darwin_x86_64.tar.gz
	fileName := strings.Join([]string{
		"edgex",
		"v" + m.GwVersion.Text,
		m.OS.Text,
		m.Arch.Text + ".tar.gz",
	}, "_")

	url := common.DownloadUrlPrefix + path.Join(m.GwVersion.Text, fileName)
	cmd := util.BuildDeployCmd(url, fileName, m.GwService)
	m.ScriptInfo.SetText(cmd)

	w := fyne.CurrentApp().NewWindow("部署脚本信息")

	button := widget.NewButton("ok", func() {
		m.ScriptInfo.Refresh()
		w.Close()
		fmt.Println(m.ScriptInfo.Text)
	})

	s := container.NewMax(m.ScriptInfo)
	bc := container.NewCenter(button)
	s.Resize(fyne.Size{720, 250})
	w.Resize(fyne.Size{720, 250})
	f := container.NewVSplit(m.ScriptInfo, bc)
	f.SetOffset(10)
	w.SetContent(f)
	w.Show()
	return err
}

func (m *PackEntry) packButtonFunc() {
	m.Infinite.Show()
	defer m.Infinite.Hide()
	if err := m.checkoutPackInfo(); err != nil {
		dialog.ShowInformation("错误", fmt.Sprintf("%s", err), m.W)
		return
	}
	if err := m.AssemblyPack(); err != nil {
		dialog.ShowInformation("错误", fmt.Sprintf("打包失败:%s", err), m.W)
		return
	} else {
		dialog.ShowInformation("提示", "生成成功", m.W)
	}
}

func (data *PackEntry) checkoutPackInfo() error {
	if data == nil {
		return fmt.Errorf("common.PackEntry is nil")
	}
	if data.Arch.Text == "" {
		return fmt.Errorf("错误:%s为空", HwArch)
	}
	if data.OS.Text == "" {
		return fmt.Errorf("错误:%s为空", Os)
	}
	if data.GwVersion.Text == "" {
		return fmt.Errorf("错误:%s为空", common.GwVersion)
	}
	if data.OutputDir.Text == "" {
		return fmt.Errorf("错误:打包输出，文件目录为空。")
	} else {
		fi, err := os.Stat(data.OutputDir.Text)
		if err != nil {
			return err
		}
		if !fi.IsDir() {
			return fmt.Errorf("错误:%s格式不对，这不是一个目录，或者没有此目录", PackOutputDir)
		}
	}
	return nil
}

func (data *PackEntry) AssemblyPack() error {
	var err error
	//edgex_v1.4.0_Darwin_x86_64.tar.gz
	info := []string{
		"edgex",
		"v" + data.GwVersion.Text,
		data.OS.Text,
		data.Arch.Text + ".tar.gz",
	}
	fileName := strings.Join(info, "_")
	filePath := common.TarPkg + fileName
	url := common.DownloadUrlPrefix + path.Join(data.GwVersion.Text, fileName)
	defer func() {
		//无论成功与否都删除下载下来的文件。
		os.Remove(filePath)
	}()
	if err = util.DownloadFile(filePath, url); err != nil {
		return err
	}

	if err = createScriptFile(PackDeployFileNamePath, fileName, data.GwVersion.Text); err != nil {
		return err
	}
	if err = createReadmeFile(PackReadmeFileNamePath); err != nil {
		return err
	}
	return util.CreateTarFile(path.Join(data.OutputDir.Text, PackTarName), common.TarPkg)
}

func createScriptFile(filePath, downloadFileName, version string) error {
	//./tedge init -s web-client -v 1.5.0
	content := []string{
		"#!/bin/bash \n",
		"tar -zxvf " + downloadFileName + "\n",
		"chmod +x tedge \n",
		"./tedge init -s " + common.EncryptFileName + " -v " + version + " \n",
		"./tedge init -s tedge-agent --agent-file=./tedge-agent --agent-executor-file=./tedge-agent-executor \n",
	}
	return createFile(filePath, content)
}

func createFile(filePath string, content []string) error {
	var err error
	defer func() {
		if err != nil {
			os.Remove(filePath)
		}
	}()
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0766)
	if err != nil {
		return err
	}
	for _, b := range content {
		_, err = f.Write([]byte(b))
		if err != nil {
			return err
		}
	}
	return nil
}

func createReadmeFile(filePath string) error {
	content := []string{
		"## 执行以下命令即可部署 \n",
		"```bash \n",
		"chmod +x deploy.sh \n",
		"./deploy.sh \n",
		"``` \n",
	}

	return createFile(filePath, content)
}

func (m *PackEntry) InitPack() fyne.CanvasObject {
	return container.New(layout.NewVBoxLayout(), m.initPackForm(), m.Infinite)
}
