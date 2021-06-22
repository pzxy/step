package util

import (
	"fmt"
	"net"
	"runtime"
	"step/gui/tuya-deploy/common"
)

func GetRemotelyMacAddrBySsh(info *common.SSHInfo) (string, error) {
	cmd := make(map[string]string, 0)
	var hwAddr = "hwAddr"
	switch runtime.GOOS {
	case common.Linux:
		cmd[hwAddr] = common.LinuxMacShell
	case common.Darwin:
		cmd[hwAddr] = common.DarwinMacShell
	default:
		return "", fmt.Errorf("%s操作系统的 GetMacAddr方法 尚未实现。", runtime.GOOS)
	}
	order := []string{hwAddr}
	info.Cmd = cmd
	info.CmdOrder = order

	hwMap, err := SSH(info)
	if err != nil {
		return "", err
	}
	return hwMap[hwAddr], nil
}

func GetLocalhostMacAddr() (string, error) {

	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, inter := range interfaces {
		hwAddr := inter.HardwareAddr.String()
		if hwAddr == "" {
			continue
		}
		return hwAddr, nil
	}

	return "", fmt.Errorf("not found mac addr")
}
