package leetcode

import (
	"math"
	"step/grammar/codeskill/log"
	"strconv"
)

/**
输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。

示例 1:

输入: n = 1
输出: [1,2,3,4,5,6,7,8,9]


说明：

用返回一个整数列表来代替打印
n 为正整数

*/
func printNumbers(n int) []int {
	if n == 0 {
		return nil
	}

	end := int(math.Pow10(n))
	ret := make([]int, 0, end)
	for i := 1; i < end; i++ {
		ret = append(ret, i)
	}

	return ret
}

func printNumbers2(n int) (ret []int) {
	if n == 0 {
		return nil
	}
	s := make([]rune, n+1)

	increment := func(s []rune) bool {
		idx := n

		for idx >= 0 {
			tmp := s[idx] + 1
			if tmp == 10 {
				s[idx] = 0
				idx--
				continue
			}
			s[idx] = tmp
			break
		}
		return true
	}
	for increment(s) {
		if s[0] == 1 {
			return
		}
		var tmpStr string
		for _, v := range s {
			tmpStr += strconv.Itoa(int(v))
		}
		i, err := strconv.Atoi(tmpStr)
		if err != nil {
			log.ErrLog(err)
			return
		}
		ret = append(ret, i)

	}

	return
}
