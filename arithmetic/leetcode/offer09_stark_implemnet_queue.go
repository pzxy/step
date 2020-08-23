package leetcode

/**
用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，
分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead操作返回 -1 )


*/
/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
type CQueue struct {
	dequeue []int
	enqueue []int
}

func Constructor() CQueue {
	return CQueue{
		enqueue: make([]int, 0),
		dequeue: make([]int, 0),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.enqueue = append(this.enqueue, value)
}

func (this *CQueue) DeleteHead() int {
	deLength := len(this.dequeue)
	enLength := len(this.enqueue)
	if deLength == 0 && enLength == 0 {
		return -1
	}

	if deLength == 0 {
		for i := enLength - 1; i >= 0; i-- {
			this.dequeue = append(this.dequeue, this.enqueue[i])
		}
		this.enqueue = make([]int, 0)
	}

	tailIdx := len(this.dequeue) - 1
	ret := this.dequeue[tailIdx]
	this.dequeue = this.dequeue[:tailIdx]
	return ret
}
