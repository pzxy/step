package util

import (
	"example.com/m/v2/common"
	"fmt"
	"log"
	"testing"
)

func TestGetMacAddrByLocalhost(t *testing.T) {
	fmt.Println(GetLocalhostMacAddr())
}

func TestGetOS(t *testing.T) {
	info := &common.SSHInfo{
		LoginData: "linkeding",
		Host:      "192.168.1.177",
		Port:      "22",
		User:      "root",
		LoginType: common.PassWordType,
	}
	string, err := GetOS(info)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(string)
}

func TestGetArch(t *testing.T) {
	info := &common.SSHInfo{
		LoginData: "linkeding",
		Host:      "192.168.1.177",
		Port:      "22",
		User:      "root",
		LoginType: common.PassWordType,
	}
	string, err := GetArch(info)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(string)
}

func TestGetOsVersion(t *testing.T) {
	info := &common.SSHInfo{
		LoginData: "root",
		Host:      "192.168.1.239",
		Port:      "22",
		User:      "root",
		LoginType: common.PassWordType,
	}
	m, err := GetOsVersion(info)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(m)
}

func TestGetMacAddr(t *testing.T) {
	info := &common.SSHInfo{
		LoginData: "linkeding",
		Host:      "192.168.1.177",
		Port:      "22",
		User:      "root",
		LoginType: common.PassWordType,
	}
	m, err := GetOsMacAddr(info)
	for s, s2 := range m {
		fmt.Println(s, s2)
	}
	if err != nil {
		log.Print(err)
		return
	}
}

func TestScpFile(t *testing.T) {
	info := &common.SSHInfo{
		LoginData: "linkeding",
		Host:      "192.168.1.177",
		Port:      "22",
		User:      "root",
		LoginType: common.PassWordType,
	}
	src := "/Users/pzxy/WorkSpace/golang/tedge/lib/tedge-tools/tedge/deploy.sh"
	dst := "/tmp/deploy.sh"
	err := ScpFile(src, dst, info)
	if err != nil {
		log.Print(err)
		return
	} else {
		fmt.Println("成功")
	}
}

func TestDoSsh(t *testing.T) {
	info := &common.SSHInfo{
		LoginData: "linkeding",
		Host:      "192.168.1.177",
		Port:      "22",
		User:      "root",
		LoginType: common.PassWordType,
	}
	cmd := "tar -zxvf /tmp/edgex_v1.5.0_Linux_x86_64.tar.gz \n" +
		"cd /tmp/ \n" +
		"pwd \n"
	s, err := DoSsh(cmd, info)
	if err != nil {
		log.Print(err)
	} else {
		fmt.Println(s)
	}
}
