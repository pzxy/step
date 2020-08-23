package main

import "fmt"

func main() {
	/**
	copy切片
	*/
	s1 := []int{1, 2}
	s2 := []int{3, 4, 5, 6}
	copy(s2, s1)      //讲s1 copy到 s2中
	fmt.Println(s2)   //[1 2 5 6]
	copy(s1, s2[1:3]) //将s2中下表为1,2 copy到 s2中
	fmt.Println(s1)   //[2 5]
	/**
	利用copy切片 删除切片元素
	*/
	fmt.Println(delSliceInt(s2, 2))

}

func delSliceInt(s []int, n int) []int {
	newSlice := make([]int, n)
	copy(newSlice, s[:n])
	newSlice = append(newSlice, s[n+1:]...)
	return newSlice
}
