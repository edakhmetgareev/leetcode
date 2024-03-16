package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return isSameTree(root.Left, root.Right)
}

func isSameTree(l *TreeNode, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}

	return l != nil && r != nil && l.Val == r.Val && isSameTree(l.Left, r.Right) && isSameTree(l.Right, r.Left)
}

func main() {
	ex1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}

	ex2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}

	root := &TreeNode{
		Val:   1,
		Left:  ex1,
		Right: ex2,
	}

	fmt.Println("Result:", isSymmetric(root))
}
