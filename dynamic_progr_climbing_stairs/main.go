package main

import (
	"fmt"
)

// Given the root of a binary tree, return its maximum depth.
//
// A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
//
//
//
// Example 1:
//
//
// Input: root = [3,9,20,null,null,15,7]
// Output: 3
// Example 2:
//
// Input: root = [1,null,2]
// Output: 2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 2,
		},
	}
	depth := maxDepth(tree)
	fmt.Printf("%+v", depth)
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func recursiveDepth(leaf *TreeNode, depth int) int {
	if leaf == nil || (leaf.Right == nil && leaf.Left == nil) {
		return depth
	}

	x := recursiveDepth(leaf.Right, depth+1)
	y := recursiveDepth(leaf.Left, depth+1)

	if x > y {
		return x
	}

	return y
}
