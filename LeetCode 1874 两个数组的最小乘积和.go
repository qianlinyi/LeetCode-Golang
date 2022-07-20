package main

import "sort"

func minProductSum(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	sort.Sort(sort.Reverse(sort.IntSlice(nums2)))
	ans := 0
	for i := 0; i < len(nums1); i++ {
		ans += nums1[i] * nums2[i]
	}
	return ans
}
