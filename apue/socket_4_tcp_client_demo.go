package main

import (
	"encoding/json"
	"fmt"
	"step/utils/log"
	"syscall"
)

func main() {
	socketClientDemo()
}

/**
由于客户端不需要固定的端口号，因此不必调用bind()，客户端的端口号由内核自动分
配。注意，客户端不是不允许调用bind()，只是没有必要调用bind()固定一个端口号，服务
器也不是必须调用bind()，但如果服务器不调用bind()，内核会自动给服务器分配监听端
口，每次启动服务器时端口号都不一样，客户端要连接服务器就会遇到麻烦。
客户端和服务器启动后可以查看链接情况：netstat -apn|grep 8000
*/
func socketClientDemo() {
	s := "hello WORLd"
	str, _ := json.Marshal(s)

	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {

		log.ErrLog(err)
		return
	}

	cliSockAddr := &syscall.SockaddrInet4{
		Addr: [4]byte{
			127, 0, 0, 1,
		},
		Port: 8890,
	}
	err = syscall.Connect(fd, cliSockAddr)
	if err != nil {
		log.ErrLog(err)
		return
	}

	syscall.Write(fd, str)
	buf := make([]byte, 128)

	n, err := syscall.Read(fd, buf)
	if err != nil {
		log.ErrLog(err)
		return
	}
	if n <= 0 { //读到零，说明收到fin，小与-1则错误
		fmt.Printf("client read n:%d \n", n)
		syscall.Close(fd)
	}

	//	syscall.Write(syscall.Stdout, buf)
	fmt.Printf("%s", buf[:n])
	syscall.Close(fd)

	return

}

func socketClientDemo2() {
	s := "hello WORLd"
	str, _ := json.Marshal(s)

	fd, _ := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)

	cliSockAddr := &syscall.SockaddrInet4{
		Addr: [4]byte{
			127, 0, 0, 1,
		},
		Port: 8890,
	}

	syscall.Connect(fd, cliSockAddr)

	syscall.Write(fd, str)
	buf := make([]byte, 128)
	n, _ := syscall.Read(fd, buf)
	fmt.Printf("client read n:%d \n", n)
	fmt.Printf("%s", buf[:n])

	syscall.Close(fd)

	return

}
