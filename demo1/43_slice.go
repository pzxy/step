package main

import "fmt"

func main() {
	order := make([]uint16, 2*5)

	pollorder := order[0:5:5]
	lockorder := order[5:][0:5:5]

	fmt.Println("len(pollorder) = ", len(pollorder))
	fmt.Println("cap(pollorder) = ", cap(pollorder))
	fmt.Println("len(lockorder) = ", len(lockorder))
	fmt.Println("cap(lockorder) = ", cap(lockorder))
}
