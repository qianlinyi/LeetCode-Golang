package main

import "strings"

func numberOfBeams(bank []string) int {
	ans, base := 0, -1
	for _, str := range bank {
		cnt := strings.Count(str, "1")
		if cnt == 0 {
			continue
		}
		if base == -1 {
			base = strings.Count(str, "1")
		} else {
			ans += cnt * base
			base = cnt
		}
	}
	return ans
}
