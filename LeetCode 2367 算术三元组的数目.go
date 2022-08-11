package main

func arithmeticTriplets(nums []int, diff int) int {
	cnt, ans := make([]int, 201), 0
	for _, num := range nums {
		cnt[num]++
	}
	for _, num := range nums {
		if num-diff >= 0 && cnt[num-diff] != 0 && num+diff <= 200 && cnt[num+diff] != 0 {
			ans++
		}
	}
	return ans
}
