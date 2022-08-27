package main

// Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

func cloneTree(root *Node) *Node {
	if root == nil {
		return nil
	}
	tree := &Node{Val: root.Val, Children: []*Node{}}
	for _, node := range root.Children {
		tree.Children = append(tree.Children, cloneTree(node))
	}
	return tree
}
