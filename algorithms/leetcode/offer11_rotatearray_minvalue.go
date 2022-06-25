package leetcode

/**
把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。

输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。

例如，数组[3,4,5,1,2] 为 [1,2,3,4,5] 的一个旋转，该数组的最小值为1。


假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组[0,1,2,4,5,6,7] 可能变为[4,5,6,7,0,1,2])。

请找出其中最小的元素。

注意数组中可能存在重复的元素

*/

func findMin(nums []int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}
	if length == 1 {
		return nums[0]
	}

	for i := 0; i < length-1; i++ {
		if nums[i] <= nums[i+1] {
			continue
		}
		return nums[i+1]
	}
	return nums[0]
}

/**
利用二分查找
*/
func findMin2(nums []int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		if nums[left] <= nums[right] {
			return nums[left]
		}
		mid := left + (right-left)>>1
		if nums[left] <= nums[mid] { // [left,mid] 连续递增，则在 [mid+1,right] 查找
			left = mid + 1
			continue
		}
		right = mid // [left,mid] 不连续，在 [left,mid] 查找
	}
	return -1
}
