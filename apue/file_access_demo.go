package main

import (
	"golang.org/x/sys/unix"
	"step/grammar/codeskill/log"
	"syscall"
)

func main() {
	accessDemo()
}

/**
access 测试文件的是否可读可写可执行，是否存在。
access 是以实际用户来访问的，比如用户pzxy创建的文件，可以使用sudo获取root权限访问，也可以直接访问
但是使用access的话，默认就是pzxy这个实际用户来访问的。
*/
func accessDemo() {
	//R_ok,W_OK,X_OK,F_OK 可读 可写 可执行 文件是否存在

	//只要报错就说明不符合测试条件
	err := syscall.Access("./test.txt", unix.W_OK) //go中没有定义F_OK，不过看它定义的其他值，推断应为0x03,相比open出发文件是否存在，这个轻量级一些
	if err != nil {
		log.ErrLog(err)
		return
	}
	//chmod 555 test.txt将文件改为不可写，继续执行上面函数，会得到 permission denied

}
