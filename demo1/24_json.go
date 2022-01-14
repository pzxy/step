package main

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"go/ast"
)

type Data24 struct {
	Id   int     `json:"id"`
	Age  float64 `json:"age"`
	Name string  `json:"name"`
}

func main() {
	var data = Data24{
		Id:   1,
		Age:  18.1,
		Name: "pzxy",
	}
	var ret interface{}

  ast.Inspect(nil, func(node ast.Node) bool {

  })
	marshal, _ := json.Marshal(data)
	json.Unmarshal()
	err :=  mapstructure.Decode(marshal, &ret)
	fmt.Println(err)
	fmt.Println(fmt.Sprintf("%#v",ret))
}
