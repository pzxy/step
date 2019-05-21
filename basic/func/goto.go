package main

import "fmt"

func main() {
	/**
	goto
	*/
	i := 5
	if i == 5 {
		goto loop
	}
	fmt.Println("跳过此步1")
loop:
	fmt.Println("调到这里")
	/**
	结合循环使用
	*/
	//2019-5-18
	j := 6

	if j == 6 {
		goto zxcv
	}
	fmt.Println("跳过此步")
zxcv:
	fmt.Println("来呀老弟")
	//除吃之外 还能这样写 continue loop  或者在循环的地方增加 break loop

}
