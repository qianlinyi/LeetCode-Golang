package main

import "strings"

func numOfStrings(patterns []string, word string) int {
	ans := 0
	for _, s := range patterns {
		if strings.Contains(word, s) == true {
			ans++
		}
	}
	return ans
}
