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
