package bytedance

import (
	"fmt"
)

type queue struct {
	head, tail int
	cap        int
	a          []int
}

func NewQueue(cap int) *queue {
	return &queue{0, 0, cap, make([]int, cap)}

}

func (q *queue) enqueue(v int) {
	if (q.tail+1)%q.cap == q.head {
		fmt.Println("queue already full")
		return
	}
	q.a[q.tail] = v
	q.tail = (q.tail + 1) % q.cap
}

func (q *queue) dequeue() int {
	if q.head == q.tail {
		fmt.Println("queue is null")
		return -1
	}
	ret := q.a[q.head]
	q.head = (q.head + 1) % q.cap
	return ret
}
