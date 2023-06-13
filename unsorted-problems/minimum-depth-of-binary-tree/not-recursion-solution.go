package minimumdepthofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	q := []*TreeNode{root}
	lvl := 1

	for len(q) > 0 {
		cnt := len(q)
		for i := 0; i < cnt; i++ {
			item := q[0]
			q = q[1:]

			if item.Left == nil && item.Right == nil {
				return lvl
			}

			if item.Left != nil {
				q = append(q, item.Left)
			}

			if item.Right != nil {
				q = append(q, item.Right)
			}
		}

		lvl++
	}
	return lvl
}
