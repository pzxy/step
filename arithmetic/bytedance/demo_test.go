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
	h.Insert(2)
	h.Insert(5)
	h.Insert(9)

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
