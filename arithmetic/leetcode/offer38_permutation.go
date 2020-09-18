package leetcode

/**
输入一个字符串，打印出该字符串中字符的所有排列。
你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。

示例:
输入：s = "abc"
输出：["abc","acb","bac","bca","cab","cba"]

限制：
1 <= s 的长度 <= 8
*/
func permutation(s string) []string {
	ret := make([]string, 0)
	if len(s) == 0 {
		return ret
	}
	startIdx := 0
	endIdx := len(s) - 1
	retMap := make(map[string]bool, 0)
	retMap[s] = true
	recExchange(startIdx, endIdx, retMap)

	for v, _ := range retMap {
		ret = append(ret, v)
	}
	return ret
}

func recExchange(startIdx, endIdx int, retMap map[string]bool) {
	if startIdx == endIdx {
		return
	}
	tmp := make([]string, 0)
	for v, _ := range retMap {
		str, ex := []rune(v), []rune(v)
		for i := startIdx; i <= endIdx; i++ {
			if str[i] == str[startIdx] {
				continue
			}
			ex[i], ex[startIdx] = ex[startIdx], ex[i]
			tmp = append(tmp, string(ex))
		}
	}
	for _, v := range tmp {
		retMap[v] = true
	}
	recExchange(startIdx+1, endIdx, retMap)
}
