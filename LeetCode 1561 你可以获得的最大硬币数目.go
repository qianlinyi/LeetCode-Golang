package main

import "sort"

func maxCoins(piles []int) int {
	sort.Ints(piles)
	n := len(piles)
	ans := 0
	for i, j := n-2, 0; j < n/3; i, j = i-2, j+1 {
		ans += piles[i]
	}
	return ans
}
