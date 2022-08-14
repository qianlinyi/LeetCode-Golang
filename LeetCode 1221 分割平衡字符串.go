package main

func balancedStringSplit(s string) int {
	cntl, cntr, ans := 0, 0, 0
	for _, c := range s {
		if c == 'L' {
			cntl++
		} else {
			cntr++
		}
		if cntl == cntr {
			ans++
		}
	}
	return ans
}

func main() {
	balancedStringSplit("123")
}
