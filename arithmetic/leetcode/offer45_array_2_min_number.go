package leetcode

import "strconv"

/**
输入一个非负整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。

示例 1:

输入: [10,2]
输出: "102"
示例2:

输入: [3,30,34,5,9]
输出: "3033459"

提示:

0 < nums.length <= 100
说明:

输出结果可能非常大，所以你需要返回一个字符串而不是整数
拼接起来的数字可能会有前导 0，最后结果不需要去掉前导 0，比如0123456，也是对的
*/
func minNumber(nums []int) (ret string) {

	length := len(nums)
	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			if strconv.Itoa(nums[j])+strconv.Itoa(nums[j+1]) < strconv.Itoa(nums[j+1])+strconv.Itoa(nums[j]) {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
		ret += strconv.Itoa(nums[length-i-1])
	}
	//if nums[length-1] == 0 { //1,2,3,4,5,6,7,8,9,0 这样的话，如果是0123456789也没问题的
	//	ret = ""
	//}
	return
}
