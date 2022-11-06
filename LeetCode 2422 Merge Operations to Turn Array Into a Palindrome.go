package main

func minimumOperations(nums []int) int {
	l, r := 0, len(nums)-1
	suml, sumr, ans := nums[l], nums[r], 0
	for l < r {
		if suml < sumr {
			l++
			suml += nums[l]
			ans += 1
		} else if suml > sumr {
			r--
			sumr += nums[r]
			ans += 1
		} else {
			l++
			r--
			suml, sumr = nums[l], nums[r]
		}
	}
	return ans
}
