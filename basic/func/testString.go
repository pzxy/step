package main

import "fmt"

//将int定义为ChipType 类型
type ChipType int

const (
	None ChipType = iota
	CPU
	GPU
)

func (c ChipType) String() string {
	switch c {
	case None:
		return "None"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"
	}
	return "N/A"
}
func main() {
	fmt.Printf("%s %d", CPU, CPU)
}

// make a slice

var slice2 []int
var s = slice2 //make a slice
