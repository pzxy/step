package bytedance

import (
	"errors"
	"fmt"
	"step/grammar/codeskill/log"
	"testing"
)

func TestDemo1P2(t *testing.T) {
	//1234
	//34
	//1268
	l1 := &ListNode{Val: 1}
	l1.Next = &ListNode{Val: 2}
	l1.Next.Next = &ListNode{Val: 3}
	l1.Next.Next.Next = &ListNode{Val: 4}

	l2 := &ListNode{Val: 3}
	l2.Next = &ListNode{Val: 4}
	target := []int{1, 2, 6, 8}
	l3 := demo1P2(l1, l2)

	for _, v := range target {
		if l3.Val != v {
			log.ErrLog(errors.New("TestDemo1P2 fail"))
			break
		}
		l3 = l3.Next
	}
}

func Test_minimumOperations(t *testing.T) {
	/**
	输入：leaves = "rrryyyrryyyrr"

	输出：2

	解释：调整两次，将中间的两片红叶替换成黄叶，得到 "rrryyyyyyyyrr"
	示例 2：

	输入：leaves = "ryr"
	*/
	strs := []string{
		"rrryyyrryyyrr", "ryr", "rrrrr", "yyyyy",
	}
	target := []int{2, 0, 1, 2}
	for i, s := range strs {
		ret := minimumOperations(s)
		if ret != target[i] {
			log.ErrLog(errors.New(fmt.Sprintf("minimumOperations fail:s:%s,target:%d,ret:%d \n\t", s, target[i], ret)))
		}
	}
}

func Test_scoreOfParentheses(t *testing.T) {
	/**
	输入： "()"
	输出： 1
	示例 2：

	输入： "(())"
	输出： 2
	示例 3：

	输入： "()()"
	输出： 2
	示例 4：

	输入： "(()(()))"
	输出： 6
	*/
	strs := []string{
		"()", "(())", "()()", "(()(()))",
	}
	target := []int{
		1, 2, 2, 6,
	}
	for i, s := range strs {
		ret := scoreOfParentheses(s)
		if ret != target[i] {
			log.ErrLog(errors.New(fmt.Sprintf("scoreOfParentheses fail:s:%s,target:%d,ret:%d \n\t", s, target[i], ret)))
		}
	}
}

func Test_countBits(t *testing.T) {
	/**
	  	输入: 2
	    输出: [0,1,1]
	    示例 2:

	    输入: 5
	    输出: [0,1,1,2,1,2]
	*/
	nums := []int{2, 5}
	target := [][]int{
		{0, 1, 1},
		{0, 1, 1, 2, 1, 2},
	}
	for i, n := range nums {
		ret := countBits(n)
		for j, v := range target[i] {
			if ret[j] == v {
				continue
			}
			log.ErrLog(errors.New(fmt.Sprintf("countBits fail:nums:%d,target:%v,ret:%v", n, target[i], ret)))
		}
	}

}

func Test_PreOrderOfNonRec(t *testing.T) {
	/**
	    3
	   / \
	  4   5
	 / \
	1   2
	*/
	r := &TreeNode{Val: 3}
	r.Left = &TreeNode{Val: 4}
	r.Right = &TreeNode{Val: 5}
	r.Left.Left = &TreeNode{Val: 1}
	r.Left.Right = &TreeNode{Val: 2}
	preOrderNonRec(r)
	fmt.Println("********")
	lastOrderNonRec(r)
	fmt.Println("********")

}

func TestNewMyMap(t *testing.T) {
	m := NewMyMap(10)
	m.Put("1", "1")
	m.Put("1", "2")
	m.Put("2", "2")
}

//259 -> 092500 00000

func TestMyHead_Insert(t *testing.T) {
	h := NewMyHead(10)
	h.insert(2)
	h.insert(5)
	h.insert(9)

	target := []int{0, 9, 2, 5, 0, 0, 0, 0, 0, 0, 0}
	for i, v := range h.a {
		if target[i] != v {
			log.ErrLog(errors.New("insert fail"))
		}
	}
}

func TestMerge4Array(t *testing.T) {
	array := []int{1, 5, 2, 6, 3, 9, 0}
	target := []int{0, 1, 2, 3, 5, 6, 9}
	Merge4Array(array, 0, len(array)-1)
	fmt.Println(array)
	for i, v := range array {
		if target[i] != v {
			log.ErrLog(errors.New("merge fail"))
		}
	}
}

