package main

import "fmt"

type Live2 interface {
	run()
}

type people2 struct{}
type animal2 struct{}

func (p people2) run() {
	fmt.Println("人跑不")
}

func (a animal2) run() {
	fmt.Println("DONGWUPAOBU")
}

func allrun(l Live2) {
	l.run()
}

func main() {
	p := people2{}
	a := animal2{}
	allrun(p)
	allrun(a)
}
