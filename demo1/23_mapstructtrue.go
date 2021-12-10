package main

import (
	"encoding/json"
	"github.com/lunny/log"
	"github.com/mitchellh/mapstructure"
)

func main() {
	msg := &MapStructMsg{
		ReqType: "123123",
		MsgId:   "msgID",
		Product: Product{
			A: "a",
			B: "b",
			C: 3,
		},
		Dps: []Dp{{A: "a"}, {A: "aa"}},
	}
	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		log.Error(err)
		return
	}
	m := make(map[string]interface{})
	err = json.Unmarshal(jsonMsg, &m)
	if err != nil {
		log.Error(err)
		return
	}
	log.Println("map====>%+v",m)
	ret := new(MapStructMsg)
	err = mapstructure.Decode(m, ret)
	if err != nil {
		log.Error(err)
		return
	}
	log.Println("====>", ret)
}

type MapStructMsg struct {
	ReqType string  `json:"reqType"`
	MsgId   string  `json:"msgId"`
	Product Product `json:"product"`
	Dps     []Dp    `json:"dps"`
}

type Product struct {
	A string `json:"a"`
	B string `json:"b"`
	C int    `json:"c"`
}

type Dp struct {
	A string `json:"a"`
}
