package averageoflevelsinbinarytree

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
func averageOfLevels(root *TreeNode) []float64 {

	res := []float64{}
	q := []*TreeNode{root}

	for len(q) > 0 {
		n := len(q)
		lvlSum := 0

		for i := 0; i < n; i++ {
			item := q[0]
			q = q[1:]
			lvlSum += item.Val

			if item.Right != nil {
				q = append(q, item.Right)
			}

			if item.Left != nil {
				q = append(q, item.Left)
			}

		}

		res = append(res, float64(lvlSum)/float64(n))
	}

	return res
}
