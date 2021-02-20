package leetcode

import (
	"math"
)

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

func String2Int32(str string) int32 {
	if len(str) == 0 {
		return -1
	}
	s := str
	if s[0] == '-' || s[0] == '+' {
		s = s[1:]
	}
	n := int32(0)
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return -1
		}
		ch = ch - '0'
		n = n*10 + ch
	}
	if math.Abs(float64(n)) > float64(^uint32(0)>>1) {
		return -1
	}
	if str[0] == '-' {
		n = -n
	}
	return n
}
