package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	e := Event{
		Id:        "1",
		DeviceId:  "1",
		ProfileId: "1",
		Created:   1,
		Origin:    1,
		Readings:  []Reading{{A: "1", B: "2"}},
	}
	A, err := json.Marshal(e)
	if err != nil {
		fmt.Println(err)
	}
	B := A
	var event Event
	err = json.Unmarshal(B, &event)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(event)

}

type Reading struct {
	A string
	B string
}
type Event struct {
	Id        string `gorm:"id;primaryKey"`
	DeviceId  string
	ProfileId string
	Created   int64
	Origin    int64
	Readings  []Reading `gorm:"-"json:"-"` // 这个逻辑从readings表中读取
}
