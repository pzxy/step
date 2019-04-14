package main

import (
	"fmt"
	"sort"
)

func main() {

	a := []int{1, 5, 4, 8, 4}
	sort.Ints(a) //升1
	b := []int{1, 5, 4, 8, 4}
	sort.Sort(sort.IntSlice(b)) //s升2
	fmt.Println(a)
	fmt.Println(b)
	//降
	sort.Sort(sort.Reverse(sort.IntSlice(a))) //sort.IntSlice(a)实现了interface接口
	fmt.Println(a)

}
