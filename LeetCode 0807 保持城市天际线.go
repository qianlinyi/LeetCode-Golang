package main

func maxIncreaseKeepingSkyline(grid [][]int) int {
	r, c, ans := len(grid), len(grid[0]), 0
	r_max, c_max := make([]int, r), make([]int, c)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if grid[i][j] > r_max[i] {
				r_max[i] = grid[i][j]
			}
		}
	}
	for i := 0; i < c; i++ {
		for j := 0; j < r; j++ {
			if grid[j][i] > c_max[i] {
				c_max[i] = grid[j][i]
			}
		}
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if r_max[i] > c_max[j] {
				ans += c_max[j] - grid[i][j]
			} else {
				ans += r_max[i] - grid[i][j]
			}
		}
	}
	return ans
}
