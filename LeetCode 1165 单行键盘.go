package main

func calculateTime(keyboard string, word string) int {
	pos, ans, pre := make([]int, 26), 0, 0
	for i, c := range keyboard {
		pos[c-97] = i
	}
	for _, c := range word {
		if pos[c-97]-pre > 0 {
			ans += pos[c-97] - pre
		} else {
			ans += pre - pos[c-97]
		}
		pre = pos[c-97]
	}
	return ans
}
