package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	first := head
	second := head.Next

	first.Next = swapPairs(second.Next)
	second.Next = first

	return second
}

func printList(head *ListNode) {
	for head != nil {
		print(head.Val, " ")
		head = head.Next
	}
	println()
}

func main() {
	head1 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	printList(swapPairs(head1))

	var head2 *ListNode
	printList(swapPairs(head2))

	head3 := &ListNode{1, nil}
	printList(swapPairs(head3))

	head4 := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	printList(swapPairs(head4))
}

//  Given a linked list, swap every two adjacent nodes and return its head. You must solve the problem without modifying the values in the list's nodes (i.e., only nodes themselves may be changed.)

//  Example 1:

//  Input: head = [1,2,3,4]

//  Output: [2,1,4,3]

//  Explanation:

//  Example 2:

//  Input: head = []

//  Output: []

//  Example 3:

//  Input: head = [1]

//  Output: [1]

//  Example 4:

//  Input: head = [1,2,3]

//  Output: [2,1,3]

//  Constraints:

//  The number of nodes in the list is in the range [0, 100].
//  0 <= Node.val <= 100
