package main

func numberOfMatches(n int) int {
	cnt := 0
	for n > 1 {
		if n%2 == 0 {
			cnt += n / 2
			n = n / 2
		} else {
			cnt += (n - 1) / 2
			n = (n + 1) / 2
		}
	}
	return cnt
}
