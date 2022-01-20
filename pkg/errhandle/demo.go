package main

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
)
//错误包
//检查错误处理：https://github.com/kisielk/errcheck

func main() {
	err := demo1()
	if err != nil {
		fmt.Println(err)
	}
}

func demo1() error {
	_, err := os.Open("asdasd")
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
