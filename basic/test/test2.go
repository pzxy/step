package main

import (
	"errors"
	"os"
)

/**
很好的错误处理的写法。
 */
func main() {
	A()
}

func a() int {
	return -1
}

func A() int {
	r := a()
	if r < 0 {
		a_err("print err info ")
	}
	return r
}

func a_err(s string) {
	errors.New(s)
	os.Exit(0) //平时的话，是不是要用return呢？
	//在平时写代码的过程中，是不是要加上什么也不做的处理呢。这样的话，有错的话，我就返回缺省值，
	// 缺省值就是什么也不做，这样在A()外面我也不用再次判断返回值。直接就可以运行了。相当于什么也不错。
	//我觉得这样写很好。
}
