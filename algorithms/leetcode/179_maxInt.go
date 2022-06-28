package leetcode

import (
	"fmt"
	"strconv"
)

/**
给定一组非负整数，重新排列它们的顺序使之组成一个最大的整数。

示例 1:

输入: [10,2]
输出: 210
示例 2:

输入: [3,30,34,5,9]
输出: 9534330
说明: 输出结果可能非常大，所以你需要返回一个字符串而不是整数。
*/
func largestNumber(nums []int) string {
	s := ""

	for i := 0; i <= len(nums)-1; i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if strconv.Itoa(nums[j])+strconv.Itoa(nums[j+1]) > strconv.Itoa(nums[j+1])+strconv.Itoa(nums[j]) {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
		s += strconv.Itoa(nums[len(nums)-1-i])
	}
	if nums[len(nums)-1] == 0 {
		s = "0"
		return s
	}
	return s
}
func largestNumber3(nums []int) string {
	s := ""
	if len(nums) == 0 {
		return ""
	}
	for i := 0; i <= len(nums)-1; i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if strconv.Itoa(nums[j])+strconv.Itoa(nums[j+1]) < strconv.Itoa(nums[j+1])+strconv.Itoa(nums[j]) {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}

	}
	for _, v := range nums {
		s += strconv.Itoa(v)
	}
	if nums[0] == 0 {
		s = "0"
		return s
	}
	return s
}
func largestNumber2(nums []int) string {
	//
	s := nums
	var ret string
	fmt.Println(s)
	for i := 0; i <= len(s)-1; i++ {
		for j := 0; j < len(s)-1-i; j++ {
			if compareAGreater(s[j], s[j+1]) {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
		ret += strconv.Itoa(s[len(s)-1-i])
		if ret == "0" {
			return ret
		}

	}

	fmt.Println(s)
	return ret
}

//核心思想 35 3 -> 35 33   , 301 2 -> 301 222 然后比较大小,这样做不对
func compareAGreater(a, b int) bool {

	s1 := strconv.Itoa(a)
	s2 := strconv.Itoa(b)

	if len(s1) > len(s2) {
		v := int(s2[len(s2)-1])
		for i := 0; i < len(s1)-len(s2); i++ {
			vv, _ := strconv.Atoi(string(v))
			s2 += strconv.Itoa(vv)
		}
	} else {
		v := int(s1[len(s1)-1])
		for i := 0; i < len(s2)-len(s1); i++ {
			vv, _ := strconv.Atoi(string(v))
			s1 += strconv.Itoa(vv)
		}
	}

	if s1 >= s2 {
		return true
	}

	return false
}
