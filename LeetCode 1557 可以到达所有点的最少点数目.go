package main

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	in, ans := make([]int, n), make([]int, 0)
	for _, edge := range edges {
		in[edge[1]]++
	}
	for i := 0; i < n; i++ {
		if in[i] == 0 {
			ans = append(ans, i)
		}
	}
	return ans
}
