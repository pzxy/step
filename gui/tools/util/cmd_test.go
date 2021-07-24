package util

import (
	"example.com/m/v2/common"
	"fmt"
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
		fmt.Println(err)
	}
	for _, s := range m2 {
		fmt.Println(s)
	}
}
