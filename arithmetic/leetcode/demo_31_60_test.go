package leetcode

import (
	"errors"
	"fmt"
	"step/grammar/codeskill/log"
	"testing"
)

func Test_validateStackSequences(t *testing.T) {
	pushed := []int{1, 2, 3, 4, 5}
	popped := []int{4, 5, 3, 2, 1}
	if !validateStackSequences(pushed, popped) {
		log.ErrLog(errors.New("validate fail"))
	}

	pushed = []int{1, 2, 3, 4, 5}
	popped = []int{4, 5, 3, 2, 2}
	if validateStackSequences(pushed, popped) {
		log.ErrLog(errors.New("validate fail"))
	}
}

func Test_levelOrder(t *testing.T) {
	/**
	    3
	   / \
	  9  20
	/  \  /  \
	1   2  15   7
	返回：
	[3,9,20,1,2,15,7]
	*/
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}
	target := []int{3, 9, 20, 1, 2, 15, 7}
	ret := levelOrder(root)
	if ret == nil {
		log.ErrLog(errors.New("offer32 level order fail"))
	}
	for i, v := range target {
		if ret[i] != v {
			log.ErrLog(errors.New("offer32 level order fail"))
		}
	}
}

func Test_levelOrder2(t *testing.T) {
	/**
	    3
	   / \
	  9  20
	/  \  /  \
	1   2  15   7
	返回：
	[[3],[9,20],[1,2,15,7]]
	*/
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}
	target := [][]int{
		{3},
		{9, 20},
		{1, 2, 15, 7},
	}
	ret := levelOrder2(root)
	if ret == nil {
		log.ErrLog(errors.New("offer32 level order fail"))
	}

	for i, v := range target {
		for j, vv := range v {
			if ret[i][j] != vv {
				log.ErrLog(errors.New("offer32 level order fail"))
			}
		}
	}
}

func Test_levelOrder3(t *testing.T) {
	/**
		    3
		   / \
		  9  20
	         /  \
		    15   7
		返回：
		[[3],[20,9],[1,2,15,7]]
	*/
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}
	target := [][]int{
		{3},
		{20, 9},
		{15, 7},
	}
	ret := levelOrder3(root)
	if ret == nil {
		log.ErrLog(errors.New("offer32 level order fail"))
	}

	for i, v := range target {
		for j, vv := range v {
			if ret[i][j] != vv {
				log.ErrLog(errors.New("offer32 level order fail"))
			}
		}
	}
}

func Test_levelOrder3ByStack(t *testing.T) {
	/**
	    3
	   / \
	  9  20
	/  \  /  \
	1   2  15   7
	返回：
	[[3],[20,9],[1,2,15,7]]
	*/
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}
	target := [][]int{
		{3},
		{20, 9},
		{1, 2, 15, 7},
	}
	ret := levelOrder3ByStack(root)
	if ret == nil {
		log.ErrLog(errors.New("offer32 level order fail"))
	}

	for i, v := range target {
		for j, vv := range v {
			if ret[i][j] != vv {
				log.ErrLog(errors.New("offer32 level order fail"))
			}
		}
	}
}

func Test_verityPostOrder(t *testing.T) {
	/**
	     5
	    / \
	   2   6
	  / \
	 1   3
	示例 1：

	输入: [1,6,3,2,5]
	输出: false
	示例 2：

	输入: [1,3,2,6,5]
	输出: true
	*/
	verity1 := []int{1, 6, 3, 2, 5}
	verity2 := []int{1, 3, 2, 6, 5}
	verity3 := []int{4, 6, 5, 7}
	if verifyPostorder(verity1) {
		log.ErrLog(errors.New("verity1 post order fail"))
	}
	if !verifyPostorder(verity2) {
		log.ErrLog(errors.New("verity2 post order fail"))
	}
	if !verifyPostorder(verity3) {
		log.ErrLog(errors.New("verity3 post order fail"))
	}
}

func Test_pathSum(t *testing.T) {
	/**
	              5
	             / \
	            4   8
	           /   / \
	          11  13  4
	         /  \    / \
	        7    2  5   1
	返回:

	[
	   [5,4,11,2],
	   [5,8,4,5]
	]
	*/
	r := &TreeNode{Val: 5}
	r.Left = &TreeNode{Val: 4}
	r.Left.Left = &TreeNode{Val: 11}
	r.Left.Left.Left = &TreeNode{Val: 7}
	r.Left.Left.Right = &TreeNode{Val: 2}
	r.Right = &TreeNode{Val: 8}
	r.Right.Left = &TreeNode{Val: 13}
	r.Right.Right = &TreeNode{Val: 4}
	r.Right.Right.Left = &TreeNode{Val: 5}
	r.Right.Right.Right = &TreeNode{Val: 1}
	target := [][]int{
		{5, 4, 11, 2},
		{5, 8, 4, 5},
	}

	ret := pathSum(r, 22)
	for i, v := range target {
		for j, vv := range v {
			if vv != ret[i][j] {
				log.ErrLog(errors.New("Test_pathSum is fail"))
			}
		}
	}
}

