package leetcode

/**

在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。
数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。
输入：
[2, 3, 1, 0, 2, 5, 3]
输出：2 或 3
*/
/**
特殊情况： 数组中的最大值小于等于数组的长度的话，可以用这种方法。
*/
func findRepeatNumber1(nums []int) int {

	if len(nums) == 0 {
		return -1
	}

	idx := 0
	length := len(nums) - 1
	for {

		if idx == nums[idx] {
			idx++
			continue
		}
		if nums[idx] == nums[nums[idx]] {
			return nums[idx]
		}

		nums[idx], nums[nums[idx]] = nums[nums[idx]], nums[idx]

		if idx == length {
			break
		}
	}
	return -1
}

/**
一般情况 ：数组中任意数据

解法1，排序后，数组中相邻的数据比较。
解法2，利用map。
*/
