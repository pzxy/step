package main

import "fmt"

func main() {
	/**
	goto
	*/
	i := 5
	if i == 8 {
		goto loop
	}
	fmt.Println("跳过此步")
loop:
	fmt.Println("调到这里")
	/**
	结合循环使用
	*/

}
