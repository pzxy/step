package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

const port47 = "7983"

func main() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
	}
	for _, address := range addrs {
		ipNet, ok := address.(*net.IPNet)
		if !ok {
			continue
		}
		if ipNet.IP.IsLoopback() {
			handleLoopback(ipNet.IP.String())
			continue
		}
		if ipNet.IP.To4() != nil {
			handleIpV4(ipNet.IP.String())
		}
	}
}

func handleLoopback(ip string) {
	if ret, err := exist(ip + ":" + port47); err == nil {
		fmt.Println(ret + "," + ip)
	}
}

func handleIpV4(ip string) {
	ipSplit := strings.Split(ip, ".")
	ipBase := strings.Join(ipSplit[:3], ".") + "."
	wg := &sync.WaitGroup{}
	for i := 1; i < 255; i++ {
		ip := ipBase + strconv.Itoa(i)
		wg.Add(1)
		go func() {
			if ret, err := exist(ip + ":" + port47); err == nil {
				fmt.Println(ret + "," + ip)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func exist(ip4 string) (string, error) {
	conn, err := net.DialTimeout("tcp", ip4, 3*time.Second)
	if err != nil {
		return "", err
	}
	defer conn.Close()
	if _, err = conn.Write([]byte("host\n")); err != nil {
		return "", err
	}
	msg := make([]byte, 1024)
	conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	if _, err = conn.Read(msg); err != nil {
		return "", err
	}
	return string(msg), nil
}
