package main

func restoreString(s string, indices []int) string {
	ans := make([]byte, len(indices))
	for index, pos := range indices {
		ans[pos] = s[index]
	}
	return string(ans)
}
