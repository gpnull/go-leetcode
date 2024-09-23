package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	// Check if there are at least k nodes left in the list
	count := 0
	current := head
	for current != nil && count < k {
		current = current.Next
		count++
	}
	if count < k {
		return head // Not enough nodes to reverse
	}

	// Reverse k nodes
	prev := (*ListNode)(nil)
	current = head
	for i := 0; i < k; i++ {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	// Recursively call for the next part of the list
	head.Next = reverseKGroup(current, k)

	return prev
}

func main() {
	// Create a sample linked list: 1 -> 2 -> 3 -> 4 -> 5
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}

	// Reverse in groups of k
	k := 2
	newHead := reverseKGroup(head, k)

	// Print the modified list
	for newHead != nil {
		fmt.Print(newHead.Val, " ")
		newHead = newHead.Next
	}
}

//  Given the head of a linked list, reverse the nodes of the list k at a time, and return the modified list.

//  k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes, in the end, should remain as it is.

//  You may not alter the values in the list's nodes, only nodes themselves may be changed.

//  Example 1:

//  Input: head = [1,2,3,4,5], k = 2
//  Output: [2,1,4,3,5]
//  Example 2:

//  Input: head = [1,2,3,4,5], k = 3
//  Output: [3,2,1,4,5]

//  Constraints:

//  The number of nodes in the list is n.
//  1 <= k <= n <= 5000
//  0 <= Node.val <= 1000

//  Follow-up: Can you solve the problem in O(1) extra memory space?
