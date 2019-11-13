package main

import (
	"container/heap"
	"fmt"
)

type headStudent []*student

type student struct {
	name string
	age  int
}

func (hs headStudent) Less(i, j int) bool {
	return hs[i].age < hs[j].age
}

func (hs headStudent) Len() int {
	return len(hs)
}

func (hs headStudent) Swap(i, j int) {
	hs[i].age, hs[j].age = hs[j].age, hs[i].age
}

func (hs *headStudent) Pop() interface{} {
	old := *hs
	n := len(*hs)
	s := old[n-1]
	*hs = old[0 : n-1]
	return s
}

func (hs *headStudent) Push(x interface{}) {
	s := x.(*student)
	*hs = append(*hs, s)
}

func main() {
	hs := headStudent{}
	hs = append(hs, &student{name: "wangkun5555", age: 17})
	hs = append(hs, &student{name: "wangkun2", age: 18})
	hs = append(hs, &student{name: "wangkun3", age: 19})
	hs = append(hs, &student{name: "wangkun344", age: 19})
	hs = append(hs, &student{name: "wangkun344", age: 19})
	hs = append(hs, &student{name: "wangkun344", age: 20})
	hs = append(hs, &student{name: "wangkun344", age: 19})
	heap.Init(&hs)
	for len(hs) > 0 {
		stu := heap.Pop(&hs).(*student)
		fmt.Printf("%s:%d\n", stu.name, stu.age)
	}

}
