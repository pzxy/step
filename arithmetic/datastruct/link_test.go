package datastruct

import "testing"

func TestLinkCircle(t *testing.T) {
	head := newLink(10)
	linkCircle2(head)
}
