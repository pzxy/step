package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"time"
)

func main() {
	netDial1()
}
func netDial1() {
	//net.Dail可以设置timeout
	//创建连接
	conn, err := net.Dial("tcp", "39.156.66.18:8888")
	//net.DialTimeout()可以代替net.Dial()设置超时时间

	conn.SetDeadline()
	//https://imhanjm.com/2017/10/29/深入理解golang时间处理(time.time)/?utm_campaign=studygolang.com&utm_medium=studygolang.com&utm_source=studygolang.com
	//发送数据
	conn.Write([]byte{0x00, 0x07, 0x00, 0x00, 0x00, 0x06, 0x01, 0x03, 0x00, 0x00, 0x00, 0x0A})
	defer conn.Close()
	buffer := make([]byte, 32)
	//读取返回数据
	result, err := conn.Read(buffer)
	if err != nil {
		fmt.Println(err)
	}

	//dial := &net.Dialer{
	//	Timeout:       0,
	//	Deadline:     time.Now(),
	//	LocalAddr:     nil,
	//	DualStack:     false,
	//	FallbackDelay: 0,
	//	KeepAlive:     0,
	//	Resolver:      nil,
	//	Cancel:        nil,
	//	Control:       nil,
	//}
	//dial.Dial()
	fmt.Println(result, string(buffer))
}
func netDial2() {
	//进行地址化转换
	tcpAddr, err := net.ResolveTCPAddr("tcp", "192.168.1.254:4001")
	//创建一个连接
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println(err)
	}
	str := []byte{32, 0, 0, 0, 7, 85, 35, 160, 176, 7, 226, 12, 18, 15, 45, 0}

	//向连接端发送一个数据
	_, err = conn.Write(str)
	if err != nil {
		fmt.Println(err)
	}
	//读取连接端返回来的数据，
	result, err := ioutil.ReadAll(conn)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(result))

	//从这两个请求可以看出，dial只需要直接填入连接地址，而dialtcp却需要提前把地址转换为对应的类型才能填入。
	// 在dial中，连接会根据你填入的连接类型，自动把地址转换为对应类型地址。
	//同时，我们还会发现我们dial中用了buffer接收数据，而在dialtcp用了ioutil.ReadAll()返回所有数据。
}
