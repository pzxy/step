package leetcode

import (
	"errors"
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
