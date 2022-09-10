package main

func minBitFlips(start int, goal int) int {
	num, ans := start^goal, 0
	for num != 0 {
		ans++
		num &= num - 1
	}
	return ans
}
