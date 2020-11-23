package main

import (
	"fmt"
)

func main() {
	orderLen := 5
	order := make([]uint16, 2*orderLen)
	fmt.Printf("order addr: %p \n", order)
	pollOrder := order[:orderLen:orderLen]
	lockOrder := order[orderLen:][:orderLen:orderLen]
	fmt.Println("len pollOrder", len(pollOrder))
	fmt.Println("cap pollOrder", cap(pollOrder))
	fmt.Println("len lockOrder", len(lockOrder))
	fmt.Println("cap lockOrder", cap(lockOrder))
	fmt.Printf("pollOrder addr: %p \n", pollOrder)
	fmt.Printf("lockOrder addr: %p \n", lockOrder)
}
