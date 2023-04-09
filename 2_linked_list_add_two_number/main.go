package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	res := addTwoNumbers(
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		}, &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 6,
				},
			},
		})

	// Result 5 -> 1 -> 0 -> 1
	print(*res)
}

/**
 * Definition for singly-linked list.
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var isOver bool
	result := &ListNode{}
	curResult := result
	for l1 != nil || l2 != nil || isOver {
		var l1Val, l2Val int
		if l1 == nil && l2 == nil && isOver {
			curResult.Next = &ListNode{
				Val: 1,
			}
			break
		}

		if l1 != nil {
			l1Val = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			l2Val = l2.Val
			l2 = l2.Next
		}

		curSum := l1Val + l2Val
		if isOver {
			curSum++
		}
		if curSum > 9 {
			isOver = true
			curSum = curSum - 10
		} else {
			isOver = false
		}
		curResult.Next = &ListNode{
			Val: curSum,
		}

		curResult = curResult.Next
	}

	return result.Next
}

func print(node ListNode) {
	fmt.Printf("%d, ", node.Val)
	if node.Next != nil {
		print(*node.Next)
	}
}
