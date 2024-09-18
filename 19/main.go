package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	first, second := dummy, dummy

	for i := 0; i <= n; i++ {
		first = first.Next
	}

	for first != nil {
		first = first.Next
		second = second.Next
	}

	second.Next = second.Next.Next
	return dummy.Next
}

func main() {
	// Example 1: head = [1,2,3,4,5], n = 2
	head1 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	result1 := removeNthFromEnd(head1, 3)
	printList(result1) // Output: [1,2,3,5]

	// Example 2: head = [1], n = 1
	head2 := &ListNode{1, nil}
	result2 := removeNthFromEnd(head2, 1)
	printList(result2) // Output: []

	// Example 3: head = [1,2], n = 1
	head3 := &ListNode{1, &ListNode{2, nil}}
	result3 := removeNthFromEnd(head3, 1)
	printList(result3) // Output: [1]
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(",")
		}
		head = head.Next
	}
	fmt.Println()
}

// Given the head of a linked list, remove the nth node from the end of the list and return its head.

// Example 1:

// Input: head = [1,2,3,4,5], n = 2
// Output: [1,2,3,5]
// Example 2:

// Input: head = [1], n = 1
// Output: []
// Example 3:

// Input: head = [1,2], n = 1
// Output: [1]

// Constraints:

// The number of nodes in the list is sz.
// 1 <= sz <= 30
// 0 <= Node.val <= 100
// 1 <= n <= sz

// Follow up: Could you do this in one pass?
