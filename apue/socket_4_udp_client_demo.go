package main

import (
	"encoding/json"
	"fmt"
	"step/grammar/codeskill/log"
	"syscall"
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
		Port: 9527,
	}
	fmt.Println("sending ... ")
	buf := make([]byte, 80)

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
	fmt.Printf("serverSock info :%+v \n", serverSockAddr)
	fmt.Printf("from server data :%v \n", buf[:n])

	syscall.Close(fd)
}
