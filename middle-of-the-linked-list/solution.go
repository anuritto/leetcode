package main

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
func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	// Floyd's tortoise and hare algorithm
	slow := head
	fast := head

	for {

		// fast pointer in the end, list length is odd

		if fast.Next == nil {
			return slow
		}
		// fast pointer in the end, list length is even
		if fast.Next.Next == nil {
			return slow.Next
		}

		// going to next iteration
		fast = fast.Next.Next
		slow = slow.Next
	}
}
