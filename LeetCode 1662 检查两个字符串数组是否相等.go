package main

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	s1, s2 := "", ""
	for _, c := range word1 {
		s1 += c
	}
	for _, c := range word2 {
		s2 += c
	}
	return s1 == s2
}
