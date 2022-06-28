package leetcode

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

/**
输入一个整数 n ，求1～n这n个整数的十进制表示中1出现的次数。

例如，输入12，1～12这些整数中包含1 的数字有1、10、11和12，1一共出现了5次。

示例 1：

输入：n = 12
输出：5
示例 2：

输入：n = 13
输出：6

限制：

1 <= n <2^31
*/

func countDigitOne(n int) int { //此方法当n很大时，会花费很长时间，超出时间限制。

	ret := 0
	for i := 1; i <= n; i++ {
		ii := i
		for ii > 0 {
			if ii%10 == 1 {
				ret++
			}
			ii = ii / 10
		}
	}
	return ret
}

func countDigitOne2(n int) int {
	digit := 1
	res := 0
	high := n / 10
	cur := n % 10
	low := 0
	for high != 0 || cur != 0 {
		if cur == 0 {
			res += high * digit
		} else if cur == 1 {
			res += high*digit + low + 1
		} else {
			res += (high + 1) * digit
		}
		low += cur * digit
		cur = high % 10
		high /= 10
		digit *= 10
	}
	return res
}

func complement(n int) string {
	if math.MaxInt16 < n || math.MinInt16 > n {
		return "NODATA"
	} //1111 1111 1111 1111
	//得到补码
	s := getBuMa(int16(n))
	var retBStr []rune
	l := len(s)
	for i := 0; i < 16-l; i++ {
		s = "0" + s
	}
	retBStr = []rune(s)
	var binaryStr string

	if n >= 0 {
		binaryStr = string(retBStr)
	} else {
		binaryStr = "-" + string(retBStr)
	}
	var xStr string
	for len(retBStr) > 0 {
		tmp := retBStr[:4]
		retBStr = retBStr[4:]
		i, _ := strconv.ParseUint(string(tmp), 2, 4)
		xStr += fmt.Sprintf("%x", i)
	}

	xStr = strings.ToUpper(xStr)
	return binaryStr + "," + xStr
}

func getBuMa(n int16) string {
	if n >= 0 {
		return fmt.Sprintf("%b", n)
	}
	f := math.Abs(float64(n))
	r := fmt.Sprintf("%b", int16(f))
	var retS []rune
	l := len(r)
	for i := 0; i < 16-l; i++ {
		r = "0" + r
	}
	retS = []rune(r)

	ret := make([]rune, 16)
	for i, v := range retS {
		if v == '0' {
			ret[i] = '1'
		}
		if v == '1' {
			ret[i] = '0'
		}
	}
	ret[0] = '1'
	for i := 15; i > 0; i-- {
		if ret[i] == '0' {
			ret[i] = '1'
			break
		}
		ret[i] = '0'
	}
	return string(ret)
}

//	s := fmt.Sprintf("%b", n)
//		r := make([]rune, 16-len(s))
//		r = append(r, []rune(s)...)
func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	s := "0123456789ABCDEF"

	ret := ""
	for num != 0 {
		index := num & 15
		ret = s[index:index+1] + ret
		num >>= 4
	}

	return ret
}
