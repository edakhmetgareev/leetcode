package main

import "fmt"

// Given the head of a singly linked list, return true if it is a palindrome or false otherwise.
//
//
//
//Example 1:
//
//
//Input: head = [1,2,2,1]
//Output: true
//Example 2:
//
//
//Input: head = [1,2]
//Output: false
//
//
//Constraints:
//
//The number of nodes in the list is in the range [1, 105].
//0 <= Node.val <= 9
//
//
//Follow up: Could you do it in O(n) time and O(1) space?

func main() {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 1,
						},
					},
				},
			},
		},
	}

	res := isPalindromeV2(list)
	fmt.Printf("Result: %+v", res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}

	listLen := 0
	current := head
	lists := make([]*ListNode, 0)
	for current != nil {
		listLen++
		lists = append(lists, current)
		current = current.Next
	}

	if listLen == 1 {
		return true
	}

	for i := 0; i < listLen/2; i++ {
		if lists[i].Val != lists[listLen-i-1].Val {
			return false
		}
	}

	return true
}

func isPalindromeV2(head *ListNode) bool {
	if head == nil {
		return false
	}

	list1, list2 := head, head
	for list2 != nil && list2.Next != nil {
		list1 = list1.Next
		list2 = list2.Next.Next
	}

	// reverse second part of the list
	list1 = reverseList(list1)

	for list1 != nil {
		if head.Val != list1.Val {
			return false
		}
		head = head.Next
		list1 = list1.Next
	}

	return true
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	var futr *ListNode

	for head != nil {
		futr = head.Next
		head.Next = prev
		prev = head
		head = futr
	}

	return prev
}
