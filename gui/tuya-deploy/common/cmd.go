package common

import (
	"io"
	"os/exec"
	"runtime"
)

//fmt.Sprintf("scp -P %v %v %v@:~/", u.Port, u.User, u.DstPath)

func CMD(c *CMDInfo) error {
	switch runtime.GOOS {
	case Darwin:
		if e := execute(nil, "bash", "-c"); e != nil {
			return e
		}
	case Windows:
		//execute(nil, "cmd", "/c")
	case Linux:
		//execute(nil, "bash", "-c")
	}
	return nil
}

func execute(outPut io.Writer, command string, params ...string) error {
	cmd := exec.Command(command, params...)
	//cmd.Stdout = outPut
	//cmd.Stderr = outPut
	return cmd.Run()
}
