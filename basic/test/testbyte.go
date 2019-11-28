package main

import (
	"fmt"
)

//测试按个读字节

func main() {
	var sbyte [1024]byte
	var byte2 []byte
	var byte3 []byte
	byte2 = []byte("你好呀")
	fmt.Printf("byte2数据: %v\n", byte2)
	copy(sbyte[:], byte2[:])
	fmt.Printf("1024数据: %v\n", len(sbyte))
	for _, data := range sbyte {
		if data != 0 {
			byte3 = append(byte3, data)
		} else {
			break
		}
	}
	fmt.Printf("byte3数据: %v\n", byte3)

}
