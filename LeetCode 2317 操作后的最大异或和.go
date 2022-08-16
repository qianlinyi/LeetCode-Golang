package main

func maximumXOR(nums []int) int {
	ans := 0
	for i := 0; i < 30; i++ {
		flag := 0
		for j := 0; j < len(nums); j++ {
			if (1<<i)&nums[j] != 0 {
				flag = 1
				break
			}
		}
		if flag == 1 {
			ans += 1 << i
		}
	}
	return ans
}
