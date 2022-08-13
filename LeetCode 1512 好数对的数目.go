package main

func numIdenticalPairs(nums []int) int {
	cnt, ans := make([]int, 101), 0
	for _, num := range nums {
		ans += cnt[num]
		cnt[num]++
	}
	return ans
}
