package main

import (
	"fmt"
)

// There is a singly-linked list head and we want to delete a node node in it.
//
// You are given the node to be deleted node. You will not be given access to the first node of head.
//
// All the values of the linked list are unique, and it is guaranteed that the given node node is not the last node in the linked list.
//
// Delete the given node. Note that by deleting the node, we do not mean removing it from memory. We mean:
//
// The value of the given node should not exist in the linked list.
// The number of nodes in the linked list should decrease by one.
// All the values before node should be in the same order.
// All the values after node should be in the same order.
// Custom testing:
//
// For the input, you should provide the entire linked list head and the node to be given node. node should not be the last node of the list and should be an actual node in the list.
// We will build the linked list and pass the node to your function.
// The output will be the entire list after calling your function.

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}
	reverseListV2(node)
	fmt.Printf("%+v", node)
}

func reverseList(head *ListNode) *ListNode {
	var reversedList *ListNode
	for head != nil {
		if reversedList == nil {
			reversedList = &ListNode{
				Val: head.Val,
			}
		} else {
			reversedList = &ListNode{
				Val:  head.Val,
				Next: reversedList,
			}
		}
		head = head.Next
	}

	return reversedList
}

func reverseListV2(head *ListNode) *ListNode {
	var reversedList *ListNode
	for head != nil {
		tmpNode := head.Next
		head.Next = reversedList
		reversedList = head
		head = tmpNode
	}

	return reversedList
}
