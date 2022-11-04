package main

func findTheWinner(n int, k int) int {
	if n == 1 {
		return 1
	}
	loser, cnt := make([]bool, n), 0
	count := 1
	for i := 0; ; i++ {
		if loser[i%n] == true {
			continue
		}
		if count%k == 0 {
			loser[i%n] = true
			cnt++
			count = 0
		}
		count++
		if cnt == n-1 {
			break
		}
	}
	for idx, status := range loser {
		if status == false {
			return idx + 1
		}
	}
	return 0
}
