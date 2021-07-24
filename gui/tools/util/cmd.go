package util

import (
	"bytes"
	"example.com/m/v2/common"
	"fmt"
	"os/exec"
	"runtime"
)

//fmt.Sprintf("scp -P %v %v %v@:~/", u.Port, u.User, u.DstPath)

func CMD(c *common.CMDInfo) (map[string]string, error) {
	if c == nil {
		return nil, fmt.Errorf("common.CMDInfo is nil")
	}
	var command, param string
	switch runtime.GOOS {
	case common.Darwin:
		command = "bash"
		param = "-c"
	case common.Windows:
		command = "bash"
		param = "/c"
	case common.Linux:
		command = "bash"
		param = "-c"
	default:
		return nil, fmt.Errorf("尚不支持此%s操作系统。", runtime.GOOS)
	}

	var err error
	var ret = make(map[string]string)
	for _, this := range c.CmdOrder {
		cmd := exec.Command(command, param, c.Cmd[this])
		cmd.Stdout = &bytes.Buffer{}
		err = cmd.Run()
		if err != nil {
			return nil, err
		}
		ret[this] = cmd.Stdout.(*bytes.Buffer).String()
	}
	return ret, nil
}
