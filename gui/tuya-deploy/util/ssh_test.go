package util

import (
	"fmt"
	"step/grammar/codeskill/log"
	"step/gui/tuya-deploy/common"
	"testing"
)

func Test_encryptExecSsh(t *testing.T) {
	info := &common.SSHInfo{
		PassWord:  "root",
		Host:      "192.168.1.239",
		Port:      "22",
		User:      "root",
		LoginType: common.PassWord,
	}
	string, err := GetRemotelyMacAddrBySsh(info)
	if err != nil {
		log.ErrLog(err)
	}
	fmt.Println(string)
}

func TestGetMacAddrByLocalhost(t *testing.T) {
	fmt.Println(GetLocalhostMacAddr())
}
