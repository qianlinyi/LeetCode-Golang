package main

func squareIsWhite(coordinates string) bool {
	add := 0
	for _, c := range coordinates {
		if '1' <= c && c <= '9' {
			add += int(c - '0')
		} else {
			add += int(c - 'a' + 1)
		}
	}
	return add%2 == 1
}
