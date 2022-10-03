package main

import "sort"

func minPairSum(nums []int) int {
	sort.Ints(nums)
	ans, n := 0, len(nums)
	for i := 0; i < n/2; i++ {
		if nums[i]+nums[n-i-1] > ans {
			ans = nums[i] + nums[n-i-1]
		}
	}
	return ans
}
