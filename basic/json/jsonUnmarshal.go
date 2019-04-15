package main

import (
	"encoding/json"
	"fmt"
	"github.com/lunny/log"
)

type People struct {
	Name string
	Age  int
}

func main() {
	//1 反序列化 解析已知类型
	var jsonP = []byte(`[
		{"Name":"wonkung","Age":23},
		{"Name":"pzxy","Age":5}
	]`)

	var people []People
	err := json.Unmarshal(jsonP, &people)
	if err != nil {
		log.Errorf("json.Unmarshal :", err)
	}
	fmt.Printf("%v", people)

	fmt.Println()
	fmt.Println("******************")
	//2 反序列化 解析未知类型
	var f interface{}
	b := []byte(`{"Name":"wonkung","Age":6,"Parents":["Yly","Whs"]}`)
	json.Unmarshal(b, &f)
	//类型断言
	for k, v := range f.(map[string]interface{}) {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", v)
		case float64:
			fmt.Println(k, "is float64", v)
		case []interface{}:
			fmt.Println(k, "is arrary :")
			for i, j := range vv {
				fmt.Println(i, j)
			}
		}
		//	fmt.Println(k,reflect.TypeOf(v))
	}
}
