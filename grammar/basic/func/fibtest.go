package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) { //Read接口比较底层
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p) //这里我们调用别的已经实现了的Read
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func fibanaci() intGen {
	var a, b = 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	f := fibanaci() //f 是 intGen类型的额 intGen实现了Read方法
	//printFileContents(f)
	var tmp int
	for {
		tmp = f()
		if tmp < 100000 {
			fmt.Println(tmp)
		} else {
			break
		}
	}

}
