package leetcode

/**
输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否为该栈的弹出顺序。
假设压入栈的所有数字均不相等。
例如，序列 {1,2,3,4,5} 是某栈的压栈序列，
序列 {4,5,3,2,1} 是该压栈序列对应的一个弹出序列，
但 {4,3,5,1,2} 就不可能是该压栈序列的弹出序列。


示例 1：

输入：pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
输出：true
解释：我们可以按以下顺序执行：
push(1), push(2), push(3), push(4), pop() -> 4,
push(5), pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1
示例 2：

输入：pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
输出：false
解释：1 不能在 2 之前弹出。


提示：

0 <= pushed.length == popped.length <= 1000
0 <= pushed[i], popped[i] < 1000
pushed是popped的排列。

*/
//validate vt. 确认；使生效；证实，验证
func validateStackSequences2(pushed []int, popped []int) bool {
	lpushed := len(popped)
	lpopped := len(pushed)
	if lpopped != lpushed {
		return false
	}

	popIdx, pushIdx := 0, 0
	var tmpStack []int
	for lpopped > popIdx {
		if lpushed > pushIdx && pushed[pushIdx] == popped[popIdx] {
			popIdx++
			pushIdx++
			continue
		}
		if len(tmpStack) != 0 && tmpStack[len(tmpStack)-1] == popped[popIdx] {
			tmpStack = tmpStack[:len(tmpStack)-1]
			popIdx++
			continue
		}

		if lpushed > pushIdx {
			tmpStack = append(tmpStack, pushed[pushIdx])
			pushIdx++
			continue
		}
		break
	}
	if len(tmpStack) == 0 {
		return true
	}
	return false
}

func validateStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, 0)
	i := 0
	for _, v := range pushed {
		stack = append(stack, v)
		for len(stack) != 0 && stack[len(stack)-1] == popped[i] { //这种写法很棒，等于for和if一起用
			stack = stack[:len(stack)-1]
			i++
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}
