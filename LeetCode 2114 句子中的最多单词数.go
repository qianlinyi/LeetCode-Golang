package main

import (
	"strings"
)

func mostWordsFound(sentences []string) int {
	ans := 0
	for _, sentence := range sentences {
		//wordlist := strings.Split(sentence, " ")
		//if ans < len(wordlist) {
		//	ans = len(wordlist)
		//}
		cnt := strings.Count(sentence, " ") + 1
		if ans < cnt {
			ans = cnt
		}
	}
	return ans
}

func main() {}
