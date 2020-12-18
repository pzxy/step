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

func Test_complement(t *testing.T) {
	nums := []int{-1, 15}
	target := []string{"-1111111111111111,FFFF", "0000000000001111,000F"}
	for i, n := range nums {
		ret := complement(n)
		if ret != target[i] {
			log.ErrLog(errors.New("complement fail"))
		}
	}
}

func Test_minNumber(t *testing.T) {
	/**
	输入: [10,2]
	输出: "102"
	示例2:

	输入: [3,30,34,5,9]
	输出: "3033459"
	*/
	nums := [][]int{
		{10, 2},
		{3, 30, 34, 5, 9},
	}
	target := []string{"102", "3033459"}
	for i, n := range nums {
		ret := minNumber(n)
		if target[i] != ret {
			log.ErrLog(errors.New(fmt.Sprintf("min number fail: n:%v,target:%s,ret:%s \t\n", n, target[i], ret)))
		}
	}
}

func Test_translateNum(t *testing.T) {
	/**
	输入: 12258
	输出: 5
	解释: 12258有5种不同的翻译，分别是"bccfi", "bwfi", "bczi", "mcfi"和"mzi"
	*/
	nums := []int{12258}
	target := []int{5}
	for i, n := range nums {
		ret := translateNum(n)
		if target[i] != ret {
			log.ErrLog(errors.New(fmt.Sprintf("translateNum fail: n:%v,target:%d,ret:%d \t\n", n, target[i], ret)))
		}
	}
}

func Test_GiftMaxValue(t *testing.T) {
	/**
	输入:
	[
	 [1,3,1],
	 [1,5,1],
	 [4,2,1]
	]
	输出: 12
	解释: 路径 1→3→5→2→1 可以拿到最多价值的礼物
	*/
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	target := 12

	ret := maxValue(grid)
	if ret != target {
		log.ErrLog(errors.New(fmt.Sprintf("maxValue fail: grid:%v,target:%d,ret:%d \t\n", grid, target, ret)))
	}

	ret2 := maxValue4dp(grid)
	if ret2 != target {
		log.ErrLog(errors.New(fmt.Sprintf("maxValue4dp fail: grid:%v,target:%d,ret2:%d \t\n", grid, target, ret2)))
	}

	ret3 := maxValue4dp2(grid)
	if ret3 != target {
		log.ErrLog(errors.New(fmt.Sprintf("maxValue4dp fail: grid:%v,target:%d,ret3:%d \t\n", grid, target, ret3)))
	}
}

func Test_lengthOfLongestSubString(t *testing.T) {

	str := []string{
		"abcabcbb", "bbbbb", "pwwkew", "", " ",
	}
	target := []int{3, 1, 3, 0, 1}
	for i, s := range str {
		ret := lengthOfLongestSubstring(s)
		if ret != target[i] {
			log.ErrLog(errors.New(fmt.Sprintf("lengthOfLongestSubString fail: s:%s,target:%d,ret:%d \t\n", s, target[i], ret)))
		}
	}
}

func Test_nthUglyNumber(t *testing.T) {
	/**
	输入: n = 10
	输出: 12
	解释: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12
	*/
	nums := []int{10}
	target := []int{12}
	for i, n := range nums {
		ret := nthUglyNumber(n)
		if ret != target[i] {
			log.ErrLog(errors.New(fmt.Sprintf("nthUglyNumber fail: n:%d,target:%d,ret:%d \t\n", n, target[i], ret)))
		}
	}
}

func Test_firstAppearOnlyOneChar(t *testing.T) {
	/**
	s = "abaccdeff"
	返回 "b"

	s = ""
	返回 " "
	*/
	strs := []string{
		"leetcode", "abaccdeff", "",
	}
	target := []byte{'l', 'b', ' '}
	for i, s := range strs {
		ret := firstUniqChar(s)
		if target[i] != ret {
			log.ErrLog(errors.New(fmt.Sprintf("firstUniqChar fail: s:%s,target:%c,ret:%c \t\n", s, target[i], ret)))
		}
	}
}

func Test_reversePairs(t *testing.T) {
	/**
	输入: [7,5,6,4]
	输出: 5
	*/
	nums := [][]int{
		{1, 3, 2, 3, 1}, {7, 5, 6, 4}, {}, {0},
	}
	target := []int{4, 5, 0, 0}
	for i, n := range nums {
		ret := reversePairs(n)
		if ret != target[i] {
			log.ErrLog(errors.New(fmt.Sprintf("reverseParis fail:n:%v,target:%d,ret:%d \t\n", n, target[i], ret)))
		}
	}
}

