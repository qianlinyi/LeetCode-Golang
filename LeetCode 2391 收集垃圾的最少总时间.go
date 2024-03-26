package main

import (
	"strings"
)

func garbageCollection(garbage []string, travel []int) int {
	kind, time := []string{"M", "G", "P"}, 0
	for _, k := range kind {
		sum, cost := 0, 0
		for i, g := range garbage {
			cnt := strings.Count(g, k)
			if cnt != 0 {
				sum += cost + cnt
				cost = 0
			}
			if i < len(travel) {
				cost += travel[i]
			}
		}
		time += sum
	}
	return time
}
