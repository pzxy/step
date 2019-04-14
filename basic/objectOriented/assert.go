package main

import "fmt"

/**
断言
*/
func main() {
	var i interface{} = 123
	i2, ok := i.(int) //断言
	if ok {
		fmt.Println(i2)
	}
}
