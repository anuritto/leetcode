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
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	if head.Next == nil {
		return false
	}

	// Floyd's tortoise and hare algorithm
	slow := head
	fast := head

	for slow != nil && fast != nil {
		slow = slow.Next

		if fast.Next != nil {
			fast = fast.Next.Next
		} else {
			fast = nil
		}

		if slow == fast {
			return true
		}

	}
	return false
}
