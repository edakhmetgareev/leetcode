package main

import "fmt"

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
			Left: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 5,
			},
		},
	}

	res := levelOrder(tree)
	fmt.Printf("levels: %v", res)
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	q := []*TreeNode{root}
	result := make([][]int, 0)
	for len(q) > 0 {
		k := len(q)
		currentLevel := make([]int, 0, k)
		for i := 0; i < k; i++ {
			current := q[0]
			q = q[1:]
			if current.Left != nil {
				q = append(q, current.Left)
			}

			if current.Right != nil {
				q = append(q, current.Right)
			}
			currentLevel = append(currentLevel, current.Val)
		}

		result = append(result, currentLevel)
	}

	return result
}
