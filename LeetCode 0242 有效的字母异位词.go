package main

func isAnagram(s string, t string) bool {
	cnt1, cnt2 := map[byte]int{}, map[byte]int{}
	l1, l2 := len(s), len(t)
	if l1 != l2 {
		return false
	}
	for i := 0; i < l1; i++ {
		cnt1[s[i]]++
	}
	for i := 0; i < l2; i++ {
		cnt2[t[i]]++
	}
	for k, v := range cnt1 {
		if cnt2[k] != v {
			return false
		}
	}
	for k, v := range cnt2 {
		if cnt1[k] != v {
			return false
		}
	}
	return true
}
