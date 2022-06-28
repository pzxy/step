package bytedance

/**
给定一个非负整数 num。对于 0 ≤ i ≤ num 范围中的每个数字 i ，计算其二进制数中的 1 的数目并将它们作为数组返回。

示例 1:

输入: 2
输出: [0,1,1]
示例 2:

输入: 5
输出: [0,1,1,2,1,2]
*/
func countBits(num int) (ret []int) {
	tmp := 0
	for tmp <= num {
		cur := tmp
		count := 0
		for cur != 0 {
			if cur&1 == 1 {
				count++
			}
			cur >>= 1
		}
		ret = append(ret, count)
		tmp++
	}
	return
}
