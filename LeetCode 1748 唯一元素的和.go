package main

func sumOfUnique(nums []int) int {
	cnt, ans := make(map[int]int, 0), 0
	for _, num := range nums {
		cnt[num]++
	}
	for _, num := range nums {
		if cnt[num] == 1 {
			ans += num
		}
	}
	return ans
}
