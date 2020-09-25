package leetcode

/**
请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度。

示例1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是"wke"，所以其长度为 3。
请注意，你的答案必须是 子串 的长度，"pwke"是一个子序列，不是子串。

提示：

s.length <= 40000
*/
func lengthOfLongestSubstring(s string) (ret int) {
	if len(s) == 0 {
		return 0
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	str := []rune(s)
	var tmp []rune
	tmp = append(tmp, str[0])
	ret = 1
	for _, v := range str[1:] {
		for j, v2 := range tmp {
			if v == v2 {
				tmp = append(tmp, v)
				tmp = tmp[j+1:]
				break
			}
			if j != len(tmp)-1 {
				continue
			}
			tmp = append(tmp, v)
		}
		ret = max(ret, len(tmp))
	}

	return
}

//去掉数组
func lengthOfLongestSubstring2(s string) (ret int) {
	if len(s) == 0 {
		return 0
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	str := []rune(s)
	left, right := 0, 1

	ret = 1
	for _, v := range str[1:] {
		for j, v2 := range str[left:right] {
			if v == v2 {
				right++
				left = left + j + 1 //加一是因为j和left初始化都是0，不加1值不会变化
				break
			}
			if j != right-left-1 {
				continue
			}
			right++
			break
		}
		ret = max(ret, right-left)
	}

	return
}
