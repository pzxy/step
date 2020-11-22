package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	//store
	x := uint32(0)
	atomic.StoreUint32(&x, 2)
	fmt.Println("store : x=", x)
	//add
	atomic.AddUint32(&x, ^uint32(0)) //减去1,补码，-1的补码是取反加1，刚好是...1111 1111，也就是0取反。
	fmt.Println("add - 1: x=", x)

	//load
	fmt.Println("load : x=", atomic.LoadUint32(&x))

	//CAS CompareAndSwap,x为1就替换x的值，并返回true，false
	fmt.Println("CompareAndSwap :", atomic.CompareAndSwapUint32(&x, uint32(1), uint32(5)))
	fmt.Println("CompareAndSwap : x=", x)

	//Swap

	fmt.Println("Swap 10 : old x=", atomic.SwapUint32(&x, 10)) //替换x值为10，并返回旧值
	fmt.Println("Swap : new x=", x)
}
