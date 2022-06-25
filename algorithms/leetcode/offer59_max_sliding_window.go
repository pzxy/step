package leetcode

import "math"

/**
给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。

示例:

输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
输出: [3,3,5,5,6,7]
解释:

  滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7

提示：

你可以假设 k 总是有效的，在输入数组不为空的情况下，1 ≤ k ≤输入数组的大小。

*/
/**
记录最大值；
如果滑动丢失的是最大值，则遍历滑动窗口更新存储的最大值；
如果滑动新增的值是最大值，则用它替换存储的最大值；
*/
func maxSlidingWindow(nums []int, k int) (ret []int) {
	lenNums := len(nums)

	getMax := func(n []int) int {
		m := -math.MaxInt64
		for _, v := range n {
			if m < v {
				m = v
			}
		}
		return m
	}
	if lenNums == 0 {
		return
	}
	if lenNums < k {
		ret = append(ret, getMax(nums))
		return
	}

	left, right := 0, k-1                //0-2
	maxV := getMax(nums[left : right+1]) //0-2
	ret = append(ret, maxV)
	right++
	for right < lenNums { //0-3
		if nums[left] == maxV {
			maxV = getMax(nums[left+1 : right+1]) //1-4
		}
		if nums[right] > maxV {
			maxV = nums[right]
		}
		ret = append(ret, maxV)
		left++
		right++
	}
	return
}
