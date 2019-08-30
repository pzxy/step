package main

import (
	"encoding/json"
	"fmt"
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

	jsonStr, _ := json.Marshal(animals)

	fmt.Println(string(jsonStr))

	var animals2 []Animal2
	animals2 = append(animals2, Animal2{Name: "动物名字", Age: 12})
	animals2 = append(animals2, Animal2{Name: "动物名字", Age: 12})

	jsonStr, _ = json.Marshal(animals2)

	fmt.Println(string(jsonStr))
}
