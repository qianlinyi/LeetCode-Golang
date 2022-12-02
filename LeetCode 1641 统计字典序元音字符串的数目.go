package main

func countVowelStrings(n int) int {
	ans := [51][5]int{}
	for i := 0; i < 5; i++ {
		ans[1][i] = 1
	}
	for i := 2; i <= n; i++ {
		ans[i][0] = ans[i-1][0]
		ans[i][1] = ans[i-1][0] + ans[i-1][1]
		ans[i][2] = ans[i-1][0] + ans[i-1][1] + ans[i-1][2]
		ans[i][3] = ans[i-1][0] + ans[i-1][1] + ans[i-1][2] + ans[i-1][3]
		ans[i][4] = ans[i-1][0] + ans[i-1][1] + ans[i-1][2] + ans[i-1][3] + ans[i-1][4]
	}
	return ans[n][0] + ans[n][1] + ans[n][2] + ans[n][3] + ans[n][4]
}
