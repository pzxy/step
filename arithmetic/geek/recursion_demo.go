package main

import "fmt"

/**
递归的要点：
1. 一个问题的解可以分解为几个子问题的解
2. 这个问题与分解之后的子问题，除了数据规模不同，求解思路完全一样
3. 存在递归终止条件

递归的关键：
求出递归公式

递归防止栈溢出：
设置递归深度，比如超过1000就返回报错

递归出问题后调试：
1.打印日志发现，递归值。
2.结合条件断点进行调试

另外王争没有讲尾递归，这个可以有效的防止栈溢出，但是不是所有编译器都支持。
*/

func main() {
	//walkTreeDemo(20)
	//2000 0000 stack overflow
	//walkLineDemo(20000000)
	walkTailDemo(20000000)
}

/**
爬楼梯问题，n个台阶，每次走1个或者2个，问多少种走法。
可以根据第一步的走法把所有走法分为两类，第一类是第一步走了 1 个台阶，
另一类是第一步走了 2 个台阶。所以 n 个台阶的走法就等于先走 1 阶后，
n-1 个台阶的走法 加上先走 2 阶后，n-2 个台阶的走法。
f(1) = 1
f(2) = 2
f(n) = f(n-1) + f(n-2)
*/

/**
1. 树形递归
*/
func walkTreeDemo(n int) {
	fmt.Println(walkTree(n))
}

func walkTree(n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}
	return walkTree(n-1) + walkTree(n-2)
}

/**
2. 线性递归
*/
func walkLinkDemo(n int) {
	fmt.Println(walkLine(1, 1, n))
}

func walkLine(a int, b int, n int) int {
	if n == 1 {
		return 1
	}
	//
	//if n == 2 {
	//	return 2
	//}
	return a + walkLine(b, a+b, n-1)
}

/**
3. 尾递归
*/
func walkTailDemo(n int) {
	fmt.Println(walkTail(1, 1, n))
}
func walkTail(a int, b int, n int) int {
	if n == 0 {
		return a
	}
	//if n == 2 {
	//	return v + 2
	//}
	return walkTail(b, a+b, n-1)
}
