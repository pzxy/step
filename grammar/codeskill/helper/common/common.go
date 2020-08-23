package common

import "encoding/json"

func Json(v interface{}) []byte {
	s, _ := json.Marshal(v)
	return s
}
