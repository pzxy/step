package main

import (
	"encoding/json"
	"fmt"
	"runtime"
)

type A struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}
type B struct {
	Age int `json:"age"`
}

func (a A) d() {

}

func main() {
	runtime.Version()
	b, _ := json.Marshal(&B{
		Age: 19,
	})

	a, _ := json.Marshal(&A{
		Name: "BB",
		Data: b,
	})

	var aObj A
	var bObj B

	json.Unmarshal(a, &aObj)
	fmt.Println(aObj.Name)

	switch m := aObj.Data.(type) {

	case []byte:
		fmt.Println("string")
		json.Unmarshal(m, &bObj)

	default:

	}

	fmt.Println(bObj.Age)

}
