package main

func maxDepth(s string) int {
	stack, ans := make([]int32, 0), 0
	for _, c := range s {
		if c == '(' {
			stack = append(stack, c)
		} else if c == ')' {
			if len(stack) > ans {
				ans = len(stack)
			}
			stack = stack[:len(stack)-1]
		}
	}
	return ans
}
