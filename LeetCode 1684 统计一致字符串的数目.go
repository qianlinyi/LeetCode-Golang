package main

type void struct{}

func countConsistentStrings(allowed string, words []string) int {
	origin_set, ans := make(map[int32]void), 0
	var exist void
	for _, c := range allowed {
		origin_set[c] = exist
	}
	for _, str := range words {
		flag := true
		for _, c := range str {
			_, ok := origin_set[c]
			if !ok {
				flag = false
				break
			}
		}
		if flag {
			ans++
		}
	}
	return ans
}
