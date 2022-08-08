package main

import "sort"

func minimumTime(jobs []int, workers []int) int {
	sort.Ints(jobs)
	sort.Ints(workers)
	ans := 0
	for i := 0; i < len(jobs); i++ {
		if (jobs[i]+workers[i]-1)/workers[i] > ans {
			ans = (jobs[i] + workers[i] - 1) / workers[i]
		}
	}
	return ans
}
