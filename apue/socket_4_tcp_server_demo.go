package main

import (
	"bytes"
	"fmt"
	"step/grammar/codeskill/log"
	"syscall"
)

func main() {
	socketServerDemo()
}

/**
由于客户端不需要固定的端口号，因此不必调用bind()，客户端的端口号由内核自动分
配。注意，客户端不是不允许调用bind()，只是没有必要调用bind()固定一个端口号，服务
器也不是必须调用bind()，但如果服务器不调用bind()，内核会自动给服务器分配监听端
口，每次启动服务器时端口号都不一样，客户端要连接服务器就会遇到麻烦。
客户端和服务器启动后可以查看链接情况：netstat -apn|grep 8000
*/
func socketServerDemo() {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		log.ErrLog(err)
		return
	}
	sockAddr := &syscall.SockaddrInet4{
		Port: 8890,
	}
	err = syscall.Bind(fd, sockAddr)
	if err != nil {
		log.ErrLog(err)
		return
	}
	err = syscall.Listen(fd, 20)

	if err != nil {
		log.ErrLog(err)
		return
	}

	fmt.Println("Accepting connections")
	buf := make([]byte, 128)
	for {
		connFd, sockAddr, err := syscall.Accept(fd)
		if err != nil {
			log.ErrLog(err)
		}
		n, err := syscall.Read(connFd, buf)
		if err != nil {
			log.ErrLog(err)
		}
		cliSockAddr := sockAddr.(*syscall.SockaddrInet4)
		fmt.Printf("client Addr:%v,Port:%d  \n", cliSockAddr.Addr, cliSockAddr.Port)
		if n <= 0 { //读到零，说明收到fin，小与-1则错误
			fmt.Printf("server read n:%d \n", n)
			syscall.Close(connFd)
		}

		syscall.Write(connFd, bytes.ToTitle(buf[:n]))

		syscall.Close(connFd)
	}
	syscall.Close(fd)
}

func socketServerDemo2() {
	fd, _ := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)

	sockAddr := &syscall.SockaddrInet4{
		Port: 8890,
	}
	syscall.Bind(fd, sockAddr)

	syscall.Listen(fd, 20)

	fmt.Println("Accepting connections")
	buf := make([]byte, 128)
	for {
		connFd, sockAddr, _ := syscall.Accept(fd)

		n, _ := syscall.Read(connFd, buf)

		cliSockAddr := sockAddr.(*syscall.SockaddrInet4)
		fmt.Printf("client Addr:%v,Port:%d  \n", cliSockAddr.Addr, cliSockAddr.Port)
		fmt.Printf("server read n:%d \n", n)

		syscall.Write(connFd, bytes.ToTitle(buf[:n]))
		syscall.Close(connFd)
	}
	syscall.Close(fd)
}
