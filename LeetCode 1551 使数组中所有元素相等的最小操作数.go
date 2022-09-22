package main

func minOperations(n int) int {
	sum := 0
	for i := n - 1; i > 0; i -= 2 {
		sum += 2 * i
	}
	return sum / 2
}
