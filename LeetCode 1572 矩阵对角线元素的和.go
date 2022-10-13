package main

func diagonalSum(mat [][]int) int {
	ans := 0
	for i, row := range mat {
		for j, num := range row {
			if i == j {
				ans += num
			} else if i+j == len(row)-1 {
				ans += num
			}
		}
	}
	return ans
}
