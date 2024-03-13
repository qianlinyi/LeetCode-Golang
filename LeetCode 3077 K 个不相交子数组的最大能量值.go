package main

import "math"

func maximumStrength(nums []int, k int) int64 {
	n := len(nums)
	s := make([]int, n+1)
	for i, x := range nums {
		s[i+1] = s[i] + x
	}
	dp := make([]int, n+1)
	for i := 1; i <= k; i++ {
		pre := dp[i-1]
		dp[i-1] = math.MinInt
		mx := math.MinInt
		w := (k - i + 1) * (i%2*2 - 1)
		for j := i; j <= n-k+1; j++ {
			mx = max(mx, pre-s[j-1]*w)
			pre = dp[j]
			dp[j] = max(dp[j-1], s[j]*w+mx)
		}
	}
	return int64(dp[n])
}
