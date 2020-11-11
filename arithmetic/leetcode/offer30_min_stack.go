package leetcode

/**
定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。

示例:

MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.min();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.min();   --> 返回 -2.

提示：

各函数的调用总次数不超过 20000 次
*/

type MinStack struct {
	Stack     []int //-2，0，-3
	HelpStack []int //-2，-2，-3
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.Stack = append(this.Stack, x)
	l := len(this.HelpStack)
	if l == 0 {
		this.HelpStack = append(this.HelpStack, x)
		return
	}
	v := this.HelpStack[l-1]
	if v > x {
		this.HelpStack = append(this.HelpStack, x)
	} else {
		this.HelpStack = append(this.HelpStack, v)
	}
}

func (this *MinStack) Pop() {
	sl := len(this.Stack)
	hl := len(this.HelpStack)
	if sl == 0 || hl == 0 {
		return
	}
	this.HelpStack = this.HelpStack[:hl-1]
	this.Stack = this.Stack[:sl-1]
}

func (this *MinStack) Top() int {
	sl := len(this.Stack)
	hl := len(this.HelpStack)
	if sl == 0 || hl == 0 {
		return -1
	}
	return this.Stack[sl-1]
}

func (this *MinStack) Min() int {
	sl := len(this.Stack)
	hl := len(this.HelpStack)
	if sl == 0 || hl == 0 {
		return -1
	}
	return this.HelpStack[hl-1]
}
