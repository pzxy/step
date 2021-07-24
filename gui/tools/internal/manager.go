package internal

import (
	"example.com/m/v2/common"
	"example.com/m/v2/internal/deploy"
	"example.com/m/v2/internal/encrypt"
	"example.com/m/v2/internal/pack"
	"example.com/m/v2/internal/ssh"
	"example.com/m/v2/internal/upload"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"os"
)

func (m *Manager) Init() {
	m.InitPkgDir()
	m.InitInstructions()
	m.InitManual()
	m.InitAuto()
}

func (m *Manager) InitPkgDir() {
	err := os.RemoveAll(common.TarPkg)
	if err != nil {
		dialog.ShowInformation("错误", fmt.Sprintf("%s", err), m.W)
		return
	}
	err = os.Mkdir(common.TarPkg, 0766)
	if err != nil {
		dialog.ShowInformation("错误", fmt.Sprintf("%s", err), m.W)
		return
	}
}

func (m *Manager) InitInstructions() {
	txt :=
		"方式一: 手动部署\n" +
			"第一步、 网关加密:按要求填写，文件名默认为 encrypted ，将文件上传到目的主机 /tmp/ 目录下。\n" +
			"第二步、 生成命令:选择要部署的网关信息，生成部署命令，将命令复制到目的主机执行即可。\n" +
			"\n" +
			"方式二: 自动部署 \n" +
			"第一步、 SSH连接:通过SSH连接目的主机。目前只支持密码登陆。 \n" +
			"第二步、 上传加密:按要求填写，文件会通过ssh连接，自动上传到目的主机 /tmp/ 目录下。\n" +
			"第三步、 网关部署:选择要部署的网关信息，生成命令，会弹出生成命令的窗口(此时可根据需要修改部署命令)。确认无误后，点击部署。等待部署结果即可。\n" +
			"\n" +
			"注意:\n" +
			"部署全部服务：全部部署设为true，初始化目录为true时，会读取加密激活文件，然后部署全部服务。\n" +
			"部署单个服务：全部部署设为false，填入对应服务版本号。部署单个服务时，无需再次上传加密文件。"

	title := widget.NewLabel("使用说明")
	intro := widget.NewLabel("介绍")
	intro.Wrapping = fyne.TextWrapBreak
	intro.SetText(txt)
	m.Instructions = container.New(layout.NewVBoxLayout(), title, widget.NewSeparator(), intro)
}

func (m *Manager) InitManual() {

	encrypt := container.New(layout.NewVBoxLayout(), widget.NewSeparator(), encrypt.NewEncrypt(m.W).InitEncrypt())
	edcode := container.New(layout.NewVBoxLayout(), widget.NewSeparator(), pack.NewPackEntry(m.W).InitPack())

	tabs := container.NewAppTabs(
		container.NewTabItem("网关加密", encrypt),
		container.NewTabItem("生成命令", edcode),
	)
	tabs.SetTabLocation(container.TabLocationTop)

	m.ManualDeploy = container.New(layout.NewVBoxLayout(), tabs)
}

func (m *Manager) InitAuto() {
	sshLogin := container.New(layout.NewVBoxLayout(), widget.NewSeparator(), ssh.NewSSHEntry(m.W).InitSSH())
	deploy := container.New(layout.NewVBoxLayout(), widget.NewSeparator(), deploy.NewDeploy(m.W).InitDeploy())
	upload := container.New(layout.NewVBoxLayout(), widget.NewSeparator(), upload.NewUpload(m.W).InitUpload())

	tabs := container.NewAppTabs(
		container.NewTabItem("SSH连接", sshLogin),
		container.NewTabItem("网关加密", upload),
		container.NewTabItem("网关部署", deploy),
	)
	tabs.SetTabLocation(container.TabLocationTop)

	m.AutoDeploy = container.New(layout.NewVBoxLayout(), tabs)
}
