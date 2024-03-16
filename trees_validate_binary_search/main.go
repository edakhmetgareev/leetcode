package main

import (
	"fmt"
	"math"
)

// Given the root of a binary tree, determine if it is a valid binary search tree (BST).
//
// A valid BST is defined as follows:
//
// The left subtree of a node contains only nodes with keys less than the node's key.
// The right subtree of a node contains only nodes with keys greater than the node's key.
// Both the left and right subtrees must also be binary search trees.
//
// Example 1:
//
// Input: root = [2,1,3]
// Output: true
// Example 2:
//
// Input: root = [5,1,4,null,null,3,6]
// Output: false
// Explanation: The root node's value is 5 but its right child's value is 4.
//
// Constraints:
//
// The number of nodes in the tree is in the range [1, 104].
// -231 <= Node.val <= 231 - 1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tree := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	isBst := isValidBST(tree)
	fmt.Printf("tree is binary: %v", isBst)
}

// Данное решение неверно, например для случая [5,4,6,null,null,3,7] .
// func isValidBST(root *TreeNode) bool {
//	if root == nil {
//		return true
//	}
//
//	if root.Left == nil && root.Right == nil {
//		return true
//	}
//
//	if root.Left != nil && root.Left.Val >= root.Val {
//		return false
//	}
//
//	if root.Right != nil && root.Val >= root.Right.Val {
//		return false
//	}
//
//	return isValidBST(root.Right) && isValidBST(root.Left)
//}

func isValidBST(root *TreeNode) bool {
	return valid(root, math.MinInt64, math.MaxInt64)
}

func valid(root *TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}
	return valid(root.Left, min, root.Val) && valid(root.Right, root.Val, max)
}
