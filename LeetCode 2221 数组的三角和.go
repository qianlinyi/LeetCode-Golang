package main

func triangularSum(nums []int) int {
	for len(nums) > 1 {
		for i := 0; i < len(nums)-1; i++ {
			nums[i] = (nums[i] + nums[i+1]) % 10
		}
		nums = nums[:len(nums)-1]
	}
	return nums[0]
}
