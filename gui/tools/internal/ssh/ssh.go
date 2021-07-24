package ssh

import (
	"example.com/m/v2/common"
	"example.com/m/v2/util"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"regexp"
)

func (m *SSHEntry) initSshForm() fyne.CanvasObject {
	m.LoginData.Password = true

	var loginType string
	inLoginTypeSelect := widget.NewSelect([]string{common.PassWordType}, func(value string) {
		m.LoginType.Text = value
		if common.KeyType == value {
			loginType = KeyPath
			m.LoginData.Password = false
		} else {
			loginType = Password
			m.LoginData.Password = true
		}
	})
	inLoginTypeSelect.SetSelectedIndex(0)

	m.User.SetText("root")

	m.Port.SetText("22")
	m.Host.SetText("192.168.1.x")
	return &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: LoadType, Widget: m.LoginType},
			{Text: User, Widget: m.User},
			{Text: loginType, Widget: m.LoginData},
			{Text: Host, Widget: m.Host},
			{Text: Port, Widget: m.Port},
		},
		OnSubmit: func() { // optional, handle form submission
			if err := m.checkoutSSH(); err != nil {
				dialog.ShowInformation("错误", fmt.Sprintf("%s", err), m.W)
				return
			}
			if err := m.pingSsh(); err != nil {
				dialog.ShowInformation("错误", fmt.Sprintf("%s", err), m.W)
				return
			}
			dialog.ShowInformation("提示", "连接成功", m.W)
		},
		SubmitText: "ssh连接",
	}
}

func (data *SSHEntry) checkoutSSH() error {
	if data == nil {
		return fmt.Errorf("EncryptEntry is nil")
	}
	if data.User.Text == "" {
		return fmt.Errorf("错误:%s为空", User)
	}
	if data.LoginData.Text == "" {
		var loginStr string
		if data.LoginType.Text == common.PassWordType {
			loginStr = Password
		} else if data.LoginType.Text == common.KeyType {
			loginStr = KeyPath
		} else {
			loginStr = "data.LoginType.Text:" + data.LoginType.Text
		}
		return fmt.Errorf("错误:%s为空", loginStr)
	}
	if data.Host.Text == "" {
		return fmt.Errorf("错误:%s为空", Host)
	}
	if data.Host.Text <= "0" || data.Host.Text > "65535" {
		return fmt.Errorf("错误:%s大小超出范围", Host)
	}
	matched, err := regexp.MatchString(`^((2(5[0-5]|[0-4]\d))|[0-1]?\d{1,2})(\.((2(5[0-5]|[0-4]\d))|[0-1]?\d{1,2})){3}$`, data.Host.Text)
	if err != nil {
		return err
	}
	if !matched {
		return fmt.Errorf("错误:%s格式不对", Host)
	}
	if data.Port.Text == "" {
		return fmt.Errorf("错误:%s为空", Port)
	}
	return nil
}

func (m *SSHEntry) pingSsh() error {
	ls := "ls"

	cmd := make(map[string]string, 0)
	cmd[ls] = "ls"
	order := []string{ls}
	info := &common.SSHInfo{
		LoginData: m.LoginData.Text,
		Host:      m.Host.Text,
		Port:      m.Port.Text,
		User:      m.User.Text,
		LoginType: m.LoginType.Text,
		Cmd:       cmd,
		CmdOrder:  order,
	}
	_, err := util.SSH(info)
	if err != nil {
		return err
	}

	return nil
}

func (m *SSHEntry) InitSSH() fyne.CanvasObject {
	return container.New(layout.NewVBoxLayout(), m.initSshForm())
}

func GetSSHLoginInfo() common.SSHInfo {
	return common.SSHInfo{
		LoginData: s.LoginData.Text,
		Host:      s.Host.Text,
		Port:      s.Port.Text,
		User:      s.User.Text,
		LoginType: s.LoginType.Text,
	}
}
