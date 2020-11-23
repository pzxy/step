package main

import "fmt"

/**
重点 copy的时候 切片不会扩容。
*/
func main() {
	copyDemo()

}

func copyDemo() {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	s2 := make([]int, 5)
	var s3 []int

	copy(s2, s1)
	copy(s3, s1)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Println("s3:", s3)
}
