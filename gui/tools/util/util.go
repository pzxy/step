package util

import (
	"archive/tar"
	"example.com/m/v2/common"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

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

func DownloadFile(fileOutputPath, url string) error {
	var err error
	defer func() {
		if err != nil {
			os.Remove(fileOutputPath)
		}
	}()
	f, err := os.OpenFile(fileOutputPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0766)
	if err != nil {
		return err
	}
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	//if resp.ContentLength <= 0 {
	//	return fmt.Errorf("错误：无法下载文件，下载路径为：%s", url)
	//}
	defer resp.Body.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func CreateTarFile(tarFilePath, tarSourceDirPath string) error {
	var err error
	defer func() {
		if err != nil {
			os.Remove(tarFilePath)
		}
	}()
	if tarSourceDirPath == "" {
		return fmt.Errorf("文件目录为空")

	}
	if tarFilePath == "" {
		tarFilePath = filepath.Base(tarSourceDirPath)
	}
	os.Remove(tarFilePath)
	f, err := os.OpenFile(tarFilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0766)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := tar.NewWriter(f)
	defer writer.Close()

	err = filepath.Walk(tarSourceDirPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}

		hr := &tar.Header{
			Name: path,
			Size: info.Size(),
			Mode: 0666,
		}

		writer.WriteHeader(hr)
		var buff [1024]byte

		for {
			n, err := f.Read(buff[:])
			writer.Write(buff[:n])
			if err != nil {
				if err == io.EOF {
					break
				}
				return err
			}
		}
		return nil

	})
	if err != nil {
		return err
	}
	return nil
}

//uname -s

func GetOS(info *common.SSHInfo) (string, error) {
	cmd := make(map[string]string, 0)
	os := "os"

	cmd[os] = "uname -s"

	info.Cmd = cmd
	info.CmdOrder = []string{os}

	hwMap, err := SSH(info)
	if err != nil {
		return "", err
	}
	return strings.Title(strings.Replace(hwMap[os], "\n", "", -1)), nil
}

// 查看是Ubuntu还是Centos

func GetOsVersion(info *common.SSHInfo) (string, error) {
	cmd := make(map[string]string, 0)
	redhat := "Redhat"
	centos := "Centos"
	ubuntu := "Ubuntu"
	debian := "Debian"
	cmd[redhat] = "cat /etc/redhat-release"
	cmd[centos] = "cat /etc/centos-release"
	cmd[ubuntu] = "cat /etc/lsb-release"
	cmd[debian] = "cat /etc/debian_version"

	for k, v := range cmd {
		info.Cmd = map[string]string{k: v}
		info.CmdOrder = []string{k}
		hwMap, err := SSH(info)
		if len(hwMap[k]) == 0 || err != nil {
			continue
		}
		switch k {
		case redhat, centos:
			return centos, nil
		case ubuntu, debian:
			return ubuntu, nil
		}
	}
	return "", fmt.Errorf("不支持此系统")
}

//uname -m

func GetArch(info *common.SSHInfo) (string, error) {
	cmd := make(map[string]string, 0)
	arch := "arch"
	cmd[arch] = "uname -m"
	info.Cmd = cmd
	info.CmdOrder = []string{arch}

	hwMap, err := SSH(info)
	if err != nil {
		return "", err
	}
	s := strings.Replace(hwMap[arch], "\n", "", -1)
	if s == "armv7l" {
		s = "armv7"
	}
	return s, nil
}

func GetOsMacAddr(info *common.SSHInfo) ([]string, error) {
	ret := make([]string, 0)
	cmd := make(map[string]string, 0)
	macAddr := "mac"
	cmd[macAddr] = "ifconfig"
	info.Cmd = cmd
	info.CmdOrder = []string{macAddr}

	hwMap, err := SSH(info)
	if err != nil {
		return ret, err
	}
	m := hwMap[macAddr]
	//'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'
	//`^([0-9a-fA-F]{2})(([/\\s:][0-9a-fA-F]{2}){5})$`
	//.*([0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5})
	//nameRe := regexp.MustCompile(`(\w*?): `)
	//macNames := nameRe.FindAllStringSubmatch(m, -1)
	addrRe := regexp.MustCompile(`([0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5})`)
	macAddrs := addrRe.FindAllStringSubmatch(m, -1)
	//
	//if len(macAddrs) == 0 || len(macNames) == 0 {
	//	return ret, fmt.Errorf("can't find mac addr")
	//}
	//var count int
	//if len(macNames) >= len(macAddrs) {
	//	count = len(macAddrs)
	//} else {
	//	count = len(macNames)
	//}
	//for i := 0; i < count; i++ {
	//	ret[macNames[i][1]] = macAddrs[i][0]
	//}

	for _, v := range macAddrs {
		ret = append(ret, v[0])
	}
	return ret, nil
}

func BuildDeployCmd(url, fileName string, m common.GwService) string {
	var dockerRegistry string
	var dockerLogin string
	if strings.TrimSpace(m.DockerRegistry.Text) != "" {
		dockerRegistry = " --harbor-addr=" + m.DockerRegistry.Text + " "
		//dockerLogin = "&& docker login " + m.DockerRegistry.Text + " \\\n "
	}
	var tedgeUUid string
	if strings.TrimSpace(m.GwVersion.Text) >= "1.6.0" {
		tedgeUUid = " --tedge-uuid=" + strings.TrimSpace(m.TedgeUuid.Text) + " "
	}
	end := " ' \n"
	cmd := "bash -c 'cd /tmp \\\n " +
		"&& systemctl start docker  \\\n " +
		dockerLogin +
		"&& wget " + url + " -O " + fileName + " \\\n " +
		"&& tar -zxvf " + fileName + " \\\n "

	if m.InitDirP {
		cmd += "&& ./tedge init -c encrypted -v " + m.GwVersion.Text + tedgeUUid + dockerRegistry + " \\\n "
	}

	if strings.TrimSpace(m.EdgeHub.Text) != "" {
		cmd += "&& ./tedge init -s " + common.EdgeHub + " -v " + m.EdgeHub.Text + dockerRegistry + " \\\n "
	}
	if strings.TrimSpace(m.Resource.Text) != "" {
		cmd += "&& ./tedge init -s " + common.Resource + " -v " + m.Resource.Text + dockerRegistry + " \\\n "
	}
	if strings.TrimSpace(m.GateWay.Text) != "" {
		cmd += "&& ./tedge init -s " + common.GateWay + " -v " + m.GateWay.Text + dockerRegistry + " \\\n "
	}
	if strings.TrimSpace(m.Expert.Text) != "" {
		cmd += "&& ./tedge init -s " + common.Expert + " -v " + m.Expert.Text + dockerRegistry + " \\\n "
	}
	if strings.TrimSpace(m.FileBeat.Text) != "" {
		cmd += "&& ./tedge init -s " + common.FileBeat + " -v " + m.FileBeat.Text + dockerRegistry + " \\\n "
	}
	if strings.TrimSpace(m.Agent.Text) != "" {
		cmd += "&& ./tedge init -s " + common.Agent + " -v " + m.Agent.Text + dockerRegistry + " \\\n "
	}
	if strings.TrimSpace(m.WebClient.Text) != "" {
		cmd += "&& ./tedge init -s " + common.WebClient + " -v " + m.WebClient.Text + dockerRegistry + " \\\n "
	}
	cmd += "&& ./tedge init -s tedge-agent --agent-file=./tedge-agent --agent-executor-file=./tedge-agent-executor \\\n"
	cmd += end
	return cmd
}
