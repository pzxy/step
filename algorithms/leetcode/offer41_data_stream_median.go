package leetcode

/**
如何得到一个数据流中的中位数？
如果从数据流中读出奇数个数值，那么中位数就是所有数值排序之后位于中间的数值。
如果从数据流中读出偶数个数值，那么中位数就是所有数值排序之后中间两个数的平均值。

例如，

[2,3,4]的中位数是 3

[2,3] 的中位数是 (2 + 3) / 2 = 2.5

设计一个支持以下两种操作的数据结构：

void addNum(int num) - 从数据流中添加一个整数到数据结构中。
double findMedian() - 返回目前所有元素的中位数。
示例 1：

输入：
["MedianFinder","addNum","addNum","findMedian","addNum","findMedian"]
[[],[1],[2],[],[3],[]]
输出：[null,null,null,1.50000,null,2.00000]
示例 2：

输入：
["MedianFinder","addNum","findMedian","addNum","findMedian"]
[[],[2],[],[3],[]]
输出：[null,null,2.00000,null,2.50000]

限制：

最多会对addNum、findMedia进行50000次调用。

*/
//type Heap struct {
//	a     []int
//	len   int
//	count int
//}
type MedianFinder struct {
	headBigTop   Heap
	headSmallTop Heap
}

/** initialize your data structure here. */
func Constructor41() MedianFinder {
	return MedianFinder{
		headBigTop:   Heap{a: make([]int, 1)},
		headSmallTop: Heap{a: make([]int, 1)},
	}
}

func (this *MedianFinder) AddNum(num int) {

	if this.headBigTop.count > this.headSmallTop.count {
		if num < this.headBigTop.top() {
			this.headSmallTop.insert4SmallTop2UpHeap(this.headBigTop.top())
			this.headBigTop.replace4BigTop2downHeap(num)
		} else {
			this.headSmallTop.insert4SmallTop2UpHeap(num)
		}
		return
	}

	if this.headBigTop.count < this.headSmallTop.count {
		if num > this.headSmallTop.top() {
			this.headBigTop.insert4BigTop2UpHeap(this.headSmallTop.top())
			this.headSmallTop.replace4SmallTop2downHeap(num)
		} else {
			this.headBigTop.insert4BigTop2UpHeap(num)
		}
		return
	}

	if this.headSmallTop.top() > num {
		this.headBigTop.insert4BigTop2UpHeap(num)
		return
	}
	this.headSmallTop.insert4SmallTop2UpHeap(num)

}

func (this *MedianFinder) FindMedian() float64 {
	if this.headSmallTop.count == 0 && this.headBigTop.count == 0 {
		return 0
	}
	if this.headSmallTop.count > this.headBigTop.count {
		return float64(this.headSmallTop.a[1])
	}
	if this.headSmallTop.count < this.headBigTop.count {
		return float64(this.headBigTop.a[1])
	}

	val := (float64(this.headBigTop.top()) + float64(this.headSmallTop.top())) / 2
	//num, _ := strconv.ParseFloat(fmt.Sprintf("%.5f", val), 64)
	return val
}

func (h *Heap) top() int {
	if h.count == 0 {
		return 0
	}
	return h.a[1]
}

func (h *Heap) insert4BigTop2UpHeap(num int) { //向上堆化
	h.a = append(h.a, num)
	h.count++
	i := h.count
	for i/2 > 0 && h.a[i] > h.a[i/2] {
		ii := i / 2
		h.a[i], h.a[ii] = h.a[ii], h.a[i]
		i = ii
	}
}

func (h *Heap) replace4BigTop2downHeap(num int) { //向下堆化
	h.a[1] = num
	i := 1
	maxPos := i
	for {
		if i*2 <= h.count && h.a[i] < h.a[2*i] {
			maxPos = 2 * i
		}
		if i*2+1 <= h.count && h.a[maxPos] < h.a[2*i+1] {
			maxPos = 2*i + 1
		}
		if i == maxPos {
			break
		}
		h.a[i], h.a[maxPos] = h.a[maxPos], h.a[i]
		i = maxPos
	}
}

func (h *Heap) replace4SmallTop2downHeap(num int) { //向下堆化
	h.a[1] = num
	i := 1
	maxPos := i
	for {
		if i*2 <= h.count && h.a[i] > h.a[2*i] {
			maxPos = 2 * i
		}
		if i*2+1 <= h.count && h.a[maxPos] > h.a[2*i+1] {
			maxPos = 2*i + 1
		}
		if i == maxPos {
			break
		}
		h.a[i], h.a[maxPos] = h.a[maxPos], h.a[i]
		i = maxPos
	}
}

func (h *Heap) insert4SmallTop2UpHeap(num int) {
	h.a = append(h.a, num)
	h.count++
	i := h.count
	for i/2 > 0 && h.a[i] < h.a[i/2] {
		ii := i / 2
		h.a[i], h.a[ii] = h.a[ii], h.a[i]
		i = ii
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
