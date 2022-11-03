package main

func replaceElements(arr []int) []int {
	n, mx := len(arr), -1
	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}
	for i := n - 1; i >= 0; i-- {
		temp := arr[i]
		arr[i] = mx
		mx = max(temp, arr[i])
	}
	return arr
}
