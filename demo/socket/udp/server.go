package main

import (
	"fmt"
	"net"
)

func main() {
	udp6667()
	udp6666()
}

func udp6666() {
	// 创建监听

	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", "0.0.0.0", 6666))
	if err != nil {
		fmt.Println("监听失败1!", err)
	}
	listener, err := net.ListenUDP("udp", addr)

	if err != nil {
		fmt.Println("监听失败!", err)
		return
	}
	defer listener.Close()

	for {
		// 读取数据
		data := make([]byte, 4096)
		read, remoteAddr, err := listener.ReadFromUDP(data)
		if err != nil {
			fmt.Println("读取数据失败!", err)
			continue
		}
		fmt.Println("6666", read, remoteAddr.String())
		fmt.Printf("6666 %s\n\n", data)

		// 发送数据
		//senddata := []byte("hello client!")
		//_, err = listener.WriteToUDP(senddata, remoteAddr)
		//if err != nil {
		//	return
		//	fmt.Println("发送数据失败!", err)
		//}
	}
}

func udp6667() {
	// 创建监听

	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", "0.0.0.0", 6667))
	if err != nil {
		fmt.Println("监听失败1!", err)
	}
	listener, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("监听失败!", err)
		return
	}
	defer listener.Close()

	for {
		// 读取数据
		data := make([]byte, 4096)
		read, remoteAddr, err := listener.ReadFromUDP(data)
		if err != nil {
			fmt.Println("读取数据失败!", err)
			continue
		}
		fmt.Println("6667", read, remoteAddr.String())
		fmt.Printf("6667 %s\n\n", data)
		//
		//// 发送数据
		//senddata := []byte("hello client!")
		//_, err = listener.WriteToUDP(senddata, remoteAddr)
		//if err != nil {
		//	return
		//	fmt.Println("发送数据失败!", err)
		//}
	}
}
