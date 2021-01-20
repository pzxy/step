package main

import (
	"encoding/json"
	"fmt"
	"step/grammar/codeskill/log"
	"syscall"
	"time"
)

func main() {
	//syscall.Select()
	socketClint4Udp()
}

func socketClint4Udp() {
	localhost := [4]byte{
		127, 0, 0, 1,
	}
	//changAddr := [4]byte{
	//	192, 168, 1, 4,
	//}
	s := "hello world \n"
	b, _ := json.Marshal(s)
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, 0)
	if err != nil {
		log.ErrLog(err)
		return
	}
	socketAddr := &syscall.SockaddrInet4{
		Addr: localhost,
		Port: 9627,
	}
	fmt.Println("sending ... ")
	buf := make([]byte, 80)
	for {
		err = syscall.Sendto(fd, b, 0, socketAddr)
		if err != nil {
			log.ErrLog(err)
			return
		}
		fmt.Println("send over ... ")
		n, serverSockAddr, err := syscall.Recvfrom(fd, buf, 0)
		if err != nil {
			log.ErrLog(err)
			return
		}
		socket := serverSockAddr.(*syscall.SockaddrInet4)
		//socket.ram是可以操作底层的套接字，叫原生套接字，可以直接改udp或者tcp的头部字段。在传输层之下
		fmt.Printf("serverSock info :%+v \n", socket.Addr)
		fmt.Printf("from server data :%s \n", buf[:n])
		fmt.Println()
		time.Sleep(time.Second)
	}
	syscall.Close(fd)
}
