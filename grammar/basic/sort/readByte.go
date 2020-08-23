package main

import "fmt"

func main() {
	srcByte := make([]byte, 0)
	srcByte = append(srcByte, 1)
	srcByte = append(srcByte, 0)
	srcByte = append(srcByte, 1)
	srcByte = append(srcByte, 0)
	srcByte = append(srcByte, 0)
	srcByte = append(srcByte, 0)
	srcByte = append(srcByte, 23)
	srcByte = append(srcByte, 67)
	srcByte = append(srcByte, 0)
	srcByte = append(srcByte, 0)
	srcByte = append(srcByte, 1)
	srcByte = append(srcByte, 0)
	srcByte = append(srcByte, 45)
	srcByte = append(srcByte, 0)
	srcByte = append(srcByte, 0)
	srcByte = append(srcByte, 0)
	srcByte = append(srcByte, 6)
	srcByte = append(srcByte)
	srcByte = append(srcByte)
	srcByte = append(srcByte)
	srcByte = append(srcByte, 0)
	fmt.Printf("%v\n", srcByte)
	fmt.Printf("%v\n", test1(srcByte))
}

func test1(srcByte []byte) []byte {
	var sendPkg []byte
	var tempPkg []byte
	for _, data := range srcByte {
		if data == 0 {
			tempPkg = append(tempPkg, data)
			continue
		}
		if data != 0 {
			if len(tempPkg) > 0 {
				sendPkg = append(sendPkg, tempPkg...)
				tempPkg = make([]byte, 0)
			}
			sendPkg = append(sendPkg, data)
		}
	}
	return sendPkg
}
