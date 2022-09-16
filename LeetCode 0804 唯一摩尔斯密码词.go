package main

type void struct{}

func uniqueMorseRepresentations(words []string) int {
	moss := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	set := make(map[string]void)
	var member void
	for _, word := range words {
		s := ""
		for _, c := range word {
			s += moss[c-'a']
		}
		set[s] = member
	}
	return len(set)
}
