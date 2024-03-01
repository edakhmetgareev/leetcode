package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CopyList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	node := &ListNode{Val: head.Val}

	node.Next = CopyList(head.Next)

	return node
}

func TreeClone(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	clone := &TreeNode{Val: root.Val}
	clone.Left = TreeClone(root.Left)
	clone.Right = TreeClone(root.Right)

	return clone
}
