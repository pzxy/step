package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	//0X0A0B0C0D
	b := int32(168496141)
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(buf.Bytes())
}
