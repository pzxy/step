package main

import (
	"fmt"
	"io/ioutil"
)

type DtTask struct {
	//接受后台的数据
	order []byte
	//拆分的固件切片Break up the firmware
	BreakFirmwarePkg [][1024]byte
}

func (t *DtTask) TryBreakFirmwarePkg() {
	var oneByte [1024]byte
	var tmpByte [1024]byte
	srcByte := t.order

	for {
		s := len(srcByte)
		if s > 0 {
			if s < 1024 {
				copy(oneByte[:], srcByte[:s])
				t.BreakFirmwarePkg = append(t.BreakFirmwarePkg, oneByte)
				oneByte = tmpByte
				srcByte = nil
			}
			if s >= 1024 {
				copy(oneByte[:], srcByte[:1024])
				t.BreakFirmwarePkg = append(t.BreakFirmwarePkg, oneByte)
				oneByte = tmpByte
				srcByte = srcByte[1024:]
			}
		} else {
			break
		}
		//	time.Sleep(time.Second)
	}
}
func main() {
	data, _ := ioutil.ReadFile("D:/workspace/Go/src/step/basic/slice/re.json")
	t := &DtTask{
		order:            data,
		BreakFirmwarePkg: nil,
	}
	t.TryBreakFirmwarePkg()

	fmt.Printf("%s\n", t.BreakFirmwarePkg[1])

}
