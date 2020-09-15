package leetcode

/**
从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。

例如:
给定二叉树:[3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回：

[3,9,20,15,7]


提示：

节点总数 <= 1000
*/

func levelOrder(root *TreeNode) []int {
	var ret []int
	if root == nil {
		return ret
	}
	queue := make([]*TreeNode, 1)
	queue[0] = root

	for len(queue) != 0 {
		r := queue[0]
		queue = queue[1:]
		ret = append(ret, r.Val)
		if r.Left != nil {
			queue = append(queue, r.Left)
		}
		if r.Right != nil {
			queue = append(queue, r.Right)
		}
	}

	return ret
}
