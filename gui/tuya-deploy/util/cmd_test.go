package util

import (
	"fmt"
	"step/grammar/codeskill/log"
	"step/gui/tuya-deploy/common"
	"testing"
)

func TestCMD(t *testing.T) {
	var ls = "ls"
	var pwd = "pwd"
	var tree = "tree"
	m := make(map[string]string)
	m[ls] = " ls -l"
	m[pwd] = " pwd"
	m[tree] = " tree"
	order := []string{pwd, ls, tree}

	cmd := &common.CMDInfo{
		Cmd:      m,
		CmdOrder: order,
	}
	m2, err := CMD(cmd)
	if err != nil {
		log.ErrLog(err)
	}
	for _, s := range m2 {
		fmt.Println(s)
	}
}
