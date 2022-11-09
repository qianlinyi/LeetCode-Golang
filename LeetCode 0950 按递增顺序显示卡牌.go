package main

import "sort"

func deckRevealedIncreasing(deck []int) []int {
	ans, index := make([]int, len(deck)), make([]int, 0)
	sort.Ints(deck)
	for i := 0; i < len(deck); i++ {
		index = append(index, i)
	}
	for _, card := range deck {
		ans[index[0]] = card
		index = index[1:]
		if len(index) != 0 {
			index = append(index, index[0])
			index = index[1:]
		}
	}
	return ans
}
