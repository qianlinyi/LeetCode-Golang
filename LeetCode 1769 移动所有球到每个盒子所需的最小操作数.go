package main

import (
	"sort"
)

// 前缀和 + 二分
func minOperations(boxes string) []int {
	pos, ans := make([]int, 0), make([]int, len(boxes))
	for i, c := range boxes {
		if c == '1' {
			pos = append(pos, i)
		}
	}
	n := len(pos)
	sum := make([]int, n)
	if n == 0 {
		return ans
	}
	sum[0] = pos[0]
	for i := 1; i < n; i++ {
		sum[i] = pos[i] + sum[i-1]
	}
	for i := 0; i < len(boxes); i++ {
		index := sort.Search(n, func(x int) bool {
			return pos[x] >= i
		})
		if index == 0 {
			ans[i] = sum[n-1] - n*i
		} else if index == n {
			ans[i] = n*i - sum[n-1]
		} else {
			ans[i] = index*i - sum[index-1] + sum[n-1] - sum[index-1] - (n-index)*i
		}
	}
	return ans
}

func main() {
}
