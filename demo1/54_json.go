package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func main() {
	str := `{"result":true,"success":true,"name":"topswapboot_pro_1.0.0.T","logno":"6b5fef11ecb12d0b1c696282d62bf778","servertime":1697599880619,"code":200,"data":true,"message":"OK"}`
	result := &VerifyResp{}
	err := json.Unmarshal([]byte(str), result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

	fmt.Println(hexString2Uint64("0x07"))
}

func hexString2Uint64(s string) uint64 {
	if len(s) < 2 {
		return 0
	}
	s = s[2:]
	ret, err := strconv.ParseUint(s, 16, 64)
	if err != nil {
		return 0
	}
	return ret
}

type VerifyResp struct {
	Code       int32  `json:"code"`
	Logno      string `json:"logno"`
	Message    string `json:"message"`
	Name       string `json:"name"`
	Result     bool   `json:"result"`
	Servertime uint64 `json:"servertime"`
}
