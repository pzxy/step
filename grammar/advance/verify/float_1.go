package main

import (
	"fmt"
	"math"
)

func main() {
	float1()
}

type User3 struct {
	user3Child *User3Child
	user5      *User5
}
type User3Child struct {
	user4 *User4
}
type User4 struct {
	u3 *User3Func
}
type User5 struct {
}
type User3Func func() int

func float1() {
	fmt.Printf("       %b \n", 12345)
	fmt.Printf("%b", math.Float32bits(12345))

}
