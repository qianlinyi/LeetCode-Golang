package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getLonelyNodes(root *TreeNode) []int {
	ans, q := make([]int, 0), []*TreeNode{root}
	for len(q) != 0 {
		node := q[0]
		if node.Left == nil && node.Right == nil {
			continue
		} else if node.Left == nil {
			ans = append(ans, node.Right.Val)
			q = append(q, node.Right)
		} else if node.Right == nil {
			ans = append(ans, node.Left.Val)
			q = append(q, node.Left)
		}
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
		q = q[1:]
	}
	return ans
}
