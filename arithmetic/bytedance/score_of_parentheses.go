package bytedance

/**
给定一个平衡括号字符串 S，按下述规则计算该字符串的分数：

() 得 1 分。
AB 得 A + B 分，其中 A 和 B 是平衡括号字符串。
(A) 得 2 * A 分，其中 A 是平衡括号字符串。


示例 1：

输入： "()"
输出： 1
示例 2：

输入： "(())"
输出： 2
示例 3：

输入： "()()"
输出： 2
示例 4：

输入： "(()(()))"
输出： 6


提示：

S 是平衡括号字符串，且只含有 ( 和 ) 。
2 <= S.length <= 50
*/
/**
事实上，我们可以发现，只有 () 会对字符串 S 贡献实质的分数，
其它的括号只会将分数乘二或者将分数累加。
因此，我们可以找到每一个 () 对应的深度 x，那么答案就是 2^x 的累加和。


其实就是（）就是1
（）（）就是两个1 相加
(())就是2的1次方，这个1是 指外面的（） 如果是（（（））），那么1就是2了，就是2的2次方。
之所以是2是因为上面规定了2*A，如果是三个就是2*2*A，A是1，就是2的2次方
*/
func scoreOfParentheses(S string) int {
	o, c := 0, uint(0)
	for i, b := range []byte(S) {
		if b == '(' {
			c++
		} else {
			c--
			if b == ')' && S[i-1] == '(' {
				o += 1 << c
			}
		}
	}
	return o
}

/**
利用栈的求法
*/
func scoreOfParentheses2(S string) int {
	return getScore(S)
}

func getScore(s string) int {
	if len(s) > 4 {
		var tmp []rune
		lenS := len(s)
		for i, v := range []rune(s) {
			lenTmp := len(tmp)
			if v == ')' {
				if tmp[lenTmp-1] == '(' {
					tmp = tmp[:lenTmp-1]
				} else {
					tmp = append(tmp, v)
				}
				if len(tmp) != 0 {
					continue
				}
				if i != lenS-1 { //第一个（ 没有了以后，还未到字符串结尾，说明符合 ()()
					return getScore(s[:i+1]) + getScore(s[i+1:])
				} else { //到了结尾，说明符合(())
					return 2 * getScore(s[1:i])
				}
			} else {
				tmp = append(tmp, v)
			}
		}

	}
	if s == "()" {
		return 1
	}
	if s == "()()" {
		return 2
	}
	if s == "(())" {
		return 2
	}
	return 0
}
