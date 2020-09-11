package leetcode

import "strings"

/**

请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。
例如，字符串"+100"、"5e2"、"-123"、"3.1416"、"-1E-16"、"0123"都表示数值，
但"12e"、"1a3.14"、"1.2.3"、"+-5"及"12e+5.4"都不是。

A[.[B]][e|EC] 或者 .B[e|EC]

*/

/**
方法一：确定有限状态自动机
预备知识

确定有限状态自动机（以下简称「自动机」）是一类计算模型。它包含一系列状态，这些状态中：

有一个特殊的状态，被称作「初始状态」。
还有一系列状态被称为「接受状态」，它们组成了一个特殊的集合。其中，一个状态可能既是「初始状态」，也是「接受状态」。
起初，这个自动机处于「初始状态」。
随后，它顺序地读取字符串中的每一个字符，并根据当前状态和读入的字符，按照某个事先约定好的「转移规则」，
从当前状态转移到下一个状态；当状态转移完成后，它就读取下一个字符。
当字符串全部读取完毕后，如果自动机处于某个「接受状态」，则判定该字符串「被接受」；否则，判定该字符串「被拒绝」。

注意：如果输入的过程中某一步转移失败了，即不存在对应的「转移规则」，此时计算将提前中止。在这种情况下我们也判定该字符串「被拒绝」。

一个自动机，总能够回答某种形式的「对于给定的输入字符串 S，判断其是否满足条件 P」的问题。在本题中，条件 P 即为「构成合法的表示数值的字符串」。

自动机驱动的编程，可以被看做一种暴力枚举方法的延伸：它穷尽了在任何一种情况下，对应任何的输入，需要做的事情。

自动机在计算机科学领域有着广泛的应用。在算法领域，它与大名鼎鼎的字符串查找算法「KMP」算法有着密切的关联；在工程领域，它是实现「正则表达式」的基础。

*/
func isNumber(s string) bool {
	if len(s) == 0 {
		return false
	}
	s = strings.TrimSpace(s)
	char := []rune(s)
	var numeric bool
	var tmpNumeric bool

	numeric, char = scanInteger(char)
	if len(char) > 0 && char[0] == '.' {
		char = char[1:]
		tmpNumeric, char = scanUnsignedInteger(char)
		numeric = tmpNumeric && numeric //这里为什么要这样？答：.123= 0.123 ; 123. = 123.0
	}
	if len(char) > 0 && (char[0] == 'e' || char[0] == 'E') {
		char = char[1:]
		tmpNumeric, char = scanInteger(char)
		numeric = numeric && tmpNumeric
	}

	return numeric && len(char) == 0
}

func scanUnsignedInteger(char []rune) (bool, []rune) {
	before := len(char)
	for len(char) > 0 && char[0] >= '0' && char[0] <= '9' {
		char = char[1:]
	}
	return before > len(char), char
}

func scanInteger(char []rune) (bool, []rune) {
	if len(char) == 0 {
		return false, char
	}
	if char[0] == '-' || char[0] == '+' {
		char = char[1:]
	}
	return scanUnsignedInteger(char)
}
