package main

func largestAltitude(gain []int) int {
	sum, ans := 0, 0
	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}
	for _, a := range gain {
		sum += a
		ans = max(ans, sum)
	}
	return ans
}
