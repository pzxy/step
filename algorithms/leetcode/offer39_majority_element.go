package leetcode

/**
数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。

示例1:

输入: [1, 2, 3, 2, 2, 2, 5, 4, 2]
输出: 2

限制：

1 <= 数组长度 <= 50000

*/
/**
解法1
*/
func majorityElement(nums []int) int { //快排后，中间的数就是超过长度的一半的数。
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		return nums[0]
	}
	return parition(nums, 0, len(nums)-1)
}

func parition(nums []int, p, r int) int {
	if p > r {
		return -1
	}

	pivot := nums[r]
	i := p
	for j, v := range nums[p:r] {
		if v < pivot {
			nums[i], nums[j+p] = nums[j+p], nums[i]
			i++
		}
	}
	nums[r], nums[i] = nums[i], nums[r]
	if i > len(nums)>>1 { //如果i大与数组长度的1/2，那么【超过长度的一半的数】一定在前面部分数组中
		return parition(nums, p, i-1)
	}
	if i < len(nums)>>1 {
		return parition(nums, i+1, r)
	}
	return nums[i]
}

/**
解法二：等于说一个数出现的次数比其他值出现的次数和还要多
[1, 2, 3, 2, 2, 2, 5, 4, 2]
两个值，一个保存数，一个保存次数。不同就减去1，相同就加1，为0时替换并设1
依次为 [1,1][2,1][3,1][2,1][2,2][2,3][2,2][2,1][2,2]
*/
func majorityElement2(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		return nums[0]
	}
	ret, count := nums[0], 1
	for _, v := range nums[1:] {
		if ret == v {
			count++
			continue
		}
		count--
		if count == 0 {
			ret, count = v, 1
		}
	}
	return ret
}
