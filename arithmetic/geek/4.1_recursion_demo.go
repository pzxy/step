package geek

import (
	"fmt"
)

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
go不支持尾递归，但是c支持。只不过我用线形递归的方法写或者用尾递归的方法写，
如果编译时不指定优化的话，都是报段错误。
但是我开启优化后，无论是线性递归还是尾递归都会被完全优化的。树形递归我没试过。
树形递归计算量非常大，有很多重复计算，可以去掉这些重复计算，或者直接改成线性递归。
*/

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
1.1 树形递归改造
这个办法真的太nice了，非常棒
*/

func walkTreeDemo2(n int) {
	m := make(map[int]int, n)
	w := walk{m}
	fmt.Println(w.walkTree2(n))
}

type walk struct {
	mMap map[int]int
}

func (w *walk) walkTree2(n int) int {
	if v, ok := w.mMap[n]; ok {
		return v
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	ret := w.walkTree2(n-1) + w.walkTree2(n-2)
	w.mMap[n] = ret
	return ret
}

/**
2. 线性递归
*/
func walkLineDemo(n int) {
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

/**
4.递归改循环
*/
func walkCirculateDemo(n int) {
	fmt.Println(walkCirculate(n))

}

func walkCirculate(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	ret := 0
	pre := 2
	prepre := 1

	for i := 3; n >= i; i++ {
		ret = pre + prepre
		prepre = pre
		pre = ret
	}

	return ret
}

/**
5. 尾调用，其实就是最后return时是一个调用函数，这个函数不必是调用者自己。
尾递归就调用者自己。
*/
