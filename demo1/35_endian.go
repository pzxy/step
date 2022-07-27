package main

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type AAAA struct {
	B string `json:"b"`
}

func main() {
	var body io.Reader
	if body == nil {
		fmt.Println("asd")
	}
	str := "eyJodW1pZGl0eSI6MTAwLCJ0ZW1wIjoyMDAsInBoVmFsdWUiOjAuMCwiZWNWYWx1ZSI6MC4wfQ=="

	b64 := base64.NewDecoder(base64.StdEncoding, strings.NewReader(str))
	bs := make([]byte, len(str))
	if n, err := b64.Read(bs); err != nil {
		panic(err)
	} else {
		bs = bs[:n]
	}
	fmt.Println(string(bs))

	base64.NewDecoder(base64.StdEncoding, strings.NewReader(str))
	bb := `{"b":"eyJodW1pZGl0eSI6MTAwLCJ0ZW1wIjoyMDAsInBoVmFsdWUiOjAuMCwiZWNWYWx1ZSI6MC4wfQ=="}`
	var a AAAA
	json.Unmarshal([]byte(bb), &a)
	fmt.Println(a)

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
