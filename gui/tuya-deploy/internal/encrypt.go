package internal

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"step/gui/tuya-deploy/common"
	"step/gui/tuya-deploy/cryptool"
)

func (m *Manager) NewEncryptLabelC() *fyne.Container {
	pid := widget.NewLabel(common.Pid)
	inPid := widget.NewEntry()

	uuid := widget.NewLabel(common.Uuid)
	inUuid := widget.NewEntry()

	deviceId := widget.NewLabel(common.DeviceID)
	inDeviceId := widget.NewEntry()

	secKey := widget.NewLabel(common.SecKey)
	inSecKey := widget.NewEntry()

	localKey := widget.NewLabel(common.LocalKey)
	inLocalKey := widget.NewEntry()

	subDeviceLimit := widget.NewLabel(common.SubDeviceLimit)
	inLimit := widget.NewEntry()

	macAddr := widget.NewLabel(common.MacAddr)
	inMacAddr := widget.NewEntry()

	m.EncryptEntry = &common.EncryptEntry{
		Pid:            inPid,
		Uuid:           inUuid,
		DeviceId:       inDeviceId,
		SecKey:         inSecKey,
		LocalKey:       inLocalKey,
		MacAddr:        inMacAddr,
		SubDeviceLimit: inLimit,
	}

	return container.New(layout.NewFormLayout(),
		pid, inPid,
		uuid, inUuid,
		deviceId, inDeviceId,
		secKey, inSecKey,
		localKey, inLocalKey,
		subDeviceLimit, inLimit,
		macAddr, inMacAddr,
	)
}

func (m *Manager) NewEncryptButtonC() *fyne.Container {
	buttonFunc := func() {
		if err := checkoutEncrypt(m.EncryptEntry); err != nil {
			dialog.ShowInformation("错误", fmt.Sprintf("%s", err), m.W)
			return
		}
		if err := cryptool.CreateConfigFile(m.EncryptEntry); err != nil {
			dialog.ShowInformation("错误", fmt.Sprintf("创建yaml文件失败:%s", err), m.W)
			return
		}
		if err := cryptool.Encrypt(m.EncryptEntry.MacAddr.Text, common.Opt, common.InputFile, common.OutputFile); err != nil {
			dialog.ShowInformation("错误", fmt.Sprintf("加密失败:%s", err), m.W)
			return
		} else {
			dialog.ShowInformation("提示", "加密成功", m.W)
		}
	}

	return container.New(layout.NewCenterLayout(), widget.NewButton(common.ButtonEncrypt, buttonFunc))

}

func checkoutEncrypt(data *common.EncryptEntry) error {
	if data == nil {
		return fmt.Errorf("EncryptEntry is nil")
	}
	if data.Pid.Text == "" {
		return fmt.Errorf("错误:%s为空", common.Pid)
	}
	if data.Uuid.Text == "" {
		return fmt.Errorf("错误:%s为空", common.Uuid)
	}
	if data.DeviceId.Text == "" {
		return fmt.Errorf("错误:%s为空", common.DeviceID)
	}
	if data.SecKey.Text == "" {
		return fmt.Errorf("错误:%s为空", common.SecKey)
	}
	if data.LocalKey.Text == "" {
		return fmt.Errorf("错误:%s为空", common.LocalKey)
	}
	if data.SubDeviceLimit.Text == "" {
		return fmt.Errorf("错误:%s为空", common.SubDeviceLimit)
	}
	if data.MacAddr.Text == "" {
		return fmt.Errorf("错误:%s为空", common.MacAddr)
	}

	return nil
}

func encryptExecSsh(data *common.UploadEntry) error {
	cmd := make(map[common.CmdKey]string, 0)
	cmd[common.Kernel] = "uname -s"
	cmd[common.HwArch] = "uname -m"
	cmd[common.HwAddr] = "ifconfig -a |\nawk '/^[a-z]/ { iface=$1; mac=$NF; next }\n    /inet addr:/ { print iface, mac }'"

	order := []common.CmdKey{common.Kernel, common.HwArch, common.HwAddr}

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
	//edgex_v1.5.0_Darwin_x86_64.tar.gz

	fmt.Println(fileName)
	return nil
}
