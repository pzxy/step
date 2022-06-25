package leetcode

/**
统计一个数字在排序数组中出现的次数。

示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: 2
示例2:

输入: nums = [5,7,7,8,8,10], target = 6
输出: 0

限制：

0 <= 数组长度 <= 50000
*/

func search(nums []int, target int) (ret int) {
	if len(nums) == 0 {
		return
	}
	l, q, r := 0, 0, len(nums)-1
	for l <= r {
		q = l + (r-l)>>1
		if nums[q] < target {
			l = q + 1
			continue
		}
		if nums[q] > target {
			r = q - 1
			continue
		}
		ret = +1
		break
	}
	q1, q2 := q+1, q-1
	for q1 < len(nums) {
		if nums[q1] == target {
			q1++
			ret += 1
			continue
		}
		break
	}
	for 0 <= q2 {
		if nums[q2] == target {
			q2--
			ret += 1
			continue
		}
		break
	}
	return
}
