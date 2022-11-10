package main

import "sort"

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	ans := 0
	for idx, num := range nums {
		if idx%2 == 0 {
			ans += num
		}
	}
	return ans
}
