package main

import "strings"

func countPrefixes(words []string, s string) int {
	ans := 0
	for _, word := range words {
		if strings.Index(s, word) == 0 {
			ans++
		}
	}
	return ans
}
