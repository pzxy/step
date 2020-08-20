package leetcode

/**
实现一个函数，将字符串中的每个空格替换成"%20"，
例如： "we are happy"
输出:  "we%20are%20happy"

为了使时间复杂度为O(n),我们设置两个
*/
func replaceSpaces(s string) string {
	if len(s) == 0 {
		return ""
	}
	const (
		space   = 32
		percent = 37
		zero    = 48
		two     = 50
	)
	var b []rune
	for _, v := range s {
		if v == space {
			b = append(b, '%') //写'%' 或者 32 都可以
			b = append(b, '2')
			b = append(b, '0')
		} else {
			b = append(b, v)
		}
	}
	return string(b)
}
