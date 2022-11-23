package main

func selfDividingNumbers(left int, right int) []int {
	ans := make([]int, 0)
	for i := left; i <= right; i++ {
		num, flag := i, true
		for num != 0 {
			x := num % 10
			if x == 0 || i%x != 0 {
				flag = false
				break
			}
			num /= 10
		}
		if flag == true {
			ans = append(ans, i)
		}
	}
	return ans
}
