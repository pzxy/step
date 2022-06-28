package leetcode

import (
	"sync"
)

/**
设计和构建一个“最近最少使用”缓存，该缓存会删除最近最少使用的项目。
缓存应该从键映射到值(允许你插入和检索特定键对应的值)，并在初始化时指定最大容量。当缓存被填满时，它应该删除最近最少使用的项目。

它应该支持以下操作： 获取数据 get 和 写入数据 put 。

获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
写入数据 put(key, value) - 如果密钥不存在，则写入其数据值。
当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。

示例:

LRUCache cache = new LRUCache( 2 /* 缓存容量 */
/**
cache.put(1, 1)
cache.put(2, 2)
cache.get(1) // 返回  1




cache.put(3, 3) // 该操作会使得密钥 2 作废
cache.get(2) // 返回 -1 (未找到)
cache.put(4, 4) // 该操作会使得密钥 1 作废
cache.get(1)    // 返回 -1 (未找到)
cache.get(3) // 返回  3
cache.get(4) // 返回  4

*/
/**
自己实现1 有问题
*/
type LRUCache2 struct {
	cap   int
	len   int
	value sync.Map //key:interface value:interface()
	count sync.Map //key:interface value:int(count)
}

func NewLRUCache(cap int) *LRUCache2 {
	return &LRUCache2{cap: cap}
}

func (l *LRUCache2) Put(key interface{}, value interface{}) {
	//key exist
	_, ok := l.value.Load(key)
	if ok {
		l.value.Store(key, value)
		return
	}
	//cap  > len
	if l.cap > l.len {
		l.value.Store(key, value)
		l.count.Store(key, 0)
		l.len++
		return
	}
	//cap <= len
	var lessUseKey interface{}
	lessCount := 1<<63 - 1
	l.count.Range(func(key interface{}, value interface{}) bool {
		if value.(int) < lessCount {
			lessCount = value.(int)
			lessUseKey = key
		}
		return true
	})

	l.value.Delete(lessUseKey)
	l.value.Store(key, value)
	l.count.Delete(lessUseKey)
	l.count.Store(key, 0)
	return
}

func (l *LRUCache2) Get(key interface{}) interface{} {
	ret, ok := l.value.Load(key)
	if !ok {
		return -1
	}
	c, ok2 := l.count.Load(key)
	if ok2 {
		l.count.Store(key, c.(int)+1)
	}
	return ret
}

/**
官方实现
*/
type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode
}

type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

func initDLinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{
		key:   key,
		value: value,
	}
}

func Constructor4(capacity int) LRUCache {
	l := LRUCache{
		cache:    map[int]*DLinkedNode{}, //头节点和尾节点都是空的。
		head:     initDLinkedNode(0, 0),
		tail:     initDLinkedNode(0, 0),
		capacity: capacity,
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key]
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; !ok {
		node := initDLinkedNode(key, value)
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.capacity {
			removed := this.removeTail()
			delete(this.cache, removed.key)
			this.size--
		}
	} else {
		node := this.cache[key]
		node.value = value
		this.moveToHead(node)
	}
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

/**
自己模仿官方实现
*/
//type LRUCache4 struct {
//	len, cap   int
//	cache      map[int]*DListNode4
//	head, tail *DListNode4
//}
//
//type DListNode4 struct { //双向链表
//	key, value int
//	next, prev *DListNode4
//}
//
//func initDListNode4(key, value int) *DListNode4 {
//	return &DListNode4{
//		key:   key,
//		value: value,
//	}
//}
//
//func NewLRUCache4(cap int) *LRUCache4 {
//	lru := &LRUCache4{
//		len:   0,
//		cap:   cap,
//		cache: make(map[int]*DListNode4),
//		head:  initDListNode4(0, 0),
//		tail:  initDListNode4(0, 0),
//	}
//
//	lru.head.next = lru.tail
//	lru.tail.prev = lru.head
//	return lru
//}
//
//func (l *LRUCache4) Get(key int) int {
//	node, ok := l.cache[key]
//	if !ok {
//		return -1
//	}
//	l.moveToHead(node)
//	return node.value
//}
//
//func (l *LRUCache4) Put(key, value int) {
//	if v, ok := l.cache[key]; ok {
//		if v.value == value {
//			l.moveToHead(v)
//			return
//		}
//	}
//
//	if l.len < l.cap {
//		node := initDListNode4(key, value)
//		l.addToHead(node)
//		l.cache[key] = node
//		l.len++
//		return
//	} else {
//		removeNode := l.removeTail()
//		delete(l.cache, removeNode.key)
//		node := initDListNode4(key, value)
//		l.cache[key] = node
//		l.addToHead(node)
//	}
//
//}
//
//func (l *LRUCache4) addToHead(node *DListNode4) {
//	node.prev = l.head
//	node.next = l.head.next
//	l.head.next.prev = node
//	l.head.next = node
//}
//
//func (l *LRUCache4) moveToHead(node *DListNode4) {
//	l.removeNode(node)
//	l.addToHead(node)
//}
//
//func (l *LRUCache4) removeNode(node *DListNode4) {
//	node.prev.next = node.next
//	node.next.prev = node.prev
//}
//
//func (l *LRUCache4) removeTail() *DListNode4 {
//	node := l.tail.prev
//	l.removeNode(node)
//	return node
//}

//第二遍实现
type LRUCache4 struct {
	len, cap   int
	cache      map[int]*DListNode4
	head, tail *DListNode4
}
type DListNode4 struct {
	key, val   int
	prev, next *DListNode4
}

func initDListNode4(key, val int) *DListNode4 {
	return &DListNode4{
		key: key,
		val: val,
	}
}

func NewLRUCache4(cap int) *LRUCache4 {
	l := &LRUCache4{
		cap:   cap,
		cache: make(map[int]*DListNode4, 0),
		head:  initDListNode4(0, 0),
		tail:  initDListNode4(-1, -1),
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (l *LRUCache4) Get(key int) int {
	if v, ok := l.cache[key]; ok {
		l.moveToHead(v)
		return v.val
	}
	return -1
}

func (l *LRUCache4) Put(key, val int) {
	if v, ok := l.cache[key]; ok {
		if v.val != val {
			node := initDListNode4(key, val)
			l.cache[key] = node
			l.removeNode(v)
			l.addToHead(node)
		} else {
			l.moveToHead(v)
		}
		return
	}
	if l.len < l.cap {
		node := initDListNode4(key, val)
		l.cache[key] = node
		l.addToHead(node)
		l.len++
		return
	}
	last := l.removeTail()
	delete(l.cache, last.key)
	node := initDListNode4(key, val)
	l.cache[key] = node
	l.addToHead(node)
}

func (l *LRUCache4) moveToHead(node *DListNode4) {
	l.removeNode(node)
	l.addToHead(node)
}
func (l *LRUCache4) removeNode(node *DListNode4) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (l *LRUCache4) addToHead(node *DListNode4) {
	node.prev = l.head
	node.next = l.head.next
	l.head.next.prev = node
	l.head.next = node
}

func (l *LRUCache4) removeTail() *DListNode4 {
	node := l.tail.prev
	l.removeNode(node)
	return node
}
