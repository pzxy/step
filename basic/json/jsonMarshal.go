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

type Animal2 struct {
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

	var animals2 []Animal2
	animals2 = append(animals2, Animal2{Name: "动物名字", Age: 12})
	animals2 = append(animals2, Animal2{Name: "动物名字", Age: 12})

	jsonStr, err = json.Marshal(animals2)
	if err != nil {
		log.Errorf(`json.Marshal:`, err)
	}
	fmt.Println(string(jsonStr))
}
