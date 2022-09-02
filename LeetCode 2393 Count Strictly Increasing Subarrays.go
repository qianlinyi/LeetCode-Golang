package main

func countSubarrays(nums []int) int64 {
	var ans, cnt int64 = 0, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			cnt++
		} else {
			ans += cnt * (cnt + 1) / 2
			cnt = 1
		}
	}
	ans += cnt * (cnt + 1) / 2
	return ans
}
