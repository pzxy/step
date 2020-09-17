package leetcode

/**
输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的循环双向链表。
要求不能创建任何新的节点，只能调整树中节点指针的指向。

为了让您更好地理解问题，以下面的二叉搜索树为例：

我们希望将这个二叉搜索树转化为双向循环链表。链表中的每个节点都有一个前驱和后继指针。
对于双向循环链表，第一个节点的前驱是最后一个节点，最后一个节点的后继是第一个节点。

特别地，我们希望可以就地完成转换操作。
当转化完成以后，树中节点的左指针需要指向前驱，
树中节点的右指针需要指向后继。还需要返回链表中的第一个节点的指针。

*/

func treeToDoublyList(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	helper(root)
	head, tail := root, root
	for head.Left != nil {
		head = head.Left
	}
	for tail.Right != nil {
		tail = tail.Right
	}
	head.Left = tail
	tail.Right = head
	return head
}

//这里如果要将pre作为值传进去，要使用pre的双重指针 **TreeNode这样的
func helper(root *TreeNode) {
	if root == nil {
		return
	}
	helper(root.Left)
	if pre != nil {
		root.Left = pre
		pre.Right = root
	}
	pre = root
	helper(root.Right)
}

var pre *TreeNode //必须在全局变量上才可以实现
