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
