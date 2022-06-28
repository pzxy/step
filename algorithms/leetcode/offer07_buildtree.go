package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。



例如，给出

前序遍历 preorder =[3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7


限制：

0 <= 节点个数 <= 5000
*/
//前序是为了找到根结点，中序是为了确定左右子树。
//但是，如果是[3,9,null,null,20,15,7]这种将空节点也存去的数组。那么可以直接构建，不需要要其他操作了。
func buildTree(preorder []int, inorder []int) *TreeNode {
	preorderLen := len(preorder)
	if preorderLen != len(inorder) {
		return nil
	}
	if preorderLen > 5000 { //0<= preorderLen <= 5000
		return nil
	}
	for k := range inorder {
		if inorder[k] == preorder[0] { //中序遍历 root (index=k)
			return &TreeNode{
				Val: preorder[0],
				//这里的preorder也使用k，k+1是因为，左子树节点个数是相同的。
				//也就是说。如果k是3，那么中序下标0，1，2，3，前面左子树有3个节点。
				//前序下标0，1，2，3，除了首位根结点下标0，后面还有三位，也就是说最后一个下标为3也就是k
				//中序截取的话，自然要街区k+1前面的数，不包括k+1
				Left:  buildTree(preorder[1:k+1], inorder[0:k]),
				Right: buildTree(preorder[k+1:], inorder[k+1:]),
			}
		}
	}
	return nil
}

func printPreorder(t *TreeNode, s []int, idx int) {
	if t == nil {
		return
	}
	//s = append(s, t.Val)这样出来结果是不对的，因为每次append后切片已经是另外一个切片了
	s[idx] = t.Val
	idx++
	printPreorder(t.Left, s, idx)
	idx++
	printPreorder(t.Right, s, idx)
}
