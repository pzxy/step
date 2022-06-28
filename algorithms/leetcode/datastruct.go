package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

type Heap struct {
	a     []int
	len   int
	count int
}
