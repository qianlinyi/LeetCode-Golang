package main

func judgeCircle(moves string) bool {
	r, l, u, d := 0, 0, 0, 0
	for _, move := range moves {
		if move == 'R' {
			r++
		} else if move == 'L' {
			l++
		} else if move == 'U' {
			u++
		} else if move == 'D' {
			d++
		}
	}
	return (l == r) && (u == d)
}
