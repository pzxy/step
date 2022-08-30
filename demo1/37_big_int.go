package main

import (
	"fmt"
	"math/big"
)

func main() {
	target := "0000ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"
	target2 := "abc"
	t, _ := new(big.Int).SetString(target, 16)
	t2, _ := new(big.Int).SetString(target2, 16)
	t3 := new(big.Int).Div(t, t2)
	fmt.Println(t3.Text(16))
}
