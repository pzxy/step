package leetcode

import (
	"fmt"
	"strconv"
)

/**
实现函数double Power(double base, int exponent)，求base的exponent次方。
不得使用库函数，同时不需要考虑大数问题。

示例 1:

输入: 2.00000, 10
输出: 1024.00000
示例2:

输入: 2.10000, 3
输出: 9.26100
示例3:

输入: 2.00000, -2
输出: 0.25000
解释: 2^-2 = 1/2^2 = 1/4 = 0.25

说明:

-100.0 <x< 100.0
n是 32 位有符号整数，其数值范围是[−231,231− 1]。
*/

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	result := myPow(x, n>>1)
	result *= result
	if n&0x1 == 1 {
		result *= x
	}
	return result
}

func myPow2(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	var ret float64 = 1

	if n >= 0 {
		for i := 0; i < n; i++ {
			ret = ret * x
		}
	} else {
		for i := 0; i < -n; i++ {
			ret = ret * x
		}
		ret = 1.0 / ret
	}
	ret, _ = strconv.ParseFloat(fmt.Sprintf("%.5f", ret), 5)
	return ret
}
func myPow3(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}
	temp := myPow(x, n>>1)
	if n&1 == 0 {
		return temp * temp
	}
	return x * temp * temp
}

func myPow4(x float64, n int) float64 {
	if x == 0 {
		return x
	}
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}
	temp := myPow4(x, n>>1) //要求一个数的16次方，我们只需要用他们两个8次方相乘。
	if n&1 == 0 {
		return temp * temp
	}
	return x * temp * temp
}

func myPow5(x float64, n int) float64 {
	if x == 0 {
		return x
	}
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}
	var ret float64 = 1.0
	for n > 0 {
		if n&1 == 1 {
			ret *= x
			n--
			continue
		}
		x *= x
		n = n >> 1
	}
	return ret
}
