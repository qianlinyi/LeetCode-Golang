package main

func digitCount(num string) bool {
	cnt := make(map[int32]int32)
	for _, c := range num {
		cnt[c]++
	}
	for idx, c := range num {
		if cnt[int32(idx)+'0'] != c-'0' {
			return false
		}
	}
	return true
}
