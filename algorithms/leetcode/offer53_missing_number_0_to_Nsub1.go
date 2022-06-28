package leetcode

/**
一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围0～n-1之内。
在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出这个数字。

示例 1:

输入: [0,1,3]
输出: 2
示例2:

输入: [0,1,2,3,4,5,6,7,9]
输出: 8

限制：

1 <= 数组长度 <= 10000
*/

func missingNumber(nums []int) (ret int) {
	lenNums := len(nums)
	if lenNums == 0 {
		return
	}
	l, p, r := 0, 0, lenNums-1

	for r >= l {
		p = l + (r-l)>>1
		if p == nums[p] {
			l = p + 1
			continue
		}
		r = p - 1
	}

	return l
}
func missingNumber2(nums []int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		m := (i + j) >> 1
		if nums[m] == m {
			i = m + 1
		} else {
			j = m - 1
		}
	}
	return i
}

func missingNumber3(nums []int) int {
	l, p, r := 0, 0, len(nums)
	for l < r {
		p = l + (r-l)>>1
		if p != nums[p] {
			//nums是有序数组，如果q和数字不相同就在左边查找
			r = p
			continue
		}
		l = p + 1
	}
	return l
}
