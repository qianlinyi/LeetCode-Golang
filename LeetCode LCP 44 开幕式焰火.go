package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type void struct{}

func numColor(root *TreeNode) int {
	set := make(map[int]void)
	var exist void
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		set[node.Val] = exist
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return len(set)
}
