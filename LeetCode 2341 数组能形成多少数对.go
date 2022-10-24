package main

func numberOfPairs(nums []int) []int {
	cnt := make(map[int]int)
	cnt1, cnt2 := 0, 0
	for _, num := range nums {
		cnt[num]++
	}
	for _, value := range cnt {
		if value%2 == 1 {
			cnt1++
		}
		cnt2 += value / 2
	}
	return []int{cnt2, cnt1}
}
