package main

func makePalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}
	half, cnt := len(s)/2, 0
	for i := 0; i < half; i++ {
		if s[i] != s[len(s)-i-1] {
			cnt++
		}
	}
	return cnt <= 2
}

func main() {
}
