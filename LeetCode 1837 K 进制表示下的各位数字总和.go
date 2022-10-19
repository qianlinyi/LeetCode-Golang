package main

func sumBase(n int, k int) int {
	ans := 0
	for n != 0 {
		ans += n % k
		n /= k
	}
	return ans
}
