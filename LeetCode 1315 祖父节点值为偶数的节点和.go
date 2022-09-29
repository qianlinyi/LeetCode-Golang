package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumEvenGrandparent(root *TreeNode) int {
	parent, ans := make(map[*TreeNode]*TreeNode), 0
	parent[root] = nil
	parent[nil] = nil
	var q = []*TreeNode{root}
	for len(q) != 0 {
		node := q[0]
		ancester := parent[parent[node]]
		if ancester != nil && ancester.Val%2 == 0 {
			ans += node.Val
		}
		q = q[1:]
		if node.Left != nil {
			q = append(q, node.Left)
			parent[node.Left] = node
		}
		if node.Right != nil {
			q = append(q, node.Right)
			parent[node.Right] = node
		}
	}
	return ans
}
