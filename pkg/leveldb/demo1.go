package main

import (
	"fmt"
	"strings"
)

func getJointIds(ids ...string) (ret string) {
	if len(ids) == 0 {
		return ""
	}
	ret = strings.Join(ids, ":") + ":"
	return
}

func main() {
	fmt.Println(getJointIds("1"))
}
