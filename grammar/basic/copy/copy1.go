package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"step/utils/log"
)

func main() {
	testReadBin4()
}

func test1() {
	var dstByte []byte
	var srcByte []byte
	dstByte = []byte("你好呀")
	srcByte = []byte("我太难了")
	s := len(dstByte)
	copy(dstByte[s:], srcByte[:])

	var buffer bytes.Buffer //Buffer是一个实现了读写方法的可变大小的字节缓冲
	buffer.Write(dstByte)
	buffer.Write(srcByte)
	dstByte = buffer.Bytes()
	fmt.Printf("data : %s ", dstByte)
}

type DTResult uint16

const dtResult DTResult = 0x01FF

func test2() {
	tmpByte := make([]byte, 0)
	buf := bytes.NewBuffer(tmpByte)
	err := binary.Write(buf, binary.LittleEndian, dtResult)
	if err != nil {
		//这里要有错都有错,不是哪个任务的问题
		log.ErrLog(err)
	}
	var failNum uint16
	fmt.Println(failNum)
	if len(buf.Bytes()) > 0 {
		failNum = uint16(buf.Bytes()[1])
	}
	fmt.Println(buf.Bytes())
	fmt.Println(failNum)
}

type Result struct {
	Code1 uint16
}

func testReadBin4() {
	r := new(Result)
	tmpByte := make([]byte, 0)
	tmpByte = append(tmpByte, 0xAA)
	tmpByte = append(tmpByte, 0x77)
	buf := bytes.NewBuffer(tmpByte)
	if err := binary.Read(buf, binary.LittleEndian, r); err != nil {
		log.ErrLog(err)
	}
	fmt.Println(r)
}
func test3() {
	var i1 int64 = 511 // [00000000 00000000 ... 00000001 11111111] = [0 0 0 0 0 0 1 255]
	s1 := make([]byte, 0)
	buf := bytes.NewBuffer(s1)
	// 数字转 []byte, 网络字节序为大端字节序
	binary.Write(buf, binary.BigEndian, i1)
	fmt.Println(buf.Bytes())
}
