package main

func findGCD(nums []int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}
	min := func(x, y int) int {
		if x < y {
			return x
		} else {
			return y
		}
	}
	gcd := func(x, y int) int {
		if x < y {
			x, y = y, x
		}
		for y != 0 {
			x, y = y, x%y
		}
		return x
	}
	mx, mn := 1, 1000
	for _, num := range nums {
		mx = max(mx, num)
		mn = min(mn, num)
	}
	return gcd(mx, mn)
}
