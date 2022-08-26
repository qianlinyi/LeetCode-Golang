package main

func shuffle(nums []int, n int) []int {
	ans := make([]int, n*2)
	for i, num := range nums[:n] {
		ans[2*i] = num
		ans[2*i+1] = nums[n+i]
	}
	return ans
}
