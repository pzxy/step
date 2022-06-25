package leetcode

/**
在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。

示例:

s = "abaccdeff"
返回 "b"

s = ""
返回 " "

限制：

0 <= s 的长度 <= 50000
*/
func firstUniqChar(s string) byte {
	if len(s) == 0 {
		return ' '
	}
	time := make([]byte, 26)
	order := make([]int, 26)
	for i, v := range s {
		time[v-97]++
		order[v-97] = i
	}

	idx := 50001
	var ret byte
	ret = ' '
	for i, v := range time {
		if v != 1 {
			continue
		}
		if idx > order[i] {
			idx = order[i]
			ret = byte(i) + 97
		}
	}

	return ret
}
