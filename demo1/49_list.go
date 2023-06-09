package main

import (
	"container/list"
	"fmt"
)

func main() {
	var l = list.List{}
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	l.PushBack(5)
	printList(&l)
	l.MoveToBack(l.Front())
	printList(&l)
	fmt.Println(l.Back().Value)
}

func printList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println("-----")
}
