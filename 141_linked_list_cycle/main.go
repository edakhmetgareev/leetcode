package main

// Given head, the head of a linked list, determine if the linked list has a cycle in it.
//
// There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.
//
// Return true if there is a cycle in the linked list. Otherwise, return false.

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list := &ListNode{}
	hasCycle(list)
}

func hasCycle(head *ListNode) bool {
	hash := make(map[*ListNode]struct{})
	current := head

	for current != nil {
		if _, ok := hash[current]; ok {
			return true
		}
		hash[current] = struct{}{}
		current = current.Next
	}

	return false
}
