package main

func subsetXORSum(nums []int) int {
	sum := 0
	for i := 0; i < (1 << len(nums)); i++ {
		xor := 0
		for j := 0; j < len(nums); j++ {
			if i&(1<<j) != 0 {
				xor ^= nums[j]
			}
		}
		sum += xor
	}
	return sum
}
