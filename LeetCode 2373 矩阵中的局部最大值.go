package main

//func max(a, b int) int {
//	if a > b {
//		return a
//	} else {
//		return b
//	}
//}

func largestLocal(grid [][]int) [][]int {
	n := len(grid)
	for i := 0; i < n-2; i++ {
		for j := 0; j < n-2; j++ {
			mx := 0
			for x := i; x < i+3; x++ {
				for y := j; y < j+3; y++ {
					mx = max(mx, grid[x][y])
				}
			}
			grid[i][j] = mx
		}
		grid[i] = grid[i][:n-2]
	}
	return grid[:n-2]
}
