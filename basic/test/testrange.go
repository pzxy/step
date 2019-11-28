package main

import "fmt"

type testArrS struct {
	testArrs []testArr
}
type testArr struct {
	Arr []string
	s   string
}

func main() {
	test1()
}
func test1() {
	t := new(testArr)
	t.Arr = append(t.Arr, "aa")
	t.Arr = append(t.Arr, "bb")
	t.Arr = append(t.Arr, "cc")
	tt := &testArrS{}

	tt.testArrs = append(tt.testArrs, *t)
	for _, order := range tt.testArrs {
		for _, d := range order.Arr {
			order.s = d
		}
	}
	fmt.Println(t)
}
func (t testArr) asdd(asd string) {

}
func test2() {
	t := new(testArr)
	t.Arr = append(t.Arr, "aa")
	t.Arr = append(t.Arr, "bb")
	t.Arr = append(t.Arr, "cc")
	tt := &testArrS{}
	tt.testArrs = append(tt.testArrs, *t)

	for _, order := range tt.testArrs {
		for i := 0; i < len(order.Arr); i++ {
			tt.testArrs[0].s = order.Arr[i]
		}
	}

	fmt.Println(t)
}
