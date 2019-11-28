package main

import (
	"bytes"
	"fmt"
)

func main() {
	trimDemo()
}

//35 func Trim（查看源代码）
//func Trim(s []byte, cutset string) []byte
//Trim 通过切割 cutset 中包含的所有前导和尾随 UTF-8 编码的 Unicode 代码点来返回 s 的子片段。

func trimDemo() {
	fmt.Printf("[%q]", bytes.Trim([]byte(" !!! Achtung! Achtung! !!! "), "! "))
	//["Achtung! Achtung"]
}

//36 func TrimFunc（查看源代码）
//func TrimFunc(s []byte, f func(r rune) bool) []byte
//TrimFunc 通过分割所有满足f(c)的前导和尾随 UTF-8 编码的 Unicode 代码点 c 来返回 s 的子片段
