package leetcode

/**
输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。

示例 1：

输入：arr = [3,2,1], k = 2
输出：[1,2] 或者 [2,1]
示例 2：

输入：arr = [0,1,2,1], k = 1
输出：[0]

限制：

0 <= k <= arr.length <= 10000
0 <= arr[i]<= 10000

*/

func getLeastNumbers(arr []int, k int) []int {
	if len(arr) == 0 || len(arr) == k {
		return arr
	}
	if len(arr) < k {
		return []int{}
	}

	return parition2(arr, k)
}

/**
解法1：快排思想，比k小的都放到k前面
求最小的k个数，用了快速排序的思想，使用快速排序的分区原地排序法，用原地排序的方式，分区点左侧的值小于等于分区点，分区点右侧的值大于等于分区点。针对分区点，有下面三种情况
若分区点的数组下标小于k，则k位于分区点的右侧，则对分区点右侧进行快速排序
若分区点的数组下标大于k，则k位于分区点的左侧，对分区点左侧进行快速排序
若分区点的数组小表等于k，则k的值等于当前坐在位置
*/

func parition2(arr []int, k int) []int {
	i := 0
	start := 0
	end := len(arr) - 1
	for {
		for j, v := range arr[start:end] {
			if v < arr[end] {
				arr[i], arr[start+j] = arr[start+j], arr[i]
				i++
			}
		}
		arr[i], arr[end] = arr[end], arr[i]

		if i > k {
			end = i - 1
		}
		if i < k {
			start = i + 1
		}
		if i == k {
			break
		}
		i = start
	}

	return arr[:k]
}

/**
解法2：用大顶堆，最上面是最大的。如果要前k小。那么创建大小为k的堆。堆顶是堆中最大的，里面的一定都比他小。
*/

//先实现堆化方法。
func getLeastNumbers2(arr []int, k int) []int {

	if len(arr) == 0 {
		return arr
	}
	if k <= 0 {
		return []int{}
	}
	if len(arr) == k {
		return arr
	}
	if len(arr) < k {
		return []int{}
	}

	h := Heap{a: make([]int, k+1), len: k, count: 0}
	for h.count != k {
		h.count++
		h.a[h.count] = arr[h.count-1]
	}
	h.build()
	for i := h.count; i < len(arr); i++ {
		if arr[i] > h.a[1] {
			continue
		}
		h.a[1] = arr[i]
		h.heapify(1)
	}
	return h.a[1:]
}
func (h *Heap) heapify(i int) {
	if h.count == 0 {
		return
	}
	maxPos := i
	for {
		if 2*i <= h.len && h.a[i] < h.a[2*i] {
			maxPos = 2 * i
		}
		if 2*i+1 <= h.len && h.a[maxPos] < h.a[2*i+1] {
			maxPos = 2*i + 1
		}
		if maxPos == i {
			break
		}
		h.a[maxPos], h.a[i] = h.a[i], h.a[maxPos]
		i = maxPos
	}
}

func (h *Heap) build() {
	for i := h.len >> 1; i >= 1; i-- {
		h.heapify(i)
	}
}

//type Heap struct {
//	a     []int
//	count int
//}
