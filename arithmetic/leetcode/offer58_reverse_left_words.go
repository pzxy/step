package leetcode

/**
字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。
请定义一个函数实现字符串左旋转操作的功能。比如，输入字符串"abcdefg"和数字2，该函数将返回左旋转两位得到的结果"cdefgab"。

示例 1：

输入: s = "abcdefg", k = 2
输出:"cdefgab"
示例 2：

输入: s = "lrloseumgh", k = 6
输出:"umghlrlose"

限制：

1 <= k < s.length <= 10000
*/
func reverseLeftWords2(s string, n int) (ret string) {
	if len(s)-1 < n {
		return
	}
	chars := []rune(s)
	ret = string(chars[n:]) + string(chars[:n])
	return
}

func reverseLeftWords3(s string, n int) (ret string) {
	if len(s)-1 < n {
		return
	}
	var tmpStr string
	for i, v := range s {
		if i < n {
			tmpStr += string(v)
			continue
		}
		ret += string(v)
	}
	ret += tmpStr
	return
}

func reverseLeftWords(s string, n int) string {
	return s[n:] + s[:n]
}