package main

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i, num := range nums {
		if p, ok := m[num]; ok {
			if i-p <= k {
				return true
			}
		}
		m[num] = i
	}
	return false
}
