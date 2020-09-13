package leetcode

/**
请完成一个函数，输入一个二叉树，该函数输出它的镜像。

例如输入：

  4
 /  \
 2   7
/ \  / \
1  3 6  9
镜像输出：

  4
 /  \
 7   2
/ \  / \
9  6 3 1



示例 1：

输入：root = [4,2,7,1,3,6,9]
输出：[4,7,2,9,6,3,1]

限制：

0 <= 节点个数 <= 1000

*/

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	//inOrderTraversal(root)
	root.Left, root.Right = root.Right, root.Left //交换不能放到中间，否则会交换了两遍同一个子节点
	mirrorTree(root.Left)
	mirrorTree(root.Right)
	return root
}

//inOrder postOrder preOrder traversal
func inOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	root.Left, root.Right = root.Right, root.Left //交换不能放到中间，否则会交换了两遍同一个子节点
	inOrderTraversal(root.Left)
	inOrderTraversal(root.Right)
}
