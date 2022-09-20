package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	q, ans := []*TreeNode{root}, 0
	for len(q) != 0 {
		node := q[0]
		if low <= node.Val && node.Val <= high {
			ans += node.Val
		}
		q = q[1:]
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
	}
	return ans
}
