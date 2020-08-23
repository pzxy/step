package main

import (
	"fmt"
	"math"
)

func main() {
	float1()
}

func float1() {
	fmt.Printf("       %b \n", 12345)

	fmt.Printf("%b", math.Float32bits(12345))

}
