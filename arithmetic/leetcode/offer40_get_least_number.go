package leetcode

/**
输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。

示例 1：

输入：arr = [3,2,1], k = 2
输出：[1,2] 或者 [2,1]
示例 2：

输入：arr = [0,1,2,1], k = 1
输出：[0]

限制：

0 <= k <= arr.length <= 10000
0 <= arr[i]<= 10000

*/

func getLeastNumbers(arr []int, k int) []int {
	if len(arr) == 0 {
		return arr
	}
	return parition2(arr, k)
}

/**
解法1：快排思想，比k小的都放到k前面
*/

func parition2(arr []int, k int) []int {
	i := 0
	for j := 0; j < len(arr); j++ {
		if arr[j] < k {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}

	}
	return arr[:i]
}
