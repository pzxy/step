package deploy

import (
	"example.com/m/v2/common"
	"example.com/m/v2/internal/ssh"
	"example.com/m/v2/util"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"os"
	"path"
	"strings"
)

func (m *DeployEntry) initDeployForm() fyne.CanvasObject {
	m.Infinite = widget.NewProgressBarInfinite()
	m.Infinite.Hide()
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
	inGwVersionSelect.SetSelectedIndex(2)
	inDockerAddrSelect := widget.NewSelect(common.DockerAddrList, func(s string) {
		m.DockerRegistry.SetText(s)
	})
	inDockerAddrSelect.SetSelectedIndex(0)
	m.DockerRegistry.SetText(common.AliyunDockerAddr)
	m.DockerPwd.Password = true
	return &widget.Form{
		Items: []*widget.FormItem{
			{Text: common.DockerRegistry, Widget: inDockerAddrSelect},
			{Text: DockerUser, Widget: m.DockerUser},
			{Text: DockerPwd, Widget: m.DockerPwd},
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
		OnSubmit: func() {
			m.deployButtonFunc()
		},
		SubmitText: "生成部署命令",
	}
}

func (m *DeployEntry) deployButtonFunc() {
	m.Infinite.Show()
	defer m.Infinite.Hide()
	//显示窗口
	if m.InitDirP == true && strings.TrimSpace(m.TedgeUuid.Text) == "" && strings.TrimSpace(m.GwVersion.Text) >= "1.6.0" {
		dialog.ShowInformation("错误", fmt.Sprintf("错误：%s为true时，%s不能为空", common.DeployAll, common.TedgeUuid), m.W)
		return
	}
	if strings.TrimSpace(m.FileBeat.Text) != "" && strings.TrimSpace(m.TedgeUuid.Text) == "" && strings.TrimSpace(m.GwVersion.Text) >= "1.6.0" {
		dialog.ShowInformation("错误", fmt.Sprintf("错误：%strue时，%s不能为空", common.FileBeat, common.TedgeUuid), m.W)
		return
	}

	c, err := m.showScriptInfo()
	if err != nil {
		dialog.ShowInformation("错误", fmt.Sprintf("%s", err), m.W)
		return
	}

	flag := <-c
	if flag == false {
		return
	}
	defer os.Remove(common.InputFile) //为了安全

	s, err := m.deploy()
	go m.showResultInfo(s)
	if err != nil {
		dialog.ShowInformation("错误", fmt.Sprintf("部署网关:%s", err), m.W)
	} else {
		dialog.ShowInformation("", fmt.Sprintf("部署成功"), m.W)
	}
}

func (m *DeployEntry) showScriptInfo() (chan bool, error) {
	var err error
	sshInfo := ssh.GetSSHLoginInfo()
	m.OS.Text, err = util.GetOS(&sshInfo)
	if err != nil {
		return nil, err
	}
	m.Arch.Text, err = util.GetArch(&sshInfo)
	if err != nil {
		return nil, err
	}
	//edgex_v1.4.0_Darwin_x86_64.tar.gz
	fileName := strings.Join([]string{
		"edgex",
		"v" + m.GwVersion.Text,
		m.OS.Text,
		m.Arch.Text + ".tar.gz",
	}, "_")

	url := common.DownloadUrlPrefix + path.Join(m.GwVersion.Text, fileName)
	if s, err := m.loginDocker(); err != nil {
		return nil, fmt.Errorf("镜像仓库的用户名或密码错误:%s", err)
	} else {
		fmt.Println("docker登陆信息:", s)
	}

	cmd := util.BuildDeployCmd(url, fileName, m.GwService)

	m.ScriptInfo.SetText(cmd)
	w := fyne.CurrentApp().NewWindow("部署脚本信息")
	c := make(chan bool)
	button := widget.NewButton("部署", func() {
		m.ScriptInfo.Refresh()
		w.Close()
		c <- true
	})
	cancel := widget.NewButton("取消", func() {
		w.Close()
		c <- false
	})

	s := container.NewMax(m.ScriptInfo)
	bc := container.NewHSplit(button, cancel)
	s.Resize(fyne.Size{720, 250})
	w.Resize(fyne.Size{720, 250})
	f := container.NewVSplit(m.ScriptInfo, bc)
	f.SetOffset(10)
	w.SetContent(f)
	w.Show()

	return c, nil
}

func (m *DeployEntry) showResultInfo(ret string) {
	w := fyne.CurrentApp().NewWindow("部署信息:")
	intro := widget.NewLabel("结果")
	intro.Wrapping = fyne.TextWrapBreak
	intro.SetText(ret)
	w.SetContent(container.NewVScroll(intro))
	w.Resize(fyne.NewSize(700, 400))
	w.Show()
}

func (m *DeployEntry) InitDeploy() fyne.CanvasObject {
	return container.New(layout.NewVBoxLayout(), m.initDeployForm(), m.Infinite)
}

func (m *DeployEntry) deploy() (string, error) {
	info := ssh.GetSSHLoginInfo()
	return util.DoSsh(m.ScriptInfo.Text, &info)
}

func (m *DeployEntry) loginDocker() (string, error) {
	user := strings.TrimSpace(m.DockerUser.Text)
	pwd := strings.TrimSpace(m.DockerPwd.Text)
	if user == "" || pwd == "" {
		return "", nil
	}
	info := ssh.GetSSHLoginInfo()
	cmd := "docker login -u " + user + " -p " + pwd + " " + m.DockerRegistry.Text
	return util.DoSsh(cmd, &info)
}

//func (m *DeployEntry) doPack() error {
//	var err error
//	info := ssh.GetSSHLoginInfo()
//	m.OS.Text, err = util.GetOS(&info)
//	if err != nil {
//		return err
//	}
//	m.Arch.Text, err = util.GetArch(&info)
//	if err != nil {
//		return err
//	}
//	dir := widget.NewEntry()
//	dir.Text = "./"
//
//	p := &pack.PackEntry{
//		W:         m.W,
//		Arch:      m.Arch,
//		OS:        m.OS,
//		GwVersion: m.GwVersion,
//		OutputDir: dir,
//	}
//	return p.AssemblyPack()
//}
