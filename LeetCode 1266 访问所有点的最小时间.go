package main

func maxInt(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

func minTimeToVisitAllPoints(points [][]int) int {
	ans := 0
	for i := 1; i < len(points); i++ {
		mx := maxInt(abs(points[i-1][0]-points[i][0]), abs(points[i-1][1]-points[i][1]))
		ans += mx
	}
	return ans
}
