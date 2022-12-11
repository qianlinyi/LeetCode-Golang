package main

func diStringMatch(s string) []int {
	mn, mx := 0, len(s)
	ans := make([]int, mx+1)
	for i, c := range s {
		if c == 'I' {
			ans[i] = mn
			mn++
		} else {
			ans[i] = mx
			mx--
		}
	}
	ans[len(s)] = mn
	return ans
}
