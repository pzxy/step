package datastruct

import (
	"fmt"
	"testing"
)

func Test_Link(t *testing.T) {
	printLink(removeNthFromEnd(newLink(10), 9))
	//linkRev()
	//printLink(orderLinkMerge(newLink(10), newLink(10)))
	//linkDeleteN()
	//fmt.Println(linkMinNode(newLink(10)).v)
	//fmt.Println(linkCircle(newLinkCircle(100)))

	//printLink(recLinkNodeInterval3(newLink(11)))
}

func Test_Stack(t *testing.T) {
	//stackDemo()
	stack2Demo()
}

func Test_Queue(t *testing.T) {
	//queueDemo()
	//queueDemo2()
	queueDemo3()
}

func Test_Recursion(t *testing.T) {
	walkTreeDemo(10)
	//walkTreeDemo2(1000)
	//2000 0000 stack overflow
	//walkLineDemo(1000)
	//walkTailDemo(20)
	walkCirculateDemo(10)
	//fmt.Println(r2(20, 0))

}

func Test_Bubble_Insert_Select_Sort(t *testing.T) {
	s := []int{2, 1, 3, 5, 2, 9, 4, 7, 6, 8, 0}
	//bubbleDemo(s)
	//insertDemo(s)
	//insertDemo2(s)
	selectDemo(s)

}

func Test_Merge_Quick_Sort(t *testing.T) {
	s := []int{2, 1, 3, 5, 2, 9, 4, 7, 6, 8, 0}
	mergeDemo(s)
	quickDemo(s)
}

func Test_Bucket_Counting_Radix(t *testing.T) {
	countingDemo()
}

func Test_BinarySearch1(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	binarySearchDemo(a)
	recBinarySearchDemo(a)
	squareRootDemo(10, 3)
}

func Test_BinarySearch2(t *testing.T) {
	a := []int{1, 2, 3, 5, 5, 6, 7, 8, 9, 10}
	searchFirstVDemo(a, 5)
	searchLastVDemo(a, 5)
	searchFirstGreaterVDemo(a, 5)
	searchLastLessVDemo(a, 4)
}

func Test_Tree(t *testing.T) {
	n := &node4{v: 5}
	n.insert(1)
	n.insert(2)
	n.insert(3)
	n.insert(4)
	n.insert(8)
	n.insert(7)
	n.insert(10)
	n.insert(9)
	//1234 5 78910
	fmt.Println(n.delete(5))

	n.inOrderTraversal()
	fmt.Println(n.find(10))
}

func Test_Head(t *testing.T) {
	headInsert()
	//headBuild()
	//headDeleteTop()
	//headSort()
}

func Test_StringMatch(t *testing.T) {
	bfDemo()
	rkDemo()
}
