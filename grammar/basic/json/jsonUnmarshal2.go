package main

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"step/grammar/codeskill/log"
)

type DtRobotsReq struct {
	DownloadUrl        string   `json:"downloadUrl"`
	Content            string   `json:"content"`
	Version            string   `json:"version"`
	MD5                string   `json:"md5"`
	UpdateSerialNumber uint32   `json:"updateSerialNumber"`
	DeviceIds          []uint32 `json:"deviceIds"`
}

var jsonP = []byte(`	{
		"content":"第二个版本",
		"deviceIds":[
		2
],
	"downloadUrl":"http://update.file.test.chipparking.com/robot/robot.bin?e=1569484833&token=lg7QUwopX7lpjq6xH4O98cocrkiTWYVo0W2PjJF7:jUqAvmB5fKstJARFi4SJiu_Vmks=",
	"md5":"5a1f3dbbae52785f58bdea8db575050b",
	"updateSerialNumber":2,
	"version":"2.1"
	}`)

type DtRobotOrder struct {
	//接受来自后台的数据
	DtRobotsReq

	//以下信息不用后台传
	//固件信息
	FirmwareDetails []byte
	//拆分后的固件包
	BreakFirmwarePkg [][1024]byte
	//固件大小 32bit
	FirmwareSize uint32
	//总包数目 uint8个数
	PkgTotal uint8
	//单包大小  uint8字节数
	PkgOneSize uint8
}

func main() {
	testJson(jsonP)
}
func testJson(msg []byte) {
	tRobotsReq := &DtRobotsReq{}

	/*testJson := make([]ChargeRobotOrder, 0)*/

	if err := json.Unmarshal(msg, tRobotsReq); err != nil {
		log.ErrLog(errors.Errorf("msg = %+v", tRobotsReq))
	}
	fmt.Printf("%v", tRobotsReq)
}
