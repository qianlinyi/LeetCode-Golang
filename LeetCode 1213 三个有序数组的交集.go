package main

import "sort"

func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {
	cnt, arr := make(map[int]int, 0), append(append(arr1, arr2...), arr3...)
	for _, num := range arr {
		cnt[num]++
	}
	ans := make([]int, 0)
	for k, v := range cnt {
		if v == 3 {
			ans = append(ans, k)
		}
	}
	sort.Ints(ans)
	return ans
}
