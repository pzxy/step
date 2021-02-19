package leetcode

import (
	"fmt"
)

func String2Int32(s string) int32 {
	s0 := s
	if s[0] == '-' || s[0] == '+' {
		s = s[1:]
		if len(s) < 1 {
			fmt.Println("fail 1")
		}
	}

	n := int32(0)
	for _, ch := range []byte(s) {
		if ch < '0' || ch > '9' {
			fmt.Println("fail 2")
		}
		n = n*10 + int32(ch)
	}
	if s0[0] == '-' {
		n = -n
	}
	return n
}

func String2Int(s string) int {
	s0 := s
	if s[0] == '-' || s[0] == '+' {
		s = s[1:]
		if len(s) < 1 {
			return -1
		}
	}
	n := 0
	for _, v := range s {
		if v < '0' || v > '9' {
			return -1
		}
		n = n*10 + int(v)
	}
	if s0[0] == '-' {
		n = -n
	}
	return n
}
