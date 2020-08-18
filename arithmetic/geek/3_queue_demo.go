package geek

import (
	"fmt"
)

/**
队列我用这里实现了3种，数组队列，链表队列，
还有循环队列，循环队列会浪费一个空间，但循环队列能避免数据搬运
此外还有阻塞队列，满或空时会阻塞，这是生产者，消费者模型。
此外还有并发队列，多线程使用，需要加锁。
*/

/**
数组顺序队列
*/
type queue struct {
	data []interface{}
	//如果队列是切片的话，其实这个用处不大。数组的话，能派上用场
	head int
	tail int
}

func newQueue() *queue {
	return &queue{make([]interface{}, 0), 0, 0}
}

func (q *queue) enqueue(n interface{}) {
	q.data = append(q.data, n)
}

func (q *queue) dequeue() interface{} {
	if len(q.data) == 0 {
		return nil
	}
	v := q.data[0]
	q.data = append(q.data[1:])
	return v
}

func queueDemo() {
	q := newQueue()
	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
}

/**
链表队列
*/
type queue2 struct {
	head *node3
	tail *node3
	len  int
}
type node3 struct {
	v    interface{}
	next *node3
}

func newQueue2() *queue2 {
	return &queue2{}
}

func (q *queue2) enqueue2(data interface{}) {
	if q == nil {
		return
	}
	if data == nil {
		return
	}

	n := &node3{data, nil}
	if q.len == 0 {
		q.head = n
		q.tail = n
	} else {
		q.tail.next = n
		q.tail = n
	}
	q.len++
}

func (q *queue2) dequeue2() interface{} {
	if q.len == 0 {
		return nil
	}
	n := q.head.v
	q.head = q.head.next
	q.len--
	return n
}
func queueDemo2() {
	q := newQueue2()
	q.enqueue2(1)
	q.enqueue2(2)
	q.enqueue2(3)
	fmt.Println(q.dequeue2())
	fmt.Println(q.dequeue2())
	fmt.Println(q.dequeue2())
	q.enqueue2(4)
	fmt.Println(q.dequeue2())
	q.enqueue2(5)
	fmt.Println(q.dequeue2())
	q.enqueue2(6)
	fmt.Println(q.dequeue2())
}

/**
循环队列，和上面例1一样，只不过我们不删除队列中内容了。
关键算法 （tail+1)%cap == head 则满了,这是个很好的思想，我们应该注意他
*/

type queue3 struct {
	head  int
	tail  int
	cap   int
	items []interface{}
}

func newQueue3(cap int) *queue3 {
	return &queue3{0, 0, cap, make([]interface{}, cap)}
}

//当队列满时，图中的 tail 指向的位置实际上是没有存储数据的。所以，循环队列会浪费一个数组的存储空间。
func (q *queue3) enqueue3(data interface{}) {
	if (q.tail+1)%q.cap == q.head {
		fmt.Printf("队列满了，%v 不能存入 \n", data)
		return
	}
	q.items[q.tail] = data
	q.tail = (q.tail + 1) % q.cap
}

func (q *queue3) dequeue3() interface{} {
	if q.head == q.tail {
		fmt.Print("队列已空,返回nil ")
		return nil
	}

	ret := q.items[q.head]
	q.head = (q.head + 1) % q.cap
	return ret
}

func queueDemo3() {
	//这里没问题，n个空间只能存n-1个数
	q := newQueue3(6)
	q.enqueue3(1)
	q.enqueue3(2)
	q.enqueue3(3)
	q.enqueue3(4)
	q.enqueue3(5)
	q.enqueue3(6)
	fmt.Println(q.dequeue3())
	fmt.Println(q.dequeue3())
	fmt.Println(q.dequeue3())
	fmt.Println(q.dequeue3())
	fmt.Println(q.dequeue3())
	fmt.Println(q.dequeue3())
	fmt.Println(q.dequeue3())
	q.enqueue3(7)
	q.enqueue3(8)
	fmt.Println(q.dequeue3())
	fmt.Println(q.dequeue3())
}
