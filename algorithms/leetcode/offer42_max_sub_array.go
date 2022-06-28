package leetcode

import "math"

/**
输入一个整型数组，数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。

要求时间复杂度为O(n)。

示例1:

输入: nums = [-2,1,-3,4,-1,2,1,-5,4]
输出: 6
解释:连续子数组[4,-1,2,1] 的和最大，为6。

提示：

1 <=arr.length <= 10^5
-100 <= arr[i] <= 100
*/
//解法一
func maxSubArray(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}
	maxSum := nums[0]
	tmpSum := maxSum
	for _, v := range nums[1:] {
		tmpSum = v + tmpSum
		if tmpSum < v { //小于本次的值，那么应该放弃前面加起来的和
			tmpSum = v
			if v > maxSum {
				maxSum = v
			}
			continue
		}
		if tmpSum > maxSum {
			maxSum = tmpSum
		}

	}
	return maxSum
}

//解法二 动态规划

func maxSubArray4dp(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = -math.MaxInt64
	}

	dp[0] = nums[0]
	ret := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = int(math.Max(float64(nums[i]), float64(nums[i]+dp[i-1])))
		if dp[i] > ret {
			ret = dp[i]
		}
	}
	return ret
}
