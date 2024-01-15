package main

func lengthOfLongestSubstring(s string) int {
	cnt := make([]int, 130)
	l, r, ans := 0, 0, 0
	for r < len(s) {
		cnt[s[r]]++
		for cnt[s[r]] > 1 {
			cnt[s[l]]--
			l++
		}
		ans = max(ans, r-l+1)
		r++
	}
	return ans
}
