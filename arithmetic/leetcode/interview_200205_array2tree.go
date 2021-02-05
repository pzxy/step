package leetcode

//0, 3, 9, 10, -1, 5, -1
func Array2Tree(data1 []int) *TreeNode {
	return array2TreeDo(data1, 0)
}

//思路：和树转数组不同，数组不方便一开始确定大小，为了不越界，所以要append。所以层级遍历树。
//这里直接前序列还原
func array2TreeDo(data []int, rootIdx int) *TreeNode {
	if rootIdx >= len(data) || data[rootIdx] == -1 {
		return nil
	}
	root := &TreeNode{Val: data[rootIdx]}
	root.Left = array2TreeDo(data, rootIdx*2+1)
	root.Right = array2TreeDo(data, rootIdx*2+2)
	return root
}
