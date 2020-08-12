package main

import "fmt"

func main() {
	n := &node4{v: 5}
	n.insert(1)
	n.insert(2)
	n.insert(3)
	n.insert(4)
	n.insert(6)
	n.insert(7)
	n.insert(8)
	n.insert(9)
	n.inOrderTraversal()
	fmt.Println(n.find(10))
}

type node4 struct {
	v     int
	left  *node4
	right *node4
}

func (t *node4) find(data int) bool {
	for t != nil {
		if t.v == data {
			return true
		}
		if t.v > data {
			if t.left == nil {
				return false
			}
			t = t.left
			continue
		}
		if t.v < data {
			if t.right == nil {
				return false
			}
			t = t.right
			continue
		}
	}
	return false
}

func (t *node4) insert(data int) {
	//如果存在如何插入?
	for t != nil {
		if t.v > data {
			if t.left == nil {
				t.left = &node4{v: data}
				return
			}
			t = t.left
			continue
		}
		if t.v < data {
			if t.right == nil {
				t.right = &node4{v: data}
				return
			}
			t = t.right
			continue
		}
	}
}
func (t *node4) delete(data int) bool {

	return false
}

func (t *node4) inOrderTraversal() {
	if t == nil {
		return
	}
	t.left.inOrderTraversal()
	fmt.Println(t.v)
	t.right.inOrderTraversal()
}
