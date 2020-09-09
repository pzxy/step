package leetcode

import (
	"errors"
	"fmt"
	errors2 "github.com/pkg/errors"
	"step/grammar/codeskill/log"
	"strconv"
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
	l := Fibonacci4(n)
	fmt.Printf("%v  %v  %v %v", i, j, k, l)
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

	n99 := &TreeNode{Val: 99}
	n98 := &TreeNode{Val: 98}
	fmt.Println(lowestCommonAncestor3(n3, n98, n99))

	fmt.Println(lowestCommonAncestor2(n3, n5, n4))

}

func Test_TwoDimensionalArray(t *testing.T) {
	a := [][]int{
		{1, 2, 8, 9},
		{2, 4, 9, 12},
		{4, 7, 10, 13},
		{6, 8, 11, 15},
	}

	f := func(n int) bool {
		for _, v1 := range a {
			for _, v2 := range v1 {
				if v2 == n {
					return true
				}
			}
		}
		return false
	}

	for i := -100; i < 100; i++ {
		if f(i) == ifFindValFormTwoDimensionalArray(a, i) {
			fmt.Printf("%d is pass \n\t", i)
			continue
		}
		log.ErrLog(errors.New(fmt.Sprintf("vaule :%d error ", i)))

	}

}

func Test_ReplaceSpaces(t *testing.T) {
	s := []string{
		" we are happy ",
		"",
		" ",
		"we are happy",
	}
	s2 := []string{
		"%20we%20are%20happy%20",
		"",
		"%20",
		"we%20are%20happy",
	}

	for i, v := range s {

		if s2[i] == replaceSpaces(v) {
			continue
		}
		log.ErrLog(errors.New(fmt.Sprintf("i: %v ,expect: %v ,fact: %v \t",
			i, s2[i], replaceSpaces(v))))
	}

}

func Test_PrintlnReversSinglyListNode(t *testing.T) {
	n := &ListNode{Val: 1}
	n.newListNode()
	s := []*ListNode{
		nil,
		n,
		&ListNode{Val: 20},
	}
	for i, v := range s {
		fmt.Printf("****%v***\n", i)
		printReversSinglyListNode(v)
	}
}

func Test_BuildTree(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	n := buildTree(preorder, inorder)
	s := make([]int, len(preorder))
	idx := 0
	printPreorder(n, s, idx)

	for i := range preorder {
		if preorder[i] == s[i] {
			continue
		}
		log.ErrLog(errors.New("构建错误"))
	}
}

func Test_2StackImplementQueue(t *testing.T) {
	c := Constructor()
	length := 10
	for i := 0; i <= length; i++ {
		c.AppendTail(i)
	}
	for i := 0; i <= length; i++ {
		if c.DeleteHead() == i {
			continue
		}
		log.ErrLog(errors.New("lower"))
	}
}

func Test_findRepeatNumber1(t *testing.T) {
	s := []int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 11,
	}
	v := findRepeatNumber1(s)
	fmt.Println(v)
	if v == -1 {
		log.ErrLog(errors.New("未找到重复数字"))
	}

}

func Test_rotateArray(t *testing.T) {
	s := []int{
		2, 2, 2, 0, 1,
	}
	fmt.Println(findMin2(s))
}

func Test_matrixArray(t *testing.T) {

	s := [][]byte{
		[]byte("ABCE"),
		[]byte("SFES"),
		[]byte("ADEE"),
	}
	word := "ABCESEEEFS"

	//s4 := [][]byte{
	//	[]byte("AA"),
	//}
	//word := "AAA"

	//s5 := [][]byte{
	//	[]byte("ABCE"),
	//	[]byte("SFCS"),
	//	[]byte("ADEE"),
	//}
	//word := "ABCB"
	//fmt.Println(exist(s, word))
	fmt.Println(exist2(s, word))

}

