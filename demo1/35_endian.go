package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	fmt.Println(IntToBytes(123))

}

func IntToBytes(n int) []byte {
	x := uint32(9798)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.LittleEndian, x)
	return bytesBuffer.Bytes()
}
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}
