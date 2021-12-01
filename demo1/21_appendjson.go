package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type AppendJson struct {
	A int
	B string
}

func main() {
	s := "{\"body\":{\"url\":\"客户分支1\"},\"position\":{\"x\":577.5714285714286,\"y\":-302}}"
	m := map[string]interface{}{
		"a": "2",
	}
	m2 := AppendJson{
		A: 1,
		B: "asd",
	}
	fmt.Println(appendJson(s, "data", m))
	fmt.Println(appendJson(s, "appendJson", m2))
}

func appendJson(jsonStr string, key string, data interface{}) string {
	if len(jsonStr) == 0 {
		jsonStr = "{}"
	}
	if len(key) == 0 {
		return jsonStr
	}
	if data == nil {
		return jsonStr
	}
	v, err := json.Marshal(data)
	if err != nil || len(v) == 0 {
		return jsonStr
	}
	suffix := strings.TrimSuffix(jsonStr, "}")
	var split string
	if len(suffix) != 1 {
		split = ","
	}
	return suffix + fmt.Sprintf(`%s"%s":%s}`, split, key, string(v))
}
