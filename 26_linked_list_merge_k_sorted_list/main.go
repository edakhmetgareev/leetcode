package main

import (
	"fmt"
	"sort"
)

// You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.
// Merge all the linked-lists into one sorted linked-list and return it.
//
// Example 1:
// Input: lists = [[1,4,5],[1,3,4],[2,6]]
// Output: [1,1,2,3,4,4,5,6]
// Explanation: The linked-lists are:
// [
//  1->4->5,
//  1->3->4,
//  2->6
// ]
// merging them into one sorted list:
// 1->1->2->3->4->4->5->6

// Example 2:
// Input: lists = []
// Output: []
// Example 3:
//
// Input: lists = [[]]
// Output: []
//
// Constraints:
//
// k == lists.length
// 0 <= k <= 104
// 0 <= lists[i].length <= 500
// -104 <= lists[i][j] <= 104
// lists[i] is sorted in ascending order.
// The sum of lists[i].length will not exceed 104.

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// [[1,4,5],[1,3,4],[2,6]]
	list := []*ListNode{
		{
			Val: 1,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 5,
				},
			},
		},
		{
			Val: 1,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
		{
			Val: 2,
			Next: &ListNode{
				Val: 6,
			},
		},
	}
	result := mergeKLists(list)
	print(*result)
	fmt.Printf("\n---------")
	// result := mergeKLists([]*ListNode{
	// 	nil,
	// 	{
	// 		Val: 0,
	// 		Next: &ListNode{
	// 			Val: 2,
	// 			Next: &ListNode{
	// 				Val: 5,
	// 			},
	// 		},
	// 	},
	// })
	// print(*result)
}

func print(node ListNode) {
	fmt.Printf("%d, ", node.Val)
	if node.Next != nil {
		print(*node.Next)
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// [[1,4,5],[1,3,4],[2,6]]
// [[1,1,4,5],[3,4],[2,6]]
func mergeKLists(lists []*ListNode) *ListNode {
	listCnt := len(lists)

	if listCnt == 0 {
		return nil
	}

	if listCnt == 1 {
		return lists[0]
	}

	var result, cur *ListNode
	for i := 0; i < listCnt; i++ {
		list := lists[i]
		if list == nil {
			continue
		}

		if result == nil {
			result = &ListNode{
				Val:  list.Val,
				Next: list.Next,
			}
			continue
		}

		cur = result
		for {
			if list.Val <= cur.Val {
				oldCur := *cur
				cur.Val = list.Val
				cur.Next = &oldCur

				if list.Next == nil {
					break
				}
				list = list.Next

				continue
			}

			if cur.Next == nil {
				cur.Next = list
				break
			}

			cur = cur.Next
		}

		// setNext(result, *lists[i])
	}

	return result
}

func mergeKListsV2(lists []*ListNode) *ListNode {
	var plainLists []*ListNode

	for _, list := range append(plainLists, lists...) {
		for {
			if list == nil {
				break
			}
			plainLists = append(plainLists, list)
			list = list.Next
		}
	}

	if len(plainLists) == 0 {
		return nil
	}

	sort.Slice(plainLists, func(i, j int) bool {
		return plainLists[i].Val < plainLists[j].Val
	})

	for i := 1; i < len(plainLists); i += 1 {
		plainLists[i-1].Next = plainLists[i]
	}

	if len(plainLists) != 0 {
		plainLists[len(plainLists)-1].Next = nil
	}

	return plainLists[0]
}

func mergeKListsV3(lists []*ListNode) *ListNode {
	var plainLists []*ListNode

	for _, list := range append(plainLists, lists...) {
		for {
			if list == nil {
				break
			}
			plainLists = append(plainLists, list)
			list = list.Next
		}
	}

	if len(plainLists) == 0 {
		return nil
	}

	sort.Slice(plainLists, func(i, j int) bool {
		return plainLists[i].Val < plainLists[j].Val
	})

	for i := 1; i < len(plainLists); i += 1 {
		plainLists[i-1].Next = plainLists[i]
	}

	if len(plainLists) != 0 {
		plainLists[len(plainLists)-1].Next = nil
	}

	return plainLists[0]
}

func setNext(result *ListNode, nextNode ListNode) {
	if nextNode.Val < result.Val {
		nextHead, nextTail := getHeadAndTail(nextNode)
		oldResult := *result
		result.Val = nextHead.Val
		result.Next = &oldResult

		if nextTail != nil {
			setNext(result, *nextTail)
		}
		return
	}
	if result.Next == nil {
		result.Next = &nextNode
		return
	}

	setNext(result.Next, nextNode)
}

func getHeadAndTail(node ListNode) (*ListNode, *ListNode) {
	return &ListNode{
		Val: node.Val,
	}, node.Next
}
