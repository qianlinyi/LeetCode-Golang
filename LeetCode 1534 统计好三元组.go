package main

func countGoodTriplets(arr []int, a int, b int, c int) int {
	abs := func(a int) int {
		if a < 0 {
			return -a
		} else {
			return a
		}
	}
	n, ans := len(arr), 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if abs(arr[i]-arr[j]) <= a && abs(arr[j]-arr[k]) <= b && abs(arr[i]-arr[k]) <= c {
					ans++
				}
			}
		}
	}
	return ans
}
