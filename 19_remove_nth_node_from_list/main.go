package main

import "fmt"

func main() {
	// result1 := myAtoi("  -0012a42")
	// fmt.Printf("Result: %d, expected: -12 \n", result1)
	result := removeNthFromEnd(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, 2)
	// result := removeNthFromEnd(&ListNode{Val: 1}, 1)

	for {
		if result == nil {
			break
		}
		if result.Next == nil {
			break
		}
		fmt.Println(result.Val)
		result = result.Next
	}
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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n < 0 {
		return head
	}

	m := make([]*ListNode, 0, 31)
	currNode := head
	i := 0
	for {
		m = append(m, currNode)
		if currNode.Next == nil {
			if n == 1 {
				if i == 0 {
					return nil
				}

				m[i-1].Next = nil
				break
			}

			if i == n-1 {
				return head.Next
			}

			if i < n {
				return head
			}

			m[i-n].Next = m[i-n+2]
			break
		}

		currNode = currNode.Next
		i++
	}

	return head
}

func removeNthFromEndNotMine(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy

	for i := 0; i <= n; i++ {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}
