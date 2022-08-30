package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	pow("", "")
}

func pow(s string, target string) {
	s, target = "hello world", "000fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"
	src, nonce := s, 0
	for verify(sha(src), target) == false {
		nonce += 1
		src = block(s, nonce)
		fmt.Println(nonce)
	}
	fmt.Println(nonce, sha(src))
}

func sha(src string) string {
	s := sha256.Sum256([]byte(src))
	ss := sha256.Sum256(s[:])
	return hex.EncodeToString(ss[:])
}

func block(s string, nonce int) string {
	return s + strconv.Itoa(nonce)
}

func verify(src string, target string) bool {
	t1, _ := new(big.Int).SetString(src, 16)
	t2, _ := new(big.Int).SetString(target, 16)
	return t1.Cmp(t2) != 1
}
