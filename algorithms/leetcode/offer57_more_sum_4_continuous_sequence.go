package leetcode

/**
输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。

序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。

示例 1：

输入：target = 9
输出：[[2,3,4],[4,5]]
示例 2：

输入：target = 15
输出：[[1,2,3,4,5],[4,5,6],[7,8]]

限制：

1 <= target <= 10^5
*/
func findContinuousSequence(target int) (ret [][]int) {
	pre := 1
	hind, sum, tmp := pre, pre, []int{pre}
	for pre <= target>>1 {
		if sum == target {
			ret = append(ret, tmp)
			pre++
			hind, sum, tmp = pre, pre, []int{pre}
			continue
		}
		if sum > target {
			pre++
			hind, sum, tmp = pre, pre, []int{pre}
			continue
		}
		hind++
		sum += hind
		tmp = append(tmp, hind)
	}
	return
}

func findContinuousSequence2(target int) (res [][]int) {
	left := 1
	right := 2
	for left < right {
		sum := (left + right) * (right - left + 1) / 2 //高斯速算
		if sum == target {
			var list []int
			for i := left; i <= right; i++ {
				list = append(list, i)
			}
			res = append(res, list)
			left++
			continue
		}
		if sum < target {
			right++
			continue
		}
		left++
	}
	return
}
