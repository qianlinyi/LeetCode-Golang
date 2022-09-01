package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func dfs(root *TreeNode) (int, int, int) {
	if root == nil {
		return 0, 0, 0
	}
	suml, cntl, ansl := dfs(root.Left)
	sumr, cntr, ansr := dfs(root.Right)
	sum, cnt, ans := suml+sumr+root.Val, cntl+cntr+1, ansl+ansr
	if root.Val == sum/cnt {
		ans++
	}
	return sum, cnt, ans
}

func averageOfSubtree(root *TreeNode) int {
	_, _, ans := dfs(root)
	return ans
}
