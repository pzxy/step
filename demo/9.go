package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"step/grammar/codeskill/log"
)

func main() {
	s := fmt.Sprintf("arp %s", "192.168.18.186")
	cmd := exec.Command("bash", "-c", s)
	//in := bytes.NewBuffer(nil)
	//cmd.Stdin = in //绑定输入
	var out bytes.Buffer
	cmd.Stdout = &out //绑定输出
	//go func() {
	//	in.WriteString("node E:/design/test.js\n") //写入你的命令，可以有多行，"\n"表示回车
	//}()
	err := cmd.Run()
	if err != nil {
		log.ErrLog(err)
	}

	fmt.Println(out.String())

}
