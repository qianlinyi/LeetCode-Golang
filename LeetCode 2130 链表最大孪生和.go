package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 快慢指针+反转链表
func pairSum(head *ListNode) int {
	slow, fast := head, head.Next
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	pre := slow.Next
	for pre.Next != nil {
		cur := pre.Next
		pre.Next = cur.Next
		cur.Next = pre
		slow.Next = cur
	}
	ans, x, y := 0, head, slow.Next
	for y != nil {
		if x.Val+y.Val > ans {
			ans = x.Val + y.Val
		}
		x, y = x.Next, y.Next
	}
	return ans
}
