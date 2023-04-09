package main

import "fmt"

// Given a binary tree, determine if it is height-balanced.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   4,
					Right: nil,
				},
				Right: &TreeNode{
					Val: 4,
					Right: &TreeNode{
						Val:   5,
						Right: nil,
					},
				},
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}

	fmt.Println("Result:", isBalanced(root))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	ln := depth(root.Left)
	rn := depth(root.Right)

	if ln-rn > 1 || rn-ln > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func depth(n *TreeNode) int {
	if n == nil {
		return 0
	}

	ld := depth(n.Left)
	rd := depth(n.Right)

	if ld > rd {
		return ld + 1
	}

	return rd + 1
}
