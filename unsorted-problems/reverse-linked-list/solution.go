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