func TestMerge4TwoArray(t *testing.T) {
	b := []int{1, 5, 2, 6, 3, 9, 0}
	a := []int{1, 5, 2, 6, 3, 9, 0}
	target := []int{0, 0, 1, 1, 2, 2, 3, 3, 5, 5, 6, 6, 9, 9}
	ret := Merge4TwoArray(a, b)
	fmt.Println(ret)
	for i, v := range ret {
		if target[i] != v {
			log.ErrLog(errors.New("merge2 fail"))
		}
	}
}

func TestMerge4ThreeArray(t *testing.T) {
	b := []int{1, 5, 2, 6, 3, 9, 0}
	a := []int{1, 5, 2, 6, 3, 9, 0}
	c := []int{1, 5, 2, 6, 3, 9, 0}
	target := []int{0, 0, 0, 1, 1, 1, 2, 2, 2, 3, 3, 3, 5, 5, 5, 6, 6, 6, 9, 9, 9}
	ret := Merge4ThreeArray(a, b, c)
	fmt.Println(ret)
	for i, v := range ret {
		if target[i] != v {
			log.ErrLog(errors.New("merge3 fail"))
		}
	}
}

func TestLongestSubstring(t *testing.T) {
	str := "adfacabcddfaa"
	target := "abcd"
	ret := LongestSubstring(str)
	fmt.Println("str:", str)
	fmt.Println("target:", target)
	fmt.Println("ret:", ret)
	for i, v := range []byte(ret) {
		if target[i] != v {
			log.ErrLog(errors.New("LongestSubstring fail"))
		}
	}
}

func TestStrConv(t *testing.T) {
	strConvDemo()
	pointDemo2()
}

func TestStockOnce(t *testing.T) {
	s := []int{100, 80, 120, 130, 70, 60, 100, 125}
	stockOnce(s)
}

func TestHeadInsert(t *testing.T) {
	fmt.Println("插入元素")

	h := NewMyHead(11)
	h.insert(1)
	h.insert(2)
	h.insert(3)
	h.insert(4)
	h.insert(5)
	h.insert(6)
	fmt.Println(h.a)
	fmt.Println()

}
func TestHeadBuild(t *testing.T) {
	fmt.Println("数组堆化")

	h := NewMyHead(10)
	h.a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	h.len = 10
	h.build()
	fmt.Println(h.a)
	fmt.Println()
}

func TestHeadDeleteTop(t *testing.T) {
	fmt.Println("删除堆顶元素")
	h := NewMyHead(10)
	h.a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	h.len = 10
	h.build()
	fmt.Println(h.a)
	fmt.Println()
	h.deleteTop()
	fmt.Println(h.a)
	fmt.Println()
	h.deleteTop()
	fmt.Println(h.a)
	fmt.Println()
}
func TestHeadSort(t *testing.T) {
	fmt.Println("堆排序顺序输出")
	h := NewMyHead(10)
	h.a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	h.len = 10
	h.build()
	fmt.Println(h.a)
	fmt.Println()
	h.sort()
	fmt.Println(h.a)
	fmt.Println()
}

func TestTreeCommonAncestor(t *testing.T) {
	n6 := &TreeNode{Val: 6}
	n0 := &TreeNode{Val: 0}
	n8 := &TreeNode{Val: 8}
	n7 := &TreeNode{Val: 7}
	n4 := &TreeNode{Val: 4}
	/**
	        3
	     5    1
	  6    2
	     7  4
	*/
	n2 := &TreeNode{Val: 2, Left: n7, Right: n4}
	n5 := &TreeNode{Val: 5, Left: n6, Right: n2}
	n1 := &TreeNode{Val: 1, Left: n0, Right: n8}
	n3 := &TreeNode{Val: 3, Left: n5, Right: n1}

	fmt.Println(findCommonNode(n3, n5, n1).Val)

}

func TestLoopQueue(t *testing.T) {
	q := NewQueue(10)
	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)
	q.enqueue(4)
	q.enqueue(5)
	q.enqueue(6)
	q.enqueue(7)
	q.enqueue(8)
	q.enqueue(9)
	q.enqueue(10)
	fmt.Println(q.a)
	q.dequeue()
	q.enqueue(10)
	fmt.Println(q.a)
	q.dequeue()
	q.enqueue(11)
	fmt.Println(q.a)
	q.dequeue()
	q.enqueue(12)
	fmt.Println(q.a)
}

func TestRemoveKdigits(t *testing.T) {
	fmt.Println(RemoveKdigits("1345213", 3))
}

func TestSortString(t *testing.T) {
	SortString("aabbdddd")
}
