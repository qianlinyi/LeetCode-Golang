package main

import "sort"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	res := make([]int, 0)
	bfs := func(node *TreeNode) []int {
		ans, q := make([]int, 0), []*TreeNode{node}
		if node == nil {
			return ans
		}
		for len(q) != 0 {
			node := q[0]
			q = q[1:]
			ans = append(ans, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		return ans
	}
	res = append(append(res, bfs(root1)...), bfs(root2)...)
	sort.Ints(res)
	return res
}
