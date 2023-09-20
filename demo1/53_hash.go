package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "operator_token=abc123&secret_key=a1b25cde5f3gh46ijkl&count=5000&bet_type=1&row_version=1346592723000"
	b := sha256.Sum256([]byte(s))
	fmt.Printf("%x", b)
}
