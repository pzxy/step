package util

import (
	"archive/tar"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
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

func DownloadFile(fileOutputPath, url string) error {
	f, err := os.OpenFile(fileOutputPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func CreateTarFile(fileName, directory string) error {
	if directory == "" {
		return fmt.Errorf("文件目录为空")

	}
	if fileName == "" {
		fileName = filepath.Base(directory)
	}
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := tar.NewWriter(f)
	defer writer.Close()

	filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			fmt.Println(err)
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
				fmt.Println(err)
				return nil
			}
		}
		return nil

	})
	return nil
}
