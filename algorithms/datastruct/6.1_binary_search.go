package datastruct

import "fmt"

/**
二分查找是我们目前为止遇到的第一个时间复杂度为 O(logn) 的算法。
后面章节我们还会讲堆、二叉树的操作等等，它们的时间复杂度也是 O(logn)。

局限性：只能查顺序数组，数据量太小并无优势，数据量太大所需数组也大。
*/
func binarySearchDemo(a []int) {
	idx := binarySearch(a, len(a), 5)
	fmt.Printf("binarySearch idx:%v \n", idx)
}

func binarySearch(a []int, n int, v int) (mid int) {
	//最大，最小，中间位置值
	//注意退出条件是<=
	//（low+high） 越界问题 ，转为low + ((high - low) >> 1)
	if n <= 1 {
		return
	}
	low := 0
	high := n - 1
	for low <= high {
		mid = low + ((high - low) >> 1)
		if a[mid] == v {
			return
		} else if a[mid] > v {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return
}
func binarySearch2(a []int, v int) (mid int) {
	if len(a) <= 1 {
		return
	}
	low, high := 0, len(a)-1
	for low <= high {
		mid = low + ((high - low) >> 1)
		if a[mid] == v {
			return
		}

		if a[mid] > v {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return
}

/**
递归实现
*/
func recBinarySearchDemo(a []int) {
	idx := recBinarySearch(a, a[0], a[len(a)-1], 5)
	fmt.Printf("recBinarySearch idx:%v\n", idx)
}
func recBinarySearch(a []int, low, high, v int) (mid int) {
	if low > high {
		return -1
	}
	mid = low + ((high - low) >> 1)
	if a[mid] == v {
		return
	}
	if a[mid] > v {
		return recBinarySearch(a, low, mid-1, v)
	} else {
		return recBinarySearch(a, mid+1, high, v)
	}
}

/**
如何编程实现“求一个数的平方根”？要求精确到小数点后 6 位。

*/

func squareRootDemo(m float64, n int) {
	//n 小数点后几位
	v := squareRoot(m, n)
	fmt.Printf("v :%v \n", v)
}

/**
大概值     | 商 | 平均值
-------- | ----- | -----
1  | 2/1=2 |（1+2）/ 2=1.5
1.5  | 2/1.5=1.3333 |（1.3333+1.5）/ 2=1.4167
1.4167 | 2/1.4167=1.4118 | （1.4167+1.4168）/ 2=1.4142
... | ... | ...
*/
func squareRoot(m float64, n int) (ret float64) {
	count := 0
	ret = m / 2
	for count <= n {
		ret = ((m / ret) + ret) / m
		count++
	}

	return
}
