package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	str2uint()
}
func str2uint() {
	s := "1.2"
	//buf := &bytes.Buffer{}
	//errors.WithStack(binary.Write(buf, binary.LittleEndian, s))

	//var FirmwareDetails [2]byte
	b, _ := hex.DecodeString(s)

	//copy(FirmwareDetails[:],b[:])

	fmt.Printf("%s", b)
}
