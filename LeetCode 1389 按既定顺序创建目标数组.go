package main

func createTargetArray(nums []int, index []int) []int {
	ans := make([]int, len(nums))
	for k, i := range index {
		copy(ans[i+1:], ans[i:])
		ans[i] = nums[k]
	}
	return ans
}
