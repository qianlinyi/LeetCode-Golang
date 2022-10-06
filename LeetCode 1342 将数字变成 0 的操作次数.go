package main

func numberOfSteps(num int) int {
	ans := 0
	for num != 0 {
		if num%2 == 0 {
			num /= 2
		} else {
			num -= 1
		}
		ans++
	}
	return ans
}