func Test_getInterSectionNode(t *testing.T) {
	/**
	intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5],
	*/
	listA := []int{4, 1, 8, 4, 5}
	listB := []int{5, 0, 1, 8, 4, 5}

	la5 := &ListNode{Val: 5}
	la4 := &ListNode{Val: 4, Next: la5}
	la3 := &ListNode{Val: 8, Next: la4}
	la2 := &ListNode{Val: 1, Next: la3}
	ListA := &ListNode{Val: 4, Next: la2}

	lb3 := &ListNode{Val: 1, Next: la3}
	lb2 := &ListNode{Val: 0, Next: lb3}
	ListB := &ListNode{Val: 5, Next: lb2}

	target := 8
	ret := getIntersectionNode(ListA, ListB)
	if target != ret.Val {
		log.ErrLog(errors.New(fmt.Sprintf("reverseParis fail:listA:%v,listB:,%v,target:%d,ret:%d \t\n",
			listA, listB, target, ret.Val)))

	}

	//getListNode := func(list []int) *ListNode {
	//	if len(list) == 0 {
	//		return nil
	//	}
	//	var ListA = &ListNode{Val: list[0]}
	//	tmpA := ListA
	//	for _, v := range list[1:] {
	//		tmpA.Next = &ListNode{Val: v}
	//		tmpA = tmpA.Next
	//	}
	//	return ListA
	//}
	//listA := []int{4, 1, 8, 4, 5}
	//listB := []int{5, 0, 1, 8, 4, 5}
	//ListA := getListNode(listA)
	//ListB := getListNode(listB)
	//target := 8
	//ret := getIntersectionNode(ListA, ListB)
	//if target != ret.Val {
	//	log.ErrLog(errors.New(fmt.Sprintf("reverseParis fail:listA:%v,listB:,%v,target:%d,ret:%d \t\n",
	//		listA, listB, target, ret.Val)))
	//
	//}
}

func Test_SearchNumberFromOrderArray(t *testing.T) {
	/**
	输入: nums = [5,7,7,8,8,10], target = 8
	输出: 2
	示例2:

	输入: nums = [5,7,7,8,8,10], target = 6
	输出: 0
	*/
	nums := [][]int{
		{5, 7, 7, 8, 8, 10},
		{5, 7, 7, 8, 8, 10},
	}
	target := []int{8, 6}
	retTarget := []int{2, 0}
	for i, num := range nums {
		ret := search(num, target[i])
		if ret != retTarget[i] {
			log.ErrLog(errors.New(fmt.Sprintf("search fail :nums:%v,target:%d,retTarget:%d,ret:%d",
				num, target[i], retTarget[i], ret)))
		}
	}
}

func Test_missingNumber(t *testing.T) {
	/**
	输入: [0,1,3]
	输出: 2
	示例2:

	输入: [0,1,2,3,4,5,6,7,9]
	输出: 8
	*/
	nums := [][]int{
		{0, 1, 2, 3, 4, 5, 6, 7, 9},
		{0, 1, 3},
		{1},
	}
	target := []int{8, 2, 0}
	for i, num := range nums {
		ret := missingNumber(num)
		if ret != target[i] {
			log.ErrLog(errors.New(fmt.Sprintf("missingNumber fail: num:%v,tartget:%d,ret:%d \t\n", num, target[i], ret)))
		}
	}
}

