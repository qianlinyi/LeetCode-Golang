package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
	ans := 0
	for head != nil {
		if head.Val == 1 {
			ans = ans<<1 | 1
		} else {
			ans = ans << 1
		}
		head = head.Next
	}
	return ans
}
