package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {
	cur, nxt := head, head.Next
	for nxt != nil {
		for nxt.Val != 0 {
			cur.Val += nxt.Val
			nxt = nxt.Next
		}
		if nxt.Next != nil {
			cur.Next = nxt
			cur = cur.Next
		} else {
			cur.Next = nil
		}
		nxt = nxt.Next
	}
	return head
}

func main() {}
