package main

import (
	"fmt"
	"sort"
)

/**
map的创建 赋值 排序 创建  删除 清空map
*/
func main() {
	/*
		a := []int{1, 5, 4, 8, 4}
		sort.Ints(a) //升1
		b := []int{1, 5, 4, 8, 4}
		sort.Sort(sort.IntSlice(b)) //s升2
		fmt.Println(a)
		fmt.Println(b)
		//降
		sort.Sort(sort.Reverse(sort.IntSlice(a))) //sort.IntSlice(a)实现了interface接口
		fmt.Println(a)*/
	//	sortMap()
	testMap()
}

//将无序的map排序  排序后 每次输出的次序都是一样的

func sortMap() {
	//创建按
	map1 := make(map[int]string) //make分配空间
	map1[1] = "长出来"
	map1[2] = "afengbaobao"
	map1[3] = "网页"

	fmt.Println(map1)

	var sli []string //生命一个切片保存map中的值
	//从delete中删除map中的值
	delete(map1, 2)

	for _, val := range map1 {
		sli = append(sli, val)
	}
	//排序
	sort.Strings(sli) //对切片进行排序

	//清空map
	map1 = make(map[int]string) //创建一个空的map相当于清空了一个map
	fmt.Println(sli)
	fmt.Println("清空map1后", map1)
}
func testMap() {
	//创建按
	map1 := make(map[int]string) //make分配空间
	map1[1] = "长出来"
	map1[2] = "afengbaobao"
	map1[3] = "网页"
	sli := make([]int, 0)
	for k, _ := range map1 {
		sli = append(sli, k)
	}
	sort.Ints(sli)
	for _, v := range sli {
		fmt.Printf("key为:%d , value为:%s\n", v, map1[v])
	}
}
