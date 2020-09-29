package leetcode

/**
一个整型数组 nums 里除两个数字之外，其他数字都出现了两次。
请写程序找出这两个只出现一次的数字。要求时间复杂度是O(n)，空间复杂度是O(1)。

示例 1：

输入：nums = [4,1,4,6]
输出：[1,6] 或 [6,1]
示例 2：

输入：nums = [1,2,10,4,1,4,3,3]
输出：[2,10] 或 [10,2]

限制：

2 <= nums.length <= 10000
*/
func singleNumbers(nums []int) []int {
	var ret int
	for _, v := range nums {
		ret ^= v //将相同的过滤掉
	}
	dev := 1
	for dev&ret == 0 { //找到第一个1
		dev <<= 1
	}
	a, b := 0, 0
	for _, v := range nums {
		if dev&v == 0 { //两个不同的肯定不会分到一块
			a ^= v
		} else {
			b ^= v
		}

	}
	return []int{a, b}
}