func Test_CopyRandomList(t *testing.T) {
	/**
	输入：head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
	输出：[[7,null],[13,0],[11,4],[10,2],[1,0]]
	*/
	n0 := &Node{Val: 7}
	n1 := &Node{Val: 13}
	n2 := &Node{Val: 11}
	n3 := &Node{Val: 10}
	n4 := &Node{Val: 1}

	n0.Next = n1
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4

	n0.Random = nil
	n1.Random = n0
	n2.Random = n4
	n3.Random = n2
	n4.Random = n0

	ret := copyRandomList(n0)
	if ret == nil {
		log.ErrLog(errors.New("copyRandomList fail"))
	}

	for ret != nil {
		if ret.Val != n0.Val {
			log.ErrLog(errors.New("copyRandomList fail"))
		}
		if n0.Random == nil && ret.Random == nil {
			ret = ret.Next
			n0 = n0.Next
			continue
		}
		if n0.Random.Val != ret.Random.Val {
			log.ErrLog(errors.New("copyRandomList fail"))
		}
		ret = ret.Next
		n0 = n0.Next
	}

}

func Test_TreeToDoublyList(t *testing.T) {
	root := &TreeNode{4, nil, nil}
	node1 := &TreeNode{2, nil, nil}
	node2 := &TreeNode{5, nil, nil}
	node3 := &TreeNode{1, nil, nil}
	node4 := &TreeNode{3, nil, nil}
	root.Left = node1
	root.Right = node2
	node1.Left = node3
	node1.Right = node4
	head := treeToDoublyList(root)
	tail := head.Left
	//从头开始遍历
	for i := 0; i <= 9; i++ {
		fmt.Printf("%d\t", head.Val)
		head = head.Right
	}
	fmt.Println()
	//从尾开始遍历
	for i := 0; i <= 9; i++ {
		fmt.Printf("%d\t", tail.Val)
		tail = tail.Left
	}
}

func Test_serializeTree4Recursion(t *testing.T) {
	/**
	    1
	   / \
	  2   3
	     / \
	    4   5
	序列化为 "[1,2,3,null,null,4,5]"
	*/
	r := &TreeNode{Val: 1}
	r.Left = &TreeNode{Val: 2}
	r.Right = &TreeNode{Val: 3}
	r.Right.Left = &TreeNode{Val: 4}
	r.Right.Right = &TreeNode{Val: 5}
	target := "1,2,#,#,3,4,#,#,5,#,#,"

	c := Constructor37()
	str := c.serialize(r)
	root := c.deserialize(str)

	if str != target {
		log.ErrLog(errors.New("serialize fail of serializeTree"))
	}
	var f func(t1 *TreeNode, t2 *TreeNode) bool

	f = func(t1 *TreeNode, t2 *TreeNode) bool {
		if t1 == nil && t2 == nil {
			return true
		}
		if t1 == nil || t2 == nil {
			return false
		}
		if t1.Val != t2.Val {
			return false
		}
		return f(t1.Left, t2.Left) && f(t1.Right, t2.Right)
	}
	if !f(r, root) {
		log.ErrLog(errors.New("deserialize fail of serializeTree"))
	}
}

func Test_serializeTree4BFS(t *testing.T) {
	/**
	    1
	   / \
	  2   3
	     / \
	    4   5
	序列化为 "[1,2,3,null,null,4,5]"
	*/
	r := &TreeNode{Val: 1}
	r.Left = &TreeNode{Val: 2}
	r.Right = &TreeNode{Val: 3}
	r.Right.Left = &TreeNode{Val: 4}
	r.Right.Right = &TreeNode{Val: 5}
	target := "1,2,#,#,3,4,#,#,5,#,#,"

	c := Constructor37()
	str := c.serialize1(r)
	root := c.deserialize1(str)

	if str != target {
		log.ErrLog(errors.New("serialize1 fail of serializeTree1"))
	}
	var f func(t1 *TreeNode, t2 *TreeNode) bool

	f = func(t1 *TreeNode, t2 *TreeNode) bool {
		if t1 == nil && t2 == nil {
			return true
		}
		if t1 == nil || t2 == nil {
			return false
		}
		if t1.Val != t2.Val {
			return false
		}
		return f(t1.Left, t2.Left) && f(t1.Right, t2.Right)
	}
	if !f(r, root) {
		log.ErrLog(errors.New("deserialize1 fail of serializeTree1"))
	}
}

func Test_permutation(t *testing.T) {
	/**
	输入：s = "abc"
	输出：["abc","acb","bac","bca","cab","cba"]
	*/
	s := "abc"
	target := []string{"abc", "acb", "bac", "bca", "cab", "cba"}
	ret := permutation(s)
	if len(target) != len(ret) {
		log.ErrLog(errors.New("permutation1 fail"))
		return
	}

	for len(target) > 0 {
		targetVal := target[0]
		target = target[1:]
		var isFind bool
		for _, v := range ret {
			if targetVal == v {
				isFind = true
				break
			}
		}
		if isFind {
			continue
		}
		log.ErrLog(errors.New("permutation1 fail"))
		return
	}

}

