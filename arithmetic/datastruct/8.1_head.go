package datastruct

import "fmt"

/**
数组中下标为 i 的节点的左子节点，就是下标为 i∗2 的节点，右子节点就是下标为 i∗2+1 的节点，父节点就是下标为 i/2​ 的节点。
*/
type head struct {
	//go里面用的切片，没必要用长度n限制自身，先这样写吧。
	a     []int
	n     int //能存数据个数，不包括a[0]
	count int //存了多少了
}

func newHead(n int) *head {
	return &head{a: make([]int, n+1), n: n, count: 0}
}
func headInsert() {
	fmt.Println("插入元素")

	h := newHead(11)
	h.insert(1)
	h.insert(2)
	h.insert(3)
	h.insert(4)
	h.insert(5)
	h.insert(6)
	fmt.Println(h.a)
	fmt.Println()

}
func headBuild() {
	fmt.Println("数组堆化")

	h := newHead(10)
	h.a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	h.count = 10
	h.build()
	fmt.Println(h.a)
	fmt.Println()
}

func headDeleteTop() {
	fmt.Println("删除堆顶元素")
	h := newHead(10)
	h.a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	h.count = 10
	h.build()
	fmt.Println(h.a)
	fmt.Println()
	h.deleteTop()
	fmt.Println(h.a)
	fmt.Println()
	h.deleteTop()
	fmt.Println(h.a)
	fmt.Println()
}
func headSort() {
	fmt.Println("堆排序顺序输出")
	h := newHead(10)
	h.a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	h.count = 10
	h.build()
	fmt.Println(h.a)
	fmt.Println()
	h.sort()
	fmt.Println(h.a)
	fmt.Println()
}

func (h *head) insert(data int) {
	if h.count >= h.n {
		return
	}
	h.count++
	h.a[h.count] = data
	i := h.count
	ii := i / 2
	for ii > 0 && h.a[i] > h.a[ii] {
		h.a[i], h.a[ii] = h.a[ii], h.a[i]
		i = ii
		ii = i / 2
	}
}

//建堆，一种是不断insert，一种就是一个好的数组然后build，确保每个跟节点对于自节点来说都是大的
func (h *head) build() {
	//我们对下标从 n/2 开始到 1 的数据进行堆化，下标是 2n​+1 到 n 的节点是叶子节点，我们不需要堆化。
	//因为叶子节点往下堆化只能自己跟自己比较
	for i := h.n / 2; i >= 1; i-- {
		h.heapify(h.n, i)
	}
}

/**

 */
func (h *head) heapify(len int, i int) {
	maxPos := i
	for {
		if 2*i <= len && h.a[i] < h.a[2*i] { //小于左子节点，获取最大值位置。
			maxPos = 2 * i
		}
		if 2*i+1 <= len && h.a[maxPos] < h.a[2*i+1] { //用最大值再和右边子节点比较。
			maxPos = 2*i + 1
		}
		if maxPos == i {
			break
		}

		h.a[i], h.a[maxPos] = h.a[maxPos], h.a[i]

		i = maxPos
	}
}

func (h *head) deleteTop() {
	if h.count <= 0 {
		return
	}
	h.a[1] = h.a[h.count] //最后一个元素覆盖堆顶元素。
	h.a = h.a[:h.count]
	h.count--

	h.heapify(h.count, 1)
}

/**
建堆结束之后，数组中的数据已经是按照大顶堆的特性来组织的。数组中的第一个元素就是堆顶，也就是最大的元素。
我们把它跟最后一个元素交换，那最大元素就放到了下标为 n 的位置。

这个过程有点类似上面讲的“删除堆顶元素”的操作，当堆顶元素移除之后，
我们把下标为 n 的元素放到堆顶，然后再通过堆化的方法，将剩下的 n−1 个元素重新构建成堆。
堆化完成之后，我们再取堆顶的元素，放到下标是 n−1 的位置，
一直重复这个过程，直到最后堆中只剩下标为 1 的一个元素，排序工作就完成了。
*/
func (h *head) sort() {
	h.build()
	k := h.n
	for k > 1 {
		h.a[1], h.a[k] = h.a[k], h.a[1]
		k--
		h.heapify(k, 1)
	}
}
