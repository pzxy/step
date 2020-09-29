package leetcode

/**
在一个数组 nums 中除一个数字只出现一次之外，其他数字都出现了三次。请找出那个只出现一次的数字。

示例 1：

输入：nums = [3,4,3,3]
输出：4
示例 2：

输入：nums = [9,1,7,9,7,9,7]
输出：1

限制：

1 <= nums.length <= 10000
1 <= nums[i] < 2^31

*/

func singleNumbers4Thrice2(nums []int) int {

	res := 0
	for i := 0; i < 32; i++ {
		count := 0
		bit := 1 << i
		for _, num := range nums {
			if num&bit != 0 {
				count++
			}
		}
		if count%3 != 0 {
			res |= bit
		}
	}
	return res
}

func singleNumbers4Thrice(nums []int) (ret int) {
	for i := 0; i < 32; i++ {
		bit := 1 << i
		count := 0
		for _, v := range nums {
			if bit&v == 0 {
				continue
			}
			count++
		}
		if count%3 != 0 {
			ret |= bit
		}
	}
	return
}
