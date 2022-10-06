package main

import "sort"

func maxWidthOfVerticalArea(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] < points[j][1]
		}
		return points[i][0] < points[j][0]
	})
	ans := 0
	for i := 1; i < len(points); i++ {
		if points[i][0]-points[i-1][0] > ans {
			ans = points[i][0] - points[i-1][0]
		}
	}
	return ans
}
