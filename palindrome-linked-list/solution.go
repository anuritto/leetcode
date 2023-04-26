package solution

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
	// dont need becouse "The number of nodes in the list is in the range [1, 105]""
	// if head == nil {
	// 	return false
	// }

	if head.Next == nil {
		return true
	}

	// our list length is even, so middle will be first node of second part
	middle := middleNode(head)

	// head of reversing second part of list
	rightCurrent := reverseList(middle)
	leftCurrent := head

	for rightCurrent != nil {
		if rightCurrent.Val != leftCurrent.Val {
			return false
		}
		rightCurrent = rightCurrent.Next
		leftCurrent = leftCurrent.Next
	}

	return true
}

// from 876 problem getting list 206
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

// from 206 problem list reversing
func reverseList(head *ListNode) *ListNode {
	cur := head
	var prev *ListNode

	for cur != nil {
		// saving next
		next := cur.Next

		// reversing link from next to prev
		cur.Next = prev

		// going next iteration
		prev = cur
		cur = next
	}

	return prev
}
