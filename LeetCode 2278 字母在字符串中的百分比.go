package main

func percentageLetter(s string, letter byte) int {
	cnt := 0
	for _, c := range s {
		if byte(c) == letter {
			cnt++
		}
	}
	return cnt * 100 / len(s)
}
