package main

import (
	"encoding/json"
	"fmt"
	"os"
	"step/utils/log"
	"syscall"
	"time"
)

func main() {
	syscall4SeekDemo()
}
func os4aOpenFile() {
	f, err := os.OpenFile("./test4.txt", os.O_RDWR|os.O_APPEND|os.O_EXCL, 0655) //里面也是调用了 syscallOpen
	if err != nil {
		log.ErrLog(err)
	}
	fmt.Printf("os.OpenFile fd:%v \n", f.Fd())

	b, _ := json.Marshal(" os open file ")
	f.Write(b) //与 syscall相比，参数少了文件描述符，因为文件描符就在 f.fd()中。
	f.Close()

}

func syscall4OpenFile() {

	//fd2, _ := syscall.Open("./test4.txt", syscall.O_CREAT, 0777)
	//fmt.Printf("syscall.Open fd:%v \n", fd2)
	b, _ := json.Marshal(" syscall open file ")

	//fd, _ := syscall.Open("./test4.txt", syscall.O_RDWR|syscall.O_APPEND, 0655)
	fd, _ := syscall.Open("./test4.txt", syscall.O_RDWR|syscall.O_APPEND|syscall.O_EXCL, 0655)
	//syscall.O_EXCL 和 create一起使用，如果文件不存在的话就创建文件，如果文件已经存在就返回-1 ，然后退出。
	//syscall.O_EXCL 和读一起使用，如果文件存在的话，就写进去，如果文件不存在的话，就返回-1。
	//如果文件不存在的话，对文件的操作都会返回-1，因为这个时候没有这个文件。而excl是相反的，只有文件存在的时候才返回-1
	fmt.Printf("syscall.Open fd:%v \n", fd)

	syscall.Write(fd, b) //这是系统调用，算是与c比较接近的吧，但是少了一个制定缓冲区大小的参数。

	syscall.Close(fd)
}

//阻塞读终端
func syscall4ReadFile1() {
	fmt.Println("这是阻塞式输入，请输入一些数据")

	b := make([]byte, 512)
	n, err := syscall.Read(syscall.Stdin, b)
	if err != nil {
		log.ErrLog(err)
		return
	}
	if n < 0 {
		return
	}

	syscall.Write(syscall.Stdout, b)
	return
}
func syscall4ReadFile2() {
	fmt.Println("这是阻塞式输入，请输入一些数据")
	b := make([]byte, 1024)
	n, err := fmt.Scanln(&b)

	if err != nil {
		log.ErrLog(err)
		return
	}
	if n < 0 {
		return
	}
	return
}

/**
非阻塞读取终端
*/
func readFileNoBlock() {
	fmt.Println("这是非阻塞读取终端...")
	fd, _ := syscall.Open("/dev/ttys002", syscall.O_RDONLY|syscall.O_NONBLOCK, 0777)
	if fd < 0 {
		fmt.Printf("fd:%v \n", fd)
		return
	}
	b := make([]byte, 20)
cc:
	n, err := syscall.Read(fd, b)
	if n < 0 {
		if err == syscall.EAGAIN {
			time.Sleep(time.Second)
			syscall.Write(syscall.Stdout, b)
			fmt.Println()
			goto cc

		}
		return
	}
	syscall.Write(syscall.Stdout, b)

	return

}

/**
seek
*/
func os4SeekDemo() {
	file, err := os.OpenFile("./seek.txt", os.O_RDWR, 0777)
	if err != nil {
		log.ErrLog(err)
	}
	file.Seek(10, os.SEEK_END)
	b, _ := json.Marshal("ppppppppp")
	file.Write(b)
	file.Close()
}
func syscall4SeekDemo() {
	fd, err := syscall.Open("./seek.txt", os.O_RDWR, 0777)
	if err != nil {
		log.ErrLog(err)
	}
	syscall.Seek(fd, 0, os.SEEK_SET)
	b, _ := json.Marshal("文件首行会覆盖，会覆盖后面的")
	syscall.Write(fd, b)

}
