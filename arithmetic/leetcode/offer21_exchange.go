package leetcode

/**
输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数位于数组的前半部分，所有偶数位于数组的后半部分。

示例：

输入：nums =[1,2,3,4]
输出：[1,3,2,4]
注：[3,1,2,4] 也是正确的答案之一。

提示：
1 <= nums.length <= 50000
1 <= nums[i] <= 10000
*/

func exchange2(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	var odd, even []int
	for _, v := range nums {
		if v&1 == 0 {
			even = append(even, v)
			continue
		}
		odd = append(odd, v)
	}
	odd = append(odd, even...)
	return odd
}

func exchange(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	head := 0
	tail := len(nums) - 1
	for head < tail {
		for head < tail && nums[head]&1 == 1 { //for比if continue速度快应该是少了continue的缘故
			head++
		}
		for head < tail && nums[tail]&1 == 0 {
			tail--
		}
		nums[head], nums[tail] = nums[tail], nums[head]
		head++
		tail--
	}
	return nums
}
