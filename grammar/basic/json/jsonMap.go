package main

import (
	"encoding/json"
	"fmt"
	"github.com/lunny/log"
)

func main() {
	jsonToMap(mapToJson())
}

//map转json
func mapToJson() string {
	map1 := make(map[string]interface{})
	map1["a"] = "hello"
	map1["b"] = "world"
	map1["c"] = 9527

	jsonMap, err := json.Marshal(map1)

	if err != nil {
		log.Errorf("json.Marshal :", err)
	}

	fmt.Println(string(jsonMap))
	fmt.Println("*************")
	return string(jsonMap)
}

//json转map
func jsonToMap(jsonMap string) {
	//fmt.Println("jsonMap Type",reflect.TypeOf(jsonMap))
	map1 := make(map[string]interface{})
	json.Unmarshal([]byte(jsonMap), &map1)

	fmt.Println("a :", map1["a"])
	fmt.Println("b :", map1["b"])
	fmt.Println("c :", map1["c"])

	fmt.Printf("%+v", map1)
}
