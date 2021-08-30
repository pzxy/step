package main

import (
	"fmt"
	"net"
)

func main() {
	// 创建连接
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", "127.0.0.1", 8080))
	if err != nil {
		fmt.Println("连接失败!1", err)
	}
	conn, err := net.DialUDP("udp4", nil, addr)
	if err != nil {
		fmt.Println("连接失败!", err)
		return
	}
	defer conn.Close()

	// 发送数据
	senddata := []byte("hello server!")
	_, err = conn.Write(senddata)
	if err != nil {
		fmt.Println("发送数据失败!", err)
		return
	}

	// 接收数据
	data := make([]byte, 4096)
	read, remoteAddr, err := conn.ReadFromUDP(data)
	if err != nil {
		fmt.Println("读取数据失败!", err)
		return
	}
	fmt.Println(read, remoteAddr)
	fmt.Printf("%s\n", data)
}
