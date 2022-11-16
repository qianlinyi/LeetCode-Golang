package main

func firstPalindrome(words []string) string {
next:
	for _, word := range words {
		for i, n := 0, len(word); i < n/2; i++ {
			if word[i] != word[n-1-i] {
				continue next
			}
		}
		return word
	}
	return ""
}
