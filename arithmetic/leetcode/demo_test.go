package leetcode

import (
	"fmt"
	"testing"
)

func Test_LengthOf(t *testing.T) {
	fmt.Println(lengthOf("aaaaadcaddd"))
}

func Test_Fibonacci(t *testing.T) {
	n := 7
	i := Fibo1(n)
	j := Fibo2(n)
	k := Fibo3(n)
	fmt.Printf("%v  %v  %v", i, j, k)
}

func Test_FindRepeatFrom100Int(t *testing.T) {
	sli := initSli(uint8(2))
	length := len(sli)
	n := len(sli) - 1

	if int(sli[n]) == length {
		fmt.Println(sli[n])
		return
	}

	for i, j := 0, n; i < j; i, j = i+1, j-1 {
		if sli[i]+sli[j] != 101 {
			if int(sli[i])+j == 101 {
				fmt.Println(sli[i])
			} else {
				fmt.Println(sli[j] + 1)
			}
			return
		}
	}
}

func Test_Prime(t *testing.T) {
	//test2()
	fmt.Print(allPrimeSum(1, 100000, 0))
}

func Test_RevSlice(t *testing.T) {
	//test2()
	revSlice()
}

func Test_RevSingleList(t *testing.T) {
	var node link
	node.value = 1
	node.next = &link{value: 2}
	node.next.next = &link{value: 3}
	node.next.next.next = &link{value: 4}
	node.next.next.next.next = &link{value: 5}
	node.next.next.next.next.next = &link{value: 6}
	node.next.next.next.next.next.next = &link{value: 7}

	//node.reverseLink().printLink()//全部反转
	//每三个反转
	node.revlink2()
	node.printLink()
}

func Test_StepGrid(t *testing.T) {
	m := 6
	n := 4
	fmt.Printf("%d * %d的方格一共 %d 种走法\n", m, n, walk(m, n))
	fmt.Printf("%d * %d的方格一共 %d 种走法\n", m, n, walks(m, n))

}

func Test_Int2Chinese(t *testing.T) {
	fmt.Println(int2Chinese(33402))
}

func Test_largestNumber(t *testing.T) {
	//s1 := []int{3, 30, 34, 5, 9}
	//s2 := []int{824, 938, 1399, 5607, 6973, 5703, 9609, 4398, 8247}
	s3 := []int{1, 2, 3, 1}
	fmt.Println(largestNumber(s3))
}

func Test_lowestCommonAncestor(t *testing.T) {
	//root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1 结果3
	// p = 5, q = 4,结果5
	/**
	  如果节点 X 存储在数组中下标为 i 的位置，下标为 2 * i 的位置存储的就是左子节点，
	  	下标为 2 * i + 1 的位置存储的就是右子节点。反过来，下标为 i/2 的位置存储就是它的父节点。
	  	通过这种方式，我们只要知道根节点存储的位置（一般情况下，为了方便计算子节点，根节点会存储在下标为 1 的位置），
	  	这样就可以通过下标计算，把整棵树都串起来。
	*/
	n6 := &TreeNode{Val: 6}
	n0 := &TreeNode{Val: 0}
	n8 := &TreeNode{Val: 8}
	n7 := &TreeNode{Val: 7}
	n4 := &TreeNode{Val: 4}

	n2 := &TreeNode{Val: 2, Left: n7, Right: n4}
	n5 := &TreeNode{Val: 5, Left: n6, Right: n2}
	n1 := &TreeNode{Val: 1, Left: n0, Right: n8}
	n3 := &TreeNode{Val: 3, Left: n5, Right: n1}

	fmt.Println(lowestCommonAncestor2(n3, n5, n1).Val)

	fmt.Println(lowestCommonAncestor2(n3, n5, n4).Val)

}
