package main

import "fmt"

//就是说这个v不在本goroutinue空间，所以他取值的时候要去主mian.go里面取这个位置的值，所以取值的时候，当前for到的v是哪个他取出来的就是哪个。
func main() {
	var a, b int
	b = incr(a)
	fmt.Println(a, b)
}
func incr(a int) (b int) {
	defer func() {
		a++
		b++
	}()
	a++
	return a
}
