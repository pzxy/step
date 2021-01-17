package main

func main() {
	s := &S{}
	Aa(s)
}

type AAA interface {
	Aa()
}

type S struct {
	a int
}

func Aa(s *S) {

}
