package main

import "fmt"

func main() {
	const sliceSize = 1000
	slice1 := make([]int, sliceSize)

	for i := 0; i < sliceSize; i++ {
		slice1[i] = i
	}
	aliseSlice := slice1 //起了一个别名
	aliseSlice[1] = 999
	copySlice := make([]int, sliceSize)
	copy(copySlice, slice1)
	fmt.Println(copySlice[1:2])
	//删除切片
	slice1 = append(slice1[:1], slice1[2:]...) //...表示将整个slice添加到slice1[:1]后面
	//
	fmt.Println(slice1[:2])
	for _, val := range slice1 {
		fmt.Printf("%d ", val)
	}

	/*
		s := []int{1, 2, 3}
		f(s)
		fmt.Println(s)*/
}
func f(s []int) {
	for index := range s {
		s[index] = 3
	}
	//val是一个副本,不能改变s中的元素
	for _, val := range s {
		val++
	}
}
