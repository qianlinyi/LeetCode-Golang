package main

func wateringPlants(plants []int, capacity int) int {
	ans, cur := 0, capacity
	for idx, plant := range plants {
		if cur >= plant {
			cur -= plant
			ans++
		} else {
			ans += 2*(1+idx) - 1
			cur = capacity - plant
		}
	}
	return ans
}
