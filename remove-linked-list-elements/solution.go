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
func removeElements(head *ListNode, val int) *ListNode {
	var prev *ListNode
	current := head
	for current != nil {
		if current.Val == val {
			if prev == nil {
				head = head.Next
			} else {
				prev.Next = current.Next
			}
		} else {
			prev = current
		}

		current = current.Next
	}
	return head
}

// with false head and without nill prev head check
func removeElements2(head *ListNode, val int) *ListNode {
	prevHead := &ListNode{
		Next: head,
	}

	prev := prevHead
	current := head

	for current != nil {
		if current.Val == val {
			// skip this node
			prev.Next = current.Next
		} else {
			// next iteration for prev point
			prev = current
		}

		// next iteration for current point
		current = current.Next
	}
	return prevHead.Next
}