func Test_robotMovingRange(t *testing.T) {
	m := 2
	n := 3
	k := 1
	count := movingCount2(m, n, k)
	if count != 3 {
		panic(fmt.Sprintf("count := %v", count))
	}
	m2 := 3
	n2 := 1
	k2 := 0
	count2 := movingCount2(m2, n2, k2)
	if count2 != 1 {
		panic(fmt.Sprintf("count2 := %v", count2))
	}

	m3 := 7
	n3 := 2
	k3 := 3
	count3 := movingCount2(m3, n3, k3)
	if count3 != 7 {
		panic(fmt.Sprintf("count3 := %v", count3))
	}
}

func Test_cuttingRope(t *testing.T) {
	//10 36
	if cuttingRope(10) != 36 {
		errors2.WithStack(errors.New("不对"))
	}
	fmt.Println()
}

func Test_1NumberInBinary(t *testing.T) {
	num := uint32(8)
	if hammingWeight(num) != 1 {
		log.ErrLog(errors.New("1不对"))
	}
	if hammingWeight2(num) != 1 {
		panic("2不对")
	}
	if hammingWeight3(num) != 1 {
		panic("3不对")
	}
	if hammingWeight4(num) != 1 {
		panic("4不对")
	}
	if hammingWeight22(num) != 1 {
		panic("22不对")
	}
}

func Test_Pow(t *testing.T) {
	x := 2.00000
	n := 10
	r := 1024.00000
	ret := myPow(x, n)
	ret, _ = strconv.ParseFloat(fmt.Sprintf("%.5f", ret), 5)

	if ret != r {
		fmt.Println(ret)
		panic("myPow1")
	}

	x = 2.10000
	n = 3
	r = 9.26100
	ret = myPow(x, n)
	ret, _ = strconv.ParseFloat(fmt.Sprintf("%.5f", ret), 5)

	if ret != r {
		fmt.Println(ret)
		panic("myPow2")
	}

	x = 2.00000
	n = -2
	r = 0.25
	ret = myPow(x, n)
	ret, _ = strconv.ParseFloat(fmt.Sprintf("%.5f", ret), 5)

	if ret != r {
		fmt.Println(ret)
		panic("myPow3")
	}

}

func Test_PrintNumber(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, v := range printNumbers(1) {
		if s[i] == v {
			fmt.Printf("%d \n", s[i])
			continue
		}
		log.ErrLog(errors.New("print number error"))
	}

	for i, v := range printNumbers2(1) {
		if s[i] == v {
			fmt.Printf("%d \n", s[i])
			continue
		}
		log.ErrLog(errors.New("print number error2"))
	}

}

func Test_deleteSingleNode(t *testing.T) {
	head := &ListNode{Val: 1}
	head.newListNode()
	deleteNode(head, 50)
	for head != nil {
		if head.Val != 50 {
			head = head.Next
			continue
		}
		log.ErrLog(errors.New("50 should already deleted"))
	}
}

func Test_deleteDuplication(t *testing.T) {
	head := &ListNode{Val: 1}
	head.newDuplicationListNode(3, 4)
	deleteDuplication2(&head)
	for head.Next != nil {
		if head != head.Next {
			fmt.Println(head.Val)
			head = head.Next
			continue
		}
		log.ErrLog(errors.New("cant have duplication value"))
	}
}

func Test_match(t *testing.T) {
	s := "mississippi"
	p := "mis*is*p*."
	if isMatch(s, p) {
		log.ErrLog(errors.New("can't match1"))
	}
	s = "aab"
	p = "c*a*b"
	if !isMatch(s, p) {
		log.ErrLog(errors.New("can't match2"))
	}
}

func Test_isNumber(t *testing.T) {
	s := []string{"1 ", "+100", "5e2", "-123", "3.1416", "-1E-16", "0123"}
	for _, v := range s {
		if !isNumber(v) {
			log.ErrLog(errors.New(fmt.Sprintf("should be number:(%v)", v)))
		}
	}

	s = []string{"12e", "1a3.14", "1.2.3", "+-5", "12e+5.4"}
	for _, v := range s {
		if isNumber(v) {
			log.ErrLog(errors.New(fmt.Sprintf("don't number:(%v)", v)))
		}
	}
}
