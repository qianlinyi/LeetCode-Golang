package main

func removeVowels(s string) string {
	ans := ""
	for _, c := range s {
		if c != 'a' && c != 'e' && c != 'i' && c != 'o' && c != 'u' {
			ans += string(c)
		}
	}
	return ans
}

func main() {}
