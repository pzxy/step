package leetcode

//寻找最长不含有重复字符的子串，里面的rune直接把中文转为一个字节，中英文一起操作
func lengthOf(s string) int {
	lastOcc := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		if lastI, ok := lastOcc[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOcc[ch] = i
	}
	return maxLength
}

func lengthOf2(s string) int {
	cMap := make(map[rune]int)
	start, maxLength := 0, 0

	for i, char := range []rune(s) {
		if index, ok := cMap[char]; ok && index >= start {
			start = index + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		cMap[char] = i
	}
	return maxLength
}
