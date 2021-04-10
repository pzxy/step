package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	b, err := json.Marshal("")
	fmt.Println(b, err)
}