func Test_kthLargest(t *testing.T) {
	/**
	输入: root = [3,1,4,null,2], k = 1
	   3
	  / \
	 1   4
	  \
	  2
	输出: 4
	示例 2:

	输入: root = [5,3,6,2,4,null,null,1], k = 3
	       5
	      / \
	     3   6
	    / \
	   2   4
	  /
	 1
	输出: 4
	*/
	r1 := &TreeNode{Val: 3}
	r1.Left = &TreeNode{Val: 1}
	r1.Right = &TreeNode{Val: 4}
	r1.Left.Right = &TreeNode{Val: 2}
	k1 := 1
	target1 := 4

	r2 := &TreeNode{Val: 5}
	r2.Left = &TreeNode{Val: 3}
	r2.Right = &TreeNode{Val: 6}
	r2.Left.Left = &TreeNode{Val: 2}
	r2.Left.Right = &TreeNode{Val: 4}
	r2.Left.Left.Left = &TreeNode{Val: 1}
	k2 := 3
	target2 := 4
	ret1 := kthLargest(r1, k1)
	if ret1 != target1 {
		log.ErrLog(errors.New(fmt.Sprintf("kthLargest1 fail:target:%d,ret:%d", target1, ret1)))
	}
	ret2 := kthLargest(r2, k2)
	if ret2 != target2 {
		log.ErrLog(errors.New(fmt.Sprintf("kthLargest2 fail:target:%d,ret:%d", target2, ret2)))
	}
}

func Test_binaryTreeMaxDepth(t *testing.T) {
	/**
	给定二叉树 [3,9,20,null,null,15,7]，
	    3
	   / \
	  9  20
	    /  \
	   15   7

	3
	*/
	r1 := &TreeNode{Val: 3}
	r1.Left = &TreeNode{Val: 9}
	r1.Right = &TreeNode{Val: 20}
	r1.Right.Left = &TreeNode{Val: 15}
	r1.Right.Right = &TreeNode{Val: 7}
	target1 := 3
	ret1 := maxDepth(r1)
	if target1 != ret1 {
		log.ErrLog(errors.New(fmt.Sprintf("maxDepth1 fail:target:%d,ret:%d", target1, ret1)))
	}
}

func Test_binaryTreeBalanced(t *testing.T) {
	/**
	给定二叉树 [3,9,20,null,null,15,7]

	    3
	   / \
	  9  20
	    /  \
	   15   7
	返回 true 。

	示例 2:

	给定二叉树 [1,2,2,3,3,null,null,4,4]

	       1
	      / \
	     2   2
	    / \
	   3   3
	  / \
	 4   4
	返回false 。
	*/
	r1 := &TreeNode{Val: 3}
	r1.Left = &TreeNode{Val: 9}
	r1.Right = &TreeNode{Val: 20}
	r1.Left.Left = &TreeNode{Val: 15}
	r1.Left.Right = &TreeNode{Val: 7}
	target1 := true

	r2 := &TreeNode{Val: 1}
	r2.Left = &TreeNode{Val: 2}
	r2.Right = &TreeNode{Val: 2}
	r2.Left.Left = &TreeNode{Val: 3}
	r2.Left.Right = &TreeNode{Val: 3}
	r2.Left.Left.Left = &TreeNode{Val: 4}
	r2.Left.Left.Right = &TreeNode{Val: 4}
	target2 := false
	ret1 := isBalanced(r1)
	if ret1 != target1 {
		log.ErrLog(errors.New(fmt.Sprintf("isBalanced1 fail:target:%v,ret:%v", target1, ret1)))
	}
	ret2 := isBalanced(r2)
	if ret2 != target2 {
		log.ErrLog(errors.New(fmt.Sprintf("isBalanced2 fail:target:%v,ret:%v", target2, ret2)))
	}

}

func Test_singleNumber4Twice(t *testing.T) {
	/**
	输入：nums = [4,1,4,6]
	输出：[1,6] 或 [6,1]
	示例 2：

	输入：nums = [1,2,10,4,1,4,3,3]
	输出：[2,10] 或 [10,2]
	*/
	nums := [][]int{
		{4, 1, 4, 6},
		{1, 2, 10, 4, 1, 4, 3, 3},
	}
	target := [][]int{
		{1, 6},
		{10, 2},
	}
	for i, num := range nums {
		ret := singleNumbers(num)
		if len(ret) != len(target[i]) {
			log.ErrLog(errors.New(fmt.Sprintf("singleNumbers fail of slice length not equal :num:%v,target:%v,ret:%v \n\t", num, target[i], ret)))
		}
		for _, v := range target[i] {
			find := false
			for _, r := range ret {
				if v == r {
					find = true
				}
			}
			if !find {
				log.ErrLog(errors.New(fmt.Sprintf("singleNumbers fail of value error:num:%v,target:%v,ret:%v \n\t", num, target[i], ret)))
				break
			}
		}
	}
}

