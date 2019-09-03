package main

import (
	"bytes"
	"fmt"
)

func main(){
	test1()
}

func  test1(){
	var dstByte []byte
	var srcByte []byte
	dstByte = []byte("你好呀")
	srcByte = []byte("我太难了")
	s := len(dstByte)
	copy(dstByte[s:],srcByte[:])


	var buffer bytes.Buffer //Buffer是一个实现了读写方法的可变大小的字节缓冲
	buffer.Write(dstByte)
	buffer.Write(srcByte)
	dstByte =buffer.Bytes()
	fmt.Printf("data : %s ",dstByte)
}
