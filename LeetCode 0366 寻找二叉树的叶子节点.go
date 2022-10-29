package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findLeaves(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return -1
		}
		l, r := dfs(node.Left), dfs(node.Right)
		depth := max(l, r) + 1
		if depth < len(ans) {
			ans[depth] = append(ans[depth], node.Val)
		} else {
			ans = append(ans, []int{node.Val})
		}
		return depth
	}
	dfs(root)
	return ans
}
