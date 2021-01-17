package main

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
)

var jsonStr = `
{
  "name": {"first": "Tom", "last": "Anderson"},
  "age":37,
  "children": ["Sara","Alex","Jack"],
  "fav.movie": "Deer Hunter",
  "friends": [
    {"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
    {"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},
    {"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
  ]
}
`

func main() {
	dic := JsonParse(jsonStr)
	out, err := json.MarshalIndent(dic, " ", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(string(out))
}

func JsonParse(jsonStr string) map[string]interface{} {
	res := make(map[string]interface{})
	jsonObj := gjson.Parse(jsonStr)

	if isObject(jsonObj) {
		parseJson2Map(jsonObj, res, "")
	} else if isArray(jsonObj) {
		for idx, obj := range jsonObj.Array() {
			key := fmt.Sprintf("%d", idx)
			parseJson2Map(obj, res, key)
		}
	}

	return res
}

func parseJson2Map(obj gjson.Result, res map[string]interface{}, prefix string) {
	if isObject(obj) {
		for key, value := range obj.Map() {
			if prefix != "" {
				key = prefix + "." + key
			}
			if isObject(value) {
				parseJson2Map(value, res, key)
			} else if isArray(value) {
				for idx, v := range value.Array() {
					fullkey := key + "." + fmt.Sprintf("%d", idx)
					parseJson2Map(v, res, fullkey)
				}
			} else {
				res[key] = value.Value()
				continue
			}
		}
	} else {
		res[prefix] = obj.Value()
	}
}

func isObject(obj gjson.Result) bool {
	if obj.IsObject() {
		return true
	}
	return false
}

func isArray(obj gjson.Result) bool {
	if obj.IsArray() {
		return true
	}
	return false
}
