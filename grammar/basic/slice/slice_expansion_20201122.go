package main

import "fmt"

func main() {
	s1 := make([]int, 5, 10)
	s2 := s1[:5]
	s3 := s1[:5:5]
	fmt.Printf("s1 :len:%v,cap:%v,%p,%v \n", len(s1), cap(s1), s1, s1)
	fmt.Printf("s2 :len:%v,cap:%v,%p,%v \n", len(s2), cap(s2), s2, s2)
	fmt.Printf("s3 :len:%v,cap:%v,%p,%v \n", len(s3), len(s3), s3, s3)
}
