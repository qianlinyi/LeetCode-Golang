package main

import (
	"sort"
)

func smallerNumbersThanCurrent(nums []int) []int {
	ans, temp := make([]int, len(nums)), make([]int, len(nums))
	for i, num := range nums {
		temp[i] = num
	}
	sort.Ints(nums)
	for i, num := range temp {
		pos := sort.Search(len(nums), func(i int) bool {
			return nums[i] >= num
		})
		ans[i] = pos
	}
	return ans
}

func main() {
}
