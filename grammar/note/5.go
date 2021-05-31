package main

import "fmt"

type People struct {
	Name string
}

//'v', 's', 'x', 'X', 'q': 这些类型都会循环打印
func (p People) String() string {

	return fmt.Sprintf("print111:%s", p)
}

func main() {
	p := &People{}
	p.Name = "123"
	fmt.Println(p.String())
}
