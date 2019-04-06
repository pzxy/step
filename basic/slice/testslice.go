package slice

import "fmt"

func main() {
	s := []int{1, 2, 3}
	f(s)
	fmt.Println(s)
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
