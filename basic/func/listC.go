package main

import (
	"container/list"
	"fmt"
)

/**
双向链表
*/
func main() {
	//实例化list
	mylist := list.New()
	fmt.Println(mylist)
	//添加
	mylist.PushFront("a")
	mylist.PushBack("b")
	mylist.PushBack("c")
	mylist.PushBack("d")
	mylist.InsertAfter("end", mylist.Back())     //mylist 最后添加end
	mylist.InsertBefore("first", mylist.Front()) //mylist 前面添加first

	//遍历
	for e := mylist.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()
	//取指定位置元素值
	queryListByN(3, mylist)

	//移动元素
	mylist.MoveAfter(mylist.Front(), mylist.Back())  //把第一个参数移到第二个参数后面
	mylist.MoveBefore(mylist.Front(), mylist.Back()) //把第一个参数移到第二个参数前面

	mylist.MoveToFront(mylist.Back()) //吧最后一个参数移到最前面
	mylist.MoveToBack(mylist.Front()) //把第一个参数移到 最后面

	//删除元素
	mylist.Remove(mylist.Front()) //删除第一个元素

	//遍历
	for e := mylist.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()

}

//取指定位置值
func queryListByN(n int, mylist *list.List) {
	curr := mylist.Front()
	if n == 1 {
		curr = mylist.Front()
	} else if n == mylist.Len() {
		curr = mylist.Back()
	} else {
		for i := 1; i < n; i++ {
			curr = curr.Next()
		}
	}
	fmt.Printf("取出第%d个元素:%s\n", n, curr.Value)
}
