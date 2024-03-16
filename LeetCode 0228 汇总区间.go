package main

import "strconv"

func summaryRanges(nums []int) []string {
	ans := make([]string, 0)
	left, right := 0, 0
	for right < len(nums) {
		for right < len(nums) && nums[right]-nums[left] == right-left {
			right++
		}
		if left == right-1 {
			ans = append(ans, strconv.Itoa(nums[left]))
		} else {
			ans = append(ans, strconv.Itoa(nums[left])+"->"+strconv.Itoa(nums[right-1]))
		}
		left = right
	}
	return ans
}
