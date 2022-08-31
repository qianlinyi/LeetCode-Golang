package main

// 双指针
func pivotArray(nums []int, pivot int) []int {
	n := len(nums)
	ans := make([]int, n)
	left, right := 0, n-1
	for _, num := range nums {
		if num < pivot {
			ans[left] = num
			left++
		} else if num > pivot {
			ans[right] = num
			right--
		}
	}
	for i := left; i < right+1; i++ {
		ans[i] = pivot
	}
	x, y := right+1, n-1
	for x < y {
		ans[x], ans[y] = ans[y], ans[x]
		x++
		y--
	}
	return ans
}
