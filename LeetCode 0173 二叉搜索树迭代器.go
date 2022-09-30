package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	ans := BSTIterator{stack: []*TreeNode{}}
	for root != nil {
		ans.stack = append(ans.stack, root)
		root = root.Left
	}
	return ans
}

func (this *BSTIterator) Next() int {
	cur := this.stack[len(this.stack)-1]
	node := cur.Right
	this.stack = this.stack[:len(this.stack)-1]
	for node != nil {
		this.stack = append(this.stack, node)
		node = node.Left
	}
	return cur.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
