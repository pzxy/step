package main

import (
	"fmt"
	"os"
	"step/grammar/codeskill/log"
)

func main() {
	forkDemo()
}

//fork这里就跳过吧，go里面不知道谁实现呢，fork的时候子进程会返回给父进程自己的pid号。
//除了pcb中的数据，其他父进程的数据和子进程的都一样。
func forkDemo() { //这是开启一个进程，不是fork
	procAttr := &os.ProcAttr{}
	var argv []string
	p, err := os.StartProcess("demo2", argv, procAttr)
	if err != nil {
		log.ErrLog(err)
		return
	}
	fmt.Printf("proAttr:%+v \n", procAttr)
	fmt.Printf("pid:%v  \n", p.Pid)
}

/**
vfork()这个函数fork函数的时候不会马上复制原来父进程的，因为开销太大，现在已经废弃了。
*/
