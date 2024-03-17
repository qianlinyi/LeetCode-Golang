package main

func isValid(s string) bool {
	stack := make([]byte, 0)
	for i := range s {
		if len(stack) == 0 {
			stack = append(stack, s[i])
		} else {
			if stack[len(stack)-1] == '(' && s[i] == ')' || stack[len(stack)-1] == '[' && s[i] == ']' || stack[len(stack)-1] == '{' && s[i] == '}' {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, s[i])
			}
		}
	}
	return len(stack) == 0
}
