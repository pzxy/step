package bytedance

import "fmt"

//一个N位数K，从N位中去掉M个数字，使剩下的数字最大，leetcode有类似题

func RemoveKdigits(str string, m int) int {

	if str[0] == '-' {

	}
	if len(str) <= m {
		return 0
	}
	x := ^uint32(0) >> 1
	fmt.Println(x)
	var q []int
	for _, s := range str {
		v := int(s - '0')
		for m > 0 && len(q) > 0 && v > q[len(q)-1] {
			q = q[:len(q)-1]
			m--
		}
		q = append(q, v)
	}
	ret := 0
	for _, v := range q {
		ret = ret*10 + v
	}
	return ret
}
