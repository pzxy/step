package leetcode

/**
输入一个递增排序的数组和一个数字s，在数组中查找两个数，使得它们的和正好是s。如果有多对数字的和等于s，则输出任意一对即可。

示例 1：

输入：nums = [2,7,11,15], target = 9
输出：[2,7] 或者 [7,2]
示例 2：

输入：nums = [10,26,30,31,47,60], target = 40
输出：[10,30] 或者 [30,10]

限制：

1 <= nums.length <= 10^5
1 <= nums[i]<= 10^6

*/
func twoSum(nums []int, target int) (ret []int) {
	if len(nums) <= 1 {
		return
	}
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]+nums[right] > target { //大玉目标值，只能右边减小，因为左边移动是增的大
			right--
			continue
		}
		if nums[left]+nums[right] < target {
			left++
			continue
		}
		ret = make([]int, 2)
		ret[0], ret[1] = nums[left], nums[right]
		return
	}
	return
}
