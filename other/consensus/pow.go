package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {
	s, target := []byte("hello"), []byte("000")
	src, nonce := sha(s), 0
	for powVerify(sha(sha(src)), target) == false {
		nonce += 1
		src = sha(data(s, []byte(strconv.Itoa(nonce))))
		fmt.Println(nonce)
	}
	fmt.Println(nonce, hex.EncodeToString(src))
}

var s256 = sha256.New()

func sha(src []byte) []byte {
	s256.Reset()
	s256.Write(src)
	return s256.Sum(nil)
}

func data(s []byte, nonce []byte) []byte {
	return append(s, nonce...)
}

func powVerify(s []byte, target []byte) bool {
	if len(s) <= len(target) {
		return false
	}
	for i, v := range target {
		if s[i] != v {
			return false
		}
	}
	return true
}
