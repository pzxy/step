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
				//Val: inorder[root],
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