func Test_singleNumber4Thrice(t *testing.T) {
	/**
	输入：nums = [3,4,3,3]
	输出：[4]
	输入：nums = [9,1,7,9,7,9,7]
	输出：[1]
	*/
	nums := [][]int{
		{3, 4, 3, 3},
		{9, 1, 7, 9, 7, 9, 7},
	}
	target := []int{
		4, 1,
	}
	for i, num := range nums {
		ret := singleNumbers4Thrice(num)
		if target[i] != ret {
			log.ErrLog(errors.New(fmt.Sprintf("singleNumbers4Thrice fail of value :num:%v,target:%v,ret:%v \n\t", num, target[i], ret)))
		}
	}
}

func Test_twoSum(t *testing.T) {
	/**
	输入：nums = [2,7,11,15], target = 9
	输出：[2,7] 或者 [7,2]
	示例 2：

	输入：nums = [10,26,30,31,47,60], target = 40
	输出：[10,30] 或者 [30,10]
	*/
	nums := [][]int{
		{2, 7, 11, 15},
		{10, 26, 30, 31, 47, 60},
	}
	target := []int{
		9, 40,
	}
	retTarget := [][]int{
		{2, 7},
		{30, 10},
	}
	for i, num := range nums {
		ret := twoSum(num, target[i])
		if len(ret) != len(retTarget[i]) {
			log.ErrLog(errors.New(fmt.Sprintf("twoSum fail of slice length not equal :num:%v,retTarget:%v,ret:%v \n\t", num, retTarget[i], ret)))
		}
		for _, v := range retTarget[i] {
			find := false
			for _, r := range ret {
				if v == r {
					find = true
				}
			}
			if !find {
				log.ErrLog(errors.New(fmt.Sprintf("twoSum fail of value error:num:%v,retTarget:%v,ret:%v \n\t", num, retTarget[i], ret)))
				break
			}
		}
	}
}

func Test_moreSum4findContinuousSequence(t *testing.T) {
	/**
	输入：target = 9
	输出：[[2,3,4],[4,5]]
	示例 2：

	输入：target = 15
	输出：[[1,2,3,4,5],[4,5,6],[7,8]]
	*/
	target := []int{9, 15}
	retTarget := [][][]int{
		{{2, 3, 4}, {4, 5}},
		{{1, 2, 3, 4, 5}, {4, 5, 6}, {7, 8}},
	}
	for i, t := range target {
		ret := findContinuousSequence(t)
		for j, v := range retTarget[i] {
			if len(v) != len(ret[j]) {
				log.ErrLog(errors.New(fmt.Sprintf("findContinuousSequence fail of length:"+
					" target:%d,retTarget:%v,ret:%v \n\t", target, retTarget[i], ret[j])))
			}
			for k, vv := range v {
				if vv != ret[j][k] {
					log.ErrLog(errors.New(fmt.Sprintf("findContinuousSequence fail of value:"+
						" target:%d,retTarget:%v,ret:%v \n\t", target, retTarget[i], ret[j])))
				}
			}
		}
	}
}

func Test_reverseWords(t *testing.T) {
	/**
	示例 1:
	输入: "the sky is blue"
	输出:"blue is sky the"
	示例 2：
	输入: " hello world! "
	输出:"world! hello"
	解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
	示例 3：
	输入: "a good  example"
	输出:"example good a"
	解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
	*/
	str := []string{"the sky    is blue", " hello world! ", "a good  example"}
	target := []string{"blue is sky the", "world! hello", "example good a"}
	for i, s := range str {
		ret := reverseWords(s)
		if ret != target[i] {
			log.ErrLog(errors.New(fmt.Sprintf("reverseWords fail: s:%s,target:%s,ret:%s \n\t", s, target[i], ret)))
		}
	}
}

func Test_reverseLeftWords(t *testing.T) {
	/**
	输入: s = "abcdefg", k = 2
	输出:"cdefgab"
	示例 2：

	输入: s = "lrloseumgh", k = 6
	输出:"umghlrlose"
	*/

	str := []string{"abcdefg", "lrloseumgh"}
	target := []int{2, 6}
	retTarget := []string{"cdefgab", "umghlrlose"}
	for i, s := range str {
		ret := reverseLeftWords(s, target[i])
		if ret != retTarget[i] {
			log.ErrLog(errors.New(fmt.Sprintf("reverseLeftWords fail: s:%s,retTarget:%s,ret:%s \n\t", s, retTarget[i], ret)))
		}
	}
}
