package main

import (
	"fmt"
	"time"
)

func main() {
	s := new(student)
	s.name = "Ben"


	for i:=0;i<10 ;i++  {
		go func(s *student) {
			cl := new(class)
			cl.s = s
			fmt.Println(&s.name)
		}(s)
	}
	time.Sleep(time.Second*3)
}

type class struct {
	s *student
}
type student struct {
	name string
}
