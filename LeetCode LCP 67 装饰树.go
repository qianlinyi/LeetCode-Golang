package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func expandBinaryTree(root *TreeNode) *TreeNode {
	if root.Left != nil {
		root.Left = &TreeNode{
			-1, expandBinaryTree(root.Left), nil,
		}
	}
	if root.Right != nil {
		root.Right = &TreeNode{
			-1, nil, expandBinaryTree(root.Right),
		}
	}
	return root
}
