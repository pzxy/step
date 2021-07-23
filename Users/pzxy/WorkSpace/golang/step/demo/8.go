package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"net"
	"strings"
)

func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		return
	}

	var mac string
	var hhhhex [][]byte
	s := "14:7d:da:e:f3:0"

	for _, v := range strings.Split(s, ":") {
		temp := []rune(v)
		if len(temp) == 0 {
			v = "00"
		}
		if len(temp) == 1 {
			v = "0" + v
		}
		b, err := hex.DecodeString(v)
		if err != nil {
			fmt.Errorf("err:%s", err)
		}
		hhhhex = append(hhhhex, b)
	}

	for _, inter := range interfaces {
		inter.HardwareAddr.String()
		mac = inter.HardwareAddr.String()
		s := strings.Split(mac, ":")
		if len(s) != len(hhhhex) {
			fmt.Println("error：", mac)
			continue
		}
		var f bool
		for i, v := range s {
			b, _ := hex.DecodeString(v)
			if bytes.Compare(b, hhhhex[i]) != 0 {
				fmt.Println("不相等")
				f = true
				break
			}
		}
		if f {
			continue
		}

		fmt.Println(mac)
	}
	//	origData, err = aes_crypto.GcmDecrypt(crypted, encryptKey(mac))
	//	if err != nil {
	//		continue
	//	}
	//	return
	//}
	//err = errors.New("no mac addr or no valid mac addr")
	return

}
