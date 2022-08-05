package main

func countKDifference(nums []int, k int) int {
	cnt, ans := make([]int, 101), 0
	for _, num := range nums {
		cnt[num]++
	}
	for _, num := range nums {
		if num-k >= 0 {
			ans += cnt[num-k]
		}
		if num+k <= 100 {
			ans += cnt[num+k]
		}
	}
	return ans / 2
}
