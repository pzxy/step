package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Setenv("ASD", "demo")
	fmt.Println(err)
	fmt.Println(os.Getenv("ASD"))
}
