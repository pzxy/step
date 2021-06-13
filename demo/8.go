package main

import (
	"fmt"
	"net"
)

func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		return
	}
	var mac string
	for _, inter := range interfaces {
		mac = inter.HardwareAddr.String()
		fmt.Println(mac == "14:7d:da:0e:f3:00")
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