func Test_majorityElement(t *testing.T) {
	/**
	输入: [1, 2, 3, 2, 2, 2, 5, 4, 2]
	输出: 2
	*/
	//s := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
	s := []int{2, 2}
	target := 2
	ret := majorityElement(s)
	if ret != target {
		log.ErrLog(errors.New("majority element fail"))
	}

	s = []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
	target = 2
	ret = majorityElement2(s)
	if ret != target {
		log.ErrLog(errors.New("majority element fail"))
	}

}

func Test_getLeastNumbers(t *testing.T) {
	/**
	输入：arr = [3,2,1], k = 2
	输出：[1,2] 或者 [2,1]

	*/
	//[0,0,1,2,4,2,2,3,1,4]
	//8
	arr := []int{3, 2, 1}
	k := 2
	target := []int{1, 2}
	ret := getLeastNumbers(arr, k)
	if len(target) != len(ret) {
		log.ErrLog(errors.New("getLeastNumbers1 fail"))
	}
	for _, v := range target {
		var isFind bool
		for _, vv := range ret {
			if v == vv {
				isFind = true
			}
		}
		if !isFind {
			log.ErrLog(errors.New("getLeastNumbers1 fail"))
			break
		}
	}

}

func Test_getLeastNumbers2(t *testing.T) {
	/**
	输入：arr = [3,2,1], k = 2
	输出：[1,2] 或者 [2,1]

	*/
	//[0,0,1,2,4,2,2,3,1,4]
	//8
	arr := []int{0, 0, 1, 2, 4, 2, 2, 3, 1, 4}
	k := 8
	target := []int{0, 0, 1, 2, 2, 2, 3, 1}
	ret := getLeastNumbers2(arr, k)
	if len(target) != len(ret) {
		log.ErrLog(errors.New("getLeastNumbers2 fail"))
	}
	for _, v := range target {
		var isFind bool
		for _, vv := range ret {
			if v == vv {
				isFind = true
			}
		}
		if !isFind {
			log.ErrLog(errors.New("getLeastNumbers2 fail"))
			break
		}
	}

}

func Test_MedianFinder(t *testing.T) {
	/**
	输入：
	["MedianFinder","addNum","addNum","findMedian","addNum","findMedian"]
	[[],[1],[2],[],[3],[]]
	输出：[null,null,null,1.50000,null,2.00000]
	示例 2：
	*/
	medianFinder := Constructor41()
	medianFinder.AddNum(1)
	medianFinder.AddNum(2)
	ret := medianFinder.FindMedian()
	if ret != 1.50000 {
		log.ErrLog(errors.New("MedianFinder1 fail"))
		return
	}
	medianFinder.AddNum(3)
	ret = medianFinder.FindMedian()
	if ret != 2.00000 {
		log.ErrLog(errors.New("MedianFinder2 fail"))
		return
	}

}

func Test_MaxSubArray(t *testing.T) {
	/**
	输入: nums = [-2,1,-3,4,-1,2,1,-5,4]
	输出: 6
	解释:连续子数组[4,-1,2,1] 的和最大，为6。
	*/
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	target := 6 //[]int{4, -1, 2, 1}
	ret := maxSubArray(nums)
	if ret != target {
		log.ErrLog(errors.New("maxSubArray fail"))
	}
}

func Test_MaxSubArray4dp(t *testing.T) {
	/**
	输入: nums = [-2,1,-3,4,-1,2,1,-5,4]
	输出: 6
	解释:连续子数组[4,-1,2,1] 的和最大，为6。
	*/
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	target := 6 //[]int{4, -1, 2, 1}
	ret := maxSubArray4dp(nums)
	if ret != target {
		log.ErrLog(errors.New("maxSubArray fail"))
	}
}

func Test_CountDigitOne(t *testing.T) {
	/**
	输入：n = 12
	输出：5
	输入：n = 13
	输出：6
	*/
	n := []int{12, 13}
	target := []int{5, 6}

	for i, v := range n {
		ret := countDigitOne(v)
		if ret != target[i] {
			log.ErrLog(errors.New(fmt.Sprintf("countDigitOne fail,n:%d,target:%d,ret:%d", v, target[i], ret)))
		}
	}

}

func Test_CountDigitOne2(t *testing.T) {
	/**
	输入：n = 12
	输出：5
	输入：n = 13
	输出：6
	*/
	n := []int{12, 13}
	target := []int{5, 6}

	for i, v := range n {
		ret := countDigitOne2(v)
		if ret != target[i] {
			log.ErrLog(errors.New(fmt.Sprintf("countDigitOne fail,n:%d,target:%d,ret:%d", v, target[i], ret)))
		}
	}

}

func Test_FindNthDigit(t *testing.T) {
	//n := 123456789101112131415
	/**
	输入：n = 3
	输出：3
	示例 2：

	输入：n = 11
	输出：0
	*/
	nums := []int{3, 11}
	target := []int{3, 0}
	for i, n := range nums {
		ret := findNthDigit(n)
		if ret != target[i] {
			log.ErrLog(errors.New(fmt.Sprintf("countDigitOne fail,n:%d,target:%d,ret:%d", n, target[i], ret)))
		}
	}

}
