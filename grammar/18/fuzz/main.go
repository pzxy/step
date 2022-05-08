package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

//func main() {
//	input := "The quick brown fox jumped over the lazy dog"
//	rev := Reverse(input)
//	doubleRev := Reverse(rev)
//	fmt.Printf("original: %q\n", input)
//	fmt.Printf("reversed: %q\n", rev)
//	fmt.Printf("reversed again: %q\n", doubleRev)
//}

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev, revErr := Reverse(input)
	doubleRev, doubleRevErr := Reverse(rev)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q, err: %v\n", rev, revErr)
	fmt.Printf("reversed again: %q, err: %v\n", doubleRev, doubleRevErr)
}

//func Reverse(s string) string {
//	b := []byte(s)
//	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
//		b[i], b[j] = b[j], b[i]
//	}
//	return string(b)
//}

// Reverse1 使用 []rune(s) 而不是字节
func Reverse1(s string) string {
	fmt.Printf("input: %q\n", s)
	r := []rune(s)
	fmt.Printf("runes: %q\n", r)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// Reverse 必须是utf8编码
func Reverse(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid UTF-8")
	}
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r), nil
}
