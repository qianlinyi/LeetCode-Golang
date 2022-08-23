package main

import "strings"

func countAsterisks(s string) int {
	ans, list := 0, strings.Split(s, "|")
	for index, s := range list {
		if index%2 == 0 {
			ans += strings.Count(s, "*")
		}
	}
	return ans
}
