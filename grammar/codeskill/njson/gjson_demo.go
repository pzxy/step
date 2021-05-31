package main

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
)

type GDemo struct {
	Name string
	Age  int
}

func main() {
	j := `{"name": "Gilbert", "age": 61} {"name": "Alexa", "age": 34} {"name": "May", "age": 57} {"name": "Deloise", "age": 44} `
	gjson.ForEachLine(j, func(line gjson.Result) bool {
		var gd GDemo
		b := []byte(line.String())
		json.Unmarshal(b, &gd)
		fmt.Println(gd)
		return true
	})

}
