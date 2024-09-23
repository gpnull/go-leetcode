package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	return dummy.Next
}

func main() {
	// Ví dụ 1
	list1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	list2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	mergedList := mergeTwoLists(list1, list2)
	printList(mergedList) // Output: 1 -> 1 -> 2 -> 3 -> 4 -> 4

	// Ví dụ 2
	list1 = nil
	list2 = nil
	mergedList = mergeTwoLists(list1, list2)
	printList(mergedList) // Output: nil

	// Ví dụ 3
	list1 = nil
	list2 = &ListNode{0, nil}
	mergedList = mergeTwoLists(list1, list2)
	printList(mergedList) // Output: 0
}

func printList(node *ListNode) {
	for node != nil {
		if node.Next != nil {
			print(node.Val, " -> ")
		} else {
			print(node.Val)
		}
		node = node.Next
	}
	println()
}

// You are given the heads of two sorted linked lists list1 and list2.

// Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.

// Return the head of the merged linked list.

// Example 1:

// Input: list1 = [1,2,4], list2 = [1,3,4]
// Output: [1,1,2,3,4,4]
// Example 2:

// Input: list1 = [], list2 = []
// Output: []
// Example 3:

// Input: list1 = [], list2 = [0]
// Output: [0]

// Constraints:

// The number of nodes in both lists is in the range [0, 50].
// -100 <= Node.val <= 100
// Both list1 and list2 are sorted in non-decreasing order.
