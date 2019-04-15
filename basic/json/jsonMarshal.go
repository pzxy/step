package main

import (
	"encoding/json"
	"fmt"
	"github.com/lunny/log"
)

type Animal struct {
	Name string `json:name`
	Age  int    `json:age`
}

func main() {
	//序列化
	var animals []Animal
	animals = append(animals, Animal{Name: "点点", Age: 12})
	animals = append(animals, Animal{Name: "豆豆", Age: 10})

	jsonStr, err := json.Marshal(animals)
	if err != nil {
		log.Errorf("json.Marshal :", err)
	}
	fmt.Println(string(jsonStr))
}
