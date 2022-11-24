package main

import "strings"

func reversePrefix(word string, ch byte) string {
	right := strings.IndexByte(word, ch)
	if right < 0 {
		return word
	}
	s := []byte(word)
	for left := 0; left < right; left++ {
		s[left], s[right] = s[right], s[left]
		right--
	}
	return string(s)
}
