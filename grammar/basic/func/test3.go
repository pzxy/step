package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type fibbRead func() int

var b int = 3

func fibb() fibbRead {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

//implement Read method

func (r fibbRead) Read(p []byte) (n int, err error) {
	s := r()
	if s > 10000 {
		return 0, io.EOF
	}
	next := fmt.Sprintf("%d\n", s)
	return strings.NewReader(next).Read(p)
}

//sacnner io.Reader

func printFile(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fibb()
	printFile(f)
}
