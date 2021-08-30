package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"
)

func main() {
	s := "00 00 00 00 00 00 00 13  00 00 00 9c 00 00 00 00  \n d0 97 66 67 6f 33 69 eb  10 b5 e9 f1 32 fd 80 2a \n c3 78 85 40 6a e6 9b 1a  29 b0 2d db 08 38 5d 7f \n 6b f3 42 ed 3e 8d cb 09  5a 80 4e f8 1a 43 62 aa  \n ac 73 6d 4f 3d 56 73 85  cb ff 6e f4 14 eb 50 cd \n c2 fb 84 59 b1 15 5f c7  5d 4b f6 69 9f 92 cb a4  \n c0 ba 52 01 48 04 5e 76  05 fa 04 98 df ea 5a ab  \n fa a2 1b 8b e2 68 f3 ba  5a 97 64 cf 44 1b e8 85  \n 17 0b 32 6d 80 3d 30 a9  0c 87 45 3b eb ee 4c df  \n 6e da 8e ea 40 e9 3b 1e  3f c1 4a 25 70 e1 82 79"
	ss := strings.Replace(s, " ", "", -1)
	sss := strings.Replace(ss, "\n", "", -1)
	fmt.Println(sss)

	b := 1079460891
	fmt.Println(uint32ToBytes(uint32(b)))
}

func uint32ToBytes(n uint32) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, n)
	return bytesBuffer.Bytes()
}
