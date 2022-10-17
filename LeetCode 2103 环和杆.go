package main

// 位运算
func countPoints(rings string) int {
	ans, status := 0, make([]int, 10)
	bits := []int{'R': 1, 'G': 2, 'B': 4}
	for i, c := range rings {
		if 48 <= c && c <= 57 {
			status[c-'0'] |= bits[rings[i-1]]
		}
	}
	for i := 0; i < 10; i++ {
		if status[i] == 7 {
			ans++
		}
	}
	return ans
}
