package mergetwosortedlists

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
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{}
	curr := result

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			curr.Next = &ListNode{list1.Val, nil}
			list1 = list1.Next
		} else {
			curr.Next = &ListNode{list2.Val, nil}
			list2 = list2.Next
		}

		curr = curr.Next
	}

	if list1 != nil {
		curr.Next = list1
	}

	if list2 != nil {
		curr.Next = list2
	}

	return result.Next

}
