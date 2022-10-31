package main

func countTriplets(arr []int) int {
	n, ans := len(arr), 0
	s := make([]int, n+1)
	for i := 0; i < n; i++ {
		s[i+1] = s[i] ^ arr[i]
	}
	for i := 0; i < n; i++ {
		for k := i + 1; k < n; k++ {
			if s[i] == s[k+1] {
				ans += k - i
			}
		}
	}
	return ans
}
