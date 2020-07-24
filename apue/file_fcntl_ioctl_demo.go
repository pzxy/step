package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/sys/unix"
	"step/utils/log"
	"syscall"
	"time"
)

func main() {
	ioctlDemo()
}

func noFcntl() {
	fmt.Println("阻塞")
	b := make([]byte, 128)
	n, err := syscall.Read(syscall.Stdin, b)
	if err != nil {
		log.ErrLog(err)
		return
	}
	if n < 0 {
		return
	}

	syscall.Write(syscall.Stdout, b)
}

func fcntlDemo() {
	b, _ := json.Marshal("asdasdasappen")
	fd, _ := syscall.Open("./test3.txt", syscall.O_RDWR, 0655)
	flags, err := unix.FcntlInt(uintptr(fd), unix.F_GETFL, unix.O_CREAT) //unix.F_GETFL unix.F_SETFL, 修改和设置文件的访问设置属性

	if err != nil {
		fmt.Printf("err1:%v \n", err)
		return
	}
	flags |= unix.O_APPEND //尝试改变这个值来观察 文件结构体中的 flag的改变
	unix.FcntlInt(uintptr(fd), unix.F_SETFL, flags)
	//if err != nil {
	//	fmt.Printf("err2:%v \n", err)
	//	return
	//}
	//n, err := syscall.Read(int(file.Fd()), b)
	//if err != nil {
	//	fmt.Printf("err:%v \n", err)
	//	return
	//}
	//if n < 0 {
	//	return
	//}

	syscall.Write(fd, b)
	syscall.Close(fd)
}

func ioctlDemo() {
	//ioctl是控制文件物理属性的函数。比如文件的flag描述了文件的读写属性，这是文件的共性。但是不同的文件可能代表不同的设备。ioctl这个
	//个函数就是专门提供用来获取驱动文件的中属性的函数。但是不同的文件，属性可能不同，所以要根据驱动的源码来做操作才行。
	fd, _ := syscall.Open("/dev/ttys001", syscall.O_RDONLY, 0777) //ttys004是终端窗口的文件，通过 who命令可以看到

	w, err := unix.IoctlGetWinsize(fd, unix.TIOCGWINSZ)
	if err != nil {
		log.ErrLog(err)
		return
	}
	//
	//v, err := unix.IoctlGetInt(fd, unix.TIOCGWINSZ)
	//if err != nil {
	//	log.ErrLog(err)
	//	return
	//
	//}
	//var w unix.Winsize
	//&w = *v

	fmt.Printf("行：%v,列：%v \n", w.Row, w.Col)
	time.Sleep(time.Second * 5)
	w.Col = 30

	unix.IoctlSetWinsize(fd, unix.TIOCGWINSZ, w) //改不掉

	ww, _ := unix.IoctlGetWinsize(fd, unix.TIOCGWINSZ)

	fmt.Printf("行：%v,列：%v \n", ww.Row, ww.Col)
	time.Sleep(time.Second * 5)
	//再次打开才能看到效果
}
