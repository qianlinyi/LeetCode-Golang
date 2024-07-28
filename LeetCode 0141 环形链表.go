package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	s := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := s[head]; ok {
			return true
		}
		s[head] = struct{}{}
		head = head.Next
	}
	return false
}

func main() {
	L := &ListNode{Val: 3}
	L.Next = &ListNode{Val: 2}
	L.Next.Next = &ListNode{Val: 0}
	L.Next.Next.Next = &ListNode{Val: -4}
	L.Next.Next.Next.Next = L
	fmt.Println(hasCycle(L))
}
