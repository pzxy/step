package main

import (
	"fmt"
	"github.com/firstrow/tcp_server"
	"net"
	"os"
	"strings"
)

const port46 = "7983"

func main() {
	server := tcp_server.New("localhost:" + port46)
	//server.OnNewClient(func(c *tcp_server.Client) {
	//	// new client connected
	//	// lets send some message
	//	//c.Send("Hello")
	//	fmt.Println("1")
	//})
	server.OnNewMessage(func(c *tcp_server.Client, message string) {
		// new message received
		//fmt.Println(message)
		if message == "host\n" {
			name, err := os.Hostname()
			if err != nil {
				_ = c.Send(err.Error())
			} else {
				_ = c.Send(name)
			}
		}
	})
	//server.OnClientConnectionClosed(func(c *tcp_server.Client, err error) {
	//	// connection with client lost
	//	fmt.Println("3")
	//})

	server.Listen()
}

func getIps() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
	}
	var ips []string
	for _, address := range addrs {
		ipNet, ok := address.(*net.IPNet)
		if !ok {
			continue
		}
		if ipNet.IP.IsLoopback() {
			continue
		}
		if ipNet.IP.To4() != nil {
			ips = append(ips, ipNet.IP.String())
		}
	}
	return strings.Join(ips, ",")
}
