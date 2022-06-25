package leetcode

import (
	"strings"
)

/**
输入一个英文句子，翻转句子中单词的顺序，但单词内字符的顺序不变。
为简单起见，标点符号和普通字母一样处理。例如输入字符串"I am a student. "，则输出"student. a am I"。

示例 1：

输入: "the sky is blue"
输出:"blue is sky the"
示例 2：

输入: " hello world! "
输出:"world! hello"
解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
示例 3：

输入: "a good  example"
输出:"example good a"
解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。

说明：

无空格字符构成一个单词。
输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
*/
func reverseWords2(s string) (ret string) {
	chars := []rune(s)
	tmpStr := make([]string, 0)
	var tmpS string
	for _, c := range chars {
		if tmpS == "" && c == ' ' { //每个单词开始前，遇到空格跳过
			continue
		}
		if c == ' ' { //每个单词后遇到的空格，跳过
			tmpStr = append(tmpStr, tmpS)
			tmpS = ""
			continue
		}
		tmpS += string(c)
	}
	if tmpS != "" { //最后一个字符串
		tmpStr = append(tmpStr, tmpS)
	}
	//反向拼接字符串
	for i := len(tmpStr) - 1; i >= 0; i-- {
		strings.Join(tmpStr, " ")
	}
	return
}

func reverseWords(s string) (ret string) {
	str := strings.Split(s, " ")
	tmpStr := make([]string, 0)
	for i := len(str) - 1; i >= 0; i-- {
		if len(str[i]) <= 0 {
			continue
		}
		tmpStr = append(tmpStr, strings.TrimSpace(str[i]))
	}
	ret = strings.Join(tmpStr, " ")
	return
}
