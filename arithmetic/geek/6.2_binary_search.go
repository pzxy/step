package main

import "fmt"

/**
二分查找变体
*/
func main() {
	a := []int{1, 2, 3, 5, 5, 6, 7, 8, 9, 10}
	searchFirstVDemo(a, 5)
	searchLastVDemo(a, 5)
	searchFirstGreaterVDemo(a, 5)
	searchLastLessVDemo(a, 4)
}

/**
变体一：查找第一个值等于给定值的元素
*/
func searchFirstVDemo(a []int, n int) {
	idx := bSearchFirstV(a, len(a), n)
	fmt.Printf("bSearchFirstV idx:%v \n", idx)

}
func bSearchFirstV(a []int, n int, v int) (mid int) {
	low := 0
	high := n - 1
	for low <= high {
		mid = low + ((high - low) >> 1)
		if a[mid] < v {
			low = mid + 1
			continue
		}
		if a[mid] > v {
			high = mid - 1
			continue
		}
		//mid ==0 说明是第一个，一定是我们要找的(因为除了大于小于，只有等于)，如果前一个mid-1等于v，说明不是第一个给定值。则 high = mid - 1
		if mid == 0 || a[mid-1] != v {
			return
		}
		high = mid - 1
	}
	return -1
}

/**
变体二：查找最后一个值等于给定值的元素
*/
func searchLastVDemo(a []int, n int) {
	idx := bSearchLastV(a, len(a), n)
	fmt.Printf("bSearchLastV idx:%v \n", idx)

}

func bSearchLastV(a []int, n int, v int) (mid int) {
	low := 0
	high := n - 1
	for low <= high {
		mid = low + ((high - low) >> 1)
		if a[mid] > v {
			high = mid - 1
			continue

		}
		if a[mid] < v {
			low = mid + 1
			continue
		}
		if mid == n-1 || a[mid+1] != v {
			return
		}
	}
	return -1
}

/**
变体三：查找第一个大于等于给定值的元素
*/
func searchFirstGreaterVDemo(a []int, n int) {
	idx := bSearchFirstGreaterV(a, len(a), n)
	fmt.Printf("bSearchFirstGreaterV idx:%v \n", idx)
}
func bSearchFirstGreaterV(a []int, n, v int) (mid int) {
	low := 0
	high := n - 1
	for low <= high {
		mid = low + ((high - low) >> 1)
		if a[mid] < v {
			low = mid + 1
			continue
		}
		if mid == n-1 || a[mid-1] < v {
			return
		}
		high = mid - 1

	}
	return -1
}

/**
变体四：查找最后一个小于等于给定值的元素
*/
func searchLastLessVDemo(a []int, n int) {
	idx := bSearchLastLessV(a, len(a), n)
	fmt.Printf("bSearchLastLessV idx:%v \n", idx)
}

func bSearchLastLessV(a []int, n, v int) (mid int) {
	low := 0
	high := n - 1
	for low <= high {
		mid = low + ((high - low) >> 1)
		if a[mid] > v {
			high = mid - 1
			continue
		}
		if mid == 0 || a[mid+1] > v {
			return
		}
		low = mid + 1
	}
	return -1
}
