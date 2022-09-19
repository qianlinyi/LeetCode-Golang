package main

func decodeMessage(key string, message string) string {
	mp, index, ans := make(map[int32]int), 0, ""
	for _, c := range key {
		if c == ' ' {
			continue
		}
		_, ok := mp[c]
		if ok {
			continue
		} else {
			mp[c] = index
			index++
		}
	}
	for _, c := range message {
		if c != ' ' {
			ans += string(rune('a' + mp[c]))
		} else {
			ans += string(c)
		}
	}
	return ans
}
