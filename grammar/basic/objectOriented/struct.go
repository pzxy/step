package main

import "fmt"

//结构体之间的继承
type People struct {
	Age  int
	Name string
}
type Teacher struct {
	Sex string
	People
}

func main() {
	t := Teacher{Sex: "男", People: People{Age: 12, Name: "王坤"}}
	fmt.Println(t)
}
