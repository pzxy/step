package bytedance

import (
	"errors"
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
