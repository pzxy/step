package leetcode

import (
	"fmt"
	"strings"
)

/**

请实现一个函数，输入一个整数，输出该数二进制表示中 1 的个数。
例如，把 9表示成二进制是 1001，有 2 位是 1。因此，如果输入 9，则该函数输出 2。

示例 1：

输入：00000000000000000000000000001011
输出：3
解释：输入的二进制串 00000000000000000000000000001011中，共有三位为 '1'。
示例 2：

输入：00000000000000000000000010000000
输出：1
解释：输入的二进制串 00000000000000000000000010000000中，共有一位为 '1'。
示例 3：

输入：11111111111111111111111111111101
输出：31
解释：输入的二进制串 11111111111111111111111111111101 中，共有 31 位为 '1'。

*/

func hammingWeight(num uint32) int {
	return strings.Count(fmt.Sprintf("%b", num), "1")
}

func hammingWeight2(num uint32) int {
	count := 0

	for 0 < num {
		if num%2 == 1 {
			count++
		}
		num /= 2
	}
	return count
}
func hammingWeight22(num uint32) int {
	count := 0

	for 0 < num {
		if num&1 == 1 {
			count++
		}
		num >>= 1
	}
	return count
}
func hammingWeight3(num uint32) int {
	count := 0
	for ; 0 < num; num >>= 1 {
		if num&1 == 1 {
			count++
		}
	}
	return count
}

func hammingWeight4(num uint32) int {
	count := 0
	for num != 0 { // 7&6 -> 0111&0110 = 0110 每次消去最后的1 ， 0110&0101 = 0100。每次都消去一个1
		num &= num - 1 //此操作能去掉一个1，减1刚好错位，然后与操作就不一样了。
		count++
	}
	return count
}
