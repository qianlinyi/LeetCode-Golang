package main

func findNumbers(nums []int) int {
	ans := 0
	for _, num := range nums {
		cnt := 0
		for num != 0 {
			num /= 10
			cnt++
		}
		if cnt%2 == 0 {
			ans++
		}
	}
	return ans
}
