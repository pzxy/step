package common

import (
	"fmt"
	"step/grammar/codeskill/log"
	"testing"
)

func Test_encryptExecSsh(t *testing.T) {
	info := &SSHInfo{
		PassWord:  "root",
		Host:      "192.168.1.239",
		Port:      "22",
		User:      "root",
		LoginType: PassWord,
	}
	string, err := GetMacAddrBySSH(info)
	if err != nil {
		log.ErrLog(err)
	}
	fmt.Println(string)
}

func TestGetMacAddrByLocalhost(t *testing.T) {
	GetMacAddrByLocalhost()
}
