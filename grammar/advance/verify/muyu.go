package main

import (
	"fmt"
	"strconv"
)

func main() {
	arrangeNumber(1234)
}

var ret []string
var mapS = make(map[string]string, 0)

func arrangeNumber(n int) {
	str := strconv.Itoa(n)
	var s []string
	for _, v := range str {
		s = append(s, string(int(v)))
	}

	if len(s) <= 1 {
		fmt.Println(s)
	}
	recMath(s, "")
	for _, v := range ret {
		fmt.Println(v)
	}
}

func recMath(s []string, str string) {
	if len(s) == 1 {
		if _, ok := mapS[str]; ok {
			return
		}
		str += s[0]
		mapS[str] = str
		ret = append(ret, str)
		return
	}
	for i, v := range s {
		tmpStr := str + v
		before := make([]string, len(s[:i]))
		after := make([]string, len(s[i+1:]))
		copy(before, s[:i])
		copy(after, s[i+1:])
		before = append(before, after...)
		recMath(before, tmpStr)
	}
}
