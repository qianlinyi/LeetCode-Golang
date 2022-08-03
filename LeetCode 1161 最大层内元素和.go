package main

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	mx, ans, layer, queue := root.Val, 1, 1, []*TreeNode{root}
	for layer = 1; len(queue) > 0; layer++ {
		sum := 0
		tmp := queue
		queue = nil
		for _, node := range tmp {
			sum += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if sum > mx {
			mx, ans = sum, layer
		}
	}
	return ans
}
